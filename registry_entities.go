package schema_registry

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

var UnsupportedEvent = errors.New("unsupported event")
var UnsupportedEventVersion = errors.New("unsupported version event")

type caster func(v int, data interface{}) (proto.Message, bool)
type deCaster func(v int, data []byte) (proto.Message, error)

type SchemaRegistry interface {
	EncodeUserStreamEvent(event string, version int, payload interface{}) ([]byte, error)
	DecodeUserStreamEvent(message []byte) (map[int]interface{}, error)
	EncodeTaskStreamEvent(event string, version int, payload interface{}) ([]byte, error)
	DecodeTaskStreamEvent(message []byte) (map[int]interface{}, error)
}

type entityStreamEvent struct {
	Version int    `json:"v"`
	Event   string `json:"e"`
	Data    []byte `json:"d"`
}
