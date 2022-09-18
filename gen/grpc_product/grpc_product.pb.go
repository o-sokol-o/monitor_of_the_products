// protoc --go_out=. --go-grpc_out=. proto/grpc_product.proto
// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/grpc_product.proto
// bad: protoc --go_out=plugins=grpc:proto proto/grpc_log.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/grpc_product.proto

package grpc_product

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

type RequestFetch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RequestFetch) Reset() {
	*x = RequestFetch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFetch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFetch) ProtoMessage() {}

func (x *RequestFetch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFetch.ProtoReflect.Descriptor instead.
func (*RequestFetch) Descriptor() ([]byte, []int) {
	return file_proto_grpc_product_proto_rawDescGZIP(), []int{0}
}

func (x *RequestFetch) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ResponseFetch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewProductCount    int32 `protobuf:"varint,1,opt,name=new_product_count,json=newProductCount,proto3" json:"new_product_count,omitempty"`
	UpdateProductCount int32 `protobuf:"varint,2,opt,name=update_product_count,json=updateProductCount,proto3" json:"update_product_count,omitempty"`
}

func (x *ResponseFetch) Reset() {
	*x = ResponseFetch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFetch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFetch) ProtoMessage() {}

func (x *ResponseFetch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFetch.ProtoReflect.Descriptor instead.
func (*ResponseFetch) Descriptor() ([]byte, []int) {
	return file_proto_grpc_product_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseFetch) GetNewProductCount() int32 {
	if x != nil {
		return x.NewProductCount
	}
	return 0
}

func (x *ResponseFetch) GetUpdateProductCount() int32 {
	if x != nil {
		return x.UpdateProductCount
	}
	return 0
}

type Paging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumPage        int32 `protobuf:"varint,1,opt,name=NumPage,proto3" json:"NumPage,omitempty"`
	ProductPerPage int32 `protobuf:"varint,2,opt,name=ProductPerPage,proto3" json:"ProductPerPage,omitempty"`
}

func (x *Paging) Reset() {
	*x = Paging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paging) ProtoMessage() {}

func (x *Paging) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paging.ProtoReflect.Descriptor instead.
func (*Paging) Descriptor() ([]byte, []int) {
	return file_proto_grpc_product_proto_rawDescGZIP(), []int{2}
}

func (x *Paging) GetNumPage() int32 {
	if x != nil {
		return x.NumPage
	}
	return 0
}

func (x *Paging) GetProductPerPage() int32 {
	if x != nil {
		return x.ProductPerPage
	}
	return 0
}

type Sorting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortField int32 `protobuf:"varint,1,opt,name=SortField,proto3" json:"SortField,omitempty"`
	Direction int32 `protobuf:"varint,2,opt,name=Direction,proto3" json:"Direction,omitempty"`
}

func (x *Sorting) Reset() {
	*x = Sorting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sorting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sorting) ProtoMessage() {}

func (x *Sorting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sorting.ProtoReflect.Descriptor instead.
func (*Sorting) Descriptor() ([]byte, []int) {
	return file_proto_grpc_product_proto_rawDescGZIP(), []int{3}
}

func (x *Sorting) GetSortField() int32 {
	if x != nil {
		return x.SortField
	}
	return 0
}

func (x *Sorting) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

type RequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging  *Paging  `protobuf:"bytes,1,opt,name=Paging,proto3" json:"Paging,omitempty"`
	Sorting *Sorting `protobuf:"bytes,2,opt,name=Sorting,proto3" json:"Sorting,omitempty"`
}

func (x *RequestList) Reset() {
	*x = RequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestList) ProtoMessage() {}

