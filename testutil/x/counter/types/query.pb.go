// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/counter/v1/query.proto

package types

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// QueryGetCountRequest defines the request type for querying x/mock count.
type QueryGetCountRequest struct {
}

func (m *QueryGetCountRequest) Reset()         { *m = QueryGetCountRequest{} }
func (m *QueryGetCountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCountRequest) ProtoMessage()    {}
func (*QueryGetCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd21727562626c9f, []int{0}
}
func (m *QueryGetCountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCountRequest.Merge(m, src)
}
func (m *QueryGetCountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCountRequest proto.InternalMessageInfo

// QueryGetCountResponse defines the response type for querying x/mock count.
type QueryGetCountResponse struct {
	TotalCount int64                  `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Amount     *cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount,omitempty"`
}

func (m *QueryGetCountResponse) Reset()         { *m = QueryGetCountResponse{} }
func (m *QueryGetCountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCountResponse) ProtoMessage()    {}
func (*QueryGetCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd21727562626c9f, []int{1}
}
func (m *QueryGetCountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCountResponse.Merge(m, src)
}
func (m *QueryGetCountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCountResponse proto.InternalMessageInfo

func (m *QueryGetCountResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryGetCountRequest)(nil), "cosmos.counter.v1.QueryGetCountRequest")
	proto.RegisterType((*QueryGetCountResponse)(nil), "cosmos.counter.v1.QueryGetCountResponse")
}

func init() { proto.RegisterFile("cosmos/counter/v1/query.proto", fileDescriptor_bd21727562626c9f) }

var fileDescriptor_bd21727562626c9f = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0x2c,
	0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x48, 0xeb, 0x41, 0xa5,
	0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2, 0xfa, 0x20, 0x16, 0x44, 0xa1,
	0x94, 0x24, 0x44, 0x61, 0x3c, 0x44, 0x02, 0xaa, 0x0b, 0xcc, 0x51, 0x12, 0xe3, 0x12, 0x09, 0x04,
	0x19, 0xe9, 0x9e, 0x5a, 0xe2, 0x0c, 0x32, 0x26, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0xa9,
	0x92, 0x4b, 0x14, 0x4d, 0xbc, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x9e, 0x8b, 0xbb, 0x24,
	0xbf, 0x24, 0x31, 0x27, 0x1e, 0x6c, 0xab, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x17, 0x58,
	0x08, 0xac, 0x50, 0xc8, 0x9e, 0x8b, 0x2d, 0x31, 0x17, 0x2c, 0xc7, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9,
	0xa4, 0x7e, 0xeb, 0x9e, 0xbc, 0x28, 0xc4, 0xce, 0xe2, 0x94, 0x6c, 0xbd, 0xcc, 0x7c, 0xfd, 0xdc,
	0xc4, 0x92, 0x0c, 0x3d, 0xcf, 0xbc, 0x92, 0x4b, 0x5b, 0x74, 0xb9, 0xa0, 0x8e, 0xf1, 0xcc, 0x2b,
	0x09, 0x82, 0x6a, 0x33, 0x4a, 0xe3, 0x62, 0x05, 0x5b, 0x2d, 0x14, 0xcb, 0xc5, 0x01, 0xb3, 0x5e,
	0x48, 0x5d, 0x0f, 0xc3, 0xb3, 0x7a, 0xd8, 0x1c, 0x2e, 0xa5, 0x41, 0x58, 0x21, 0xc4, 0x27, 0x4e,
	0xfe, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9a, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x8f, 0x02, 0x10, 0xa5, 0x5b, 0x9c, 0x92, 0xad,
	0x5f, 0x92, 0x5a, 0x5c, 0x52, 0x5a, 0x92, 0x99, 0xa3, 0x5f, 0x01, 0x8f, 0x98, 0x92, 0xca, 0x82,
	0xd4, 0xe2, 0x24, 0x36, 0x70, 0x90, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x0e, 0xb7,
	0x96, 0xb7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// GetCount queries the parameters of x/Counter module.
	GetCount(ctx context.Context, in *QueryGetCountRequest, opts ...grpc.CallOption) (*QueryGetCountResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetCount(ctx context.Context, in *QueryGetCountRequest, opts ...grpc.CallOption) (*QueryGetCountResponse, error) {
	out := new(QueryGetCountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.counter.v1.Query/GetCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// GetCount queries the parameters of x/Counter module.
	GetCount(context.Context, *QueryGetCountRequest) (*QueryGetCountResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetCount(ctx context.Context, req *QueryGetCountRequest) (*QueryGetCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCount not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.counter.v1.Query/GetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetCount(ctx, req.(*QueryGetCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.counter.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCount",
			Handler:    _Query_GetCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/counter/v1/query.proto",
}

func (m *QueryGetCountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetCountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != nil {
		{
			size := m.Amount.Size()
			i -= size
			if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.TotalCount != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.TotalCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetCountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetCountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalCount != 0 {
		n += 1 + sovQuery(uint64(m.TotalCount))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetCountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCount", wireType)
			}
			m.TotalCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.Amount = &v
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
