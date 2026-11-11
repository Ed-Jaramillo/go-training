package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"time"

	"go-server/models"
)

// Store is what the service needs from persistence. It is defined here, next to its
// only user, and the Postgres repository satisfies it.
type Store interface {
	Insert(ctx context.Context, u models.ShortURL) error
	ByCode(ctx context.Context, code string) (models.ShortURL, error)
	RecordHit(ctx context.Context, code string) (string, error)
	Delete(ctx context.Context, code string) error
}

const (
	maxGenerateAttempts = 5
	maxURLLength        = 2048
)

var vanityPattern = regexp.MustCompile(`^[A-Za-z0-9_-]{3,32}$`)

// reservedCodes are routes, so no vanity slug may use them.
var reservedCodes = map[string]bool{
	"healthz": true,
}

// Service holds the rules: how codes are made, what counts as a valid URL or slug,
// and whether a redirect is a hit.
type Service struct {
	store Store
}

// New builds a Service on top of a Store.
func New(store Store) *Service {
	return &Service{store: store}
}

// Shorten stores target under vanity if one is given, or under a generated code otherwise.
// A taken vanity slug returns models.ErrCodeTaken. A generated collision is retried.
func (s *Service) Shorten(ctx context.Context, target, vanity string) (models.ShortURL, error) {
	if err := validateTarget(target); err != nil {
		return models.ShortURL{}, err
	}

	if vanity != "" {
		if err := validateVanity(vanity); err != nil {
			return models.ShortURL{}, err
		}
		u := newRecord(vanity, target)
		if err := s.store.Insert(ctx, u); err != nil {
			return models.ShortURL{}, fmt.Errorf("saving vanity slug: %w", err)
		}
		return u, nil
	}

	for attempt := 1; attempt <= maxGenerateAttempts; attempt++ {
		code, err := models.GenerateShortURL()
		if err != nil {
			return models.ShortURL{}, fmt.Errorf("generating code: %w", err)
		}
		u := newRecord(code, target)
		err = s.store.Insert(ctx, u)
		if err == nil {
			return u, nil
		}
		if !errors.Is(err, models.ErrCodeTaken) {
			return models.ShortURL{}, fmt.Errorf("saving generated code: %w", err)
		}
	}

	// Not wrapping ErrCodeTaken on purpose: the client did not pick a slug, so this is a server error.
	return models.ShortURL{}, fmt.Errorf("no free code after %d attempts", maxGenerateAttempts)
}

// Resolve returns the target for code and counts the lookup as a hit.
func (s *Service) Resolve(ctx context.Context, code string) (string, error) {
	return s.store.RecordHit(ctx, code)
}

// Get returns the record for code without counting a hit.
func (s *Service) Get(ctx context.Context, code string) (models.ShortURL, error) {
	return s.store.ByCode(ctx, code)
}

// Delete removes code, or returns models.ErrNotFound.
func (s *Service) Delete(ctx context.Context, code string) error {
	return s.store.Delete(ctx, code)
}

func newRecord(code, target string) models.ShortURL {
	return models.ShortURL{
		Code:      code,
		Target:    target,
		CreatedAt: time.Now().UTC(),
	}
}

func validateTarget(target string) error {
	if len(target) > maxURLLength {
		return fmt.Errorf("%w: longer than %d characters", models.ErrInvalidURL, maxURLLength)
	}
	u, err := url.Parse(target)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErrInvalidURL, err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("%w: scheme must be http or https", models.ErrInvalidURL)
	}
	if u.Host == "" {
		return fmt.Errorf("%w: missing host", models.ErrInvalidURL)
	}
	return nil
}

func validateVanity(code string) error {
	if !vanityPattern.MatchString(code) {
		return fmt.Errorf("%w: use 3-32 letters, digits, '-' or '_'", models.ErrInvalidCode)
	}
	if reservedCodes[code] {
		return fmt.Errorf("%w: %q is reserved", models.ErrInvalidCode, code)
	}
	return nil
}
