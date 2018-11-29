// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=pb.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "pb.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "pb.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "pb.AddressBook")
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor_1eb1a68c9dd6d429) }

var fileDescriptor_1eb1a68c9dd6d429 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x9a, 0x06, 0x3b, 0x91, 0x52, 0x87, 0x52, 0x42, 0x2f, 0x86, 0x9c, 0x02, 0xc2,
	0x16, 0xeb, 0xd9, 0x83, 0x85, 0x82, 0xa2, 0x35, 0x65, 0xa9, 0x78, 0x94, 0x84, 0x8c, 0x35, 0x34,
	0xc9, 0x2e, 0xc9, 0xe6, 0xd0, 0xa7, 0xf2, 0x15, 0x25, 0xbb, 0xa9, 0x8a, 0xb7, 0x7f, 0x67, 0x3e,
	0x66, 0xbf, 0x19, 0xb8, 0x4c, 0xb2, 0xac, 0xa6, 0xa6, 0x49, 0x85, 0x38, 0x30, 0x59, 0x0b, 0x25,
	0xd0, 0x96, 0xe9, 0xfc, 0x6a, 0x2f, 0xc4, 0xbe, 0xa0, 0x85, 0xae, 0xa4, 0xed, 0xc7, 0x42, 0xe5,
	0x25, 0x35, 0x2a, 0x29, 0xa5, 0x81, 0xc2, 0x2f, 0x1b, 0xdc, 0x2d, 0xd5, 0x8d, 0xa8, 0x10, 0xc1,
	0xa9, 0x92, 0x92, 0x7c, 0x2b, 0xb0, 0xa2, 0x11, 0xd7, 0x19, 0xc7, 0x60, 0xe7, 0x99, 0x6f, 0x07,
	0x56, 0x34, 0xe4, 0x76, 0x9e, 0xe1, 0x14, 0x86, 0x54, 0x26, 0x79, 0xe1, 0x0f, 0x34, 0x64, 0x1e,
	0xc8, 0xc0, 0x95, 0x9f, 0xa2, 0xa2, 0xc6, 0x77, 0x82, 0x41, 0xe4, 0x2d, 0x67, 0x4c, 0xa6, 0xcc,
	0x4c, 0x65, 0xdb, 0xae, 0xf1, 0xd2, 0x96, 0x29, 0xd5, 0xbc, 0xa7, 0xf0, 0x0e, 0x2e, 0x8a, 0xa4,
	0x51, 0xef, 0xad, 0xcc, 0x12, 0x45, 0x99, 0x3f, 0x0c, 0xac, 0xc8, 0x5b, 0xce, 0x99, 0x91, 0x65,
	0x27, 0x59, 0xb6, 0x3b, 0xc9, 0x72, 0xaf, 0xe3, 0x5f, 0x0d, 0x3e, 0x8f, 0xc1, 0xfb, 0x33, 0x15,
	0x67, 0xe0, 0x56, 0x3a, 0xf5, 0xe6, 0xfd, 0x0b, 0x23, 0x70, 0xd4, 0x51, 0x92, 0xb6, 0x1f, 0x2f,
	0xa7, 0xff, 0x9d, 0x76, 0x47, 0x49, 0x5c, 0x13, 0xe1, 0x35, 0x8c, 0x7e, 0x4a, 0x08, 0xe0, 0x6e,
	0xe2, 0xd5, 0xe3, 0xf3, 0x7a, 0x72, 0x86, 0xe7, 0xe0, 0x3c, 0xc4, 0x9b, 0xf5, 0xc4, 0xea, 0xd2,
	0x5b, 0xcc, 0x9f, 0x26, 0x76, 0x78, 0x03, 0xde, 0xbd, 0xb9, 0xf5, 0x4a, 0x88, 0x03, 0x86, 0xe0,
	0x4a, 0x12, 0xb2, 0xe8, 0xee, 0xd6, 0xed, 0x0e, 0xbf, 0xff, 0xf0, 0xbe, 0x93, 0xba, 0x7a, 0xa3,
	0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x3c, 0x9c, 0x00, 0xa5, 0x01, 0x00, 0x00,
}
