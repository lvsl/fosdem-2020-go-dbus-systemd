package main

import (
	"github.com/godbus/dbus/v5"
	"log"
	"net"
	"os"
)

const ServerAddr = "/tmp/dbus-p2p.sock"

func handler(c net.Conn) {
	log.Printf("Client connected: [%s]", c.RemoteAddr().Network())
	defer c.Close()

	msg, err := dbus.DecodeMessage(c)
	if err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	log.Printf("D-Bus msg: %s", msg)
}

func main() {
	if err := os.RemoveAll(ServerAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handler(conn)
	}
}
