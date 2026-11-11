package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"go-server/models"
	"go-server/service"
)

// memStore is an in-memory service.Store for exercising the routes without a database.
type memStore struct {
	mu   sync.Mutex
	rows map[string]models.ShortURL
}

func (m *memStore) Insert(_ context.Context, u models.ShortURL) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.rows[u.Code]; ok {
		return models.ErrCodeTaken
	}
	m.rows[u.Code] = u
	return nil
}

func (m *memStore) ByCode(_ context.Context, code string) (models.ShortURL, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	u, ok := m.rows[code]
	if !ok {
		return models.ShortURL{}, models.ErrNotFound
	}
	return u, nil
}

func (m *memStore) RecordHit(_ context.Context, code string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	u, ok := m.rows[code]
	if !ok {
		return "", models.ErrNotFound
	}
	u.HitCount++
	m.rows[code] = u
	return u.Target, nil
}

func (m *memStore) Delete(_ context.Context, code string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.rows[code]; !ok {
		return models.ErrNotFound
	}
	delete(m.rows, code)
	return nil
}

func newTestServer(t *testing.T) http.Handler {
	t.Helper()
	store := &memStore{rows: make(map[string]models.ShortURL)}
	return New(service.New(store), "http://short.test").Routes()
}

func do(t *testing.T, h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestRoutes(t *testing.T) {
	h := newTestServer(t)

	// Create a vanity link so the later steps have a known code.
	created := do(t, h, http.MethodPost, "/", `{"url":"https://example.com/page","code":"docs"}`)
	if created.Code != http.StatusCreated {
		t.Fatalf("POST / status = %d, body %s", created.Code, created.Body)
	}
	var body shortenResponse
	if err := json.Unmarshal(created.Body.Bytes(), &body); err != nil {
		t.Fatalf("decoding POST / body: %v", err)
	}
	if body.Code != "docs" || body.ShortURL != "http://short.test/docs" {
		t.Errorf("POST / body = %+v", body)
	}

	tests := []struct {
		name       string
		method     string
		path       string
		body       string
		wantStatus int
		wantLoc    string
	}{
		{name: "redirect is 301", method: http.MethodGet, path: "/docs", wantStatus: http.StatusMovedPermanently, wantLoc: "https://example.com/page"},
		{name: "unknown code is 404", method: http.MethodGet, path: "/nope", wantStatus: http.StatusNotFound},
		{name: "duplicate vanity is 409", method: http.MethodPost, path: "/", body: `{"url":"https://example.com/b","code":"docs"}`, wantStatus: http.StatusConflict},
		{name: "bad url is 400", method: http.MethodPost, path: "/", body: `{"url":"javascript:alert(1)"}`, wantStatus: http.StatusBadRequest},
		{name: "bad JSON is 400", method: http.MethodPost, path: "/", body: `{"url":`, wantStatus: http.StatusBadRequest},
		{name: "get record", method: http.MethodGet, path: "/api/urls/docs", wantStatus: http.StatusOK},
		{name: "health check", method: http.MethodGet, path: "/healthz", wantStatus: http.StatusOK},
		{name: "delete existing is 204", method: http.MethodDelete, path: "/api/urls/docs", wantStatus: http.StatusNoContent},
		{name: "delete missing is 404", method: http.MethodDelete, path: "/api/urls/docs", wantStatus: http.StatusNotFound},
		{name: "after delete redirect is 404", method: http.MethodGet, path: "/docs", wantStatus: http.StatusNotFound},
		{name: "GET on POST-only root is 405", method: http.MethodGet, path: "/", wantStatus: http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := do(t, h, tt.method, tt.path, tt.body)
			if rec.Code != tt.wantStatus {
				t.Fatalf("%s %s status = %d, want %d, body %s", tt.method, tt.path, rec.Code, tt.wantStatus, rec.Body)
			}
			if tt.wantLoc != "" && rec.Header().Get("Location") != tt.wantLoc {
				t.Errorf("Location = %q, want %q", rec.Header().Get("Location"), tt.wantLoc)
			}
		})
	}
}

func TestWriteErrorHidesInternalDetail(t *testing.T) {
	h := &Handler{}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/x", nil)

	h.writeError(rec, req, errors.New("connection refused to 10.0.0.5"))

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "10.0.0.5") {
		t.Errorf("internal error leaked to client: %s", rec.Body)
	}
}
