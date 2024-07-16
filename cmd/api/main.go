package main

import (
	"fmt"
	"github.com/xorwise/nedozerov-todo/internal/server"
	"log/slog"
)

func main() {

	server := server.NewServer()
	slog.Info("Starting server", "address", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
