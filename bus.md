# D-Bus Message Bus
Message Bus is a part of [D-Bus specification](https://dbus.freedesktop.org/doc/dbus-specification.html#message-bus)
and defines the following:
* Bus naming scheme (unique and well-known)
* Two standard buses: system and session
* Message routing via match rules
* A set of methods in a well-known `org.freedesktop.DBus` interface

D-Bus message bus reference implementation - [`dbus-daemon`](https://dbus.freedesktop.org/doc/dbus-daemon.1.html) is not considered to be part of the spec.

There's an alternative, compatible implementation of message bus specification for Linux - [`dbus-broker`](https://github.com/bus1/dbus-broker/wiki#using-dbus-broker). Goal of `dbus-broker` is to provide high performance and reliability.

Go's [`godbus`](https://github.com/godbus/dbus) provides functionality to connect to a Message Bus via [`dbus.Connect()`](https://www.godoc.org/github.com/godbus/dbus#Connect).
