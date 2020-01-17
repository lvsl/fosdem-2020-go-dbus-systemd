# D-Bus Message Bus
Message bus is defined in [D-Bus spec](https://dbus.freedesktop.org/doc/dbus-specification.html#message-bus)
and specify the following:
* Bus naming scheme
* Message routing
* Defines a set of well known interfaces

D-Bus message bus reference implementation - [`dbus-daemon`](https://dbus.freedesktop.org/doc/dbus-daemon.1.html) is not considered to be part of the spec.

There's an alternative implementation of message bus specification - [`dbus-broker`](https://github.com/bus1/dbus-broker/wiki#using-dbus-broker). Goal of `dbus-broker` is to provide high performance and reliability.
