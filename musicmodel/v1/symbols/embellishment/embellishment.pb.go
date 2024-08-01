// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: musicmodel/v1/symbols/embellishment/embellishment.proto

package embellishment

import (
	pitch "github.com/tomvodi/limepipes-music-model/musicmodel/v1/pitch"
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

type Type int32

const (
	Type_NoEmbellishment   Type = 0
	Type_SingleGrace       Type = 1
	Type_Doubling          Type = 2
	Type_Strike            Type = 3
	Type_Grip              Type = 4
	Type_Taorluath         Type = 5
	Type_Bubbly            Type = 6
	Type_GraceBirl         Type = 7
	Type_ABirl             Type = 8
	Type_Birl              Type = 9
	Type_ThrowD            Type = 10
	Type_Pele              Type = 11
	Type_DoubleStrike      Type = 12
	Type_TripleStrike      Type = 13
	Type_GTripleStrike     Type = 14
	Type_ThumbTripleStrike Type = 15
	Type_HalfTripleStrike  Type = 16
	Type_DoubleGrace       Type = 17
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:  "NoEmbellishment",
		1:  "SingleGrace",
		2:  "Doubling",
		3:  "Strike",
		4:  "Grip",
		5:  "Taorluath",
		6:  "Bubbly",
		7:  "GraceBirl",
		8:  "ABirl",
		9:  "Birl",
		10: "ThrowD",
		11: "Pele",
		12: "DoubleStrike",
		13: "TripleStrike",
		14: "GTripleStrike",
		15: "ThumbTripleStrike",
		16: "HalfTripleStrike",
		17: "DoubleGrace",
	}
	Type_value = map[string]int32{
		"NoEmbellishment":   0,
		"SingleGrace":       1,
		"Doubling":          2,
		"Strike":            3,
		"Grip":              4,
		"Taorluath":         5,
		"Bubbly":            6,
		"GraceBirl":         7,
		"ABirl":             8,
		"Birl":              9,
		"ThrowD":            10,
		"Pele":              11,
		"DoubleStrike":      12,
		"TripleStrike":      13,
		"GTripleStrike":     14,
		"ThumbTripleStrike": 15,
		"HalfTripleStrike":  16,
		"DoubleGrace":       17,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_musicmodel_v1_symbols_embellishment_embellishment_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_musicmodel_v1_symbols_embellishment_embellishment_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescGZIP(), []int{0}
}

type Variant int32

const (
	Variant_NoVariant Variant = 0
	Variant_G         Variant = 1
	Variant_Half      Variant = 2
	Variant_Thumb     Variant = 3
)

// Enum value maps for Variant.
var (
	Variant_name = map[int32]string{
		0: "NoVariant",
		1: "G",
		2: "Half",
		3: "Thumb",
	}
	Variant_value = map[string]int32{
		"NoVariant": 0,
		"G":         1,
		"Half":      2,
		"Thumb":     3,
	}
)

func (x Variant) Enum() *Variant {
	p := new(Variant)
	*p = x
	return p
}

func (x Variant) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Variant) Descriptor() protoreflect.EnumDescriptor {
	return file_musicmodel_v1_symbols_embellishment_embellishment_proto_enumTypes[1].Descriptor()
}

func (Variant) Type() protoreflect.EnumType {
	return &file_musicmodel_v1_symbols_embellishment_embellishment_proto_enumTypes[1]
}

func (x Variant) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Variant.Descriptor instead.
func (Variant) EnumDescriptor() ([]byte, []int) {
	return file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescGZIP(), []int{1}
}

type Weight int32

const (
	Weight_NoWeight Weight = 0
	Weight_Light    Weight = 1
	Weight_Heavy    Weight = 2
)

// Enum value maps for Weight.
var (
	Weight_name = map[int32]string{
		0: "NoWeight",
		1: "Light",
		2: "Heavy",
	}
	Weight_value = map[string]int32{
		"NoWeight": 0,
		"Light":    1,
		"Heavy":    2,
	}
)

func (x Weight) Enum() *Weight {
	p := new(Weight)
	*p = x
	return p
}

