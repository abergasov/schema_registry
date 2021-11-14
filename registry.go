package schema_registry

import (
	"encoding/json"

	"schema_registry/pkg/grpc/user"

	"google.golang.org/protobuf/proto"
)

type Registry struct {
	supportedUserVersions map[int]bool
}

func InitRegistry(supportedUserVersions []int) *Registry {
	result := &Registry{
		supportedUserVersions: make(map[int]bool),
	}
	for i := range supportedUserVersions {
		result.supportedUserVersions[supportedUserVersions[i]] = true
	}
	return result
}

func (r *Registry) EncodeUserStreamEvent(event string, version int, payload interface{}) ([]byte, error) {
	result := &entityStreamEvent{
		Event:   event,
		Version: version,
	}
	if ok, _ := r.supportedUserVersions[version]; !ok {
		return nil, UnsupportedEventVersion
	}

	var err error
	switch version {
	case 1:
		u, ok := payload.(*user.UserStreamV1)
		if ok {
			result.Data, err = proto.Marshal(u)
		}
	}
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

func (r *Registry) DecodeUserStreamEvent(message []byte) (map[int]interface{}, error) {
	var result entityStreamEvent
	err := json.Unmarshal(message, &result)
	if err != nil {
		return nil, err
	}
	if ok, _ := r.supportedUserVersions[result.Version]; !ok {
		return nil, UnsupportedEvent
	}
	switch result.Version {
	case 1:
		var usr user.UserStreamV1
		if err = proto.Unmarshal(result.Data, &usr); err != nil {
			return nil, err
		}
		return map[int]interface{}{1: &usr}, nil
	}
	return nil, UnsupportedEvent
}
