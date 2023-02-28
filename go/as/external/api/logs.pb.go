// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/logs.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 日志：记录某用户(author)的操作行为（行为名称，汉语描述，便于理解),其中，记录日志方法表示其操作的行为(post/get/delete/put),url表示API的路径，IP表示操作用户的电脑IP(192.168.1.18).
// 日志列表和日志详情：如果用户操作的行为没用数据变化（比如登录)，则不产生详情数据。如果操作前后又数据变化，则详情的数据（其实是个json）将记录前后的数据变化情况。
type Log struct {
	// log ID. Will be set automatically on create.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 行为名称
	ActTag string `protobuf:"bytes,2,opt,name=act_tag,json=actTag,proto3" json:"act_tag,omitempty"`
	// API的请求路径
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// 记录日志方法
	LogMethod string `protobuf:"bytes,4,opt,name=log_method,json=logMethod,proto3" json:"log_method,omitempty"`
	// 操作IP
	LogIp string `protobuf:"bytes,5,opt,name=log_ip,json=logIp,proto3" json:"log_ip,omitempty"`
	// 0、未启用 1、已启用 2、已停用
	Status int32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	// 行为操作者
	Author string `protobuf:"bytes,7,opt,name=author,proto3" json:"author,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Remark    string               `protobuf:"bytes,10,opt,name=remark,proto3" json:"remark,omitempty"`
	// API的请求路径
	LogData              string   `protobuf:"bytes,11,opt,name=log_data,json=logData,proto3" json:"log_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_7865615dda6b0285, []int{0}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Log) GetActTag() string {
	if m != nil {
		return m.ActTag
	}
	return ""
}

func (m *Log) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Log) GetLogMethod() string {
	if m != nil {
		return m.LogMethod
	}
	return ""
}

func (m *Log) GetLogIp() string {
	if m != nil {
		return m.LogIp
	}
	return ""
}

func (m *Log) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Log) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Log) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Log) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Log) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Log) GetLogData() string {
	if m != nil {
		return m.LogData
	}
	return ""
}

type CreateLogsRequest struct {
	// Test object to create.
	Log                  *Log     `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLogsRequest) Reset()         { *m = CreateLogsRequest{} }
func (m *CreateLogsRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLogsRequest) ProtoMessage()    {}
func (*CreateLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7865615dda6b0285, []int{1}
}

func (m *CreateLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLogsRequest.Unmarshal(m, b)
}
func (m *CreateLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLogsRequest.Marshal(b, m, deterministic)
}
func (m *CreateLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLogsRequest.Merge(m, src)
}
func (m *CreateLogsRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLogsRequest.Size(m)
}
func (m *CreateLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLogsRequest proto.InternalMessageInfo

func (m *CreateLogsRequest) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type CreateLogsResponse struct {
	// Test ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLogsResponse) Reset()         { *m = CreateLogsResponse{} }
func (m *CreateLogsResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLogsResponse) ProtoMessage()    {}
func (*CreateLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7865615dda6b0285, []int{2}
}

func (m *CreateLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLogsResponse.Unmarshal(m, b)
}
func (m *CreateLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLogsResponse.Marshal(b, m, deterministic)
}
func (m *CreateLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLogsResponse.Merge(m, src)
}
func (m *CreateLogsResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLogsResponse.Size(m)
}
func (m *CreateLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLogsResponse proto.InternalMessageInfo

func (m *CreateLogsResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetLogsRequest struct {
	// Test ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLogsRequest) Reset()         { *m = GetLogsRequest{} }
func (m *GetLogsRequest) String() string { return proto.CompactTextString(m) }
func (*GetLogsRequest) ProtoMessage()    {}
func (*GetLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7865615dda6b0285, []int{3}
}

func (m *GetLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLogsRequest.Unmarshal(m, b)
}
func (m *GetLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLogsRequest.Marshal(b, m, deterministic)
}
func (m *GetLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLogsRequest.Merge(m, src)
}
func (m *GetLogsRequest) XXX_Size() int {
	return xxx_messageInfo_GetLogsRequest.Size(m)
}
func (m *GetLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLogsRequest proto.InternalMessageInfo

func (m *GetLogsRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetLogsResponse struct {
	// Test object.
	Log *Log `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetLogsResponse) Reset()         { *m = GetLogsResponse{} }
