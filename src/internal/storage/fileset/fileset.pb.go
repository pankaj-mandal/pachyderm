// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/storage/fileset/fileset.proto

package fileset

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	index "github.com/pachyderm/pachyderm/v2/src/internal/storage/fileset/index"
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

type Metadata struct {
	// Types that are valid to be assigned to Value:
	//	*Metadata_Primitive
	//	*Metadata_Composite
	Value                isMetadata_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_22dc3e2e3017d669, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

type isMetadata_Value interface {
	isMetadata_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Metadata_Primitive struct {
	Primitive *Primitive `protobuf:"bytes,1,opt,name=primitive,proto3,oneof" json:"primitive,omitempty"`
}
type Metadata_Composite struct {
	Composite *Composite `protobuf:"bytes,2,opt,name=composite,proto3,oneof" json:"composite,omitempty"`
}

func (*Metadata_Primitive) isMetadata_Value() {}
func (*Metadata_Composite) isMetadata_Value() {}

func (m *Metadata) GetValue() isMetadata_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Metadata) GetPrimitive() *Primitive {
	if x, ok := m.GetValue().(*Metadata_Primitive); ok {
		return x.Primitive
	}
	return nil
}

func (m *Metadata) GetComposite() *Composite {
	if x, ok := m.GetValue().(*Metadata_Composite); ok {
		return x.Composite
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Metadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Metadata_Primitive)(nil),
		(*Metadata_Composite)(nil),
	}
}

