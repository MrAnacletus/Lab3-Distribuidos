// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: comunicacion.proto

package proto

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comunicacion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comunicacion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Jugada struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Jugada int32 `protobuf:"varint,2,opt,name=Jugada,proto3" json:"Jugada,omitempty"`
}

func (x *Jugada) Reset() {
	*x = Jugada{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comunicacion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jugada) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jugada) ProtoMessage() {}

func (x *Jugada) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jugada.ProtoReflect.Descriptor instead.
func (*Jugada) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{2}
}

func (x *Jugada) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Jugada) GetJugada() int32 {
	if x != nil {
		return x.Jugada
	}
	return 0
}

type Resultado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Estado int32 `protobuf:"varint,2,opt,name=Estado,proto3" json:"Estado,omitempty"`
}

func (x *Resultado) Reset() {
	*x = Resultado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comunicacion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resultado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resultado) ProtoMessage() {}

func (x *Resultado) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resultado.ProtoReflect.Descriptor instead.
func (*Resultado) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{3}
}

func (x *Resultado) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Resultado) GetEstado() int32 {
	if x != nil {
		return x.Estado
	}
	return 0
}

type Jugada2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID1     []int32 `protobuf:"varint,1,rep,packed,name=ID1,proto3" json:"ID1,omitempty"`
	ID2     []int32 `protobuf:"varint,2,rep,packed,name=ID2,proto3" json:"ID2,omitempty"`
	Jugada1 []int32 `protobuf:"varint,3,rep,packed,name=Jugada1,proto3" json:"Jugada1,omitempty"`
	Jugada2 []int32 `protobuf:"varint,4,rep,packed,name=Jugada2,proto3" json:"Jugada2,omitempty"`
}

func (x *Jugada2) Reset() {
	*x = Jugada2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comunicacion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jugada2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jugada2) ProtoMessage() {}

func (x *Jugada2) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jugada2.ProtoReflect.Descriptor instead.
func (*Jugada2) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{4}
}

func (x *Jugada2) GetID1() []int32 {
	if x != nil {
		return x.ID1
	}
	return nil
}

func (x *Jugada2) GetID2() []int32 {
	if x != nil {
		return x.ID2
	}
	return nil
}

func (x *Jugada2) GetJugada1() []int32 {
	if x != nil {
		return x.Jugada1
	}
	return nil
}

func (x *Jugada2) GetJugada2() []int32 {
	if x != nil {
		return x.Jugada2
	}
	return nil
}

type Jugada3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID1     int32 `protobuf:"varint,1,opt,name=ID1,proto3" json:"ID1,omitempty"`
	ID2     int32 `protobuf:"varint,2,opt,name=ID2,proto3" json:"ID2,omitempty"`
	Jugada1 int32 `protobuf:"varint,3,opt,name=Jugada1,proto3" json:"Jugada1,omitempty"`
	Jugada2 int32 `protobuf:"varint,4,opt,name=Jugada2,proto3" json:"Jugada2,omitempty"`
}

func (x *Jugada3) Reset() {
	*x = Jugada3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comunicacion_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jugada3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jugada3) ProtoMessage() {}

func (x *Jugada3) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jugada3.ProtoReflect.Descriptor instead.
func (*Jugada3) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{5}
}

func (x *Jugada3) GetID1() int32 {
	if x != nil {
		return x.ID1
	}
	return 0
}

func (x *Jugada3) GetID2() int32 {
	if x != nil {
		return x.ID2
	}
	return 0
}

func (x *Jugada3) GetJugada1() int32 {
	if x != nil {
		return x.Jugada1
	}
	return 0
}

func (x *Jugada3) GetJugada2() int32 {
	if x != nil {
		return x.Jugada2
	}
	return 0
}

var File_comunicacion_proto protoreflect.FileDescriptor

var file_comunicacion_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x26,
	0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x06, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x22, 0x33, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x61, 0x0a,
	0x07, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x31, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x49, 0x44, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44,
	0x32, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x49, 0x44, 0x32, 0x12, 0x18, 0x0a, 0x07,
	0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x31, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x4a,
	0x75, 0x67, 0x61, 0x64, 0x61, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61,
	0x32, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x32,
	0x22, 0x61, 0x0a, 0x07, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x49,
	0x44, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x49, 0x44, 0x31, 0x12, 0x10, 0x0a,
	0x03, 0x49, 0x44, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x49, 0x44, 0x32, 0x12,
	0x18, 0x0a, 0x07, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x75, 0x67,
	0x61, 0x64, 0x61, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4a, 0x75, 0x67, 0x61,
	0x64, 0x61, 0x32, 0x32, 0x41, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x72, 0x41, 0x6e, 0x61, 0x63, 0x6c, 0x65, 0x74, 0x75, 0x73,
	0x2f, 0x4c, 0x61, 0x62, 0x33, 0x2d, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x69, 0x64,
	0x6f, 0x73, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comunicacion_proto_rawDescOnce sync.Once
	file_comunicacion_proto_rawDescData = file_comunicacion_proto_rawDesc
)

func file_comunicacion_proto_rawDescGZIP() []byte {
	file_comunicacion_proto_rawDescOnce.Do(func() {
		file_comunicacion_proto_rawDescData = protoimpl.X.CompressGZIP(file_comunicacion_proto_rawDescData)
	})
	return file_comunicacion_proto_rawDescData
}

var file_comunicacion_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_comunicacion_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: grpc.HelloRequest
	(*HelloReply)(nil),   // 1: grpc.HelloReply
	(*Jugada)(nil),       // 2: grpc.Jugada
	(*Resultado)(nil),    // 3: grpc.Resultado
	(*Jugada2)(nil),      // 4: grpc.Jugada2
	(*Jugada3)(nil),      // 5: grpc.Jugada3
}
var file_comunicacion_proto_depIdxs = []int32{
	0, // 0: grpc.BrokerService.SayHello:input_type -> grpc.HelloRequest
	1, // 1: grpc.BrokerService.SayHello:output_type -> grpc.HelloReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comunicacion_proto_init() }
func file_comunicacion_proto_init() {
	if File_comunicacion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comunicacion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_comunicacion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_comunicacion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jugada); i {
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
		file_comunicacion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resultado); i {
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
		file_comunicacion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jugada2); i {
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
		file_comunicacion_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jugada3); i {
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
			RawDescriptor: file_comunicacion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comunicacion_proto_goTypes,
		DependencyIndexes: file_comunicacion_proto_depIdxs,
		MessageInfos:      file_comunicacion_proto_msgTypes,
	}.Build()
	File_comunicacion_proto = out.File
	file_comunicacion_proto_rawDesc = nil
	file_comunicacion_proto_goTypes = nil
	file_comunicacion_proto_depIdxs = nil
}
