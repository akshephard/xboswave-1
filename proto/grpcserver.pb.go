// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcserver.proto

package xbospb

/*
framing messages for exposing GRPC servers over WAVEMQ
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type GRPCServerMessage struct {
	UnaryCall            *UnaryCall         `protobuf:"bytes,1,opt,name=UnaryCall,proto3" json:"UnaryCall,omitempty"`
	UnaryResponse        *UnaryResponse     `protobuf:"bytes,2,opt,name=UnaryResponse,proto3" json:"UnaryResponse,omitempty"`
	StreamingCall        *StreamingCall     `protobuf:"bytes,3,opt,name=StreamingCall,proto3" json:"StreamingCall,omitempty"`
	StreamingResponse    *StreamingResponse `protobuf:"bytes,4,opt,name=StreamingResponse,proto3" json:"StreamingResponse,omitempty"`
	Ping                 *Ping              `protobuf:"bytes,5,opt,name=Ping,proto3" json:"Ping,omitempty"`
	Pong                 *Pong              `protobuf:"bytes,6,opt,name=Pong,proto3" json:"Pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GRPCServerMessage) Reset()         { *m = GRPCServerMessage{} }
func (m *GRPCServerMessage) String() string { return proto.CompactTextString(m) }
func (*GRPCServerMessage) ProtoMessage()    {}
func (*GRPCServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcserver_7ff8f6f8f3d8c463, []int{0}
}
func (m *GRPCServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GRPCServerMessage.Unmarshal(m, b)
}
func (m *GRPCServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GRPCServerMessage.Marshal(b, m, deterministic)
}
func (dst *GRPCServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GRPCServerMessage.Merge(dst, src)
}
func (m *GRPCServerMessage) XXX_Size() int {
	return xxx_messageInfo_GRPCServerMessage.Size(m)
}
func (m *GRPCServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GRPCServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GRPCServerMessage proto.InternalMessageInfo

func (m *GRPCServerMessage) GetUnaryCall() *UnaryCall {
	if m != nil {
		return m.UnaryCall
	}
	return nil
}

func (m *GRPCServerMessage) GetUnaryResponse() *UnaryResponse {
	if m != nil {
		return m.UnaryResponse
	}
	return nil
}

func (m *GRPCServerMessage) GetStreamingCall() *StreamingCall {
	if m != nil {
		return m.StreamingCall
	}
	return nil
}

func (m *GRPCServerMessage) GetStreamingResponse() *StreamingResponse {
	if m != nil {
		return m.StreamingResponse
	}
	return nil
}

func (m *GRPCServerMessage) GetPing() *Ping {
	if m != nil {
		return m.Ping
	}
	return nil
}

func (m *GRPCServerMessage) GetPong() *Pong {
	if m != nil {
		return m.Pong
	}
	return nil
}

type Ping struct {
	ActiveQueries        []int64  `protobuf:"varint,1,rep,packed,name=active_queries,json=activeQueries,proto3" json:"active_queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcserver_7ff8f6f8f3d8c463, []int{1}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (dst *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(dst, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetActiveQueries() []int64 {
	if m != nil {
		return m.ActiveQueries
	}
	return nil
}

type Pong struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcserver_7ff8f6f8f3d8c463, []int{2}
}
func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (dst *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(dst, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

type UnaryCall struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	QueryId              int64    `protobuf:"varint,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	Payload              *any.Any `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnaryCall) Reset()         { *m = UnaryCall{} }
func (m *UnaryCall) String() string { return proto.CompactTextString(m) }
func (*UnaryCall) ProtoMessage()    {}
func (*UnaryCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcserver_7ff8f6f8f3d8c463, []int{3}
}
func (m *UnaryCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnaryCall.Unmarshal(m, b)
}
func (m *UnaryCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnaryCall.Marshal(b, m, deterministic)
}
func (dst *UnaryCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnaryCall.Merge(dst, src)
}
func (m *UnaryCall) XXX_Size() int {
	return xxx_messageInfo_UnaryCall.Size(m)
}
func (m *UnaryCall) XXX_DiscardUnknown() {
	xxx_messageInfo_UnaryCall.DiscardUnknown(m)
}

var xxx_messageInfo_UnaryCall proto.InternalMessageInfo

func (m *UnaryCall) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *UnaryCall) GetQueryId() int64 {
	if m != nil {
		return m.QueryId
	}
	return 0
}

func (m *UnaryCall) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type UnaryResponse struct {
	QueryId              int64    `protobuf:"varint,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Payload              *any.Any `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnaryResponse) Reset()         { *m = UnaryResponse{} }
func (m *UnaryResponse) String() string { return proto.CompactTextString(m) }
func (*UnaryResponse) ProtoMessage()    {}
func (*UnaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcserver_7ff8f6f8f3d8c463, []int{4}
}
func (m *UnaryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnaryResponse.Unmarshal(m, b)
}
func (m *UnaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnaryResponse.Marshal(b, m, deterministic)
}
func (dst *UnaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnaryResponse.Merge(dst, src)
}
func (m *UnaryResponse) XXX_Size() int {
	return xxx_messageInfo_UnaryResponse.Size(m)
}
func (m *UnaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnaryResponse proto.InternalMessageInfo

func (m *UnaryResponse) GetQueryId() int64 {
	if m != nil {
		return m.QueryId
	}
	return 0
}

func (m *UnaryResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UnaryResponse) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type StreamingCall struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	QueryId              int64    `protobuf:"varint,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	Payload              *any.Any `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingCall) Reset()         { *m = StreamingCall{} }
func (m *StreamingCall) String() string { return proto.CompactTextString(m) }
func (*StreamingCall) ProtoMessage()    {}
func (*StreamingCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcserver_7ff8f6f8f3d8c463, []int{5}
}
func (m *StreamingCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingCall.Unmarshal(m, b)
}
func (m *StreamingCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingCall.Marshal(b, m, deterministic)
}
func (dst *StreamingCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingCall.Merge(dst, src)
}
func (m *StreamingCall) XXX_Size() int {
	return xxx_messageInfo_StreamingCall.Size(m)
}
func (m *StreamingCall) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingCall.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingCall proto.InternalMessageInfo

func (m *StreamingCall) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *StreamingCall) GetQueryId() int64 {
	if m != nil {
		return m.QueryId
	}
	return 0
}

func (m *StreamingCall) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type StreamingResponse struct {
	QueryId int64  `protobuf:"varint,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// true if the stream is done
	Finished             bool     `protobuf:"varint,3,opt,name=finished,proto3" json:"finished,omitempty"`
	Payload              *any.Any `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcserver_7ff8f6f8f3d8c463, []int{6}
}
func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (dst *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(dst, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetQueryId() int64 {
	if m != nil {
		return m.QueryId
	}
	return 0
}

func (m *StreamingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *StreamingResponse) GetFinished() bool {
	if m != nil {
		return m.Finished
	}
	return false
}

func (m *StreamingResponse) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type TestParams struct {
	X                    string   `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestParams) Reset()         { *m = TestParams{} }
func (m *TestParams) String() string { return proto.CompactTextString(m) }
func (*TestParams) ProtoMessage()    {}
func (*TestParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcserver_7ff8f6f8f3d8c463, []int{7}
}
func (m *TestParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestParams.Unmarshal(m, b)
}
func (m *TestParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestParams.Marshal(b, m, deterministic)
}
func (dst *TestParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestParams.Merge(dst, src)
}
func (m *TestParams) XXX_Size() int {
	return xxx_messageInfo_TestParams.Size(m)
}
func (m *TestParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TestParams.DiscardUnknown(m)
}

var xxx_messageInfo_TestParams proto.InternalMessageInfo

func (m *TestParams) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

type TestResponse struct {
	X                    string   `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcserver_7ff8f6f8f3d8c463, []int{8}
}
func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (dst *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(dst, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func init() {
	proto.RegisterType((*GRPCServerMessage)(nil), "xbospb.GRPCServerMessage")
	proto.RegisterType((*Ping)(nil), "xbospb.Ping")
	proto.RegisterType((*Pong)(nil), "xbospb.Pong")
	proto.RegisterType((*UnaryCall)(nil), "xbospb.UnaryCall")
	proto.RegisterType((*UnaryResponse)(nil), "xbospb.UnaryResponse")
	proto.RegisterType((*StreamingCall)(nil), "xbospb.StreamingCall")
	proto.RegisterType((*StreamingResponse)(nil), "xbospb.StreamingResponse")
	proto.RegisterType((*TestParams)(nil), "xbospb.TestParams")
	proto.RegisterType((*TestResponse)(nil), "xbospb.TestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	TestUnary(ctx context.Context, in *TestParams, opts ...grpc.CallOption) (*TestResponse, error)
	TestStream(ctx context.Context, in *TestParams, opts ...grpc.CallOption) (Test_TestStreamClient, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) TestUnary(ctx context.Context, in *TestParams, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/xbospb.Test/TestUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) TestStream(ctx context.Context, in *TestParams, opts ...grpc.CallOption) (Test_TestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[0], "/xbospb.Test/TestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testTestStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Test_TestStreamClient interface {
	Recv() (*TestResponse, error)
	grpc.ClientStream
}

type testTestStreamClient struct {
	grpc.ClientStream
}

func (x *testTestStreamClient) Recv() (*TestResponse, error) {
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	TestUnary(context.Context, *TestParams) (*TestResponse, error)
	TestStream(*TestParams, Test_TestStreamServer) error
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_TestUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).TestUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xbospb.Test/TestUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).TestUnary(ctx, req.(*TestParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_TestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TestParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServer).TestStream(m, &testTestStreamServer{stream})
}

type Test_TestStreamServer interface {
	Send(*TestResponse) error
	grpc.ServerStream
}

type testTestStreamServer struct {
	grpc.ServerStream
}

func (x *testTestStreamServer) Send(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xbospb.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestUnary",
			Handler:    _Test_TestUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestStream",
			Handler:       _Test_TestStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpcserver.proto",
}

func init() { proto.RegisterFile("grpcserver.proto", fileDescriptor_grpcserver_7ff8f6f8f3d8c463) }

var fileDescriptor_grpcserver_7ff8f6f8f3d8c463 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x6b, 0x97, 0x35, 0x8f, 0x0e, 0x51, 0xab, 0xa0, 0x34, 0xe2, 0x50, 0x45, 0x42, 0xda,
	0x85, 0x14, 0x8d, 0x03, 0x12, 0x9c, 0xd0, 0x0e, 0x13, 0x07, 0xa4, 0xe2, 0xc1, 0x79, 0x72, 0x16,
	0xcf, 0x8b, 0x94, 0xda, 0xc1, 0xce, 0xa6, 0x46, 0x7c, 0x09, 0xbe, 0x17, 0x5f, 0x0a, 0xe5, 0xb9,
	0xc9, 0xea, 0x96, 0x03, 0xbd, 0x70, 0xb2, 0xde, 0xfb, 0xfd, 0x79, 0xff, 0x0c, 0xcf, 0xa5, 0xa9,
	0x6e, 0xac, 0x30, 0x0f, 0xc2, 0xa4, 0x95, 0xd1, 0xb5, 0xa6, 0xc1, 0x3a, 0xd3, 0xb6, 0xca, 0xe2,
	0x99, 0xd4, 0x5a, 0x96, 0x62, 0x81, 0xd9, 0xec, 0xfe, 0x76, 0xc1, 0x55, 0xe3, 0x28, 0xc9, 0xef,
	0x23, 0x98, 0x5c, 0xb2, 0xe5, 0xc5, 0x15, 0xea, 0xbe, 0x08, 0x6b, 0xb9, 0x14, 0x74, 0x01, 0xe1,
	0x77, 0xc5, 0x4d, 0x73, 0xc1, 0xcb, 0x32, 0x22, 0x73, 0x72, 0xf6, 0xf4, 0x7c, 0x92, 0x3a, 0xb3,
	0xb4, 0x07, 0xd8, 0x23, 0x87, 0x7e, 0x84, 0x53, 0x0c, 0x98, 0xb0, 0x95, 0x56, 0x56, 0x44, 0x47,
	0x28, 0x7a, 0xe1, 0x89, 0x3a, 0x90, 0xf9, 0xdc, 0x56, 0x7c, 0x55, 0x1b, 0xc1, 0x57, 0x85, 0x92,
	0x58, 0x71, 0xe0, 0x8b, 0x3d, 0x90, 0xf9, 0x5c, 0x7a, 0x09, 0x93, 0x3e, 0xd1, 0x57, 0x1f, 0xa2,
	0xc1, 0x6c, 0xcf, 0xa0, 0xef, 0x60, 0x5f, 0x43, 0xe7, 0x30, 0x5c, 0x16, 0x4a, 0x46, 0xc7, 0xa8,
	0x1d, 0x77, 0xda, 0x36, 0xc7, 0x10, 0x41, 0x86, 0x56, 0x32, 0x0a, 0x76, 0x18, 0x1a, 0x19, 0x5a,
	0xc9, 0xe4, 0x8d, 0xf3, 0xa0, 0xaf, 0xe1, 0x19, 0xbf, 0xa9, 0x8b, 0x07, 0x71, 0xfd, 0xe3, 0x5e,
	0x98, 0x42, 0xd8, 0x88, 0xcc, 0x07, 0x67, 0x03, 0x76, 0xea, 0xb2, 0x5f, 0x5d, 0x32, 0x09, 0x9c,
	0x61, 0xa2, 0xb6, 0xd6, 0x4d, 0x5f, 0x42, 0xb0, 0x12, 0xf5, 0x9d, 0xce, 0x71, 0xf1, 0x21, 0xdb,
	0x44, 0x74, 0x06, 0xa3, 0xd6, 0xac, 0xb9, 0x2e, 0x72, 0xdc, 0xee, 0x80, 0x9d, 0x60, 0xfc, 0x39,
	0xa7, 0x29, 0x9c, 0x54, 0xbc, 0x29, 0x35, 0xcf, 0x37, 0xab, 0x9b, 0xa6, 0xee, 0xe2, 0x69, 0x77,
	0xf1, 0xf4, 0x93, 0x6a, 0x58, 0x47, 0x4a, 0xaa, 0x9d, 0x6b, 0x79, 0xde, 0xc4, 0xf7, 0x9e, 0xc2,
	0xb1, 0x30, 0x46, 0x1b, 0xac, 0x19, 0x32, 0x17, 0x1c, 0x5c, 0xd1, 0xec, 0x9c, 0xf8, 0x7f, 0x4c,
	0xf9, 0x8b, 0xfc, 0xe5, 0x6b, 0x1c, 0x3e, 0x6a, 0x0c, 0xa3, 0xdb, 0x42, 0x15, 0xf6, 0x4e, 0xb8,
	0xba, 0x23, 0xd6, 0xc7, 0xdb, 0x2d, 0x0d, 0xff, 0xa5, 0xa5, 0x18, 0xe0, 0x9b, 0xb0, 0xf5, 0x92,
	0x1b, 0xbe, 0xb2, 0x74, 0x0c, 0x64, 0xbd, 0x19, 0x9f, 0xac, 0x93, 0x57, 0x30, 0x6e, 0xb1, 0xbe,
	0x51, 0x0f, 0x3d, 0xff, 0x09, 0xc3, 0x16, 0xa5, 0xef, 0x21, 0x6c, 0x5f, 0x3c, 0x1f, 0xa5, 0xdd,
	0x17, 0x7c, 0x34, 0x8d, 0xa7, 0xdb, 0xb9, 0xce, 0x2c, 0x79, 0x42, 0x3f, 0xb8, 0xd2, 0x6e, 0x21,
	0x87, 0x28, 0xdf, 0x92, 0x2c, 0xc0, 0x69, 0xde, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x29,
	0xe4, 0x28, 0x62, 0x04, 0x00, 0x00,
}
