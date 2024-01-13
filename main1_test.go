package main

import (
	"ascii-art-web/pkg/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	server.PathHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}
