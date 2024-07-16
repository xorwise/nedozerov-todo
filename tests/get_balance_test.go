package tests

import (
	"github.com/xorwise/nedozerov-todo/internal/server"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetBalanceHandler(t *testing.T) {
	s := server.NewServer()
	server := httptest.NewServer(s.Handler)
	defer server.Close()
	resp, err := http.Post(server.URL+"/accounts", "application/json", nil)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	_, err = http.Post(server.URL+"/accounts/1/deposit", "application/json", strings.NewReader(`{"amount": 10}`))
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	resp, err = http.Get(server.URL + "/accounts/1/balance")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
}
