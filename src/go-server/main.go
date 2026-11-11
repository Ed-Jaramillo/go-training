package main

/*Here I am going to write the instructions for my assignment
I need to create an URL shorterner in Go.


An HTTP service backed by a SQL database. Small enough to finish in a
couple of sittings, and broad enough that once you have it working you have the shape or nearly
every Go service you will ever write.

This is the base and it will be followed by more tasks like
  * A CLI against the same database
  * Concurrency
  * Background work
  * Caching
  * Middleware
  * gRPC

  POST: Takes {"url": "https://…"}
  GET:  Looks up the code and redirects 301 to the target. 404 if
  there's no such code
  GET:  Returns the record as JSON- target, created time, hit count, last accessed
  DELETE: Removes it. 204 on success, 404 if it wasn't there
  GET:  Returns 200. No database call -- this one answers whether the process is alive, not whether the database is

  Generated codes are eight random alphanumeric characters unless you have a reason to
  prefer something else. Vanity slugs will collide eventually, so decide what that response looks like.


  This project done looks like

  * A server binary that starts, and shuts down cleanly on Ctrl+C
  * A schema.sql in the repo
  * A README.md how to run it, what env vars it needs, the sync-versus-background
  note on hit counting, and the ORM note if that applies
  * go build ./... and go vet ./... clean

  Once this is standing up, the follow-on exercises attach to it directly, and by
  then mos of the Go you need for real work will already be under your hands.
*/

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"go-server/handlers"
	"go-server/service"
	"go-server/storage"
)

const (
	shutdownTimeout = 10 * time.Second
	pingTimeout     = 5 * time.Second
)

func main() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN is required, e.g. postgres://user:pass@localhost:5432/shortener?sslmode=disable")
	}
	port := getenv("PORT", "8080")
	baseURL := strings.TrimRight(getenv("BASE_URL", "http://localhost:"+port), "/")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("opening database: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pingCtx, cancel := context.WithTimeout(ctx, pingTimeout)
	err = db.PingContext(pingCtx)
	cancel()
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}

	svc := service.New(storage.NewPostgres(db))
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           handlers.New(svc, baseURL).Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- srv.ListenAndServe()
	}()
	log.Printf("listening on %s, short URLs use %s", srv.Addr, baseURL)

	select {
	case err := <-serverErr:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("serving: %v", err)
		}
	case <-ctx.Done():
		log.Print("shutdown signal received, draining requests")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("shutdown: %v", err)
		}
		cancel()
	}

	if err := db.Close(); err != nil {
		log.Printf("closing database: %v", err)
	}
	log.Print("stopped")
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
