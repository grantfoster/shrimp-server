package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grantfoster/shrimp-server/app/helpers"
	"github.com/grantfoster/shrimp-server/app/servers/httpserver"
	"github.com/grantfoster/shrimp-server/app/servers/udpserver"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	helpers.SetupLogging()

	go udpserver.ListenUDP()
	go httpserver.ListenHttp()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	// time for cleanup before exit
	slog.Info("server shutdown", "datetime", time.Now().Format(time.RFC3339))
}
