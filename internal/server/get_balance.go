package server

import (
	"encoding/json"
	"github.com/xorwise/nedozerov-todo/internal/domain"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid id")
		slog.Warn("invalid id", "account_id", id)
		return
	}
	account, ok := accounts[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("account not found")
		slog.Warn("account not found", "account_id", id)
		return
	}
	ch := make(chan float64)
	go func() {
		ch <- account.GetBalance()
	}()

	select {
	case balance := <-ch:
		w.WriteHeader(http.StatusOK)
		response := domain.BalanceResponse{Balance: balance}
		json.NewEncoder(w).Encode(response)
	case <-r.Context().Done():
		w.WriteHeader(http.StatusRequestTimeout)
		slog.Warn("request timeout", "account_id", id)
	}
}
