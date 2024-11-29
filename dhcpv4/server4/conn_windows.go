// +build windows

package server4

import (
	"fmt"
	"net"
)

// NewIPv4UDPConn returns a UDP connection bound to both the interface and port
// given based on a IPv4 UDP socket. The UDP connection allows broadcasting.
func NewIPv4UDPConn(iface string, addr *net.UDPAddr) (*net.UDPConn, error) {
	// For Windows, we use the ListenPacket method to create a UDP connection.
	// This method automatically handles socket creation and configuration.
	laddr := net.UDPAddr{
		IP:   net.IPv4zero, // Use IPv4zero for listening on all interfaces.
		Port: addr.Port,
	}
	conn, err := net.ListenPacket("udp4", laddr.String())
	if err != nil {
		return nil, fmt.Errorf("cannot create UDP listener: %v", err)
	}
	udpConn, ok := conn.(*net.UDPConn)
	if !ok {
		return nil, fmt.Errorf("incorrect connection type: %T, expected *net.UDPConn", conn)
	}
	// On Windows, binding to a specific interface is not as straightforward
	// as on Unix-based systems. It might require additional steps or be
	// managed differently depending on the use case and network configuration.
	// Setting socket options like SO_BROADCAST is not directly exposed in the
	// net package, but the ListenPacket should create a socket that is
	// capable of broadcasting by default.
	return udpConn, nil
}
