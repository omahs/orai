// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/provider/types/query_dsource.proto

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

// DataSourceInfoReq is the request type for the Query/DataSourceInfo RPC method
type DataSourceInfoReq struct {
	// address is the address of the contract to query
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DataSourceInfoReq) Reset()         { *m = DataSourceInfoReq{} }
func (m *DataSourceInfoReq) String() string { return proto.CompactTextString(m) }
func (*DataSourceInfoReq) ProtoMessage()    {}
func (*DataSourceInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87619281e010ebb, []int{0}
}
func (m *DataSourceInfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataSourceInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataSourceInfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataSourceInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSourceInfoReq.Merge(m, src)
}
func (m *DataSourceInfoReq) XXX_Size() int {
	return m.Size()
}
func (m *DataSourceInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSourceInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_DataSourceInfoReq proto.InternalMessageInfo

func (m *DataSourceInfoReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// DataSourceInfoRes is the response type for the Query/DataSourceInfo RPC method
type DataSourceInfoRes struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Owner is the address who is allowed to make further changes to the data source.
	Owner       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty"`
	Contract    string                                        `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	Description string                                        `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Fees        github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,6,rep,name=fees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fees" json:"fees"`
}

func (m *DataSourceInfoRes) Reset()         { *m = DataSourceInfoRes{} }
func (m *DataSourceInfoRes) String() string { return proto.CompactTextString(m) }
func (*DataSourceInfoRes) ProtoMessage()    {}
func (*DataSourceInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87619281e010ebb, []int{1}
}
func (m *DataSourceInfoRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataSourceInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataSourceInfoRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataSourceInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSourceInfoRes.Merge(m, src)
}
func (m *DataSourceInfoRes) XXX_Size() int {
	return m.Size()
}
func (m *DataSourceInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSourceInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_DataSourceInfoRes proto.InternalMessageInfo

func (m *DataSourceInfoRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataSourceInfoRes) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *DataSourceInfoRes) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *DataSourceInfoRes) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DataSourceInfoRes) GetFees() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fees
	}
	return nil
}

// ListDataSourcesReq is the request type for the Query/ListDataSources RPC method
type ListDataSourcesReq struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Page  int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *ListDataSourcesReq) Reset()         { *m = ListDataSourcesReq{} }
func (m *ListDataSourcesReq) String() string { return proto.CompactTextString(m) }
func (*ListDataSourcesReq) ProtoMessage()    {}
func (*ListDataSourcesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87619281e010ebb, []int{2}
}
func (m *ListDataSourcesReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListDataSourcesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListDataSourcesReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListDataSourcesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDataSourcesReq.Merge(m, src)
}
func (m *ListDataSourcesReq) XXX_Size() int {
	return m.Size()
}
func (m *ListDataSourcesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDataSourcesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListDataSourcesReq proto.InternalMessageInfo

func (m *ListDataSourcesReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListDataSourcesReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListDataSourcesReq) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// ListDataSourcesRes is the response type for the Query/ListDataSources RPC method
type ListDataSourcesRes struct {
	AIDataSources []AIDataSource `protobuf:"bytes,1,rep,name=AIDataSources,proto3" json:"AIDataSources" json:"data_sources"`
	Count         int64          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *ListDataSourcesRes) Reset()         { *m = ListDataSourcesRes{} }
func (m *ListDataSourcesRes) String() string { return proto.CompactTextString(m) }
func (*ListDataSourcesRes) ProtoMessage()    {}
func (*ListDataSourcesRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87619281e010ebb, []int{3}
}
func (m *ListDataSourcesRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListDataSourcesRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListDataSourcesRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListDataSourcesRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDataSourcesRes.Merge(m, src)
}
func (m *ListDataSourcesRes) XXX_Size() int {
	return m.Size()
}
func (m *ListDataSourcesRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDataSourcesRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListDataSourcesRes proto.InternalMessageInfo

func (m *ListDataSourcesRes) GetAIDataSources() []AIDataSource {
	if m != nil {
		return m.AIDataSources
	}
	return nil
}

func (m *ListDataSourcesRes) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*DataSourceInfoReq)(nil), "oraichain.orai.provider.DataSourceInfoReq")
	proto.RegisterType((*DataSourceInfoRes)(nil), "oraichain.orai.provider.DataSourceInfoRes")
	proto.RegisterType((*ListDataSourcesReq)(nil), "oraichain.orai.provider.ListDataSourcesReq")
	proto.RegisterType((*ListDataSourcesRes)(nil), "oraichain.orai.provider.ListDataSourcesRes")
}

func init() {
	proto.RegisterFile("x/provider/types/query_dsource.proto", fileDescriptor_d87619281e010ebb)
}

