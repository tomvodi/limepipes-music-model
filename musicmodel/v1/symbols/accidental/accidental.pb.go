// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: musicmodel/v1/symbols/accidental/accidental.proto

package accidental

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

type Accidental int32

const (
	Accidental_NoAccidental Accidental = 0
	Accidental_Sharp        Accidental = 1
	Accidental_Flat         Accidental = 2
	Accidental_Natural      Accidental = 3
)

// Enum value maps for Accidental.
var (
	Accidental_name = map[int32]string{
		0: "NoAccidental",
		1: "Sharp",
		2: "Flat",
		3: "Natural",
	}
	Accidental_value = map[string]int32{
		"NoAccidental": 0,
		"Sharp":        1,
		"Flat":         2,
		"Natural":      3,
	}
)

func (x Accidental) Enum() *Accidental {
	p := new(Accidental)
	*p = x
	return p
}

func (x Accidental) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Accidental) Descriptor() protoreflect.EnumDescriptor {
	return file_musicmodel_v1_symbols_accidental_accidental_proto_enumTypes[0].Descriptor()
}

func (Accidental) Type() protoreflect.EnumType {
	return &file_musicmodel_v1_symbols_accidental_accidental_proto_enumTypes[0]
}

func (x Accidental) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Accidental.Descriptor instead.
func (Accidental) EnumDescriptor() ([]byte, []int) {
	return file_musicmodel_v1_symbols_accidental_accidental_proto_rawDescGZIP(), []int{0}
}

var File_musicmodel_v1_symbols_accidental_accidental_proto protoreflect.FileDescriptor

var file_musicmodel_v1_symbols_accidental_accidental_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2a, 0x40, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x6f, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x72, 0x70, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x74, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x61,
	0x74, 0x75, 0x72, 0x61, 0x6c, 0x10, 0x03, 0x42, 0xa5, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x42, 0x0f, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6f, 0x6d, 0x76, 0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70, 0x69, 0x70, 0x65,
	0x73, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0xa2, 0x02, 0x04,
	0x4d, 0x56, 0x53, 0x41, 0xaa, 0x02, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x41, 0x63, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0xca, 0x02, 0x20, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x5c,
	0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0xe2, 0x02, 0x2c, 0x4d, 0x75, 0x73,
	0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x73, 0x5c, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23, 0x4d, 0x75, 0x73, 0x69,
	0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x73, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_musicmodel_v1_symbols_accidental_accidental_proto_rawDescOnce sync.Once
	file_musicmodel_v1_symbols_accidental_accidental_proto_rawDescData = file_musicmodel_v1_symbols_accidental_accidental_proto_rawDesc
)

func file_musicmodel_v1_symbols_accidental_accidental_proto_rawDescGZIP() []byte {
	file_musicmodel_v1_symbols_accidental_accidental_proto_rawDescOnce.Do(func() {
		file_musicmodel_v1_symbols_accidental_accidental_proto_rawDescData = protoimpl.X.CompressGZIP(file_musicmodel_v1_symbols_accidental_accidental_proto_rawDescData)
	})
	return file_musicmodel_v1_symbols_accidental_accidental_proto_rawDescData
}

var file_musicmodel_v1_symbols_accidental_accidental_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_musicmodel_v1_symbols_accidental_accidental_proto_goTypes = []any{
	(Accidental)(0), // 0: musicmodel.v1.symbols.accidental.Accidental
}
var file_musicmodel_v1_symbols_accidental_accidental_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_musicmodel_v1_symbols_accidental_accidental_proto_init() }
func file_musicmodel_v1_symbols_accidental_accidental_proto_init() {
	if File_musicmodel_v1_symbols_accidental_accidental_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_musicmodel_v1_symbols_accidental_accidental_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_musicmodel_v1_symbols_accidental_accidental_proto_goTypes,
		DependencyIndexes: file_musicmodel_v1_symbols_accidental_accidental_proto_depIdxs,
		EnumInfos:         file_musicmodel_v1_symbols_accidental_accidental_proto_enumTypes,
	}.Build()
	File_musicmodel_v1_symbols_accidental_accidental_proto = out.File
	file_musicmodel_v1_symbols_accidental_accidental_proto_rawDesc = nil
	file_musicmodel_v1_symbols_accidental_accidental_proto_goTypes = nil
	file_musicmodel_v1_symbols_accidental_accidental_proto_depIdxs = nil
}
