package common

import (
	"bytes"
	"encoding/gob"
)

func ToGOB64(val interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(val); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func FromGOB64(val interface{}, buf []byte) error {
	dec := gob.NewDecoder(bytes.NewReader(buf))
	if err := dec.Decode(val); err != nil {
		return err
	}
	return nil
}
