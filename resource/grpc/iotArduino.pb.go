// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iotArduino.proto

/*
Package arduino_service is a generated protocol buffer package.

It is generated from these files:
	iotArduino.proto

It has these top-level messages:
	ArduinoRequest
	GetTemperatureResponse
	PowerResponse
	ArduinoResponse
*/
package arduino_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ArduinoRequest struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
}

func (m *ArduinoRequest) Reset()                    { *m = ArduinoRequest{} }
func (m *ArduinoRequest) String() string            { return proto.CompactTextString(m) }
func (*ArduinoRequest) ProtoMessage()               {}
func (*ArduinoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ArduinoRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArduinoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArduinoRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type GetTemperatureResponse struct {
	Temperature float32 `protobuf:"fixed32,1,opt,name=temperature" json:"temperature,omitempty"`
}

func (m *GetTemperatureResponse) Reset()                    { *m = GetTemperatureResponse{} }
func (m *GetTemperatureResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTemperatureResponse) ProtoMessage()               {}
func (*GetTemperatureResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetTemperatureResponse) GetTemperature() float32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

type PowerResponse struct {
	Power bool `protobuf:"varint,1,opt,name=power" json:"power,omitempty"`
}

func (m *PowerResponse) Reset()                    { *m = PowerResponse{} }
func (m *PowerResponse) String() string            { return proto.CompactTextString(m) }
func (*PowerResponse) ProtoMessage()               {}
func (*PowerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PowerResponse) GetPower() bool {
	if m != nil {
		return m.Power
	}
	return false
}

type ArduinoResponse struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Auto bool   `protobuf:"varint,4,opt,name=auto" json:"auto,omitempty"`
}

func (m *ArduinoResponse) Reset()                    { *m = ArduinoResponse{} }
func (m *ArduinoResponse) String() string            { return proto.CompactTextString(m) }
func (*ArduinoResponse) ProtoMessage()               {}
func (*ArduinoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ArduinoResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArduinoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArduinoResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ArduinoResponse) GetAuto() bool {
	if m != nil {
		return m.Auto
	}
	return false
}

