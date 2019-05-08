package packet

import (
	"justanother.org/protocolhelper/chat"
	"justanother.org/protocolhelper/protocol/codecs"
)

// StatusRequest represents a packet
type StatusRequest struct{}

// ID returns the packet ID
func (p StatusRequest) ID() int { return 0x00 }

// StatusResponse represents a packet
type StatusResponse struct {
	Status struct {
		Version struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		} `json:"version"`

		Players struct {
			Max    int `json:"max"`
			Online int `json:"online"`
		} `json:"players"`

		Description chat.TextComponent `json:"description"`
	}
}

// ID returns the packet ID
func (p StatusResponse) ID() int { return 0x00 }

// StatusPing represents a packet
type StatusPing struct {
	Payload codecs.Long
}

// ID returns the packet ID
func (p StatusPing) ID() int { return 0x01 }

// StatusPong represents a packet
type StatusPong struct {
	Payload codecs.Long
}

// ID returns the packet ID
func (p StatusPong) ID() int { return 0x01 }
