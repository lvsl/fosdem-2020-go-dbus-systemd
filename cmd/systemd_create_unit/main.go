package main

import (
	"github.com/coreos/go-systemd/v22/dbus"
	"log"
)

func main() {
	conn, err := dbus.NewUserConnection()
	if err != nil {
		log.Fatalf("can't connect to --user systemd: %v", err)
	}
	defer conn.Close()

	units, err := conn.ListUnits()
	if err != nil {
		log.Fatalf("can't list units: %v", err)
	}

	log.Printf("Loadede units: %+v", units)
}
