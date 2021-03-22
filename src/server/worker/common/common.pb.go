// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/worker/common/common.proto

package common

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	pfs "github.com/pachyderm/pachyderm/v2/src/pfs"
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

type Input struct {
	FileInfo             *pfs.FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo,proto3" json:"file_info,omitempty"`
	ParentCommit         *pfs.Commit   `protobuf:"bytes,5,opt,name=parent_commit,json=parentCommit,proto3" json:"parent_commit,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	JoinOn               string        `protobuf:"bytes,8,opt,name=join_on,json=joinOn,proto3" json:"join_on,omitempty"`
	OuterJoin            bool          `protobuf:"varint,11,opt,name=outer_join,json=outerJoin,proto3" json:"outer_join,omitempty"`
	GroupBy              string        `protobuf:"bytes,10,opt,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	Lazy                 bool          `protobuf:"varint,3,opt,name=lazy,proto3" json:"lazy,omitempty"`
	Branch               string        `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	GitURL               string        `protobuf:"bytes,6,opt,name=git_url,json=gitUrl,proto3" json:"git_url,omitempty"`
	EmptyFiles           bool          `protobuf:"varint,7,opt,name=empty_files,json=emptyFiles,proto3" json:"empty_files,omitempty"`
	S3                   bool          `protobuf:"varint,9,opt,name=s3,proto3" json:"s3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fb6c79ddd9db74, []int{0}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Input.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return m.Size()
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetFileInfo() *pfs.FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

func (m *Input) GetParentCommit() *pfs.Commit {
	if m != nil {
		return m.ParentCommit
	}
	return nil
}

func (m *Input) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Input) GetJoinOn() string {
	if m != nil {
		return m.JoinOn
	}
	return ""
}

func (m *Input) GetOuterJoin() bool {
	if m != nil {
		return m.OuterJoin
	}
	return false
}

func (m *Input) GetGroupBy() string {
	if m != nil {
		return m.GroupBy
	}
	return ""
}

func (m *Input) GetLazy() bool {
	if m != nil {
		return m.Lazy
	}
	return false
}

func (m *Input) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Input) GetGitURL() string {
	if m != nil {
		return m.GitURL
	}
	return ""
}

func (m *Input) GetEmptyFiles() bool {
	if m != nil {
		return m.EmptyFiles
	}
	return false
}

func (m *Input) GetS3() bool {
	if m != nil {
		return m.S3
	}
	return false
}

func init() {
	proto.RegisterType((*Input)(nil), "common.Input")
}

func init() { proto.RegisterFile("server/worker/common/common.proto", fileDescriptor_91fb6c79ddd9db74) }

