package service

import (
	"context"
	"errors"
	"sync"
	"testing"

	"go-server/models"
)

// fakeStore is an in-memory Store. takenFor makes Insert report a collision for the first n calls.
type fakeStore struct {
	mu        sync.Mutex
	rows      map[string]models.ShortURL
	collideN  int
	insertErr error
}

func newFakeStore() *fakeStore {
	return &fakeStore{rows: make(map[string]models.ShortURL)}
}

func (f *fakeStore) Insert(_ context.Context, u models.ShortURL) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.insertErr != nil {
		return f.insertErr
	}
	if f.collideN > 0 {
		f.collideN--
		return models.ErrCodeTaken
	}
	if _, ok := f.rows[u.Code]; ok {
		return models.ErrCodeTaken
	}
	f.rows[u.Code] = u
	return nil
}

func (f *fakeStore) ByCode(_ context.Context, code string) (models.ShortURL, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	u, ok := f.rows[code]
	if !ok {
		return models.ShortURL{}, models.ErrNotFound
	}
	return u, nil
}

func (f *fakeStore) RecordHit(_ context.Context, code string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	u, ok := f.rows[code]
	if !ok {
		return "", models.ErrNotFound
	}
	u.HitCount++
	f.rows[code] = u
	return u.Target, nil
}

func (f *fakeStore) Delete(_ context.Context, code string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.rows[code]; !ok {
		return models.ErrNotFound
	}
	delete(f.rows, code)
	return nil
}

func TestShorten(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		vanity    string
		seed      []string // codes already in the store
		collideN  int
		insertErr error
		wantErr   error
		wantCode  string // exact code, only for vanity cases
	}{
		{name: "generated code", target: "https://example.com/a"},
		{name: "vanity code", target: "https://example.com/a", vanity: "my-link", wantCode: "my-link"},
		{name: "vanity already taken", target: "https://example.com/a", vanity: "taken", seed: []string{"taken"}, wantErr: models.ErrCodeTaken},
		{name: "generated collision is retried", target: "https://example.com/a", collideN: 2},
		{name: "generated collisions exhaust attempts", target: "https://example.com/a", collideN: maxGenerateAttempts, wantErr: errExhausted},
		{name: "missing scheme", target: "example.com/a", wantErr: models.ErrInvalidURL},
		{name: "javascript scheme", target: "javascript:alert(1)", wantErr: models.ErrInvalidURL},
		{name: "missing host", target: "https://", wantErr: models.ErrInvalidURL},
		{name: "too long", target: "https://example.com/" + string(make([]byte, maxURLLength)), wantErr: models.ErrInvalidURL},
		{name: "vanity too short", target: "https://example.com/a", vanity: "ab", wantErr: models.ErrInvalidCode},
		{name: "vanity with space", target: "https://example.com/a", vanity: "a b c", wantErr: models.ErrInvalidCode},
		{name: "reserved vanity", target: "https://example.com/a", vanity: "healthz", wantErr: models.ErrInvalidCode},
		{name: "store failure is not a client error", target: "https://example.com/a", insertErr: errors.New("db down"), wantErr: errStore},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := newFakeStore()
			store.collideN = tt.collideN
			store.insertErr = tt.insertErr
			for _, code := range tt.seed {
				store.rows[code] = models.ShortURL{Code: code}
			}
			svc := New(store)

			u, err := svc.Shorten(context.Background(), tt.target, tt.vanity)

			switch {
			case tt.wantErr == errExhausted, tt.wantErr == errStore:
				if err == nil || errors.Is(err, models.ErrCodeTaken) {
					t.Fatalf("Shorten() error = %v, want a non-client error", err)
				}
				return
			case tt.wantErr != nil:
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Shorten() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Shorten() unexpected error: %v", err)
			}
			if u.Target != tt.target {
				t.Errorf("Target = %q, want %q", u.Target, tt.target)
			}
			if tt.wantCode != "" && u.Code != tt.wantCode {
				t.Errorf("Code = %q, want %q", u.Code, tt.wantCode)
			}
			if _, ok := store.rows[u.Code]; !ok {
				t.Errorf("code %q was not saved", u.Code)
			}
		})
	}
}

// errExhausted and errStore are sentinels for table rows whose failure is not one of the models errors.
var (
	errExhausted = errors.New("exhausted")
	errStore     = errors.New("store")
)

func TestResolveCountsHit(t *testing.T) {
	store := newFakeStore()
	store.rows["abc"] = models.ShortURL{Code: "abc", Target: "https://example.com"}
	svc := New(store)

	target, err := svc.Resolve(context.Background(), "abc")
	if err != nil {
		t.Fatalf("Resolve() error: %v", err)
	}
	if target != "https://example.com" {
		t.Errorf("Resolve() = %q, want %q", target, "https://example.com")
	}
	if got := store.rows["abc"].HitCount; got != 1 {
		t.Errorf("HitCount = %d, want 1", got)
	}
}

func TestNotFound(t *testing.T) {
	svc := New(newFakeStore())
	ctx := context.Background()

	if _, err := svc.Resolve(ctx, "nope"); !errors.Is(err, models.ErrNotFound) {
		t.Errorf("Resolve() error = %v, want ErrNotFound", err)
	}
	if _, err := svc.Get(ctx, "nope"); !errors.Is(err, models.ErrNotFound) {
		t.Errorf("Get() error = %v, want ErrNotFound", err)
	}
	if err := svc.Delete(ctx, "nope"); !errors.Is(err, models.ErrNotFound) {
		t.Errorf("Delete() error = %v, want ErrNotFound", err)
	}
}
