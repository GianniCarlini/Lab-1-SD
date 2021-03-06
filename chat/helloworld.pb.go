// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: helloworld.proto

package helloworld

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The request message containing the user's name.
type PacketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor    int64  `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda   string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino  string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Tipo     string `protobuf:"bytes,6,opt,name=tipo,proto3" json:"tipo,omitempty"`
}

func (x *PacketRequest) Reset() {
	*x = PacketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketRequest) ProtoMessage() {}

func (x *PacketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketRequest.ProtoReflect.Descriptor instead.
func (*PacketRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *PacketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PacketRequest) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *PacketRequest) GetValor() int64 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *PacketRequest) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *PacketRequest) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *PacketRequest) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

// The response message containing the greetings
type PacketReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Nseg    int64  `protobuf:"varint,2,opt,name=nseg,proto3" json:"nseg,omitempty"`
}

func (x *PacketReply) Reset() {
	*x = PacketReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketReply) ProtoMessage() {}

func (x *PacketReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketReply.ProtoReflect.Descriptor instead.
func (*PacketReply) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *PacketReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PacketReply) GetNseg() int64 {
	if x != nil {
		return x.Nseg
	}
	return 0
}

//----------------------------------------------------------
type CamionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete   string `protobuf:"bytes,1,opt,name=id_paquete,json=idPaquete,proto3" json:"id_paquete,omitempty"`
	Seguimiento int64  `protobuf:"varint,2,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	Tipo        string `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor       int64  `protobuf:"varint,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Intentos    int64  `protobuf:"varint,5,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado      int64  `protobuf:"varint,6,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *CamionRequest) Reset() {
	*x = CamionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CamionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CamionRequest) ProtoMessage() {}

func (x *CamionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CamionRequest.ProtoReflect.Descriptor instead.
func (*CamionRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *CamionRequest) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *CamionRequest) GetSeguimiento() int64 {
	if x != nil {
		return x.Seguimiento
	}
	return 0
}

func (x *CamionRequest) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *CamionRequest) GetValor() int64 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *CamionRequest) GetIntentos() int64 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *CamionRequest) GetEstado() int64 {
	if x != nil {
		return x.Estado
	}
	return 0
}

type CamionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete   string `protobuf:"bytes,1,opt,name=id_paquete,json=idPaquete,proto3" json:"id_paquete,omitempty"`
	Seguimiento int64  `protobuf:"varint,2,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	Tipo        string `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor       int64  `protobuf:"varint,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Intentos    int64  `protobuf:"varint,5,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado      int64  `protobuf:"varint,6,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *CamionReply) Reset() {
	*x = CamionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CamionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CamionReply) ProtoMessage() {}

func (x *CamionReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CamionReply.ProtoReflect.Descriptor instead.
func (*CamionReply) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *CamionReply) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *CamionReply) GetSeguimiento() int64 {
	if x != nil {
		return x.Seguimiento
	}
	return 0
}

func (x *CamionReply) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *CamionReply) GetValor() int64 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *CamionReply) GetIntentos() int64 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *CamionReply) GetEstado() int64 {
	if x != nil {
		return x.Estado
	}
	return 0
}

//-----------------------------------------------------------
type SeguimientoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codigo int64 `protobuf:"varint,1,opt,name=codigo,proto3" json:"codigo,omitempty"`
}

func (x *SeguimientoRequest) Reset() {
	*x = SeguimientoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeguimientoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeguimientoRequest) ProtoMessage() {}

func (x *SeguimientoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeguimientoRequest.ProtoReflect.Descriptor instead.
func (*SeguimientoRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{4}
}

func (x *SeguimientoRequest) GetCodigo() int64 {
	if x != nil {
		return x.Codigo
	}
	return 0
}

type SeguimientoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estadoseguimiento string `protobuf:"bytes,1,opt,name=estadoseguimiento,proto3" json:"estadoseguimiento,omitempty"`
}

func (x *SeguimientoReply) Reset() {
	*x = SeguimientoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeguimientoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeguimientoReply) ProtoMessage() {}

func (x *SeguimientoReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeguimientoReply.ProtoReflect.Descriptor instead.
func (*SeguimientoReply) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{5}
}

