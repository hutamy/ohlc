// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ohlc.proto

package __

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

type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockCode string `protobuf:"bytes,1,opt,name=stock_code,json=stockCode,proto3" json:"stock_code,omitempty"`
	Prev      int32  `protobuf:"varint,2,opt,name=prev,proto3" json:"prev,omitempty"`
	Open      int32  `protobuf:"varint,3,opt,name=open,proto3" json:"open,omitempty"`
	High      int32  `protobuf:"varint,4,opt,name=high,proto3" json:"high,omitempty"`
	Low       int32  `protobuf:"varint,5,opt,name=low,proto3" json:"low,omitempty"`
	Close     int32  `protobuf:"varint,6,opt,name=close,proto3" json:"close,omitempty"`
	Average   int32  `protobuf:"varint,7,opt,name=average,proto3" json:"average,omitempty"`
	Volume    int32  `protobuf:"varint,8,opt,name=volume,proto3" json:"volume,omitempty"`
	Value     int32  `protobuf:"varint,9,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ohlc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_ohlc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_ohlc_proto_rawDescGZIP(), []int{0}
}

func (x *Summary) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

func (x *Summary) GetPrev() int32 {
	if x != nil {
		return x.Prev
	}
	return 0
}

func (x *Summary) GetOpen() int32 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Summary) GetHigh() int32 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Summary) GetLow() int32 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Summary) GetClose() int32 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *Summary) GetAverage() int32 {
	if x != nil {
		return x.Average
	}
	return 0
}

func (x *Summary) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Summary) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockCode string `protobuf:"bytes,1,opt,name=stock_code,json=stockCode,proto3" json:"stock_code,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Price     int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity  int32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ohlc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_ohlc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_ohlc_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

func (x *Transaction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Transaction) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Transaction) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type StockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockCode string `protobuf:"bytes,1,opt,name=stock_code,json=stockCode,proto3" json:"stock_code,omitempty"`
}

func (x *StockRequest) Reset() {
	*x = StockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ohlc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockRequest) ProtoMessage() {}

func (x *StockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ohlc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockRequest.ProtoReflect.Descriptor instead.
func (*StockRequest) Descriptor() ([]byte, []int) {
	return file_ohlc_proto_rawDescGZIP(), []int{2}
}

func (x *StockRequest) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

var File_ohlc_proto protoreflect.FileDescriptor

var file_ohlc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6f, 0x68, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6f, 0x68,
	0x6c, 0x63, 0x22, 0xd4, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x72, 0x65, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x72, 0x65,
	0x76, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x72, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2d, 0x0a,
	0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x3b, 0x0a, 0x0b,
	0x4f, 0x48, 0x4c, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4f, 0x48, 0x4c, 0x43, 0x12, 0x12, 0x2e, 0x6f, 0x68, 0x6c, 0x63, 0x2e, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6f, 0x68, 0x6c,
	0x63, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ohlc_proto_rawDescOnce sync.Once
	file_ohlc_proto_rawDescData = file_ohlc_proto_rawDesc
)

func file_ohlc_proto_rawDescGZIP() []byte {
	file_ohlc_proto_rawDescOnce.Do(func() {
		file_ohlc_proto_rawDescData = protoimpl.X.CompressGZIP(file_ohlc_proto_rawDescData)
	})
	return file_ohlc_proto_rawDescData
}

var file_ohlc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ohlc_proto_goTypes = []interface{}{
	(*Summary)(nil),      // 0: ohlc.Summary
	(*Transaction)(nil),  // 1: ohlc.Transaction
	(*StockRequest)(nil), // 2: ohlc.StockRequest
}
var file_ohlc_proto_depIdxs = []int32{
	2, // 0: ohlc.OHLCService.GetOHLC:input_type -> ohlc.StockRequest
	0, // 1: ohlc.OHLCService.GetOHLC:output_type -> ohlc.Summary
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ohlc_proto_init() }
func file_ohlc_proto_init() {
	if File_ohlc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ohlc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary); i {
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
		file_ohlc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_ohlc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockRequest); i {
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
			RawDescriptor: file_ohlc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ohlc_proto_goTypes,
		DependencyIndexes: file_ohlc_proto_depIdxs,
		MessageInfos:      file_ohlc_proto_msgTypes,
	}.Build()
	File_ohlc_proto = out.File
	file_ohlc_proto_rawDesc = nil
	file_ohlc_proto_goTypes = nil
	file_ohlc_proto_depIdxs = nil
}
