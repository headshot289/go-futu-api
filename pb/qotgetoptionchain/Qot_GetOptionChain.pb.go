// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Qot_GetOptionChain.proto

package qotgetoptionchain

import (
	_ "github.com/headshot289/go-futu-api/pb/common"
	qotcommon "github.com/headshot289/go-futu-api/pb/qotcommon"
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

type OptionCondType int32

const (
	OptionCondType_OptionCondType_Unknow  OptionCondType = 0
	OptionCondType_OptionCondType_WithIn  OptionCondType = 1 //价内
	OptionCondType_OptionCondType_Outside OptionCondType = 2 //价外
)

// Enum value maps for OptionCondType.
var (
	OptionCondType_name = map[int32]string{
		0: "OptionCondType_Unknow",
		1: "OptionCondType_WithIn",
		2: "OptionCondType_Outside",
	}
	OptionCondType_value = map[string]int32{
		"OptionCondType_Unknow":  0,
		"OptionCondType_WithIn":  1,
		"OptionCondType_Outside": 2,
	}
)

func (x OptionCondType) Enum() *OptionCondType {
	p := new(OptionCondType)
	*p = x
	return p
}

func (x OptionCondType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OptionCondType) Descriptor() protoreflect.EnumDescriptor {
	return file_Qot_GetOptionChain_proto_enumTypes[0].Descriptor()
}

func (OptionCondType) Type() protoreflect.EnumType {
	return &file_Qot_GetOptionChain_proto_enumTypes[0]
}

func (x OptionCondType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *OptionCondType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = OptionCondType(num)
	return nil
}

// Deprecated: Use OptionCondType.Descriptor instead.
func (OptionCondType) EnumDescriptor() ([]byte, []int) {
	return file_Qot_GetOptionChain_proto_rawDescGZIP(), []int{0}
}

//以下为数据字段筛选，可选字段，不填表示不过滤
type DataFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImpliedVolatilityMin *float64 `protobuf:"fixed64,1,opt,name=impliedVolatilityMin" json:"impliedVolatilityMin,omitempty"` //隐含波动率过滤起点（精确到小数点后 0 位，超出部分会被舍弃）
	ImpliedVolatilityMax *float64 `protobuf:"fixed64,2,opt,name=impliedVolatilityMax" json:"impliedVolatilityMax,omitempty"` //隐含波动率过滤终点（精确到小数点后 0 位，超出部分会被舍弃）
	DeltaMin             *float64 `protobuf:"fixed64,3,opt,name=deltaMin" json:"deltaMin,omitempty"`                         //希腊值 Delta过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
	DeltaMax             *float64 `protobuf:"fixed64,4,opt,name=deltaMax" json:"deltaMax,omitempty"`                         //希腊值 Delta过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
	GammaMin             *float64 `protobuf:"fixed64,5,opt,name=gammaMin" json:"gammaMin,omitempty"`                         //希腊值 Gamma过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
	GammaMax             *float64 `protobuf:"fixed64,6,opt,name=gammaMax" json:"gammaMax,omitempty"`                         //希腊值 Gamma过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
	VegaMin              *float64 `protobuf:"fixed64,7,opt,name=vegaMin" json:"vegaMin,omitempty"`                           //希腊值 Vega过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
	VegaMax              *float64 `protobuf:"fixed64,8,opt,name=vegaMax" json:"vegaMax,omitempty"`                           //希腊值 Vega过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
	ThetaMin             *float64 `protobuf:"fixed64,9,opt,name=thetaMin" json:"thetaMin,omitempty"`                         //希腊值 Theta过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
	ThetaMax             *float64 `protobuf:"fixed64,10,opt,name=thetaMax" json:"thetaMax,omitempty"`                        //希腊值 Theta过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
	RhoMin               *float64 `protobuf:"fixed64,11,opt,name=rhoMin" json:"rhoMin,omitempty"`                            //希腊值 Rho过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
	RhoMax               *float64 `protobuf:"fixed64,12,opt,name=rhoMax" json:"rhoMax,omitempty"`                            //希腊值 Rho过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
	NetOpenInterestMin   *float64 `protobuf:"fixed64,13,opt,name=netOpenInterestMin" json:"netOpenInterestMin,omitempty"`    //净未平仓合约数过滤起点（精确到小数点后 0 位，超出部分会被舍弃）
	NetOpenInterestMax   *float64 `protobuf:"fixed64,14,opt,name=netOpenInterestMax" json:"netOpenInterestMax,omitempty"`    //净未平仓合约数过滤终点（精确到小数点后 0 位，超出部分会被舍弃）
	OpenInterestMin      *float64 `protobuf:"fixed64,15,opt,name=openInterestMin" json:"openInterestMin,omitempty"`          //未平仓合约数过滤起点（精确到小数点后 0 位，超出部分会被舍弃）
	OpenInterestMax      *float64 `protobuf:"fixed64,16,opt,name=openInterestMax" json:"openInterestMax,omitempty"`          //未平仓合约数过滤终点（精确到小数点后 0 位，超出部分会被舍弃）
	VolMin               *float64 `protobuf:"fixed64,17,opt,name=volMin" json:"volMin,omitempty"`                            //成交量过滤起点（精确到小数点后 0 位，超出部分会被舍弃）
	VolMax               *float64 `protobuf:"fixed64,18,opt,name=volMax" json:"volMax,omitempty"`                            //成交量过滤终点（精确到小数点后 0 位，超出部分会被舍弃）
}

