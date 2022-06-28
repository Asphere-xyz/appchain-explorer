package common

import (
	"encoding/binary"
	"github.com/google/uuid"
)

func Uint64FromUUID(id uuid.UUID) uint64 {
	return binary.BigEndian.Uint64(id[:8])
}

func BinaryFromStringUUID(u string) []byte {
	bytes, _ := uuid.MustParse(u).MarshalBinary()
	return bytes
}

func TryParseUUID(u string) uuid.UUID {
	res, _ := uuid.Parse(u)
	return res
}
