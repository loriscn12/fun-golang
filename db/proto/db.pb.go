// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/db.proto

package db

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname              string   `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_903744a76531f483, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserRequest) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

type GetUserResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname              string   `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_903744a76531f483, []int{1}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserResponse) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

type AddUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname              string   `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserRequest) Reset()         { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()    {}
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_903744a76531f483, []int{2}
}

func (m *AddUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRequest.Unmarshal(m, b)
}
func (m *AddUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRequest.Marshal(b, m, deterministic)
}
func (m *AddUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRequest.Merge(m, src)
}
func (m *AddUserRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserRequest.Size(m)
}
func (m *AddUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRequest proto.InternalMessageInfo

func (m *AddUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddUserRequest) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

type AddUserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserResponse) Reset()         { *m = AddUserResponse{} }
func (m *AddUserResponse) String() string { return proto.CompactTextString(m) }
func (*AddUserResponse) ProtoMessage()    {}
func (*AddUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_903744a76531f483, []int{3}
}

func (m *AddUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserResponse.Unmarshal(m, b)
}
func (m *AddUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserResponse.Marshal(b, m, deterministic)
}
func (m *AddUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserResponse.Merge(m, src)
}
func (m *AddUserResponse) XXX_Size() int {
	return xxx_messageInfo_AddUserResponse.Size(m)
}
func (m *AddUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserResponse proto.InternalMessageInfo

func (m *AddUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "db.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "db.GetUserResponse")
	proto.RegisterType((*AddUserRequest)(nil), "db.AddUserRequest")
	proto.RegisterType((*AddUserResponse)(nil), "db.AddUserResponse")
}

func init() { proto.RegisterFile("proto/db.proto", fileDescriptor_903744a76531f483) }

var fileDescriptor_903744a76531f483 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x49, 0xd2, 0x03, 0x33, 0x84, 0x98, 0x52, 0x92, 0x94, 0xec, 0xb8, 0xf8, 0xdc,
	0x53, 0x4b, 0x42, 0x8b, 0x53, 0x8b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8,
	0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09,
	0x2e, 0xf6, 0xe2, 0xd2, 0x22, 0xb0, 0x30, 0x13, 0x58, 0x18, 0xc6, 0x55, 0xb2, 0xe7, 0xe2, 0x87,
	0xeb, 0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x25, 0xd1, 0x00, 0x3b, 0x2e, 0x3e, 0xc7, 0x94, 0x14,
	0xf2, 0x1d, 0xa0, 0xc8, 0xc5, 0x0f, 0xd7, 0x0f, 0x75, 0x00, 0x1f, 0x17, 0x53, 0x66, 0x0a, 0x54,
	0x3b, 0x53, 0x66, 0x8a, 0x91, 0x03, 0x17, 0x87, 0x4b, 0x62, 0x49, 0x62, 0x52, 0x62, 0x71, 0xaa,
	0x90, 0x09, 0x17, 0x3b, 0x54, 0xb9, 0x90, 0x90, 0x5e, 0x4a, 0x92, 0x1e, 0xaa, 0xdd, 0x52, 0xc2,
	0x28, 0x62, 0x10, 0xf3, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x01, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xa3, 0x85, 0x74, 0xb9, 0x42, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
}

type databaseClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseClient(cc *grpc.ClientConn) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, "/db.Database/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
type DatabaseServer interface {
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

func _Database_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Database/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Database_AddUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/db.proto",
}
