// Copyright 2012 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.1
// source: get_updates_caller_info.proto

package sync_pb

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

// This enum was deprecated in M28.  The preferred represenation of this
// information is now the GetUpdatesOrigin enum, which is defined in
// sync_enums.proto.
// TODO(crbug.com/510165): Remove all values except for UNKNOWN and stop
// filling the field once the server doesn't depend on it anymore.
type GetUpdatesCallerInfo_GetUpdatesSource int32

const (
	GetUpdatesCallerInfo_UNKNOWN      GetUpdatesCallerInfo_GetUpdatesSource = 0 // The source was not set by the caller.
	GetUpdatesCallerInfo_FIRST_UPDATE GetUpdatesCallerInfo_GetUpdatesSource = 1 // First request after browser restart.  Not to
	// be confused with "NEW_CLIENT".
	GetUpdatesCallerInfo_LOCAL                    GetUpdatesCallerInfo_GetUpdatesSource = 2 // The source of the update was a local change.
	GetUpdatesCallerInfo_NOTIFICATION             GetUpdatesCallerInfo_GetUpdatesSource = 3 // The source of the update was a p2p notification.
	GetUpdatesCallerInfo_PERIODIC                 GetUpdatesCallerInfo_GetUpdatesSource = 4 // The source of the update was periodic polling.
	GetUpdatesCallerInfo_SYNC_CYCLE_CONTINUATION  GetUpdatesCallerInfo_GetUpdatesSource = 5 // The source of the update was a
	GetUpdatesCallerInfo_NEWLY_SUPPORTED_DATATYPE GetUpdatesCallerInfo_GetUpdatesSource = 7 // The client is in configuration mode
	// because it's syncing all datatypes, and
	// support for a new datatype was recently
	// released via a software auto-update.
	GetUpdatesCallerInfo_MIGRATION GetUpdatesCallerInfo_GetUpdatesSource = 8 // The client is in configuration mode because a
	// MIGRATION_DONE error previously returned by the
	// server necessitated resynchronization.
	GetUpdatesCallerInfo_NEW_CLIENT GetUpdatesCallerInfo_GetUpdatesSource = 9 // The client is in configuration mode because the
	// user enabled sync for the first time.  Not to be
	// confused with FIRST_UPDATE.
	GetUpdatesCallerInfo_RECONFIGURATION GetUpdatesCallerInfo_GetUpdatesSource = 10 // The client is in configuration mode because the
	// user opted to sync a different set of datatypes.
	GetUpdatesCallerInfo_DATATYPE_REFRESH GetUpdatesCallerInfo_GetUpdatesSource = 11 // A datatype has requested a refresh. This is
	// typically used when datatype's have custom
	// sync UI, e.g. sessions.
	GetUpdatesCallerInfo_RETRY        GetUpdatesCallerInfo_GetUpdatesSource = 13 // A follow-up GU to pick up updates missed by previous GU.
	GetUpdatesCallerInfo_PROGRAMMATIC GetUpdatesCallerInfo_GetUpdatesSource = 14 // The client is programmatically enabling/disabling
)

// Enum value maps for GetUpdatesCallerInfo_GetUpdatesSource.
var (
	GetUpdatesCallerInfo_GetUpdatesSource_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "FIRST_UPDATE",
		2:  "LOCAL",
		3:  "NOTIFICATION",
		4:  "PERIODIC",
		5:  "SYNC_CYCLE_CONTINUATION",
		7:  "NEWLY_SUPPORTED_DATATYPE",
		8:  "MIGRATION",
		9:  "NEW_CLIENT",
		10: "RECONFIGURATION",
		11: "DATATYPE_REFRESH",
		13: "RETRY",
		14: "PROGRAMMATIC",
	}
	GetUpdatesCallerInfo_GetUpdatesSource_value = map[string]int32{
		"UNKNOWN":                  0,
		"FIRST_UPDATE":             1,
		"LOCAL":                    2,
		"NOTIFICATION":             3,
		"PERIODIC":                 4,
		"SYNC_CYCLE_CONTINUATION":  5,
		"NEWLY_SUPPORTED_DATATYPE": 7,
		"MIGRATION":                8,
		"NEW_CLIENT":               9,
		"RECONFIGURATION":          10,
		"DATATYPE_REFRESH":         11,
		"RETRY":                    13,
		"PROGRAMMATIC":             14,
	}
)

func (x GetUpdatesCallerInfo_GetUpdatesSource) Enum() *GetUpdatesCallerInfo_GetUpdatesSource {
	p := new(GetUpdatesCallerInfo_GetUpdatesSource)
	*p = x
	return p
}

func (x GetUpdatesCallerInfo_GetUpdatesSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetUpdatesCallerInfo_GetUpdatesSource) Descriptor() protoreflect.EnumDescriptor {
	return file_get_updates_caller_info_proto_enumTypes[0].Descriptor()
}

