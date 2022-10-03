package database

import (
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

func unmarshallProto[T proto.Message, B string | []byte](b B, t T) {
	if err := proto.Unmarshal([]byte(b), t); err != nil {
		log.Panicf("failed to unmarshasl proto from redis: %+v", err)
	}
}

func marshallProto[T proto.Message](t T) []byte {
	bytes, err := proto.Marshal(t)
	if err != nil {
		log.Panicf("failed to marshal proto: %+v", err)
	}
	return bytes
}
