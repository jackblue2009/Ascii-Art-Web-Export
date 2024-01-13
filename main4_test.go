package main

import (
	"io"
	"net/http"
	"os/exec"
	"testing"
)

func TestMyWebApp(t *testing.T) {
	image := "asciiartwebdockerize"
	container := exec.Command("docker", "run", image)
	err := container.Start()
	if err != nil {
		t.Errorf("Failed to start Docker image: %v", err)
	}
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Errorf("Failed to make HTTP request: %v", err)
	}
	_, err1 := io.ReadAll(response.Body)
	if err1 != nil {
		t.Errorf("Failed to read HTTP response: %v", err)
	}
	response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}
