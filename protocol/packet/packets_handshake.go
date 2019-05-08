package packet

import "justanother.org/protocolhelper/protocol/codecs"

// Handshake represents a packet
type Handshake struct {
	ProtocolVersion codecs.VarInt
	ServerAddress   codecs.String
	ServerPort      codecs.UnsignedShort
	NextState       codecs.VarInt
}

// ID returns the packet ID
func (p Handshake) ID() int { return 0x00 }
