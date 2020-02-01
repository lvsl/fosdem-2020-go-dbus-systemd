package main

import (
	"github.com/godbus/dbus/v5"
	"log"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatalf("can't connect: %v", err)
	}
	defer conn.Close()

	obj := conn.Object("org.freedesktop.DBus", "/")
	call := obj.Call("org.freedesktop.DBus.ListNames", 0)

	var result []string
	if err := call.Store(&result); err != nil {
		log.Fatalf("can't complete the call: %v", err)
	}

	log.Printf("Call returned: %+v", result)
}
