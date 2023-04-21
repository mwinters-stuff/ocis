// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/tx/v1beta1/resources.proto

package txv1beta1

import (
	fmt "fmt"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta13 "github.com/cs3org/go-cs3apis/cs3/sharing/ocm/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
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

// Status represents transfer status.
type Status int32

const (
	Status_STATUS_INVALID Status = 0
	// The destination could not be found.
	Status_STATUS_DESTINATION_NOT_FOUND Status = 1
	// A new data transfer
	Status_STATUS_TRANSFER_NEW Status = 2
	// The data transfer is awaiting acceptance from the destination
	Status_STATUS_TRANSFER_AWAITING_ACCEPTANCE Status = 3
	// The data transfer is accepted by the destination.
	Status_STATUS_TRANSFER_ACCEPTED Status = 4
	// The data transfer has started and not yet completed.
	Status_STATUS_TRANSFER_IN_PROGRESS Status = 5
	// The data transfer has completed.
	Status_STATUS_TRANSFER_COMPLETE Status = 6
	// The data transfer has failed.
	Status_STATUS_TRANSFER_FAILED Status = 7
	// The data transfer had been cancelled.
	Status_STATUS_TRANSFER_CANCELLED Status = 8
	// The request for cancelling the data transfer has failed.
	Status_STATUS_TRANSFER_CANCEL_FAILED Status = 9
	// The transfer has expired somewhere down the line.
	Status_STATUS_TRANSFER_EXPIRED Status = 10
)

var Status_name = map[int32]string{
	0:  "STATUS_INVALID",
	1:  "STATUS_DESTINATION_NOT_FOUND",
	2:  "STATUS_TRANSFER_NEW",
	3:  "STATUS_TRANSFER_AWAITING_ACCEPTANCE",
	4:  "STATUS_TRANSFER_ACCEPTED",
	5:  "STATUS_TRANSFER_IN_PROGRESS",
	6:  "STATUS_TRANSFER_COMPLETE",
	7:  "STATUS_TRANSFER_FAILED",
	8:  "STATUS_TRANSFER_CANCELLED",
	9:  "STATUS_TRANSFER_CANCEL_FAILED",
	10: "STATUS_TRANSFER_EXPIRED",
}

var Status_value = map[string]int32{
	"STATUS_INVALID":                      0,
	"STATUS_DESTINATION_NOT_FOUND":        1,
	"STATUS_TRANSFER_NEW":                 2,
	"STATUS_TRANSFER_AWAITING_ACCEPTANCE": 3,
	"STATUS_TRANSFER_ACCEPTED":            4,
	"STATUS_TRANSFER_IN_PROGRESS":         5,
	"STATUS_TRANSFER_COMPLETE":            6,
	"STATUS_TRANSFER_FAILED":              7,
	"STATUS_TRANSFER_CANCELLED":           8,
	"STATUS_TRANSFER_CANCEL_FAILED":       9,
	"STATUS_TRANSFER_EXPIRED":             10,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e6e9f5a73c9e5a84, []int{0}
}

// TxId uniquely identifies a transfer in the transfer provider namespace.
type TxId struct {
	// REQUIRED.
	// The internal transfer id used by the service implementor
	// to uniquely identity the transfer in the internal
	// implementation of the service.
	OpaqueId             string   `protobuf:"bytes,1,opt,name=opaque_id,json=opaqueId,proto3" json:"opaque_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxId) Reset()         { *m = TxId{} }
func (m *TxId) String() string { return proto.CompactTextString(m) }
func (*TxId) ProtoMessage()    {}
func (*TxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e9f5a73c9e5a84, []int{0}
}

func (m *TxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxId.Unmarshal(m, b)
}
func (m *TxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxId.Marshal(b, m, deterministic)
}
func (m *TxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxId.Merge(m, src)
}
func (m *TxId) XXX_Size() int {
	return xxx_messageInfo_TxId.Size(m)
}
func (m *TxId) XXX_DiscardUnknown() {
	xxx_messageInfo_TxId.DiscardUnknown(m)
}

var xxx_messageInfo_TxId proto.InternalMessageInfo

func (m *TxId) GetOpaqueId() string {
	if m != nil {
		return m.OpaqueId
	}
	return ""
}

// TxInfo represents information about a transfer.
type TxInfo struct {
	// REQUIRED.
	// The transfer identifier.
	Id *TxId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// REQUIRED.
	// The transfer status. Eg.: STATUS_TRANSFER_FAILED.
	// Note: the description field may provide additional information on the transfer status.
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=cs3.tx.v1beta1.Status" json:"status,omitempty"`
	// REQUIRED.
	// The destination (receiver of the transfer)
	Grantee *v1beta1.Grantee `protobuf:"bytes,3,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// REQUIRED.
	// Uniquely identifies a principal who initiates the transfer creation.
	Creator *v1beta11.UserId `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	// REQUIRED.
	// Creation time of the transfer.
	Ctime *v1beta12.Timestamp `protobuf:"bytes,5,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// OPTIONAL.
	// Information to describe the transfer status.
	// Eg. may contain information about a transfer failure.
	// Meant to be human-readable.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// REQUIRED.
	// Opaque unique identifier of the share on which the transfer is based.
	ShareId              *v1beta13.ShareId `protobuf:"bytes,7,opt,name=share_id,json=shareId,proto3" json:"share_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TxInfo) Reset()         { *m = TxInfo{} }
func (m *TxInfo) String() string { return proto.CompactTextString(m) }
func (*TxInfo) ProtoMessage()    {}
func (*TxInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6e9f5a73c9e5a84, []int{1}
}

func (m *TxInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxInfo.Unmarshal(m, b)
}
func (m *TxInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxInfo.Marshal(b, m, deterministic)
}
func (m *TxInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxInfo.Merge(m, src)
}
func (m *TxInfo) XXX_Size() int {
	return xxx_messageInfo_TxInfo.Size(m)
}
func (m *TxInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TxInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TxInfo proto.InternalMessageInfo

func (m *TxInfo) GetId() *TxId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TxInfo) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_INVALID
}

func (m *TxInfo) GetGrantee() *v1beta1.Grantee {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *TxInfo) GetCreator() *v1beta11.UserId {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *TxInfo) GetCtime() *v1beta12.Timestamp {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *TxInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TxInfo) GetShareId() *v1beta13.ShareId {
	if m != nil {
		return m.ShareId
	}
	return nil
}

func init() {
	proto.RegisterEnum("cs3.tx.v1beta1.Status", Status_name, Status_value)
	proto.RegisterType((*TxId)(nil), "cs3.tx.v1beta1.TxId")
	proto.RegisterType((*TxInfo)(nil), "cs3.tx.v1beta1.TxInfo")
}

func init() { proto.RegisterFile("cs3/tx/v1beta1/resources.proto", fileDescriptor_e6e9f5a73c9e5a84) }

var fileDescriptor_e6e9f5a73c9e5a84 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0x69, 0xb6, 0xb5, 0xeb, 0x37, 0xa9, 0x8a, 0x0c, 0xda, 0xc2, 0xfe, 0x40, 0xb7, 0x81,
	0x36, 0x10, 0x4a, 0xb5, 0xf5, 0x72, 0x17, 0x28, 0x6b, 0xbd, 0xc9, 0x52, 0x49, 0x2b, 0xc7, 0xfb,
	0x03, 0x9a, 0x14, 0x65, 0x89, 0x19, 0xb9, 0x68, 0x5d, 0x6c, 0x77, 0xca, 0x5e, 0x87, 0x4b, 0x1e,
	0x85, 0x27, 0xe0, 0x2d, 0x78, 0x05, 0x14, 0x27, 0x2d, 0xd0, 0x6d, 0xdc, 0xb5, 0xdf, 0xf9, 0x9d,
	0x23, 0x7f, 0xc7, 0x31, 0xbc, 0x88, 0x55, 0xbb, 0xa5, 0xb3, 0xd6, 0xed, 0xc1, 0x35, 0xd7, 0xd1,
	0x41, 0x4b, 0x72, 0x25, 0x26, 0x32, 0xe6, 0xca, 0x1d, 0x4b, 0xa1, 0x05, 0x6a, 0xc4, 0xaa, 0xed,
	0xea, 0xcc, 0x2d, 0xf5, 0xf5, 0x37, 0x39, 0x9f, 0x26, 0x7c, 0xa4, 0x53, 0x7d, 0xd7, 0x9a, 0x28,
	0x2e, 0x1f, 0xb3, 0xae, 0xef, 0xe5, 0xa8, 0xfa, 0x12, 0xc9, 0x74, 0x74, 0xd3, 0x12, 0xf1, 0xf0,
	0x51, 0xf0, 0x9d, 0x01, 0xb5, 0x90, 0xd1, 0x0d, 0x6f, 0x8d, 0xa5, 0xb8, 0x4d, 0x93, 0xff, 0xc4,
	0x6e, 0x99, 0x13, 0xdf, 0x8d, 0xb9, 0x9a, 0x21, 0xe6, 0x5f, 0x21, 0xef, 0xec, 0xc2, 0x22, 0xcb,
	0x48, 0x82, 0x36, 0xa0, 0x2e, 0xc6, 0xd1, 0xd7, 0x09, 0x0f, 0xd3, 0xc4, 0xa9, 0x34, 0x2b, 0xfb,
	0x75, 0xba, 0x5c, 0x0c, 0x48, 0xb2, 0xf3, 0xcb, 0x82, 0x2a, 0xcb, 0xc8, 0xe8, 0xb3, 0x40, 0xaf,
	0xc0, 0x2a, 0x81, 0x95, 0xc3, 0x67, 0xee, 0xbf, 0xdb, 0xba, 0x79, 0x12, 0xb5, 0xd2, 0x04, 0xb9,
	0x50, 0x55, 0x3a, 0xd2, 0x13, 0xe5, 0x58, 0xcd, 0xca, 0x7e, 0xe3, 0x70, 0x75, 0x9e, 0x0c, 0x8c,
	0x4a, 0x4b, 0x0a, 0xbd, 0x87, 0xda, 0x8d, 0x8c, 0x46, 0x9a, 0x73, 0x67, 0xc1, 0x44, 0xbf, 0x36,
	0x86, 0x72, 0x49, 0x77, 0xba, 0xe4, 0xcc, 0x7e, 0x5a, 0xc0, 0x74, 0xea, 0x42, 0x47, 0x50, 0x8b,
	0x25, 0x8f, 0xb4, 0x90, 0xce, 0xa2, 0x09, 0xd8, 0x36, 0x01, 0xd3, 0xe6, 0xdd, 0xbc, 0xf9, 0x99,
	0xfb, 0x4c, 0x71, 0x49, 0x12, 0x3a, 0x75, 0xa0, 0x43, 0x58, 0x8a, 0x75, 0x3a, 0xe4, 0xce, 0x92,
	0xb1, 0x6e, 0x16, 0x87, 0x35, 0x25, 0xcd, 0x36, 0x4b, 0x87, 0x5c, 0xe9, 0x68, 0x38, 0xa6, 0x05,
	0x8a, 0x9a, 0xb0, 0x92, 0x70, 0x15, 0xcb, 0x74, 0xac, 0x53, 0x31, 0x72, 0xaa, 0xa6, 0xb1, 0xbf,
	0x47, 0xe8, 0x08, 0x96, 0xf3, 0xdb, 0x34, 0x85, 0xd6, 0x4c, 0x70, 0xb3, 0x58, 0xaa, 0xb8, 0x62,
	0x57, 0xc4, 0xc3, 0x3f, 0x75, 0xe4, 0x60, 0x7e, 0x24, 0x55, 0xfc, 0x78, 0xfb, 0xd3, 0x82, 0x6a,
	0xd1, 0x11, 0x42, 0xd0, 0x08, 0x98, 0xc7, 0xce, 0x82, 0x90, 0xf8, 0xe7, 0x5e, 0x8f, 0x74, 0xed,
	0x27, 0xa8, 0x09, 0x9b, 0xe5, 0xac, 0x8b, 0x03, 0x46, 0x7c, 0x8f, 0x91, 0xbe, 0x1f, 0xfa, 0x7d,
	0x16, 0x9e, 0xf4, 0xcf, 0xfc, 0xae, 0x5d, 0x41, 0x6b, 0xf0, 0xb4, 0x24, 0x18, 0xf5, 0xfc, 0xe0,
	0x04, 0xd3, 0xd0, 0xc7, 0x17, 0xb6, 0x85, 0xf6, 0x60, 0x77, 0x5e, 0xf0, 0x2e, 0x3c, 0xc2, 0x88,
	0x7f, 0x1a, 0x7a, 0x9d, 0x0e, 0x1e, 0x30, 0xcf, 0xef, 0x60, 0x7b, 0x01, 0x6d, 0x82, 0x73, 0x0f,
	0x34, 0x3a, 0xee, 0xda, 0x8b, 0xe8, 0x25, 0x6c, 0xcc, 0xab, 0xc4, 0x0f, 0x07, 0xb4, 0x7f, 0x4a,
	0x71, 0x10, 0xd8, 0x4b, 0x0f, 0xd9, 0x3b, 0xfd, 0x0f, 0x83, 0x1e, 0x66, 0xd8, 0xae, 0xa2, 0x75,
	0x58, 0x9d, 0x57, 0x4f, 0x3c, 0xd2, 0xc3, 0x5d, 0xbb, 0x86, 0xb6, 0xe0, 0xf9, 0x3d, 0x67, 0x7e,
	0xa6, 0x5e, 0x2e, 0x2f, 0xa3, 0x6d, 0xd8, 0x7a, 0x58, 0x9e, 0x26, 0xd4, 0xd1, 0x06, 0xac, 0xcd,
	0x23, 0xf8, 0x72, 0x40, 0x28, 0xee, 0xda, 0x70, 0xfc, 0x11, 0x50, 0x2c, 0x86, 0x73, 0x1f, 0xe4,
	0x71, 0x83, 0x4e, 0xdf, 0xcd, 0x20, 0x7f, 0x17, 0x83, 0xca, 0xa7, 0xba, 0xce, 0x4a, 0xf1, 0x9b,
	0xb5, 0xd0, 0x61, 0x97, 0xdf, 0xad, 0x46, 0x47, 0xb5, 0x5d, 0x96, 0xb9, 0xe7, 0x07, 0xc7, 0xf9,
	0xf8, 0x87, 0x19, 0x5c, 0xb1, 0xec, 0xaa, 0x1c, 0x5c, 0x57, 0xcd, 0x9b, 0x6a, 0xff, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x0d, 0x00, 0x76, 0xc4, 0x26, 0x04, 0x00, 0x00,
}