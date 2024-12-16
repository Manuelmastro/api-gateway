// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: pkg/order/pb/order.proto

package pb

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

// Messages for MakeOrder
type MakeOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`   // ID of the user placing the order
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // Delivery address
}

func (x *MakeOrderRequest) Reset() {
	*x = MakeOrderRequest{}
	mi := &file_pkg_order_pb_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MakeOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeOrderRequest) ProtoMessage() {}

func (x *MakeOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_pb_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeOrderRequest.ProtoReflect.Descriptor instead.
func (*MakeOrderRequest) Descriptor() ([]byte, []int) {
	return file_pkg_order_pb_order_proto_rawDescGZIP(), []int{0}
}

func (x *MakeOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MakeOrderRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type MakeOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`  // Status of the order creation
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Success or error message
	OrderId string `protobuf:"bytes,3,opt,name=orderId,proto3" json:"orderId,omitempty"` // ID of the created order
}

func (x *MakeOrderResponse) Reset() {
	*x = MakeOrderResponse{}
	mi := &file_pkg_order_pb_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MakeOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeOrderResponse) ProtoMessage() {}

func (x *MakeOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_pb_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeOrderResponse.ProtoReflect.Descriptor instead.
func (*MakeOrderResponse) Descriptor() ([]byte, []int) {
	return file_pkg_order_pb_order_proto_rawDescGZIP(), []int{1}
}

func (x *MakeOrderResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *MakeOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MakeOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

// Messages for ListOrders
type ListOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"` // ID of the user whose orders are to be listed
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	mi := &file_pkg_order_pb_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_pb_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_pkg_order_pb_order_proto_rawDescGZIP(), []int{2}
}

func (x *ListOrdersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"` // List of orders
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	mi := &file_pkg_order_pb_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_pb_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_pkg_order_pb_order_proto_rawDescGZIP(), []int{3}
}

func (x *ListOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

// Order structure
type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     string     `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`           // Unique ID for the order
	UserId      string     `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`             // ID of the user who placed the order
	Products    []*Product `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`         // List of products in the order
	Address     string     `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`           // Delivery address
	TotalAmount float32    `protobuf:"fixed32,5,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"` // Total cost of the order
	OrderStatus string     `protobuf:"bytes,6,opt,name=orderStatus,proto3" json:"orderStatus,omitempty"`   // Status of the order (e.g., "Pending", "Completed")
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_pkg_order_pb_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_pb_order_proto_msgTypes[4]
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
	return file_pkg_order_pb_order_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Order) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Order) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Order) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

// Product structure
type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string  `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`     // Unique ID for the product
	ProductName string  `protobuf:"bytes,2,opt,name=productName,proto3" json:"productName,omitempty"` // Name of the product
	Price       float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`           // Price of the product
	Quantity    int32   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`      // Quantity of the product
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_pkg_order_pb_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_pb_order_proto_msgTypes[5]
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
	return file_pkg_order_pb_order_proto_rawDescGZIP(), []int{5}
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_pkg_order_pb_order_proto protoreflect.FileDescriptor

var file_pkg_order_pb_order_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x44, 0x0a, 0x10, 0x4d, 0x61, 0x6b, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5f, 0x0a, 0x11, 0x4d, 0x61, 0x6b, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x22, 0xc3, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x7b, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x32, 0x95, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x4d, 0x61, 0x6b, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_order_pb_order_proto_rawDescOnce sync.Once
	file_pkg_order_pb_order_proto_rawDescData = file_pkg_order_pb_order_proto_rawDesc
)

func file_pkg_order_pb_order_proto_rawDescGZIP() []byte {
	file_pkg_order_pb_order_proto_rawDescOnce.Do(func() {
		file_pkg_order_pb_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_order_pb_order_proto_rawDescData)
	})
	return file_pkg_order_pb_order_proto_rawDescData
}

var file_pkg_order_pb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_order_pb_order_proto_goTypes = []any{
	(*MakeOrderRequest)(nil),   // 0: order.MakeOrderRequest
	(*MakeOrderResponse)(nil),  // 1: order.MakeOrderResponse
	(*ListOrdersRequest)(nil),  // 2: order.ListOrdersRequest
	(*ListOrdersResponse)(nil), // 3: order.ListOrdersResponse
	(*Order)(nil),              // 4: order.Order
	(*Product)(nil),            // 5: order.Product
}
var file_pkg_order_pb_order_proto_depIdxs = []int32{
	4, // 0: order.ListOrdersResponse.orders:type_name -> order.Order
	5, // 1: order.Order.products:type_name -> order.Product
	0, // 2: order.OrderService.MakeOrder:input_type -> order.MakeOrderRequest
	2, // 3: order.OrderService.ListOrders:input_type -> order.ListOrdersRequest
	1, // 4: order.OrderService.MakeOrder:output_type -> order.MakeOrderResponse
	3, // 5: order.OrderService.ListOrders:output_type -> order.ListOrdersResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_order_pb_order_proto_init() }
func file_pkg_order_pb_order_proto_init() {
	if File_pkg_order_pb_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_order_pb_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_order_pb_order_proto_goTypes,
		DependencyIndexes: file_pkg_order_pb_order_proto_depIdxs,
		MessageInfos:      file_pkg_order_pb_order_proto_msgTypes,
	}.Build()
	File_pkg_order_pb_order_proto = out.File
	file_pkg_order_pb_order_proto_rawDesc = nil
	file_pkg_order_pb_order_proto_goTypes = nil
	file_pkg_order_pb_order_proto_depIdxs = nil
}
