// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

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

type Modulation int32

const (
	// LoRa
	Modulation_LORA Modulation = 0
	// FSK
	Modulation_FSK Modulation = 1
	// LR-FHSS
	Modulation_LR_FHSS Modulation = 2
)

var Modulation_name = map[int32]string{
	0: "LORA",
	1: "FSK",
	2: "LR_FHSS",
}

var Modulation_value = map[string]int32{
	"LORA":    0,
	"FSK":     1,
	"LR_FHSS": 2,
}

func (x Modulation) String() string {
	return proto.EnumName(Modulation_name, int32(x))
}

func (Modulation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type Region int32

const (
	// EU868
	Region_EU868 Region = 0
	// US915
	Region_US915 Region = 2
	// CN779
	Region_CN779 Region = 3
	// EU433
	Region_EU433 Region = 4
	// AU915
	Region_AU915 Region = 5
	// CN470
	Region_CN470 Region = 6
	// AS923
	Region_AS923 Region = 7
	// AS923 with -1.80 MHz frequency offset
	Region_AS923_2 Region = 12
	// AS923 with -6.60 MHz frequency offset
	Region_AS923_3 Region = 13
	// AS923 with -5.90 MHz frequency offset
	Region_AS923_4 Region = 14
	// KR920
	Region_KR920 Region = 8
	// IN865
	Region_IN865 Region = 9
	// RU864
	Region_RU864 Region = 10
	// ISM2400 (LoRaWAN 2.4 GHz)
	Region_ISM2400 Region = 11
)

var Region_name = map[int32]string{
	0:  "EU868",
	2:  "US915",
	3:  "CN779",
	4:  "EU433",
	5:  "AU915",
	6:  "CN470",
	7:  "AS923",
	12: "AS923_2",
	13: "AS923_3",
	14: "AS923_4",
	8:  "KR920",
	9:  "IN865",
	10: "RU864",
	11: "ISM2400",
}

var Region_value = map[string]int32{
	"EU868":   0,
	"US915":   2,
	"CN779":   3,
	"EU433":   4,
	"AU915":   5,
	"CN470":   6,
	"AS923":   7,
	"AS923_2": 12,
	"AS923_3": 13,
	"AS923_4": 14,
	"KR920":   8,
	"IN865":   9,
	"RU864":   10,
	"ISM2400": 11,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}

func (Region) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

type MType int32

const (
	// JoinRequest.
	MType_JoinRequest MType = 0
	// JoinAccept.
	MType_JoinAccept MType = 1
	// UnconfirmedDataUp.
	MType_UnconfirmedDataUp MType = 2
	// UnconfirmedDataDown.
	MType_UnconfirmedDataDown MType = 3
	// ConfirmedDataUp.
	MType_ConfirmedDataUp MType = 4
	// ConfirmedDataDown.
	MType_ConfirmedDataDown MType = 5
	// RejoinRequest.
	MType_RejoinRequest MType = 6
	// Proprietary.
	MType_Proprietary MType = 7
)

var MType_name = map[int32]string{
	0: "JoinRequest",
	1: "JoinAccept",
	2: "UnconfirmedDataUp",
	3: "UnconfirmedDataDown",
	4: "ConfirmedDataUp",
	5: "ConfirmedDataDown",
	6: "RejoinRequest",
	7: "Proprietary",
}

var MType_value = map[string]int32{
	"JoinRequest":         0,
	"JoinAccept":          1,
	"UnconfirmedDataUp":   2,
	"UnconfirmedDataDown": 3,
	"ConfirmedDataUp":     4,
	"ConfirmedDataDown":   5,
	"RejoinRequest":       6,
	"Proprietary":         7,
}

func (x MType) String() string {
	return proto.EnumName(MType_name, int32(x))
}

func (MType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

type LocationSource int32

const (
	// Unknown.
	LocationSource_UNKNOWN LocationSource = 0
	// GPS.
	LocationSource_GPS LocationSource = 1
	// Manually configured.
	LocationSource_CONFIG LocationSource = 2
	// Geo resolver (TDOA).
	LocationSource_GEO_RESOLVER_TDOA LocationSource = 3
	// Geo resolver (RSSI).
	LocationSource_GEO_RESOLVER_RSSI LocationSource = 4
	// Geo resolver (GNSS).
	LocationSource_GEO_RESOLVER_GNSS LocationSource = 5
	// Geo resolver (WIFI).
	LocationSource_GEO_RESOLVER_WIFI LocationSource = 6
)

var LocationSource_name = map[int32]string{
	0: "UNKNOWN",
	1: "GPS",
	2: "CONFIG",
	3: "GEO_RESOLVER_TDOA",
	4: "GEO_RESOLVER_RSSI",
	5: "GEO_RESOLVER_GNSS",
	6: "GEO_RESOLVER_WIFI",
}

var LocationSource_value = map[string]int32{
	"UNKNOWN":           0,
	"GPS":               1,
	"CONFIG":            2,
	"GEO_RESOLVER_TDOA": 3,
	"GEO_RESOLVER_RSSI": 4,
	"GEO_RESOLVER_GNSS": 5,
	"GEO_RESOLVER_WIFI": 6,
}

func (x LocationSource) String() string {
	return proto.EnumName(LocationSource_name, int32(x))
}

func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

type KeyEnvelope struct {
	// KEK label.
	KekLabel string `protobuf:"bytes,1,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	// AES key (when the kek_label is set, this key is encrypted using a key
	// known to the join-server and application-server.
	// For more information please refer to the LoRaWAN Backend Interface
	// 'Key Transport Security' section.
	AesKey               []byte   `protobuf:"bytes,2,opt,name=aes_key,json=aesKey,proto3" json:"aes_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEnvelope) Reset()         { *m = KeyEnvelope{} }
func (m *KeyEnvelope) String() string { return proto.CompactTextString(m) }
func (*KeyEnvelope) ProtoMessage()    {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *KeyEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyEnvelope.Unmarshal(m, b)
}
func (m *KeyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyEnvelope.Marshal(b, m, deterministic)
}
func (m *KeyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEnvelope.Merge(m, src)
}
func (m *KeyEnvelope) XXX_Size() int {
	return xxx_messageInfo_KeyEnvelope.Size(m)
}
func (m *KeyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEnvelope proto.InternalMessageInfo

func (m *KeyEnvelope) GetKekLabel() string {
	if m != nil {
		return m.KekLabel
	}
	return ""
}

func (m *KeyEnvelope) GetAesKey() []byte {
	if m != nil {
		return m.AesKey
	}
	return nil
}

type Location struct {
	// Latitude.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude float64 `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Location source.
	Source LocationSource `protobuf:"varint,4,opt,name=source,proto3,enum=common.LocationSource" json:"source,omitempty"`
	// Accuracy (in meters).
	Accuracy             uint32   `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetSource() LocationSource {
	if m != nil {
		return m.Source
	}
	return LocationSource_UNKNOWN
}

func (m *Location) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func init() {
	proto.RegisterEnum("common.Modulation", Modulation_name, Modulation_value)
	proto.RegisterEnum("common.Region", Region_name, Region_value)
	proto.RegisterEnum("common.MType", MType_name, MType_value)
	proto.RegisterEnum("common.LocationSource", LocationSource_name, LocationSource_value)
	proto.RegisterType((*KeyEnvelope)(nil), "common.KeyEnvelope")
	proto.RegisterType((*Location)(nil), "common.Location")
}

func init() {
	proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6)
}

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x5f, 0x53, 0xda, 0x4e,
	0x14, 0x75, 0xf9, 0x13, 0xe0, 0xa2, 0x78, 0x5d, 0xe7, 0xf7, 0x93, 0x69, 0xfb, 0xc0, 0xf8, 0xc4,
	0x30, 0x2d, 0x50, 0x40, 0x85, 0x47, 0x44, 0xa0, 0x14, 0x0c, 0xcc, 0xc6, 0xd4, 0x69, 0x5f, 0x98,
	0x25, 0x6e, 0x31, 0x25, 0x66, 0xd3, 0x10, 0xec, 0xf0, 0x25, 0xfa, 0x39, 0x3a, 0x9d, 0x7e, 0xc8,
	0xce, 0x26, 0x28, 0x55, 0x9f, 0xf6, 0x9c, 0x73, 0xcf, 0xfd, 0xb3, 0xb3, 0x7b, 0xe1, 0xd0, 0x92,
	0x77, 0x77, 0xd2, 0xad, 0x44, 0x47, 0xd9, 0xf3, 0x65, 0x20, 0xa9, 0x16, 0xb1, 0xe3, 0x0e, 0x64,
	0x87, 0x62, 0xdd, 0x75, 0xef, 0x85, 0x23, 0x3d, 0x41, 0x5f, 0x43, 0x66, 0x21, 0x16, 0x53, 0x87,
	0xcf, 0x84, 0x93, 0x27, 0x05, 0x52, 0xcc, 0xb0, 0xf4, 0x42, 0x2c, 0x46, 0x8a, 0xd3, 0x23, 0x48,
	0x71, 0xb1, 0x9c, 0x2e, 0xc4, 0x3a, 0x1f, 0x2b, 0x90, 0xe2, 0x2e, 0xd3, 0xb8, 0x58, 0x0e, 0xc5,
	0xfa, 0xf8, 0x0f, 0x81, 0xf4, 0x48, 0x5a, 0x3c, 0xb0, 0xa5, 0x4b, 0x5f, 0x41, 0xda, 0xe1, 0x81,
	0x1d, 0xac, 0x6e, 0x44, 0x58, 0x81, 0xb0, 0x47, 0x4e, 0xdf, 0x40, 0xc6, 0x91, 0xee, 0x3c, 0x0a,
	0xc6, 0xc2, 0xe0, 0x56, 0x50, 0x99, 0xdc, 0xd9, 0x64, 0xc6, 0xa3, 0xcc, 0x07, 0x4e, 0xcb, 0xa0,
	0x2d, 0xe5, 0xca, 0xb7, 0x44, 0x3e, 0x51, 0x20, 0xc5, 0x5c, 0xed, 0xff, 0xf2, 0xe6, 0x3a, 0x0f,
	0x7d, 0x8d, 0x30, 0xca, 0x36, 0xae, 0xb0, 0x96, 0x65, 0xad, 0x7c, 0x6e, 0xad, 0xf3, 0xc9, 0x02,
	0x29, 0xee, 0xb1, 0x47, 0x5e, 0x7a, 0x0b, 0x70, 0x29, 0x6f, 0x56, 0x4e, 0x34, 0x6f, 0x1a, 0x12,
	0xa3, 0x31, 0x6b, 0xe3, 0x0e, 0x4d, 0x41, 0xbc, 0x67, 0x0c, 0x91, 0xd0, 0x2c, 0xa4, 0x46, 0x6c,
	0xda, 0xfb, 0x60, 0x18, 0x18, 0x2b, 0xfd, 0x26, 0xa0, 0x31, 0x31, 0x57, 0xd6, 0x0c, 0x24, 0xbb,
	0x66, 0xf3, 0xb4, 0x89, 0x3b, 0x0a, 0x9a, 0x46, 0xeb, 0xfd, 0x09, 0xc6, 0x14, 0xec, 0xe8, 0x67,
	0x67, 0x2d, 0x8c, 0x47, 0x86, 0x46, 0xbd, 0x8e, 0x09, 0x05, 0xdb, 0xa6, 0x32, 0x24, 0x23, 0x43,
	0xe3, 0xac, 0x8a, 0x5a, 0xa8, 0x1a, 0xad, 0x5a, 0x1d, 0x53, 0xaa, 0x49, 0x08, 0xa7, 0x35, 0xdc,
	0xdd, 0x92, 0x3a, 0xee, 0x6d, 0x49, 0x03, 0x73, 0x2a, 0x63, 0xc8, 0x5a, 0xb5, 0x2a, 0xa6, 0x15,
	0x1c, 0xe8, 0xcd, 0xd3, 0x13, 0xcc, 0x28, 0xc8, 0xcc, 0xe6, 0x69, 0x03, 0x41, 0xb9, 0x07, 0xc6,
	0x65, 0xad, 0x51, 0xad, 0x62, 0xb6, 0xf4, 0x8b, 0x40, 0xf2, 0xf2, 0x6a, 0xed, 0x09, 0xba, 0x0f,
	0xd9, 0x8f, 0xd2, 0x76, 0x99, 0xf8, 0xbe, 0x12, 0xcb, 0x00, 0x77, 0x68, 0x0e, 0x40, 0x09, 0x6d,
	0xcb, 0x12, 0x5e, 0x80, 0x84, 0xfe, 0x07, 0x07, 0xa6, 0x6b, 0x49, 0xf7, 0xab, 0xed, 0xdf, 0x89,
	0x9b, 0x0b, 0x1e, 0x70, 0xd3, 0xc3, 0x18, 0x3d, 0x82, 0xc3, 0x67, 0xf2, 0x85, 0xfc, 0xe1, 0x62,
	0x9c, 0x1e, 0xc2, 0x7e, 0xe7, 0x99, 0x3b, 0xa1, 0x8a, 0x74, 0x5e, 0x78, 0x93, 0xf4, 0x00, 0xf6,
	0x98, 0xf8, 0xf6, 0x4f, 0x7b, 0x4d, 0xcd, 0x33, 0xf1, 0xa5, 0xe7, 0xdb, 0x22, 0xe0, 0xfe, 0x1a,
	0x53, 0xa5, 0x9f, 0x04, 0x72, 0x4f, 0x1f, 0x4f, 0x5d, 0xc5, 0xd4, 0x87, 0xfa, 0xf8, 0x5a, 0x8f,
	0x5e, 0xa3, 0x3f, 0x31, 0x90, 0x50, 0x00, 0xad, 0x33, 0xd6, 0x7b, 0x83, 0x3e, 0xc6, 0x54, 0xbf,
	0x7e, 0x77, 0x3c, 0x65, 0x5d, 0x63, 0x3c, 0xfa, 0xd4, 0x65, 0xd3, 0xab, 0x8b, 0x71, 0x1b, 0xe3,
	0x2f, 0x64, 0x66, 0x18, 0x83, 0x68, 0xba, 0x27, 0x72, 0x5f, 0x37, 0x0c, 0x4c, 0xbe, 0x90, 0xaf,
	0x07, 0xbd, 0x01, 0x6a, 0xe7, 0x9f, 0x21, 0x6f, 0xcb, 0xb2, 0x75, 0x6b, 0xfb, 0xde, 0x32, 0xe0,
	0xd6, 0xa2, 0xcc, 0x3d, 0x7b, 0xf3, 0xcb, 0xce, 0xb3, 0x9d, 0xf0, 0x9c, 0xa8, 0xdd, 0x99, 0x90,
	0x2f, 0xe5, 0xb9, 0x1d, 0xdc, 0xae, 0x66, 0x2a, 0x5a, 0x99, 0xf9, 0xd2, 0xe2, 0xdc, 0xaf, 0x6c,
	0x13, 0xdf, 0x71, 0xcf, 0xae, 0xcc, 0x65, 0xe5, 0xbe, 0xbe, 0xd9, 0xb9, 0x99, 0x16, 0x2e, 0x5d,
	0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x6c, 0x54, 0x97, 0x8b, 0x03, 0x00, 0x00,
}