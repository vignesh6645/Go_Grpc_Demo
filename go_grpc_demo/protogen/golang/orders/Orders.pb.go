// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: orders/Orders.proto

package orders

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ProductType int32

const (
	ProductType_UNKNOWN ProductType = 0
	ProductType_FOOD    ProductType = 1
	ProductType_DRINK   ProductType = 2
)

// Enum value maps for ProductType.
var (
	ProductType_name = map[int32]string{
		0: "UNKNOWN",
		1: "FOOD",
		2: "DRINK",
	}
	ProductType_value = map[string]int32{
		"UNKNOWN": 0,
		"FOOD":    1,
		"DRINK":   2,
	}
)

func (x ProductType) Enum() *ProductType {
	p := new(ProductType)
	*p = x
	return p
}

func (x ProductType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductType) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_Orders_proto_enumTypes[0].Descriptor()
}

func (ProductType) Type() protoreflect.EnumType {
	return &file_orders_Orders_proto_enumTypes[0]
}

func (x ProductType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductType.Descriptor instead.
func (ProductType) EnumDescriptor() ([]byte, []int) {
	return file_orders_Orders_proto_rawDescGZIP(), []int{0}
}

// import "proto/product/product.proto";
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_orders_Orders_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_orders_Orders_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_orders_Orders_proto_rawDescGZIP(), []int{0}
}

type PayloadWithSingleOrder struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PayloadWithSingleOrder) Reset() {
	*x = PayloadWithSingleOrder{}
	mi := &file_orders_Orders_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PayloadWithSingleOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadWithSingleOrder) ProtoMessage() {}

func (x *PayloadWithSingleOrder) ProtoReflect() protoreflect.Message {
	mi := &file_orders_Orders_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadWithSingleOrder.ProtoReflect.Descriptor instead.
func (*PayloadWithSingleOrder) Descriptor() ([]byte, []int) {
	return file_orders_Orders_proto_rawDescGZIP(), []int{1}
}

func (x *PayloadWithSingleOrder) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       uint64                 `protobuf:"varint,1,opt,name=order_id,proto3" json:"order_id,omitempty"`
	CustomerId    uint64                 `protobuf:"varint,2,opt,name=customer_id,proto3" json:"customer_id,omitempty"`
	IsActive      bool                   `protobuf:"varint,3,opt,name=is_active,proto3" json:"is_active,omitempty"`
	OrderDate     *timestamp.Timestamp   `protobuf:"bytes,5,opt,name=order_date,proto3" json:"order_date,omitempty"`
	Products      []*Product             `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_orders_Orders_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orders_Orders_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_orders_Orders_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Order) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Order) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Order) GetOrderDate() *timestamp.Timestamp {
	if x != nil {
		return x.OrderDate
	}
	return nil
}

func (x *Order) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint64                 `protobuf:"varint,1,opt,name=product_id,proto3" json:"product_id,omitempty"`
	ProductName   string                 `protobuf:"bytes,2,opt,name=product_name,proto3" json:"product_name,omitempty"`
	ProductType   ProductType            `protobuf:"varint,5,opt,name=product_type,proto3,enum=ProductType" json:"product_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_orders_Orders_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_orders_Orders_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_orders_Orders_proto_rawDescGZIP(), []int{3}
}

func (x *Product) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Product) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Product) GetProductType() ProductType {
	if x != nil {
		return x.ProductType
	}
	return ProductType_UNKNOWN
}

var File_orders_Orders_proto protoreflect.FileDescriptor

var file_orders_Orders_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x36, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x69, 0x74, 0x68, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xc5, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x3a, 0x0a,
	0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22,
	0x7f, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2a, 0x2f, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x4f, 0x4f, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x49, 0x4e, 0x4b, 0x10,
	0x02, 0x32, 0x37, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x57, 0x69, 0x74, 0x68, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x67, 0x6e, 0x65, 0x73, 0x68,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_Orders_proto_rawDescOnce sync.Once
	file_orders_Orders_proto_rawDescData = file_orders_Orders_proto_rawDesc
)

func file_orders_Orders_proto_rawDescGZIP() []byte {
	file_orders_Orders_proto_rawDescOnce.Do(func() {
		file_orders_Orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_Orders_proto_rawDescData)
	})
	return file_orders_Orders_proto_rawDescData
}

var file_orders_Orders_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orders_Orders_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orders_Orders_proto_goTypes = []any{
	(ProductType)(0),               // 0: ProductType
	(*Empty)(nil),                  // 1: Empty
	(*PayloadWithSingleOrder)(nil), // 2: PayloadWithSingleOrder
	(*Order)(nil),                  // 3: Order
	(*Product)(nil),                // 4: Product
	(*timestamp.Timestamp)(nil),    // 5: google.protobuf.Timestamp
}
var file_orders_Orders_proto_depIdxs = []int32{
	3, // 0: PayloadWithSingleOrder.order:type_name -> Order
	5, // 1: Order.order_date:type_name -> google.protobuf.Timestamp
	4, // 2: Order.products:type_name -> Product
	0, // 3: Product.product_type:type_name -> ProductType
	2, // 4: Orders.AddOrder:input_type -> PayloadWithSingleOrder
	1, // 5: Orders.AddOrder:output_type -> Empty
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_orders_Orders_proto_init() }
func file_orders_Orders_proto_init() {
	if File_orders_Orders_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orders_Orders_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_Orders_proto_goTypes,
		DependencyIndexes: file_orders_Orders_proto_depIdxs,
		EnumInfos:         file_orders_Orders_proto_enumTypes,
		MessageInfos:      file_orders_Orders_proto_msgTypes,
	}.Build()
	File_orders_Orders_proto = out.File
	file_orders_Orders_proto_rawDesc = nil
	file_orders_Orders_proto_goTypes = nil
	file_orders_Orders_proto_depIdxs = nil
}
