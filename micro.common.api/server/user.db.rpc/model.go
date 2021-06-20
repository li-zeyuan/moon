package userdbrpc

import (
	"github.com/li-zeyuan/micro/micro.common.api/model"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	ServerNameUserDbRpc = "user.db.rpc"
	AddressProfileServer = "localhost:7072"

	UrlProfileSave = "/profile.ProfileService/Save"
)

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Passport string `protobuf:"bytes,3,opt,name=passport,proto3" json:"passport,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (p *Profile) Reset() {

}

func (p *Profile) String() string{
	return ""
}

func (p *Profile) ProtoMessage() {

}

type SaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles []*Profile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
}


func (p *SaveReq) Reset() {

}

func (p *SaveReq) String() string{
	return ""
}

func (p *SaveReq) ProtoMessage() {

}


type SaveResp struct {
	model.BaseResponse
}