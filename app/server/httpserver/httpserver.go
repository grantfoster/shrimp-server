package httpserver

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const Port = 7777

func ListenHttp() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
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
	})

	slog.Info("started http server", "port", Port)
	http.ListenAndServe(fmt.Sprintf(":%v", Port), nil)
}
