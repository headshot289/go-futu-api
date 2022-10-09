// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Qot_GetUserSecurityGroup.proto

package qotgetusersecuritygroup

import (
	_ "github.com/headshot289/go-futu-api/pb/common"
	_ "github.com/headshot289/go-futu-api/pb/qotcommon"
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

// 自选股分组类型
type GroupType int32

const (
	GroupType_GroupType_Unknown GroupType = 0 // 未知
	GroupType_GroupType_Custom  GroupType = 1 // 自定义分组
	GroupType_GroupType_System  GroupType = 2 // 系统分组
	GroupType_GroupType_All     GroupType = 3 // 全部分组
)

// Enum value maps for GroupType.
var (
	GroupType_name = map[int32]string{
		0: "GroupType_Unknown",
		1: "GroupType_Custom",
		2: "GroupType_System",
		3: "GroupType_All",
	}
	GroupType_value = map[string]int32{
		"GroupType_Unknown": 0,
		"GroupType_Custom":  1,
		"GroupType_System":  2,
		"GroupType_All":     3,
	}
)

func (x GroupType) Enum() *GroupType {
	p := new(GroupType)
	*p = x
	return p
}

func (x GroupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupType) Descriptor() protoreflect.EnumDescriptor {
	return file_Qot_GetUserSecurityGroup_proto_enumTypes[0].Descriptor()
}

func (GroupType) Type() protoreflect.EnumType {
	return &file_Qot_GetUserSecurityGroup_proto_enumTypes[0]
}

func (x GroupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GroupType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GroupType(num)
	return nil
}

// Deprecated: Use GroupType.Descriptor instead.
func (GroupType) EnumDescriptor() ([]byte, []int) {
	return file_Qot_GetUserSecurityGroup_proto_rawDescGZIP(), []int{0}
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupType *int32 `protobuf:"varint,1,req,name=groupType" json:"groupType,omitempty"` // GroupType,自选股分组类型。
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Qot_GetUserSecurityGroup_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetGroupType() int32 {
	if x != nil && x.GroupType != nil {
		return *x.GroupType
	}
	return 0
}

type GroupData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName *string `protobuf:"bytes,1,req,name=groupName" json:"groupName,omitempty"`  // 自选股分组名字
	GroupType *int32  `protobuf:"varint,2,req,name=groupType" json:"groupType,omitempty"` // GroupType,自选股分组类型。
}

func (x *GroupData) Reset() {
	*x = GroupData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupData) ProtoMessage() {}

func (x *GroupData) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupData.ProtoReflect.Descriptor instead.
func (*GroupData) Descriptor() ([]byte, []int) {
	return file_Qot_GetUserSecurityGroup_proto_rawDescGZIP(), []int{1}
}

func (x *GroupData) GetGroupName() string {
	if x != nil && x.GroupName != nil {
		return *x.GroupName
	}
	return ""
}

func (x *GroupData) GetGroupType() int32 {
	if x != nil && x.GroupType != nil {
		return *x.GroupType
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupList []*GroupData `protobuf:"bytes,1,rep,name=groupList" json:"groupList,omitempty"` // 自选股分组列表
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Qot_GetUserSecurityGroup_proto_rawDescGZIP(), []int{2}
}

func (x *S2C) GetGroupList() []*GroupData {
	if x != nil {
		return x.GroupList
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Qot_GetUserSecurityGroup_proto_rawDescGZIP(), []int{3}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetUserSecurityGroup_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Qot_GetUserSecurityGroup_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Qot_GetUserSecurityGroup_proto protoreflect.FileDescriptor

var file_Qot_GetUserSecurityGroup_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x03, 0x43, 0x32,
	0x53, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x47, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12,
	0x41, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x3a, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a,
	0x03, 0x63, 0x32, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x8d,
	0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34,
	0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74,
	0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a,
	0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x2a, 0x61,
	0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x41, 0x6c, 0x6c, 0x10,
	0x03, 0x42, 0x52, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x67,
	0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x6f,
	0x74, 0x67, 0x65, 0x74, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x70,
}

var (
	file_Qot_GetUserSecurityGroup_proto_rawDescOnce sync.Once
	file_Qot_GetUserSecurityGroup_proto_rawDescData = file_Qot_GetUserSecurityGroup_proto_rawDesc
)

func file_Qot_GetUserSecurityGroup_proto_rawDescGZIP() []byte {
	file_Qot_GetUserSecurityGroup_proto_rawDescOnce.Do(func() {
		file_Qot_GetUserSecurityGroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetUserSecurityGroup_proto_rawDescData)
	})
	return file_Qot_GetUserSecurityGroup_proto_rawDescData
}

var file_Qot_GetUserSecurityGroup_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Qot_GetUserSecurityGroup_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Qot_GetUserSecurityGroup_proto_goTypes = []interface{}{
	(GroupType)(0),    // 0: Qot_GetUserSecurityGroup.GroupType
	(*C2S)(nil),       // 1: Qot_GetUserSecurityGroup.C2S
	(*GroupData)(nil), // 2: Qot_GetUserSecurityGroup.GroupData
	(*S2C)(nil),       // 3: Qot_GetUserSecurityGroup.S2C
	(*Request)(nil),   // 4: Qot_GetUserSecurityGroup.Request
	(*Response)(nil),  // 5: Qot_GetUserSecurityGroup.Response
}
var file_Qot_GetUserSecurityGroup_proto_depIdxs = []int32{
	2, // 0: Qot_GetUserSecurityGroup.S2C.groupList:type_name -> Qot_GetUserSecurityGroup.GroupData
	1, // 1: Qot_GetUserSecurityGroup.Request.c2s:type_name -> Qot_GetUserSecurityGroup.C2S
	3, // 2: Qot_GetUserSecurityGroup.Response.s2c:type_name -> Qot_GetUserSecurityGroup.S2C
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Qot_GetUserSecurityGroup_proto_init() }
func file_Qot_GetUserSecurityGroup_proto_init() {
	if File_Qot_GetUserSecurityGroup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetUserSecurityGroup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_Qot_GetUserSecurityGroup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupData); i {
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
		file_Qot_GetUserSecurityGroup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_Qot_GetUserSecurityGroup_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Qot_GetUserSecurityGroup_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_Qot_GetUserSecurityGroup_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetUserSecurityGroup_proto_goTypes,
		DependencyIndexes: file_Qot_GetUserSecurityGroup_proto_depIdxs,
		EnumInfos:         file_Qot_GetUserSecurityGroup_proto_enumTypes,
		MessageInfos:      file_Qot_GetUserSecurityGroup_proto_msgTypes,
	}.Build()
	File_Qot_GetUserSecurityGroup_proto = out.File
	file_Qot_GetUserSecurityGroup_proto_rawDesc = nil
	file_Qot_GetUserSecurityGroup_proto_goTypes = nil
	file_Qot_GetUserSecurityGroup_proto_depIdxs = nil
}