func (x *DataFilter) Reset() {
	*x = DataFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetOptionChain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataFilter) ProtoMessage() {}

func (x *DataFilter) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetOptionChain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataFilter.ProtoReflect.Descriptor instead.
func (*DataFilter) Descriptor() ([]byte, []int) {
	return file_Qot_GetOptionChain_proto_rawDescGZIP(), []int{0}
}

func (x *DataFilter) GetImpliedVolatilityMin() float64 {
	if x != nil && x.ImpliedVolatilityMin != nil {
		return *x.ImpliedVolatilityMin
	}
	return 0
}

func (x *DataFilter) GetImpliedVolatilityMax() float64 {
	if x != nil && x.ImpliedVolatilityMax != nil {
		return *x.ImpliedVolatilityMax
	}
	return 0
}

func (x *DataFilter) GetDeltaMin() float64 {
	if x != nil && x.DeltaMin != nil {
		return *x.DeltaMin
	}
	return 0
}

func (x *DataFilter) GetDeltaMax() float64 {
	if x != nil && x.DeltaMax != nil {
		return *x.DeltaMax
	}
	return 0
}

func (x *DataFilter) GetGammaMin() float64 {
	if x != nil && x.GammaMin != nil {
		return *x.GammaMin
	}
	return 0
}

func (x *DataFilter) GetGammaMax() float64 {
	if x != nil && x.GammaMax != nil {
		return *x.GammaMax
	}
	return 0
}

func (x *DataFilter) GetVegaMin() float64 {
	if x != nil && x.VegaMin != nil {
		return *x.VegaMin
	}
	return 0
}

func (x *DataFilter) GetVegaMax() float64 {
	if x != nil && x.VegaMax != nil {
		return *x.VegaMax
	}
	return 0
}

func (x *DataFilter) GetThetaMin() float64 {
	if x != nil && x.ThetaMin != nil {
		return *x.ThetaMin
	}
	return 0
}

func (x *DataFilter) GetThetaMax() float64 {
	if x != nil && x.ThetaMax != nil {
		return *x.ThetaMax
	}
	return 0
}

func (x *DataFilter) GetRhoMin() float64 {
	if x != nil && x.RhoMin != nil {
		return *x.RhoMin
	}
	return 0
}

func (x *DataFilter) GetRhoMax() float64 {
	if x != nil && x.RhoMax != nil {
		return *x.RhoMax
	}
	return 0
}

func (x *DataFilter) GetNetOpenInterestMin() float64 {
	if x != nil && x.NetOpenInterestMin != nil {
		return *x.NetOpenInterestMin
	}
	return 0
}

