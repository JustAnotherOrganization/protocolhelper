package packet

import (
	"justanother.org/protocolhelper/chat"
	"justanother.org/protocolhelper/protocol/codecs"
)

// LoginStart represents a packet
type LoginStart struct {
	Username codecs.String
}

// ID returns the packet ID
func (p LoginStart) ID() int { return 0x00 }

// LoginSuccess represents a packet
type LoginSuccess struct {
	UUID     codecs.String
	Username codecs.String
}

// ID returns the packet ID
func (p LoginSuccess) ID() int { return 0x02 }

// LoginDisconnect represents a packet
type LoginDisconnect struct {
	Chat chat.TextComponent
}

// ID returns the packet ID
func (p LoginDisconnect) ID() int { return 0x00 }
