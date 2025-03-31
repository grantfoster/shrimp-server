package udpserver

import (
	"log/slog"
	"net"
	"os"
	"strings"
	"time"
)

const (
	Address      = "localhost"
	ListenerPort = 6666
	TickRate     = 1 // how many minimum updates per second
)

func ListenUDP() {
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

	slog.Info("started udp server",
		"port", ListenerPort,
	)

	for {
		conn.SetReadDeadline(time.Now().Add((1 / TickRate) * time.Second))
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			opErr, ok := err.(*net.OpError)
			if ok == false {
				slog.Error(opErr.Error())
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
