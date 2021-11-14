package schema_registry_test

import (
	"testing"

	"schema_registry"
	"schema_registry/pkg/grpc/user"

	"github.com/google/uuid"
)

const (
	testRole  = "jedi"
	testEvent = "test"
)

func TestRegistry_DecodeUserStreamEvent(t *testing.T) {
	testUser := &user.UserStreamV1{
		Version:  666,
		PublicID: uuid.New().String(),
		UserRole: testRole,
		Active:   true,
		UserName: "ivan drago",
		UserMail: "super.ivan@gmail.com",
	}
	regio := schema_registry.InitRegistry([]int{1})
	msg, err := regio.EncodeUserStreamEvent(testEvent, 1, testUser)
	if err != nil {
		t.Fatal("unexpected err here", err)
	}

	regio2 := schema_registry.InitRegistry([]int{1})
	data, err := regio2.DecodeUserStreamEvent(msg)
	if err != nil {
		t.Fatal("unexpected err here", err)
	}
	if len(data) == 0 {
		t.Fatal("map is empty", err)
	}
	usr, ok := data[1].(*user.UserStreamV1)
	if !ok {
		t.Fatal("wrong repsonse")
	}
	match := usr.Version == testUser.Version && usr.UserName == testUser.UserName
	if !match {
		t.Fatal("users mismatch")
	}
}