var fileDescriptor_d87619281e010ebb = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x6e, 0x13, 0x41,
	0x10, 0xf5, 0xc5, 0x4e, 0x04, 0x6b, 0x28, 0x58, 0x22, 0x71, 0xb8, 0xb8, 0xb3, 0x4e, 0x20, 0x2c,
	0x44, 0x76, 0x65, 0xe8, 0x68, 0x90, 0x4d, 0x24, 0x14, 0x89, 0xea, 0xe8, 0x68, 0xac, 0xf5, 0xde,
	0xc6, 0x59, 0xc0, 0x3b, 0x97, 0x9d, 0x75, 0xc0, 0x7f, 0xc1, 0x17, 0xd0, 0xd1, 0xf0, 0x25, 0x29,
	0x53, 0x52, 0x1d, 0xc8, 0xfe, 0x83, 0x94, 0x54, 0xe8, 0x76, 0x4d, 0x70, 0x48, 0x2c, 0xd1, 0xdc,
	0xbd, 0x99, 0x7d, 0x33, 0xf3, 0xe6, 0x69, 0xc8, 0x83, 0x4f, 0xbc, 0xb4, 0x70, 0xa2, 0x0b, 0x65,
	0xb9, 0x9b, 0x97, 0x0a, 0xf9, 0xf1, 0x4c, 0xd9, 0xf9, 0xa8, 0x40, 0x98, 0x59, 0xa9, 0x58, 0x69,
	0xc1, 0x01, 0xbd, 0x07, 0x56, 0x68, 0x79, 0x24, 0xb4, 0x61, 0x35, 0x62, 0x7f, 0x4a, 0x3a, 0xbb,
	0x13, 0x98, 0x80, 0xe7, 0xf0, 0x1a, 0x05, 0x7a, 0x27, 0x91, 0x80, 0x53, 0x40, 0x3e, 0x16, 0xa8,
	0xf8, 0x49, 0x7f, 0xac, 0x9c, 0xe8, 0x73, 0x09, 0xda, 0xac, 0xde, 0xd3, 0x2b, 0x43, 0xfd, 0x77,
	0x54, 0x60, 0x20, 0x64, 0x8f, 0xc8, 0x9d, 0x7d, 0xe1, 0xc4, 0x1b, 0xaf, 0xe1, 0xc0, 0x1c, 0x42,
	0xae, 0x8e, 0x29, 0x25, 0x2d, 0x23, 0xa6, 0x2a, 0x8e, 0xba, 0x51, 0xef, 0x66, 0xee, 0x71, 0xf6,
	0x65, 0xeb, 0x2a, 0x13, 0xaf, 0x63, 0xd2, 0x57, 0x64, 0x1b, 0x3e, 0x1a, 0x65, 0xe3, 0xad, 0x6e,
	0xd4, 0xbb, 0x35, 0xec, 0xff, 0xaa, 0xd2, 0xbd, 0x89, 0x76, 0x47, 0xb3, 0x31, 0x93, 0x30, 0xe5,
	0x2b, 0xc5, 0xe1, 0xb7, 0x87, 0xc5, 0xfb, 0x20, 0x89, 0x0d, 0xa4, 0x1c, 0x14, 0x85, 0x55, 0x88,
	0x79, 0xa8, 0xa7, 0x1d, 0x72, 0x43, 0x82, 0x71, 0x56, 0x48, 0x17, 0x37, 0xfd, 0x80, 0x8b, 0x98,
	0x76, 0x49, 0xbb, 0x50, 0x28, 0xad, 0x2e, 0x9d, 0x06, 0x13, 0xb7, 0xfc, 0xf3, 0x7a, 0x8a, 0x1a,
	0xd2, 0x3a, 0x54, 0x0a, 0xe3, 0x9d, 0x6e, 0xb3, 0xd7, 0x7e, 0x7a, 0x9f, 0x85, 0x81, 0xac, 0x76,
	0x8a, 0xad, 0x9c, 0x62, 0x2f, 0x41, 0x9b, 0xe1, 0x8b, 0xd3, 0x2a, 0x6d, 0x9c, 0x57, 0x69, 0xfb,
	0x1d, 0x82, 0x79, 0x9e, 0xd5, 0x45, 0xd9, 0xb7, 0x1f, 0x69, 0xef, 0x3f, 0x34, 0xd7, 0xf5, 0x98,
	0xfb, 0x39, 0x59, 0x4e, 0xe8, 0x6b, 0x8d, 0xee, 0xaf, 0x47, 0xb8, 0xc1, 0xca, 0x3a, 0x57, 0x8a,
	0x89, 0xf2, 0xfe, 0x34, 0x73, 0x8f, 0xe9, 0x2e, 0xd9, 0xfe, 0xa0, 0xa7, 0x3a, 0x2c, 0xda, 0xcc,
	0x43, 0x90, 0x7d, 0x8d, 0xae, 0x69, 0x8a, 0x74, 0x4e, 0x6e, 0x0f, 0x0e, 0xd6, 0x72, 0x71, 0xe4,
	0x77, 0x7c, 0xc8, 0x36, 0x1c, 0x0f, 0x5b, 0x67, 0x0f, 0x9f, 0xd4, 0xfb, 0x2e, 0xaa, 0xf4, 0x72,
	0x8f, 0xf3, 0x2a, 0xbd, 0x1b, 0x0c, 0x28, 0x84, 0x13, 0xa3, 0x70, 0x94, 0x98, 0xe5, 0x97, 0x59,
	0xb5, 0x4e, 0x09, 0x33, 0xe3, 0x56, 0xe2, 0x43, 0x30, 0xdc, 0x3f, 0x5d, 0x24, 0xd1, 0xd9, 0x22,
	0x89, 0x7e, 0x2e, 0x92, 0xe8, 0xf3, 0x32, 0x69, 0x9c, 0x2d, 0x93, 0xc6, 0xf7, 0x65, 0xd2, 0x78,
	0xfb, 0x78, 0xcd, 0xc5, 0x0b, 0x75, 0x1e, 0xf1, 0x7f, 0x4f, 0x73, 0xbc, 0xe3, 0x4f, 0xf2, 0xd9,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x6a, 0x6c, 0xaa, 0x2a, 0x03, 0x00, 0x00,
}

