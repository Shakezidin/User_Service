// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: coordinator.proto

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

type CoodinatorViewPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId int64 `protobuf:"varint,1,opt,name=packageId,proto3" json:"packageId,omitempty"`
}

func (x *CoodinatorViewPackage) Reset() {
	*x = CoodinatorViewPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoodinatorViewPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoodinatorViewPackage) ProtoMessage() {}

func (x *CoodinatorViewPackage) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoodinatorViewPackage.ProtoReflect.Descriptor instead.
func (*CoodinatorViewPackage) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{0}
}

func (x *CoodinatorViewPackage) GetPackageId() int64 {
	if x != nil {
		return x.PackageId
	}
	return 0
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryName string `protobuf:"bytes,1,opt,name=categoryName,proto3" json:"categoryName,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{1}
}

func (x *Category) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

type Destinations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationName string `protobuf:"bytes,1,opt,name=DestinationName,proto3" json:"DestinationName,omitempty"`
	Description     string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	PackageID       int64  `protobuf:"varint,3,opt,name=PackageID,proto3" json:"PackageID,omitempty"`
	MinPrice        int64  `protobuf:"varint,4,opt,name=MinPrice,proto3" json:"MinPrice,omitempty"`
	MaxCapacity     int64  `protobuf:"varint,5,opt,name=MaxCapacity,proto3" json:"MaxCapacity,omitempty"`
	Image           string `protobuf:"bytes,6,opt,name=Image,proto3" json:"Image,omitempty"`
	DestinationId   int32  `protobuf:"varint,7,opt,name=DestinationId,proto3" json:"DestinationId,omitempty"`
}

func (x *Destinations) Reset() {
	*x = Destinations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Destinations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Destinations) ProtoMessage() {}

func (x *Destinations) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Destinations.ProtoReflect.Descriptor instead.
func (*Destinations) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{2}
}

func (x *Destinations) GetDestinationName() string {
	if x != nil {
		return x.DestinationName
	}
	return ""
}

func (x *Destinations) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Destinations) GetPackageID() int64 {
	if x != nil {
		return x.PackageID
	}
	return 0
}

func (x *Destinations) GetMinPrice() int64 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *Destinations) GetMaxCapacity() int64 {
	if x != nil {
		return x.MaxCapacity
	}
	return 0
}

func (x *Destinations) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Destinations) GetDestinationId() int32 {
	if x != nil {
		return x.DestinationId
	}
	return 0
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Startlocation    string          `protobuf:"bytes,2,opt,name=startlocation,proto3" json:"startlocation,omitempty"`
	Endlocation      string          `protobuf:"bytes,3,opt,name=endlocation,proto3" json:"endlocation,omitempty"`
	Startdatetime    string          `protobuf:"bytes,4,opt,name=startdatetime,proto3" json:"startdatetime,omitempty"`
	Enddatetime      string          `protobuf:"bytes,5,opt,name=enddatetime,proto3" json:"enddatetime,omitempty"`
	Price            int32           `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
	Image            string          `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
	DestinationCount int32           `protobuf:"varint,8,opt,name=destinationCount,proto3" json:"destinationCount,omitempty"`
	Destination      string          `protobuf:"bytes,9,opt,name=destination,proto3" json:"destination,omitempty"`
	PackageId        int64           `protobuf:"varint,10,opt,name=PackageId,proto3" json:"PackageId,omitempty"`
	Description      string          `protobuf:"bytes,11,opt,name=Description,proto3" json:"Description,omitempty"`
	Category         *Category       `protobuf:"bytes,12,opt,name=category,proto3" json:"category,omitempty"`
	Destinations     []*Destinations `protobuf:"bytes,13,rep,name=Destinations,proto3" json:"Destinations,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{3}
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetStartlocation() string {
	if x != nil {
		return x.Startlocation
	}
	return ""
}

func (x *Package) GetEndlocation() string {
	if x != nil {
		return x.Endlocation
	}
	return ""
}

func (x *Package) GetStartdatetime() string {
	if x != nil {
		return x.Startdatetime
	}
	return ""
}

func (x *Package) GetEnddatetime() string {
	if x != nil {
		return x.Enddatetime
	}
	return ""
}

func (x *Package) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Package) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Package) GetDestinationCount() int32 {
	if x != nil {
		return x.DestinationCount
	}
	return 0
}

