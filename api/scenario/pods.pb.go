// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: scenario/pods.proto

package scenario

import (
	proto "github.com/golang/protobuf/proto"
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

// The status of the pods
// See corev1.PodPhase
type PodStatus int32

const (
	PodStatus_POD_PENDING   PodStatus = 0
	PodStatus_POD_RUNNING   PodStatus = 1
	PodStatus_POD_SUCCEEDED PodStatus = 2
	PodStatus_POD_FAILED    PodStatus = 3
	PodStatus_POD_UNKNOWN   PodStatus = 4
)

// Enum value maps for PodStatus.
var (
	PodStatus_name = map[int32]string{
		0: "POD_PENDING",
		1: "POD_RUNNING",
		2: "POD_SUCCEEDED",
		3: "POD_FAILED",
		4: "POD_UNKNOWN",
	}
	PodStatus_value = map[string]int32{
		"POD_PENDING":   0,
		"POD_RUNNING":   1,
		"POD_SUCCEEDED": 2,
		"POD_FAILED":    3,
		"POD_UNKNOWN":   4,
	}
)

func (x PodStatus) Enum() *PodStatus {
	p := new(PodStatus)
	*p = x
	return p
}

func (x PodStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PodStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_scenario_pods_proto_enumTypes[0].Descriptor()
}

func (PodStatus) Type() protoreflect.EnumType {
	return &file_scenario_pods_proto_enumTypes[0]
}

func (x PodStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PodStatus.Descriptor instead.
func (PodStatus) EnumDescriptor() ([]byte, []int) {
	return file_scenario_pods_proto_rawDescGZIP(), []int{0}
}

var File_scenario_pods_proto protoreflect.FileDescriptor

var file_scenario_pods_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2a, 0x61, 0x0a, 0x09, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4f, 0x44, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4f, 0x44, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x4f, 0x44, 0x5f, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4f, 0x44, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4f, 0x44, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x04, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x2d, 0x72,
	0x65, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x63, 0x2d, 0x65,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scenario_pods_proto_rawDescOnce sync.Once
	file_scenario_pods_proto_rawDescData = file_scenario_pods_proto_rawDesc
)

func file_scenario_pods_proto_rawDescGZIP() []byte {
	file_scenario_pods_proto_rawDescOnce.Do(func() {
		file_scenario_pods_proto_rawDescData = protoimpl.X.CompressGZIP(file_scenario_pods_proto_rawDescData)
	})
	return file_scenario_pods_proto_rawDescData
}

var file_scenario_pods_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_scenario_pods_proto_goTypes = []interface{}{
	(PodStatus)(0), // 0: apate.scenario.PodStatus
}
var file_scenario_pods_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scenario_pods_proto_init() }
func file_scenario_pods_proto_init() {
	if File_scenario_pods_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scenario_pods_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scenario_pods_proto_goTypes,
		DependencyIndexes: file_scenario_pods_proto_depIdxs,
		EnumInfos:         file_scenario_pods_proto_enumTypes,
	}.Build()
	File_scenario_pods_proto = out.File
	file_scenario_pods_proto_rawDesc = nil
	file_scenario_pods_proto_goTypes = nil
	file_scenario_pods_proto_depIdxs = nil
}
