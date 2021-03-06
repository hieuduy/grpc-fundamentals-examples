// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: src/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CalculateSumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *CalculateSumRequest) Reset() {
	*x = CalculateSumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateSumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateSumRequest) ProtoMessage() {}

func (x *CalculateSumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculatorpb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateSumRequest.ProtoReflect.Descriptor instead.
func (*CalculateSumRequest) Descriptor() ([]byte, []int) {
	return file_src_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculateSumRequest) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *CalculateSumRequest) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type CalculateSumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *CalculateSumResponse) Reset() {
	*x = CalculateSumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateSumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateSumResponse) ProtoMessage() {}

func (x *CalculateSumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculatorpb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateSumResponse.ProtoReflect.Descriptor instead.
func (*CalculateSumResponse) Descriptor() ([]byte, []int) {
	return file_src_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateSumResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type DecomposePrimeFactorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *DecomposePrimeFactorRequest) Reset() {
	*x = DecomposePrimeFactorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculatorpb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecomposePrimeFactorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecomposePrimeFactorRequest) ProtoMessage() {}

func (x *DecomposePrimeFactorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculatorpb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecomposePrimeFactorRequest.ProtoReflect.Descriptor instead.
func (*DecomposePrimeFactorRequest) Descriptor() ([]byte, []int) {
	return file_src_calculatorpb_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *DecomposePrimeFactorRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type DecomposePrimeFactorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeFactor int64 `protobuf:"varint,1,opt,name=prime_factor,json=primeFactor,proto3" json:"prime_factor,omitempty"`
}

func (x *DecomposePrimeFactorResponse) Reset() {
	*x = DecomposePrimeFactorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculatorpb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecomposePrimeFactorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecomposePrimeFactorResponse) ProtoMessage() {}

func (x *DecomposePrimeFactorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculatorpb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecomposePrimeFactorResponse.ProtoReflect.Descriptor instead.
func (*DecomposePrimeFactorResponse) Descriptor() ([]byte, []int) {
	return file_src_calculatorpb_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *DecomposePrimeFactorResponse) GetPrimeFactor() int64 {
	if x != nil {
		return x.PrimeFactor
	}
	return 0
}

type CalculateSumEvenNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *CalculateSumEvenNumberRequest) Reset() {
	*x = CalculateSumEvenNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculatorpb_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateSumEvenNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateSumEvenNumberRequest) ProtoMessage() {}

func (x *CalculateSumEvenNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculatorpb_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateSumEvenNumberRequest.ProtoReflect.Descriptor instead.
func (*CalculateSumEvenNumberRequest) Descriptor() ([]byte, []int) {
	return file_src_calculatorpb_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *CalculateSumEvenNumberRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type CalculateSumEvenNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *CalculateSumEvenNumberResponse) Reset() {
	*x = CalculateSumEvenNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculatorpb_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateSumEvenNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateSumEvenNumberResponse) ProtoMessage() {}

func (x *CalculateSumEvenNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculatorpb_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateSumEvenNumberResponse.ProtoReflect.Descriptor instead.
func (*CalculateSumEvenNumberResponse) Descriptor() ([]byte, []int) {
	return file_src_calculatorpb_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *CalculateSumEvenNumberResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type FindMaxNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FindMaxNumberRequest) Reset() {
	*x = FindMaxNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculatorpb_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaxNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxNumberRequest) ProtoMessage() {}

func (x *FindMaxNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculatorpb_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxNumberRequest.ProtoReflect.Descriptor instead.
func (*FindMaxNumberRequest) Descriptor() ([]byte, []int) {
	return file_src_calculatorpb_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *FindMaxNumberRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaxNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxNumber int32 `protobuf:"varint,1,opt,name=max_number,json=maxNumber,proto3" json:"max_number,omitempty"`
}

func (x *FindMaxNumberResponse) Reset() {
	*x = FindMaxNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculatorpb_calculator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaxNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxNumberResponse) ProtoMessage() {}

func (x *FindMaxNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculatorpb_calculator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxNumberResponse.ProtoReflect.Descriptor instead.
func (*FindMaxNumberResponse) Descriptor() ([]byte, []int) {
	return file_src_calculatorpb_calculator_proto_rawDescGZIP(), []int{7}
}

func (x *FindMaxNumberResponse) GetMaxNumber() int32 {
	if x != nil {
		return x.MaxNumber
	}
	return 0
}

var File_src_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_src_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22,
	0x5d, 0x0a, 0x13, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x28,
	0x0a, 0x14, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x35, 0x0a, 0x1b, 0x44, 0x65, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x41, 0x0a, 0x1c, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x6d,
	0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0x37, 0x0a, 0x1d, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x1e, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22,
	0x2e, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x36, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61,
	0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0xa8, 0x03, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a,
	0x0c, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x12, 0x1f, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6d, 0x0a, 0x14, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x27, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x73, 0x0a, 0x16, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d,
	0x45, 0x76, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x5a, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61,
	0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_src_calculatorpb_calculator_proto_rawDescData = file_src_calculatorpb_calculator_proto_rawDesc
)

func file_src_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_src_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_src_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_calculatorpb_calculator_proto_rawDescData)
	})
	return file_src_calculatorpb_calculator_proto_rawDescData
}

var file_src_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_src_calculatorpb_calculator_proto_goTypes = []interface{}{
	(*CalculateSumRequest)(nil),            // 0: calculator.CalculateSumRequest
	(*CalculateSumResponse)(nil),           // 1: calculator.CalculateSumResponse
	(*DecomposePrimeFactorRequest)(nil),    // 2: calculator.DecomposePrimeFactorRequest
	(*DecomposePrimeFactorResponse)(nil),   // 3: calculator.DecomposePrimeFactorResponse
	(*CalculateSumEvenNumberRequest)(nil),  // 4: calculator.CalculateSumEvenNumberRequest
	(*CalculateSumEvenNumberResponse)(nil), // 5: calculator.CalculateSumEvenNumberResponse
	(*FindMaxNumberRequest)(nil),           // 6: calculator.FindMaxNumberRequest
	(*FindMaxNumberResponse)(nil),          // 7: calculator.FindMaxNumberResponse
}
var file_src_calculatorpb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorService.CalculateSum:input_type -> calculator.CalculateSumRequest
	2, // 1: calculator.CalculatorService.DecomposePrimeFactor:input_type -> calculator.DecomposePrimeFactorRequest
	4, // 2: calculator.CalculatorService.CalculateSumEvenNumber:input_type -> calculator.CalculateSumEvenNumberRequest
	6, // 3: calculator.CalculatorService.FindMaxNumber:input_type -> calculator.FindMaxNumberRequest
	1, // 4: calculator.CalculatorService.CalculateSum:output_type -> calculator.CalculateSumResponse
	3, // 5: calculator.CalculatorService.DecomposePrimeFactor:output_type -> calculator.DecomposePrimeFactorResponse
	5, // 6: calculator.CalculatorService.CalculateSumEvenNumber:output_type -> calculator.CalculateSumEvenNumberResponse
	7, // 7: calculator.CalculatorService.FindMaxNumber:output_type -> calculator.FindMaxNumberResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_calculatorpb_calculator_proto_init() }
func file_src_calculatorpb_calculator_proto_init() {
	if File_src_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateSumRequest); i {
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
		file_src_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateSumResponse); i {
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
		file_src_calculatorpb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecomposePrimeFactorRequest); i {
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
		file_src_calculatorpb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecomposePrimeFactorResponse); i {
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
		file_src_calculatorpb_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateSumEvenNumberRequest); i {
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
		file_src_calculatorpb_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateSumEvenNumberResponse); i {
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
		file_src_calculatorpb_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaxNumberRequest); i {
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
		file_src_calculatorpb_calculator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaxNumberResponse); i {
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
			RawDescriptor: file_src_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_src_calculatorpb_calculator_proto_depIdxs,
		MessageInfos:      file_src_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_src_calculatorpb_calculator_proto = out.File
	file_src_calculatorpb_calculator_proto_rawDesc = nil
	file_src_calculatorpb_calculator_proto_goTypes = nil
	file_src_calculatorpb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	CalculateSum(ctx context.Context, in *CalculateSumRequest, opts ...grpc.CallOption) (*CalculateSumResponse, error)
	// Server Streaming
	DecomposePrimeFactor(ctx context.Context, in *DecomposePrimeFactorRequest, opts ...grpc.CallOption) (CalculatorService_DecomposePrimeFactorClient, error)
	// Client Streaming
	CalculateSumEvenNumber(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_CalculateSumEvenNumberClient, error)
	// Bi-Directional Streaming
	FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaxNumberClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) CalculateSum(ctx context.Context, in *CalculateSumRequest, opts ...grpc.CallOption) (*CalculateSumResponse, error) {
	out := new(CalculateSumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/CalculateSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) DecomposePrimeFactor(ctx context.Context, in *DecomposePrimeFactorRequest, opts ...grpc.CallOption) (CalculatorService_DecomposePrimeFactorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/DecomposePrimeFactor", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceDecomposePrimeFactorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_DecomposePrimeFactorClient interface {
	Recv() (*DecomposePrimeFactorResponse, error)
	grpc.ClientStream
}

type calculatorServiceDecomposePrimeFactorClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceDecomposePrimeFactorClient) Recv() (*DecomposePrimeFactorResponse, error) {
	m := new(DecomposePrimeFactorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) CalculateSumEvenNumber(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_CalculateSumEvenNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/CalculateSumEvenNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceCalculateSumEvenNumberClient{stream}
	return x, nil
}

type CalculatorService_CalculateSumEvenNumberClient interface {
	Send(*CalculateSumEvenNumberRequest) error
	CloseAndRecv() (*CalculateSumEvenNumberResponse, error)
	grpc.ClientStream
}

type calculatorServiceCalculateSumEvenNumberClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceCalculateSumEvenNumberClient) Send(m *CalculateSumEvenNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceCalculateSumEvenNumberClient) CloseAndRecv() (*CalculateSumEvenNumberResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalculateSumEvenNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaxNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMaxNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaxNumberClient{stream}
	return x, nil
}

type CalculatorService_FindMaxNumberClient interface {
	Send(*FindMaxNumberRequest) error
	Recv() (*FindMaxNumberResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaxNumberClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaxNumberClient) Send(m *FindMaxNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaxNumberClient) Recv() (*FindMaxNumberResponse, error) {
	m := new(FindMaxNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	CalculateSum(context.Context, *CalculateSumRequest) (*CalculateSumResponse, error)
	// Server Streaming
	DecomposePrimeFactor(*DecomposePrimeFactorRequest, CalculatorService_DecomposePrimeFactorServer) error
	// Client Streaming
	CalculateSumEvenNumber(CalculatorService_CalculateSumEvenNumberServer) error
	// Bi-Directional Streaming
	FindMaxNumber(CalculatorService_FindMaxNumberServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) CalculateSum(context.Context, *CalculateSumRequest) (*CalculateSumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateSum not implemented")
}
func (*UnimplementedCalculatorServiceServer) DecomposePrimeFactor(*DecomposePrimeFactorRequest, CalculatorService_DecomposePrimeFactorServer) error {
	return status.Errorf(codes.Unimplemented, "method DecomposePrimeFactor not implemented")
}
func (*UnimplementedCalculatorServiceServer) CalculateSumEvenNumber(CalculatorService_CalculateSumEvenNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculateSumEvenNumber not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMaxNumber(CalculatorService_FindMaxNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaxNumber not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_CalculateSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateSumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).CalculateSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/CalculateSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).CalculateSum(ctx, req.(*CalculateSumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_DecomposePrimeFactor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DecomposePrimeFactorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).DecomposePrimeFactor(m, &calculatorServiceDecomposePrimeFactorServer{stream})
}

type CalculatorService_DecomposePrimeFactorServer interface {
	Send(*DecomposePrimeFactorResponse) error
	grpc.ServerStream
}

type calculatorServiceDecomposePrimeFactorServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceDecomposePrimeFactorServer) Send(m *DecomposePrimeFactorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_CalculateSumEvenNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).CalculateSumEvenNumber(&calculatorServiceCalculateSumEvenNumberServer{stream})
}

type CalculatorService_CalculateSumEvenNumberServer interface {
	SendAndClose(*CalculateSumEvenNumberResponse) error
	Recv() (*CalculateSumEvenNumberRequest, error)
	grpc.ServerStream
}

type calculatorServiceCalculateSumEvenNumberServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceCalculateSumEvenNumberServer) SendAndClose(m *CalculateSumEvenNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceCalculateSumEvenNumberServer) Recv() (*CalculateSumEvenNumberRequest, error) {
	m := new(CalculateSumEvenNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaxNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaxNumber(&calculatorServiceFindMaxNumberServer{stream})
}

type CalculatorService_FindMaxNumberServer interface {
	Send(*FindMaxNumberResponse) error
	Recv() (*FindMaxNumberRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaxNumberServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaxNumberServer) Send(m *FindMaxNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaxNumberServer) Recv() (*FindMaxNumberRequest, error) {
	m := new(FindMaxNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateSum",
			Handler:    _CalculatorService_CalculateSum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DecomposePrimeFactor",
			Handler:       _CalculatorService_DecomposePrimeFactor_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CalculateSumEvenNumber",
			Handler:       _CalculatorService_CalculateSumEvenNumber_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaxNumber",
			Handler:       _CalculatorService_FindMaxNumber_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/calculatorpb/calculator.proto",
}