var fileDescriptor_91fb6c79ddd9db74 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0xeb, 0xd3, 0x30,
	0x1c, 0xc6, 0x69, 0xdd, 0xfa, 0xe7, 0xbb, 0x3f, 0x87, 0x20, 0x1a, 0x07, 0x6e, 0x53, 0x2f, 0x03,
	0xa1, 0xc5, 0xed, 0xe6, 0x71, 0x82, 0x3a, 0x11, 0x84, 0xc2, 0x2e, 0x5e, 0x4a, 0x5b, 0xd3, 0x2e,
	0xda, 0x26, 0x21, 0x49, 0x27, 0xf5, 0x15, 0x7a, 0xdc, 0x2b, 0x10, 0xe9, 0x2b, 0x91, 0x24, 0x3b,
	0x78, 0xf8, 0x9d, 0xf2, 0x3c, 0x9f, 0xe4, 0xc9, 0x43, 0xf2, 0x85, 0x17, 0x8a, 0xc8, 0x2b, 0x91,
	0xe9, 0x4f, 0x2e, 0x7f, 0x10, 0x99, 0x56, 0xbc, 0xeb, 0x38, 0xbb, 0x2f, 0x89, 0x90, 0x5c, 0x73,
	0x14, 0x38, 0xb7, 0x5a, 0x88, 0x5a, 0xa5, 0xa2, 0x56, 0x0e, 0xaf, 0x1e, 0x37, 0xbc, 0xe1, 0x56,
	0xa6, 0x46, 0x39, 0xfa, 0xf2, 0xe6, 0xc3, 0xf4, 0xc4, 0x44, 0xaf, 0xd1, 0x6b, 0x88, 0x6b, 0xda,
	0x92, 0x9c, 0xb2, 0x9a, 0x63, 0x6f, 0xeb, 0xed, 0x66, 0xfb, 0x65, 0x22, 0x6a, 0xb5, 0x4f, 0xde,
	0xd3, 0x96, 0x9c, 0x58, 0xcd, 0xb3, 0xa8, 0xbe, 0x2b, 0xf4, 0x06, 0x16, 0xa2, 0x90, 0x84, 0xe9,
	0xdc, 0x94, 0x51, 0x8d, 0xa7, 0x36, 0x30, 0x77, 0x81, 0x77, 0x96, 0x65, 0x73, 0x77, 0xc4, 0x39,
	0x84, 0x60, 0xc2, 0x8a, 0x8e, 0x60, 0x7f, 0xeb, 0xed, 0xe2, 0xcc, 0x6a, 0xf4, 0x14, 0xc2, 0xef,
	0x9c, 0xb2, 0x9c, 0x33, 0x1c, 0x59, 0x1c, 0x18, 0xfb, 0x85, 0xa1, 0xe7, 0x00, 0xbc, 0xd7, 0x44,
	0xe6, 0xc6, 0xe3, 0xd9, 0xd6, 0xdb, 0x45, 0x59, 0x6c, 0xc9, 0x27, 0x4e, 0x19, 0x7a, 0x06, 0x51,
	0x23, 0x79, 0x2f, 0xf2, 0x72, 0xc0, 0x60, 0x83, 0xa1, 0xf5, 0xc7, 0xc1, 0xd4, 0xb4, 0xc5, 0xaf,
	0x01, 0x3f, 0xb2, 0x19, 0xab, 0xd1, 0x13, 0x08, 0x4a, 0x59, 0xb0, 0xea, 0x82, 0x27, 0xae, 0xc5,
	0x39, 0xf4, 0x0a, 0xc2, 0x86, 0xea, 0xbc, 0x97, 0x2d, 0x0e, 0xcc, 0xc6, 0x11, 0xc6, 0x3f, 0x9b,
	0xe0, 0x03, 0xd5, 0xe7, 0xec, 0x73, 0x16, 0x34, 0x54, 0x9f, 0x65, 0x8b, 0x36, 0x30, 0x23, 0x9d,
	0xd0, 0x43, 0x6e, 0x1e, 0xaf, 0x70, 0x68, 0xef, 0x05, 0x8b, 0xcc, 0xc7, 0x28, 0xb4, 0x04, 0x5f,
	0x1d, 0x70, 0x6c, 0xb9, 0xaf, 0x0e, 0xc7, 0x8f, 0xbf, 0xc7, 0xb5, 0x77, 0x1b, 0xd7, 0xde, 0xdf,
	0x71, 0xed, 0x7d, 0x7d, 0xdb, 0x50, 0x7d, 0xe9, 0xcb, 0xa4, 0xe2, 0x5d, 0x2a, 0x8a, 0xea, 0x32,
	0x7c, 0x23, 0xf2, 0x7f, 0x75, 0xdd, 0xa7, 0x4a, 0x56, 0xe9, 0x43, 0x63, 0x2d, 0x03, 0x3b, 0xa3,
	0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x38, 0x4d, 0xb3, 0xf5, 0x01, 0x00, 0x00,
}

func (m *Input) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Input) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OuterJoin {
		i--
		if m.OuterJoin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.GroupBy) > 0 {
		i -= len(m.GroupBy)
		copy(dAtA[i:], m.GroupBy)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.GroupBy)))
		i--
		dAtA[i] = 0x52
	}
	if m.S3 {
		i--
		if m.S3 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.JoinOn) > 0 {
		i -= len(m.JoinOn)
		copy(dAtA[i:], m.JoinOn)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.JoinOn)))
		i--
		dAtA[i] = 0x42
	}
	if m.EmptyFiles {
		i--
		if m.EmptyFiles {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.GitURL) > 0 {
		i -= len(m.GitURL)
		copy(dAtA[i:], m.GitURL)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.GitURL)))
		i--
		dAtA[i] = 0x32
	}
	if m.ParentCommit != nil {
		{
			size, err := m.ParentCommit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Branch) > 0 {
		i -= len(m.Branch)
		copy(dAtA[i:], m.Branch)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Branch)))
		i--
		dAtA[i] = 0x22
	}
	if m.Lazy {
		i--
		if m.Lazy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.FileInfo != nil {
		{
			size, err := m.FileInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Input) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FileInfo != nil {
		l = m.FileInfo.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Lazy {
		n += 2
	}
	l = len(m.Branch)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.ParentCommit != nil {
		l = m.ParentCommit.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.GitURL)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.EmptyFiles {
		n += 2
	}
	l = len(m.JoinOn)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.S3 {
		n += 2
	}
	l = len(m.GroupBy)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.OuterJoin {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Input) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Input: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FileInfo == nil {
				m.FileInfo = &pfs.FileInfo{}
			}
			if err := m.FileInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lazy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Lazy = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Branch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Branch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ParentCommit == nil {
				m.ParentCommit = &pfs.Commit{}
			}
			if err := m.ParentCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmptyFiles", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EmptyFiles = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoinOn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JoinOn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field S3", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.S3 = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OuterJoin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OuterJoin = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
