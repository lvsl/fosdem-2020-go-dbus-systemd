package main

import (
	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/google/uuid"
	"log"
)

const intro = `
<node>
	<interface name="com.github.lvsl.Worker">
		<method name="DoWork">
			<arg direction="out" type="s"/>
		</method>
	</interface>` + introspect.IntrospectDataString + `</node>`

type Worker struct{}

func (w Worker) DoWork() (string, *dbus.Error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return "", dbus.MakeFailedError(err)
	}

	// schedule some work here ...

	return token.String(), nil
}

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatalf("can't connect to session bus: %v", err)
	}
	defer conn.Close()

	w := Worker{}

	// Export object on the bus
	conn.Export(w, "/", "com.github.lvsl.Worker")
	conn.Export(introspect.Introspectable(intro), "/", "org.freedesktop.DBus.Introspectable")

	// register on a bus
	reply, err := conn.RequestName("com.github.lvsl.Worker", dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatalf("can't request a name: %v", err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		log.Fatalf("name taken?")
	}

	// serve...
    // dbus-send --session --print-reply --type=method_call
    // --dest=com.github.lvsl.Worker / com.github.lvsl.Worker.DoWork
	select {}
}
