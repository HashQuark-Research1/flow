// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/node_version_info.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NodeVersionInfo struct {
	Semver               string   `protobuf:"bytes,1,opt,name=semver,proto3" json:"semver,omitempty"`
	Commit               []byte   `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	SporkId              []byte   `protobuf:"bytes,3,opt,name=spork_id,json=sporkId,proto3" json:"spork_id,omitempty"`
	ProtocolVersion      uint32   `protobuf:"varint,4,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeVersionInfo) Reset()         { *m = NodeVersionInfo{} }
func (m *NodeVersionInfo) String() string { return proto.CompactTextString(m) }
func (*NodeVersionInfo) ProtoMessage()    {}
func (*NodeVersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9e447351cb4a8a, []int{0}
}

func (m *NodeVersionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeVersionInfo.Unmarshal(m, b)
}
func (m *NodeVersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeVersionInfo.Marshal(b, m, deterministic)
}
func (m *NodeVersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeVersionInfo.Merge(m, src)
}
func (m *NodeVersionInfo) XXX_Size() int {
	return xxx_messageInfo_NodeVersionInfo.Size(m)
}
func (m *NodeVersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeVersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeVersionInfo proto.InternalMessageInfo

func (m *NodeVersionInfo) GetSemver() string {
	if m != nil {
		return m.Semver
	}
	return ""
}

func (m *NodeVersionInfo) GetCommit() []byte {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *NodeVersionInfo) GetSporkId() []byte {
	if m != nil {
		return m.SporkId
	}
	return nil
}

func (m *NodeVersionInfo) GetProtocolVersion() uint32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeVersionInfo)(nil), "flow.entities.NodeVersionInfo")
}

func init() {
	proto.RegisterFile("flow/entities/node_version_info.proto", fileDescriptor_ff9e447351cb4a8a)
}

var fileDescriptor_ff9e447351cb4a8a = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xcb, 0xc9, 0x2f,
	0xd7, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0xcf, 0xcb, 0x4f, 0x49, 0x8d, 0x2f,
	0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x8b, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x05, 0x29, 0xd3, 0x83, 0x29, 0x53, 0x6a, 0x67, 0xe4, 0xe2, 0xf7, 0xcb, 0x4f,
	0x49, 0x0d, 0x83, 0xa8, 0xf4, 0xcc, 0x4b, 0xcb, 0x17, 0x12, 0xe3, 0x62, 0x2b, 0x4e, 0xcd, 0x2d,
	0x4b, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x40, 0xe2, 0xc9, 0xf9, 0xb9,
	0xb9, 0x99, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x50, 0x9e, 0x90, 0x24, 0x17, 0x47,
	0x71, 0x41, 0x7e, 0x51, 0x76, 0x7c, 0x66, 0x8a, 0x04, 0x33, 0x58, 0x86, 0x1d, 0xcc, 0xf7, 0x4c,
	0x11, 0xd2, 0xe4, 0x12, 0x00, 0x5b, 0x9b, 0x9c, 0x9f, 0x03, 0x73, 0x8c, 0x04, 0x8b, 0x02, 0xa3,
	0x06, 0x6f, 0x10, 0x3f, 0x4c, 0x1c, 0x6a, 0xb3, 0x53, 0x00, 0x97, 0x4c, 0x7e, 0x51, 0xba, 0x5e,
	0x7e, 0x1e, 0xd8, 0x81, 0x60, 0xd9, 0xa4, 0xd2, 0x34, 0xb8, 0x4b, 0xa3, 0x0c, 0xd2, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x21, 0x8a, 0xf4, 0xc1, 0x04, 0x4c, 0xa5, 0x7e,
	0x7a, 0xbe, 0x3e, 0x4a, 0x10, 0x24, 0xb1, 0x81, 0xa5, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x97, 0x38, 0x86, 0xbc, 0x1a, 0x01, 0x00, 0x00,
}
