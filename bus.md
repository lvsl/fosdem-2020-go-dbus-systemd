# D-Bus Message Bus
Message Bus is a part of [D-Bus specification](https://dbus.freedesktop.org/doc/dbus-specification.html#message-bus)
and defines the following:
* Bus naming scheme (unique and well-known)
* Two standard buses: system and session
* Message routing via match rules
* A set of methods in a well-known `org.freedesktop.DBus` interface

D-Bus message bus reference implementation - [`dbus-daemon`](https://dbus.freedesktop.org/doc/dbus-daemon.1.html) is not considered to be part of the spec.

There's an alternative, compatible implementation of message bus specification for Linux - [`dbus-broker`](https://github.com/bus1/dbus-broker/wiki#using-dbus-broker). Goal of `dbus-broker` is to provide high performance and reliability.

## Systemd integration
D-Bus is tightly integrated with systemd. It's possible to start a new bus "session" via systemd socket activation,
as well as activate a systemd service by sending a message to the bus.

```
(systemd-cgls output on Ubuntu 18.04):
Control group /:
-.slice
├─user.slice
│ └─user-1000.slice
│   ├─user@1000.service
│   │ ├─init.scope
│   │ │ ├─1594 /lib/systemd/systemd --user
│   │ │ └─1595 (sd-pam)
│   │ └─dbus.service
│   │   └─1744 /usr/bin/dbus-daemon --session --address=systemd: --nofork --nopidfile --systemd-activation ...
...
```

## Go interface
Go's [`godbus`](https://github.com/godbus/dbus/tree/v5.0.3) provides functionality to connect to a Message Bus via [`dbus.SessionBus()`](https://www.godoc.org/github.com/godbus/dbus#SessionBus) and [`dbus.SystemBus()`](https://www.godoc.org/github.com/godbus/dbus#SystemBus).

### Integration testing
For testing with isolated D-Bus instance there is a [`dbus-run-session`](https://dbus.freedesktop.org/doc/dbus-run-session.1.html) tool.

Note: this is different from [`dbus-launch`](https://dbus.freedesktop.org/doc/dbus-launch.1.html) tool, which for starting session bus along with X11 session.
