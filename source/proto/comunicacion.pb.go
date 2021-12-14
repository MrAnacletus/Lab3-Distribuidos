// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/comunicacion.proto

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
		mi := &file_proto_comunicacion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comunicacion_proto_msgTypes[0]
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
	return file_proto_comunicacion_proto_rawDescGZIP(), []int{0}
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
		mi := &file_proto_comunicacion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comunicacion_proto_msgTypes[1]
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
	return file_proto_comunicacion_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ComandoSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comando string `protobuf:"bytes,1,opt,name=Comando,proto3" json:"Comando,omitempty"`
	Vector  string `protobuf:"bytes,2,opt,name=vector,proto3" json:"vector,omitempty"`
}

func (x *ComandoSend) Reset() {
	*x = ComandoSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comunicacion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComandoSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComandoSend) ProtoMessage() {}

func (x *ComandoSend) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comunicacion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComandoSend.ProtoReflect.Descriptor instead.
func (*ComandoSend) Descriptor() ([]byte, []int) {
	return file_proto_comunicacion_proto_rawDescGZIP(), []int{2}
}

func (x *ComandoSend) GetComando() string {
	if x != nil {
		return x.Comando
	}
	return ""
}

func (x *ComandoSend) GetVector() string {
	if x != nil {
		return x.Vector
	}
	return ""
}

type ComandoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comando string `protobuf:"bytes,1,opt,name=Comando,proto3" json:"Comando,omitempty"`
	Vector  string `protobuf:"bytes,2,opt,name=vector,proto3" json:"vector,omitempty"`
}

func (x *ComandoReply) Reset() {
	*x = ComandoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comunicacion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComandoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComandoReply) ProtoMessage() {}

func (x *ComandoReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comunicacion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComandoReply.ProtoReflect.Descriptor instead.
func (*ComandoReply) Descriptor() ([]byte, []int) {
	return file_proto_comunicacion_proto_rawDescGZIP(), []int{3}
}

func (x *ComandoReply) GetComando() string {
	if x != nil {
		return x.Comando
	}
	return ""
}

func (x *ComandoReply) GetVector() string {
	if x != nil {
		return x.Vector
	}
	return ""
}

type Rebeldes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numero string `protobuf:"bytes,1,opt,name=numero,proto3" json:"numero,omitempty"`
	Vector string `protobuf:"bytes,2,opt,name=vector,proto3" json:"vector,omitempty"`
}

func (x *Rebeldes) Reset() {
	*x = Rebeldes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comunicacion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rebeldes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rebeldes) ProtoMessage() {}

func (x *Rebeldes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comunicacion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rebeldes.ProtoReflect.Descriptor instead.
func (*Rebeldes) Descriptor() ([]byte, []int) {
	return file_proto_comunicacion_proto_rawDescGZIP(), []int{4}
}

func (x *Rebeldes) GetNumero() string {
	if x != nil {
		return x.Numero
	}
	return ""
}

func (x *Rebeldes) GetVector() string {
	if x != nil {
		return x.Vector
	}
	return ""
}

var File_proto_comunicacion_proto protoreflect.FileDescriptor

var file_proto_comunicacion_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x63, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63,
	0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x40, 0x0a,
	0x0c, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22,
	0x3a, 0x0a, 0x08, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0xb0, 0x01, 0x0a, 0x0d,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x35, 0x0a, 0x0d, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f,
	0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x11, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72,
	0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x4c, 0x65, 0x69, 0x61, 0x12, 0x11, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x53, 0x65, 0x6e, 0x64, 0x1a, 0x0e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x32, 0xa3,
	0x03, 0x0a, 0x0e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x36, 0x0a, 0x0d, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x43, 0x6f, 0x6d, 0x61, 0x6e,
	0x64, 0x6f, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64,
	0x6f, 0x53, 0x65, 0x6e, 0x64, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d,
	0x61, 0x6e, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x11, 0x45, 0x6e, 0x76,
	0x69, 0x61, 0x72, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x4c, 0x65, 0x69, 0x61, 0x12, 0x11,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x53, 0x65, 0x6e,
	0x64, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65,
	0x73, 0x12, 0x35, 0x0a, 0x0d, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x72, 0x4d, 0x65, 0x72,
	0x67, 0x65, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x13, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x72, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x45, 0x6e, 0x76, 0x69, 0x6f, 0x12,
	0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x12, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x43,
	0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x53, 0x65, 0x6e, 0x64, 0x1a, 0x12,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x4d, 0x65, 0x72, 0x67, 0x61, 0x6e, 0x12, 0x12, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x40, 0x0a, 0x17, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x43, 0x6f, 0x6d, 0x61,
	0x6e, 0x64, 0x6f, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x11, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x53, 0x65, 0x6e, 0x64,
	0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x72, 0x41, 0x6e, 0x61, 0x63, 0x6c, 0x65, 0x74, 0x75, 0x73, 0x2f, 0x4c,
	0x61, 0x62, 0x33, 0x2d, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x73,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_comunicacion_proto_rawDescOnce sync.Once
	file_proto_comunicacion_proto_rawDescData = file_proto_comunicacion_proto_rawDesc
)

