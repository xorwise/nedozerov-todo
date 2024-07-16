package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		parts := strings.Split(r.URL.Path, "/")
		var id string
		if len(parts) > 2 {
			id = parts[2]
		}
		slog.Info(fmt.Sprintf("%s %s %s %s", r.Method, r.RequestURI, r.Proto, time.Since(start)), "account_id", id)
	})
}
