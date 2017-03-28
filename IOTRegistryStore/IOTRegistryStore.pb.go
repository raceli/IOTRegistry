// Code generated by protoc-gen-go.
// source: IOTRegistryStore.proto
// DO NOT EDIT!

/*
Package IOTRegistryStore is a generated protocol buffer package.

It is generated from these files:
	IOTRegistryStore.proto

It has these top-level messages:
	Identities
	Alias
	Things
	Spec
*/
package IOTRegistryStore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Identities struct {
	RegistrantName string `protobuf:"bytes,1,opt,name=RegistrantName" json:"RegistrantName,omitempty"`
	Pubkey         []byte `protobuf:"bytes,3,opt,name=Pubkey,proto3" json:"Pubkey,omitempty"`
}

func (m *Identities) Reset()         { *m = Identities{} }
func (m *Identities) String() string { return proto.CompactTextString(m) }
func (*Identities) ProtoMessage()    {}

type Alias struct {
	Nonce []byte `protobuf:"bytes,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}

type Things struct {
	Alias          []string `protobuf:"bytes,1,rep,name=Alias" json:"Alias,omitempty"`
	RegistrantName string   `protobuf:"bytes,2,opt,name=RegistrantName" json:"RegistrantName,omitempty"`
	Data           string   `protobuf:"bytes,3,opt,name=Data" json:"Data,omitempty"`
	SpecName       string   `protobuf:"bytes,4,opt,name=SpecName" json:"SpecName,omitempty"`
}

func (m *Things) Reset()         { *m = Things{} }
func (m *Things) String() string { return proto.CompactTextString(m) }
func (*Things) ProtoMessage()    {}

type Spec struct {
	RegistrantName string `protobuf:"bytes,2,opt,name=RegistrantName" json:"RegistrantName,omitempty"`
	Data           string `protobuf:"bytes,1,opt,name=Data" json:"Data,omitempty"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
