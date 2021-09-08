// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kvm.proto

package protobuf

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

type AgentIdentity struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentIdentity) Reset()         { *m = AgentIdentity{} }
func (m *AgentIdentity) String() string { return proto.CompactTextString(m) }
func (*AgentIdentity) ProtoMessage()    {}
func (*AgentIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_5863643a03e0ce6f, []int{0}
}

func (m *AgentIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentIdentity.Unmarshal(m, b)
}
func (m *AgentIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentIdentity.Marshal(b, m, deterministic)
}
func (m *AgentIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentIdentity.Merge(m, src)
}
func (m *AgentIdentity) XXX_Size() int {
	return xxx_messageInfo_AgentIdentity.Size(m)
}
func (m *AgentIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_AgentIdentity proto.InternalMessageInfo

func (m *AgentIdentity) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

type Status struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_5863643a03e0ce6f, []int{1}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Status) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type PolicyData struct {
	PolicyName           string   `protobuf:"bytes,1,opt,name=policyName,proto3" json:"policyName,omitempty"`
	PolicyData           []byte   `protobuf:"bytes,2,opt,name=policyData,proto3" json:"policyData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyData) Reset()         { *m = PolicyData{} }
func (m *PolicyData) String() string { return proto.CompactTextString(m) }
func (*PolicyData) ProtoMessage()    {}
func (*PolicyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5863643a03e0ce6f, []int{2}
}

func (m *PolicyData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyData.Unmarshal(m, b)
}
func (m *PolicyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyData.Marshal(b, m, deterministic)
}
func (m *PolicyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyData.Merge(m, src)
}
func (m *PolicyData) XXX_Size() int {
	return xxx_messageInfo_PolicyData.Size(m)
}
func (m *PolicyData) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyData.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyData proto.InternalMessageInfo

func (m *PolicyData) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *PolicyData) GetPolicyData() []byte {
	if m != nil {
		return m.PolicyData
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentIdentity)(nil), "kvm.agentIdentity")
	proto.RegisterType((*Status)(nil), "kvm.status")
	proto.RegisterType((*PolicyData)(nil), "kvm.policyData")
}

func init() { proto.RegisterFile("kvm.proto", fileDescriptor_5863643a03e0ce6f) }

var fileDescriptor_5863643a03e0ce6f = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0xc5, 0x62, 0xc7, 0x8a, 0xb0, 0xa0, 0x94, 0x1e, 0x44, 0x72, 0x0a, 0x0a, 0x89,
	0xe8, 0xc5, 0xab, 0xd2, 0x8b, 0xd4, 0x88, 0x44, 0xe8, 0xc1, 0xdb, 0x26, 0x1d, 0xe3, 0x12, 0x37,
	0x5b, 0x66, 0x27, 0x81, 0xfe, 0x7b, 0xe9, 0x90, 0xb4, 0xcd, 0xed, 0xcd, 0x37, 0x8f, 0xb7, 0xfb,
	0x06, 0x26, 0x55, 0x6b, 0xe3, 0x0d, 0x39, 0x76, 0x6a, 0x54, 0xb5, 0x36, 0xbc, 0x87, 0x0b, 0x5d,
	0x62, 0xcd, 0x6f, 0x6b, 0xac, 0xd9, 0xf0, 0x56, 0xcd, 0xe1, 0xcc, 0x74, 0x7a, 0x16, 0xdc, 0x06,
	0xd1, 0x24, 0xdb, 0xcf, 0xe1, 0x02, 0xc6, 0x9e, 0x35, 0x37, 0x5e, 0x5d, 0xf7, 0x4a, 0x3c, 0xa7,
	0x59, 0xcf, 0x43, 0x98, 0x22, 0x91, 0xa3, 0x14, 0xbd, 0xd7, 0x25, 0xce, 0x4e, 0x24, 0x61, 0xc0,
	0xc2, 0x77, 0x80, 0x8d, 0xfb, 0x33, 0xc5, 0x76, 0xa1, 0x59, 0xab, 0x9b, 0x7e, 0xfa, 0xd0, 0x16,
	0xbb, 0x17, 0x8f, 0xc8, 0x61, 0xbf, 0x73, 0x4b, 0xde, 0x34, 0x3b, 0x22, 0x8f, 0x0e, 0x46, 0xcb,
	0x55, 0xaa, 0x9e, 0xe1, 0x8a, 0xb0, 0x34, 0x9e, 0x91, 0x5e, 0x06, 0x7d, 0x54, 0xbc, 0x6b, 0x3c,
	0xe8, 0x38, 0x3f, 0x17, 0xd6, 0x7d, 0x39, 0x06, 0xf0, 0x58, 0xaf, 0x3f, 0x25, 0x52, 0x5d, 0xca,
	0xea, 0x90, 0x3f, 0xf0, 0x46, 0xc1, 0x43, 0xf0, 0x7a, 0xf7, 0x1d, 0x95, 0x86, 0x7f, 0x9b, 0x3c,
	0x2e, 0x9c, 0x4d, 0xaa, 0x26, 0x47, 0x4d, 0xd6, 0x51, 0xb2, 0x5c, 0xa5, 0x5f, 0x48, 0xad, 0x29,
	0x30, 0x91, 0xfb, 0xe6, 0xcd, 0x4f, 0x3e, 0x16, 0xf5, 0xf4, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xab,
	0x18, 0x54, 0xdd, 0x76, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVMClient is the client API for KVM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVMClient interface {
	RegisterAgentIdentity(ctx context.Context, in *AgentIdentity, opts ...grpc.CallOption) (*Status, error)
	SendPolicy(ctx context.Context, opts ...grpc.CallOption) (KVM_SendPolicyClient, error)
}

type kVMClient struct {
	cc *grpc.ClientConn
}

func NewKVMClient(cc *grpc.ClientConn) KVMClient {
	return &kVMClient{cc}
}

func (c *kVMClient) RegisterAgentIdentity(ctx context.Context, in *AgentIdentity, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/kvm.KVM/registerAgentIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMClient) SendPolicy(ctx context.Context, opts ...grpc.CallOption) (KVM_SendPolicyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KVM_serviceDesc.Streams[0], "/kvm.KVM/sendPolicy", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVMSendPolicyClient{stream}
	return x, nil
}

type KVM_SendPolicyClient interface {
	Send(*PolicyData) error
	Recv() (*Status, error)
	grpc.ClientStream
}

type kVMSendPolicyClient struct {
	grpc.ClientStream
}

func (x *kVMSendPolicyClient) Send(m *PolicyData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kVMSendPolicyClient) Recv() (*Status, error) {
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KVMServer is the server API for KVM service.
type KVMServer interface {
	RegisterAgentIdentity(context.Context, *AgentIdentity) (*Status, error)
	SendPolicy(KVM_SendPolicyServer) error
}

// UnimplementedKVMServer can be embedded to have forward compatible implementations.
type UnimplementedKVMServer struct {
}

func (*UnimplementedKVMServer) RegisterAgentIdentity(ctx context.Context, req *AgentIdentity) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAgentIdentity not implemented")
}
func (*UnimplementedKVMServer) SendPolicy(srv KVM_SendPolicyServer) error {
	return status.Errorf(codes.Unimplemented, "method SendPolicy not implemented")
}

func RegisterKVMServer(s *grpc.Server, srv KVMServer) {
	s.RegisterService(&_KVM_serviceDesc, srv)
}

func _KVM_RegisterAgentIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServer).RegisterAgentIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvm.KVM/RegisterAgentIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServer).RegisterAgentIdentity(ctx, req.(*AgentIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVM_SendPolicy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KVMServer).SendPolicy(&kVMSendPolicyServer{stream})
}

type KVM_SendPolicyServer interface {
	Send(*Status) error
	Recv() (*PolicyData, error)
	grpc.ServerStream
}

type kVMSendPolicyServer struct {
	grpc.ServerStream
}

func (x *kVMSendPolicyServer) Send(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kVMSendPolicyServer) Recv() (*PolicyData, error) {
	m := new(PolicyData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _KVM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvm.KVM",
	HandlerType: (*KVMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerAgentIdentity",
			Handler:    _KVM_RegisterAgentIdentity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sendPolicy",
			Handler:       _KVM_SendPolicy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "kvm.proto",
}
