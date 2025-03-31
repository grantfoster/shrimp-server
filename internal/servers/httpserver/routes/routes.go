package routes

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/grantfoster/shrimp-server/internal/servers/httpserver/middlewares"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", middlewares.BasicAuth(func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		slog.Info("http request received", "body", string(body))

		// Send a response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		_, err = w.Write([]byte("pong"))
		if err != nil {
			slog.Error("Error writing response:", "err", err)
		}
	}))
}
