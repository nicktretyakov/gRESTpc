// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: bank.proto

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

type Bank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string                 `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	AvailableSlot int64                  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
}

func (x *Bank) Reset() {
	*x = Bank{}
	mi := &file_bank_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bank) ProtoMessage() {}

func (x *Bank) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bank.ProtoReflect.Descriptor instead.
func (*Bank) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{0}
}

func (x *Bank) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bank) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bank) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Bank) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Bank) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Bank) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Bank) GetAvailableSlot() int64 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

type ListBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From          string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int64                  `protobuf:"varint,5,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
	Name          string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListBankRequest) Reset() {
	*x = ListBankRequest{}
	mi := &file_bank_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBankRequest) ProtoMessage() {}

func (x *ListBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBankRequest.ProtoReflect.Descriptor instead.
func (*ListBankRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{1}
}

func (x *ListBankRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ListBankRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ListBankRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ListBankRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ListBankRequest) GetAvailableSlot() int64 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

func (x *ListBankRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListBankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banks []*Bank `protobuf:"bytes,1,rep,name=banks,proto3" json:"banks,omitempty"`
}

func (x *ListBankResponse) Reset() {
	*x = ListBankResponse{}
	mi := &file_bank_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBankResponse) ProtoMessage() {}

func (x *ListBankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBankResponse.ProtoReflect.Descriptor instead.
func (*ListBankResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{2}
}

func (x *ListBankResponse) GetBanks() []*Bank {
	if x != nil {
		return x.Banks
	}
	return nil
}

var File_bank_proto protoreflect.FileDescriptor

var file_bank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x38, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x05, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x32, 0xa0, 0x01, 0x0a,
	0x0b, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x1a, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x28, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x6e, 0x6b,
	0x1a, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x6e, 0x6b,
	0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_proto_rawDescOnce sync.Once
	file_bank_proto_rawDescData = file_bank_proto_rawDesc
)

func file_bank_proto_rawDescGZIP() []byte {
	file_bank_proto_rawDescOnce.Do(func() {
		file_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_proto_rawDescData)
	})
	return file_bank_proto_rawDescData
}

var file_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bank_proto_goTypes = []any{
	(*Bank)(nil),                  // 0: training.Bank
	(*ListBankRequest)(nil),       // 1: training.ListBankRequest
	(*ListBankResponse)(nil),      // 2: training.ListBankResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_bank_proto_depIdxs = []int32{
	3, // 0: training.Bank.date:type_name -> google.protobuf.Timestamp
	3, // 1: training.ListBankRequest.date:type_name -> google.protobuf.Timestamp
	0, // 2: training.ListBankResponse.banks:type_name -> training.Bank
	0, // 3: training.BankService.Create:input_type -> training.Bank
	0, // 4: training.BankService.Update:input_type -> training.Bank
	1, // 5: training.BankService.List:input_type -> training.ListBankRequest
	0, // 6: training.BankService.Create:output_type -> training.Bank
	0, // 7: training.BankService.Update:output_type -> training.Bank
	2, // 8: training.BankService.List:output_type -> training.ListBankResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bank_proto_init() }
func file_bank_proto_init() {
	if File_bank_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_proto_goTypes,
		DependencyIndexes: file_bank_proto_depIdxs,
		MessageInfos:      file_bank_proto_msgTypes,
	}.Build()
	File_bank_proto = out.File
	file_bank_proto_rawDesc = nil
	file_bank_proto_goTypes = nil
	file_bank_proto_depIdxs = nil
}
