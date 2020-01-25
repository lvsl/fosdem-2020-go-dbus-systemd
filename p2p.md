# D-Bus Message Protocol

Message protocol is a part of [D-Bus specification](https://dbus.freedesktop.org/doc/dbus-specification.html#introduction) and defines the following:
* Data types and wire format
* Message protocol
* Authentication protocol
* Supported transports
* Set of standard interfaces in `org.freedesktop.DBus`
* XML-based introspection schema format

## Message Format
<table>
  <tr>
   <td>
    HEADER:
    <table>
      <tr>
        <td><tt>l</tt> or <tt>B</tt></td>
        <td>TYPE</td>
        <td>FLAGS</td>
        <td>VERSION</td>
      </tr>
      <tr>
        <td colspan="4">LENGTH OF BODY</td>
      </tr>
      <tr>
        <td colspan="4">SERIAL COOKIE</td>
      </tr>
    </table>
   </td>
   <td>
     TYPES:
     <ul>
       <li>METHOD CALL</li>
       <li>METHOD RETURN</li>
       <li>ERROR</li>
       <li>SIGNAL</li>
       <li>INVALID</li>
     </ul>
    </td>
  </tr>

  <tr>
    <td>
      HEADER FIELDS(n):
      <table>
        <tr>
          <td>CODE<sub>1</sub></td>
          <td>BODY<sub>1</sub></td>
        </tr>
        <tr><td colspan="2">...</td></tr>
        <tr>
          <td>CODE<sub>n</sub></td>
          <td>BODY<sub>n</sub></td>
        </tr>
      </table>
    </td>
    <td>
      CODES:
      <ul>
        <li>PATH ()</li>
        <li>INTERFACE ()</li>
        <li>MEMBER ()</li>
        <li>ERROR_NAME ()</li>
        <li>REPLY_SERIAL ()</li>
        <li>DESTINATION ()</li>
        <li>SENDER ()</li>
        <li>SIGNATURE ()</li>
        <li>UNIX_FDS ()</li>
      </ul>
    </td>
  </tr>
</table>

## Peer-to-peer communication over D-Bus
D-Bus does not have a concept of client and server.
Peers can communicate directly over one of supported transports.

The process to setup connection is different when connecting to a peer directly or connecting to a message bus.
Peer-to-peer D-Bus is not very practical and only useful for resolving circular dependencies and testing.

For example `systemctl` (systemd CLI) communicate with systemd via D-Bus daemon,
but if called as root, connects directly to `systemd` daemon.
This is to resolve cirular dependency between D-Bus daemon and systemd daemon.

### Go Interface
[Godbus](https://github.com/godbus/dbus/tree/v5.0.3) does not have any special functionality to setup direct peer-to-peer connection.

Peers that act as a client can use [`Conn.Dial()`](https://github.com/godbus/dbus/blob/v5.0.3/conn.go#L158) method.

Peers that act as a server must configure transport layer directly (e.g. open Unix Socket) and handle D-Bus messages with [`Message.Decode()`](https://github.com/godbus/dbus/blob/v5.0.3/message.go#L125) and [`Message.EncodeTo()`](https://github.com/godbus/dbus/blob/v5.0.3/message.go#L213).
 
Take at sample [server](https://github.com/lvsl/fosdem-2020-go-dbus-systemd/blob/master/cmd/p2p/server.go) and [client](https://github.com/lvsl/fosdem-2020-go-dbus-systemd/blob/master/cmd/p2p/client.go).
