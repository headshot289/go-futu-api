// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Trd_ModifyOrder.proto

package trdmodifyorder

import (
	common "github.com/headshot289/go-futu-api/pb/common"
	trdcommon "github.com/headshot289/go-futu-api/pb/trdcommon"
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

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketID      *common.PacketID     `protobuf:"bytes,1,req,name=packetID" json:"packetID,omitempty"`            //交易写操作防重放攻击
	Header        *trdcommon.TrdHeader `protobuf:"bytes,2,req,name=header" json:"header,omitempty"`                //交易公共参数头
	OrderID       *uint64              `protobuf:"varint,3,req,name=orderID" json:"orderID,omitempty"`             //订单号，forAll为true时，传0
	ModifyOrderOp *int32               `protobuf:"varint,4,req,name=modifyOrderOp" json:"modifyOrderOp,omitempty"` //修改操作类型，参见Trd_Common.ModifyOrderOp的枚举定义
	ForAll        *bool                `protobuf:"varint,5,opt,name=forAll" json:"forAll,omitempty"`               //是否对此业务账户的全部订单操作，true是，false否(对单个订单)，无此字段代表false，仅对单个订单
	//下面的字段仅针对单个订单，且modifyOrderOp为ModifyOrderOp_Normal有效
	Qty   *float64 `protobuf:"fixed64,8,opt,name=qty" json:"qty,omitempty"`     //数量，期权单位是"张"（精确到小数点后 0 位，超出部分会被舍弃）
	Price *float64 `protobuf:"fixed64,9,opt,name=price" json:"price,omitempty"` //价格，（证券账户精确到小数点后 3 位，期货账户精确到小数点后 9 位，超出部分会被舍弃）
	//以下为调整价格使用，都传才有效，对港、A股有意义，因为港股有价位，A股2位精度，美股可不传
	AdjustPrice        *bool    `protobuf:"varint,10,opt,name=adjustPrice" json:"adjustPrice,omitempty"`                //是否调整价格，如果价格不合法，是否调整到合法价位，true调整，false不调整
	AdjustSideAndLimit *float64 `protobuf:"fixed64,11,opt,name=adjustSideAndLimit" json:"adjustSideAndLimit,omitempty"` //调整方向和调整幅度百分比限制，正数代表向上调整，负数代表向下调整，具体值代表调整幅度限制，如：0.015代表向上调整且幅度不超过1.5%；-0.01代表向下调整且幅度不超过1%
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Trd_ModifyOrder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Trd_ModifyOrder_proto_msgTypes[0]
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
	return file_Trd_ModifyOrder_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetPacketID() *common.PacketID {
	if x != nil {
		return x.PacketID
	}
	return nil
}

func (x *C2S) GetHeader() *trdcommon.TrdHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *C2S) GetOrderID() uint64 {
	if x != nil && x.OrderID != nil {
		return *x.OrderID
	}
	return 0
}

func (x *C2S) GetModifyOrderOp() int32 {
	if x != nil && x.ModifyOrderOp != nil {
		return *x.ModifyOrderOp
	}
	return 0
}

func (x *C2S) GetForAll() bool {
	if x != nil && x.ForAll != nil {
		return *x.ForAll
	}
	return false
}

func (x *C2S) GetQty() float64 {
	if x != nil && x.Qty != nil {
		return *x.Qty
	}
	return 0
}

func (x *C2S) GetPrice() float64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *C2S) GetAdjustPrice() bool {
	if x != nil && x.AdjustPrice != nil {
		return *x.AdjustPrice
	}
	return false
}

