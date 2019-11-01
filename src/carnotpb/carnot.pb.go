// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/carnotpb/carnot.proto

package carnotpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	proto1 "pixielabs.ai/pixielabs/src/common/uuid/proto"
	proto2 "pixielabs.ai/pixielabs/src/table_store/proto"
	reflect "reflect"
	strings "strings"
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

type RowBatchRequest struct {
	Address       string               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	QueryID       *proto1.UUID         `protobuf:"bytes,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	DestinationId string               `protobuf:"bytes,3,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	RowBatch      *proto2.RowBatchData `protobuf:"bytes,4,opt,name=row_batch,json=rowBatch,proto3" json:"row_batch,omitempty"`
}

func (m *RowBatchRequest) Reset()      { *m = RowBatchRequest{} }
func (*RowBatchRequest) ProtoMessage() {}
func (*RowBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9054907422204f4e, []int{0}
}
func (m *RowBatchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RowBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RowBatchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RowBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowBatchRequest.Merge(m, src)
}
func (m *RowBatchRequest) XXX_Size() int {
	return m.Size()
}
func (m *RowBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RowBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RowBatchRequest proto.InternalMessageInfo

func (m *RowBatchRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RowBatchRequest) GetQueryID() *proto1.UUID {
	if m != nil {
		return m.QueryID
	}
	return nil
}

func (m *RowBatchRequest) GetDestinationId() string {
	if m != nil {
		return m.DestinationId
	}
	return ""
}

func (m *RowBatchRequest) GetRowBatch() *proto2.RowBatchData {
	if m != nil {
		return m.RowBatch
	}
	return nil
}

type RowBatchResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *RowBatchResponse) Reset()      { *m = RowBatchResponse{} }
func (*RowBatchResponse) ProtoMessage() {}
func (*RowBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9054907422204f4e, []int{1}
}
func (m *RowBatchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RowBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RowBatchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RowBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowBatchResponse.Merge(m, src)
}
func (m *RowBatchResponse) XXX_Size() int {
	return m.Size()
}
func (m *RowBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RowBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RowBatchResponse proto.InternalMessageInfo

func (m *RowBatchResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RowBatchResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RowBatchRequest)(nil), "pl.carnotpb.RowBatchRequest")
	proto.RegisterType((*RowBatchResponse)(nil), "pl.carnotpb.RowBatchResponse")
}

func init() { proto.RegisterFile("src/carnotpb/carnot.proto", fileDescriptor_9054907422204f4e) }

