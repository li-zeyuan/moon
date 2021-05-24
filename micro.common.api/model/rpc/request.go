package rpc

import "google.golang.org/protobuf/runtime/protoimpl"

type UpsertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}