func (m *GetLogsResponse) String() string { return proto.CompactTextString(m) }
func (*GetLogsResponse) ProtoMessage()    {}
func (*GetLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7865615dda6b0285, []int{4}
}

func (m *GetLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLogsResponse.Unmarshal(m, b)
}
func (m *GetLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLogsResponse.Marshal(b, m, deterministic)
}
func (m *GetLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLogsResponse.Merge(m, src)
}
func (m *GetLogsResponse) XXX_Size() int {
	return xxx_messageInfo_GetLogsResponse.Size(m)
}
func (m *GetLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLogsResponse proto.InternalMessageInfo

func (m *GetLogsResponse) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *GetLogsResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetLogsResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type DeleteLogsRequest struct {
	// Test ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLogsRequest) Reset()         { *m = DeleteLogsRequest{} }
func (m *DeleteLogsRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteLogsRequest) ProtoMessage()    {}
func (*DeleteLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7865615dda6b0285, []int{5}
}

func (m *DeleteLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLogsRequest.Unmarshal(m, b)
}
func (m *DeleteLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLogsRequest.Marshal(b, m, deterministic)
}
func (m *DeleteLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLogsRequest.Merge(m, src)
}
func (m *DeleteLogsRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteLogsRequest.Size(m)
}
func (m *DeleteLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLogsRequest proto.InternalMessageInfo

func (m *DeleteLogsRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListLogsRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListLogsRequest) Reset()         { *m = ListLogsRequest{} }
func (m *ListLogsRequest) String() string { return proto.CompactTextString(m) }
func (*ListLogsRequest) ProtoMessage()    {}
func (*ListLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7865615dda6b0285, []int{6}
}

func (m *ListLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLogsRequest.Unmarshal(m, b)
}
func (m *ListLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLogsRequest.Marshal(b, m, deterministic)
}
func (m *ListLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLogsRequest.Merge(m, src)
}
func (m *ListLogsRequest) XXX_Size() int {
	return xxx_messageInfo_ListLogsRequest.Size(m)
}
func (m *ListLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListLogsRequest proto.InternalMessageInfo

func (m *ListLogsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListLogsRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListLogsResponse struct {
	// Total number of test.
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Test within the result-set.
	Result               []*Log   `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListLogsResponse) Reset()         { *m = ListLogsResponse{} }
func (m *ListLogsResponse) String() string { return proto.CompactTextString(m) }
func (*ListLogsResponse) ProtoMessage()    {}
func (*ListLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7865615dda6b0285, []int{7}
}

func (m *ListLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLogsResponse.Unmarshal(m, b)
}
func (m *ListLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLogsResponse.Marshal(b, m, deterministic)
}
func (m *ListLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLogsResponse.Merge(m, src)
}
func (m *ListLogsResponse) XXX_Size() int {
	return xxx_messageInfo_ListLogsResponse.Size(m)
}
func (m *ListLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListLogsResponse proto.InternalMessageInfo

func (m *ListLogsResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListLogsResponse) GetResult() []*Log {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Log)(nil), "api.Log")
	proto.RegisterType((*CreateLogsRequest)(nil), "api.CreateLogsRequest")
	proto.RegisterType((*CreateLogsResponse)(nil), "api.CreateLogsResponse")
	proto.RegisterType((*GetLogsRequest)(nil), "api.GetLogsRequest")
	proto.RegisterType((*GetLogsResponse)(nil), "api.GetLogsResponse")
	proto.RegisterType((*DeleteLogsRequest)(nil), "api.DeleteLogsRequest")
	proto.RegisterType((*ListLogsRequest)(nil), "api.ListLogsRequest")
	proto.RegisterType((*ListLogsResponse)(nil), "api.ListLogsResponse")
}

func init() {
	proto.RegisterFile("as/external/api/logs.proto", fileDescriptor_7865615dda6b0285)
}

var fileDescriptor_7865615dda6b0285 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0xe3, 0x26, 0x6d, 0x4e, 0xa4, 0xb4, 0x9d, 0xdb, 0xa6, 0xbe, 0xbe, 0xf7, 0xaa, 0xb9,
	0x86, 0x45, 0x54, 0x09, 0x5b, 0x6a, 0x17, 0x08, 0x36, 0xa8, 0x3f, 0x50, 0x90, 0x82, 0xa8, 0x42,
	0xd9, 0xb0, 0x89, 0x26, 0xf6, 0x74, 0x32, 0xaa, 0x9d, 0x31, 0x9e, 0xe3, 0x0a, 0x84, 0xd8, 0xf0,
	0x0a, 0x3c, 0x02, 0x6f, 0xc2, 0x0b, 0xb0, 0xe0, 0x15, 0x78, 0x10, 0x34, 0xe3, 0x49, 0xdb, 0x38,
	0x12, 0x08, 0x76, 0x39, 0xdf, 0xf9, 0xce, 0xe7, 0x73, 0xbe, 0x4f, 0x19, 0xf0, 0xa9, 0x8a, 0xd8,
	0x5b, 0x64, 0xc5, 0x8c, 0xa6, 0x11, 0xcd, 0x45, 0x94, 0x4a, 0xae, 0xc2, 0xbc, 0x90, 0x28, 0x89,
	0x4b, 0x73, 0xe1, 0xff, 0xcb, 0xa5, 0xe4, 0x29, 0x33, 0x3d, 0x3a, 0x9b, 0x49, 0xa4, 0x28, 0xe4,
	0xcc, 0x52, 0xfc, 0x5d, 0xdb, 0x35, 0xd5, 0xa4, 0xbc, 0x88, 0x50, 0x64, 0x4c, 0x21, 0xcd, 0x72,
	0x4b, 0xf8, 0xa7, 0x4e, 0x60, 0x59, 0x8e, 0xef, 0xaa, 0x66, 0xf0, 0xb5, 0x01, 0xee, 0x50, 0x72,
	0xd2, 0x85, 0x86, 0x48, 0x3c, 0xa7, 0xef, 0x0c, 0xdc, 0x51, 0x43, 0x24, 0x64, 0x07, 0x56, 0x69,
	0x8c, 0x63, 0xa4, 0xdc, 0x6b, 0xf4, 0x9d, 0x41, 0x7b, 0xd4, 0xa2, 0x31, 0x9e, 0x53, 0x4e, 0x36,
	0xc0, 0x2d, 0x8b, 0xd4, 0x73, 0x0d, 0xa8, 0x7f, 0x92, 0xff, 0x00, 0x52, 0xc9, 0xc7, 0x19, 0xc3,
	0xa9, 0x4c, 0xbc, 0x15, 0xd3, 0x68, 0xa7, 0x92, 0x3f, 0x37, 0x00, 0xd9, 0x86, 0x96, 0x6e, 0x8b,
	0xdc, 0x6b, 0x9a, 0x56, 0x33, 0x95, 0xfc, 0x59, 0x4e, 0x7a, 0xd0, 0x52, 0x48, 0xb1, 0x54, 0x5e,
	0xab, 0xef, 0x0c, 0x9a, 0x23, 0x5b, 0x69, 0x9c, 0x96, 0x38, 0x95, 0x85, 0xb7, 0x6a, 0xbf, 0x6b,
	0x2a, 0xf2, 0x00, 0x20, 0x2e, 0x18, 0x45, 0x96, 0x8c, 0x29, 0x7a, 0x6b, 0x7d, 0x67, 0xd0, 0xd9,
	0xf7, 0xc3, 0xea, 0xb4, 0x70, 0x7e, 0x5a, 0x78, 0x3e, 0xbf, 0x7d, 0xd4, 0xb6, 0xec, 0x43, 0xd4,
	0xa3, 0x65, 0x9e, 0xcc, 0x47, 0xdb, 0xbf, 0x1e, 0xb5, 0xec, 0x43, 0xd4, 0xdb, 0x14, 0x2c, 0xa3,
	0xc5, 0xa5, 0x07, 0xd5, 0x36, 0x55, 0x45, 0xfe, 0x86, 0x35, 0x7d, 0x54, 0x42, 0x91, 0x7a, 0x1d,
	0xd3, 0x59, 0x4d, 0x25, 0x3f, 0xa1, 0x48, 0x83, 0x08, 0x36, 0x8f, 0xcd, 0xa7, 0x87, 0x92, 0xab,
	0x11, 0x7b, 0x53, 0x32, 0x85, 0xc4, 0x07, 0x37, 0x95, 0xdc, 0xf8, 0xdb, 0xd9, 0x5f, 0x0b, 0x69,
	0x2e, 0xc2, 0xa1, 0xe4, 0x23, 0x0d, 0x06, 0x77, 0x81, 0xdc, 0x1e, 0x50, 0xb9, 0x9c, 0x29, 0x56,
	0x0f, 0x24, 0xe8, 0x43, 0xf7, 0x94, 0xe1, 0x6d, 0xcd, 0x3a, 0xe3, 0xb3, 0x03, 0xeb, 0xd7, 0x14,
	0xab, 0xf2, 0x93, 0xef, 0xd6, 0x1c, 0x6d, 0xfc, 0xb9, 0xa3, 0xee, 0x6f, 0x38, 0x1a, 0xdc, 0x81,
	0xcd, 0x13, 0x96, 0xb2, 0x45, 0x7b, 0xea, 0xa7, 0x3c, 0x82, 0xf5, 0xa1, 0x50, 0x0b, 0xd7, 0x6e,
	0x41, 0x33, 0x15, 0x99, 0x40, 0xcb, 0xaa, 0x0a, 0x9d, 0x8f, 0xbc, 0xb8, 0x50, 0xac, 0xda, 0xdf,
	0x1d, 0xd9, 0x2a, 0x78, 0x05, 0x1b, 0x37, 0x02, 0xd6, 0x8b, 0x5d, 0xe8, 0xa0, 0x44, 0x9a, 0x8e,
	0x63, 0x59, 0xce, 0xe6, 0x3a, 0x60, 0xa0, 0x63, 0x8d, 0x90, 0xbe, 0x0e, 0x5b, 0x95, 0xa9, 0x16,
	0x73, 0x17, 0xfc, 0xb2, 0xf8, 0xfe, 0x97, 0x06, 0x74, 0xb4, 0xe6, 0x4b, 0x56, 0x5c, 0x89, 0x98,
	0x91, 0x17, 0xd0, 0xaa, 0xa2, 0x23, 0x3d, 0xc3, 0x5d, 0x0a, 0xde, 0xdf, 0x59, 0xc2, 0xab, 0x6d,
	0x82, 0xad, 0x8f, 0xdf, 0xbe, 0x7f, 0x6a, 0x74, 0x83, 0xf6, 0xf5, 0x5f, 0xfe, 0xa1, 0xb3, 0x47,
	0x9e, 0x82, 0x7b, 0xca, 0x90, 0xfc, 0x65, 0xa6, 0x16, 0xf3, 0xf6, 0xb7, 0x16, 0x41, 0xab, 0xd3,
	0x33, 0x3a, 0x1b, 0xa4, 0x7b, 0xad, 0x13, 0xbd, 0x17, 0xc9, 0x07, 0x72, 0x06, 0xad, 0xca, 0x67,
	0xbb, 0xda, 0x92, 0xe9, 0x7e, 0x6f, 0x29, 0xb0, 0xc7, 0xfa, 0x61, 0x98, 0x2b, 0xee, 0xd5, 0x15,
	0x9f, 0xc0, 0x8a, 0xf6, 0x94, 0x54, 0x7b, 0xd4, 0xf2, 0xf1, 0xb7, 0x6b, 0xa8, 0x5d, 0x6f, 0xd3,
	0x88, 0x75, 0xc8, 0xcd, 0x99, 0x47, 0x02, 0xfe, 0x17, 0x32, 0x8c, 0xa7, 0xa2, 0xc8, 0x15, 0xd2,
	0xf8, 0xd2, 0x0c, 0x52, 0x15, 0xce, 0xdf, 0x40, 0x5d, 0x1f, 0xb5, 0xb5, 0xca, 0x99, 0x5e, 0xec,
	0xcc, 0x79, 0x7d, 0x9f, 0x0b, 0x9c, 0x96, 0x93, 0x30, 0x96, 0x59, 0x34, 0x29, 0x64, 0x4c, 0x69,
	0x11, 0xdd, 0xcc, 0xdf, 0xd3, 0xca, 0x5c, 0x46, 0x57, 0x07, 0x51, 0xed, 0x25, 0x9d, 0xb4, 0xcc,
	0x69, 0x07, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x65, 0x43, 0xc6, 0xbe, 0x63, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogsServiceClient is the client API for LogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogsServiceClient interface {
	// Create creates logs.
	Create(ctx context.Context, in *CreateLogsRequest, opts ...grpc.CallOption) (*CreateLogsResponse, error)
	// Get returns the logs matching the given id.
	Get(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error)
	// Delete deletes the logs matching the given id.
	Delete(ctx context.Context, in *DeleteLogsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available logs.
	List(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error)
}

type logsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogsServiceClient(cc grpc.ClientConnInterface) LogsServiceClient {
	return &logsServiceClient{cc}
}

func (c *logsServiceClient) Create(ctx context.Context, in *CreateLogsRequest, opts ...grpc.CallOption) (*CreateLogsResponse, error) {
	out := new(CreateLogsResponse)
	err := c.cc.Invoke(ctx, "/api.LogsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsServiceClient) Get(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error) {
	out := new(GetLogsResponse)
	err := c.cc.Invoke(ctx, "/api.LogsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsServiceClient) Delete(ctx context.Context, in *DeleteLogsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.LogsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsServiceClient) List(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error) {
	out := new(ListLogsResponse)
	err := c.cc.Invoke(ctx, "/api.LogsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogsServiceServer is the server API for LogsService service.
type LogsServiceServer interface {
	// Create creates logs.
	Create(context.Context, *CreateLogsRequest) (*CreateLogsResponse, error)
	// Get returns the logs matching the given id.
	Get(context.Context, *GetLogsRequest) (*GetLogsResponse, error)
	// Delete deletes the logs matching the given id.
	Delete(context.Context, *DeleteLogsRequest) (*empty.Empty, error)
	// List lists the available logs.
	List(context.Context, *ListLogsRequest) (*ListLogsResponse, error)
}

// UnimplementedLogsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogsServiceServer struct {
}

func (*UnimplementedLogsServiceServer) Create(ctx context.Context, req *CreateLogsRequest) (*CreateLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedLogsServiceServer) Get(ctx context.Context, req *GetLogsRequest) (*GetLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedLogsServiceServer) Delete(ctx context.Context, req *DeleteLogsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedLogsServiceServer) List(ctx context.Context, req *ListLogsRequest) (*ListLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterLogsServiceServer(s *grpc.Server, srv LogsServiceServer) {
	s.RegisterService(&_LogsService_serviceDesc, srv)
}

func _LogsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LogsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServiceServer).Create(ctx, req.(*CreateLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LogsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServiceServer).Get(ctx, req.(*GetLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LogsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServiceServer).Delete(ctx, req.(*DeleteLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LogsService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServiceServer).List(ctx, req.(*ListLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.LogsService",
	HandlerType: (*LogsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LogsService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LogsService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LogsService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LogsService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/logs.proto",
}