func (x *SeguimientoReply) GetEstadoseguimiento() string {
	if x != nil {
		return x.Estadoseguimiento
	}
	return ""
}

//-----------------------------------------------------------
type OdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OdRequest) Reset() {
	*x = OdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OdRequest) ProtoMessage() {}

func (x *OdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OdRequest.ProtoReflect.Descriptor instead.
func (*OdRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{6}
}

func (x *OdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origen  string `protobuf:"bytes,1,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino string `protobuf:"bytes,2,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *OdReply) Reset() {
	*x = OdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OdReply) ProtoMessage() {}

func (x *OdReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OdReply.ProtoReflect.Descriptor instead.
func (*OdReply) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{7}
}

func (x *OdReply) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *OdReply) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

var File_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x97,
	0x01, 0x0a, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x22, 0x3b, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x73, 0x65, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x6e, 0x73, 0x65, 0x67, 0x22, 0xae, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x50,
	0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d,
	0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x65, 0x67,
	0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x6d, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x50, 0x61,
	0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69,
	0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x2c, 0x0a, 0x12, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69,
	0x65, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x64,
	0x69, 0x67, 0x6f, 0x22, 0x40, 0x0a, 0x10, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e,
	0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d,
	0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x4f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x3b, 0x0a, 0x07, 0x4f, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x32,
	0x93, 0x02, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x53, 0x65,
	0x6e, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x06, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74,
	0x6f, 0x12, 0x1e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53,
	0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53,
	0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x02, 0x4f, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4f, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_proto_rawDescData = file_helloworld_proto_rawDesc
)

func file_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_proto_rawDescData)
	})
	return file_helloworld_proto_rawDescData
}

