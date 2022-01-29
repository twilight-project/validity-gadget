// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validitygadgetchain/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgSnapshotProposal struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Iteration string `protobuf:"bytes,2,opt,name=iteration,proto3" json:"iteration,omitempty"`
	LogSnap   string `protobuf:"bytes,3,opt,name=logSnap,proto3" json:"logSnap,omitempty"`
}

func (m *MsgSnapshotProposal) Reset()         { *m = MsgSnapshotProposal{} }
func (m *MsgSnapshotProposal) String() string { return proto.CompactTextString(m) }
func (*MsgSnapshotProposal) ProtoMessage()    {}
func (*MsgSnapshotProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_92fcf06d3612c316, []int{0}
}
func (m *MsgSnapshotProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSnapshotProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSnapshotProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSnapshotProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSnapshotProposal.Merge(m, src)
}
func (m *MsgSnapshotProposal) XXX_Size() int {
	return m.Size()
}
func (m *MsgSnapshotProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSnapshotProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSnapshotProposal proto.InternalMessageInfo

func (m *MsgSnapshotProposal) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSnapshotProposal) GetIteration() string {
	if m != nil {
		return m.Iteration
	}
	return ""
}

func (m *MsgSnapshotProposal) GetLogSnap() string {
	if m != nil {
		return m.LogSnap
	}
	return ""
}

type MsgSnapshotProposalResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgSnapshotProposalResponse) Reset()         { *m = MsgSnapshotProposalResponse{} }
func (m *MsgSnapshotProposalResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSnapshotProposalResponse) ProtoMessage()    {}
func (*MsgSnapshotProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92fcf06d3612c316, []int{1}
}
func (m *MsgSnapshotProposalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSnapshotProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSnapshotProposalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSnapshotProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSnapshotProposalResponse.Merge(m, src)
}
func (m *MsgSnapshotProposalResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSnapshotProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSnapshotProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSnapshotProposalResponse proto.InternalMessageInfo

func (m *MsgSnapshotProposalResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgSnapshotProposal)(nil), "twilightproject.validitygadgetchain.validitygadgetchain.MsgSnapshotProposal")
	proto.RegisterType((*MsgSnapshotProposalResponse)(nil), "twilightproject.validitygadgetchain.validitygadgetchain.MsgSnapshotProposalResponse")
}

func init() { proto.RegisterFile("validitygadgetchain/tx.proto", fileDescriptor_92fcf06d3612c316) }

