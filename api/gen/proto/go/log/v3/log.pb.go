// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: log/v3/log.proto

package logv3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Level int32

const (
	Level_LEVEL_UNKNOWN Level = 0
	Level_LEVEL_TRACE   Level = 1
	Level_LEVEL_DEBUG   Level = 2
	Level_LEVEL_INFO    Level = 3
	Level_LEVEL_WARN    Level = 4
	Level_LEVEL_ERROR   Level = 5
	Level_LEVEL_FATAL   Level = 6
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "LEVEL_UNKNOWN",
		1: "LEVEL_TRACE",
		2: "LEVEL_DEBUG",
		3: "LEVEL_INFO",
		4: "LEVEL_WARN",
		5: "LEVEL_ERROR",
		6: "LEVEL_FATAL",
	}
	Level_value = map[string]int32{
		"LEVEL_UNKNOWN": 0,
		"LEVEL_TRACE":   1,
		"LEVEL_DEBUG":   2,
		"LEVEL_INFO":    3,
		"LEVEL_WARN":    4,
		"LEVEL_ERROR":   5,
		"LEVEL_FATAL":   6,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_log_v3_log_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_log_v3_log_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_log_v3_log_proto_rawDescGZIP(), []int{0}
}

type Format int32

const (
	Format_FORMAT_UNKNOWN   Format = 0
	Format_FORMAT_JSON      Format = 1
	Format_FORMAT_SYSLOG    Format = 2
	Format_FORMAT_KEY_VALUE Format = 3
)

// Enum value maps for Format.
var (
	Format_name = map[int32]string{
		0: "FORMAT_UNKNOWN",
		1: "FORMAT_JSON",
		2: "FORMAT_SYSLOG",
		3: "FORMAT_KEY_VALUE",
	}
	Format_value = map[string]int32{
		"FORMAT_UNKNOWN":   0,
		"FORMAT_JSON":      1,
		"FORMAT_SYSLOG":    2,
		"FORMAT_KEY_VALUE": 3,
	}
)

func (x Format) Enum() *Format {
	p := new(Format)
	*p = x
	return p
}

func (x Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Format) Descriptor() protoreflect.EnumDescriptor {
	return file_log_v3_log_proto_enumTypes[1].Descriptor()
}

func (Format) Type() protoreflect.EnumType {
	return &file_log_v3_log_proto_enumTypes[1]
}

func (x Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Format.Descriptor instead.
func (Format) EnumDescriptor() ([]byte, []int) {
	return file_log_v3_log_proto_rawDescGZIP(), []int{1}
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Metadata  map[string]string      `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Level     Level                  `protobuf:"varint,3,opt,name=level,proto3,enum=log.v3.Level" json:"level,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Format    Format                 `protobuf:"varint,5,opt,name=format,proto3,enum=log.v3.Format" json:"format,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_v3_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_log_v3_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_log_v3_log_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Log) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Log) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_LEVEL_UNKNOWN
}

func (x *Log) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Log) GetFormat() Format {
	if x != nil {
		return x.Format
	}
	return Format_FORMAT_UNKNOWN
}

var File_log_v3_log_proto protoreflect.FileDescriptor

var file_log_v3_log_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x03,
	0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x7e, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x52,
	0x41, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44,
	0x45, 0x42, 0x55, 0x47, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x57, 0x41, 0x52, 0x4e, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x06, 0x2a, 0x56, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54,
	0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x5f, 0x53, 0x59, 0x53, 0x4c, 0x4f, 0x47, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x54, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x03,
	0x42, 0x88, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x42,
	0x08, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x33, 0x31, 0x33, 0x33, 0x33, 0x33, 0x33, 0x37,
	0x2f, 0x62, 0x6d, 0x72, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x33, 0x3b, 0x6c,
	0x6f, 0x67, 0x76, 0x33, 0xa2, 0x02, 0x03, 0x4c, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x4c, 0x6f, 0x67,
	0x2e, 0x56, 0x33, 0xca, 0x02, 0x06, 0x4c, 0x6f, 0x67, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x12, 0x4c,
	0x6f, 0x67, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x07, 0x4c, 0x6f, 0x67, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_log_v3_log_proto_rawDescOnce sync.Once
	file_log_v3_log_proto_rawDescData = file_log_v3_log_proto_rawDesc
)

func file_log_v3_log_proto_rawDescGZIP() []byte {
	file_log_v3_log_proto_rawDescOnce.Do(func() {
		file_log_v3_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_v3_log_proto_rawDescData)
	})
	return file_log_v3_log_proto_rawDescData
}

var file_log_v3_log_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_log_v3_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_log_v3_log_proto_goTypes = []interface{}{
	(Level)(0),                    // 0: log.v3.Level
	(Format)(0),                   // 1: log.v3.Format
	(*Log)(nil),                   // 2: log.v3.Log
	nil,                           // 3: log.v3.Log.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_log_v3_log_proto_depIdxs = []int32{
	3, // 0: log.v3.Log.metadata:type_name -> log.v3.Log.MetadataEntry
	0, // 1: log.v3.Log.level:type_name -> log.v3.Level
	4, // 2: log.v3.Log.timestamp:type_name -> google.protobuf.Timestamp
	1, // 3: log.v3.Log.format:type_name -> log.v3.Format
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_log_v3_log_proto_init() }
func file_log_v3_log_proto_init() {
	if File_log_v3_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_v3_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_log_v3_log_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_log_v3_log_proto_goTypes,
		DependencyIndexes: file_log_v3_log_proto_depIdxs,
		EnumInfos:         file_log_v3_log_proto_enumTypes,
		MessageInfos:      file_log_v3_log_proto_msgTypes,
	}.Build()
	File_log_v3_log_proto = out.File
	file_log_v3_log_proto_rawDesc = nil
	file_log_v3_log_proto_goTypes = nil
	file_log_v3_log_proto_depIdxs = nil
}