func (x *C2S) GetAdjustSideAndLimit() float64 {
	if x != nil && x.AdjustSideAndLimit != nil {
		return *x.AdjustSideAndLimit
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *trdcommon.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`    //交易公共参数头
	OrderID *uint64              `protobuf:"varint,2,req,name=orderID" json:"orderID,omitempty"` //订单号
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Trd_ModifyOrder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Trd_ModifyOrder_proto_msgTypes[1]
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
	return file_Trd_ModifyOrder_proto_rawDescGZIP(), []int{1}
}

func (x *S2C) GetHeader() *trdcommon.TrdHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *S2C) GetOrderID() uint64 {
	if x != nil && x.OrderID != nil {
		return *x.OrderID
	}
	return 0
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
		mi := &file_Trd_ModifyOrder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Trd_ModifyOrder_proto_msgTypes[2]
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
	return file_Trd_ModifyOrder_proto_rawDescGZIP(), []int{2}
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

	//以下3个字段每条协议都有，注释说明在InitConnect.proto中
	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
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
		mi := &file_Trd_ModifyOrder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Trd_ModifyOrder_proto_msgTypes[3]
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
	return file_Trd_ModifyOrder_proto_rawDescGZIP(), []int{3}
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

var File_Trd_ModifyOrder_proto protoreflect.FileDescriptor

var file_Trd_ModifyOrder_proto_rawDesc = []byte{
	0x0a, 0x15, 0x54, 0x72, 0x64, 0x5f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x54, 0x72, 0x64, 0x5f, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x54, 0x72, 0x64, 0x5f, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x03, 0x43, 0x32, 0x53,
	0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x44, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x2d,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x54, 0x72, 0x64, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x64, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x70, 0x18, 0x04, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0d,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66,
	0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x12, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x41, 0x6e, 0x64,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x61, 0x64, 0x6a,
	0x75, 0x73, 0x74, 0x53, 0x69, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x4e, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x72, 0x64, 0x5f, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x31, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x03, 0x63, 0x32,
	0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x54, 0x72, 0x64, 0x5f, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63,
	0x32, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05,
	0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x26, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x54, 0x72, 0x64, 0x5f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x42, 0x49, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72,
	0x69, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x6f,
	0x72, 0x64, 0x65, 0x72,
}

var (
	file_Trd_ModifyOrder_proto_rawDescOnce sync.Once
	file_Trd_ModifyOrder_proto_rawDescData = file_Trd_ModifyOrder_proto_rawDesc
)

func file_Trd_ModifyOrder_proto_rawDescGZIP() []byte {
	file_Trd_ModifyOrder_proto_rawDescOnce.Do(func() {
		file_Trd_ModifyOrder_proto_rawDescData = protoimpl.X.CompressGZIP(file_Trd_ModifyOrder_proto_rawDescData)
	})
	return file_Trd_ModifyOrder_proto_rawDescData
}

var file_Trd_ModifyOrder_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Trd_ModifyOrder_proto_goTypes = []interface{}{
	(*C2S)(nil),                 // 0: Trd_ModifyOrder.C2S
	(*S2C)(nil),                 // 1: Trd_ModifyOrder.S2C
	(*Request)(nil),             // 2: Trd_ModifyOrder.Request
	(*Response)(nil),            // 3: Trd_ModifyOrder.Response
	(*common.PacketID)(nil),     // 4: Common.PacketID
	(*trdcommon.TrdHeader)(nil), // 5: Trd_Common.TrdHeader
}
var file_Trd_ModifyOrder_proto_depIdxs = []int32{
	4, // 0: Trd_ModifyOrder.C2S.packetID:type_name -> Common.PacketID
	5, // 1: Trd_ModifyOrder.C2S.header:type_name -> Trd_Common.TrdHeader
	5, // 2: Trd_ModifyOrder.S2C.header:type_name -> Trd_Common.TrdHeader
	0, // 3: Trd_ModifyOrder.Request.c2s:type_name -> Trd_ModifyOrder.C2S
	1, // 4: Trd_ModifyOrder.Response.s2c:type_name -> Trd_ModifyOrder.S2C
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Trd_ModifyOrder_proto_init() }
func file_Trd_ModifyOrder_proto_init() {
	if File_Trd_ModifyOrder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Trd_ModifyOrder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Trd_ModifyOrder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Trd_ModifyOrder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_Trd_ModifyOrder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_Trd_ModifyOrder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Trd_ModifyOrder_proto_goTypes,
		DependencyIndexes: file_Trd_ModifyOrder_proto_depIdxs,
		MessageInfos:      file_Trd_ModifyOrder_proto_msgTypes,
	}.Build()
	File_Trd_ModifyOrder_proto = out.File
	file_Trd_ModifyOrder_proto_rawDesc = nil
	file_Trd_ModifyOrder_proto_goTypes = nil
	file_Trd_ModifyOrder_proto_depIdxs = nil
}
