// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	client.proto

It has these top-level messages:
	StoreRequest
	StoreResponse
	RetrieveRequest
	RetrieveResponse
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StoreRequest struct {
	ID   string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *StoreRequest) Reset()                    { *m = StoreRequest{} }
func (m *StoreRequest) String() string            { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()               {}
func (*StoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StoreRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *StoreRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type StoreResponse struct {
	ID  string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *StoreResponse) Reset()                    { *m = StoreResponse{} }
func (m *StoreResponse) String() string            { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()               {}
func (*StoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StoreResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *StoreResponse) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type RetrieveRequest struct {
	ID  string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *RetrieveRequest) Reset()                    { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()               {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RetrieveRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *RetrieveRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type RetrieveResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RetrieveResponse) Reset()                    { *m = RetrieveResponse{} }
func (m *RetrieveResponse) String() string            { return proto.CompactTextString(m) }
func (*RetrieveResponse) ProtoMessage()               {}
func (*RetrieveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RetrieveResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*StoreRequest)(nil), "protobuf.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "protobuf.StoreResponse")
	proto.RegisterType((*RetrieveRequest)(nil), "protobuf.RetrieveRequest")
	proto.RegisterType((*RetrieveResponse)(nil), "protobuf.RetrieveResponse")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a,
	0x46, 0x5c, 0x3c, 0xc1, 0x25, 0xf9, 0x45, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x7c, 0x5c, 0x4c, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x9e, 0x2e, 0x42,
	0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6,
	0x92, 0x21, 0x17, 0x2f, 0x54, 0x4f, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x86, 0x26, 0x01, 0x2e,
	0xe6, 0xec, 0xd4, 0x4a, 0xa8, 0x1e, 0x10, 0x53, 0xc9, 0x98, 0x8b, 0x3f, 0x28, 0xb5, 0xa4, 0x28,
	0x33, 0xb5, 0x0c, 0xa7, 0x4d, 0x98, 0x9a, 0xd4, 0xb8, 0x04, 0x10, 0x9a, 0xa0, 0x56, 0xc1, 0xdc,
	0xc3, 0x88, 0x70, 0x4f, 0x12, 0x1b, 0xd8, 0x37, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e,
	0xbb, 0xb7, 0x68, 0xe4, 0x00, 0x00, 0x00,
}