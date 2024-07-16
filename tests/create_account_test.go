package tests

import (
	"github.com/xorwise/nedozerov-todo/internal/server"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreationHandler(t *testing.T) {
	s := server.NewServer()
	server := httptest.NewServer(s.Handler)
	defer server.Close()
	resp, err := http.Post(server.URL+"/accounts", "application/json", nil)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expectedBody := `{"account_id":1}`
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	got := strings.TrimSpace(string(body))
	if got != expectedBody {
		t.Errorf("expected body %q; got %q", expectedBody, body)
	}
}
