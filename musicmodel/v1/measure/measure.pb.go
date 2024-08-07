// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: musicmodel/v1/measure/measure.proto

package measure

import (
	barline "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/barline"
	symbols "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
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

type Measure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeftBarline    *barline.Barline  `protobuf:"bytes,1,opt,name=left_barline,json=leftBarline,proto3" json:"left_barline,omitempty"`
	RightBarline   *barline.Barline  `protobuf:"bytes,2,opt,name=right_barline,json=rightBarline,proto3" json:"right_barline,omitempty"`
	Time           *TimeSignature    `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Symbols        []*symbols.Symbol `protobuf:"bytes,4,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Comments       []string          `protobuf:"bytes,5,rep,name=comments,proto3" json:"comments,omitempty"`
	InlineText     []string          `protobuf:"bytes,6,rep,name=inline_text,json=inlineText,proto3" json:"inline_text,omitempty"`
	ImportMessages []*ImportMessage  `protobuf:"bytes,7,rep,name=import_messages,json=importMessages,proto3" json:"import_messages,omitempty"`
}

func (x *Measure) Reset() {
	*x = Measure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicmodel_v1_measure_measure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measure) ProtoMessage() {}

func (x *Measure) ProtoReflect() protoreflect.Message {
	mi := &file_musicmodel_v1_measure_measure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measure.ProtoReflect.Descriptor instead.
func (*Measure) Descriptor() ([]byte, []int) {
	return file_musicmodel_v1_measure_measure_proto_rawDescGZIP(), []int{0}
}

func (x *Measure) GetLeftBarline() *barline.Barline {
	if x != nil {
		return x.LeftBarline
	}
	return nil
}

func (x *Measure) GetRightBarline() *barline.Barline {
	if x != nil {
		return x.RightBarline
	}
	return nil
}

func (x *Measure) GetTime() *TimeSignature {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Measure) GetSymbols() []*symbols.Symbol {
	if x != nil {
		return x.Symbols
	}
	return nil
}

func (x *Measure) GetComments() []string {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *Measure) GetInlineText() []string {
	if x != nil {
		return x.InlineText
	}
	return nil
}

func (x *Measure) GetImportMessages() []*ImportMessage {
	if x != nil {
		return x.ImportMessages
	}
	return nil
}

var File_musicmodel_v1_measure_measure_proto protoreflect.FileDescriptor

var file_musicmodel_v1_measure_measure_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x1a, 0x23, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x72, 0x6c,
	0x69, 0x6e, 0x65, 0x2f, 0x62, 0x61, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x03,
	0x0a, 0x07, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x6c, 0x65, 0x66,
	0x74, 0x5f, 0x62, 0x61, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x61, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x61, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x0b, 0x6c, 0x65, 0x66, 0x74, 0x42, 0x61, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x43, 0x0a, 0x0d,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x61, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x62, 0x61, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x61, 0x72, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x0c, 0x72, 0x69, 0x67, 0x68, 0x74, 0x42, 0x61, 0x72, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x07, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x4d, 0x0a, 0x0f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x0e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x42, 0xde, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x42, 0x0c,
	0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6d, 0x76, 0x6f,
	0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70, 0x69, 0x70, 0x65, 0x73, 0x2d, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0xa2, 0x02, 0x03,
	0x4d, 0x56, 0x4d, 0xaa, 0x02, 0x15, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x56, 0x31, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0xca, 0x02, 0x15, 0x4d, 0x75,
	0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0xe2, 0x02, 0x21, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5c, 0x56, 0x31, 0x5c, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_musicmodel_v1_measure_measure_proto_rawDescOnce sync.Once
	file_musicmodel_v1_measure_measure_proto_rawDescData = file_musicmodel_v1_measure_measure_proto_rawDesc
)

func file_musicmodel_v1_measure_measure_proto_rawDescGZIP() []byte {
	file_musicmodel_v1_measure_measure_proto_rawDescOnce.Do(func() {
		file_musicmodel_v1_measure_measure_proto_rawDescData = protoimpl.X.CompressGZIP(file_musicmodel_v1_measure_measure_proto_rawDescData)
	})
	return file_musicmodel_v1_measure_measure_proto_rawDescData
}

var file_musicmodel_v1_measure_measure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_musicmodel_v1_measure_measure_proto_goTypes = []any{
	(*Measure)(nil),         // 0: musicmodel.v1.measure.Measure
	(*barline.Barline)(nil), // 1: musicmodel.v1.barline.Barline
	(*TimeSignature)(nil),   // 2: musicmodel.v1.measure.TimeSignature
	(*symbols.Symbol)(nil),  // 3: musicmodel.v1.symbols.Symbol
	(*ImportMessage)(nil),   // 4: musicmodel.v1.measure.ImportMessage
}
var file_musicmodel_v1_measure_measure_proto_depIdxs = []int32{
	1, // 0: musicmodel.v1.measure.Measure.left_barline:type_name -> musicmodel.v1.barline.Barline
	1, // 1: musicmodel.v1.measure.Measure.right_barline:type_name -> musicmodel.v1.barline.Barline
	2, // 2: musicmodel.v1.measure.Measure.time:type_name -> musicmodel.v1.measure.TimeSignature
	3, // 3: musicmodel.v1.measure.Measure.symbols:type_name -> musicmodel.v1.symbols.Symbol
	4, // 4: musicmodel.v1.measure.Measure.import_messages:type_name -> musicmodel.v1.measure.ImportMessage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_musicmodel_v1_measure_measure_proto_init() }
func file_musicmodel_v1_measure_measure_proto_init() {
	if File_musicmodel_v1_measure_measure_proto != nil {
		return
	}
	file_musicmodel_v1_measure_time_signature_proto_init()
	file_musicmodel_v1_measure_import_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_musicmodel_v1_measure_measure_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Measure); i {
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
			RawDescriptor: file_musicmodel_v1_measure_measure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_musicmodel_v1_measure_measure_proto_goTypes,
		DependencyIndexes: file_musicmodel_v1_measure_measure_proto_depIdxs,
		MessageInfos:      file_musicmodel_v1_measure_measure_proto_msgTypes,
	}.Build()
	File_musicmodel_v1_measure_measure_proto = out.File
	file_musicmodel_v1_measure_measure_proto_rawDesc = nil
	file_musicmodel_v1_measure_measure_proto_goTypes = nil
	file_musicmodel_v1_measure_measure_proto_depIdxs = nil
}
