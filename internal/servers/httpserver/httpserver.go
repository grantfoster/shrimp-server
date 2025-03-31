package httpserver

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/grantfoster/shrimp-server/internal/servers/httpserver/routes"
)

const Port = 7777

func ListenHttp() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	slog.Info("started http server", "port", Port)
	http.ListenAndServe(fmt.Sprintf(":%v", Port), mux)
}
