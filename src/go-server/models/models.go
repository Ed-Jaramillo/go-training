package models

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// Errors shared across layers. The repository translates driver errors into these,
// the service decides when they happen, and the handler maps them to status codes.
var (
	ErrNotFound    = errors.New("short url not found")
	ErrCodeTaken   = errors.New("code already taken")
	ErrInvalidURL  = errors.New("invalid url")
	ErrInvalidCode = errors.New("invalid code")
)

// ShortURL is one row of the urls table.
type ShortURL struct {
	Code           string     `json:"code"`
	Target         string     `json:"target"`
	CreatedAt      time.Time  `json:"created_at"`
	HitCount       int64      `json:"hit_count"`
	LastAccessedAt *time.Time `json:"last_accessed_at"`
}

const (
	codeLength = 8
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// GenerateShortURL returns a random 8-character code.
func GenerateShortURL() (string, error) {
	return generateKey(codeLength)
}

// generateKey picks each character with crypto/rand.Int. Taking a random byte modulo 62
// would favour the first 8 characters, because 256 is not a multiple of 62.
func generateKey(length int) (string, error) {
	max := big.NewInt(int64(len(alphabet)))
	key := make([]byte, length)

	for i := range key {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", fmt.Errorf("reading random source: %w", err)
		}
		key[i] = alphabet[n.Int64()]
	}

	return string(key), nil
}
