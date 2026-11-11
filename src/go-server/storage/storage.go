package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"

	"go-server/models"
)

const uniqueViolation = "23505"

// Postgres is the only code that speaks SQL. Every query lives here.
type Postgres struct {
	db *sql.DB
}

// NewPostgres wraps an open connection pool.
func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

// Insert saves a new row. A primary-key collision returns models.ErrCodeTaken.
func (p *Postgres) Insert(ctx context.Context, u models.ShortURL) error {
	_, err := p.db.ExecContext(ctx,
		`INSERT INTO urls (code, target_url, created_at) VALUES ($1, $2, $3)`,
		u.Code, u.Target, u.CreatedAt,
	)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
		return models.ErrCodeTaken
	}
	if err != nil {
		return fmt.Errorf("inserting code %q: %w", u.Code, err)
	}
	return nil
}

// ByCode returns the full record without changing it.
func (p *Postgres) ByCode(ctx context.Context, code string) (models.ShortURL, error) {
	var (
		u    models.ShortURL
		last sql.NullTime
	)
	err := p.db.QueryRowContext(ctx,
		`SELECT code, target_url, created_at, hit_count, last_accessed_at FROM urls WHERE code = $1`,
		code,
	).Scan(&u.Code, &u.Target, &u.CreatedAt, &u.HitCount, &last)
	if errors.Is(err, sql.ErrNoRows) {
		return models.ShortURL{}, models.ErrNotFound
	}
	if err != nil {
		return models.ShortURL{}, fmt.Errorf("selecting code %q: %w", code, err)
	}
	if last.Valid {
		u.LastAccessedAt = &last.Time
	}
	return u, nil
}

// RecordHit bumps hit_count and last_accessed_at, and returns the target, in one statement.
func (p *Postgres) RecordHit(ctx context.Context, code string) (string, error) {
	var target string
	err := p.db.QueryRowContext(ctx,
		`UPDATE urls
		    SET hit_count = hit_count + 1, last_accessed_at = now()
		  WHERE code = $1
		  RETURNING target_url`,
		code,
	).Scan(&target)
	if errors.Is(err, sql.ErrNoRows) {
		return "", models.ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("recording hit for code %q: %w", code, err)
	}
	return target, nil
}

// Delete removes a row, or returns models.ErrNotFound if there was none.
func (p *Postgres) Delete(ctx context.Context, code string) error {
	res, err := p.db.ExecContext(ctx, `DELETE FROM urls WHERE code = $1`, code)
	if err != nil {
		return fmt.Errorf("deleting code %q: %w", code, err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("counting deleted rows for code %q: %w", code, err)
	}
	if n == 0 {
		return models.ErrNotFound
	}
	return nil
}