func file_proto_comunicacion_proto_rawDescGZIP() []byte {
	file_proto_comunicacion_proto_rawDescOnce.Do(func() {
		file_proto_comunicacion_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_comunicacion_proto_rawDescData)
	})
	return file_proto_comunicacion_proto_rawDescData
}

var file_proto_comunicacion_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_comunicacion_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: grpc.HelloRequest
	(*HelloReply)(nil),   // 1: grpc.HelloReply
	(*ComandoSend)(nil),  // 2: grpc.ComandoSend
	(*ComandoReply)(nil), // 3: grpc.ComandoReply
	(*Rebeldes)(nil),     // 4: grpc.Rebeldes
}
var file_proto_comunicacion_proto_depIdxs = []int32{
	0,  // 0: grpc.BrokerService.SayHello:input_type -> grpc.HelloRequest
	0,  // 1: grpc.BrokerService.EnviarComando:input_type -> grpc.HelloRequest
	2,  // 2: grpc.BrokerService.EnviarComandoLeia:input_type -> grpc.ComandoSend
	2,  // 3: grpc.FulcrumService.EnviarComando:input_type -> grpc.ComandoSend
	2,  // 4: grpc.FulcrumService.EnviarComandoLeia:input_type -> grpc.ComandoSend
	0,  // 5: grpc.FulcrumService.InformarMerge:input_type -> grpc.HelloRequest
	0,  // 6: grpc.FulcrumService.InformarInicioEnvio:input_type -> grpc.HelloRequest
	2,  // 7: grpc.FulcrumService.EnviarComandoMerge:input_type -> grpc.ComandoSend
	0,  // 8: grpc.FulcrumService.Mergan:input_type -> grpc.HelloRequest
	2,  // 9: grpc.FulcrumService.EnviarComandoMergeFinal:input_type -> grpc.ComandoSend
	1,  // 10: grpc.BrokerService.SayHello:output_type -> grpc.HelloReply
	1,  // 11: grpc.BrokerService.EnviarComando:output_type -> grpc.HelloReply
	4,  // 12: grpc.BrokerService.EnviarComandoLeia:output_type -> grpc.Rebeldes
	3,  // 13: grpc.FulcrumService.EnviarComando:output_type -> grpc.ComandoReply
	4,  // 14: grpc.FulcrumService.EnviarComandoLeia:output_type -> grpc.Rebeldes
	1,  // 15: grpc.FulcrumService.InformarMerge:output_type -> grpc.HelloReply
	1,  // 16: grpc.FulcrumService.InformarInicioEnvio:output_type -> grpc.HelloReply
	3,  // 17: grpc.FulcrumService.EnviarComandoMerge:output_type -> grpc.ComandoReply
	1,  // 18: grpc.FulcrumService.Mergan:output_type -> grpc.HelloReply
	3,  // 19: grpc.FulcrumService.EnviarComandoMergeFinal:output_type -> grpc.ComandoReply
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_comunicacion_proto_init() }
func file_proto_comunicacion_proto_init() {
	if File_proto_comunicacion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_comunicacion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_comunicacion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_comunicacion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComandoSend); i {
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
		file_proto_comunicacion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComandoReply); i {
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
		file_proto_comunicacion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rebeldes); i {
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
			RawDescriptor: file_proto_comunicacion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_comunicacion_proto_goTypes,
		DependencyIndexes: file_proto_comunicacion_proto_depIdxs,
		MessageInfos:      file_proto_comunicacion_proto_msgTypes,
	}.Build()
	File_proto_comunicacion_proto = out.File
	file_proto_comunicacion_proto_rawDesc = nil
	file_proto_comunicacion_proto_goTypes = nil
	file_proto_comunicacion_proto_depIdxs = nil
}
