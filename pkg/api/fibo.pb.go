// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.18.0
// source: api/proto/fibo.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FiboReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *FiboReq) Reset() {
	*x = FiboReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_fibo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiboReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiboReq) ProtoMessage() {}

func (x *FiboReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_fibo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiboReq.ProtoReflect.Descriptor instead.
func (*FiboReq) Descriptor() ([]byte, []int) {
	return file_api_proto_fibo_proto_rawDescGZIP(), []int{0}
}

func (x *FiboReq) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *FiboReq) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type FiboResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []int64 `protobuf:"varint,1,rep,packed,name=result,proto3" json:"result,omitempty"`
}

func (x *FiboResp) Reset() {
	*x = FiboResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_fibo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiboResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiboResp) ProtoMessage() {}

func (x *FiboResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_fibo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiboResp.ProtoReflect.Descriptor instead.
func (*FiboResp) Descriptor() ([]byte, []int) {
	return file_api_proto_fibo_proto_rawDescGZIP(), []int{1}
}

func (x *FiboResp) GetResult() []int64 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_api_proto_fibo_proto protoreflect.FileDescriptor

var file_api_proto_fibo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x62, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x25, 0x0a, 0x07, 0x46,
	0x69, 0x62, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x01, 0x79, 0x22, 0x22, 0x0a, 0x08, 0x46, 0x69, 0x62, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x30, 0x0a, 0x07, 0x46, 0x69, 0x62, 0x6f, 0x4e, 0x75,
	0x6d, 0x12, 0x25, 0x0a, 0x04, 0x46, 0x69, 0x62, 0x6f, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x46, 0x69, 0x62, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69,
	0x62, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_fibo_proto_rawDescOnce sync.Once
	file_api_proto_fibo_proto_rawDescData = file_api_proto_fibo_proto_rawDesc
)

func file_api_proto_fibo_proto_rawDescGZIP() []byte {
	file_api_proto_fibo_proto_rawDescOnce.Do(func() {
		file_api_proto_fibo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_fibo_proto_rawDescData)
	})
	return file_api_proto_fibo_proto_rawDescData
}

var file_api_proto_fibo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_fibo_proto_goTypes = []interface{}{
	(*FiboReq)(nil),  // 0: api.FiboReq
	(*FiboResp)(nil), // 1: api.FiboResp
}
var file_api_proto_fibo_proto_depIdxs = []int32{
	0, // 0: api.FiboNum.Fibo:input_type -> api.FiboReq
	1, // 1: api.FiboNum.Fibo:output_type -> api.FiboResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_fibo_proto_init() }
func file_api_proto_fibo_proto_init() {
	if File_api_proto_fibo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_fibo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiboReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_fibo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiboResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_fibo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_fibo_proto_goTypes,
		DependencyIndexes: file_api_proto_fibo_proto_depIdxs,
		MessageInfos:      file_api_proto_fibo_proto_msgTypes,
	}.Build()
	File_api_proto_fibo_proto = out.File
	file_api_proto_fibo_proto_rawDesc = nil
	file_api_proto_fibo_proto_goTypes = nil
	file_api_proto_fibo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FiboNumClient is the client API for FiboNum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FiboNumClient interface {
	Fibo(ctx context.Context, in *FiboReq, opts ...grpc.CallOption) (*FiboResp, error)
}

type fiboNumClient struct {
	cc grpc.ClientConnInterface
}

func NewFiboNumClient(cc grpc.ClientConnInterface) FiboNumClient {
	return &fiboNumClient{cc}
}

func (c *fiboNumClient) Fibo(ctx context.Context, in *FiboReq, opts ...grpc.CallOption) (*FiboResp, error) {
	out := new(FiboResp)
	err := c.cc.Invoke(ctx, "/api.FiboNum/Fibo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FiboNumServer is the server API for FiboNum service.
type FiboNumServer interface {
	Fibo(context.Context, *FiboReq) (*FiboResp, error)
}

// UnimplementedFiboNumServer can be embedded to have forward compatible implementations.
type UnimplementedFiboNumServer struct {
}

func (*UnimplementedFiboNumServer) Fibo(context.Context, *FiboReq) (*FiboResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fibo not implemented")
}

func RegisterFiboNumServer(s *grpc.Server, srv FiboNumServer) {
	s.RegisterService(&_FiboNum_serviceDesc, srv)
}

func _FiboNum_Fibo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FiboReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiboNumServer).Fibo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FiboNum/Fibo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiboNumServer).Fibo(ctx, req.(*FiboReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FiboNum_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.FiboNum",
	HandlerType: (*FiboNumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fibo",
			Handler:    _FiboNum_Fibo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/fibo.proto",
}