func (x Weight) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Weight) Descriptor() protoreflect.EnumDescriptor {
	return file_musicmodel_v1_symbols_embellishment_embellishment_proto_enumTypes[2].Descriptor()
}

func (Weight) Type() protoreflect.EnumType {
	return &file_musicmodel_v1_symbols_embellishment_embellishment_proto_enumTypes[2]
}

func (x Weight) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Weight.Descriptor instead.
func (Weight) EnumDescriptor() ([]byte, []int) {
	return file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescGZIP(), []int{2}
}

type Embellishment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=musicmodel.v1.symbols.embellishment.Type" json:"type,omitempty"`
	// Pitch is set for embellishments that have a pitch regardless of the melody note
	// following it (e.g. single g-grace)
	// Other embellishments have their pitch defined by the melody note following it
	// (e.g. doubling) because a d-doubling can only precede a d-melody note.
	// In these cases, Pitch is set to NoPitch
	Pitch   pitch.Pitch `protobuf:"varint,2,opt,name=pitch,proto3,enum=musicmodel.v1.pitch.Pitch" json:"pitch,omitempty"`
	Variant Variant     `protobuf:"varint,3,opt,name=variant,proto3,enum=musicmodel.v1.symbols.embellishment.Variant" json:"variant,omitempty"`
	Weight  Weight      `protobuf:"varint,4,opt,name=weight,proto3,enum=musicmodel.v1.symbols.embellishment.Weight" json:"weight,omitempty"`
}

func (x *Embellishment) Reset() {
	*x = Embellishment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicmodel_v1_symbols_embellishment_embellishment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embellishment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embellishment) ProtoMessage() {}

func (x *Embellishment) ProtoReflect() protoreflect.Message {
	mi := &file_musicmodel_v1_symbols_embellishment_embellishment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embellishment.ProtoReflect.Descriptor instead.
func (*Embellishment) Descriptor() ([]byte, []int) {
	return file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescGZIP(), []int{0}
}

func (x *Embellishment) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_NoEmbellishment
}

func (x *Embellishment) GetPitch() pitch.Pitch {
	if x != nil {
		return x.Pitch
	}
	return pitch.Pitch(0)
}

func (x *Embellishment) GetVariant() Variant {
	if x != nil {
		return x.Variant
	}
	return Variant_NoVariant
}

func (x *Embellishment) GetWeight() Weight {
	if x != nil {
		return x.Weight
	}
	return Weight_NoWeight
}

var File_musicmodel_v1_symbols_embellishment_embellishment_proto protoreflect.FileDescriptor

var file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDesc = []byte{
	0x0a, 0x37, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x2e, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69,
	0x74, 0x63, 0x68, 0x2f, 0x70, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8d, 0x02, 0x0a, 0x0d, 0x45, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x30, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x50, 0x69, 0x74, 0x63, 0x68, 0x52, 0x05, 0x70, 0x69, 0x74,
	0x63, 0x68, 0x12, 0x46, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x6d, 0x62, 0x65,
	0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x73, 0x2e, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2a,
	0x9a, 0x02, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x6f, 0x45, 0x6d,
	0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x63, 0x65, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x72, 0x69, 0x70,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x61, 0x6f, 0x72, 0x6c, 0x75, 0x61, 0x74, 0x68, 0x10,
	0x05, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x75, 0x62, 0x62, 0x6c, 0x79, 0x10, 0x06, 0x12, 0x0d, 0x0a,
	0x09, 0x47, 0x72, 0x61, 0x63, 0x65, 0x42, 0x69, 0x72, 0x6c, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x42, 0x69, 0x72, 0x6c, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x69, 0x72, 0x6c, 0x10,
	0x09, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x68, 0x72, 0x6f, 0x77, 0x44, 0x10, 0x0a, 0x12, 0x08, 0x0a,
	0x04, 0x50, 0x65, 0x6c, 0x65, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x53, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x72, 0x69,
	0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x10, 0x0d, 0x12, 0x11, 0x0a, 0x0d, 0x47,
	0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x10, 0x0e, 0x12, 0x15,
	0x0a, 0x11, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6b, 0x65, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x61, 0x6c, 0x66, 0x54, 0x72, 0x69,
	0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x63, 0x65, 0x10, 0x11, 0x2a, 0x34, 0x0a, 0x07,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x6f, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x61, 0x6c, 0x66, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x10, 0x03, 0x2a, 0x2c, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0c, 0x0a, 0x08,
	0x4e, 0x6f, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x69,
	0x67, 0x68, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x65, 0x61, 0x76, 0x79, 0x10, 0x02,
	0x42, 0xbb, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x65,
	0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x12, 0x45, 0x6d,
	0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x6f, 0x6d, 0x76, 0x6f, 0x64, 0x69, 0x2f, 0x6c, 0x69, 0x6d, 0x65, 0x70, 0x69, 0x70, 0x65, 0x73,
	0x2d, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x73, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0xa2, 0x02, 0x04, 0x4d, 0x56, 0x53, 0x45, 0xaa, 0x02, 0x23, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e,
	0x45, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x23,
	0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x5c, 0x45, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0xe2, 0x02, 0x2f, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5c, 0x56, 0x31, 0x5c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x5c, 0x45, 0x6d, 0x62, 0x65,
	0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x26, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x3a,
	0x3a, 0x45, 0x6d, 0x62, 0x65, 0x6c, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescOnce sync.Once
	file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescData = file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDesc
)

