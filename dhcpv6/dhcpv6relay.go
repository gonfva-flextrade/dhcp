package dhcpv6

import (
	"fmt"
	"net"
)

const RelayHeaderSize = 34

type DHCPv6Relay struct {
	messageType MessageType
	hopCount    uint8
	linkAddr    net.IP
	peerAddr    net.IP
	payload     []byte // TODO implement relay payload
}

func (r *DHCPv6Relay) Type() MessageType {
	return r.messageType
}

func (r *DHCPv6Relay) MessageTypeToString() string {
	return MessageTypeToString(r.messageType)
}

func (r *DHCPv6Relay) Summary() string {
	ret := fmt.Sprintf(
		"DHCPv6Relay\n"+
			"  messageType=%v\n"+
			"  hopcount=%v\n"+
			"  linkaddr=%v\n"+
			"  peeraddr=%v\n",
		r.MessageTypeToString(),
		r.hopCount,
		r.linkAddr,
		r.peerAddr,
	)
	return ret
}

func (r *DHCPv6Relay) ToBytes() []byte {
	ret := make([]byte, RelayHeaderSize)
	ret[0] = byte(r.messageType)
	ret[1] = byte(r.hopCount)
	copy(ret[2:18], r.peerAddr)
	copy(ret[18:34], r.linkAddr)
	ret = append(ret, r.payload...)

	return ret
}