func (x *DataFilter) GetNetOpenInterestMax() float64 {
	if x != nil && x.NetOpenInterestMax != nil {
		return *x.NetOpenInterestMax
	}
	return 0
}

func (x *DataFilter) GetOpenInterestMin() float64 {
	if x != nil && x.OpenInterestMin != nil {
		return *x.OpenInterestMin
	}
	return 0
}

func (x *DataFilter) GetOpenInterestMax() float64 {
	if x != nil && x.OpenInterestMax != nil {
		return *x.OpenInterestMax
	}
	return 0
}

func (x *DataFilter) GetVolMin() float64 {
	if x != nil && x.VolMin != nil {
		return *x.VolMin
	}
	return 0
}

func (x *DataFilter) GetVolMax() float64 {
	if x != nil && x.VolMax != nil {
		return *x.VolMax
	}
	return 0
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner           *qotcommon.Security `protobuf:"bytes,1,req,name=owner" json:"owner,omitempty"`                      //期权标的股，目前仅支持传入港美正股以及恒指国指
	IndexOptionType *int32              `protobuf:"varint,6,opt,name=indexOptionType" json:"indexOptionType,omitempty"` //Qot_Common.IndexOptionType，指数期权的类型，仅用于恒指国指
	Type            *int32              `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`                       //Qot_Common.OptionType，期权类型，可选字段，不指定则表示都返回
	Condition       *int32              `protobuf:"varint,3,opt,name=condition" json:"condition,omitempty"`             //OptionCondType，价内价外，可选字段，不指定则表示都返回
	BeginTime       *string             `protobuf:"bytes,4,req,name=beginTime" json:"beginTime,omitempty"`              //期权到期日开始时间
	EndTime         *string             `protobuf:"bytes,5,req,name=endTime" json:"endTime,omitempty"`                  //期权到期日结束时间，时间跨度最多一个月
	DataFilter      *DataFilter         `protobuf:"bytes,7,opt,name=dataFilter" json:"dataFilter,omitempty"`            //数据字段筛选
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetOptionChain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetOptionChain_proto_msgTypes[1]
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
	return file_Qot_GetOptionChain_proto_rawDescGZIP(), []int{1}
}

func (x *C2S) GetOwner() *qotcommon.Security {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *C2S) GetIndexOptionType() int32 {
	if x != nil && x.IndexOptionType != nil {
		return *x.IndexOptionType
	}
	return 0
}

func (x *C2S) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *C2S) GetCondition() int32 {
	if x != nil && x.Condition != nil {
		return *x.Condition
	}
	return 0
}

func (x *C2S) GetBeginTime() string {
	if x != nil && x.BeginTime != nil {
		return *x.BeginTime
	}
	return ""
}

func (x *C2S) GetEndTime() string {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return ""
}

func (x *C2S) GetDataFilter() *DataFilter {
	if x != nil {
		return x.DataFilter
	}
	return nil
}

type OptionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Call *qotcommon.SecurityStaticInfo `protobuf:"bytes,1,opt,name=call" json:"call,omitempty"` //看涨期权，不一定有该字段，由请求条件决定
	Put  *qotcommon.SecurityStaticInfo `protobuf:"bytes,2,opt,name=put" json:"put,omitempty"`   //看跌期权，不一定有该字段，由请求条件决定
}

func (x *OptionItem) Reset() {
	*x = OptionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetOptionChain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionItem) ProtoMessage() {}

func (x *OptionItem) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetOptionChain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionItem.ProtoReflect.Descriptor instead.
func (*OptionItem) Descriptor() ([]byte, []int) {
	return file_Qot_GetOptionChain_proto_rawDescGZIP(), []int{2}
}

func (x *OptionItem) GetCall() *qotcommon.SecurityStaticInfo {
	if x != nil {
		return x.Call
	}
	return nil
}

func (x *OptionItem) GetPut() *qotcommon.SecurityStaticInfo {
	if x != nil {
		return x.Put
	}
	return nil
}

type OptionChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StrikeTime      *string       `protobuf:"bytes,1,req,name=strikeTime" json:"strikeTime,omitempty"`             //行权日
	Option          []*OptionItem `protobuf:"bytes,2,rep,name=option" json:"option,omitempty"`                     //期权信息
	StrikeTimestamp *float64      `protobuf:"fixed64,3,opt,name=strikeTimestamp" json:"strikeTimestamp,omitempty"` //行权日时间戳
}

func (x *OptionChain) Reset() {
	*x = OptionChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetOptionChain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionChain) ProtoMessage() {}

func (x *OptionChain) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetOptionChain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionChain.ProtoReflect.Descriptor instead.
func (*OptionChain) Descriptor() ([]byte, []int) {
	return file_Qot_GetOptionChain_proto_rawDescGZIP(), []int{3}
}

func (x *OptionChain) GetStrikeTime() string {
	if x != nil && x.StrikeTime != nil {
		return *x.StrikeTime
	}
	return ""
}

func (x *OptionChain) GetOption() []*OptionItem {
	if x != nil {
		return x.Option
	}
	return nil
}

func (x *OptionChain) GetStrikeTimestamp() float64 {
	if x != nil && x.StrikeTimestamp != nil {
		return *x.StrikeTimestamp
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionChain []*OptionChain `protobuf:"bytes,1,rep,name=optionChain" json:"optionChain,omitempty"` //期权链
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetOptionChain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetOptionChain_proto_msgTypes[4]
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
	return file_Qot_GetOptionChain_proto_rawDescGZIP(), []int{4}
}

func (x *S2C) GetOptionChain() []*OptionChain {
	if x != nil {
		return x.OptionChain
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
		mi := &file_Qot_GetOptionChain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetOptionChain_proto_msgTypes[5]
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
	return file_Qot_GetOptionChain_proto_rawDescGZIP(), []int{5}
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

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType，返回结果
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
		mi := &file_Qot_GetOptionChain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetOptionChain_proto_msgTypes[6]
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
	return file_Qot_GetOptionChain_proto_rawDescGZIP(), []int{6}
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

var File_Qot_GetOptionChain_proto protoreflect.FileDescriptor

var file_Qot_GetOptionChain_proto_rawDesc = []byte{
	0x0a, 0x18, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x51, 0x6f, 0x74, 0x5f,
	0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x0c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f,
	0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4,
	0x04, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a,
	0x14, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x4d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x69, 0x6d, 0x70,
	0x6c, 0x69, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69,
	0x6e, 0x12, 0x32, 0x0a, 0x14, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x14, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x4d, 0x61, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x69,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x61, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x61, 0x78, 0x12, 0x1a, 0x0a,
	0x08, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x4d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x4d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d,
	0x6d, 0x61, 0x4d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x67, 0x61, 0x6d,
	0x6d, 0x61, 0x4d, 0x61, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x67, 0x61, 0x4d, 0x69, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x76, 0x65, 0x67, 0x61, 0x4d, 0x69, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x67, 0x61, 0x4d, 0x61, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x76, 0x65, 0x67, 0x61, 0x4d, 0x61, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x4d, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x4d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x65, 0x74, 0x61, 0x4d, 0x61,
	0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x68, 0x65, 0x74, 0x61, 0x4d, 0x61,
	0x78, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x68, 0x6f, 0x4d, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x72, 0x68, 0x6f, 0x4d, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x68, 0x6f,
	0x4d, 0x61, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x68, 0x6f, 0x4d, 0x61,
	0x78, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x6e,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4d, 0x69,
	0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x78, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x6e,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4d, 0x61,
	0x78, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x4d, 0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6f,
	0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x78, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x4d, 0x69, 0x6e, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x4d, 0x69, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x6f, 0x6c, 0x4d, 0x61, 0x78, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76,
	0x6f, 0x6c, 0x4d, 0x61, 0x78, 0x22, 0x85, 0x02, 0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x2a, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51,
	0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x72, 0x0a,
	0x0a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x32, 0x0a, 0x04, 0x63,
	0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x51, 0x6f, 0x74, 0x5f,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x12,
	0x30, 0x0a, 0x03, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x51,
	0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x70, 0x75,
	0x74, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x36, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x72,
	0x69, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x48, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12, 0x41, 0x0a, 0x0b, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x52, 0x0b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x34, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03,
	0x63, 0x32, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x2a, 0x62, 0x0a,
	0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x15, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x57, 0x69, 0x74,
	0x68, 0x49, 0x6e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x10,
	0x02, 0x42, 0x4c, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x67,
	0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x6f,
	0x74, 0x67, 0x65, 0x74, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x68, 0x61, 0x69, 0x6e,
}

var (
	file_Qot_GetOptionChain_proto_rawDescOnce sync.Once
	file_Qot_GetOptionChain_proto_rawDescData = file_Qot_GetOptionChain_proto_rawDesc
)

func file_Qot_GetOptionChain_proto_rawDescGZIP() []byte {
	file_Qot_GetOptionChain_proto_rawDescOnce.Do(func() {
		file_Qot_GetOptionChain_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetOptionChain_proto_rawDescData)
	})
	return file_Qot_GetOptionChain_proto_rawDescData
}

var file_Qot_GetOptionChain_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Qot_GetOptionChain_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Qot_GetOptionChain_proto_goTypes = []interface{}{
	(OptionCondType)(0),                  // 0: Qot_GetOptionChain.OptionCondType
	(*DataFilter)(nil),                   // 1: Qot_GetOptionChain.DataFilter
	(*C2S)(nil),                          // 2: Qot_GetOptionChain.C2S
	(*OptionItem)(nil),                   // 3: Qot_GetOptionChain.OptionItem
	(*OptionChain)(nil),                  // 4: Qot_GetOptionChain.OptionChain
	(*S2C)(nil),                          // 5: Qot_GetOptionChain.S2C
	(*Request)(nil),                      // 6: Qot_GetOptionChain.Request
	(*Response)(nil),                     // 7: Qot_GetOptionChain.Response
	(*qotcommon.Security)(nil),           // 8: Qot_Common.Security
	(*qotcommon.SecurityStaticInfo)(nil), // 9: Qot_Common.SecurityStaticInfo
}
var file_Qot_GetOptionChain_proto_depIdxs = []int32{
	8, // 0: Qot_GetOptionChain.C2S.owner:type_name -> Qot_Common.Security
	1, // 1: Qot_GetOptionChain.C2S.dataFilter:type_name -> Qot_GetOptionChain.DataFilter
	9, // 2: Qot_GetOptionChain.OptionItem.call:type_name -> Qot_Common.SecurityStaticInfo
	9, // 3: Qot_GetOptionChain.OptionItem.put:type_name -> Qot_Common.SecurityStaticInfo
	3, // 4: Qot_GetOptionChain.OptionChain.option:type_name -> Qot_GetOptionChain.OptionItem
	4, // 5: Qot_GetOptionChain.S2C.optionChain:type_name -> Qot_GetOptionChain.OptionChain
	2, // 6: Qot_GetOptionChain.Request.c2s:type_name -> Qot_GetOptionChain.C2S
	5, // 7: Qot_GetOptionChain.Response.s2c:type_name -> Qot_GetOptionChain.S2C
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_Qot_GetOptionChain_proto_init() }
func file_Qot_GetOptionChain_proto_init() {
	if File_Qot_GetOptionChain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetOptionChain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataFilter); i {
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
		file_Qot_GetOptionChain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_GetOptionChain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionItem); i {
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
		file_Qot_GetOptionChain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionChain); i {
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
		file_Qot_GetOptionChain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_GetOptionChain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_GetOptionChain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_Qot_GetOptionChain_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetOptionChain_proto_goTypes,
		DependencyIndexes: file_Qot_GetOptionChain_proto_depIdxs,
		EnumInfos:         file_Qot_GetOptionChain_proto_enumTypes,
		MessageInfos:      file_Qot_GetOptionChain_proto_msgTypes,
	}.Build()
	File_Qot_GetOptionChain_proto = out.File
	file_Qot_GetOptionChain_proto_rawDesc = nil
	file_Qot_GetOptionChain_proto_goTypes = nil
	file_Qot_GetOptionChain_proto_depIdxs = nil
}
