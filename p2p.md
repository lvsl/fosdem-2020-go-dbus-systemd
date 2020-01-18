# D-Bus Message Protocol

Message protocol is a part of [D-Bus specification] and defines the following:
* Data types and wire format
* Message protocol
* Authentication protocol
* Supported transports
* Set of standard interfaces in `org.freedesktop.DBus`
* XML-based introspection schema format

## Peer-to-peer communication over D-Bus
D-Bus does not have a concept of client and server.
Peers can communicate directly over one of supported transports.

The process to setup connection is different when connecting to a peer directly or connecting to a message bus.
Peer-to-peer D-Bus is not very practical and only useful for resolving circular dependencies and testing.

### Go Interface
[Godbus]() does not have any special functionality to setup peer-to-peer connection.

Peers that act as a client can use [`Conn.Dial()`]() method.

Peers that act as a server must configure transport layer directly, and handle D-Bus messages with
 
