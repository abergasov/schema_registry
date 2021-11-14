// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: schema/task/task_stream/2.proto

package task

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskStreamV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicID     string `protobuf:"bytes,1,opt,name=publicID,proto3" json:"publicID,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	JiraID       string `protobuf:"bytes,3,opt,name=jiraID,proto3" json:"jiraID,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Author       string `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	AssignCost   int64  `protobuf:"varint,6,opt,name=assignCost,proto3" json:"assignCost,omitempty"`
	DoneCost     int64  `protobuf:"varint,7,opt,name=doneCost,proto3" json:"doneCost,omitempty"`
	Status       string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	PublicStatus string `protobuf:"bytes,9,opt,name=publicStatus,proto3" json:"publicStatus,omitempty"`
	CreatedAt    string `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	AssignedTo   string `protobuf:"bytes,11,opt,name=assignedTo,proto3" json:"assignedTo,omitempty"`
	AssignedAt   string `protobuf:"bytes,12,opt,name=assignedAt,proto3" json:"assignedAt,omitempty"`
}

func (x *TaskStreamV2) Reset() {
	*x = TaskStreamV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_task_task_stream_2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStreamV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStreamV2) ProtoMessage() {}

func (x *TaskStreamV2) ProtoReflect() protoreflect.Message {
	mi := &file_schema_task_task_stream_2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStreamV2.ProtoReflect.Descriptor instead.
func (*TaskStreamV2) Descriptor() ([]byte, []int) {
	return file_schema_task_task_stream_2_proto_rawDescGZIP(), []int{0}
}

func (x *TaskStreamV2) GetPublicID() string {
	if x != nil {
		return x.PublicID
	}
	return ""
}

func (x *TaskStreamV2) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskStreamV2) GetJiraID() string {
	if x != nil {
		return x.JiraID
	}
	return ""
}

func (x *TaskStreamV2) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskStreamV2) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *TaskStreamV2) GetAssignCost() int64 {
	if x != nil {
		return x.AssignCost
	}
	return 0
}

func (x *TaskStreamV2) GetDoneCost() int64 {
	if x != nil {
		return x.DoneCost
	}
	return 0
}

func (x *TaskStreamV2) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TaskStreamV2) GetPublicStatus() string {
	if x != nil {
		return x.PublicStatus
	}
	return ""
}

func (x *TaskStreamV2) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TaskStreamV2) GetAssignedTo() string {
	if x != nil {
		return x.AssignedTo
	}
	return ""
}

func (x *TaskStreamV2) GetAssignedAt() string {
	if x != nil {
		return x.AssignedAt
	}
	return ""
}

var File_schema_task_task_stream_2_proto protoreflect.FileDescriptor

var file_schema_task_task_stream_2_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0xe8, 0x02, 0x0a, 0x0c,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x32, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6a, 0x69, 0x72, 0x61, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6a, 0x69, 0x72, 0x61, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x64, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x54, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_task_task_stream_2_proto_rawDescOnce sync.Once
	file_schema_task_task_stream_2_proto_rawDescData = file_schema_task_task_stream_2_proto_rawDesc
)

func file_schema_task_task_stream_2_proto_rawDescGZIP() []byte {
	file_schema_task_task_stream_2_proto_rawDescOnce.Do(func() {
		file_schema_task_task_stream_2_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_task_task_stream_2_proto_rawDescData)
	})
	return file_schema_task_task_stream_2_proto_rawDescData
}

var file_schema_task_task_stream_2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_task_task_stream_2_proto_goTypes = []interface{}{
	(*TaskStreamV2)(nil), // 0: protobuf.TaskStreamV2
}
var file_schema_task_task_stream_2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_task_task_stream_2_proto_init() }
func file_schema_task_task_stream_2_proto_init() {
	if File_schema_task_task_stream_2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_task_task_stream_2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStreamV2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_task_task_stream_2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_task_task_stream_2_proto_goTypes,
		DependencyIndexes: file_schema_task_task_stream_2_proto_depIdxs,
		MessageInfos:      file_schema_task_task_stream_2_proto_msgTypes,
	}.Build()
	File_schema_task_task_stream_2_proto = out.File
	file_schema_task_task_stream_2_proto_rawDesc = nil
	file_schema_task_task_stream_2_proto_goTypes = nil
	file_schema_task_task_stream_2_proto_depIdxs = nil
}
