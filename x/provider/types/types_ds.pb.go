// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/provider/types/types_ds.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type AIDataSource struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	// Owner is the address who is allowed to make further changes to the data source.
	Owner       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty"`
	Description string                                        `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Fees        github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,5,rep,name=fees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fees" json:"transaction_fee"`
}

func (m *AIDataSource) Reset()         { *m = AIDataSource{} }
func (m *AIDataSource) String() string { return proto.CompactTextString(m) }
func (*AIDataSource) ProtoMessage()    {}
func (*AIDataSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_246318b67ba64c0e, []int{0}
}
func (m *AIDataSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AIDataSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AIDataSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AIDataSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AIDataSource.Merge(m, src)
}
func (m *AIDataSource) XXX_Size() int {
	return m.Size()
}
func (m *AIDataSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AIDataSource.DiscardUnknown(m)
}

var xxx_messageInfo_AIDataSource proto.InternalMessageInfo

func (m *AIDataSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AIDataSource) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *AIDataSource) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *AIDataSource) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AIDataSource) GetFees() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fees
	}
	return nil
}

func init() {
	proto.RegisterType((*AIDataSource)(nil), "oraichain.orai.provider.AIDataSource")
}

func init() { proto.RegisterFile("x/provider/types/types_ds.proto", fileDescriptor_246318b67ba64c0e) }

var fileDescriptor_246318b67ba64c0e = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x6e, 0xea, 0x30,
	0x14, 0x86, 0x13, 0x08, 0x57, 0xf7, 0x06, 0xa6, 0xe8, 0xaa, 0x4d, 0x19, 0x92, 0x88, 0x29, 0xaa,
	0x84, 0x2d, 0xda, 0x8d, 0x0d, 0x8a, 0x54, 0xb5, 0x23, 0xdd, 0xba, 0x20, 0xc7, 0x31, 0xe0, 0x56,
	0xf8, 0x44, 0xb6, 0xa1, 0xe5, 0x2d, 0x3a, 0x76, 0xec, 0xcc, 0x93, 0x30, 0x32, 0x76, 0xa2, 0x15,
	0xbc, 0x41, 0xc7, 0x4e, 0x55, 0x62, 0x8a, 0x50, 0xa7, 0x2e, 0xc9, 0xe7, 0x5f, 0x3e, 0xff, 0x39,
	0xbf, 0x8f, 0x1b, 0x3e, 0xe2, 0x4c, 0xc2, 0x8c, 0xa7, 0x4c, 0x62, 0x3d, 0xcf, 0x98, 0x32, 0xdf,
	0x41, 0xaa, 0x50, 0x26, 0x41, 0x83, 0x77, 0x0c, 0x92, 0x70, 0x3a, 0x26, 0x5c, 0xa0, 0x9c, 0xd0,
	0xf7, 0xed, 0x7a, 0x40, 0x41, 0x4d, 0x40, 0xe1, 0x84, 0x28, 0x86, 0x67, 0xad, 0x84, 0x69, 0xd2,
	0xc2, 0x14, 0xb8, 0x30, 0x85, 0xf5, 0xff, 0x23, 0x18, 0x41, 0x81, 0x38, 0x27, 0xa3, 0x36, 0x16,
	0x25, 0xb7, 0xd6, 0xb9, 0xea, 0x11, 0x4d, 0x6e, 0x60, 0x2a, 0x29, 0xf3, 0x3c, 0xd7, 0x11, 0x64,
	0xc2, 0x7c, 0x3b, 0xb2, 0xe3, 0x7f, 0xfd, 0x82, 0xbd, 0xba, 0xfb, 0x97, 0x82, 0xd0, 0x92, 0x50,
	0xed, 0x97, 0x0a, 0x7d, 0x7f, 0xf6, 0x2e, 0xdd, 0x0a, 0x3c, 0x08, 0x26, 0xfd, 0x72, 0x64, 0xc7,
	0xb5, 0x6e, 0xeb, 0x73, 0x1d, 0x36, 0x47, 0x5c, 0x8f, 0xa7, 0x09, 0xa2, 0x30, 0xc1, 0xbb, 0xa1,
	0xcc, 0xaf, 0xa9, 0xd2, 0x7b, 0x93, 0x07, 0x75, 0x28, 0xed, 0xa4, 0xa9, 0x64, 0x4a, 0xf5, 0x4d,
	0xbd, 0x17, 0xb9, 0xd5, 0x94, 0x29, 0x2a, 0x79, 0xa6, 0x39, 0x08, 0xdf, 0x29, 0xfa, 0x1c, 0x4a,
	0xde, 0xdc, 0x75, 0x86, 0x8c, 0x29, 0xbf, 0x12, 0x95, 0xe3, 0xea, 0xd9, 0x09, 0x32, 0xa6, 0x28,
	0x0f, 0x8c, 0x76, 0x81, 0xd1, 0x05, 0x70, 0xd1, 0xbd, 0x5e, 0xae, 0x43, 0xeb, 0x63, 0x1d, 0x1e,
	0xdd, 0x29, 0x10, 0xed, 0x86, 0x96, 0x44, 0x28, 0x42, 0x73, 0x8f, 0xc1, 0x90, 0xb1, 0xc6, 0xe2,
	0x2d, 0x8c, 0x7f, 0x31, 0x62, 0x6e, 0xa5, 0xfa, 0x45, 0xcb, 0xb6, 0xf3, 0xfc, 0x12, 0xda, 0xdd,
	0xde, 0x72, 0x13, 0xd8, 0xab, 0x4d, 0x60, 0xbf, 0x6f, 0x02, 0xfb, 0x69, 0x1b, 0x58, 0xab, 0x6d,
	0x60, 0xbd, 0x6e, 0x03, 0xeb, 0xf6, 0xf4, 0xc0, 0x6f, 0xbf, 0xa0, 0x82, 0xf0, 0xcf, 0x85, 0x26,
	0x7f, 0x8a, 0x97, 0x3f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x14, 0x4c, 0xba, 0x69, 0xeb, 0x01,
	0x00, 0x00,
}

func (m *AIDataSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AIDataSource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AIDataSource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesDs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTypesDs(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypesDs(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintTypesDs(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypesDs(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypesDs(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypesDs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AIDataSource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypesDs(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovTypesDs(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypesDs(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTypesDs(uint64(l))
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovTypesDs(uint64(l))
		}
	}
	return n
}

func sovTypesDs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypesDs(x uint64) (n int) {
	return sovTypesDs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AIDataSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesDs
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
			return fmt.Errorf("proto: AIDataSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AIDataSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
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
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
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
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
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
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, types.Coin{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesDs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypesDs
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
func skipTypesDs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypesDs
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
					return 0, ErrIntOverflowTypesDs
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
					return 0, ErrIntOverflowTypesDs
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
				return 0, ErrInvalidLengthTypesDs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypesDs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypesDs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypesDs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypesDs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypesDs = fmt.Errorf("proto: unexpected end of group")
)
