// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: musicmodel/v1/boundary/boundary.proto

package boundary

import (
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

type Boundary int32

const (
	Boundary_NoBoundary Boundary = 0
	Boundary_Start      Boundary = 1
	Boundary_End        Boundary = 2
)

// Enum value maps for Boundary.
var (
	Boundary_name = map[int32]string{
		0: "NoBoundary",
		1: "Start",
		2: "End",
	}
	Boundary_value = map[string]int32{
		"NoBoundary": 0,
		"Start":      1,
		"End":        2,
	}
)

func (x Boundary) Enum() *Boundary {
	p := new(Boundary)
	*p = x
	return p
}

func (x Boundary) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Boundary) Descriptor() protoreflect.EnumDescriptor {
	return file_musicmodel_v1_boundary_boundary_proto_enumTypes[0].Descriptor()
}

func (Boundary) Type() protoreflect.EnumType {
	return &file_musicmodel_v1_boundary_boundary_proto_enumTypes[0]
}

func (x Boundary) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Boundary.Descriptor instead.
func (Boundary) EnumDescriptor() ([]byte, []int) {
	return file_musicmodel_v1_boundary_boundary_proto_rawDescGZIP(), []int{0}
}

var File_musicmodel_v1_boundary_boundary_proto protoreflect.FileDescriptor

var file_musicmodel_v1_boundary_boundary_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2a,
	0x2e, 0x0a, 0x08, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x0a, 0x4e,
	0x6f, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x10, 0x02, 0x42,
	0xe6, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x42, 0x0d,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6d, 0x76,
	0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70, 0x69, 0x70, 0x65, 0x73, 0x2d, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0xa2, 0x02, 0x03, 0x4d, 0x56, 0x42, 0xaa, 0x02, 0x16, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x56, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0xca,
	0x02, 0x16, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0xe2, 0x02, 0x22, 0x4d, 0x75, 0x73, 0x69, 0x63,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18,
	0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_musicmodel_v1_boundary_boundary_proto_rawDescOnce sync.Once
	file_musicmodel_v1_boundary_boundary_proto_rawDescData = file_musicmodel_v1_boundary_boundary_proto_rawDesc
)

func file_musicmodel_v1_boundary_boundary_proto_rawDescGZIP() []byte {
	file_musicmodel_v1_boundary_boundary_proto_rawDescOnce.Do(func() {
		file_musicmodel_v1_boundary_boundary_proto_rawDescData = protoimpl.X.CompressGZIP(file_musicmodel_v1_boundary_boundary_proto_rawDescData)
	})
	return file_musicmodel_v1_boundary_boundary_proto_rawDescData
}

var file_musicmodel_v1_boundary_boundary_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_musicmodel_v1_boundary_boundary_proto_goTypes = []any{
	(Boundary)(0), // 0: musicmodel.v1.boundary.Boundary
}
var file_musicmodel_v1_boundary_boundary_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_musicmodel_v1_boundary_boundary_proto_init() }
func file_musicmodel_v1_boundary_boundary_proto_init() {
	if File_musicmodel_v1_boundary_boundary_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_musicmodel_v1_boundary_boundary_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_musicmodel_v1_boundary_boundary_proto_goTypes,
		DependencyIndexes: file_musicmodel_v1_boundary_boundary_proto_depIdxs,
		EnumInfos:         file_musicmodel_v1_boundary_boundary_proto_enumTypes,
	}.Build()
	File_musicmodel_v1_boundary_boundary_proto = out.File
	file_musicmodel_v1_boundary_boundary_proto_rawDesc = nil
	file_musicmodel_v1_boundary_boundary_proto_goTypes = nil
	file_musicmodel_v1_boundary_boundary_proto_depIdxs = nil
}
