// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iotTempService.proto

/*
Package iottempservice is a generated protocol buffer package.

It is generated from these files:
	iotTempService.proto

It has these top-level messages:
	Thermostat
	Zone
	Room
	FindOneByIdReqest
	ThermostatArray
*/
package iot_temp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Thermostat_DeviceType int32

const (
	Thermostat_OTHER        Thermostat_DeviceType = 0
	Thermostat_ARDUINO      Thermostat_DeviceType = 1
	Thermostat_RASPBERRY_PI Thermostat_DeviceType = 2
)

var Thermostat_DeviceType_name = map[int32]string{
	0: "OTHER",
	1: "ARDUINO",
	2: "RASPBERRY_PI",
}
var Thermostat_DeviceType_value = map[string]int32{
	"OTHER":        0,
	"ARDUINO":      1,
	"RASPBERRY_PI": 2,
}

func (x Thermostat_DeviceType) String() string {
	return proto.EnumName(Thermostat_DeviceType_name, int32(x))
}
func (Thermostat_DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Thermostat struct {
	Id        uint64                     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address   string                     `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Port      int32                      `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	Type      Thermostat_DeviceType      `protobuf:"varint,5,opt,name=type,enum=com.avegao.iot.temp.service.Thermostat_DeviceType" json:"type,omitempty"`
	Zone      *Zone                      `protobuf:"bytes,6,opt,name=zone" json:"zone,omitempty"`
	Room      *Room                      `protobuf:"bytes,7,opt,name=room" json:"room,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	DeletedAt *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt" json:"deleted_at,omitempty"`
}

func (m *Thermostat) Reset()                    { *m = Thermostat{} }
func (m *Thermostat) String() string            { return proto.CompactTextString(m) }
func (*Thermostat) ProtoMessage()               {}
func (*Thermostat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Thermostat) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Thermostat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Thermostat) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Thermostat) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Thermostat) GetType() Thermostat_DeviceType {
	if m != nil {
		return m.Type
	}
	return Thermostat_OTHER
}

func (m *Thermostat) GetZone() *Zone {
	if m != nil {
		return m.Zone
	}
	return nil
}

func (m *Thermostat) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *Thermostat) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Thermostat) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Thermostat) GetDeletedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type Zone struct {
	Id        uint64                     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	DeletedAt *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt" json:"deleted_at,omitempty"`
}

func (m *Zone) Reset()                    { *m = Zone{} }
func (m *Zone) String() string            { return proto.CompactTextString(m) }
func (*Zone) ProtoMessage()               {}
func (*Zone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Zone) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Zone) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Zone) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Zone) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Zone) GetDeletedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type Room struct {
	Id        uint64                     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Zone      *Zone                      `protobuf:"bytes,3,opt,name=zone" json:"zone,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	DeletedAt *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt" json:"deleted_at,omitempty"`
}

func (m *Room) Reset()                    { *m = Room{} }
func (m *Room) String() string            { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()               {}
func (*Room) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Room) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Room) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Room) GetZone() *Zone {
	if m != nil {
		return m.Zone
	}
	return nil
}

func (m *Room) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Room) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Room) GetDeletedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type FindOneByIdReqest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *FindOneByIdReqest) Reset()                    { *m = FindOneByIdReqest{} }
func (m *FindOneByIdReqest) String() string            { return proto.CompactTextString(m) }
func (*FindOneByIdReqest) ProtoMessage()               {}
func (*FindOneByIdReqest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FindOneByIdReqest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ThermostatArray struct {
	Thermostats []*Thermostat `protobuf:"bytes,1,rep,name=thermostats" json:"thermostats,omitempty"`
}

func (m *ThermostatArray) Reset()                    { *m = ThermostatArray{} }
func (m *ThermostatArray) String() string            { return proto.CompactTextString(m) }
func (*ThermostatArray) ProtoMessage()               {}
func (*ThermostatArray) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ThermostatArray) GetThermostats() []*Thermostat {
	if m != nil {
		return m.Thermostats
	}
	return nil
}

