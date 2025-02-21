package main

import (
	"log"
	"net"
)

const Port = 6666

func main() {
	addr := net.UDPAddr{
		Port: 6666,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("server started on port %v", Port)

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Print(err)
		}

		log.Println(n)
		log.Println(addr)
	}
}