var fileDescriptor_92fcf06d3612c316 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xeb, 0x14, 0x81, 0xea, 0x01, 0xa1, 0xb0, 0x44, 0x50, 0x59, 0xa8, 0x13, 0x4b, 0x12,
	0x09, 0x06, 0x76, 0x66, 0x22, 0xa1, 0xc0, 0xc4, 0xe6, 0x26, 0x96, 0x73, 0x28, 0xe4, 0x2c, 0xfb,
	0x80, 0xf6, 0x2d, 0x78, 0x1a, 0x46, 0x66, 0xc6, 0x8e, 0x8c, 0x28, 0x79, 0x11, 0x44, 0xd2, 0x88,
	0x01, 0x4f, 0x88, 0xf1, 0xce, 0xbe, 0xfb, 0xfe, 0x5f, 0xff, 0xf1, 0xf9, 0x93, 0xac, 0xa1, 0x04,
	0x5a, 0x6b, 0x59, 0x6a, 0x45, 0x45, 0x25, 0xa1, 0x49, 0x69, 0x95, 0x18, 0x8b, 0x84, 0xe1, 0x05,
	0x3d, 0x43, 0x0d, 0xba, 0x22, 0x63, 0xf1, 0x5e, 0x15, 0x94, 0x78, 0x7e, 0xfb, 0x7a, 0x0b, 0xcd,
	0x0f, 0x33, 0xa7, 0x6f, 0x1a, 0x69, 0x5c, 0x85, 0x74, 0x6d, 0xd1, 0xa0, 0x93, 0x75, 0x18, 0xf1,
	0xbd, 0xc2, 0x2a, 0x49, 0x68, 0x23, 0x76, 0xc2, 0x4e, 0x67, 0xf9, 0x58, 0x86, 0x73, 0x3e, 0x03,
	0x52, 0x56, 0x12, 0x60, 0x13, 0x05, 0xfd, 0xdb, 0x4f, 0xe3, 0x7b, 0xae, 0xc6, 0x7e, 0x5d, 0x34,
	0x1d, 0xe6, 0xb6, 0xe5, 0x22, 0xe6, 0xc7, 0x1e, 0x50, 0xae, 0x9c, 0xc1, 0xc6, 0xa9, 0x70, 0x9f,
	0x07, 0x50, 0xf6, 0xac, 0x9d, 0x3c, 0x80, 0xf2, 0xec, 0x8d, 0xf1, 0x69, 0xe6, 0x74, 0xf8, 0xca,
	0xf8, 0xc1, 0x2f, 0x75, 0x57, 0xc9, 0x1f, 0xed, 0x26, 0x1e, 0x09, 0x47, 0xb7, 0xff, 0xb9, 0x6d,
	0x34, 0x74, 0xa9, 0xdf, 0x5b, 0xc1, 0x36, 0xad, 0x60, 0x9f, 0xad, 0x60, 0x2f, 0x9d, 0x98, 0x6c,
	0x3a, 0x31, 0xf9, 0xe8, 0xc4, 0xe4, 0x2e, 0xd3, 0x40, 0xd5, 0xe3, 0x32, 0x29, 0xf0, 0x21, 0x1d,
	0xc9, 0xf1, 0x16, 0x9d, 0x8e, 0x98, 0x78, 0xe0, 0xc4, 0x43, 0xce, 0xab, 0xd4, 0x9b, 0xfe, 0xda,
	0x28, 0xb7, 0xdc, 0xed, 0x2f, 0xe0, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x28, 0x2a, 0xab, 0x4b,
	0x21, 0x02, 0x00, 0x00,
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
	SnapshotProposal(ctx context.Context, in *MsgSnapshotProposal, opts ...grpc.CallOption) (*MsgSnapshotProposalResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SnapshotProposal(ctx context.Context, in *MsgSnapshotProposal, opts ...grpc.CallOption) (*MsgSnapshotProposalResponse, error) {
	out := new(MsgSnapshotProposalResponse)
	err := c.cc.Invoke(ctx, "/twilightproject.validitygadgetchain.validitygadgetchain.Msg/SnapshotProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SnapshotProposal(context.Context, *MsgSnapshotProposal) (*MsgSnapshotProposalResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SnapshotProposal(ctx context.Context, req *MsgSnapshotProposal) (*MsgSnapshotProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SnapshotProposal not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SnapshotProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSnapshotProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SnapshotProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twilightproject.validitygadgetchain.validitygadgetchain.Msg/SnapshotProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SnapshotProposal(ctx, req.(*MsgSnapshotProposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "twilightproject.validitygadgetchain.validitygadgetchain.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SnapshotProposal",
			Handler:    _Msg_SnapshotProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "validitygadgetchain/tx.proto",
}

func (m *MsgSnapshotProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSnapshotProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSnapshotProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LogSnap) > 0 {
		i -= len(m.LogSnap)
		copy(dAtA[i:], m.LogSnap)
		i = encodeVarintTx(dAtA, i, uint64(len(m.LogSnap)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Iteration) > 0 {
		i -= len(m.Iteration)
		copy(dAtA[i:], m.Iteration)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Iteration)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSnapshotProposalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSnapshotProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSnapshotProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
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
func (m *MsgSnapshotProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Iteration)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.LogSnap)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSnapshotProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSnapshotProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSnapshotProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSnapshotProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Iteration", wireType)
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
			m.Iteration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogSnap", wireType)
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
			m.LogSnap = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSnapshotProposalResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSnapshotProposalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSnapshotProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