func (GetUpdatesCallerInfo_GetUpdatesSource) Type() protoreflect.EnumType {
	return &file_get_updates_caller_info_proto_enumTypes[0]
}

func (x GetUpdatesCallerInfo_GetUpdatesSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GetUpdatesCallerInfo_GetUpdatesSource) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GetUpdatesCallerInfo_GetUpdatesSource(num)
	return nil
}

// Deprecated: Use GetUpdatesCallerInfo_GetUpdatesSource.Descriptor instead.
func (GetUpdatesCallerInfo_GetUpdatesSource) EnumDescriptor() ([]byte, []int) {
	return file_get_updates_caller_info_proto_rawDescGZIP(), []int{0, 0}
}

type GetUpdatesCallerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in get_updates_caller_info.proto.
	Source *GetUpdatesCallerInfo_GetUpdatesSource `protobuf:"varint,1,req,name=source,enum=sync_pb.GetUpdatesCallerInfo_GetUpdatesSource" json:"source,omitempty"`
	// True only if notifications were enabled for this GetUpdateMessage.
	// TODO(crbug.com/510165): Move this bool out of GetUpdatesCallerInfo so that
	// GetUpdatesCallerInfo can eventually be removed.
	NotificationsEnabled *bool `protobuf:"varint,2,opt,name=notifications_enabled,json=notificationsEnabled" json:"notifications_enabled,omitempty"`
}

func (x *GetUpdatesCallerInfo) Reset() {
	*x = GetUpdatesCallerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_updates_caller_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpdatesCallerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdatesCallerInfo) ProtoMessage() {}

func (x *GetUpdatesCallerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_get_updates_caller_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdatesCallerInfo.ProtoReflect.Descriptor instead.
func (*GetUpdatesCallerInfo) Descriptor() ([]byte, []int) {
	return file_get_updates_caller_info_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in get_updates_caller_info.proto.
func (x *GetUpdatesCallerInfo) GetSource() GetUpdatesCallerInfo_GetUpdatesSource {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return GetUpdatesCallerInfo_UNKNOWN
}

func (x *GetUpdatesCallerInfo) GetNotificationsEnabled() bool {
	if x != nil && x.NotificationsEnabled != nil {
		return *x.NotificationsEnabled
	}
	return false
}

var File_get_updates_caller_info_proto protoreflect.FileDescriptor

var file_get_updates_caller_info_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x98, 0x03, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0e, 0x32, 0x2e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x15, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x22, 0xfe, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10,
	0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x49, 0x43, 0x10,
	0x04, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x1c,
	0x0a, 0x18, 0x4e, 0x45, 0x57, 0x4c, 0x59, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45,
	0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09,
	0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x4e,
	0x45, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x52,
	0x45, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a,
	0x12, 0x14, 0x0a, 0x10, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x46,
	0x52, 0x45, 0x53, 0x48, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x54, 0x52, 0x59, 0x10,
	0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x4d, 0x41, 0x54, 0x49,
	0x43, 0x10, 0x0e, 0x42, 0x36, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
	0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
}

var (
	file_get_updates_caller_info_proto_rawDescOnce sync.Once
	file_get_updates_caller_info_proto_rawDescData = file_get_updates_caller_info_proto_rawDesc
)

func file_get_updates_caller_info_proto_rawDescGZIP() []byte {
	file_get_updates_caller_info_proto_rawDescOnce.Do(func() {
		file_get_updates_caller_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_updates_caller_info_proto_rawDescData)
	})
	return file_get_updates_caller_info_proto_rawDescData
}

var file_get_updates_caller_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_get_updates_caller_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_get_updates_caller_info_proto_goTypes = []interface{}{
	(GetUpdatesCallerInfo_GetUpdatesSource)(0), // 0: sync_pb.GetUpdatesCallerInfo.GetUpdatesSource
	(*GetUpdatesCallerInfo)(nil),               // 1: sync_pb.GetUpdatesCallerInfo
}
var file_get_updates_caller_info_proto_depIdxs = []int32{
	0, // 0: sync_pb.GetUpdatesCallerInfo.source:type_name -> sync_pb.GetUpdatesCallerInfo.GetUpdatesSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_get_updates_caller_info_proto_init() }
func file_get_updates_caller_info_proto_init() {
	if File_get_updates_caller_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_get_updates_caller_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpdatesCallerInfo); i {
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
			RawDescriptor: file_get_updates_caller_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_updates_caller_info_proto_goTypes,
		DependencyIndexes: file_get_updates_caller_info_proto_depIdxs,
		EnumInfos:         file_get_updates_caller_info_proto_enumTypes,
		MessageInfos:      file_get_updates_caller_info_proto_msgTypes,
	}.Build()
	File_get_updates_caller_info_proto = out.File
	file_get_updates_caller_info_proto_rawDesc = nil
	file_get_updates_caller_info_proto_goTypes = nil
	file_get_updates_caller_info_proto_depIdxs = nil
}
