package main

import (
	"log/slog"
	"net"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

const (
	Address      = "127.0.0.1"
	ListenerPort = 6666
	TickRate     = 1 // how many minimum updates per second
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}
	if os.Getenv("LOG_JSON") == "true" {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		slog.SetDefault(logger)
	} else {
		logger := slog.New(tint.NewHandler(os.Stdout, nil))
		slog.SetDefault(logger)
	}

	addr := net.UDPAddr{
		Port: ListenerPort,
		IP:   net.ParseIP(Address),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(0)
	}
	defer conn.Close()

	slog.Info("server started",
		"port", ListenerPort,
		"addr", Address,
	)

	for {
		conn.SetReadDeadline(time.Now().Add((1 / TickRate) * time.Second))
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			serr, ok := err.(*net.OpError)
			if ok == false {
				slog.Error(serr.Error())
			}
		}
		if n > 0 {
			slog.Info("udp packet received",
				"payload", strings.TrimSpace(string(buffer[:n])),
				"addr", addr.IP,
				"port", addr.Port,
			)
		}
	}
}
