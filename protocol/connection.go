package protocol

import (
	"bytes"
	"io"
	"net"
	"reflect"

	"justanother.org/protocolhelper/protocol/codecs"
	"justanother.org/protocolhelper/protocol/packet"
	"justanother.org/protocolhelper/util"
)

type Connection struct {
	rw io.ReadWriteCloser

	State    State
	Protocol uint16
}

// State is the gameplay state.
type State uint8

// Different gameplay states.
const (
	Handshake State = iota
	Status
	Login
	Play
)

// NewConnection will wrap the net.Conn in a Connection struct
func NewConnection(conn net.Conn) *Connection {
	return &Connection{rw: conn}
}

// Next will read the next packet.
func (c *Connection) Next() (packet.Holder, error) {
	p, err := c.read()
	if err != nil {
		return nil, err
	}

	return c.decode(p)
}

// Write will write the packet h to the connection
func (c *Connection) Write(h packet.Holder) (int, error) {
	data, err := c.encode(h)
	if err != nil {
		return -1, err
	}

	err = util.WriteVarInt(c.rw, data.Len())
	if err != nil {
		return -1, err
	}

	n, err := data.WriteTo(c.rw)
	if err != nil {
		return int(n), err
	}

	return int(n), nil
}

// Close will attempt to close the connection
func (c *Connection) Close() error {
	if closer, ok := c.rw.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}

func (c *Connection) read() (*Packet, error) {
	length, err := util.ReadVarInt(c.rw)
	if err != nil {
		return nil, err
	}

	if length < 0 || length > 1048576 { // 2^(21-1)
		return nil, ErrInvalidPacketLength
	}

	payload := make([]byte, length)
	_, err = io.ReadFull(c.rw, payload)

	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(payload)
	id, err := util.ReadVarInt(buffer)

	if err != nil {
		return nil, err
	}

	return &Packet{
		ID:        id,
		Direction: Serverbound,
		Data:      *buffer,
	}, nil
}

func (c *Connection) decode(p *Packet) (packet.Holder, error) {

	packetType, err := GetPacketType(p.Direction, c.State, p.ID)
	if err != nil {
		return nil, err
	}

	inst := reflect.New(packetType).Elem()

	for i := 0; i < inst.NumField(); i++ {
		field := inst.Field(i)

		codec, ok := field.Interface().(codecs.Codec)
		if !ok {
			return nil, codecs.UnknownCodecType
		}

		value, err := codec.Decode(&p.Data)
		if err != nil {
			return nil, err
		}

		field.Set(reflect.ValueOf(value))
	}

	return inst.Interface().(packet.Holder), nil
}

func (c *Connection) encode(h packet.Holder) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	util.WriteVarInt(buffer, h.ID())

	value := reflect.ValueOf(h)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		codec, ok := field.Interface().(codecs.Codec)
		if !ok {
			if field.Kind() == reflect.Struct {
				codec = codecs.JSON{V: field.Interface()}
			} else {
				return nil, codecs.UnknownCodecType
			}
		}

		if err := codec.Encode(buffer); err != nil {
			return nil, err
		}
	}

	return buffer, nil
}