var file_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_helloworld_proto_goTypes = []interface{}{
	(*PacketRequest)(nil),      // 0: helloworld.PacketRequest
	(*PacketReply)(nil),        // 1: helloworld.PacketReply
	(*CamionRequest)(nil),      // 2: helloworld.CamionRequest
	(*CamionReply)(nil),        // 3: helloworld.CamionReply
	(*SeguimientoRequest)(nil), // 4: helloworld.SeguimientoRequest
	(*SeguimientoReply)(nil),   // 5: helloworld.SeguimientoReply
	(*OdRequest)(nil),          // 6: helloworld.OdRequest
	(*OdReply)(nil),            // 7: helloworld.OdReply
}
var file_helloworld_proto_depIdxs = []int32{
	0, // 0: helloworld.Packet.SendPacket:input_type -> helloworld.PacketRequest
	2, // 1: helloworld.Packet.Camion:input_type -> helloworld.CamionRequest
	4, // 2: helloworld.Packet.Seguimiento:input_type -> helloworld.SeguimientoRequest
	6, // 3: helloworld.Packet.Od:input_type -> helloworld.OdRequest
	1, // 4: helloworld.Packet.SendPacket:output_type -> helloworld.PacketReply
	3, // 5: helloworld.Packet.Camion:output_type -> helloworld.CamionReply
	5, // 6: helloworld.Packet.Seguimiento:output_type -> helloworld.SeguimientoReply
	7, // 7: helloworld.Packet.Od:output_type -> helloworld.OdReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_proto_init() }
func file_helloworld_proto_init() {
	if File_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketRequest); i {
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
		file_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketReply); i {
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
		file_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CamionRequest); i {
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
		file_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CamionReply); i {
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
		file_helloworld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeguimientoRequest); i {
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
		file_helloworld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeguimientoReply); i {
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
		file_helloworld_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OdRequest); i {
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
		file_helloworld_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OdReply); i {
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
			RawDescriptor: file_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_proto = out.File
	file_helloworld_proto_rawDesc = nil
	file_helloworld_proto_goTypes = nil
	file_helloworld_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PacketClient is the client API for Packet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PacketClient interface {
	// Sends a greeting
	SendPacket(ctx context.Context, in *PacketRequest, opts ...grpc.CallOption) (*PacketReply, error)
	Camion(ctx context.Context, opts ...grpc.CallOption) (Packet_CamionClient, error)
	Seguimiento(ctx context.Context, in *SeguimientoRequest, opts ...grpc.CallOption) (*SeguimientoReply, error)
	Od(ctx context.Context, in *OdRequest, opts ...grpc.CallOption) (*OdReply, error)
}

type packetClient struct {
	cc grpc.ClientConnInterface
}

func NewPacketClient(cc grpc.ClientConnInterface) PacketClient {
	return &packetClient{cc}
}

func (c *packetClient) SendPacket(ctx context.Context, in *PacketRequest, opts ...grpc.CallOption) (*PacketReply, error) {
	out := new(PacketReply)
	err := c.cc.Invoke(ctx, "/helloworld.Packet/SendPacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packetClient) Camion(ctx context.Context, opts ...grpc.CallOption) (Packet_CamionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Packet_serviceDesc.Streams[0], "/helloworld.Packet/Camion", opts...)
	if err != nil {
		return nil, err
	}
	x := &packetCamionClient{stream}
	return x, nil
}

type Packet_CamionClient interface {
	Send(*CamionRequest) error
	Recv() (*CamionReply, error)
	grpc.ClientStream
}

type packetCamionClient struct {
	grpc.ClientStream
}

func (x *packetCamionClient) Send(m *CamionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *packetCamionClient) Recv() (*CamionReply, error) {
	m := new(CamionReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *packetClient) Seguimiento(ctx context.Context, in *SeguimientoRequest, opts ...grpc.CallOption) (*SeguimientoReply, error) {
	out := new(SeguimientoReply)
	err := c.cc.Invoke(ctx, "/helloworld.Packet/Seguimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packetClient) Od(ctx context.Context, in *OdRequest, opts ...grpc.CallOption) (*OdReply, error) {
	out := new(OdReply)
	err := c.cc.Invoke(ctx, "/helloworld.Packet/Od", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PacketServer is the server API for Packet service.
type PacketServer interface {
	// Sends a greeting
	SendPacket(context.Context, *PacketRequest) (*PacketReply, error)
	Camion(Packet_CamionServer) error
	Seguimiento(context.Context, *SeguimientoRequest) (*SeguimientoReply, error)
	Od(context.Context, *OdRequest) (*OdReply, error)
}

// UnimplementedPacketServer can be embedded to have forward compatible implementations.
type UnimplementedPacketServer struct {
}

func (*UnimplementedPacketServer) SendPacket(context.Context, *PacketRequest) (*PacketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPacket not implemented")
}
func (*UnimplementedPacketServer) Camion(Packet_CamionServer) error {
	return status.Errorf(codes.Unimplemented, "method Camion not implemented")
}
func (*UnimplementedPacketServer) Seguimiento(context.Context, *SeguimientoRequest) (*SeguimientoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seguimiento not implemented")
}
func (*UnimplementedPacketServer) Od(context.Context, *OdRequest) (*OdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Od not implemented")
}

func RegisterPacketServer(s *grpc.Server, srv PacketServer) {
	s.RegisterService(&_Packet_serviceDesc, srv)
}

func _Packet_SendPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PacketServer).SendPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Packet/SendPacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PacketServer).SendPacket(ctx, req.(*PacketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Packet_Camion_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PacketServer).Camion(&packetCamionServer{stream})
}

type Packet_CamionServer interface {
	Send(*CamionReply) error
	Recv() (*CamionRequest, error)
	grpc.ServerStream
}

type packetCamionServer struct {
	grpc.ServerStream
}

func (x *packetCamionServer) Send(m *CamionReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *packetCamionServer) Recv() (*CamionRequest, error) {
	m := new(CamionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Packet_Seguimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeguimientoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PacketServer).Seguimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Packet/Seguimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PacketServer).Seguimiento(ctx, req.(*SeguimientoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Packet_Od_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PacketServer).Od(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Packet/Od",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PacketServer).Od(ctx, req.(*OdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Packet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Packet",
	HandlerType: (*PacketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPacket",
			Handler:    _Packet_SendPacket_Handler,
		},
		{
			MethodName: "Seguimiento",
			Handler:    _Packet_Seguimiento_Handler,
		},
		{
			MethodName: "Od",
			Handler:    _Packet_Od_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Camion",
			Handler:       _Packet_Camion_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
