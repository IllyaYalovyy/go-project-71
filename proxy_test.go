package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IllyaYalovyy/go-project-71/reverseproxy"
)

func TestProxyIntegration(t *testing.T) {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer backendServer.Close()
	go func() {
		proxy, err := reverseproxy.New("", 8080, backendServer.URL)
		if err != nil {
			panic(err)
		}
		if proxy.Run() != nil {
			panic(err)
		}
	}()

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("Failed to make a request to the reverse proxy: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	expectedResponse := "Hello, World!"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if string(body) != expectedResponse {
		t.Errorf("Expected response body: %s, got: %s", expectedResponse, string(body))
	}
}
