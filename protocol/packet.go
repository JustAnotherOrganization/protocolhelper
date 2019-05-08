package protocol

import (
	"bytes"
	"errors"
)

// Packet is a base packet.
type Packet struct {
	ID        int
	Direction Direction
	Data      bytes.Buffer
}

// Direction is the direction of the packet
type Direction int

// Directions
const (
	Serverbound Direction = iota
	Clientbound
)

// Possible Errors.
var (
	ErrUnknownPacketType   = errors.New("unknown packet type")
	ErrInvalidPacketLength = errors.New("received packet is below zero or above maximum size")
)
