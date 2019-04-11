// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/storageregistry/v0alpha/resources.proto

package storageregistryv0alphapb

import (
	fmt "fmt"
	types "github.com/cernbox/go-cs3apis/cs3/types"
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

type ProviderInfo struct {
	// OPTIONAL.
	// Opaque information (containing storage-specific information).
	// For example, additional metadata attached to the resource.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The storage provider id that will become part of the
	// resource id.
	// For example, if the provider_id is "home", resources obtained
	// from this storage provider will have a resource id like "home:1234".
	ProviderId string `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	// REQUIRED.
	// The public path prefix this storage provider handles.
	// In Unix literature, the mount path.
	// For example, if the provider_path is "/home", resources obtained
	// from this storage provier will have a resource path lik "/home/docs".
	ProviderPath string `protobuf:"bytes,3,opt,name=provider_path,json=providerPath,proto3" json:"provider_path,omitempty"`
	// REQUIRED.
	// The address where the storage provider can be reached.
	// For example, tcp://localhost:1099.
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// OPTIONAL.
	// Information to describe the functionalities
	// offered by the storage provider. Meant to be read
	// by humans.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// REQUIRED.
	// List of available methods.
	Features             *ProviderInfo_Features `protobuf:"bytes,6,opt,name=features,proto3" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ProviderInfo) Reset()         { *m = ProviderInfo{} }
func (m *ProviderInfo) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo) ProtoMessage()    {}
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90ce9d99d00c316, []int{0}
}

func (m *ProviderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderInfo.Unmarshal(m, b)
}
func (m *ProviderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderInfo.Marshal(b, m, deterministic)
}
func (m *ProviderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderInfo.Merge(m, src)
}
func (m *ProviderInfo) XXX_Size() int {
	return xxx_messageInfo_ProviderInfo.Size(m)
}
func (m *ProviderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderInfo proto.InternalMessageInfo

func (m *ProviderInfo) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *ProviderInfo) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *ProviderInfo) GetProviderPath() string {
	if m != nil {
		return m.ProviderPath
	}
	return ""
}

func (m *ProviderInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProviderInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProviderInfo) GetFeatures() *ProviderInfo_Features {
	if m != nil {
		return m.Features
	}
	return nil
}

