// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repository/repository.proto

package repository // import "github.com/ProtocolONE/payone-repository/pkg/proto/repository"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import billing "github.com/ProtocolONE/payone-repository/pkg/proto/billing"

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

type Result struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_repository_7471a48105177bc5, []int{0}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

type ConvertRequest struct {
	CurrencyFrom         int32    `protobuf:"varint,1,opt,name=currency_from,json=currencyFrom,proto3" json:"currency_from,omitempty"`
	CurrencyTo           int32    `protobuf:"varint,2,opt,name=currency_to,json=currencyTo,proto3" json:"currency_to,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvertRequest) Reset()         { *m = ConvertRequest{} }
func (m *ConvertRequest) String() string { return proto.CompactTextString(m) }
func (*ConvertRequest) ProtoMessage()    {}
func (*ConvertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_repository_7471a48105177bc5, []int{1}
}
func (m *ConvertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertRequest.Unmarshal(m, b)
}
func (m *ConvertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertRequest.Marshal(b, m, deterministic)
}
func (dst *ConvertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertRequest.Merge(dst, src)
}
func (m *ConvertRequest) XXX_Size() int {
	return xxx_messageInfo_ConvertRequest.Size(m)
}
func (m *ConvertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertRequest proto.InternalMessageInfo

func (m *ConvertRequest) GetCurrencyFrom() int32 {
	if m != nil {
		return m.CurrencyFrom
	}
	return 0
}

func (m *ConvertRequest) GetCurrencyTo() int32 {
	if m != nil {
		return m.CurrencyTo
	}
	return 0
}

func (m *ConvertRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ConvertResponse struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvertResponse) Reset()         { *m = ConvertResponse{} }
func (m *ConvertResponse) String() string { return proto.CompactTextString(m) }
func (*ConvertResponse) ProtoMessage()    {}
func (*ConvertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_repository_7471a48105177bc5, []int{2}
}
func (m *ConvertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertResponse.Unmarshal(m, b)
}
func (m *ConvertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertResponse.Marshal(b, m, deterministic)
}
func (dst *ConvertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertResponse.Merge(dst, src)
}
func (m *ConvertResponse) XXX_Size() int {
	return xxx_messageInfo_ConvertResponse.Size(m)
}
func (m *ConvertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertResponse proto.InternalMessageInfo

func (m *ConvertResponse) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*Result)(nil), "repository.Result")
	proto.RegisterType((*ConvertRequest)(nil), "repository.ConvertRequest")
	proto.RegisterType((*ConvertResponse)(nil), "repository.ConvertResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RepositoryClient is the client API for Repository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryClient interface {
	UpdateOrder(ctx context.Context, in *billing.Order, opts ...grpc.CallOption) (*Result, error)
	ConvertAmount(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error)
	GetConvertRate(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error)
	UpdateMerchant(ctx context.Context, in *billing.Merchant, opts ...grpc.CallOption) (*Result, error)
}

type repositoryClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryClient(cc *grpc.ClientConn) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) UpdateOrder(ctx context.Context, in *billing.Order, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/repository.Repository/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) ConvertAmount(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error) {
	out := new(ConvertResponse)
	err := c.cc.Invoke(ctx, "/repository.Repository/ConvertAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) GetConvertRate(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error) {
	out := new(ConvertResponse)
	err := c.cc.Invoke(ctx, "/repository.Repository/GetConvertRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) UpdateMerchant(ctx context.Context, in *billing.Merchant, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/repository.Repository/UpdateMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServer is the server API for Repository service.
type RepositoryServer interface {
	UpdateOrder(context.Context, *billing.Order) (*Result, error)
	ConvertAmount(context.Context, *ConvertRequest) (*ConvertResponse, error)
	GetConvertRate(context.Context, *ConvertRequest) (*ConvertResponse, error)
	UpdateMerchant(context.Context, *billing.Merchant) (*Result, error)
}

func RegisterRepositoryServer(s *grpc.Server, srv RepositoryServer) {
	s.RegisterService(&_Repository_serviceDesc, srv)
}

func _Repository_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.Repository/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).UpdateOrder(ctx, req.(*billing.Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_ConvertAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).ConvertAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.Repository/ConvertAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).ConvertAmount(ctx, req.(*ConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_GetConvertRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).GetConvertRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.Repository/GetConvertRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).GetConvertRate(ctx, req.(*ConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_UpdateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.Merchant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).UpdateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.Repository/UpdateMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).UpdateMerchant(ctx, req.(*billing.Merchant))
	}
	return interceptor(ctx, in, info, handler)
}

var _Repository_serviceDesc = grpc.ServiceDesc{
	ServiceName: "repository.Repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateOrder",
			Handler:    _Repository_UpdateOrder_Handler,
		},
		{
			MethodName: "ConvertAmount",
			Handler:    _Repository_ConvertAmount_Handler,
		},
		{
			MethodName: "GetConvertRate",
			Handler:    _Repository_GetConvertRate_Handler,
		},
		{
			MethodName: "UpdateMerchant",
			Handler:    _Repository_UpdateMerchant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repository/repository.proto",
}

func init() {
	proto.RegisterFile("repository/repository.proto", fileDescriptor_repository_7471a48105177bc5)
}

var fileDescriptor_repository_7471a48105177bc5 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x9b, 0x8a, 0x45, 0xa6, 0x36, 0xe2, 0x82, 0x52, 0xd2, 0x83, 0x25, 0x5e, 0xea, 0xc1,
	0x04, 0xec, 0xc9, 0x83, 0x88, 0x8a, 0x0a, 0x8a, 0x56, 0x82, 0x5e, 0xbc, 0x48, 0x9a, 0x8e, 0x6d,
	0x30, 0xd9, 0x59, 0x37, 0x13, 0xa1, 0x0f, 0xe2, 0xfb, 0x8a, 0xf9, 0xd7, 0x15, 0x7a, 0xf3, 0xb4,
	0x3b, 0xdf, 0x7e, 0xcc, 0xec, 0x6f, 0x66, 0x60, 0xa0, 0x51, 0x51, 0x16, 0x33, 0xe9, 0xa5, 0xbf,
	0xba, 0x7a, 0x4a, 0x13, 0x93, 0x80, 0x95, 0xe2, 0xec, 0x4d, 0xe3, 0x24, 0x89, 0xe5, 0xdc, 0xaf,
	0xce, 0xd2, 0xe2, 0x6e, 0x41, 0x27, 0xc0, 0x2c, 0x4f, 0xd8, 0x95, 0x60, 0x5f, 0x91, 0xfc, 0x42,
	0xcd, 0x01, 0x7e, 0xe6, 0x98, 0xb1, 0x38, 0x84, 0x5e, 0x94, 0x6b, 0x8d, 0x32, 0x5a, 0xbe, 0xbd,
	0x6b, 0x4a, 0xfb, 0xd6, 0xd0, 0x1a, 0x6d, 0x06, 0xdb, 0xb5, 0x78, 0xa3, 0x29, 0x15, 0x07, 0xd0,
	0x6d, 0x4c, 0x4c, 0xfd, 0x76, 0x61, 0x81, 0x5a, 0x7a, 0x26, 0xb1, 0x0f, 0x9d, 0x30, 0xa5, 0x5c,
	0x72, 0x7f, 0x63, 0x68, 0x8d, 0xac, 0xa0, 0x8a, 0xdc, 0x23, 0xd8, 0x69, 0xea, 0x65, 0x8a, 0x64,
	0x86, 0x86, 0xd5, 0x32, 0xad, 0x27, 0xdf, 0x6d, 0x80, 0xa0, 0x41, 0x11, 0x63, 0xe8, 0xbe, 0xa8,
	0x59, 0xc8, 0x38, 0xd1, 0x33, 0xd4, 0xc2, 0xf6, 0x6a, 0xa4, 0x22, 0x76, 0x84, 0x67, 0x34, 0xa2,
	0x82, 0x6b, 0x89, 0x3b, 0xe8, 0x55, 0xe5, 0x2e, 0x8a, 0xa4, 0xc2, 0x31, 0x6d, 0x7f, 0xc9, 0x9d,
	0xc1, 0xda, 0xb7, 0xf2, 0x97, 0x6e, 0x4b, 0xdc, 0x83, 0x7d, 0x8b, 0x5c, 0xeb, 0x21, 0xe3, 0x7f,
	0x92, 0x9d, 0x82, 0x5d, 0xd2, 0x3c, 0xa0, 0x8e, 0x16, 0xa1, 0x64, 0xb1, 0xdb, 0x00, 0xd5, 0xd2,
	0x7a, 0xa6, 0xcb, 0xf3, 0xd7, 0xb3, 0x79, 0xcc, 0x8b, 0x7c, 0xea, 0x45, 0x94, 0xfa, 0x4f, 0xbf,
	0x03, 0x8d, 0x28, 0x99, 0x3c, 0x5e, 0xfb, 0x2a, 0x5c, 0x92, 0xc4, 0x63, 0x63, 0x39, 0xd4, 0xc7,
	0xdc, 0x2f, 0x46, 0x6e, 0xac, 0xc9, 0xb4, 0x53, 0x28, 0xe3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7f, 0x69, 0x83, 0x0c, 0x46, 0x02, 0x00, 0x00,
}
