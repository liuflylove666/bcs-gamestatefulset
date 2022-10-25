// Code generated by protoc-gen-go.
// source: integrator.proto
// DO NOT EDIT!

/*
Package integrator is a generated protocol buffer package.

It is generated from these files:
	integrator.proto

It has these top-level messages:
	IntegrateReq
	IntegrateResp
*/
package integrator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bk-bscp/internal/protocol/common"

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

type IntegrateReq struct {
	Seq      uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Metadata string `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	Operator string `protobuf:"bytes,3,opt,name=operator" json:"operator,omitempty"`
}

func (m *IntegrateReq) Reset()                    { *m = IntegrateReq{} }
func (m *IntegrateReq) String() string            { return proto.CompactTextString(m) }
func (*IntegrateReq) ProtoMessage()               {}
func (*IntegrateReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IntegrateReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *IntegrateReq) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *IntegrateReq) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type IntegrateResp struct {
	Seq        uint64         `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ErrCode    common.ErrCode `protobuf:"varint,2,opt,name=errCode,enum=common.ErrCode" json:"errCode,omitempty"`
	ErrMsg     string         `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
	Bid        string         `protobuf:"bytes,4,opt,name=bid" json:"bid,omitempty"`
	Appid      string         `protobuf:"bytes,5,opt,name=appid" json:"appid,omitempty"`
	Clusterid  string         `protobuf:"bytes,6,opt,name=clusterid" json:"clusterid,omitempty"`
	Zoneid     string         `protobuf:"bytes,7,opt,name=zoneid" json:"zoneid,omitempty"`
	Cfgsetid   string         `protobuf:"bytes,8,opt,name=cfgsetid" json:"cfgsetid,omitempty"`
	Commitid   string         `protobuf:"bytes,9,opt,name=commitid" json:"commitid,omitempty"`
	Strategyid string         `protobuf:"bytes,10,opt,name=strategyid" json:"strategyid,omitempty"`
	Releaseid  string         `protobuf:"bytes,11,opt,name=releaseid" json:"releaseid,omitempty"`
	Locker     string         `protobuf:"bytes,12,opt,name=locker" json:"locker,omitempty"`
	LockTime   string         `protobuf:"bytes,13,opt,name=lockTime" json:"lockTime,omitempty"`
}

func (m *IntegrateResp) Reset()                    { *m = IntegrateResp{} }
func (m *IntegrateResp) String() string            { return proto.CompactTextString(m) }
func (*IntegrateResp) ProtoMessage()               {}
func (*IntegrateResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IntegrateResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *IntegrateResp) GetErrCode() common.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return common.ErrCode_E_OK
}

func (m *IntegrateResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *IntegrateResp) GetBid() string {
	if m != nil {
		return m.Bid
	}
	return ""
}

func (m *IntegrateResp) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *IntegrateResp) GetClusterid() string {
	if m != nil {
		return m.Clusterid
	}
	return ""
}

func (m *IntegrateResp) GetZoneid() string {
	if m != nil {
		return m.Zoneid
	}
	return ""
}

func (m *IntegrateResp) GetCfgsetid() string {
	if m != nil {
		return m.Cfgsetid
	}
	return ""
}

func (m *IntegrateResp) GetCommitid() string {
	if m != nil {
		return m.Commitid
	}
	return ""
}

func (m *IntegrateResp) GetStrategyid() string {
	if m != nil {
		return m.Strategyid
	}
	return ""
}

func (m *IntegrateResp) GetReleaseid() string {
	if m != nil {
		return m.Releaseid
	}
	return ""
}

func (m *IntegrateResp) GetLocker() string {
	if m != nil {
		return m.Locker
	}
	return ""
}

func (m *IntegrateResp) GetLockTime() string {
	if m != nil {
		return m.LockTime
	}
	return ""
}

func init() {
	proto.RegisterType((*IntegrateReq)(nil), "integrator.IntegrateReq")
	proto.RegisterType((*IntegrateResp)(nil), "integrator.IntegrateResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Integrator service

type IntegratorClient interface {
	Integrate(ctx context.Context, in *IntegrateReq, opts ...grpc.CallOption) (*IntegrateResp, error)
}

type integratorClient struct {
	cc *grpc.ClientConn
}

func NewIntegratorClient(cc *grpc.ClientConn) IntegratorClient {
	return &integratorClient{cc}
}

func (c *integratorClient) Integrate(ctx context.Context, in *IntegrateReq, opts ...grpc.CallOption) (*IntegrateResp, error) {
	out := new(IntegrateResp)
	err := grpc.Invoke(ctx, "/integrator.Integrator/Integrate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Integrator service

type IntegratorServer interface {
	Integrate(context.Context, *IntegrateReq) (*IntegrateResp, error)
}

func RegisterIntegratorServer(s *grpc.Server, srv IntegratorServer) {
	s.RegisterService(&_Integrator_serviceDesc, srv)
}

func _Integrator_Integrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegrateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegratorServer).Integrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integrator.Integrator/Integrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegratorServer).Integrate(ctx, req.(*IntegrateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Integrator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "integrator.Integrator",
	HandlerType: (*IntegratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Integrate",
			Handler:    _Integrator_Integrate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrator.proto",
}

func init() { proto.RegisterFile("integrator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0x3d, 0x4f, 0xc3, 0x30,
	0x14, 0xa4, 0xdf, 0xcd, 0xa3, 0x85, 0xca, 0x42, 0xc8, 0x54, 0x08, 0x55, 0x9d, 0xca, 0xd0, 0x56,
	0x2a, 0xff, 0x00, 0xc4, 0xd0, 0x01, 0x09, 0x45, 0x0c, 0xac, 0x49, 0xfc, 0x88, 0xac, 0x26, 0x71,
	0x6a, 0x9b, 0x01, 0x7e, 0x33, 0x3f, 0x02, 0x3d, 0x3b, 0x49, 0x33, 0xc0, 0x14, 0xdf, 0x5d, 0x7c,
	0x77, 0x2f, 0x79, 0x30, 0x93, 0x85, 0xc5, 0x54, 0x47, 0x56, 0xe9, 0x4d, 0xa9, 0x95, 0x55, 0x0c,
	0x4e, 0xcc, 0x7c, 0x1d, 0x1f, 0xd6, 0xb1, 0x49, 0xca, 0x2d, 0x71, 0xba, 0x88, 0xb2, 0xad, 0x7b,
	0x27, 0x51, 0xd9, 0x36, 0x51, 0x79, 0xae, 0x8a, 0xea, 0xe1, 0xaf, 0x2e, 0xdf, 0x61, 0xb2, 0xaf,
	0x2e, 0x63, 0x88, 0x47, 0x36, 0x83, 0x9e, 0xc1, 0x23, 0xef, 0x2c, 0x3a, 0xab, 0x7e, 0x48, 0x47,
	0x36, 0x87, 0x71, 0x8e, 0x36, 0x12, 0x91, 0x8d, 0x78, 0x77, 0xd1, 0x59, 0x05, 0x61, 0x83, 0x49,
	0x53, 0x25, 0xba, 0x60, 0xde, 0xf3, 0x5a, 0x8d, 0x97, 0x3f, 0x5d, 0x98, 0xb6, 0xac, 0x4d, 0xf9,
	0x87, 0xf7, 0x3d, 0x8c, 0x50, 0xeb, 0x27, 0x25, 0xd0, 0x59, 0x5f, 0xec, 0x2e, 0x37, 0x55, 0xbb,
	0x67, 0x4f, 0x87, 0xb5, 0xce, 0xae, 0x61, 0x88, 0x5a, 0xbf, 0x98, 0xb4, 0x0a, 0xaa, 0x10, 0x99,
	0xc6, 0x52, 0xf0, 0xbe, 0x23, 0xe9, 0xc8, 0xae, 0x60, 0x10, 0x95, 0xa5, 0x14, 0x7c, 0xe0, 0x38,
	0x0f, 0xd8, 0x2d, 0x04, 0x49, 0xf6, 0x69, 0x2c, 0x6a, 0x29, 0xf8, 0xd0, 0x29, 0x27, 0x82, 0xdc,
	0xbf, 0x55, 0x81, 0x52, 0xf0, 0x91, 0x77, 0xf7, 0x88, 0x06, 0x4c, 0x3e, 0x52, 0x83, 0x56, 0x0a,
	0x3e, 0xf6, 0x03, 0xd6, 0xd8, 0x69, 0x2a, 0xcf, 0x25, 0x69, 0x41, 0xa5, 0x55, 0x98, 0xdd, 0x01,
	0x18, 0x4b, 0x83, 0xa7, 0x5f, 0x52, 0x70, 0x70, 0x6a, 0x8b, 0xa1, 0x36, 0x1a, 0x33, 0x8c, 0x0c,
	0x45, 0x9e, 0xfb, 0x36, 0x0d, 0x41, 0x6d, 0x32, 0x95, 0x1c, 0x50, 0xf3, 0x89, 0x6f, 0xe3, 0x11,
	0x25, 0xd2, 0xe9, 0x4d, 0xe6, 0xc8, 0xa7, 0x3e, 0xb1, 0xc6, 0xbb, 0x57, 0x80, 0x7d, 0xb3, 0x05,
	0xec, 0x11, 0x82, 0xe6, 0xdb, 0x33, 0xbe, 0x69, 0x6d, 0x4c, 0xfb, 0x6f, 0xcf, 0x6f, 0xfe, 0x51,
	0x4c, 0xb9, 0x3c, 0x8b, 0x87, 0x6e, 0x43, 0x1e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x99, 0x8b,
	0x40, 0x26, 0x70, 0x02, 0x00, 0x00,
}