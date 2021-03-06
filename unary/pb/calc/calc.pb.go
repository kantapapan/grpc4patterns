// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc.proto

// package名

package calc

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

// リクエストで送る値。「１、２」はそのデータの番号。
type SumRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

// レスポンスで送る値
type SumReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumReply) Reset()         { *m = SumReply{} }
func (m *SumReply) String() string { return proto.CompactTextString(m) }
func (*SumReply) ProtoMessage()    {}
func (*SumReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{1}
}

func (m *SumReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumReply.Unmarshal(m, b)
}
func (m *SumReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumReply.Marshal(b, m, deterministic)
}
func (m *SumReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumReply.Merge(m, src)
}
func (m *SumReply) XXX_Size() int {
	return xxx_messageInfo_SumReply.Size(m)
}
func (m *SumReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SumReply.DiscardUnknown(m)
}

var xxx_messageInfo_SumReply proto.InternalMessageInfo

func (m *SumReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calc.SumRequest")
	proto.RegisterType((*SumReply)(nil), "calc.SumReply")
}

func init() { proto.RegisterFile("calc.proto", fileDescriptor_a2b9900dc883ea68) }

var fileDescriptor_a2b9900dc883ea68 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0xcc, 0x49,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x34, 0xb8, 0xb8, 0x82, 0x4b,
	0x73, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x78, 0xb8, 0x18, 0x13, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x83, 0x18, 0x13, 0x41, 0xbc, 0x24, 0x09, 0x26, 0x08, 0x2f, 0x49, 0x49, 0x85,
	0x8b, 0x03, 0xac, 0xb2, 0x20, 0xa7, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31,
	0x3d, 0x15, 0xac, 0x9a, 0x33, 0x08, 0xc6, 0x35, 0x32, 0xe4, 0x62, 0x71, 0x4e, 0xcc, 0x49, 0x16,
	0xd2, 0xe4, 0x62, 0x0e, 0x2e, 0xcd, 0x15, 0x12, 0xd0, 0x03, 0xdb, 0x88, 0xb0, 0x42, 0x8a, 0x0f,
	0x49, 0xa4, 0x20, 0xa7, 0x52, 0x89, 0xc1, 0x49, 0x83, 0x4b, 0x3c, 0x33, 0x5f, 0x2f, 0xbd, 0xa8,
	0x20, 0x59, 0x2f, 0xb5, 0x22, 0x31, 0xb7, 0x20, 0x27, 0xb5, 0x58, 0xaf, 0x38, 0x13, 0x44, 0x3b,
	0x71, 0x07, 0x83, 0xe9, 0x00, 0x90, 0x83, 0x03, 0x18, 0x93, 0xd8, 0xc0, 0x2e, 0x37, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0xcd, 0xed, 0x49, 0xc7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
}

type calcClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcClient(cc grpc.ClientConnInterface) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := c.cc.Invoke(ctx, "/calc.Calc/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServer is the server API for Calc service.
type CalcServer interface {
	Sum(context.Context, *SumRequest) (*SumReply, error)
}

// UnimplementedCalcServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServer struct {
}

func (*UnimplementedCalcServer) Sum(ctx context.Context, req *SumRequest) (*SumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterCalcServer(s *grpc.Server, srv CalcServer) {
	s.RegisterService(&_Calc_serviceDesc, srv)
}

func _Calc_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calc/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Calc_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}
