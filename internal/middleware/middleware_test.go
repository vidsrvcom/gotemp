package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)

	handler := Logger(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	logOutput := buf.String()
	if !strings.Contains(logOutput, "GET") {
		t.Errorf("Logger() should log request method, got: %s", logOutput)
	}
	if !strings.Contains(logOutput, "/test") {
		t.Errorf("Logger() should log request path, got: %s", logOutput)
	}
}

func TestRecoverer(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)

	handler := Recoverer(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Recoverer() status = %v, want %v", rec.Code, http.StatusInternalServerError)
	}

	logOutput := buf.String()
	if !strings.Contains(logOutput, "Panic recovered") {
		t.Errorf("Recoverer() should log panic, got: %s", logOutput)
	}
}

func TestCORS(t *testing.T) {
	handler := CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Test regular request
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("CORS() should set Access-Control-Allow-Origin header")
	}

	// Test preflight request
	req = httptest.NewRequest(http.MethodOptions, "/test", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("CORS() OPTIONS status = %v, want %v", rec.Code, http.StatusOK)
	}
}

func TestContentType(t *testing.T) {
	handler := ContentType("application/json")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("ContentType() Content-Type = %v, want application/json", rec.Header().Get("Content-Type"))
	}
}
