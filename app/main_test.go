package main

import (
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	rec := httptest.NewRecorder()

	healthz(rec, req) // llama a tu handler

	if rec.Body.String() != "ok" {
		t.Errorf("esperaba ok, recibí %q", rec.Body.String())
	}
}

func TestReadyz(t *testing.T) {
	req := httptest.NewRequest("GET", "/readyz", nil)
	rec := httptest.NewRecorder()

	readyz(rec, req) // llama a tu handler

	if rec.Body.String() != "ready" {
		t.Errorf("esperaba ready, recibí %q", rec.Body.String())
	}
}
