package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHelloGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	Hello(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestHelloMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()
	Hello(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}

func TestSecretGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/secret", nil)
	rr := httptest.NewRecorder()
	Secret(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestSecretMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/secret", nil)
	rr := httptest.NewRecorder()
	Secret(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}

func TestConfigMapMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/configmap", nil)
	rr := httptest.NewRecorder()
	ConfigMap(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}

func TestConfigMapFileNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/configmap", nil)
	rr := httptest.NewRecorder()
	ConfigMap(rr, req)
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected 500 when file missing, got %d", rr.Code)
	}
}

func TestHealthzBefore10s(t *testing.T) {
	startedAt = time.Now()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()
	Healthz(rr, req)
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected 500 before 10s, got %d", rr.Code)
	}
}

func TestHealthzAfter10s(t *testing.T) {
	startedAt = time.Now().Add(-11 * time.Second)
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()
	Healthz(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 after 10s, got %d", rr.Code)
	}
}

func TestHealthzMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/healthz", nil)
	rr := httptest.NewRecorder()
	Healthz(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}