func (x *RequestList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestList.ProtoReflect.Descriptor instead.
func (*RequestList) Descriptor() ([]byte, []int) {
	return file_proto_grpc_product_proto_rawDescGZIP(), []int{4}
}

func (x *RequestList) GetPaging() *Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *RequestList) GetSorting() *Sorting {
	if x != nil {
		return x.Sorting
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Cost        int32                  `protobuf:"varint,2,opt,name=Cost,proto3" json:"Cost,omitempty"`
	ChangeCount int32                  `protobuf:"varint,3,opt,name=ChangeCount,proto3" json:"ChangeCount,omitempty"`
	UpdateAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_proto_grpc_product_proto_rawDescGZIP(), []int{5}
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetCost() int32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Product) GetChangeCount() int32 {
	if x != nil {
		return x.ChangeCount
	}
	return 0
}

func (x *Product) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

type ResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=Products,proto3" json:"Products,omitempty"`
}

func (x *ResponseList) Reset() {
	*x = ResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseList) ProtoMessage() {}

func (x *ResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseList.ProtoReflect.Descriptor instead.
func (*ResponseList) Descriptor() ([]byte, []int) {
	return file_proto_grpc_product_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseList) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_proto_grpc_product_proto protoreflect.FileDescriptor

var file_proto_grpc_product_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x6d, 0x0a, 0x0d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x2a, 0x0a, 0x11,
	0x6e, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x06, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4e, 0x75, 0x6d, 0x50, 0x61, 0x67, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x07, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a,
	0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x52, 0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x07, 0x53, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x07, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x8b, 0x01, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x36, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x9a, 0x01, 0x0a,
	0x13, 0x47, 0x72, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x46, 0x65, 0x74, 0x63, 0x68, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_product_proto_rawDescOnce sync.Once
	file_proto_grpc_product_proto_rawDescData = file_proto_grpc_product_proto_rawDesc
)

func file_proto_grpc_product_proto_rawDescGZIP() []byte {
	file_proto_grpc_product_proto_rawDescOnce.Do(func() {
		file_proto_grpc_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_product_proto_rawDescData)
	})
	return file_proto_grpc_product_proto_rawDescData
}

var file_proto_grpc_product_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_grpc_product_proto_goTypes = []interface{}{
	(*RequestFetch)(nil),          // 0: grpc_product.RequestFetch
	(*ResponseFetch)(nil),         // 1: grpc_product.ResponseFetch
	(*Paging)(nil),                // 2: grpc_product.Paging
	(*Sorting)(nil),               // 3: grpc_product.Sorting
	(*RequestList)(nil),           // 4: grpc_product.RequestList
	(*Product)(nil),               // 5: grpc_product.Product
	(*ResponseList)(nil),          // 6: grpc_product.ResponseList
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_proto_grpc_product_proto_depIdxs = []int32{
	2, // 0: grpc_product.RequestList.Paging:type_name -> grpc_product.Paging
	3, // 1: grpc_product.RequestList.Sorting:type_name -> grpc_product.Sorting
	7, // 2: grpc_product.Product.UpdateAt:type_name -> google.protobuf.Timestamp
	5, // 3: grpc_product.ResponseList.Products:type_name -> grpc_product.Product
	0, // 4: grpc_product.GrpcProductsService.Fetch:input_type -> grpc_product.RequestFetch
	4, // 5: grpc_product.GrpcProductsService.List:input_type -> grpc_product.RequestList
	1, // 6: grpc_product.GrpcProductsService.Fetch:output_type -> grpc_product.ResponseFetch
	6, // 7: grpc_product.GrpcProductsService.List:output_type -> grpc_product.ResponseList
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_grpc_product_proto_init() }
func file_proto_grpc_product_proto_init() {
	if File_proto_grpc_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFetch); i {
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
		file_proto_grpc_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFetch); i {
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
		file_proto_grpc_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paging); i {
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
		file_proto_grpc_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sorting); i {
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
		file_proto_grpc_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestList); i {
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
		file_proto_grpc_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_proto_grpc_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseList); i {
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
			RawDescriptor: file_proto_grpc_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_product_proto_goTypes,
		DependencyIndexes: file_proto_grpc_product_proto_depIdxs,
		MessageInfos:      file_proto_grpc_product_proto_msgTypes,
	}.Build()
	File_proto_grpc_product_proto = out.File
	file_proto_grpc_product_proto_rawDesc = nil
	file_proto_grpc_product_proto_goTypes = nil
	file_proto_grpc_product_proto_depIdxs = nil
}