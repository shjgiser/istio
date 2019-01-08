// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/rate_limit/v2/rate_limit.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import ratelimit "github.com/envoyproxy/go-control-plane/envoy/api/v2/ratelimit"
import v2 "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v2"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RateLimit struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// The rate limit descriptor list to use in the rate limit service request.
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *time.Duration `protobuf:"bytes,4,opt,name=timeout,proto3,stdduration" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	// [#comment:TODO(ramaraochavali): Make this required as part of cleanup of deprecated ratelimit
	// service config in bootstrap.]
	RateLimitService     *v2.RateLimitServiceConfig `protobuf:"bytes,6,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_rate_limit_2ae29e2b9f4c0f4c, []int{0}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(dst, src)
}
func (m *RateLimit) XXX_Size() int {
	return m.Size()
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimit) GetTimeout() *time.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.network.rate_limit.v2.RateLimit")
}
func (m *RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StatPrefix) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.StatPrefix)))
		i += copy(dAtA[i:], m.StatPrefix)
	}
	if len(m.Domain) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.Domain)))
		i += copy(dAtA[i:], m.Domain)
	}
	if len(m.Descriptors) > 0 {
		for _, msg := range m.Descriptors {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRateLimit(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Timeout != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.Timeout)))
		n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.Timeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.FailureModeDeny {
		dAtA[i] = 0x28
		i++
		if m.FailureModeDeny {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RateLimitService != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(m.RateLimitService.Size()))
		n2, err := m.RateLimitService.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRateLimit(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if len(m.Descriptors) > 0 {
		for _, e := range m.Descriptors {
			l = e.Size()
			n += 1 + l + sovRateLimit(uint64(l))
		}
	}
	if m.Timeout != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.Timeout)
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.FailureModeDeny {
		n += 2
	}
	if m.RateLimitService != nil {
		l = m.RateLimitService.Size()
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRateLimit(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRateLimit(x uint64) (n int) {
	return sovRateLimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRateLimit
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Descriptors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.Timeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureModeDeny", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailureModeDeny = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RateLimitService == nil {
				m.RateLimitService = &v2.RateLimitServiceConfig{}
			}
			if err := m.RateLimitService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRateLimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRateLimit
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
func skipRateLimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
				return 0, ErrInvalidLengthRateLimit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRateLimit
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
				next, err := skipRateLimit(dAtA[start:])
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
	ErrInvalidLengthRateLimit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRateLimit   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/network/rate_limit/v2/rate_limit.proto", fileDescriptor_rate_limit_2ae29e2b9f4c0f4c)
}

var fileDescriptor_rate_limit_2ae29e2b9f4c0f4c = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0xcf, 0x8e, 0x94, 0x30,
	0x18, 0x4f, 0x99, 0xd9, 0x71, 0xa7, 0x1c, 0x5c, 0x1b, 0x13, 0x71, 0x0f, 0x88, 0x9a, 0x18, 0x5c,
	0x93, 0x36, 0xe2, 0x49, 0x8f, 0x38, 0x47, 0x4d, 0x0c, 0x9e, 0xf4, 0x42, 0xba, 0xc3, 0x07, 0x69,
	0x04, 0x4a, 0x4a, 0xa9, 0xce, 0x6b, 0x78, 0xf2, 0x59, 0x3c, 0x79, 0xf4, 0xe8, 0x1b, 0x68, 0xe6,
	0xe6, 0x43, 0x98, 0x18, 0x28, 0x0c, 0xb3, 0x73, 0xfb, 0xb5, 0xbf, 0x3f, 0x5f, 0xbf, 0x1f, 0xe0,
	0x57, 0x50, 0x1b, 0xb9, 0x63, 0x5b, 0x59, 0xe7, 0xa2, 0x60, 0xb9, 0x28, 0x35, 0x28, 0x56, 0x83,
	0xfe, 0x2c, 0xd5, 0x27, 0xa6, 0xb8, 0x86, 0xb4, 0x14, 0x95, 0xd0, 0xcc, 0x44, 0x47, 0x27, 0xda,
	0x28, 0xa9, 0x25, 0x79, 0x3a, 0x78, 0xa9, 0xf5, 0x52, 0xeb, 0xa5, 0xa3, 0x97, 0x1e, 0xa9, 0x4d,
	0x74, 0xf9, 0xc4, 0x8e, 0xe1, 0x8d, 0x98, 0x92, 0x6c, 0xec, 0x01, 0xd9, 0xc8, 0xcb, 0xc7, 0x37,
	0x9e, 0x33, 0xeb, 0x7a, 0x53, 0xd9, 0x8e, 0x22, 0xbf, 0x90, 0xb2, 0x28, 0x81, 0x0d, 0xa7, 0xeb,
	0x2e, 0x67, 0x59, 0xa7, 0xb8, 0x16, 0xb2, 0x1e, 0xf9, 0x7b, 0x86, 0x97, 0x22, 0xe3, 0x1a, 0xd8,
	0x04, 0x46, 0xe2, 0x6e, 0x21, 0x0b, 0x39, 0x40, 0xd6, 0x23, 0x7b, 0xfb, 0xe8, 0x9f, 0x83, 0xd7,
	0x09, 0xd7, 0xf0, 0xa6, 0x9f, 0x44, 0xae, 0xb0, 0xdb, 0x6a, 0xae, 0xd3, 0x46, 0x41, 0x2e, 0xbe,
	0x78, 0x28, 0x40, 0xe1, 0x3a, 0x5e, 0x7f, 0xff, 0xfb, 0x63, 0xb1, 0x54, 0x4e, 0x80, 0x12, 0xdc,
	0xb3, 0xef, 0x06, 0x92, 0x3c, 0xc4, 0xab, 0x4c, 0x56, 0x5c, 0xd4, 0x9e, 0x73, 0x2a, 0x1b, 0x09,
	0xf2, 0x01, 0xbb, 0x19, 0xb4, 0x5b, 0x25, 0x1a, 0x2d, 0x55, 0xeb, 0x2d, 0x82, 0x45, 0xe8, 0x46,
	0xcf, 0xa8, 0x6d, 0x8e, 0x37, 0x82, 0x9a, 0x88, 0xce, 0x25, 0x1c, 0x9e, 0xb1, 0x39, 0x78, 0x62,
	0xdc, 0x87, 0x9e, 0x7d, 0x45, 0xce, 0x39, 0x4a, 0x8e, 0xb3, 0xc8, 0x4b, 0x7c, 0x4b, 0x8b, 0x0a,
	0x64, 0xa7, 0xbd, 0x65, 0x80, 0x42, 0x37, 0xba, 0x4f, 0x6d, 0x31, 0x74, 0x2a, 0x86, 0x6e, 0xc6,
	0x62, 0xe2, 0xe5, 0xb7, 0xdf, 0x0f, 0x50, 0x32, 0xe9, 0xc9, 0x15, 0xbe, 0x93, 0x73, 0x51, 0x76,
	0x0a, 0xd2, 0x4a, 0x66, 0x90, 0x66, 0x50, 0xef, 0xbc, 0xb3, 0x00, 0x85, 0xe7, 0xc9, 0xed, 0x91,
	0x78, 0x2b, 0x33, 0xd8, 0x40, 0xbd, 0x23, 0x29, 0x26, 0xf3, 0xb7, 0x4c, 0x5b, 0x50, 0x46, 0x6c,
	0xc1, 0x5b, 0x0d, 0x13, 0x9f, 0xd3, 0x1b, 0xbf, 0xc0, 0xbc, 0x88, 0x89, 0xe6, 0x5d, 0xde, 0x5b,
	0xcb, 0xeb, 0x41, 0x93, 0x5c, 0xa8, 0x93, 0xfb, 0xf8, 0xe2, 0xe7, 0xde, 0x47, 0xbf, 0xf6, 0x3e,
	0xfa, 0xb3, 0xf7, 0xd1, 0x47, 0xc7, 0x44, 0xd7, 0xab, 0x61, 0x81, 0x17, 0xff, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x53, 0x89, 0x60, 0xa9, 0x9d, 0x02, 0x00, 0x00,
}
