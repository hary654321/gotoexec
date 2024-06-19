// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: implant.proto

package grpcapi

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

//Command消息包含两个字段，一个用于维护操作系统的命令；一个用于维护命令执行的输出
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Ip string
	In  string `protobuf:"bytes,1,opt,name=In,proto3" json:"In,omitempty"`
	Out string `protobuf:"bytes,2,opt,name=Out,proto3" json:"Out,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetIn() string {
	if x != nil {
		return x.In
	}
	return ""
}

func (x *Command) GetOut() string {
	if x != nil {
		return x.Out
	}
	return ""
}

//Empty 用来代替null的空消息 定义这个Empty类型是由于gRPC不显式地允许空值
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{1}
}

type SleepTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int32 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *SleepTime) Reset() {
	*x = SleepTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SleepTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SleepTime) ProtoMessage() {}

func (x *SleepTime) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SleepTime.ProtoReflect.Descriptor instead.
func (*SleepTime) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{2}
}

func (x *SleepTime) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_implant_proto protoreflect.FileDescriptor

var file_implant_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x22, 0x2b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4f, 0x75, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f,
	0x0a, 0x09, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32,
	0x9f, 0x01, 0x0a, 0x07, 0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0c, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2e, 0x0a,
	0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x0e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x32, 0x6d, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x0a, 0x52, 0x75,
	0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x0c,
	0x53, 0x65, 0x74, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65,
	0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_implant_proto_rawDescOnce sync.Once
	file_implant_proto_rawDescData = file_implant_proto_rawDesc
)

func file_implant_proto_rawDescGZIP() []byte {
	file_implant_proto_rawDescOnce.Do(func() {
		file_implant_proto_rawDescData = protoimpl.X.CompressGZIP(file_implant_proto_rawDescData)
	})
	return file_implant_proto_rawDescData
}

var file_implant_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_implant_proto_goTypes = []interface{}{
	(*Command)(nil),   // 0: grpcapi.Command
	(*Empty)(nil),     // 1: grpcapi.Empty
	(*SleepTime)(nil), // 2: grpcapi.SleepTime
}
var file_implant_proto_depIdxs = []int32{
	1, // 0: grpcapi.Implant.FetchCommand:input_type -> grpcapi.Empty
	0, // 1: grpcapi.Implant.SendOutput:input_type -> grpcapi.Command
	1, // 2: grpcapi.Implant.GetSleepTime:input_type -> grpcapi.Empty
	0, // 3: grpcapi.Admin.RunCommand:input_type -> grpcapi.Command
	2, // 4: grpcapi.Admin.SetSleepTime:input_type -> grpcapi.SleepTime
	0, // 5: grpcapi.Implant.FetchCommand:output_type -> grpcapi.Command
	1, // 6: grpcapi.Implant.SendOutput:output_type -> grpcapi.Empty
	2, // 7: grpcapi.Implant.GetSleepTime:output_type -> grpcapi.SleepTime
	0, // 8: grpcapi.Admin.RunCommand:output_type -> grpcapi.Command
	1, // 9: grpcapi.Admin.SetSleepTime:output_type -> grpcapi.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_implant_proto_init() }
func file_implant_proto_init() {
	if File_implant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_implant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_implant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_implant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SleepTime); i {
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
			RawDescriptor: file_implant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_implant_proto_goTypes,
		DependencyIndexes: file_implant_proto_depIdxs,
		MessageInfos:      file_implant_proto_msgTypes,
	}.Build()
	File_implant_proto = out.File
	file_implant_proto_rawDesc = nil
	file_implant_proto_goTypes = nil
	file_implant_proto_depIdxs = nil
}
