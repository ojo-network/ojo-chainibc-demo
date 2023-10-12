// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ojochainibcdemo/oracleintegration/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgRequestPrice struct {
	Symbols []string         `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Type    PriceRequestType `protobuf:"varint,2,opt,name=type,proto3,enum=ojochainibcdemo.oracleintegration.PriceRequestType" json:"type,omitempty"`
	From    string           `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
}

func (m *MsgRequestPrice) Reset()         { *m = MsgRequestPrice{} }
func (m *MsgRequestPrice) String() string { return proto.CompactTextString(m) }
func (*MsgRequestPrice) ProtoMessage()    {}
func (*MsgRequestPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_262163a9edac2ece, []int{0}
}
func (m *MsgRequestPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestPrice.Merge(m, src)
}
func (m *MsgRequestPrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestPrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestPrice proto.InternalMessageInfo

func (m *MsgRequestPrice) GetSymbols() []string {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *MsgRequestPrice) GetType() PriceRequestType {
	if m != nil {
		return m.Type
	}
	return PRICE_REQUEST_RATE
}

func (m *MsgRequestPrice) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

type MsgRequestPriceResponse struct {
}

func (m *MsgRequestPriceResponse) Reset()         { *m = MsgRequestPriceResponse{} }
func (m *MsgRequestPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestPriceResponse) ProtoMessage()    {}
func (*MsgRequestPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_262163a9edac2ece, []int{1}
}
func (m *MsgRequestPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestPriceResponse.Merge(m, src)
}
func (m *MsgRequestPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestPriceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRequestPrice)(nil), "ojochainibcdemo.oracleintegration.MsgRequestPrice")
	proto.RegisterType((*MsgRequestPriceResponse)(nil), "ojochainibcdemo.oracleintegration.MsgRequestPriceResponse")
}

func init() {
	proto.RegisterFile("ojochainibcdemo/oracleintegration/tx.proto", fileDescriptor_262163a9edac2ece)
}

var fileDescriptor_262163a9edac2ece = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xca, 0xcf, 0xca, 0x4f,
	0xce, 0x48, 0xcc, 0xcc, 0xcb, 0x4c, 0x4a, 0x4e, 0x49, 0xcd, 0xcd, 0xd7, 0xcf, 0x2f, 0x4a, 0x4c,
	0xce, 0x49, 0xcd, 0xcc, 0x2b, 0x49, 0x4d, 0x2f, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0xa9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x44, 0x53, 0xab, 0x87, 0xa1, 0x56, 0x4a, 0x8f,
	0xb0, 0x71, 0x05, 0x89, 0xc9, 0xd9, 0xa9, 0x25, 0x10, 0x23, 0x95, 0x3a, 0x18, 0xb9, 0xf8, 0x7d,
	0x8b, 0xd3, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x02, 0x8a, 0x32, 0x93, 0x53, 0x85, 0x24,
	0xb8, 0xd8, 0x8b, 0x2b, 0x73, 0x93, 0xf2, 0x73, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0x38, 0x83,
	0x60, 0x5c, 0x21, 0x77, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e,
	0x23, 0x63, 0x3d, 0x82, 0xee, 0xd1, 0x03, 0x9b, 0x08, 0x35, 0x3d, 0xa4, 0xb2, 0x20, 0x35, 0x08,
	0x6c, 0x80, 0x90, 0x10, 0x17, 0x4b, 0x5a, 0x51, 0x7e, 0xae, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x98, 0xad, 0x24, 0xc9, 0x25, 0x8e, 0xe6, 0x92, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0xa3, 0x56, 0x46, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x3a, 0x2e, 0x1e, 0x14, 0x97, 0x1a,
	0x11, 0xe1, 0x02, 0x34, 0x33, 0xa5, 0xac, 0x48, 0xd7, 0x03, 0x73, 0x87, 0x53, 0xcc, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x39, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0xe7, 0x67, 0xe5, 0xeb, 0xe6, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0x83,
	0xd9, 0x30, 0xcb, 0x74, 0xc1, 0x11, 0x52, 0x81, 0x2d, 0x86, 0x2b, 0x0b, 0x52, 0x8b, 0x93, 0xd8,
	0xc0, 0x51, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x19, 0xb1, 0x0c, 0x15, 0x13, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RequestPrice(ctx context.Context, in *MsgRequestPrice, opts ...grpc.CallOption) (*MsgRequestPriceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RequestPrice(ctx context.Context, in *MsgRequestPrice, opts ...grpc.CallOption) (*MsgRequestPriceResponse, error) {
	out := new(MsgRequestPriceResponse)
	err := c.cc.Invoke(ctx, "/ojochainibcdemo.oracleintegration.Msg/RequestPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RequestPrice(context.Context, *MsgRequestPrice) (*MsgRequestPriceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RequestPrice(ctx context.Context, req *MsgRequestPrice) (*MsgRequestPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPrice not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RequestPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ojochainibcdemo.oracleintegration.Msg/RequestPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestPrice(ctx, req.(*MsgRequestPrice))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ojochainibcdemo.oracleintegration.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestPrice",
			Handler:    _Msg_RequestPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ojochainibcdemo/oracleintegration/tx.proto",
}

func (m *MsgRequestPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Type != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRequestPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for _, s := range m.Symbols {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Type != 0 {
		n += 1 + sovTx(uint64(m.Type))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRequestPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRequestPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequestPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbols = append(m.Symbols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= PriceRequestType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRequestPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequestPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