func init() {
	proto.RegisterType((*Thermostat)(nil), "com.avegao.iot.temp.service.Thermostat")
	proto.RegisterType((*Zone)(nil), "com.avegao.iot.temp.service.Zone")
	proto.RegisterType((*Room)(nil), "com.avegao.iot.temp.service.Room")
	proto.RegisterType((*FindOneByIdReqest)(nil), "com.avegao.iot.temp.service.FindOneByIdReqest")
	proto.RegisterType((*ThermostatArray)(nil), "com.avegao.iot.temp.service.ThermostatArray")
	proto.RegisterEnum("com.avegao.iot.temp.service.Thermostat_DeviceType", Thermostat_DeviceType_name, Thermostat_DeviceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ThermostatService service

type ThermostatServiceClient interface {
	FindAll(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ThermostatArray, error)
	FindOneById(ctx context.Context, in *FindOneByIdReqest, opts ...grpc.CallOption) (*Thermostat, error)
}

type thermostatServiceClient struct {
	cc *grpc.ClientConn
}

func NewThermostatServiceClient(cc *grpc.ClientConn) ThermostatServiceClient {
	return &thermostatServiceClient{cc}
}

func (c *thermostatServiceClient) FindAll(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ThermostatArray, error) {
	out := new(ThermostatArray)
	err := grpc.Invoke(ctx, "/com.avegao.iot.temp.service.ThermostatService/FindAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thermostatServiceClient) FindOneById(ctx context.Context, in *FindOneByIdReqest, opts ...grpc.CallOption) (*Thermostat, error) {
	out := new(Thermostat)
	err := grpc.Invoke(ctx, "/com.avegao.iot.temp.service.ThermostatService/FindOneById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ThermostatService service

type ThermostatServiceServer interface {
	FindAll(context.Context, *google_protobuf1.Empty) (*ThermostatArray, error)
	FindOneById(context.Context, *FindOneByIdReqest) (*Thermostat, error)
}

func RegisterThermostatServiceServer(s *grpc.Server, srv ThermostatServiceServer) {
	s.RegisterService(&_ThermostatService_serviceDesc, srv)
}

func _ThermostatService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThermostatServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.avegao.iot.temp.service.ThermostatService/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThermostatServiceServer).FindAll(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThermostatService_FindOneById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneByIdReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThermostatServiceServer).FindOneById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.avegao.iot.temp.service.ThermostatService/FindOneById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThermostatServiceServer).FindOneById(ctx, req.(*FindOneByIdReqest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ThermostatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.avegao.iot.temp.service.ThermostatService",
	HandlerType: (*ThermostatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _ThermostatService_FindAll_Handler,
		},
		{
			MethodName: "FindOneById",
			Handler:    _ThermostatService_FindOneById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iotTempService.proto",
}

func init() { proto.RegisterFile("iotTempService.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x8f, 0xd2, 0x40,
	0x18, 0x87, 0x29, 0x14, 0x90, 0x17, 0xb3, 0xb2, 0x13, 0x63, 0x1a, 0xf6, 0x60, 0xad, 0x07, 0x7b,
	0x70, 0xa7, 0x09, 0x46, 0x13, 0xe3, 0xc1, 0x94, 0x2c, 0x1b, 0xb9, 0xc8, 0x3a, 0x8b, 0x07, 0x37,
	0x26, 0x9b, 0x42, 0x5f, 0xa1, 0x09, 0x65, 0xea, 0x74, 0xd8, 0x04, 0x3f, 0x86, 0x9f, 0xcc, 0xbb,
	0x9f, 0xc5, 0xc4, 0xcc, 0xb4, 0x95, 0x5d, 0x49, 0xd8, 0xc2, 0x6d, 0x98, 0xf7, 0xf7, 0xd0, 0x79,
	0x9f, 0xf9, 0x03, 0x8f, 0x23, 0x2e, 0xc7, 0x18, 0x27, 0x97, 0x28, 0x6e, 0xa2, 0x29, 0xd2, 0x44,
	0x70, 0xc9, 0xc9, 0xc9, 0x94, 0xc7, 0x34, 0xb8, 0xc1, 0x59, 0xc0, 0x69, 0xc4, 0x25, 0x95, 0x18,
	0x27, 0x34, 0xcd, 0x22, 0xdd, 0xa7, 0x33, 0xce, 0x67, 0x0b, 0xf4, 0x74, 0x74, 0xb2, 0xfa, 0xe6,
	0xc9, 0x28, 0xc6, 0x54, 0x06, 0x71, 0x92, 0xd1, 0xdd, 0x93, 0xff, 0x03, 0x18, 0x27, 0x72, 0x9d,
	0x15, 0x9d, 0x3f, 0x35, 0x80, 0xf1, 0x1c, 0x45, 0xcc, 0x53, 0x19, 0x48, 0x72, 0x04, 0xd5, 0x28,
	0xb4, 0x0c, 0xdb, 0x70, 0x4d, 0x56, 0x8d, 0x42, 0x42, 0xc0, 0x5c, 0x06, 0x31, 0x5a, 0x55, 0xdb,
	0x70, 0x5b, 0x4c, 0x8f, 0x89, 0x05, 0xcd, 0x20, 0x0c, 0x05, 0xa6, 0xa9, 0x55, 0xd3, 0xd3, 0xc5,
	0x4f, 0x95, 0x4e, 0xb8, 0x90, 0x96, 0x69, 0x1b, 0x6e, 0x9d, 0xe9, 0x31, 0x39, 0x07, 0x53, 0xae,
	0x13, 0xb4, 0xea, 0xb6, 0xe1, 0x1e, 0xf5, 0x7a, 0x74, 0x47, 0x2b, 0x74, 0xb3, 0x10, 0x7a, 0x86,
	0x6a, 0x66, 0xbc, 0x4e, 0x90, 0x69, 0x9e, 0xbc, 0x06, 0xf3, 0x07, 0x5f, 0xa2, 0xd5, 0xb0, 0x0d,
	0xb7, 0xdd, 0x7b, 0xb6, 0xf3, 0x7f, 0xae, 0xf8, 0x12, 0x99, 0x8e, 0x2b, 0x4c, 0x70, 0x1e, 0x5b,
	0xcd, 0x12, 0x18, 0xe3, 0x3c, 0x66, 0x3a, 0x4e, 0xde, 0x02, 0x4c, 0x05, 0x06, 0x12, 0xc3, 0xeb,
	0x40, 0x5a, 0x0f, 0x34, 0xdc, 0xa5, 0x99, 0x48, 0x5a, 0x88, 0xa4, 0xe3, 0xc2, 0x34, 0x6b, 0xe5,
	0x69, 0x5f, 0x2a, 0x74, 0x95, 0x84, 0x05, 0xda, 0xba, 0x1f, 0xcd, 0xd3, 0x19, 0x1a, 0xe2, 0x02,
	0x73, 0x14, 0xee, 0x47, 0xf3, 0xb4, 0x2f, 0x9d, 0x37, 0x00, 0x1b, 0x65, 0xa4, 0x05, 0xf5, 0xd1,
	0xf8, 0xc3, 0x80, 0x75, 0x2a, 0xa4, 0x0d, 0x4d, 0x9f, 0x9d, 0x7d, 0x1e, 0x7e, 0x1c, 0x75, 0x0c,
	0xd2, 0x81, 0x87, 0xcc, 0xbf, 0xbc, 0xe8, 0x0f, 0x18, 0xfb, 0x72, 0x7d, 0x31, 0xec, 0x54, 0x9d,
	0xdf, 0x06, 0x98, 0x4a, 0x57, 0xa9, 0x9d, 0xbf, 0x6b, 0xa5, 0x76, 0xb8, 0x15, 0xf3, 0x70, 0x2b,
	0xf5, 0x7d, 0xac, 0xfc, 0xac, 0x82, 0xa9, 0x76, 0xb5, 0x54, 0x77, 0xc5, 0x09, 0xab, 0xed, 0x77,
	0xc2, 0xee, 0x4a, 0x31, 0x0f, 0x97, 0x52, 0x3f, 0x5c, 0x4a, 0x63, 0x1f, 0x29, 0xcf, 0xe1, 0xf8,
	0x3c, 0x5a, 0x86, 0xa3, 0x25, 0xf6, 0xd7, 0xc3, 0x90, 0xe1, 0x77, 0x4c, 0xb7, 0x2e, 0xbe, 0xf3,
	0x15, 0x1e, 0x6d, 0x6e, 0xa3, 0x2f, 0x44, 0xb0, 0x26, 0x43, 0x68, 0xcb, 0x7f, 0x53, 0xa9, 0x65,
	0xd8, 0x35, 0xb7, 0xdd, 0x7b, 0x51, 0xf2, 0x42, 0xb3, 0xdb, 0x6c, 0xef, 0x97, 0x01, 0xc7, 0x9b,
	0x5a, 0xfe, 0xd8, 0x91, 0x4f, 0xd0, 0x54, 0x0b, 0xf3, 0x17, 0x0b, 0xf2, 0x64, 0xab, 0x95, 0x81,
	0x7a, 0xb4, 0xba, 0x2f, 0x4b, 0x7e, 0x4e, 0xaf, 0xd8, 0xa9, 0x90, 0x39, 0xb4, 0x6f, 0xf5, 0x4a,
	0xe8, 0x4e, 0x7c, 0xcb, 0x4a, 0xb7, 0x6c, 0x77, 0x4e, 0xa5, 0xef, 0x5f, 0xbd, 0x9f, 0x45, 0x72,
	0xbe, 0x9a, 0x28, 0xc4, 0xcb, 0x10, 0x2f, 0xe2, 0xf2, 0x54, 0x21, 0xa7, 0x39, 0xe2, 0x09, 0x4c,
	0xf9, 0x4a, 0x4c, 0xd1, 0x9b, 0x89, 0x64, 0xfa, 0x2e, 0xe2, 0x52, 0x55, 0xf3, 0xe2, 0xa4, 0xa1,
	0x9b, 0x7d, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x94, 0x94, 0x92, 0x71, 0x05, 0x06, 0x00, 0x00,
}
