package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"go-server/models"
	"go-server/service"
)

const maxBodyBytes = 1 << 20

// Handler knows about HTTP: verbs, status codes and JSON. It calls only the service.
type Handler struct {
	svc     *service.Service
	baseURL string
}

// New builds a Handler. baseURL is used to build the short_url in responses.
func New(svc *service.Service, baseURL string) *Handler {
	return &Handler{svc: svc, baseURL: baseURL}
}

// Routes returns the HTTP routes. The Go 1.22 mux handles method matching and path values.
func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /{$}", h.shorten)
	mux.HandleFunc("GET /healthz", h.healthz)
	mux.HandleFunc("GET /api/urls/{code}", h.getURL)
	mux.HandleFunc("DELETE /api/urls/{code}", h.deleteURL)
	mux.HandleFunc("GET /{code}", h.redirect)
	return mux
}

type shortenRequest struct {
	URL  string `json:"url"`
	Code string `json:"code"`
}

type shortenResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) shorten(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid JSON body"})
		return
	}

	u, err := h.svc.Shorten(r.Context(), req.URL, req.Code)
	if err != nil {
		h.writeError(w, r, err)
		return
	}

	writeJSON(w, http.StatusCreated, shortenResponse{
		Code:     u.Code,
		ShortURL: h.baseURL + "/" + u.Code,
	})
}

func (h *Handler) redirect(w http.ResponseWriter, r *http.Request) {
	target, err := h.svc.Resolve(r.Context(), r.PathValue("code"))
	if err != nil {
		h.writeError(w, r, err)
		return
	}

	http.Redirect(w, r, target, http.StatusMovedPermanently)
}

func (h *Handler) getURL(w http.ResponseWriter, r *http.Request) {
	u, err := h.svc.Get(r.Context(), r.PathValue("code"))
	if err != nil {
		h.writeError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, u)
}

func (h *Handler) deleteURL(w http.ResponseWriter, r *http.Request) {
	if err := h.svc.Delete(r.Context(), r.PathValue("code")); err != nil {
		h.writeError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// healthz does not touch the database. It answers whether the process is alive.
func (h *Handler) healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// writeError maps service errors to status codes. Unknown errors are logged and hidden from the client.
func (h *Handler) writeError(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.Is(err, models.ErrInvalidURL), errors.Is(err, models.ErrInvalidCode):
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
	case errors.Is(err, models.ErrCodeTaken):
		writeJSON(w, http.StatusConflict, errorResponse{Error: "that code is already taken"})
	case errors.Is(err, models.ErrNotFound):
		writeJSON(w, http.StatusNotFound, errorResponse{Error: "short url not found"})
	default:
		log.Printf("%s %s: %v", r.Method, r.URL.Path, err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal error"})
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("writing response: %v", err)
	}
}