func init() {
	proto.RegisterType((*ArduinoRequest)(nil), "arduino_service.ArduinoRequest")
	proto.RegisterType((*GetTemperatureResponse)(nil), "arduino_service.GetTemperatureResponse")
	proto.RegisterType((*PowerResponse)(nil), "arduino_service.PowerResponse")
	proto.RegisterType((*ArduinoResponse)(nil), "arduino_service.ArduinoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Arduino service

type ArduinoClient interface {
	GetTemperature(ctx context.Context, in *ArduinoRequest, opts ...grpc.CallOption) (*GetTemperatureResponse, error)
	IsPower(ctx context.Context, in *ArduinoRequest, opts ...grpc.CallOption) (*PowerResponse, error)
	PowerOff(ctx context.Context, in *ArduinoRequest, opts ...grpc.CallOption) (*PowerResponse, error)
	PowerOn(ctx context.Context, in *ArduinoRequest, opts ...grpc.CallOption) (*PowerResponse, error)
}

type arduinoClient struct {
	cc *grpc.ClientConn
}

func NewArduinoClient(cc *grpc.ClientConn) ArduinoClient {
	return &arduinoClient{cc}
}

func (c *arduinoClient) GetTemperature(ctx context.Context, in *ArduinoRequest, opts ...grpc.CallOption) (*GetTemperatureResponse, error) {
	out := new(GetTemperatureResponse)
	err := grpc.Invoke(ctx, "/arduino_service.Arduino/GetTemperature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoClient) IsPower(ctx context.Context, in *ArduinoRequest, opts ...grpc.CallOption) (*PowerResponse, error) {
	out := new(PowerResponse)
	err := grpc.Invoke(ctx, "/arduino_service.Arduino/IsPower", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoClient) PowerOff(ctx context.Context, in *ArduinoRequest, opts ...grpc.CallOption) (*PowerResponse, error) {
	out := new(PowerResponse)
	err := grpc.Invoke(ctx, "/arduino_service.Arduino/PowerOff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoClient) PowerOn(ctx context.Context, in *ArduinoRequest, opts ...grpc.CallOption) (*PowerResponse, error) {
	out := new(PowerResponse)
	err := grpc.Invoke(ctx, "/arduino_service.Arduino/PowerOn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Arduino service

type ArduinoServer interface {
	GetTemperature(context.Context, *ArduinoRequest) (*GetTemperatureResponse, error)
	IsPower(context.Context, *ArduinoRequest) (*PowerResponse, error)
	PowerOff(context.Context, *ArduinoRequest) (*PowerResponse, error)
	PowerOn(context.Context, *ArduinoRequest) (*PowerResponse, error)
}

func RegisterArduinoServer(s *grpc.Server, srv ArduinoServer) {
	s.RegisterService(&_Arduino_serviceDesc, srv)
}

func _Arduino_GetTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArduinoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoServer).GetTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino_service.Arduino/GetTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoServer).GetTemperature(ctx, req.(*ArduinoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arduino_IsPower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArduinoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoServer).IsPower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino_service.Arduino/IsPower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoServer).IsPower(ctx, req.(*ArduinoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arduino_PowerOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArduinoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoServer).PowerOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino_service.Arduino/PowerOff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoServer).PowerOff(ctx, req.(*ArduinoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arduino_PowerOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArduinoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoServer).PowerOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino_service.Arduino/PowerOn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoServer).PowerOn(ctx, req.(*ArduinoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Arduino_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arduino_service.Arduino",
	HandlerType: (*ArduinoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTemperature",
			Handler:    _Arduino_GetTemperature_Handler,
		},
		{
			MethodName: "IsPower",
			Handler:    _Arduino_IsPower_Handler,
		},
		{
			MethodName: "PowerOff",
			Handler:    _Arduino_PowerOff_Handler,
		},
		{
			MethodName: "PowerOn",
			Handler:    _Arduino_PowerOn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iotArduino.proto",
}

func init() { proto.RegisterFile("iotArduino.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0x3a, 0xed, 0x7c, 0xb2, 0x76, 0x04, 0x19, 0xc5, 0x83, 0x96, 0x82, 0xb8, 0x53,
	0x41, 0xbd, 0x79, 0x73, 0x07, 0xc7, 0xe6, 0xc1, 0x12, 0x76, 0xd2, 0x83, 0xc4, 0xf5, 0x4d, 0x02,
	0xb6, 0xa9, 0x69, 0x3a, 0xbf, 0x8f, 0x1f, 0xc7, 0x4f, 0x25, 0x4d, 0x6b, 0xb5, 0x0e, 0x11, 0xdc,
	0xed, 0xe5, 0xcf, 0x2f, 0x3f, 0x92, 0x3f, 0x0f, 0x86, 0x42, 0xea, 0x2b, 0x95, 0x94, 0x22, 0x93,
	0x51, 0xae, 0xa4, 0x96, 0xd4, 0xe3, 0xf5, 0xf1, 0xa1, 0x40, 0xb5, 0x16, 0x4b, 0x0c, 0xaf, 0xc1,
	0x6d, 0x08, 0x86, 0x2f, 0x25, 0x16, 0x9a, 0xba, 0x40, 0x44, 0xe2, 0x5b, 0x81, 0x35, 0x1e, 0x30,
	0x22, 0x12, 0x4a, 0xa1, 0x97, 0xf1, 0x14, 0x7d, 0x12, 0x58, 0xe3, 0x3d, 0x66, 0x66, 0x3a, 0x04,
	0xbb, 0x54, 0xcf, 0xbe, 0x6d, 0xa2, 0x6a, 0x0c, 0x2f, 0x61, 0x34, 0x45, 0xbd, 0xc0, 0x34, 0x47,
	0xc5, 0x75, 0xa9, 0x90, 0x61, 0x91, 0xcb, 0xac, 0x40, 0x1a, 0xc0, 0xbe, 0xfe, 0x8a, 0x8d, 0x98,
	0xb0, 0xef, 0x51, 0x78, 0x02, 0x83, 0x58, 0xbe, 0xa2, 0x6a, 0xaf, 0x1c, 0xc0, 0x4e, 0x5e, 0x05,
	0x06, 0xee, 0xb3, 0xfa, 0x10, 0xde, 0x83, 0xd7, 0x3e, 0xb5, 0x01, 0xff, 0xf5, 0xd6, 0x8a, 0xe2,
	0xa5, 0x96, 0x7e, 0xcf, 0xd8, 0xcd, 0x7c, 0xfe, 0x4e, 0xc0, 0x69, 0xec, 0xf4, 0x0e, 0xdc, 0xee,
	0x5f, 0xe8, 0x71, 0xf4, 0xa3, 0xb7, 0xa8, 0x5b, 0xda, 0xe1, 0xe9, 0x06, 0xf0, 0x4b, 0x1b, 0x73,
	0x70, 0x66, 0x85, 0xf9, 0xed, 0xdf, 0xd2, 0xa3, 0x0d, 0xa0, 0x5b, 0xd3, 0x0d, 0xf4, 0x4d, 0x70,
	0xbb, 0x5a, 0x6d, 0x2f, 0x9b, 0x83, 0x53, 0xcb, 0xb2, 0xad, 0x5d, 0x93, 0x33, 0x18, 0x2d, 0x65,
	0x1a, 0xf1, 0x35, 0x3e, 0x71, 0x19, 0x09, 0xa9, 0x3f, 0xf9, 0x89, 0x37, 0x6b, 0x37, 0x32, 0xae,
	0x16, 0x32, 0xb6, 0xde, 0x88, 0xcd, 0x16, 0xd3, 0xc7, 0x5d, 0xb3, 0x9f, 0x17, 0x1f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0xbd, 0xba, 0x34, 0xb3, 0x02, 0x00, 0x00,
}