var fileDescriptor_9054907422204f4e = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x63, 0x40, 0xb4, 0x75, 0x75, 0xdc, 0x29, 0x53, 0xa9, 0xc0, 0x54, 0x95, 0x4e, 0xea,
	0x82, 0x2b, 0x1d, 0x03, 0x1b, 0x43, 0x54, 0x21, 0x45, 0x4c, 0xe4, 0xb8, 0x85, 0xa5, 0xf2, 0xbf,
	0x4b, 0x23, 0x25, 0x71, 0xce, 0x76, 0xee, 0xc4, 0xc6, 0x47, 0xe0, 0x63, 0xf0, 0x51, 0xd8, 0xe8,
	0x78, 0x13, 0xa2, 0xee, 0xc2, 0xd8, 0x8f, 0x80, 0xec, 0xc4, 0xa2, 0x42, 0xba, 0xed, 0x7d, 0xf2,
	0xfe, 0xfc, 0xbe, 0xcf, 0x93, 0x17, 0x3e, 0xd7, 0x8a, 0x2d, 0x19, 0x51, 0xb5, 0x34, 0x0d, 0xed,
	0x0b, 0xdc, 0x28, 0x69, 0x64, 0x3c, 0x6e, 0x4a, 0x1c, 0x3a, 0xd3, 0xd7, 0x79, 0x61, 0x36, 0x2d,
	0xc5, 0x4c, 0x56, 0xcb, 0x5c, 0xe6, 0x72, 0xe9, 0x19, 0xda, 0x5e, 0x7b, 0xe5, 0x85, 0xaf, 0xba,
	0xb7, 0xd3, 0x99, 0x1f, 0x2b, 0xab, 0x4a, 0xd6, 0xcb, 0xb6, 0x2d, 0x78, 0x87, 0xfb, 0xb2, 0x27,
	0xe6, 0x8e, 0x30, 0x84, 0x96, 0x62, 0xad, 0x8d, 0x54, 0xa2, 0x27, 0x34, 0xdb, 0x88, 0x8a, 0x74,
	0xcc, 0xfc, 0x27, 0x80, 0xa7, 0x99, 0xbc, 0x4b, 0x88, 0x61, 0x9b, 0x4c, 0xdc, 0xb4, 0x42, 0x9b,
	0x78, 0x02, 0x07, 0x84, 0x73, 0x25, 0xb4, 0x9e, 0x80, 0x19, 0x58, 0x8c, 0xb2, 0x20, 0xe3, 0xb7,
	0x70, 0x78, 0xd3, 0x0a, 0xf5, 0x65, 0x5d, 0xf0, 0xc9, 0xa3, 0x19, 0x58, 0x8c, 0x2f, 0x4e, 0x71,
	0x53, 0x62, 0xb7, 0xb3, 0xa1, 0xf8, 0xea, 0x2a, 0x5d, 0x25, 0x63, 0xfb, 0xeb, 0xd5, 0xe0, 0xa3,
	0x83, 0xd2, 0x55, 0x36, 0xf0, 0x74, 0xca, 0xe3, 0x73, 0xf8, 0x8c, 0x0b, 0x6d, 0x8a, 0x9a, 0x98,
	0x42, 0xd6, 0xee, 0xf9, 0x63, 0x3f, 0xf9, 0xe4, 0xe8, 0x6b, 0xca, 0xe3, 0x04, 0x8e, 0x94, 0xbc,
	0x5b, 0x53, 0xe7, 0x66, 0xf2, 0xc4, 0x2f, 0x38, 0x77, 0x0b, 0x8e, 0x42, 0xe0, 0xce, 0x7e, 0x43,
	0x71, 0xb0, 0xbd, 0x22, 0x86, 0x64, 0x43, 0xd5, 0xab, 0xf9, 0x7b, 0x78, 0xf6, 0x2f, 0x90, 0x6e,
	0x64, 0xad, 0x85, 0x4b, 0xa4, 0x5b, 0xc6, 0x42, 0xa2, 0x61, 0x16, 0xa4, 0xeb, 0x54, 0x42, 0x6b,
	0x92, 0x0b, 0x1f, 0x68, 0x94, 0x05, 0x79, 0xc1, 0xe1, 0xc9, 0x07, 0x51, 0xde, 0x16, 0xf5, 0xa5,
	0x50, 0xb7, 0x05, 0x13, 0xf1, 0x25, 0x3c, 0xfb, 0xa4, 0x48, 0xad, 0xaf, 0x85, 0x0a, 0x0b, 0xe2,
	0x17, 0xf8, 0xe8, 0x82, 0xf8, 0xbf, 0x1f, 0x39, 0x7d, 0xf9, 0x40, 0xb7, 0x73, 0x35, 0x8f, 0x16,
	0x20, 0x79, 0xb7, 0xdd, 0xa1, 0xe8, 0x7e, 0x87, 0xa2, 0xc3, 0x0e, 0x81, 0xaf, 0x16, 0x81, 0xef,
	0x16, 0x81, 0x1f, 0x16, 0x81, 0xad, 0x45, 0xe0, 0xb7, 0x45, 0xe0, 0x8f, 0x45, 0xd1, 0xc1, 0x22,
	0xf0, 0x6d, 0x8f, 0xa2, 0xed, 0x1e, 0x45, 0xf7, 0x7b, 0x14, 0x7d, 0x1e, 0x86, 0xa1, 0xf4, 0xa9,
	0x3f, 0xe3, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0xc5, 0xc7, 0x87, 0x65, 0x02, 0x00,
	0x00,
}