func (m *DataSourceInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataSourceInfoReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataSourceInfoReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQueryDsource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataSourceInfoRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataSourceInfoRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataSourceInfoRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintQueryDsource(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintQueryDsource(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintQueryDsource(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQueryDsource(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQueryDsource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListDataSourcesReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListDataSourcesReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListDataSourcesReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		i = encodeVarintQueryDsource(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x18
	}
	if m.Page != 0 {
		i = encodeVarintQueryDsource(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQueryDsource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListDataSourcesRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListDataSourcesRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListDataSourcesRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintQueryDsource(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AIDataSources) > 0 {
		for iNdEx := len(m.AIDataSources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AIDataSources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryDsource(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryDsource(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryDsource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataSourceInfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQueryDsource(uint64(l))
	}
	return n
}

func (m *DataSourceInfoRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQueryDsource(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQueryDsource(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovQueryDsource(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovQueryDsource(uint64(l))
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovQueryDsource(uint64(l))
		}
	}
	return n
}

func (m *ListDataSourcesReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQueryDsource(uint64(l))
	}
	if m.Page != 0 {
		n += 1 + sovQueryDsource(uint64(m.Page))
	}
	if m.Limit != 0 {
		n += 1 + sovQueryDsource(uint64(m.Limit))
	}
	return n
}

func (m *ListDataSourcesRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AIDataSources) > 0 {
		for _, e := range m.AIDataSources {
			l = e.Size()
			n += 1 + l + sovQueryDsource(uint64(l))
		}
	}
	if m.Count != 0 {
		n += 1 + sovQueryDsource(uint64(m.Count))
	}
	return n
}

func sovQueryDsource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryDsource(x uint64) (n int) {
	return sovQueryDsource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataSourceInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDsource
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
			return fmt.Errorf("proto: DataSourceInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataSourceInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
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
				return ErrInvalidLengthQueryDsource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDsource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDsource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDsource
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
func (m *DataSourceInfoRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDsource
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
			return fmt.Errorf("proto: DataSourceInfoRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataSourceInfoRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
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
				return ErrInvalidLengthQueryDsource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDsource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
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
				return ErrInvalidLengthQueryDsource
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDsource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
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
				return ErrInvalidLengthQueryDsource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDsource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
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
				return ErrInvalidLengthQueryDsource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDsource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
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
				return ErrInvalidLengthQueryDsource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDsource
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
			skippy, err := skipQueryDsource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDsource
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
func (m *ListDataSourcesReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDsource
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
			return fmt.Errorf("proto: ListDataSourcesReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListDataSourcesReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
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
				return ErrInvalidLengthQueryDsource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDsource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDsource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDsource
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
func (m *ListDataSourcesRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDsource
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
			return fmt.Errorf("proto: ListDataSourcesRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListDataSourcesRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AIDataSources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
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
				return ErrInvalidLengthQueryDsource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDsource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AIDataSources = append(m.AIDataSources, AIDataSource{})
			if err := m.AIDataSources[len(m.AIDataSources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDsource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDsource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDsource
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
func skipQueryDsource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryDsource
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
					return 0, ErrIntOverflowQueryDsource
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
					return 0, ErrIntOverflowQueryDsource
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
				return 0, ErrInvalidLengthQueryDsource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryDsource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryDsource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryDsource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryDsource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryDsource = fmt.Errorf("proto: unexpected end of group")
)