// REQUIRED.
// Represents the list of features available
// on this storage provider. If the feature is not supported,
// the related service methods MUST return CODE_UNIMPLEMENTED.
type ProviderInfo_Features struct {
	Recycle              bool     `protobuf:"varint,1,opt,name=recycle,proto3" json:"recycle,omitempty"`
	FileVersions         bool     `protobuf:"varint,2,opt,name=file_versions,json=fileVersions,proto3" json:"file_versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderInfo_Features) Reset()         { *m = ProviderInfo_Features{} }
func (m *ProviderInfo_Features) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo_Features) ProtoMessage()    {}
func (*ProviderInfo_Features) Descriptor() ([]byte, []int) {
	return fileDescriptor_d90ce9d99d00c316, []int{0, 0}
}

func (m *ProviderInfo_Features) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderInfo_Features.Unmarshal(m, b)
}
func (m *ProviderInfo_Features) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderInfo_Features.Marshal(b, m, deterministic)
}
func (m *ProviderInfo_Features) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderInfo_Features.Merge(m, src)
}
func (m *ProviderInfo_Features) XXX_Size() int {
	return xxx_messageInfo_ProviderInfo_Features.Size(m)
}
func (m *ProviderInfo_Features) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderInfo_Features.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderInfo_Features proto.InternalMessageInfo

func (m *ProviderInfo_Features) GetRecycle() bool {
	if m != nil {
		return m.Recycle
	}
	return false
}

func (m *ProviderInfo_Features) GetFileVersions() bool {
	if m != nil {
		return m.FileVersions
	}
	return false
}

func init() {
	proto.RegisterType((*ProviderInfo)(nil), "cs3.storageregistryv0alpha.ProviderInfo")
	proto.RegisterType((*ProviderInfo_Features)(nil), "cs3.storageregistryv0alpha.ProviderInfo.Features")
}

func init() {
	proto.RegisterFile("cs3/storageregistry/v0alpha/resources.proto", fileDescriptor_d90ce9d99d00c316)
}

var fileDescriptor_d90ce9d99d00c316 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4e, 0xc2, 0x30,
	0x1c, 0xc7, 0xb3, 0xa9, 0x38, 0x0b, 0x1e, 0x9c, 0x9a, 0x34, 0x3b, 0x28, 0xd1, 0x0b, 0xc6, 0x64,
	0xa0, 0x7b, 0x02, 0x20, 0x48, 0x76, 0x30, 0x2c, 0x1d, 0x21, 0x6a, 0x4c, 0x48, 0xd9, 0x7e, 0xc0,
	0x12, 0xa4, 0xb5, 0x2d, 0x24, 0xbc, 0x0e, 0x47, 0x1f, 0xc5, 0xbb, 0xef, 0x63, 0xd6, 0x6d, 0x04,
	0x0d, 0x7a, 0x59, 0xd2, 0xcf, 0xf7, 0x4f, 0x7e, 0xbf, 0xae, 0xe8, 0x36, 0x92, 0x5e, 0x5d, 0x2a,
	0x26, 0xe8, 0x04, 0x04, 0x4c, 0x12, 0xa9, 0xc4, 0xaa, 0xbe, 0x6c, 0xd0, 0x19, 0x9f, 0xd2, 0xba,
	0x00, 0xc9, 0x16, 0x22, 0x02, 0xe9, 0x72, 0xc1, 0x14, 0xb3, 0x9d, 0x48, 0x7a, 0xee, 0x2f, 0x73,
	0xee, 0x75, 0xce, 0xd3, 0x22, 0xb5, 0xe2, 0x20, 0xb3, 0x6f, 0x16, 0xb9, 0xfa, 0x32, 0x51, 0x25,
	0x10, 0x6c, 0x99, 0xc4, 0x20, 0xfc, 0xf9, 0x98, 0xd9, 0x37, 0xa8, 0xc4, 0x38, 0x7d, 0x5f, 0x00,
	0x36, 0xaa, 0x46, 0xad, 0x7c, 0x7f, 0xe2, 0xa6, 0xa5, 0x59, 0xa4, 0xa7, 0x05, 0x92, 0x1b, 0xec,
	0x4b, 0x54, 0xe6, 0x79, 0x74, 0x98, 0xc4, 0xd8, 0xac, 0x1a, 0xb5, 0x23, 0x82, 0x0a, 0xe4, 0xc7,
	0xf6, 0x35, 0x3a, 0xde, 0x18, 0x38, 0x55, 0x53, 0xbc, 0xa7, 0x2d, 0x95, 0x02, 0x06, 0x54, 0x4d,
	0x6d, 0x8c, 0x0e, 0x69, 0x1c, 0x0b, 0x90, 0x12, 0xef, 0x6b, 0xb9, 0x38, 0xda, 0x55, 0x54, 0x8e,
	0x41, 0x46, 0x22, 0xe1, 0x2a, 0x61, 0x73, 0x7c, 0xa0, 0xd5, 0x6d, 0x64, 0x3f, 0x22, 0x6b, 0x0c,
	0x54, 0x2d, 0x04, 0x48, 0x5c, 0xd2, 0xe3, 0xde, 0xb9, 0x7f, 0xdf, 0x81, 0xbb, 0xbd, 0xa8, 0xfb,
	0x90, 0x07, 0xc9, 0xa6, 0xc2, 0xf1, 0x91, 0x55, 0xd0, 0x74, 0x2c, 0x01, 0xd1, 0x2a, 0x9a, 0x65,
	0x17, 0x61, 0x91, 0xe2, 0x98, 0x6e, 0x35, 0x4e, 0x66, 0x30, 0x5c, 0x82, 0x90, 0x09, 0x9b, 0x4b,
	0xbd, 0xb8, 0x45, 0x2a, 0x29, 0x1c, 0xe4, 0xac, 0xb5, 0x36, 0xd0, 0x45, 0xc4, 0xde, 0xfe, 0x99,
	0xa6, 0x75, 0x16, 0xfe, 0xe4, 0x41, 0xfa, 0x43, 0x02, 0xe3, 0x05, 0xef, 0xf6, 0xf3, 0xd1, 0xda,
	0x3c, 0x6d, 0xb7, 0x7a, 0x4f, 0x61, 0xbf, 0x47, 0x9a, 0xdd, 0x0e, 0xe9, 0x74, 0xfd, 0xb0, 0x4f,
	0x9e, 0x3f, 0x4c, 0xa7, 0x1d, 0x7a, 0x6e, 0x5e, 0x46, 0xf2, 0xd0, 0xa0, 0xd1, 0x4c, 0x43, 0x9f,
	0x5a, 0x7c, 0xdd, 0x2d, 0x8e, 0x4a, 0xfa, 0x0d, 0x78, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0x6c, 0x10, 0x8d, 0x65, 0x02, 0x00, 0x00,
}