type Composite struct {
	Layers               []string `protobuf:"bytes,1,rep,name=layers,proto3" json:"layers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Composite) Reset()         { *m = Composite{} }
func (m *Composite) String() string { return proto.CompactTextString(m) }
func (*Composite) ProtoMessage()    {}
func (*Composite) Descriptor() ([]byte, []int) {
	return fileDescriptor_22dc3e2e3017d669, []int{1}
}
func (m *Composite) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Composite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Composite.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Composite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Composite.Merge(m, src)
}
func (m *Composite) XXX_Size() int {
	return m.Size()
}
func (m *Composite) XXX_DiscardUnknown() {
	xxx_messageInfo_Composite.DiscardUnknown(m)
}

var xxx_messageInfo_Composite proto.InternalMessageInfo

func (m *Composite) GetLayers() []string {
	if m != nil {
		return m.Layers
	}
	return nil
}

type Primitive struct {
	Deletive             *index.Index `protobuf:"bytes,1,opt,name=deletive,proto3" json:"deletive,omitempty"`
	Additive             *index.Index `protobuf:"bytes,2,opt,name=additive,proto3" json:"additive,omitempty"`
	SizeBytes            int64        `protobuf:"varint,3,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Primitive) Reset()         { *m = Primitive{} }
func (m *Primitive) String() string { return proto.CompactTextString(m) }
func (*Primitive) ProtoMessage()    {}
func (*Primitive) Descriptor() ([]byte, []int) {
	return fileDescriptor_22dc3e2e3017d669, []int{2}
}
func (m *Primitive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Primitive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Primitive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Primitive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Primitive.Merge(m, src)
}
func (m *Primitive) XXX_Size() int {
	return m.Size()
}
func (m *Primitive) XXX_DiscardUnknown() {
	xxx_messageInfo_Primitive.DiscardUnknown(m)
}

var xxx_messageInfo_Primitive proto.InternalMessageInfo

func (m *Primitive) GetDeletive() *index.Index {
	if m != nil {
		return m.Deletive
	}
	return nil
}

func (m *Primitive) GetAdditive() *index.Index {
	if m != nil {
		return m.Additive
	}
	return nil
}

func (m *Primitive) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

type TestCacheValue struct {
	FileSetId            string   `protobuf:"bytes,1,opt,name=file_set_id,json=fileSetId,proto3" json:"file_set_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestCacheValue) Reset()         { *m = TestCacheValue{} }
func (m *TestCacheValue) String() string { return proto.CompactTextString(m) }
func (*TestCacheValue) ProtoMessage()    {}
func (*TestCacheValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_22dc3e2e3017d669, []int{3}
}
func (m *TestCacheValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestCacheValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestCacheValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestCacheValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestCacheValue.Merge(m, src)
}
func (m *TestCacheValue) XXX_Size() int {
	return m.Size()
}
func (m *TestCacheValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TestCacheValue.DiscardUnknown(m)
}

var xxx_messageInfo_TestCacheValue proto.InternalMessageInfo

func (m *TestCacheValue) GetFileSetId() string {
	if m != nil {
		return m.FileSetId
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "fileset.Metadata")
	proto.RegisterType((*Composite)(nil), "fileset.Composite")
	proto.RegisterType((*Primitive)(nil), "fileset.Primitive")
	proto.RegisterType((*TestCacheValue)(nil), "fileset.TestCacheValue")
}

func init() {
	proto.RegisterFile("internal/storage/fileset/fileset.proto", fileDescriptor_22dc3e2e3017d669)
}

var fileDescriptor_22dc3e2e3017d669 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0xb6, 0xe5, 0x6b, 0xbb, 0x5b, 0xf1, 0xb0, 0x07, 0x09, 0x82, 0xa1, 0x44, 0x90, 0xe0,
	0x21, 0x91, 0x7a, 0xf7, 0xd0, 0x5e, 0x2c, 0x28, 0x48, 0x14, 0x0f, 0x5e, 0xca, 0x36, 0x3b, 0xb6,
	0x0b, 0x69, 0x37, 0xec, 0x4e, 0x8b, 0x55, 0xf0, 0xf7, 0x79, 0xf4, 0x27, 0x48, 0x7e, 0x89, 0x6c,
	0xd3, 0xc4, 0x22, 0xf6, 0x32, 0xbb, 0x33, 0xef, 0x3d, 0xe6, 0x3d, 0x86, 0x9d, 0xa9, 0x05, 0x82,
	0x59, 0x88, 0x2c, 0xb6, 0xa8, 0x8d, 0x98, 0x42, 0xfc, 0xac, 0x32, 0xb0, 0x80, 0xd5, 0x1b, 0xe5,
	0x46, 0xa3, 0xe6, 0xed, 0x6d, 0x7b, 0x7c, 0xbe, 0x57, 0xa0, 0x16, 0x12, 0x5e, 0xca, 0x5a, 0x8a,
	0x82, 0x37, 0xd6, 0xb9, 0x05, 0x14, 0x52, 0xa0, 0xe0, 0x7d, 0x46, 0x73, 0xa3, 0xe6, 0x0a, 0xd5,
	0x0a, 0x3c, 0xd2, 0x23, 0x61, 0xb7, 0xcf, 0xa3, 0x6a, 0xc7, 0x5d, 0x85, 0x5c, 0xff, 0x4b, 0x7e,
	0x68, 0x4e, 0x93, 0xea, 0x79, 0xae, 0xad, 0x42, 0xf0, 0x1a, 0xbf, 0x34, 0xc3, 0x0a, 0x71, 0x9a,
	0x9a, 0x36, 0x68, 0xb3, 0xff, 0x2b, 0x91, 0x2d, 0x21, 0x38, 0x65, 0xb4, 0xa6, 0xf0, 0x23, 0xd6,
	0xca, 0xc4, 0x1a, 0x8c, 0xf5, 0x48, 0xaf, 0x19, 0xd2, 0x64, 0xdb, 0x05, 0xef, 0x8c, 0xd6, 0xbb,
	0x79, 0xc8, 0x3a, 0x12, 0x32, 0xd8, 0x71, 0x78, 0x10, 0x95, 0x71, 0x46, 0xae, 0x26, 0x35, 0xea,
	0x98, 0x42, 0xca, 0x32, 0x4b, 0xe3, 0x2f, 0x66, 0x85, 0xf2, 0x13, 0xc6, 0xac, 0x7a, 0x85, 0xf1,
	0x64, 0x8d, 0x60, 0xbd, 0x66, 0x8f, 0x84, 0xcd, 0x84, 0xba, 0xc9, 0xc0, 0x0d, 0x82, 0x0b, 0x76,
	0xf8, 0x00, 0x16, 0x87, 0x22, 0x9d, 0xc1, 0xa3, 0xb3, 0xcd, 0x7d, 0xd6, 0x75, 0x09, 0xc7, 0x16,
	0x70, 0xac, 0xe4, 0xc6, 0x07, 0x4d, 0xa8, 0x1b, 0xdd, 0x03, 0x8e, 0xe4, 0xe0, 0xe6, 0xa3, 0xf0,
	0xc9, 0x67, 0xe1, 0x93, 0xaf, 0xc2, 0x27, 0x4f, 0x57, 0x53, 0x85, 0xb3, 0xe5, 0x24, 0x4a, 0xf5,
	0x3c, 0xce, 0x45, 0x3a, 0x5b, 0x4b, 0x30, 0xbb, 0xbf, 0x55, 0x3f, 0xb6, 0x26, 0x8d, 0xf7, 0xdd,
	0x6c, 0xd2, 0xda, 0x1c, 0xea, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xfe, 0xe9, 0x14, 0x07,
	0x02, 0x00, 0x00,
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Metadata_Primitive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata_Primitive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Primitive != nil {
		{
			size, err := m.Primitive.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFileset(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Metadata_Composite) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata_Composite) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Composite != nil {
		{
			size, err := m.Composite.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFileset(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Composite) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Composite) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Composite) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Layers) > 0 {
		for iNdEx := len(m.Layers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Layers[iNdEx])
			copy(dAtA[i:], m.Layers[iNdEx])
			i = encodeVarintFileset(dAtA, i, uint64(len(m.Layers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Primitive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Primitive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Primitive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SizeBytes != 0 {
		i = encodeVarintFileset(dAtA, i, uint64(m.SizeBytes))
		i--
		dAtA[i] = 0x18
	}
	if m.Additive != nil {
		{
			size, err := m.Additive.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFileset(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Deletive != nil {
		{
			size, err := m.Deletive.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFileset(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestCacheValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestCacheValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestCacheValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.FileSetId) > 0 {
		i -= len(m.FileSetId)
		copy(dAtA[i:], m.FileSetId)
		i = encodeVarintFileset(dAtA, i, uint64(len(m.FileSetId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFileset(dAtA []byte, offset int, v uint64) int {
	offset -= sovFileset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Metadata_Primitive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Primitive != nil {
		l = m.Primitive.Size()
		n += 1 + l + sovFileset(uint64(l))
	}
	return n
}
func (m *Metadata_Composite) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Composite != nil {
		l = m.Composite.Size()
		n += 1 + l + sovFileset(uint64(l))
	}
	return n
}
func (m *Composite) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Layers) > 0 {
		for _, s := range m.Layers {
			l = len(s)
			n += 1 + l + sovFileset(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Primitive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Deletive != nil {
		l = m.Deletive.Size()
		n += 1 + l + sovFileset(uint64(l))
	}
	if m.Additive != nil {
		l = m.Additive.Size()
		n += 1 + l + sovFileset(uint64(l))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovFileset(uint64(m.SizeBytes))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TestCacheValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FileSetId)
	if l > 0 {
		n += 1 + l + sovFileset(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFileset(x uint64) (n int) {
	return sovFileset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileset
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Primitive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
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
				return ErrInvalidLengthFileset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFileset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Primitive{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Metadata_Primitive{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Composite", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
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
				return ErrInvalidLengthFileset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFileset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Composite{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Metadata_Composite{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFileset
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
func (m *Composite) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileset
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
			return fmt.Errorf("proto: Composite: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Composite: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Layers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
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
				return ErrInvalidLengthFileset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Layers = append(m.Layers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFileset
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
func (m *Primitive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileset
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
			return fmt.Errorf("proto: Primitive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Primitive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deletive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
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
				return ErrInvalidLengthFileset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFileset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Deletive == nil {
				m.Deletive = &index.Index{}
			}
			if err := m.Deletive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Additive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
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
				return ErrInvalidLengthFileset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFileset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Additive == nil {
				m.Additive = &index.Index{}
			}
			if err := m.Additive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFileset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFileset
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
func (m *TestCacheValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileset
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
			return fmt.Errorf("proto: TestCacheValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestCacheValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileSetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
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
				return ErrInvalidLengthFileset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileSetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFileset
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
func skipFileset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFileset
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
					return 0, ErrIntOverflowFileset
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
					return 0, ErrIntOverflowFileset
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
				return 0, ErrInvalidLengthFileset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFileset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFileset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFileset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFileset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFileset = fmt.Errorf("proto: unexpected end of group")
)