func (this *RowBatchRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RowBatchRequest)
	if !ok {
		that2, ok := that.(RowBatchRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if !this.QueryID.Equal(that1.QueryID) {
		return false
	}
	if this.DestinationId != that1.DestinationId {
		return false
	}
	if !this.RowBatch.Equal(that1.RowBatch) {
		return false
	}
	return true
}
func (this *RowBatchResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RowBatchResponse)
	if !ok {
		that2, ok := that.(RowBatchResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Success != that1.Success {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *RowBatchRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&carnotpb.RowBatchRequest{")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	if this.QueryID != nil {
		s = append(s, "QueryID: "+fmt.Sprintf("%#v", this.QueryID)+",\n")
	}
	s = append(s, "DestinationId: "+fmt.Sprintf("%#v", this.DestinationId)+",\n")
	if this.RowBatch != nil {
		s = append(s, "RowBatch: "+fmt.Sprintf("%#v", this.RowBatch)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RowBatchResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&carnotpb.RowBatchResponse{")
	s = append(s, "Success: "+fmt.Sprintf("%#v", this.Success)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCarnot(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KelvinServiceClient is the client API for KelvinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KelvinServiceClient interface {
	TransferRowBatch(ctx context.Context, opts ...grpc.CallOption) (KelvinService_TransferRowBatchClient, error)
}

type kelvinServiceClient struct {
	cc *grpc.ClientConn
}

func NewKelvinServiceClient(cc *grpc.ClientConn) KelvinServiceClient {
	return &kelvinServiceClient{cc}
}

func (c *kelvinServiceClient) TransferRowBatch(ctx context.Context, opts ...grpc.CallOption) (KelvinService_TransferRowBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KelvinService_serviceDesc.Streams[0], "/pl.carnotpb.KelvinService/TransferRowBatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &kelvinServiceTransferRowBatchClient{stream}
	return x, nil
}

type KelvinService_TransferRowBatchClient interface {
	Send(*RowBatchRequest) error
	CloseAndRecv() (*RowBatchResponse, error)
	grpc.ClientStream
}

type kelvinServiceTransferRowBatchClient struct {
	grpc.ClientStream
}

func (x *kelvinServiceTransferRowBatchClient) Send(m *RowBatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kelvinServiceTransferRowBatchClient) CloseAndRecv() (*RowBatchResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RowBatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KelvinServiceServer is the server API for KelvinService service.
type KelvinServiceServer interface {
	TransferRowBatch(KelvinService_TransferRowBatchServer) error
}

// UnimplementedKelvinServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKelvinServiceServer struct {
}

func (*UnimplementedKelvinServiceServer) TransferRowBatch(srv KelvinService_TransferRowBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method TransferRowBatch not implemented")
}

func RegisterKelvinServiceServer(s *grpc.Server, srv KelvinServiceServer) {
	s.RegisterService(&_KelvinService_serviceDesc, srv)
}

func _KelvinService_TransferRowBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KelvinServiceServer).TransferRowBatch(&kelvinServiceTransferRowBatchServer{stream})
}

type KelvinService_TransferRowBatchServer interface {
	SendAndClose(*RowBatchResponse) error
	Recv() (*RowBatchRequest, error)
	grpc.ServerStream
}

type kelvinServiceTransferRowBatchServer struct {
	grpc.ServerStream
}

func (x *kelvinServiceTransferRowBatchServer) SendAndClose(m *RowBatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kelvinServiceTransferRowBatchServer) Recv() (*RowBatchRequest, error) {
	m := new(RowBatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _KelvinService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pl.carnotpb.KelvinService",
	HandlerType: (*KelvinServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransferRowBatch",
			Handler:       _KelvinService_TransferRowBatch_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "src/carnotpb/carnot.proto",
}

func (m *RowBatchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RowBatchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RowBatchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RowBatch != nil {
		{
			size, err := m.RowBatch.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCarnot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.DestinationId) > 0 {
		i -= len(m.DestinationId)
		copy(dAtA[i:], m.DestinationId)
		i = encodeVarintCarnot(dAtA, i, uint64(len(m.DestinationId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.QueryID != nil {
		{
			size, err := m.QueryID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCarnot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCarnot(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RowBatchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RowBatchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RowBatchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintCarnot(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCarnot(dAtA []byte, offset int, v uint64) int {
	offset -= sovCarnot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RowBatchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCarnot(uint64(l))
	}
	if m.QueryID != nil {
		l = m.QueryID.Size()
		n += 1 + l + sovCarnot(uint64(l))
	}
	l = len(m.DestinationId)
	if l > 0 {
		n += 1 + l + sovCarnot(uint64(l))
	}
	if m.RowBatch != nil {
		l = m.RowBatch.Size()
		n += 1 + l + sovCarnot(uint64(l))
	}
	return n
}

func (m *RowBatchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCarnot(uint64(l))
	}
	return n
}

func sovCarnot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCarnot(x uint64) (n int) {
	return sovCarnot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RowBatchRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RowBatchRequest{`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`QueryID:` + strings.Replace(fmt.Sprintf("%v", this.QueryID), "UUID", "proto1.UUID", 1) + `,`,
		`DestinationId:` + fmt.Sprintf("%v", this.DestinationId) + `,`,
		`RowBatch:` + strings.Replace(fmt.Sprintf("%v", this.RowBatch), "RowBatchData", "proto2.RowBatchData", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RowBatchResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RowBatchResponse{`,
		`Success:` + fmt.Sprintf("%v", this.Success) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCarnot(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RowBatchRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCarnot
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
			return fmt.Errorf("proto: RowBatchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RowBatchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
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
				return ErrInvalidLengthCarnot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCarnot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
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
				return ErrInvalidLengthCarnot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCarnot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.QueryID == nil {
				m.QueryID = &proto1.UUID{}
			}
			if err := m.QueryID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
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
				return ErrInvalidLengthCarnot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCarnot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RowBatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
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
				return ErrInvalidLengthCarnot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCarnot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RowBatch == nil {
				m.RowBatch = &proto2.RowBatchData{}
			}
			if err := m.RowBatch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCarnot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCarnot
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCarnot
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
func (m *RowBatchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCarnot
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
			return fmt.Errorf("proto: RowBatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RowBatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
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
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarnot
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
				return ErrInvalidLengthCarnot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCarnot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCarnot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCarnot
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCarnot
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
func skipCarnot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCarnot
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
					return 0, ErrIntOverflowCarnot
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
					return 0, ErrIntOverflowCarnot
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
				return 0, ErrInvalidLengthCarnot
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCarnot
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCarnot
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
				next, err := skipCarnot(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCarnot
				}
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
	ErrInvalidLengthCarnot = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCarnot   = fmt.Errorf("proto: integer overflow")
)
