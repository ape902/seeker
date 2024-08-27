// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: usercenter.proto

package user_center_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserCenterUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NickName string `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Rule     int32  `protobuf:"varint,4,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *UserCenterUserResp) Reset() {
	*x = UserCenterUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCenterUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCenterUserResp) ProtoMessage() {}

func (x *UserCenterUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCenterUserResp.ProtoReflect.Descriptor instead.
func (*UserCenterUserResp) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{0}
}

func (x *UserCenterUserResp) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserCenterUserResp) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserCenterUserResp) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserCenterUserResp) GetRule() int32 {
	if x != nil {
		return x.Rule
	}
	return 0
}

type UserCenterUserAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*UserCenterUserResp `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UserCenterUserAll) Reset() {
	*x = UserCenterUserAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCenterUserAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCenterUserAll) ProtoMessage() {}

func (x *UserCenterUserAll) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCenterUserAll.ProtoReflect.Descriptor instead.
func (*UserCenterUserAll) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{1}
}

func (x *UserCenterUserAll) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserCenterUserAll) GetData() []*UserCenterUserResp {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserCenterUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	NickName string `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Rule     int32  `protobuf:"varint,5,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *UserCenterUserInfo) Reset() {
	*x = UserCenterUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCenterUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCenterUserInfo) ProtoMessage() {}

func (x *UserCenterUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCenterUserInfo.ProtoReflect.Descriptor instead.
func (*UserCenterUserInfo) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{2}
}

func (x *UserCenterUserInfo) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserCenterUserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserCenterUserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCenterUserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserCenterUserInfo) GetRule() int32 {
	if x != nil {
		return x.Rule
	}
	return 0
}

type UserCenterPageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Size  int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *UserCenterPageInfo) Reset() {
	*x = UserCenterPageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCenterPageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCenterPageInfo) ProtoMessage() {}

func (x *UserCenterPageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCenterPageInfo.ProtoReflect.Descriptor instead.
func (*UserCenterPageInfo) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{3}
}

func (x *UserCenterPageInfo) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *UserCenterPageInfo) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type UserCenterIDS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *UserCenterIDS) Reset() {
	*x = UserCenterIDS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCenterIDS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCenterIDS) ProtoMessage() {}

func (x *UserCenterIDS) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCenterIDS.ProtoReflect.Descriptor instead.
func (*UserCenterIDS) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{4}
}

func (x *UserCenterIDS) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UserCenterMobile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *UserCenterMobile) Reset() {
	*x = UserCenterMobile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCenterMobile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCenterMobile) ProtoMessage() {}

func (x *UserCenterMobile) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCenterMobile.ProtoReflect.Descriptor instead.
func (*UserCenterMobile) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{5}
}

func (x *UserCenterMobile) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type LoginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	CodeNum  string `protobuf:"bytes,3,opt,name=codeNum,proto3" json:"codeNum,omitempty"`
}

func (x *LoginInfo) Reset() {
	*x = LoginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginInfo) ProtoMessage() {}

func (x *LoginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginInfo.ProtoReflect.Descriptor instead.
func (*LoginInfo) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{6}
}

func (x *LoginInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LoginInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginInfo) GetCodeNum() string {
	if x != nil {
		return x.CodeNum
	}
	return ""
}

type UserCenterIsExists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExist bool `protobuf:"varint,1,opt,name=isExist,proto3" json:"isExist,omitempty"`
}

func (x *UserCenterIsExists) Reset() {
	*x = UserCenterIsExists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCenterIsExists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCenterIsExists) ProtoMessage() {}

func (x *UserCenterIsExists) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCenterIsExists.ProtoReflect.Descriptor instead.
func (*UserCenterIsExists) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{7}
}

func (x *UserCenterIsExists) GetIsExist() bool {
	if x != nil {
		return x.IsExist
	}
	return false
}

var File_usercenter_proto protoreflect.FileDescriptor

var file_usercenter_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6c, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x52, 0x0a,
	0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x88, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x3e, 0x0a, 0x12,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x21, 0x0a, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x44, 0x53, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0x2a, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x59, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x2e, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x32, 0xd3, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x35, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x44, 0x53, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x12, 0x36, 0x0a, 0x0c, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x1a, 0x13, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x39, 0x0a, 0x0f, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x1a, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x12, 0x5a, 0x10,
	0x2e, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usercenter_proto_rawDescOnce sync.Once
	file_usercenter_proto_rawDescData = file_usercenter_proto_rawDesc
)

func file_usercenter_proto_rawDescGZIP() []byte {
	file_usercenter_proto_rawDescOnce.Do(func() {
		file_usercenter_proto_rawDescData = protoimpl.X.CompressGZIP(file_usercenter_proto_rawDescData)
	})
	return file_usercenter_proto_rawDescData
}

var file_usercenter_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_usercenter_proto_goTypes = []interface{}{
	(*UserCenterUserResp)(nil), // 0: UserCenterUserResp
	(*UserCenterUserAll)(nil),  // 1: UserCenterUserAll
	(*UserCenterUserInfo)(nil), // 2: UserCenterUserInfo
	(*UserCenterPageInfo)(nil), // 3: UserCenterPageInfo
	(*UserCenterIDS)(nil),      // 4: UserCenterIDS
	(*UserCenterMobile)(nil),   // 5: UserCenterMobile
	(*LoginInfo)(nil),          // 6: LoginInfo
	(*UserCenterIsExists)(nil), // 7: UserCenterIsExists
	(*emptypb.Empty)(nil),      // 8: google.protobuf.Empty
}
var file_usercenter_proto_depIdxs = []int32{
	0, // 0: UserCenterUserAll.data:type_name -> UserCenterUserResp
	2, // 1: User.Create:input_type -> UserCenterUserInfo
	2, // 2: User.Update:input_type -> UserCenterUserInfo
	4, // 3: User.DeleteByIds:input_type -> UserCenterIDS
	3, // 4: User.FindPage:input_type -> UserCenterPageInfo
	5, // 5: User.FindByMobile:input_type -> UserCenterMobile
	5, // 6: User.IsExistByMobile:input_type -> UserCenterMobile
	8, // 7: User.Create:output_type -> google.protobuf.Empty
	8, // 8: User.Update:output_type -> google.protobuf.Empty
	8, // 9: User.DeleteByIds:output_type -> google.protobuf.Empty
	1, // 10: User.FindPage:output_type -> UserCenterUserAll
	2, // 11: User.FindByMobile:output_type -> UserCenterUserInfo
	7, // 12: User.IsExistByMobile:output_type -> UserCenterIsExists
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_usercenter_proto_init() }
func file_usercenter_proto_init() {
	if File_usercenter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usercenter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCenterUserResp); i {
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
		file_usercenter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCenterUserAll); i {
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
		file_usercenter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCenterUserInfo); i {
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
		file_usercenter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCenterPageInfo); i {
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
		file_usercenter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCenterIDS); i {
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
		file_usercenter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCenterMobile); i {
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
		file_usercenter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginInfo); i {
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
		file_usercenter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCenterIsExists); i {
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
			RawDescriptor: file_usercenter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usercenter_proto_goTypes,
		DependencyIndexes: file_usercenter_proto_depIdxs,
		MessageInfos:      file_usercenter_proto_msgTypes,
	}.Build()
	File_usercenter_proto = out.File
	file_usercenter_proto_rawDesc = nil
	file_usercenter_proto_goTypes = nil
	file_usercenter_proto_depIdxs = nil
}