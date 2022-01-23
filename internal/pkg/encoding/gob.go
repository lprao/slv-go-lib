package encoding

import (
	"bytes"
	"encoding/gob"
)

// Gob encoding helper interface
type GobEncoding interface {
	// Encodes the given integer using Gob encoder
	EncodeInt(int) ([]byte, error)

	// Decodes a buffer to integer usign Gob decoder
	DecodeInt(bytes.Buffer) (int, error)
}

// Encodes the given integer using Gob encoder
func EncodeInt(a int) ([]byte, error) {
	var e bytes.Buffer

	enc := gob.NewEncoder(&e)
	err := enc.Encode(a)
	if err != nil {
		return []byte{}, err
	}

	return e.Bytes(), nil
}

// Decodes a buffer to integer usign Gob decoder
func DecodeInt(b bytes.Buffer) (int, error) {
	var i int

	d := gob.NewDecoder(&b)
	err := d.Decode(&i)
	if err != nil {
		return 0, err
	}

	return i, nil
}
