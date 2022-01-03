// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: gRPC/Proto.proto

package gRPC

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

type IncrementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     int32  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	RequestID int32  `protobuf:"varint,2,opt,name=requestID,proto3" json:"requestID,omitempty"`
	ClientID  string `protobuf:"bytes,3,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (x *IncrementRequest) Reset() {
	*x = IncrementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_Proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementRequest) ProtoMessage() {}

func (x *IncrementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_Proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementRequest.ProtoReflect.Descriptor instead.
func (*IncrementRequest) Descriptor() ([]byte, []int) {
	return file_gRPC_Proto_proto_rawDescGZIP(), []int{0}
}

func (x *IncrementRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *IncrementRequest) GetRequestID() int32 {
	if x != nil {
		return x.RequestID
	}
	return 0
}

func (x *IncrementRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

type IncrementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IncrementResponse) Reset() {
	*x = IncrementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_Proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementResponse) ProtoMessage() {}

func (x *IncrementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_Proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementResponse.ProtoReflect.Descriptor instead.
func (*IncrementResponse) Descriptor() ([]byte, []int) {
	return file_gRPC_Proto_proto_rawDescGZIP(), []int{1}
}

func (x *IncrementResponse) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_Proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_Proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_gRPC_Proto_proto_rawDescGZIP(), []int{2}
}

func (x *Ping) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_Proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_Proto_proto_msgTypes[3]
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
	return file_gRPC_Proto_proto_rawDescGZIP(), []int{3}
}

var File_gRPC_Proto_proto protoreflect.FileDescriptor

var file_gRPC_Proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x10, 0x49, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x29, 0x0a,
	0x11, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x7e, 0x0a, 0x0f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x40, 0x0a, 0x09, 0x49, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0a, 0x50, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_Proto_proto_rawDescOnce sync.Once
	file_gRPC_Proto_proto_rawDescData = file_gRPC_Proto_proto_rawDesc
)

func file_gRPC_Proto_proto_rawDescGZIP() []byte {
	file_gRPC_Proto_proto_rawDescOnce.Do(func() {
		file_gRPC_Proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_Proto_proto_rawDescData)
	})
	return file_gRPC_Proto_proto_rawDescData
}

var file_gRPC_Proto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gRPC_Proto_proto_goTypes = []interface{}{
	(*IncrementRequest)(nil),  // 0: Proto.IncrementRequest
	(*IncrementResponse)(nil), // 1: Proto.IncrementResponse
	(*Ping)(nil),              // 2: Proto.Ping
	(*Empty)(nil),             // 3: Proto.Empty
}
var file_gRPC_Proto_proto_depIdxs = []int32{
	0, // 0: Proto.IncrementSystem.Increment:input_type -> Proto.IncrementRequest
	2, // 1: Proto.IncrementSystem.PingServer:input_type -> Proto.Ping
	1, // 2: Proto.IncrementSystem.Increment:output_type -> Proto.IncrementResponse
	3, // 3: Proto.IncrementSystem.PingServer:output_type -> Proto.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gRPC_Proto_proto_init() }
func file_gRPC_Proto_proto_init() {
	if File_gRPC_Proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gRPC_Proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementRequest); i {
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
		file_gRPC_Proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementResponse); i {
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
		file_gRPC_Proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_gRPC_Proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gRPC_Proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_Proto_proto_goTypes,
		DependencyIndexes: file_gRPC_Proto_proto_depIdxs,
		MessageInfos:      file_gRPC_Proto_proto_msgTypes,
	}.Build()
	File_gRPC_Proto_proto = out.File
	file_gRPC_Proto_proto_rawDesc = nil
	file_gRPC_Proto_proto_goTypes = nil
	file_gRPC_Proto_proto_depIdxs = nil
}
