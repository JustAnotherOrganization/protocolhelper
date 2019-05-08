package codecs

import (
	"errors"
	"io"
)

// ErrUnknownCodecType is an error that happens when there is not a codec for that type.
var ErrUnknownCodecType = errors.New("unknown codec type")

// Codec is an interface for all supported Codecs
// Any packet to be encoded or decoded should have its types consist of codecs
type Codec interface {
	Decode(r io.Reader) (interface{}, error)
	Encode(w io.Writer) error
}
