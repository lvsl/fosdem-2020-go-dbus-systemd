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

	jobDone := make(chan string)
	props := []dbus.Property{} // TODO: fill these in
	jobid, err := conn.StartTransientUnit("coolunit.service", "fail", props, jobDone)
	if err != nil {
		log.Fatalf("can't list units: %v", err)
	}

	log.Printf("Started job: %v", jobid)

	status := <-jobDone

	log.Printf("Job done: %+v", status)
}
