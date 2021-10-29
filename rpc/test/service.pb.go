// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: rpc/test/service.proto

package test

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

type GetRpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdInRequest string `protobuf:"bytes,1,opt,name=id_in_request,json=idInRequest,proto3" json:"id_in_request,omitempty"`
}

func (x *GetRpcRequest) Reset() {
	*x = GetRpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_test_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRpcRequest) ProtoMessage() {}

func (x *GetRpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_test_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRpcRequest.ProtoReflect.Descriptor instead.
func (*GetRpcRequest) Descriptor() ([]byte, []int) {
	return file_rpc_test_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRpcRequest) GetIdInRequest() string {
	if x != nil {
		return x.IdInRequest
	}
	return ""
}

type GetRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GetRpcResponse) Reset() {
	*x = GetRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_test_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRpcResponse) ProtoMessage() {}

func (x *GetRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_test_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRpcResponse.ProtoReflect.Descriptor instead.
func (*GetRpcResponse) Descriptor() ([]byte, []int) {
	return file_rpc_test_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRpcResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type PostRpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestData_1 string `protobuf:"bytes,1,opt,name=request_data_1,json=requestData1,proto3" json:"request_data_1,omitempty"`
	RequestData_2 string `protobuf:"bytes,2,opt,name=request_data_2,json=requestData2,proto3" json:"request_data_2,omitempty"`
}

func (x *PostRpcRequest) Reset() {
	*x = PostRpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_test_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRpcRequest) ProtoMessage() {}

func (x *PostRpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_test_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRpcRequest.ProtoReflect.Descriptor instead.
func (*PostRpcRequest) Descriptor() ([]byte, []int) {
	return file_rpc_test_service_proto_rawDescGZIP(), []int{2}
}

func (x *PostRpcRequest) GetRequestData_1() string {
	if x != nil {
		return x.RequestData_1
	}
	return ""
}

func (x *PostRpcRequest) GetRequestData_2() string {
	if x != nil {
		return x.RequestData_2
	}
	return ""
}

type PostRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PostRpcResponse) Reset() {
	*x = PostRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_test_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRpcResponse) ProtoMessage() {}

func (x *PostRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_test_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRpcResponse.ProtoReflect.Descriptor instead.
func (*PostRpcResponse) Descriptor() ([]byte, []int) {
	return file_rpc_test_service_proto_rawDescGZIP(), []int{3}
}

func (x *PostRpcResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_rpc_test_service_proto protoreflect.FileDescriptor

var file_rpc_test_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x22, 0x33,
	0x0a, 0x0d, 0x67, 0x65, 0x74, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0d, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x5c, 0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x31, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x32, 0x22,
	0x2d, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x73,
	0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x52, 0x70, 0x63,
	0x12, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x52, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x65, 0x74,
	0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x70,
	0x6f, 0x73, 0x74, 0x52, 0x70, 0x63, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x3b,
	0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_test_service_proto_rawDescOnce sync.Once
	file_rpc_test_service_proto_rawDescData = file_rpc_test_service_proto_rawDesc
)

func file_rpc_test_service_proto_rawDescGZIP() []byte {
	file_rpc_test_service_proto_rawDescOnce.Do(func() {
		file_rpc_test_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_test_service_proto_rawDescData)
	})
	return file_rpc_test_service_proto_rawDescData
}

var file_rpc_test_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_test_service_proto_goTypes = []interface{}{
	(*GetRpcRequest)(nil),   // 0: test.getRpcRequest
	(*GetRpcResponse)(nil),  // 1: test.getRpcResponse
	(*PostRpcRequest)(nil),  // 2: test.postRpcRequest
	(*PostRpcResponse)(nil), // 3: test.postRpcResponse
}
var file_rpc_test_service_proto_depIdxs = []int32{
	0, // 0: test.test.getRpc:input_type -> test.getRpcRequest
	2, // 1: test.test.postRpc:input_type -> test.postRpcRequest
	1, // 2: test.test.getRpc:output_type -> test.getRpcResponse
	3, // 3: test.test.postRpc:output_type -> test.postRpcResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_test_service_proto_init() }
func file_rpc_test_service_proto_init() {
	if File_rpc_test_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_test_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRpcRequest); i {
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
		file_rpc_test_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRpcResponse); i {
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
		file_rpc_test_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRpcRequest); i {
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
		file_rpc_test_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRpcResponse); i {
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
			RawDescriptor: file_rpc_test_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_test_service_proto_goTypes,
		DependencyIndexes: file_rpc_test_service_proto_depIdxs,
		MessageInfos:      file_rpc_test_service_proto_msgTypes,
	}.Build()
	File_rpc_test_service_proto = out.File
	file_rpc_test_service_proto_rawDesc = nil
	file_rpc_test_service_proto_goTypes = nil
	file_rpc_test_service_proto_depIdxs = nil
}