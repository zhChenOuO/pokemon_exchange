// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/v2/commonpb/pagination.proto

package common

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Pagination struct {
	Page       int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage    int64 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	TotalCount int64 `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	TotalPage  int64 `protobuf:"varint,4,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85f055e4b337911, []int{0}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return m.Size()
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pagination) GetPerPage() int64 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *Pagination) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *Pagination) GetTotalPage() int64 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

func init() {
	proto.RegisterType((*Pagination)(nil), "commonpbv2.Pagination")
	golang_proto.RegisterType((*Pagination)(nil), "commonpbv2.Pagination")
}

func init() { proto.RegisterFile("pb/v2/commonpb/pagination.proto", fileDescriptor_f85f055e4b337911) }
func init() {
	golang_proto.RegisterFile("pb/v2/commonpb/pagination.proto", fileDescriptor_f85f055e4b337911)
}

var fileDescriptor_f85f055e4b337911 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x63, 0x5a, 0xf1, 0xe7, 0x60, 0xca, 0x54, 0x90, 0x88, 0x11, 0x13, 0x0b, 0xb1, 0xd4,
	0x6e, 0x8c, 0x20, 0xf6, 0x8a, 0x91, 0xa5, 0xb2, 0x83, 0xb9, 0x5a, 0x24, 0x3e, 0xcb, 0x38, 0xe5,
	0x19, 0xd8, 0x18, 0x59, 0xe9, 0x93, 0x30, 0x32, 0xf6, 0x11, 0x50, 0x78, 0x11, 0x94, 0x8b, 0x40,
	0xea, 0x76, 0xdf, 0xef, 0xe7, 0xcf, 0xd2, 0x1d, 0xc8, 0x60, 0xd4, 0x6a, 0xaa, 0x2a, 0x6a, 0x1a,
	0xf2, 0xc1, 0xa8, 0xa0, 0xd1, 0x79, 0x9d, 0x1c, 0xf9, 0x32, 0x44, 0x4a, 0x94, 0xc3, 0x9f, 0x5a,
	0x4d, 0x4f, 0x2e, 0xd1, 0xa5, 0x65, 0x6b, 0xca, 0x8a, 0x1a, 0x85, 0x84, 0xa4, 0xf8, 0x89, 0x69,
	0x1f, 0x39, 0x71, 0xe0, 0x69, 0xa8, 0x9e, 0xbf, 0x0a, 0x80, 0xf9, 0xff, 0x7f, 0x79, 0x0e, 0xe3,
	0xa0, 0xd1, 0x4e, 0xc4, 0x99, 0xb8, 0x18, 0xdd, 0xf1, 0x9c, 0x1f, 0xc3, 0x7e, 0xb0, 0x71, 0xc1,
	0x7c, 0x87, 0xf9, 0x5e, 0xb0, 0x71, 0xde, 0x2b, 0x09, 0x87, 0x89, 0x92, 0xae, 0x17, 0x15, 0xb5,
	0x3e, 0x4d, 0x46, 0x6c, 0x81, 0xd1, 0x4d, 0x4f, 0xf2, 0x53, 0x18, 0xd2, 0xd0, 0x1e, 0xb3, 0x3f,
	0x60, 0xd2, 0xf7, 0xaf, 0x8e, 0x36, 0x1f, 0x32, 0x7b, 0x5b, 0xcb, 0xec, 0x7d, 0x2d, 0xb3, 0xeb,
	0xdb, 0xaf, 0xae, 0x10, 0x9b, 0xae, 0x10, 0xdf, 0x5d, 0x21, 0x3e, 0x7f, 0x0a, 0x71, 0x3f, 0x43,
	0x97, 0x6a, 0x6d, 0xca, 0x67, 0x57, 0x3f, 0x45, 0x7a, 0xb0, 0xfd, 0x56, 0x65, 0x7a, 0x51, 0x48,
	0xb5, 0xf6, 0xa8, 0x90, 0xc2, 0xd2, 0x46, 0xb5, 0x7d, 0x1e, 0xb3, 0xcb, 0x9b, 0xcd, 0x7e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x80, 0xc0, 0x42, 0xf0, 0x37, 0x01, 0x00, 0x00,
}

func (m *Pagination) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pagination) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pagination) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalPage != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.TotalPage))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalCount != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.TotalCount))
		i--
		dAtA[i] = 0x18
	}
	if m.PerPage != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.PerPage))
		i--
		dAtA[i] = 0x10
	}
	if m.Page != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPagination(dAtA []byte, offset int, v uint64) int {
	offset -= sovPagination(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pagination) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Page != 0 {
		n += 1 + sovPagination(uint64(m.Page))
	}
	if m.PerPage != 0 {
		n += 1 + sovPagination(uint64(m.PerPage))
	}
	if m.TotalCount != 0 {
		n += 1 + sovPagination(uint64(m.TotalCount))
	}
	if m.TotalPage != 0 {
		n += 1 + sovPagination(uint64(m.TotalPage))
	}
	return n
}

func sovPagination(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPagination(x uint64) (n int) {
	return sovPagination(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pagination) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: Pagination: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pagination: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerPage", wireType)
			}
			m.PerPage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PerPage |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCount", wireType)
			}
			m.TotalCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalPage", wireType)
			}
			m.TotalPage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalPage |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPagination
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPagination
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
func skipPagination(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
				return 0, ErrInvalidLengthPagination
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPagination
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPagination
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPagination        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPagination          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPagination = fmt.Errorf("proto: unexpected end of group")
)
