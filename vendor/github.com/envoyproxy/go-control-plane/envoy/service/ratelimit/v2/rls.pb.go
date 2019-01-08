// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/ratelimit/v2/rls.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import ratelimit "github.com/envoyproxy/go-control-plane/envoy/api/v2/ratelimit"
import _ "github.com/lyft/protoc-gen-validate/validate"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RateLimitResponse_Code int32

const (
	// The response code is not known.
	RateLimitResponse_UNKNOWN RateLimitResponse_Code = 0
	// The response code to notify that the number of requests are under limit.
	RateLimitResponse_OK RateLimitResponse_Code = 1
	// The response code to notify that the number of requests are over limit.
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

var RateLimitResponse_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}
var RateLimitResponse_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x RateLimitResponse_Code) String() string {
	return proto.EnumName(RateLimitResponse_Code_name, int32(x))
}
func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rls_fcfbfd85039b3384, []int{1, 0}
}

type RateLimitResponse_RateLimit_Unit int32

const (
	// The time unit is not known.
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	// The time unit representing a second.
	RateLimitResponse_RateLimit_SECOND RateLimitResponse_RateLimit_Unit = 1
	// The time unit representing a minute.
	RateLimitResponse_RateLimit_MINUTE RateLimitResponse_RateLimit_Unit = 2
	// The time unit representing an hour.
	RateLimitResponse_RateLimit_HOUR RateLimitResponse_RateLimit_Unit = 3
	// The time unit representing a day.
	RateLimitResponse_RateLimit_DAY RateLimitResponse_RateLimit_Unit = 4
)

var RateLimitResponse_RateLimit_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}
var RateLimitResponse_RateLimit_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return proto.EnumName(RateLimitResponse_RateLimit_Unit_name, int32(x))
}
func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rls_fcfbfd85039b3384, []int{1, 0, 0}
}

