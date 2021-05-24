package rpc

import "google.golang.org/protobuf/runtime/protoimpl"

type UpsertResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DmError  int32  `protobuf:"varint,1,opt,name=dm_error,json=dmError,proto3" json:"dm_error,omitempty"`
	ErrorMsg string `protobuf:"bytes,2,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
}
