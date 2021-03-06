// Code generated by protoc-gen-go. DO NOT EDIT.
// source: check.proto

package check

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CheckReq struct {
	Book                 string   `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckReq) Reset()         { *m = CheckReq{} }
func (m *CheckReq) String() string { return proto.CompactTextString(m) }
func (*CheckReq) ProtoMessage()    {}
func (*CheckReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{0}
}

func (m *CheckReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckReq.Unmarshal(m, b)
}
func (m *CheckReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckReq.Marshal(b, m, deterministic)
}
func (m *CheckReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckReq.Merge(m, src)
}
func (m *CheckReq) XXX_Size() int {
	return xxx_messageInfo_CheckReq.Size(m)
}
func (m *CheckReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckReq proto.InternalMessageInfo

func (m *CheckReq) GetBook() string {
	if m != nil {
		return m.Book
	}
	return ""
}

type CheckResp struct {
	Found                bool     `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Price                int64    `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResp) Reset()         { *m = CheckResp{} }
func (m *CheckResp) String() string { return proto.CompactTextString(m) }
func (*CheckResp) ProtoMessage()    {}
func (*CheckResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{1}
}

func (m *CheckResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResp.Unmarshal(m, b)
}
func (m *CheckResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResp.Marshal(b, m, deterministic)
}
func (m *CheckResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResp.Merge(m, src)
}
func (m *CheckResp) XXX_Size() int {
	return xxx_messageInfo_CheckResp.Size(m)
}
func (m *CheckResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResp proto.InternalMessageInfo

func (m *CheckResp) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *CheckResp) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*CheckReq)(nil), "check.checkReq")
	proto.RegisterType((*CheckResp)(nil), "check.checkResp")
}

func init() { proto.RegisterFile("check.proto", fileDescriptor_d8d3c606fb107336) }

var fileDescriptor_d8d3c606fb107336 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xce, 0x48, 0x4d,
	0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xe4, 0xb8, 0x38, 0xc0,
	0x8c, 0xa0, 0xd4, 0x42, 0x21, 0x21, 0x2e, 0x96, 0xa4, 0xfc, 0xfc, 0x6c, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x9c, 0x8b, 0x13, 0x2a, 0x5f, 0x5c, 0x20, 0x24, 0xc2, 0xc5,
	0x9a, 0x96, 0x5f, 0x9a, 0x97, 0x02, 0x56, 0xc1, 0x11, 0x04, 0xe1, 0x80, 0x44, 0x0b, 0x8a, 0x32,
	0x93, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x23, 0x53, 0x2e, 0x76, 0xb0,
	0xc6, 0xd4, 0x22, 0x21, 0x2d, 0x2e, 0x88, 0x65, 0x42, 0xfc, 0x7a, 0x10, 0x17, 0xc0, 0x6c, 0x94,
	0x12, 0x40, 0x15, 0x28, 0x2e, 0x48, 0x62, 0x03, 0xbb, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x6e, 0x6f, 0xa7, 0x1d, 0xac, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheckerClient is the client API for Checker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckerClient interface {
	Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckResp, error)
}

type checkerClient struct {
	cc *grpc.ClientConn
}

func NewCheckerClient(cc *grpc.ClientConn) CheckerClient {
	return &checkerClient{cc}
}

func (c *checkerClient) Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckResp, error) {
	out := new(CheckResp)
	err := c.cc.Invoke(ctx, "/check.checker/check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckerServer is the server API for Checker service.
type CheckerServer interface {
	Check(context.Context, *CheckReq) (*CheckResp, error)
}

// UnimplementedCheckerServer can be embedded to have forward compatible implementations.
type UnimplementedCheckerServer struct {
}

func (*UnimplementedCheckerServer) Check(ctx context.Context, req *CheckReq) (*CheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterCheckerServer(s *grpc.Server, srv CheckerServer) {
	s.RegisterService(&_Checker_serviceDesc, srv)
}

func _Checker_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/check.checker/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Check(ctx, req.(*CheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Checker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "check.checker",
	HandlerType: (*CheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "check",
			Handler:    _Checker_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "check.proto",
}
