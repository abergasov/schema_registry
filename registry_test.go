package schema_registry_test

import (
	"testing"

	"github.com/abergasov/schema_registry"
	"github.com/abergasov/schema_registry/pkg/grpc/task"
	"github.com/abergasov/schema_registry/pkg/grpc/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	testEvent = "test"
)

func TestRegistry_UserStream(t *testing.T) {
	testUser := &user.UserAccountV1{PublicID: uuid.New().String()}
	regio := schema_registry.InitRegistry([]int{1})
	msg, err := regio.EncodeUserStreamEvent(testEvent, 1, testUser)
	assert.Nil(t, err)

	regio2 := schema_registry.InitRegistry([]int{1})
	data, err := regio2.DecodeUserStreamEvent(msg)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, len(data), 1)

	usr, ok := data[1].(*user.UserAccountV1)
	assert.Equal(t, ok, true)
	assert.Equal(t, usr.PublicID, testUser.PublicID)
}

func TestRegistry_TaskStream(t *testing.T) {
	taskV1 := &task.TaskV1{PublicID: uuid.New().String()}
	taskV2 := &task.TaskV2{PublicID: uuid.New().String()}
	regio := schema_registry.InitRegistry([]int{1, 2})
	dataT1, err := regio.EncodeTaskStreamEvent(testEvent, 1, taskV1)
	assert.Nil(t, err)

	regio2 := schema_registry.InitRegistry([]int{1})
	msgMapV1, err := regio2.DecodeTaskStreamEvent(dataT1)

	assert.Nil(t, err)
	assert.NotNil(t, msgMapV1)
	assert.Equal(t, len(msgMapV1), 1)
	tsk1, ok := msgMapV1[1].(*task.TaskV1)
	assert.Equal(t, ok, true)
	assert.Equal(t, tsk1.PublicID, taskV1.PublicID)

	dataT2, err := regio.EncodeTaskStreamEvent(testEvent, 2, taskV2)
	assert.Nil(t, err)
	msgMapV2, err := regio2.DecodeTaskStreamEvent(dataT2)
	assert.NotNil(t, err)
	assert.Nil(t, msgMapV2)

	regio3 := schema_registry.InitRegistry([]int{2})
	msgMapV2, err = regio3.DecodeTaskStreamEvent(dataT2)
	assert.Nil(t, err)
	assert.NotNil(t, msgMapV2)
	assert.Equal(t, len(msgMapV2), 1)
	tsk2, ok := msgMapV2[2].(*task.TaskV2)
	assert.Equal(t, ok, true)
	assert.Equal(t, tsk2.PublicID, taskV2.PublicID)
}
