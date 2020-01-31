package main

import (
	"io"
	"log"
	"net"
	"os"
)

const ServerAddr = "/tmp/dbus-hello.sock"

func handler(c net.Conn) {
	log.Printf("Client connected: [%s]", c.RemoteAddr().Network())
	defer c.Close()

	// auth
	w, err := c.Write([]byte("OK 1234\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Written %v bytes to client\n", w)
	w, err = c.Write([]byte("AGREE_UNIX_FD\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Written %v bytes to client\n", w)

	written, err := io.Copy(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Written %v bytes to stdout\n", written)
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
