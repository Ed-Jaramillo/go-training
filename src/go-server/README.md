# URL shortener

An HTTP service in Go that stores short codes in PostgreSQL.

Status: builds, `go vet` is clean, and unit tests pass with `-race`. The Postgres repository compiles but has not yet been run against a live database.

## Run it

Create a database, apply the schema, then start the server.

```bash
createdb shortener
```

```bash
psql "postgres://user:pass@localhost:5432/shortener?sslmode=disable" -f schema.sql
```

```bash
DATABASE_DSN="postgres://user:pass@localhost:5432/shortener?sslmode=disable" go run .
```

The server stops cleanly on Ctrl+C (SIGINT) or SIGTERM. It stops accepting connections, waits up to 10 seconds for in-flight requests, then closes the database pool.

## Environment variables

| Variable       | Required | Default                    | Purpose                                              |
|----------------|----------|----------------------------|------------------------------------------------------|
| `DATABASE_DSN` | yes      | none                       | PostgreSQL connection string (pgx format)            |
| `PORT`         | no       | `8080`                     | Port the HTTP server listens on                      |
| `BASE_URL`     | no       | `http://localhost:$PORT`   | Prefix for `short_url` in responses. Trailing `/` is trimmed. |

The schema is not applied automatically. `schema.sql` is run by hand.

## API

| Method   | Path               | Result                                                                 |
|----------|--------------------|------------------------------------------------------------------------|
| `POST`   | `/`                | `201` with `{"code": "...", "short_url": "..."}`                       |
| `GET`    | `/{code}`          | `301` redirect to the target. `404` if the code does not exist.        |
| `GET`    | `/api/urls/{code}` | `200` with the record. `404` if the code does not exist.               |
| `DELETE` | `/api/urls/{code}` | `204` on success. `404` if the code does not exist.                    |
| `GET`    | `/healthz`         | `200`. Does not query the database.                                    |

Create a short URL with a generated code:

```bash
curl -X POST localhost:8080/ -d '{"url":"https://example.com/some/long/page"}'
```

Create one with a vanity code:

```bash
curl -X POST localhost:8080/ -d '{"url":"https://example.com/docs","code":"docs"}'
```

Errors come back as JSON: `{"error": "..."}`. Status codes:

- `400`: invalid JSON, a URL that is not `http`/`https` with a host, or a malformed vanity code.
- `409`: the vanity code is already taken.
- `500`: anything unexpected. The detail is written to the server log, not sent to the client.

## Codes

- Generated codes are 8 characters from `a-z`, `A-Z`, `0-9`. Each character is drawn with `crypto/rand.Int`. A plain random byte modulo 62 would favour the first eight characters, because 256 is not a multiple of 62.
- A generated code that collides with an existing one is retried, up to 5 attempts. If all 5 collide, the request returns `500`.
- A vanity code must be 3 to 32 characters from `A-Z`, `a-z`, `0-9`, `-` and `_`. `healthz` is reserved because it is a route.
- A vanity code that already exists returns `409` and is never overwritten.

## Hit counting: synchronous

Each redirect runs one statement, `UPDATE urls SET hit_count = hit_count + 1, last_accessed_at = now() WHERE code = $1 RETURNING target_url`, inside the request. That statement both finds the target and records the hit, so the redirect needs no second query.

The alternative is to push the hit onto a channel and let a background goroutine write it. That keeps writes off the redirect path, which lowers latency under load. It also means buffered hits can be lost if the process crashes, and shutdown has to drain the channel before closing the pool.

This service writes synchronously. The write is cheap and the counts stay exact for every request that completes. Shutdown is safe because `http.Server.Shutdown` waits for in-flight requests, and each of them finishes its update before the pool closes. The cost is one database write per redirect. Revisit this if redirect latency becomes a problem.

Two consequences to know about:

- A `HEAD` request to `/{code}` also counts as a hit, because Go's `GET` pattern matches `HEAD`.
- Browsers cache `301` responses, so a repeat visit may never reach the server. Hit counts understate real traffic.

## ORM: none

The service uses `database/sql` with the `pgx` driver (`github.com/jackc/pgx/v5/stdlib`). The repository has five SQL statements. None is complex enough to justify an ORM, and an ORM would hide the single `UPDATE ... RETURNING` that makes the hit count cheap. I did not compare any ORM's generated SQL against these, so this is a judgment from the statement count, not a measured result.

## Layout

- `main.go`: reads configuration, opens the pool, starts the server, handles shutdown.
- `handlers/`: HTTP only. Decodes JSON, maps errors to status codes, and writes responses.
- `service/`: the rules (URL and vanity validation, code generation with retries, what counts as a hit). It depends on the `Store` interface, not on Postgres.
- `storage/`: the Postgres repository. It is the only code that contains SQL.
- `models/`: the `ShortURL` struct, the code generator, and the shared error values.
- `schema.sql`: the table definition.

## Tests

```bash
go test -race ./...
```

- `models`: generator length, alphabet and uniqueness over 10,000 runs.
- `service`: validation rules, vanity collisions, generated-code retries, hit counting, not-found paths. Uses an in-memory fake `Store`.
- `handlers`: every route and status code, run against an in-memory `Store`.

`storage` has no tests. They would need a running PostgreSQL instance.

## Known limitations

- No authentication, rate limiting or URL allow-listing. Anyone who can reach the server can create links.
- No migrations. Schema changes are made by editing `schema.sql` and applying it by hand.
- The Postgres repository has not been exercised against a live database yet.
- `created_at` is set by the application, not by a database default.
