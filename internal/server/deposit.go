package server

import (
	"encoding/json"
	"github.com/xorwise/nedozerov-todo/internal/domain"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) DepositHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(`{"error": "invalid id"}`)
		slog.Warn("invalid id", "account_id", id)
		return
	}
	account, ok := accounts[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(`{"error": "account not found"}`)
		slog.Warn("account not found", "account_id", id)
		return
	}

	var request domain.DepositRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(`{"error": "invalid request"}`)
		slog.Warn("invalid request", "account_id", id)
		return
	}

	ch := make(chan error)

	go func() {
		err := account.Deposit(request.Amount)
		ch <- err
	}()

	select {
	case err := <-ch:
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(`{"error": "` + err.Error() + `"}`)
			slog.Warn(err.Error(), "account_id", id)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(account)
	case <-r.Context().Done():
		w.WriteHeader(http.StatusRequestTimeout)
		slog.Warn("request timeout", "account_id", id)
	}
}