func file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescGZIP() []byte {
	file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescOnce.Do(func() {
		file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescData = protoimpl.X.CompressGZIP(file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescData)
	})
	return file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDescData
}

var file_musicmodel_v1_symbols_embellishment_embellishment_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_musicmodel_v1_symbols_embellishment_embellishment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_musicmodel_v1_symbols_embellishment_embellishment_proto_goTypes = []any{
	(Type)(0),             // 0: musicmodel.v1.symbols.embellishment.Type
	(Variant)(0),          // 1: musicmodel.v1.symbols.embellishment.Variant
	(Weight)(0),           // 2: musicmodel.v1.symbols.embellishment.Weight
	(*Embellishment)(nil), // 3: musicmodel.v1.symbols.embellishment.Embellishment
	(pitch.Pitch)(0),      // 4: musicmodel.v1.pitch.Pitch
}
var file_musicmodel_v1_symbols_embellishment_embellishment_proto_depIdxs = []int32{
	0, // 0: musicmodel.v1.symbols.embellishment.Embellishment.type:type_name -> musicmodel.v1.symbols.embellishment.Type
	4, // 1: musicmodel.v1.symbols.embellishment.Embellishment.pitch:type_name -> musicmodel.v1.pitch.Pitch
	1, // 2: musicmodel.v1.symbols.embellishment.Embellishment.variant:type_name -> musicmodel.v1.symbols.embellishment.Variant
	2, // 3: musicmodel.v1.symbols.embellishment.Embellishment.weight:type_name -> musicmodel.v1.symbols.embellishment.Weight
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_musicmodel_v1_symbols_embellishment_embellishment_proto_init() }
func file_musicmodel_v1_symbols_embellishment_embellishment_proto_init() {
	if File_musicmodel_v1_symbols_embellishment_embellishment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_musicmodel_v1_symbols_embellishment_embellishment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Embellishment); i {
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
			RawDescriptor: file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_musicmodel_v1_symbols_embellishment_embellishment_proto_goTypes,
		DependencyIndexes: file_musicmodel_v1_symbols_embellishment_embellishment_proto_depIdxs,
		EnumInfos:         file_musicmodel_v1_symbols_embellishment_embellishment_proto_enumTypes,
		MessageInfos:      file_musicmodel_v1_symbols_embellishment_embellishment_proto_msgTypes,
	}.Build()
	File_musicmodel_v1_symbols_embellishment_embellishment_proto = out.File
	file_musicmodel_v1_symbols_embellishment_embellishment_proto_rawDesc = nil
	file_musicmodel_v1_symbols_embellishment_embellishment_proto_goTypes = nil
	file_musicmodel_v1_symbols_embellishment_embellishment_proto_depIdxs = nil
}