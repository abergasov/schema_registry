package schema_registry

import "errors"

var UnsupportedEvent = errors.New("unsupported event")
var UnsupportedEventVersion = errors.New("unsupported version event")

type SchemaRegistry interface {
	EncodeUserStreamEvent(event string, version int, payload interface{}) ([]byte, error)
	DecodeUserStreamEvent(message []byte) (map[int]interface{}, error)
}

type entityStreamEvent struct {
	Version int    `json:"v"`
	Event   string `json:"e"`
	Data    []byte `json:"d"`
}
