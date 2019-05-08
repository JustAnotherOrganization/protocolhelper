package codecs

import (
	"encoding/json"
	"io"

	"justanother.org/protocolhelper/util"
)

// String is the codec for strings
type String string

// Decode will decode the type
func (s String) Decode(r io.Reader) (interface{}, error) {
	str, err := util.ReadString(r)
	return String(str), err
}

// Encode will encode the type
func (s String) Encode(w io.Writer) error {
	return util.WriteString(w, string(s))
}

// JSON is the codec for JSON encoded objects (technically strings)
type JSON struct {
	V interface{}
}

// Decode will decode the type
func (j JSON) Decode(r io.Reader) (interface{}, error) {
	s, err := util.ReadString(r)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(s), &j.V); err != nil {
		return nil, err
	}

	return j.V, err
}

// Encode will encode the type
func (j JSON) Encode(w io.Writer) error {
	data, err := json.Marshal(j.V)
	if err != nil {
		return err
	}

	str := String(string(data))
	return str.Encode(w)
}

// VarInt is the codec for ints
type VarInt int

// Decode will decode the type
func (v VarInt) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadVarInt(r)
	return VarInt(i), err
}

// Encode will encode the type
func (v VarInt) Encode(w io.Writer) error {
	return util.WriteVarInt(w, int(v))
}

// Boolean is the codec for bools
type Boolean bool

// Decode will decode the type
func (b Boolean) Decode(r io.Reader) (interface{}, error) {
	l, err := util.ReadBool(r)
	return Boolean(l), err
}

// Encode will encode the type
func (b Boolean) Encode(w io.Writer) error {
	return util.WriteBool(w, bool(b))
}

// Byte is the codec for bytes
type Byte byte

// Decode will decode the type
func (b Byte) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadInt8(r)
	return Byte(i), err
}

// Encode will encode the type
func (b Byte) Encode(w io.Writer) error {
	return util.WriteInt8(w, int8(b))
}

// UnsignedByte is the codec for uint8s
type UnsignedByte uint8

// Decode will decode the type
func (b UnsignedByte) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadUint8(r)
	return UnsignedByte(i), err
}

// Encode will encode the type
func (b UnsignedByte) Encode(w io.Writer) error {
	return util.WriteUint8(w, uint8(b))
}

// ByteArray is the codec for arrays of bytes
type ByteArray []byte

// Decode will decode the type
func (b ByteArray) Decode(r io.Reader) (interface{}, error) {
	l, err := util.ReadVarInt(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, int(l))
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// Encode will encode the type
func (b ByteArray) Encode(w io.Writer) error {
	err := util.WriteVarInt(w, len(b))
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

// Short is the codec for int16s
type Short int16

// Decode will decode the type
func (s Short) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadInt16(r)
	return Short(i), err
}

// Encode will encode the type
func (s Short) Encode(w io.Writer) error {
	return util.WriteInt16(w, int16(s))
}

// UnsignedShort is the codec for uint16s
type UnsignedShort uint16

// Decode will decode the type
func (s UnsignedShort) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadUint16(r)
	return UnsignedShort(i), err
}

// Encode will encode the type
func (s UnsignedShort) Encode(w io.Writer) error {
	return util.WriteUint16(w, uint16(s))
}

// Int is the codec for int32s
type Int int32

// Decode will decode the type
func (i Int) Decode(r io.Reader) (interface{}, error) {
	i32, err := util.ReadInt32(r)
	return Int(i32), err
}

// Encode will encode the type
func (i Int) Encode(w io.Writer) error {
	return util.WriteInt32(w, int32(i))
}

// UnsignedInt is the codec for int32s
// TODO: why isn't this a uint32?
type UnsignedInt int32

// Decode will decode the type
func (i UnsignedInt) Decode(r io.Reader) (interface{}, error) {
	i32, err := util.ReadUint32(r)
	return UnsignedInt(i32), err
}

// Encode will encode the type
func (i UnsignedInt) Encode(w io.Writer) error {
	return util.WriteUint32(w, uint32(i))
}

// Long is the codec for int64s
type Long int64

// Decode will decode the type
func (l Long) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadInt64(r)
	return Long(i), err
}

// Encode will encode the type
func (l Long) Encode(w io.Writer) error {
	return util.WriteInt64(w, int64(l))
}

// UnsignedLong is the codec for uint64s
type UnsignedLong uint64

// Decode will decode the type
func (l UnsignedLong) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadUint64(r)
	return UnsignedLong(i), err
}

// Encode will encode the type
func (l UnsignedLong) Encode(w io.Writer) error {
	return util.WriteUint64(w, uint64(l))
}

// Float is the codec for float32s
type Float float32

// Decode will decode the type
func (f Float) Decode(r io.Reader) (interface{}, error) {
	ft, err := util.ReadInt64(r)
	return Float(ft), err
}

// Encode will encode the type
func (f Float) Encode(w io.Writer) error {
	return util.WriteFloat32(w, float32(f))
}

// Double is the codec for float64s
type Double float64

// Decode will decode the type
func (d Double) Decode(r io.Reader) (interface{}, error) {
	f, err := util.ReadFloat64(r)
	return Double(f), err
}

// Encode will encode the type
func (d Double) Encode(w io.Writer) error {
	return util.WriteFloat64(w, float64(d))
}