// Main message for a rate limit request. The rate limit service is designed to be fully generic
// in the sense that it can operate on arbitrary hierarchical key/value pairs. The loaded
// configuration will parse the request and find the most specific limit to apply. In addition,
// a RateLimitRequest can contain multiple "descriptors" to limit on. When multiple descriptors
// are provided, the server will limit on *ALL* of them and return an OVER_LIMIT response if any
// of them are over limit. This enables more complex application level rate limiting scenarios
// if desired.
type RateLimitRequest struct {
	// All rate limit requests must specify a domain. This enables the configuration to be per
	// application without fear of overlap. E.g., "envoy".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// All rate limit requests must specify at least one RateLimitDescriptor. Each descriptor is
	// processed by the service (see below). If any of the descriptors are over limit, the entire
	// request is considered to be over limit.
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// Rate limit requests can optionally specify the number of hits a request adds to the matched
	// limit. If the value is not set in the message, a request increases the matched limit by 1.
	HitsAddend           uint32   `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitRequest) Reset()         { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()    {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rls_fcfbfd85039b3384, []int{0}
}
func (m *RateLimitRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRequest.Merge(dst, src)
}
func (m *RateLimitRequest) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRequest proto.InternalMessageInfo

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitRequest) GetHitsAddend() uint32 {
	if m != nil {
		return m.HitsAddend
	}
	return 0
}

// A response from a ShouldRateLimit call.
type RateLimitResponse struct {
	// The overall response code which takes into account all of the descriptors that were passed
	// in the RateLimitRequest message.
	OverallCode RateLimitResponse_Code `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"overall_code,omitempty"`
	// A list of DescriptorStatus messages which matches the length of the descriptor list passed
	// in the RateLimitRequest. This can be used by the caller to determine which individual
	// descriptors failed and/or what the currently configured limits are for all of them.
	Statuses []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	// A list of headers to add to the response
	Headers              []*core.HeaderValue `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RateLimitResponse) Reset()         { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()    {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rls_fcfbfd85039b3384, []int{1}
}
func (m *RateLimitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse.Merge(dst, src)
}
func (m *RateLimitResponse) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse proto.InternalMessageInfo

func (m *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if m != nil {
		return m.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func (m *RateLimitResponse) GetHeaders() []*core.HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Defines an actual rate limit in terms of requests per unit of time and the unit itself.
type RateLimitResponse_RateLimit struct {
	// The number of requests per unit of time.
	RequestsPerUnit uint32 `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	// The unit of time.
	Unit                 RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitResponse_RateLimit) Reset()         { *m = RateLimitResponse_RateLimit{} }
func (m *RateLimitResponse_RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_RateLimit) ProtoMessage()    {}
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_rls_fcfbfd85039b3384, []int{1, 0}
}
func (m *RateLimitResponse_RateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitResponse_RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitResponse_RateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RateLimitResponse_RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_RateLimit.Merge(dst, src)
}
func (m *RateLimitResponse_RateLimit) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitResponse_RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_RateLimit proto.InternalMessageInfo

func (m *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

func (m *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if m != nil {
		return m.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	// The response code for an individual descriptor.
	Code RateLimitResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"code,omitempty"`
	// The current limit as configured by the server. Useful for debugging, etc.
	CurrentLimit *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	// The limit remaining in the current time unit.
	LimitRemaining       uint32   `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitResponse_DescriptorStatus) Reset()         { *m = RateLimitResponse_DescriptorStatus{} }
func (m *RateLimitResponse_DescriptorStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_DescriptorStatus) ProtoMessage()    {}
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_rls_fcfbfd85039b3384, []int{1, 1}
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RateLimitResponse_DescriptorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.Merge(dst, src)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitResponse_DescriptorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_DescriptorStatus proto.InternalMessageInfo

func (m *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if m != nil {
		return m.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func init() {
	proto.RegisterType((*RateLimitRequest)(nil), "envoy.service.ratelimit.v2.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "envoy.service.ratelimit.v2.RateLimitResponse")
	proto.RegisterType((*RateLimitResponse_RateLimit)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.RateLimit")
	proto.RegisterType((*RateLimitResponse_DescriptorStatus)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus")
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_Code", RateLimitResponse_Code_name, RateLimitResponse_Code_value)
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit", RateLimitResponse_RateLimit_Unit_name, RateLimitResponse_RateLimit_Unit_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v2.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v2/rls.proto",
}

func (m *RateLimitRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Domain) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRls(dAtA, i, uint64(len(m.Domain)))
		i += copy(dAtA[i:], m.Domain)
	}
	if len(m.Descriptors) > 0 {
		for _, msg := range m.Descriptors {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRls(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.HitsAddend != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.HitsAddend))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RateLimitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OverallCode != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.OverallCode))
	}
	if len(m.Statuses) > 0 {
		for _, msg := range m.Statuses {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRls(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Headers) > 0 {
		for _, msg := range m.Headers {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRls(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RateLimitResponse_RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitResponse_RateLimit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RequestsPerUnit != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.RequestsPerUnit))
	}
	if m.Unit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.Unit))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RateLimitResponse_DescriptorStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitResponse_DescriptorStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.Code))
	}
	if m.CurrentLimit != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.CurrentLimit.Size()))
		n1, err := m.CurrentLimit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.LimitRemaining != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.LimitRemaining))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRls(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RateLimitRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovRls(uint64(l))
	}
	if len(m.Descriptors) > 0 {
		for _, e := range m.Descriptors {
			l = e.Size()
			n += 1 + l + sovRls(uint64(l))
		}
	}
	if m.HitsAddend != 0 {
		n += 1 + sovRls(uint64(m.HitsAddend))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RateLimitResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OverallCode != 0 {
		n += 1 + sovRls(uint64(m.OverallCode))
	}
	if len(m.Statuses) > 0 {
		for _, e := range m.Statuses {
			l = e.Size()
			n += 1 + l + sovRls(uint64(l))
		}
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovRls(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RateLimitResponse_RateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestsPerUnit != 0 {
		n += 1 + sovRls(uint64(m.RequestsPerUnit))
	}
	if m.Unit != 0 {
		n += 1 + sovRls(uint64(m.Unit))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RateLimitResponse_DescriptorStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovRls(uint64(m.Code))
	}
	if m.CurrentLimit != nil {
		l = m.CurrentLimit.Size()
		n += 1 + l + sovRls(uint64(l))
	}
	if m.LimitRemaining != 0 {
		n += 1 + sovRls(uint64(m.LimitRemaining))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRls(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRls(x uint64) (n int) {
	return sovRls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimitRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Descriptors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Descriptors = append(m.Descriptors, &ratelimit.RateLimitDescriptor{})
			if err := m.Descriptors[len(m.Descriptors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HitsAddend", wireType)
			}
			m.HitsAddend = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HitsAddend |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimitResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OverallCode", wireType)
			}
			m.OverallCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OverallCode |= (RateLimitResponse_Code(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statuses = append(m.Statuses, &RateLimitResponse_DescriptorStatus{})
			if err := m.Statuses[len(m.Statuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &core.HeaderValue{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimitResponse_RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestsPerUnit", wireType)
			}
			m.RequestsPerUnit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestsPerUnit |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			m.Unit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Unit |= (RateLimitResponse_RateLimit_Unit(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimitResponse_DescriptorStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DescriptorStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DescriptorStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (RateLimitResponse_Code(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CurrentLimit == nil {
				m.CurrentLimit = &RateLimitResponse_RateLimit{}
			}
			if err := m.CurrentLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitRemaining", wireType)
			}
			m.LimitRemaining = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LimitRemaining |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRls
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRls
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRls
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRls
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRls
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRls(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRls = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRls   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/service/ratelimit/v2/rls.proto", fileDescriptor_rls_fcfbfd85039b3384)
}

var fileDescriptor_rls_fcfbfd85039b3384 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x9b, 0xdd, 0xb8, 0x6d, 0x4f, 0xfa, 0x27, 0x9d, 0x0b, 0x5d, 0x82, 0xac, 0x65, 0x11,
	0x2d, 0x56, 0x13, 0x88, 0x17, 0x7a, 0x21, 0x85, 0xda, 0x56, 0x5a, 0xda, 0xee, 0x96, 0xd9, 0x6e,
	0xc5, 0x22, 0x84, 0xe9, 0xe6, 0x60, 0x07, 0xd2, 0x24, 0xce, 0x4c, 0x02, 0xde, 0xfb, 0x14, 0x3e,
	0x91, 0xde, 0xf9, 0x02, 0x82, 0xf6, 0x49, 0x24, 0x93, 0x6c, 0x76, 0xab, 0x08, 0xae, 0xde, 0x65,
	0xce, 0x39, 0xdf, 0x6f, 0xcf, 0x77, 0xce, 0xcc, 0xc2, 0x7d, 0x8c, 0xf3, 0xe4, 0x83, 0x27, 0x51,
	0xe4, 0x7c, 0x84, 0x9e, 0x60, 0x0a, 0x23, 0x7e, 0xc5, 0x95, 0x97, 0xfb, 0x9e, 0x88, 0xa4, 0x9b,
	0x8a, 0x44, 0x25, 0xc4, 0xd1, 0x55, 0x6e, 0x55, 0xe5, 0xd6, 0x55, 0x6e, 0xee, 0x3b, 0x77, 0x4b,
	0x02, 0x4b, 0x79, 0xa1, 0x19, 0x25, 0x02, 0xbd, 0x0b, 0x26, 0xb1, 0x54, 0x3a, 0x0f, 0x6e, 0x64,
	0x27, 0xf8, 0x09, 0xa2, 0xac, 0xbb, 0x93, 0xb3, 0x88, 0x87, 0x4c, 0xa1, 0x37, 0xfe, 0x28, 0x13,
	0xdd, 0x4f, 0x06, 0xd8, 0x94, 0x29, 0x3c, 0x2a, 0x8a, 0x29, 0xbe, 0xcf, 0x50, 0x2a, 0x72, 0x1b,
	0x5a, 0x61, 0x72, 0xc5, 0x78, 0xdc, 0x36, 0xd6, 0x8d, 0x8d, 0x45, 0x5a, 0x9d, 0xc8, 0x31, 0x58,
	0x21, 0xca, 0x91, 0xe0, 0xa9, 0x4a, 0x84, 0x6c, 0x37, 0xd6, 0x9b, 0x1b, 0x96, 0xbf, 0xe9, 0x96,
	0xdd, 0xb3, 0x94, 0xbb, 0xb9, 0x3f, 0xd5, 0x7c, 0x8d, 0xdd, 0xad, 0x35, 0x74, 0x5a, 0x4f, 0xee,
	0x81, 0x75, 0xc9, 0x95, 0x0c, 0x58, 0x18, 0x62, 0x1c, 0xb6, 0x9b, 0xeb, 0xc6, 0xc6, 0x32, 0x85,
	0x22, 0xb4, 0xad, 0x23, 0xdd, 0x6f, 0xb7, 0x60, 0x6d, 0xaa, 0x39, 0x99, 0x26, 0xb1, 0x44, 0x32,
	0x84, 0xa5, 0x24, 0x47, 0xc1, 0xa2, 0x28, 0x18, 0x25, 0x21, 0xea, 0x1e, 0x57, 0x7c, 0xdf, 0xfd,
	0xf3, 0x10, 0xdd, 0xdf, 0x20, 0xee, 0x4e, 0x12, 0x22, 0xb5, 0x2a, 0x4e, 0x71, 0x20, 0xe7, 0xb0,
	0x20, 0x15, 0x53, 0x99, 0xc4, 0xb1, 0xb3, 0xad, 0xd9, 0x90, 0x13, 0x9b, 0x03, 0xcd, 0xa1, 0x35,
	0x8f, 0x3c, 0x87, 0xf9, 0x4b, 0x64, 0x21, 0x0a, 0xd9, 0x6e, 0x6a, 0x74, 0xe7, 0xe6, 0xd0, 0x8a,
	0xb5, 0xba, 0xfb, 0xba, 0xe2, 0x8c, 0x45, 0x19, 0xd2, 0x71, 0xb9, 0xf3, 0xc5, 0x80, 0xc5, 0xfa,
	0xa7, 0xc8, 0x23, 0x58, 0x13, 0xe5, 0x8e, 0x64, 0x90, 0xa2, 0x08, 0xb2, 0x98, 0x2b, 0xed, 0x7f,
	0x99, 0xae, 0x8e, 0x13, 0x27, 0x28, 0x86, 0x31, 0x57, 0xe4, 0x04, 0x4c, 0x9d, 0x6e, 0xe8, 0xf1,
	0xbc, 0x98, 0xcd, 0x4b, 0x1d, 0x71, 0x0b, 0x16, 0xd5, 0xa4, 0xee, 0x16, 0x98, 0x9a, 0x6c, 0xc1,
	0xfc, 0xb0, 0x77, 0xd8, 0xeb, 0xbf, 0xee, 0xd9, 0x73, 0x04, 0xa0, 0x35, 0xd8, 0xdb, 0xe9, 0xf7,
	0x76, 0x6d, 0xa3, 0xf8, 0x3e, 0x3e, 0xe8, 0x0d, 0x4f, 0xf7, 0xec, 0x06, 0x59, 0x00, 0x73, 0xbf,
	0x3f, 0xa4, 0x76, 0x93, 0xcc, 0x43, 0x73, 0x77, 0xfb, 0x8d, 0x6d, 0x3a, 0x3f, 0x0c, 0xb0, 0x7f,
	0x1d, 0x12, 0x79, 0x05, 0xe6, 0x7f, 0x6e, 0x51, 0xeb, 0xc9, 0x5b, 0x58, 0x1e, 0x65, 0x42, 0x60,
	0xac, 0x02, 0x2d, 0xd0, 0xbe, 0x2d, 0xff, 0xd9, 0x3f, 0xfa, 0xa6, 0x4b, 0x15, 0xad, 0x1c, 0xfc,
	0x43, 0x58, 0xd5, 0xaa, 0x40, 0x60, 0xf1, 0x12, 0x78, 0xfc, 0xae, 0xba, 0xae, 0x2b, 0x51, 0xa9,
	0xaf, 0xa2, 0xdd, 0x4d, 0x30, 0xf5, 0x6d, 0xba, 0x31, 0xa3, 0x16, 0x34, 0xfa, 0x87, 0xb6, 0x41,
	0x56, 0x00, 0xfa, 0x67, 0x7b, 0x34, 0x38, 0x3a, 0x38, 0x3e, 0x38, 0xb5, 0x1b, 0xfe, 0xc7, 0xe9,
	0xc7, 0x37, 0x28, 0x3b, 0x24, 0x29, 0xac, 0x0e, 0x2e, 0x93, 0x2c, 0x0a, 0x27, 0x6b, 0x7f, 0xfc,
	0x97, 0x26, 0xf4, 0x05, 0x70, 0x9e, 0xcc, 0x64, 0xb9, 0x3b, 0xf7, 0xd2, 0xfe, 0x7c, 0xdd, 0x31,
	0xbe, 0x5e, 0x77, 0x8c, 0xef, 0xd7, 0x1d, 0xe3, 0xbc, 0x91, 0xfb, 0x17, 0x2d, 0xfd, 0xe7, 0xf0,
	0xf4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x20, 0x5c, 0x25, 0xbf, 0x04, 0x00, 0x00,
}
