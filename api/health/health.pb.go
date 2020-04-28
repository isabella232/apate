// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: health/health.proto

package health

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status int32

const (
	Status_UNKNOWN      Status = 0
	Status_HEALTHY      Status = 1
	Status_UNHEALTHY    Status = 2
	Status_DISCONNECTED Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "HEALTHY",
		2: "UNHEALTHY",
		3: "DISCONNECTED",
	}
	Status_value = map[string]int32{
		"UNKNOWN":      0,
		"HEALTHY":      1,
		"UNHEALTHY":    2,
		"DISCONNECTED": 3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_health_health_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_health_health_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{0}
}

type NodeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeUuid string `protobuf:"bytes,1,opt,name=node_uuid,json=nodeUuid,proto3" json:"node_uuid,omitempty"`
	Status   Status `protobuf:"varint,2,opt,name=Status,proto3,enum=apate.health.Status" json:"Status,omitempty"`
}

func (x *NodeStatus) Reset() {
	*x = NodeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatus) ProtoMessage() {}

func (x *NodeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatus.ProtoReflect.Descriptor instead.
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{0}
}

func (x *NodeStatus) GetNodeUuid() string {
	if x != nil {
		return x.NodeUuid
	}
	return ""
}

func (x *NodeStatus) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

var File_health_health_proto protoreflect.FileDescriptor

var file_health_health_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x61, 0x74, 0x65, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x57, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70,
	0x61, 0x74, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x43, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x55, 0x4e, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x32, 0x50,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x46, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x61, 0x74, 0x65,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x2d, 0x72, 0x65, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x64, 0x63, 0x2d, 0x65, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2d, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_health_proto_rawDescOnce sync.Once
	file_health_health_proto_rawDescData = file_health_health_proto_rawDesc
)

func file_health_health_proto_rawDescGZIP() []byte {
	file_health_health_proto_rawDescOnce.Do(func() {
		file_health_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_health_proto_rawDescData)
	})
	return file_health_health_proto_rawDescData
}

var file_health_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_health_health_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_health_health_proto_goTypes = []interface{}{
	(Status)(0),         // 0: apate.health.Status
	(*NodeStatus)(nil),  // 1: apate.health.NodeStatus
	(*empty.Empty)(nil), // 2: google.protobuf.Empty
}
var file_health_health_proto_depIdxs = []int32{
	0, // 0: apate.health.NodeStatus.Status:type_name -> apate.health.Status
	1, // 1: apate.health.Health.HealthStream:input_type -> apate.health.NodeStatus
	2, // 2: apate.health.Health.HealthStream:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_health_health_proto_init() }
func file_health_health_proto_init() {
	if File_health_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatus); i {
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
			RawDescriptor: file_health_health_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_health_proto_goTypes,
		DependencyIndexes: file_health_health_proto_depIdxs,
		EnumInfos:         file_health_health_proto_enumTypes,
		MessageInfos:      file_health_health_proto_msgTypes,
	}.Build()
	File_health_health_proto = out.File
	file_health_health_proto_rawDesc = nil
	file_health_health_proto_goTypes = nil
	file_health_health_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	HealthStream(ctx context.Context, opts ...grpc.CallOption) (Health_HealthStreamClient, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) HealthStream(ctx context.Context, opts ...grpc.CallOption) (Health_HealthStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Health_serviceDesc.Streams[0], "/apate.health.Health/HealthStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthHealthStreamClient{stream}
	return x, nil
}

type Health_HealthStreamClient interface {
	Send(*NodeStatus) error
	Recv() (*empty.Empty, error)
	grpc.ClientStream
}

type healthHealthStreamClient struct {
	grpc.ClientStream
}

func (x *healthHealthStreamClient) Send(m *NodeStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *healthHealthStreamClient) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	HealthStream(Health_HealthStreamServer) error
}

// UnimplementedHealthServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (*UnimplementedHealthServer) HealthStream(Health_HealthStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HealthStream not implemented")
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_HealthStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HealthServer).HealthStream(&healthHealthStreamServer{stream})
}

type Health_HealthStreamServer interface {
	Send(*empty.Empty) error
	Recv() (*NodeStatus, error)
	grpc.ServerStream
}

type healthHealthStreamServer struct {
	grpc.ServerStream
}

func (x *healthHealthStreamServer) Send(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *healthHealthStreamServer) Recv() (*NodeStatus, error) {
	m := new(NodeStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apate.health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HealthStream",
			Handler:       _Health_HealthStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "health/health.proto",
}
