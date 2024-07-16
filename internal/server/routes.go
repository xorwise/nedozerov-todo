package server

import (
	"encoding/json"
	"github.com/xorwise/nedozerov-todo/internal/domain"
	"github.com/xorwise/nedozerov-todo/internal/services"
	"log/slog"
	"net/http"
	"sync"
)

var accounts = make(map[int]domain.BankAccount)
var nextID = 1
var mu = &sync.Mutex{}

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /accounts", s.CreateAccountHandler)
	mux.HandleFunc("POST /accounts/{id}/deposit", s.DepositHandler)
	mux.HandleFunc("POST /accounts/{id}/withdraw", s.WithdrawHandler)
	mux.HandleFunc("GET /accounts/{id}/balance", s.GetBalanceHandler)

	return mux
}

func (s *Server) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	if _, ok := accounts[nextID]; ok {
		w.WriteHeader(http.StatusConflict)
		slog.Warn("account already exists", "account_id", nextID)
		return
	}
	account := services.NewAccount(nextID)
	accounts[nextID] = account
	var response domain.AccountCreationResponse
	response.AccountID = nextID
	nextID++
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
