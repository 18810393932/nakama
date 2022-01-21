/**
* Created by GoLand
* @file api.pb.go
* @version: 1.0.0
* @author 李锦 <lijin@shihuituan.com>
* @date 2022/1/21 4:22 下午
* @desc api.pb.go
 */

package auth

import (
	"github.com/heroiclabs/nakama-common/api"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 116)

// AccountOculus Send a oculus ID to the server. Used with authenticate/link/unlink.
type AccountOculus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A Oculus identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Extra information that will be bundled in the session token.
	Vars map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AccountOculus) Reset() {
	*x = AccountOculus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountOculus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountOculus) ProtoMessage() {}

func (x *AccountOculus) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountOculus.ProtoReflect.Descriptor instead.
func (*AccountOculus) Descriptor() ([]byte, []int) {
	b, _ := new(api.AuthenticateSteamRequest).Descriptor()
	return b, []int{100}
}

func (x *AccountOculus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountOculus) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

//===============================================================================================

// AuthenticateOculusRequest Authenticate against the server with a oculus ID.
type AuthenticateOculusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The oculus account details.
	Account *AccountOculus `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Register the account if the user does not already exist.
	Create *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	// Set the username on the account at register. Must be unique.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AuthenticateOculusRequest) Reset() {
	*x = AuthenticateOculusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateOculusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateOculusRequest) ProtoMessage() {}

func (x *AuthenticateOculusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateOculusRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateOculusRequest) Descriptor() ([]byte, []int) {
	b, _ := new(api.AuthenticateSteamRequest).Descriptor()
	return b, []int{101}
}

func (x *AuthenticateOculusRequest) GetAccount() *AccountOculus {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *AuthenticateOculusRequest) GetCreate() *wrapperspb.BoolValue {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *AuthenticateOculusRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}
