# D-Bus Message Bus
Message Bus is a part of [D-Bus specification](https://dbus.freedesktop.org/doc/dbus-specification.html#message-bus)
and defines the following:
* Bus naming scheme
* Message routing
* A set of well known interfaces

D-Bus message bus reference implementation - [`dbus-daemon`](https://dbus.freedesktop.org/doc/dbus-daemon.1.html) is not considered to be part of the spec.

There's an alternative, compatible implementation of message bus specification for Linux - [`dbus-broker`](https://github.com/bus1/dbus-broker/wiki#using-dbus-broker). Goal of `dbus-broker` is to provide high performance and reliability.

Go's [`godbus`](https://github.com/godbus/dbus) provides functionality to connect to a Message Bus via [`dbus.Connect()`](https://www.godoc.org/github.com/godbus/dbus#Connect).
