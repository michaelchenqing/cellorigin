// Code generated by protoc-gen-go.
// source: model.proto
// DO NOT EDIT!

package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TestAccountData struct {
	Name string `protobuf:"bytes,10,opt,name=Name,json=name" json:"Name,omitempty"`
}

func (m *TestAccountData) Reset()                    { *m = TestAccountData{} }
func (m *TestAccountData) String() string            { return proto.CompactTextString(m) }
func (*TestAccountData) ProtoMessage()               {}
func (*TestAccountData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type TestEquipData struct {
	ModelKey    int64  `protobuf:"varint,1,opt,name=ModelKey,json=modelKey" json:"ModelKey,omitempty"`
	ModelDelete bool   `protobuf:"varint,2,opt,name=ModelDelete,json=modelDelete" json:"ModelDelete,omitempty"`
	Name        string `protobuf:"bytes,10,opt,name=Name,json=name" json:"Name,omitempty"`
}

func (m *TestEquipData) Reset()                    { *m = TestEquipData{} }
func (m *TestEquipData) String() string            { return proto.CompactTextString(m) }
func (*TestEquipData) ProtoMessage()               {}
func (*TestEquipData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type TestRoleData struct {
	ModelKey    int64   `protobuf:"varint,1,opt,name=ModelKey,json=modelKey" json:"ModelKey,omitempty"`
	ModelDelete bool    `protobuf:"varint,2,opt,name=ModelDelete,json=modelDelete" json:"ModelDelete,omitempty"`
	HP          int32   `protobuf:"varint,10,opt,name=HP,json=hP" json:"HP,omitempty"`
	EquipList   []int64 `protobuf:"varint,20,rep,name=EquipList,json=equipList" json:"EquipList,omitempty"`
}

func (m *TestRoleData) Reset()                    { *m = TestRoleData{} }
func (m *TestRoleData) String() string            { return proto.CompactTextString(m) }
func (*TestRoleData) ProtoMessage()               {}
func (*TestRoleData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type ModelACK struct {
	Account *TestAccountData `protobuf:"bytes,10,opt,name=Account,json=account" json:"Account,omitempty"`
	Role    []*TestRoleData  `protobuf:"bytes,20,rep,name=Role,json=role" json:"Role,omitempty"`
	Equip   []*TestEquipData `protobuf:"bytes,30,rep,name=Equip,json=equip" json:"Equip,omitempty"`
}

func (m *ModelACK) Reset()                    { *m = ModelACK{} }
func (m *ModelACK) String() string            { return proto.CompactTextString(m) }
func (*ModelACK) ProtoMessage()               {}
func (*ModelACK) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *ModelACK) GetAccount() *TestAccountData {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *ModelACK) GetRole() []*TestRoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *ModelACK) GetEquip() []*TestEquipData {
	if m != nil {
		return m.Equip
	}
	return nil
}

type TestModelCommandACK struct {
	Command string `protobuf:"bytes,1,opt,name=Command,json=command" json:"Command,omitempty"`
}

func (m *TestModelCommandACK) Reset()                    { *m = TestModelCommandACK{} }
func (m *TestModelCommandACK) String() string            { return proto.CompactTextString(m) }
func (*TestModelCommandACK) ProtoMessage()               {}
func (*TestModelCommandACK) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func init() {
	proto.RegisterType((*TestAccountData)(nil), "gamedef.TestAccountData")
	proto.RegisterType((*TestEquipData)(nil), "gamedef.TestEquipData")
	proto.RegisterType((*TestRoleData)(nil), "gamedef.TestRoleData")
	proto.RegisterType((*ModelACK)(nil), "gamedef.ModelACK")
	proto.RegisterType((*TestModelCommandACK)(nil), "gamedef.TestModelCommandACK")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x51, 0xed, 0x4a, 0xc3, 0x40,
	0x10, 0x24, 0x69, 0x62, 0x92, 0x8d, 0x1f, 0x70, 0x7e, 0x70, 0x88, 0x48, 0x08, 0x08, 0x0a, 0x12,
	0x21, 0x3e, 0x41, 0x69, 0x05, 0xa1, 0x2a, 0xe5, 0xf0, 0x05, 0xce, 0x64, 0xd5, 0x42, 0x92, 0xab,
	0xed, 0xf5, 0x87, 0xbe, 0x88, 0xaf, 0xeb, 0xdd, 0xe6, 0x2a, 0x54, 0xfc, 0xd9, 0x7f, 0x37, 0xbb,
	0xb3, 0x33, 0x93, 0x09, 0xa4, 0xad, 0xaa, 0xb1, 0x29, 0xe6, 0x0b, 0xa5, 0x15, 0x8b, 0xde, 0x64,
	0x8b, 0x35, 0xbe, 0xe6, 0x17, 0x70, 0xf0, 0x8c, 0x4b, 0x3d, 0xac, 0x2a, 0xb5, 0xea, 0xf4, 0x58,
	0x6a, 0xc9, 0x18, 0x04, 0x4f, 0x66, 0xcb, 0x21, 0xf3, 0x2e, 0x13, 0x11, 0x74, 0xe6, 0x9d, 0x4b,
	0xd8, 0xb3, 0xb4, 0xbb, 0x8f, 0xd5, 0x6c, 0x4e, 0xa4, 0x53, 0x88, 0x1f, 0xad, 0xde, 0x04, 0x3f,
	0xb9, 0x67, 0x88, 0x03, 0x11, 0xb7, 0x0e, 0xb3, 0x0c, 0x52, 0xda, 0x8d, 0xb1, 0x41, 0x8d, 0xdc,
	0x37, 0xeb, 0x58, 0xf4, 0xf6, 0xfd, 0xe8, 0x5f, 0x8b, 0x2f, 0xd8, 0xb5, 0x16, 0x42, 0x35, 0xb8,
	0x05, 0x87, 0x7d, 0xf0, 0xef, 0xa7, 0xa4, 0x1f, 0x0a, 0xff, 0x7d, 0xca, 0xce, 0x20, 0xa1, 0xf0,
	0x0f, 0xb3, 0xa5, 0xe6, 0x47, 0xd9, 0xc0, 0xc8, 0x25, 0xb8, 0x1e, 0xe4, 0xdf, 0x9e, 0x33, 0x1b,
	0x8e, 0x26, 0xac, 0x84, 0xc8, 0xd5, 0x41, 0xf7, 0x69, 0xc9, 0x0b, 0xd7, 0x56, 0xf1, 0xa7, 0x2a,
	0x11, 0xc9, 0x1e, 0xb0, 0x2b, 0x08, 0x6c, 0x70, 0x52, 0x4e, 0xcb, 0xe3, 0x8d, 0x83, 0xf5, 0x17,
	0x89, 0x60, 0x61, 0x5e, 0xec, 0x1a, 0x42, 0x4a, 0xc2, 0xcf, 0x89, 0x7b, 0xb2, 0xc1, 0xfd, 0x2d,
	0x58, 0x84, 0x94, 0x2e, 0xbf, 0x81, 0x43, 0x3b, 0xa7, 0x70, 0x23, 0xd5, 0xb6, 0xb2, 0xab, 0x6d,
	0x46, 0x0e, 0x91, 0x43, 0xd4, 0x4d, 0x22, 0xa2, 0xaa, 0x87, 0x2f, 0x3b, 0xf4, 0x83, 0x6f, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x93, 0xdf, 0xdc, 0xef, 0x01, 0x00, 0x00,
}