func (x *Package) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Package) GetPackageId() int64 {
	if x != nil {
		return x.PackageId
	}
	return 0
}

func (x *Package) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Package) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Package) GetDestinations() []*Destinations {
	if x != nil {
		return x.Destinations
	}
	return nil
}

type AllPackages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AllPackages) Reset() {
	*x = AllPackages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllPackages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllPackages) ProtoMessage() {}

func (x *AllPackages) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllPackages.ProtoReflect.Descriptor instead.
func (*AllPackages) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{4}
}

type PackagesResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packages []*Package `protobuf:"bytes,1,rep,name=Packages,proto3" json:"Packages,omitempty"`
	Token    string     `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PackagesResponce) Reset() {
	*x = PackagesResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackagesResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackagesResponce) ProtoMessage() {}

func (x *PackagesResponce) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackagesResponce.ProtoReflect.Descriptor instead.
func (*PackagesResponce) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{5}
}

func (x *PackagesResponce) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

func (x *PackagesResponce) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_coordinator_proto protoreflect.FileDescriptor

var file_coordinator_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x35, 0x0a, 0x15, 0x43, 0x6f, 0x6f, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x2e,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf2,
	0x01, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x28, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4d, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x78, 0x43, 0x61, 0x70, 0x61,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x4d, 0x61, 0x78, 0x43,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xc7, 0x03, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x64,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x6e, 0x64, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x2a, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x34, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0c, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x0d, 0x0a,
	0x0b, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x22, 0x51, 0x0a, 0x10,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x08, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32,
	0x8b, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x3a, 0x0a, 0x11, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x16, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6f, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coordinator_proto_rawDescOnce sync.Once
	file_coordinator_proto_rawDescData = file_coordinator_proto_rawDesc
)

func file_coordinator_proto_rawDescGZIP() []byte {
	file_coordinator_proto_rawDescOnce.Do(func() {
		file_coordinator_proto_rawDescData = protoimpl.X.CompressGZIP(file_coordinator_proto_rawDescData)
	})
	return file_coordinator_proto_rawDescData
}

var file_coordinator_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_coordinator_proto_goTypes = []interface{}{
	(*CoodinatorViewPackage)(nil), // 0: pb.CoodinatorViewPackage
	(*Category)(nil),              // 1: pb.Category
	(*Destinations)(nil),          // 2: pb.Destinations
	(*Package)(nil),               // 3: pb.Package
	(*AllPackages)(nil),           // 4: pb.AllPackages
	(*PackagesResponce)(nil),      // 5: pb.PackagesResponce
}
var file_coordinator_proto_depIdxs = []int32{
	1, // 0: pb.Package.category:type_name -> pb.Category
	2, // 1: pb.Package.Destinations:type_name -> pb.Destinations
	3, // 2: pb.PackagesResponce.Packages:type_name -> pb.Package
	4, // 3: pb.Coordinator.AvailablePackages:input_type -> pb.AllPackages
	0, // 4: pb.Coordinator.CoordinatorViewPackage:input_type -> pb.CoodinatorViewPackage
	5, // 5: pb.Coordinator.AvailablePackages:output_type -> pb.PackagesResponce
	3, // 6: pb.Coordinator.CoordinatorViewPackage:output_type -> pb.Package
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_coordinator_proto_init() }
func file_coordinator_proto_init() {
	if File_coordinator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coordinator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoodinatorViewPackage); i {
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
		file_coordinator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_coordinator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Destinations); i {
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
		file_coordinator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_coordinator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllPackages); i {
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
		file_coordinator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackagesResponce); i {
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
			RawDescriptor: file_coordinator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coordinator_proto_goTypes,
		DependencyIndexes: file_coordinator_proto_depIdxs,
		MessageInfos:      file_coordinator_proto_msgTypes,
	}.Build()
	File_coordinator_proto = out.File
	file_coordinator_proto_rawDesc = nil
	file_coordinator_proto_goTypes = nil
	file_coordinator_proto_depIdxs = nil
}
