package packet

import (
	"justanother.org/protocolhelper/chat"
	"justanother.org/protocolhelper/protocol/codecs"
)

// PlayKeepAlive represents a packet
type PlayKeepAlive struct {
	AliveID codecs.VarInt
}

// ID returns the packet ID
func (p PlayKeepAlive) ID() int { return 0x1F }

// PlayChatMessage represents a packet
// TODO: Verify that chat.TextComponent is being read correctly.
//  Even though we don't plan on using these types, we will want
//  to have a good example.
type PlayChatMessage struct {
	Chat     chat.TextComponent
	Position codecs.Byte
}

// ID returns the packet ID
func (p PlayChatMessage) ID() int { return 0x0F }

// PlayJoinGame represents a packet
type PlayJoinGame struct {
	EntityID   codecs.Int
	Gamemode   codecs.UnsignedByte
	Dimension  codecs.Int
	Difficulty codecs.UnsignedByte
	MaxPlayers codecs.UnsignedByte
	LevelType  codecs.String
	Debug      codecs.Boolean
}

// ID returns the packet ID
func (p PlayJoinGame) ID() int { return 0x23 }

// PlaySpawnPosition represents a packet
type PlaySpawnPosition struct {
	Location codecs.Long
}

// ID returns the packet ID
func (p PlaySpawnPosition) ID() int { return 0x43 }

// PlayPositionAndLook represents a packet
type PlayPositionAndLook struct {
	X     codecs.Double
	Y     codecs.Double
	Z     codecs.Double
	Yaw   codecs.Float
	Pitch codecs.Float
	Flags codecs.Byte
	Data  codecs.VarInt
}

// ID returns the packet ID
func (p PlayPositionAndLook) ID() int { return 0x2E }
