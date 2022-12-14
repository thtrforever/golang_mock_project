// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: customer-service.proto

package pb

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

type GrpcPingCustomerGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GrpcPingCustomerGrpcRequest) Reset() {
	*x = GrpcPingCustomerGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcPingCustomerGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcPingCustomerGrpcRequest) ProtoMessage() {}

func (x *GrpcPingCustomerGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcPingCustomerGrpcRequest.ProtoReflect.Descriptor instead.
func (*GrpcPingCustomerGrpcRequest) Descriptor() ([]byte, []int) {
	return file_customer_service_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcPingCustomerGrpcRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GrpcPingCustomerGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GrpcPingCustomerGrpcResponse) Reset() {
	*x = GrpcPingCustomerGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcPingCustomerGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcPingCustomerGrpcResponse) ProtoMessage() {}

func (x *GrpcPingCustomerGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcPingCustomerGrpcResponse.ProtoReflect.Descriptor instead.
func (*GrpcPingCustomerGrpcResponse) Descriptor() ([]byte, []int) {
	return file_customer_service_proto_rawDescGZIP(), []int{1}
}

func (x *GrpcPingCustomerGrpcResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GrpcCreateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FirstName   string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *GrpcCreateCustomerRequest) Reset() {
	*x = GrpcCreateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcCreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcCreateCustomerRequest) ProtoMessage() {}

func (x *GrpcCreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcCreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*GrpcCreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_service_proto_rawDescGZIP(), []int{2}
}

func (x *GrpcCreateCustomerRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GrpcCreateCustomerRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GrpcCreateCustomerRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GrpcCreateCustomerRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GrpcCreateCustomerRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GrpcCreateCustomerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GrpcCreateCustomerRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type GrpcCreateCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Username string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Active   bool                   `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	CreateAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *GrpcCreateCustomerResponse) Reset() {
	*x = GrpcCreateCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcCreateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcCreateCustomerResponse) ProtoMessage() {}

func (x *GrpcCreateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcCreateCustomerResponse.ProtoReflect.Descriptor instead.
func (*GrpcCreateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_service_proto_rawDescGZIP(), []int{3}
}

func (x *GrpcCreateCustomerResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GrpcCreateCustomerResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GrpcCreateCustomerResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *GrpcCreateCustomerResponse) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

type GrpcUpdateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FirstName   string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *GrpcUpdateCustomerRequest) Reset() {
	*x = GrpcUpdateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcUpdateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcUpdateCustomerRequest) ProtoMessage() {}

func (x *GrpcUpdateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcUpdateCustomerRequest.ProtoReflect.Descriptor instead.
func (*GrpcUpdateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_service_proto_rawDescGZIP(), []int{4}
}

func (x *GrpcUpdateCustomerRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GrpcUpdateCustomerRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GrpcUpdateCustomerRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GrpcUpdateCustomerRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GrpcUpdateCustomerRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GrpcUpdateCustomerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GrpcUpdateCustomerRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type GrpcUpdateCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GrpcUpdateCustomerResponse) Reset() {
	*x = GrpcUpdateCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcUpdateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcUpdateCustomerResponse) ProtoMessage() {}

func (x *GrpcUpdateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcUpdateCustomerResponse.ProtoReflect.Descriptor instead.
func (*GrpcUpdateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_service_proto_rawDescGZIP(), []int{5}
}

func (x *GrpcUpdateCustomerResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_customer_service_proto protoreflect.FileDescriptor

var file_customer_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x1b, 0x47, 0x72, 0x70, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x36, 0x0a, 0x1c, 0x47, 0x72, 0x70, 0x63, 0x50, 0x69, 0x6e, 0x67,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xe2, 0x01, 0x0a,
	0x19, 0x47, 0x72, 0x70, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0xa1, 0x01, 0x0a, 0x1a, 0x47, 0x72, 0x70, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x37, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0xe2, 0x01, 0x0a, 0x19, 0x47, 0x72, 0x70, 0x63, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x1a, 0x47, 0x72,
	0x70, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x32, 0x8d, 0x02, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x72, 0x70,
	0x63, 0x12, 0x57, 0x0a, 0x10, 0x50, 0x69, 0x6e, 0x67, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x47, 0x72, 0x70, 0x63, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x50,
	0x69, 0x6e, 0x67, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x70, 0x63,
	0x50, 0x69, 0x6e, 0x67, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x17, 0x5a, 0x15, 0x74, 0x68, 0x69, 0x6e, 0x64, 0x76, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_customer_service_proto_rawDescOnce sync.Once
	file_customer_service_proto_rawDescData = file_customer_service_proto_rawDesc
)

func file_customer_service_proto_rawDescGZIP() []byte {
	file_customer_service_proto_rawDescOnce.Do(func() {
		file_customer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_service_proto_rawDescData)
	})
	return file_customer_service_proto_rawDescData
}

var file_customer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_customer_service_proto_goTypes = []interface{}{
	(*GrpcPingCustomerGrpcRequest)(nil),  // 0: pb.GrpcPingCustomerGrpcRequest
	(*GrpcPingCustomerGrpcResponse)(nil), // 1: pb.GrpcPingCustomerGrpcResponse
	(*GrpcCreateCustomerRequest)(nil),    // 2: pb.GrpcCreateCustomerRequest
	(*GrpcCreateCustomerResponse)(nil),   // 3: pb.GrpcCreateCustomerResponse
	(*GrpcUpdateCustomerRequest)(nil),    // 4: pb.GrpcUpdateCustomerRequest
	(*GrpcUpdateCustomerResponse)(nil),   // 5: pb.GrpcUpdateCustomerResponse
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
}
var file_customer_service_proto_depIdxs = []int32{
	6, // 0: pb.GrpcCreateCustomerResponse.create_at:type_name -> google.protobuf.Timestamp
	0, // 1: pb.CustomerGrpc.PingCustomerGrpc:input_type -> pb.GrpcPingCustomerGrpcRequest
	2, // 2: pb.CustomerGrpc.CreateCustomer:input_type -> pb.GrpcCreateCustomerRequest
	4, // 3: pb.CustomerGrpc.UpdateCustomer:input_type -> pb.GrpcUpdateCustomerRequest
	1, // 4: pb.CustomerGrpc.PingCustomerGrpc:output_type -> pb.GrpcPingCustomerGrpcResponse
	3, // 5: pb.CustomerGrpc.CreateCustomer:output_type -> pb.GrpcCreateCustomerResponse
	5, // 6: pb.CustomerGrpc.UpdateCustomer:output_type -> pb.GrpcUpdateCustomerResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customer_service_proto_init() }
func file_customer_service_proto_init() {
	if File_customer_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customer_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcPingCustomerGrpcRequest); i {
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
		file_customer_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcPingCustomerGrpcResponse); i {
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
		file_customer_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcCreateCustomerRequest); i {
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
		file_customer_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcCreateCustomerResponse); i {
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
		file_customer_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcUpdateCustomerRequest); i {
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
		file_customer_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcUpdateCustomerResponse); i {
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
			RawDescriptor: file_customer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_service_proto_goTypes,
		DependencyIndexes: file_customer_service_proto_depIdxs,
		MessageInfos:      file_customer_service_proto_msgTypes,
	}.Build()
	File_customer_service_proto = out.File
	file_customer_service_proto_rawDesc = nil
	file_customer_service_proto_goTypes = nil
	file_customer_service_proto_depIdxs = nil
}
