package main

import (
	"log"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.Dial("unix:path=/tmp/dbus-p2p.sock")
	if err != nil {
		log.Fatalf("can't connect to session bus: %v", err)
	}
	defer conn.Close()

	msg := &dbus.Message{
		Type: dbus.TypeMethodCall,
		Headers: map[dbus.HeaderField]dbus.Variant{
			dbus.FieldDestination: dbus.MakeVariant("org.freedesktop.Notifications"),
			dbus.FieldPath:        dbus.MakeVariant(dbus.ObjectPath("/org/freedesktop/Notifications")),
			dbus.FieldInterface:   dbus.MakeVariant("org.freedesktop.Notifications"),
			dbus.FieldMember:      dbus.MakeVariant("Notify"),
			dbus.FieldSignature:   dbus.MakeVariant(dbus.ParseSignatureMust("susssasa{sv}i")),
		},
		Body: []interface{}{
			"app_name",
			uint32(0),
			"dialog-information",
			"Notification",
			"This is the body of a notification",
			[]string{"ok", "Ok"},
			map[string]dbus.Variant{
				"sound-name": dbus.MakeVariant("dialog-information"),
			},
			int32(-1),
		},
	}

	call := conn.Send(msg, nil)

	if call.Err != nil {
		log.Fatalf("can't send the  call: %v", call.Err)
	}

	log.Printf("Response body: %+v", call.Body)
}
