// Copyright 2013 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for managed user shared settings.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: managed_user_shared_setting_specifics.proto

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

// Properties of managed user shared setting sync objects.
type ManagedUserSharedSettingSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The MU ID for the managed user to whom the setting applies.
	MuId *string `protobuf:"bytes,1,opt,name=mu_id,json=muId" json:"mu_id,omitempty"`
	// The key of the setting.
	Key *string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	// The setting value. The setting is a JSON encoding of an arbitrary
	// Javascript value.
	Value *string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	// This flag is set by the server to acknowledge that it has committed a
	// change to a setting.
	Acknowledged *bool `protobuf:"varint,4,opt,name=acknowledged,def=0" json:"acknowledged,omitempty"`
}

// Default values for ManagedUserSharedSettingSpecifics fields.
const (
	Default_ManagedUserSharedSettingSpecifics_Acknowledged = bool(false)
)

func (x *ManagedUserSharedSettingSpecifics) Reset() {
	*x = ManagedUserSharedSettingSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managed_user_shared_setting_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagedUserSharedSettingSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagedUserSharedSettingSpecifics) ProtoMessage() {}

func (x *ManagedUserSharedSettingSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_managed_user_shared_setting_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagedUserSharedSettingSpecifics.ProtoReflect.Descriptor instead.
func (*ManagedUserSharedSettingSpecifics) Descriptor() ([]byte, []int) {
	return file_managed_user_shared_setting_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *ManagedUserSharedSettingSpecifics) GetMuId() string {
	if x != nil && x.MuId != nil {
		return *x.MuId
	}
	return ""
}

func (x *ManagedUserSharedSettingSpecifics) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *ManagedUserSharedSettingSpecifics) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

func (x *ManagedUserSharedSettingSpecifics) GetAcknowledged() bool {
	if x != nil && x.Acknowledged != nil {
		return *x.Acknowledged
	}
	return Default_ManagedUserSharedSettingSpecifics_Acknowledged
}

var File_managed_user_shared_setting_specifics_proto protoreflect.FileDescriptor

var file_managed_user_shared_setting_specifics_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x8b, 0x01, 0x0a, 0x21, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x13, 0x0a, 0x05,
	0x6d, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x75, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x0c, 0x61, 0x63, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x3a,
	0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0c, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x64, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50,
	0x01,
}

var (
	file_managed_user_shared_setting_specifics_proto_rawDescOnce sync.Once
	file_managed_user_shared_setting_specifics_proto_rawDescData = file_managed_user_shared_setting_specifics_proto_rawDesc
)

func file_managed_user_shared_setting_specifics_proto_rawDescGZIP() []byte {
	file_managed_user_shared_setting_specifics_proto_rawDescOnce.Do(func() {
		file_managed_user_shared_setting_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_managed_user_shared_setting_specifics_proto_rawDescData)
	})
	return file_managed_user_shared_setting_specifics_proto_rawDescData
}

var file_managed_user_shared_setting_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_managed_user_shared_setting_specifics_proto_goTypes = []interface{}{
	(*ManagedUserSharedSettingSpecifics)(nil), // 0: sync_pb.ManagedUserSharedSettingSpecifics
}
var file_managed_user_shared_setting_specifics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managed_user_shared_setting_specifics_proto_init() }
func file_managed_user_shared_setting_specifics_proto_init() {
	if File_managed_user_shared_setting_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managed_user_shared_setting_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagedUserSharedSettingSpecifics); i {
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
			RawDescriptor: file_managed_user_shared_setting_specifics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managed_user_shared_setting_specifics_proto_goTypes,
		DependencyIndexes: file_managed_user_shared_setting_specifics_proto_depIdxs,
		MessageInfos:      file_managed_user_shared_setting_specifics_proto_msgTypes,
	}.Build()
	File_managed_user_shared_setting_specifics_proto = out.File
	file_managed_user_shared_setting_specifics_proto_rawDesc = nil
	file_managed_user_shared_setting_specifics_proto_goTypes = nil
	file_managed_user_shared_setting_specifics_proto_depIdxs = nil
}
