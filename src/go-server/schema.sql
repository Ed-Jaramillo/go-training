-- Schema for the URL shortener. Apply it by hand; the service never runs DDL.
--   psql "$DATABASE_DSN" -f schema.sql

CREATE TABLE IF NOT EXISTS urls (
    code             TEXT        PRIMARY KEY,
    target_url       TEXT        NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL,
    hit_count        BIGINT      NOT NULL DEFAULT 0 CHECK (hit_count >= 0),
    last_accessed_at TIMESTAMPTZ
);
