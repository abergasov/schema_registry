package schema_registry

import (
	"encoding/json"

	"schema_registry/pkg/grpc/task"
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

func (r *Registry) encodeEvent(event string, version int, payload interface{}, caster caster) ([]byte, error) {
	result := &entityStreamEvent{
		Event:   event,
		Version: version,
	}
	if ok, _ := r.supportedUserVersions[version]; !ok {
		return nil, UnsupportedEventVersion
	}
	msg, ok := caster(version, payload)
	if !ok {
		return nil, UnsupportedEvent
	}
	var err error
	result.Data, err = proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

func (r *Registry) decodeEvent(message []byte, caster deCaster) (map[int]interface{}, error) {
	var result entityStreamEvent
	err := json.Unmarshal(message, &result)
	if err != nil {
		return nil, err
	}
	if ok, _ := r.supportedUserVersions[result.Version]; !ok {
		return nil, UnsupportedEvent
	}
	payload, err := caster(result.Version, result.Data)
	if err != nil {
		return nil, err
	}
	return map[int]interface{}{result.Version: payload}, nil
}

func (r *Registry) EncodeUserStreamEvent(event string, version int, payload interface{}) ([]byte, error) {
	return r.encodeEvent(event, version, payload, func(v int, data interface{}) (proto.Message, bool) {
		switch v {
		case 1:
			usr, ok := payload.(*user.UserStreamV1)
			return usr, ok
		}
		return nil, false
	})
}

func (r *Registry) DecodeUserStreamEvent(message []byte) (map[int]interface{}, error) {
	return r.decodeEvent(message, func(v int, data []byte) (proto.Message, error) {
		switch v {
		case 1:
			var usr user.UserStreamV1
			err := proto.Unmarshal(data, &usr)
			return &usr, err
		}
		return nil, UnsupportedEvent
	})
}

func (r *Registry) EncodeTaskStreamEvent(event string, version int, payload interface{}) ([]byte, error) {
	return r.encodeEvent(event, version, payload, func(v int, data interface{}) (proto.Message, bool) {
		switch v {
		case 1:
			tsk, ok := payload.(*task.TaskStreamV1)
			return tsk, ok
		case 2:
			tsk, ok := payload.(*task.TaskStreamV2)
			return tsk, ok
		}
		return nil, false
	})
}

func (r *Registry) DecodeTaskStreamEvent(message []byte) (map[int]interface{}, error) {
	return r.decodeEvent(message, func(v int, data []byte) (proto.Message, error) {
		switch v {
		case 1:
			var usr task.TaskStreamV1
			err := proto.Unmarshal(data, &usr)
			return &usr, err
		case 2:
			var usr task.TaskStreamV2
			err := proto.Unmarshal(data, &usr)
			return &usr, err
		}
		return nil, UnsupportedEvent
	})
}
