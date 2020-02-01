# The D-Bus daemon
Reference implementation of D-Bus message bus is [dbus-daemon](https://dbus.freedesktop.org/doc/dbus-daemon.1.html).

## Monitoring daemon
Message bus provides a set on standard interaces which can be introspected with:
```
$ dbus-daemon --introspect
```

To query these interfaces, use [dbus-send](https://dbus.freedesktop.org/doc/dbus-send.1.html).

To see who's on the system bus:
```
$ dbus-send --system \
            --print-reply \
            --type=method_call \
            --dest=org.freedesktop.DBus / org.freedesktop.DBus.ListNames
```

Get some runtime stats from the daemon:
```
$ dbus-send --session \
            --dest=org.freedesktop.DBus \
            --type=method_call \
            --print-reply \
            /org/freedesktop/DBus org.freedesktop.DBus.Debug.Stats.GetStats
```
