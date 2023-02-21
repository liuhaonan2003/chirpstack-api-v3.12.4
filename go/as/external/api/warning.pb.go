// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/warning.proto

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

type Warning struct {
	// Warning ID.
	// Will be set automatically on create.
	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// Optional remark to store with the warning.
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Status               int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	OrganizationId       int64    `protobuf:"varint,5,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	OrganizationName     string   `protobuf:"bytes,6,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	DeviceType           int32    `protobuf:"varint,7,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	SerialNumber         string   `protobuf:"bytes,8,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Codes                string   `protobuf:"bytes,9,opt,name=codes,proto3" json:"codes,omitempty"`
	WarningTime          int64    `protobuf:"varint,10,opt,name=warning_time,json=warningTime,proto3" json:"warning_time,omitempty"`
	Title                string   `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	WarningLevel         int32    `protobuf:"varint,13,opt,name=warning_level,json=warningLevel,proto3" json:"warning_level,omitempty"`
	WarningCount         int32    `protobuf:"varint,14,opt,name=warning_count,json=warningCount,proto3" json:"warning_count,omitempty"`
	WorkOrderId          int64    `protobuf:"varint,15,opt,name=work_order_id,json=workOrderId,proto3" json:"work_order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Warning) Reset()         { *m = Warning{} }
func (m *Warning) String() string { return proto.CompactTextString(m) }
func (*Warning) ProtoMessage()    {}
func (*Warning) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{0}
}

func (m *Warning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Warning.Unmarshal(m, b)
}
func (m *Warning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Warning.Marshal(b, m, deterministic)
}
func (m *Warning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Warning.Merge(m, src)
}
func (m *Warning) XXX_Size() int {
	return xxx_messageInfo_Warning.Size(m)
}
func (m *Warning) XXX_DiscardUnknown() {
	xxx_messageInfo_Warning.DiscardUnknown(m)
}

var xxx_messageInfo_Warning proto.InternalMessageInfo

func (m *Warning) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Warning) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Warning) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Warning) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Warning) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *Warning) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *Warning) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

func (m *Warning) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *Warning) GetCodes() string {
	if m != nil {
		return m.Codes
	}
	return ""
}

func (m *Warning) GetWarningTime() int64 {
	if m != nil {
		return m.WarningTime
	}
	return 0
}

func (m *Warning) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Warning) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Warning) GetWarningLevel() int32 {
	if m != nil {
		return m.WarningLevel
	}
	return 0
}

func (m *Warning) GetWarningCount() int32 {
	if m != nil {
		return m.WarningCount
	}
	return 0
}

func (m *Warning) GetWorkOrderId() int64 {
	if m != nil {
		return m.WorkOrderId
	}
	return 0
}

type WarningListItem struct {
	// User ID.
	// Will be set automatically on create.
	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// remark of the warning.
	Remark           string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Status           int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	OrganizationId   int64  `protobuf:"varint,5,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	OrganizationName string `protobuf:"bytes,6,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	DeviceType       int32  `protobuf:"varint,7,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	SerialNumber     string `protobuf:"bytes,8,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Codes            string `protobuf:"bytes,9,opt,name=codes,proto3" json:"codes,omitempty"`
	WarningTime      int64  `protobuf:"varint,10,opt,name=warning_time,json=warningTime,proto3" json:"warning_time,omitempty"`
	Title            string `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	Description      string `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	WarningLevel     int32  `protobuf:"varint,13,opt,name=warning_level,json=warningLevel,proto3" json:"warning_level,omitempty"`
	WarningCount     int32  `protobuf:"varint,14,opt,name=warning_count,json=warningCount,proto3" json:"warning_count,omitempty"`
	WorkOrderId      int64  `protobuf:"varint,15,opt,name=work_order_id,json=workOrderId,proto3" json:"work_order_id,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WarningListItem) Reset()         { *m = WarningListItem{} }
func (m *WarningListItem) String() string { return proto.CompactTextString(m) }
func (*WarningListItem) ProtoMessage()    {}
func (*WarningListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{1}
}

func (m *WarningListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WarningListItem.Unmarshal(m, b)
}
func (m *WarningListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WarningListItem.Marshal(b, m, deterministic)
}
func (m *WarningListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WarningListItem.Merge(m, src)
}
func (m *WarningListItem) XXX_Size() int {
	return xxx_messageInfo_WarningListItem.Size(m)
}
func (m *WarningListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_WarningListItem.DiscardUnknown(m)
}

var xxx_messageInfo_WarningListItem proto.InternalMessageInfo

func (m *WarningListItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WarningListItem) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *WarningListItem) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *WarningListItem) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *WarningListItem) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *WarningListItem) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *WarningListItem) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

func (m *WarningListItem) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *WarningListItem) GetCodes() string {
	if m != nil {
		return m.Codes
	}
	return ""
}

func (m *WarningListItem) GetWarningTime() int64 {
	if m != nil {
		return m.WarningTime
	}
	return 0
}

func (m *WarningListItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *WarningListItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WarningListItem) GetWarningLevel() int32 {
	if m != nil {
		return m.WarningLevel
	}
	return 0
}

func (m *WarningListItem) GetWarningCount() int32 {
	if m != nil {
		return m.WarningCount
	}
	return 0
}

func (m *WarningListItem) GetWorkOrderId() int64 {
	if m != nil {
		return m.WorkOrderId
	}
	return 0
}

func (m *WarningListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *WarningListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateWarningRequest struct {
	// Warning object to create.
	Warning              *Warning `protobuf:"bytes,1,opt,name=warning,proto3" json:"warning,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWarningRequest) Reset()         { *m = CreateWarningRequest{} }
func (m *CreateWarningRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWarningRequest) ProtoMessage()    {}
func (*CreateWarningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{2}
}

func (m *CreateWarningRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWarningRequest.Unmarshal(m, b)
}
func (m *CreateWarningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWarningRequest.Marshal(b, m, deterministic)
}
func (m *CreateWarningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWarningRequest.Merge(m, src)
}
func (m *CreateWarningRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWarningRequest.Size(m)
}
func (m *CreateWarningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWarningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWarningRequest proto.InternalMessageInfo

func (m *CreateWarningRequest) GetWarning() *Warning {
	if m != nil {
		return m.Warning
	}
	return nil
}

type CreateWarningResponse struct {
	// Warning ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWarningResponse) Reset()         { *m = CreateWarningResponse{} }
func (m *CreateWarningResponse) String() string { return proto.CompactTextString(m) }
func (*CreateWarningResponse) ProtoMessage()    {}
func (*CreateWarningResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{3}
}

func (m *CreateWarningResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWarningResponse.Unmarshal(m, b)
}
func (m *CreateWarningResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWarningResponse.Marshal(b, m, deterministic)
}
func (m *CreateWarningResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWarningResponse.Merge(m, src)
}
func (m *CreateWarningResponse) XXX_Size() int {
	return xxx_messageInfo_CreateWarningResponse.Size(m)
}
func (m *CreateWarningResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWarningResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWarningResponse proto.InternalMessageInfo

func (m *CreateWarningResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetWarningRequest struct {
	// Warning ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWarningRequest) Reset()         { *m = GetWarningRequest{} }
func (m *GetWarningRequest) String() string { return proto.CompactTextString(m) }
func (*GetWarningRequest) ProtoMessage()    {}
func (*GetWarningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{4}
}

func (m *GetWarningRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWarningRequest.Unmarshal(m, b)
}
func (m *GetWarningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWarningRequest.Marshal(b, m, deterministic)
}
func (m *GetWarningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWarningRequest.Merge(m, src)
}
func (m *GetWarningRequest) XXX_Size() int {
	return xxx_messageInfo_GetWarningRequest.Size(m)
}
func (m *GetWarningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWarningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWarningRequest proto.InternalMessageInfo

func (m *GetWarningRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetWarningResponse struct {
	// Warning object.
	Warning *Warning `protobuf:"bytes,1,opt,name=warning,proto3" json:"warning,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetWarningResponse) Reset()         { *m = GetWarningResponse{} }
func (m *GetWarningResponse) String() string { return proto.CompactTextString(m) }
func (*GetWarningResponse) ProtoMessage()    {}
func (*GetWarningResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{5}
}

func (m *GetWarningResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWarningResponse.Unmarshal(m, b)
}
func (m *GetWarningResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWarningResponse.Marshal(b, m, deterministic)
}
func (m *GetWarningResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWarningResponse.Merge(m, src)
}
func (m *GetWarningResponse) XXX_Size() int {
	return xxx_messageInfo_GetWarningResponse.Size(m)
}
func (m *GetWarningResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWarningResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWarningResponse proto.InternalMessageInfo

func (m *GetWarningResponse) GetWarning() *Warning {
	if m != nil {
		return m.Warning
	}
	return nil
}

func (m *GetWarningResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetWarningResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateWarningRequest struct {
	// Warning object to update.
	Warning              *Warning `protobuf:"bytes,1,opt,name=warning,proto3" json:"warning,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWarningRequest) Reset()         { *m = UpdateWarningRequest{} }
func (m *UpdateWarningRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateWarningRequest) ProtoMessage()    {}
func (*UpdateWarningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{6}
}

func (m *UpdateWarningRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWarningRequest.Unmarshal(m, b)
}
func (m *UpdateWarningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWarningRequest.Marshal(b, m, deterministic)
}
func (m *UpdateWarningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWarningRequest.Merge(m, src)
}
func (m *UpdateWarningRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateWarningRequest.Size(m)
}
func (m *UpdateWarningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWarningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWarningRequest proto.InternalMessageInfo

func (m *UpdateWarningRequest) GetWarning() *Warning {
	if m != nil {
		return m.Warning
	}
	return nil
}

type DeleteWarningRequest struct {
	// Warning ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteWarningRequest) Reset()         { *m = DeleteWarningRequest{} }
func (m *DeleteWarningRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteWarningRequest) ProtoMessage()    {}
func (*DeleteWarningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{7}
}

func (m *DeleteWarningRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteWarningRequest.Unmarshal(m, b)
}
func (m *DeleteWarningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteWarningRequest.Marshal(b, m, deterministic)
}
func (m *DeleteWarningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteWarningRequest.Merge(m, src)
}
func (m *DeleteWarningRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteWarningRequest.Size(m)
}
func (m *DeleteWarningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteWarningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteWarningRequest proto.InternalMessageInfo

func (m *DeleteWarningRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListWarningRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWarningRequest) Reset()         { *m = ListWarningRequest{} }
func (m *ListWarningRequest) String() string { return proto.CompactTextString(m) }
func (*ListWarningRequest) ProtoMessage()    {}
func (*ListWarningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{8}
}

func (m *ListWarningRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWarningRequest.Unmarshal(m, b)
}
func (m *ListWarningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWarningRequest.Marshal(b, m, deterministic)
}
func (m *ListWarningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWarningRequest.Merge(m, src)
}
func (m *ListWarningRequest) XXX_Size() int {
	return xxx_messageInfo_ListWarningRequest.Size(m)
}
func (m *ListWarningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWarningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWarningRequest proto.InternalMessageInfo

func (m *ListWarningRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListWarningRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListWarningResponse struct {
	// Total number of warning.
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Warning within the result-set.
	Result               []*WarningListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListWarningResponse) Reset()         { *m = ListWarningResponse{} }
func (m *ListWarningResponse) String() string { return proto.CompactTextString(m) }
func (*ListWarningResponse) ProtoMessage()    {}
func (*ListWarningResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba362b698662e925, []int{9}
}

func (m *ListWarningResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWarningResponse.Unmarshal(m, b)
}
func (m *ListWarningResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWarningResponse.Marshal(b, m, deterministic)
}
func (m *ListWarningResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWarningResponse.Merge(m, src)
}
func (m *ListWarningResponse) XXX_Size() int {
	return xxx_messageInfo_ListWarningResponse.Size(m)
}
func (m *ListWarningResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWarningResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWarningResponse proto.InternalMessageInfo

func (m *ListWarningResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListWarningResponse) GetResult() []*WarningListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Warning)(nil), "api.Warning")
	proto.RegisterType((*WarningListItem)(nil), "api.WarningListItem")
	proto.RegisterType((*CreateWarningRequest)(nil), "api.CreateWarningRequest")
	proto.RegisterType((*CreateWarningResponse)(nil), "api.CreateWarningResponse")
	proto.RegisterType((*GetWarningRequest)(nil), "api.GetWarningRequest")
	proto.RegisterType((*GetWarningResponse)(nil), "api.GetWarningResponse")
	proto.RegisterType((*UpdateWarningRequest)(nil), "api.UpdateWarningRequest")
	proto.RegisterType((*DeleteWarningRequest)(nil), "api.DeleteWarningRequest")
	proto.RegisterType((*ListWarningRequest)(nil), "api.ListWarningRequest")
	proto.RegisterType((*ListWarningResponse)(nil), "api.ListWarningResponse")
}

func init() {
	proto.RegisterFile("as/external/api/warning.proto", fileDescriptor_ba362b698662e925)
}

var fileDescriptor_ba362b698662e925 = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0x96, 0x3d, 0xbb, 0x5e, 0xb6, 0xec, 0xf5, 0xc6, 0x1d, 0xb3, 0x99, 0x1d, 0x7e, 0xd6, 0x99,
	0xa0, 0xc4, 0x0a, 0x30, 0x23, 0x6d, 0x0e, 0x08, 0x0e, 0x48, 0xd9, 0x80, 0xa2, 0x95, 0xa2, 0x24,
	0x98, 0xe5, 0x47, 0x5c, 0xac, 0xf6, 0x4c, 0xad, 0xb7, 0xe5, 0x99, 0xe9, 0xa1, 0xbb, 0x67, 0xc3,
	0x82, 0x72, 0xe1, 0x15, 0x78, 0x13, 0xde, 0x82, 0x33, 0x37, 0xc4, 0x91, 0x07, 0x41, 0xfd, 0x33,
	0xc2, 0x1e, 0x1b, 0x01, 0x7b, 0x44, 0x39, 0x56, 0xd5, 0x57, 0x5f, 0x7d, 0x5d, 0x55, 0x3d, 0xd3,
	0xf0, 0x16, 0x95, 0x31, 0x7e, 0xa7, 0x50, 0x14, 0x34, 0x8b, 0x69, 0xc9, 0xe2, 0x17, 0x54, 0x14,
	0xac, 0x98, 0x47, 0xa5, 0xe0, 0x8a, 0x13, 0x8f, 0x96, 0x2c, 0x78, 0x73, 0xce, 0xf9, 0x3c, 0x43,
	0x13, 0xa6, 0x45, 0xc1, 0x15, 0x55, 0x8c, 0x17, 0xd2, 0x42, 0x82, 0x23, 0x17, 0x35, 0xd6, 0xac,
	0x3a, 0x8f, 0x15, 0xcb, 0x51, 0x2a, 0x9a, 0x97, 0x0e, 0xf0, 0x46, 0x13, 0x80, 0x79, 0xa9, 0xae,
	0x6c, 0x30, 0xfc, 0xdd, 0x83, 0x9d, 0xaf, 0x6c, 0x49, 0xd2, 0x87, 0x36, 0x4b, 0xfd, 0xd6, 0xa8,
	0x35, 0xf6, 0x26, 0x6d, 0x96, 0x92, 0x03, 0xe8, 0xd0, 0x4a, 0x5d, 0x70, 0xe1, 0xb7, 0x47, 0xad,
	0xf1, 0xee, 0xc4, 0x59, 0xda, 0x2f, 0x30, 0xa7, 0x62, 0xe1, 0x7b, 0xd6, 0x6f, 0x2d, 0xed, 0x97,
	0x8a, 0xaa, 0x4a, 0xfa, 0x5b, 0xa3, 0xd6, 0x78, 0x7b, 0xe2, 0x2c, 0x72, 0x0f, 0xf6, 0xb9, 0x98,
	0xd3, 0x82, 0x7d, 0x6f, 0x84, 0x4f, 0x59, 0xea, 0x6f, 0x9b, 0x22, 0xfd, 0x65, 0xf7, 0x69, 0x4a,
	0xde, 0x85, 0xc1, 0x0a, 0xb0, 0xa0, 0x39, 0xfa, 0x1d, 0x53, 0xe3, 0xc6, 0x72, 0xe0, 0x29, 0xcd,
	0x91, 0x1c, 0x41, 0x37, 0xc5, 0x4b, 0x96, 0xe0, 0x54, 0x5d, 0x95, 0xe8, 0xef, 0x98, 0x92, 0x60,
	0x5d, 0x67, 0x57, 0x25, 0x92, 0x3b, 0xb0, 0x27, 0x51, 0x30, 0x9a, 0x4d, 0x8b, 0x2a, 0x9f, 0xa1,
	0xf0, 0x5f, 0x33, 0x4c, 0x3d, 0xeb, 0x7c, 0x6a, 0x7c, 0x64, 0x08, 0xdb, 0x09, 0x4f, 0x51, 0xfa,
	0xbb, 0x26, 0x68, 0x0d, 0x72, 0x1b, 0x7a, 0x6e, 0x0e, 0x53, 0xdd, 0x4d, 0x1f, 0x8c, 0xdc, 0xae,
	0xf3, 0x9d, 0xb1, 0x1c, 0x75, 0xa2, 0x62, 0x2a, 0x43, 0xbf, 0x6b, 0x13, 0x8d, 0x41, 0x46, 0x5a,
	0x94, 0x4c, 0x04, 0x2b, 0xb5, 0x4e, 0xbf, 0x67, 0x62, 0xcb, 0x2e, 0xad, 0xaa, 0xa6, 0xce, 0xf0,
	0x12, 0x33, 0x7f, 0xcf, 0x08, 0xaf, 0xeb, 0x3d, 0xd1, 0xbe, 0x65, 0x50, 0xc2, 0xab, 0x42, 0xf9,
	0xfd, 0x15, 0xd0, 0x23, 0xed, 0x23, 0x21, 0xec, 0xbd, 0xe0, 0x62, 0x31, 0xe5, 0x22, 0x45, 0xa1,
	0x9b, 0xba, 0xef, 0x54, 0x72, 0xb1, 0x78, 0xa6, 0x7d, 0xa7, 0x69, 0xf8, 0xdb, 0x16, 0xec, 0xbb,
	0xf1, 0x3e, 0x61, 0x52, 0x9d, 0x2a, 0xcc, 0x5f, 0x8d, 0xf9, 0x7f, 0x37, 0x66, 0xf2, 0x21, 0x40,
	0x22, 0x90, 0x2a, 0x4c, 0xa7, 0x54, 0xf9, 0x37, 0x46, 0xad, 0x71, 0xf7, 0x38, 0x88, 0xec, 0xbd,
	0x8f, 0xea, 0x7b, 0x1f, 0x9d, 0xd5, 0x1f, 0x86, 0xc9, 0xae, 0x43, 0x3f, 0x54, 0x3a, 0xb5, 0x2a,
	0xd3, 0x3a, 0x75, 0xf0, 0xcf, 0xa9, 0x0e, 0xfd, 0x50, 0x85, 0x1f, 0xc3, 0xf0, 0x91, 0xe1, 0x71,
	0x1b, 0x36, 0xc1, 0x6f, 0x2b, 0x94, 0x8a, 0xdc, 0x85, 0x1d, 0x77, 0x02, 0xb3, 0x65, 0xdd, 0xe3,
	0x5e, 0x44, 0x4b, 0x16, 0xd5, 0xa8, 0x3a, 0x18, 0xde, 0x83, 0xd7, 0x1b, 0xf9, 0xb2, 0xe4, 0x85,
	0xc4, 0xe6, 0x86, 0x86, 0x77, 0x60, 0xf0, 0x18, 0x55, 0xa3, 0x4a, 0x13, 0xf4, 0x73, 0x0b, 0xc8,
	0x32, 0xca, 0x71, 0xfd, 0x4b, 0x31, 0x8d, 0x16, 0xb6, 0xaf, 0xdf, 0x42, 0xef, 0x3f, 0xb6, 0xf0,
	0x0b, 0x63, 0x5c, 0xb3, 0x85, 0x77, 0x61, 0xf8, 0x09, 0x66, 0xb8, 0x96, 0xdf, 0x6c, 0xce, 0x09,
	0x10, 0x7d, 0xff, 0x1b, 0xa8, 0x21, 0x6c, 0x67, 0x2c, 0x67, 0xca, 0x01, 0xad, 0xa1, 0xef, 0x37,
	0x3f, 0x3f, 0x97, 0x68, 0xbb, 0xe0, 0x4d, 0x9c, 0x15, 0xa6, 0x70, 0x73, 0x85, 0xc3, 0x35, 0xf8,
	0x08, 0xba, 0x8a, 0x2b, 0x9a, 0xb9, 0x15, 0xb6, 0x54, 0x60, 0x5c, 0x76, 0x81, 0xdf, 0xd3, 0xdf,
	0x11, 0x59, 0x65, 0x9a, 0xcf, 0x1b, 0x77, 0x8f, 0x87, 0xcb, 0x47, 0xa9, 0xbf, 0x4a, 0x13, 0x87,
	0x39, 0xfe, 0xc5, 0x83, 0xbe, 0x8b, 0x7d, 0x8e, 0x42, 0xdf, 0x72, 0xf2, 0x35, 0x74, 0xec, 0x9e,
	0x90, 0x43, 0x93, 0xba, 0x69, 0xe9, 0x82, 0x60, 0x53, 0xc8, 0x4a, 0x0c, 0x6f, 0xfd, 0xf8, 0xeb,
	0x1f, 0x3f, 0xb5, 0x07, 0x61, 0x6f, 0xf9, 0x0f, 0xfb, 0x51, 0xeb, 0x3e, 0xf9, 0x0c, 0xbc, 0xc7,
	0xa8, 0xc8, 0x81, 0xc9, 0x5d, 0x5b, 0xb1, 0xe0, 0xd6, 0x9a, 0xdf, 0x11, 0x1e, 0x1a, 0xc2, 0x9b,
	0x64, 0xb0, 0x4c, 0x18, 0xff, 0xc0, 0xd2, 0x97, 0x84, 0x42, 0xc7, 0x4e, 0xd4, 0x89, 0xdd, 0x34,
	0xde, 0xe0, 0x60, 0x6d, 0x3b, 0x3e, 0xd5, 0xff, 0xe4, 0xf0, 0x1d, 0xc3, 0xfb, 0x76, 0x70, 0xb8,
	0xca, 0x5b, 0xbf, 0x09, 0x58, 0xfa, 0x52, 0xab, 0xfe, 0x12, 0x3a, 0x76, 0xe8, 0xae, 0xc4, 0xa6,
	0x0d, 0xf8, 0xdb, 0x12, 0x4e, 0xfa, 0xfd, 0x0d, 0xd2, 0x9f, 0xc1, 0x96, 0x1e, 0x07, 0xb1, 0xc7,
	0x5e, 0xdf, 0x97, 0xc0, 0x5f, 0x0f, 0xb8, 0x86, 0x0c, 0x0d, 0x6b, 0x9f, 0xac, 0x74, 0xf8, 0x24,
	0x83, 0xdb, 0x8c, 0x47, 0xc9, 0x05, 0x13, 0xa5, 0x54, 0x34, 0x59, 0x98, 0x74, 0x2a, 0xa3, 0xfa,
	0xc1, 0xa3, 0xed, 0x93, 0x9e, 0xe3, 0x7a, 0xae, 0x75, 0x3e, 0x6f, 0x7d, 0xf3, 0xc1, 0x9c, 0xa9,
	0x8b, 0x6a, 0x16, 0x25, 0x3c, 0x8f, 0x67, 0x82, 0x27, 0x94, 0x8a, 0xf8, 0x2f, 0x8a, 0xf7, 0x35,
	0xff, 0x9c, 0xc7, 0x97, 0x0f, 0xe2, 0xc6, 0xcb, 0x69, 0xd6, 0x31, 0x27, 0x7d, 0xf0, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3d, 0x96, 0xe9, 0x53, 0x53, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WarningServiceClient is the client API for WarningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WarningServiceClient interface {
	// Create creates warning.
	Create(ctx context.Context, in *CreateWarningRequest, opts ...grpc.CallOption) (*CreateWarningResponse, error)
	// Get returns the warning matching the given id.
	Get(ctx context.Context, in *GetWarningRequest, opts ...grpc.CallOption) (*GetWarningResponse, error)
	// Update updates the given warning.
	Update(ctx context.Context, in *UpdateWarningRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the warning matching the given id.
	Delete(ctx context.Context, in *DeleteWarningRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available warning.
	List(ctx context.Context, in *ListWarningRequest, opts ...grpc.CallOption) (*ListWarningResponse, error)
}

type warningServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWarningServiceClient(cc grpc.ClientConnInterface) WarningServiceClient {
	return &warningServiceClient{cc}
}

func (c *warningServiceClient) Create(ctx context.Context, in *CreateWarningRequest, opts ...grpc.CallOption) (*CreateWarningResponse, error) {
	out := new(CreateWarningResponse)
	err := c.cc.Invoke(ctx, "/api.WarningService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warningServiceClient) Get(ctx context.Context, in *GetWarningRequest, opts ...grpc.CallOption) (*GetWarningResponse, error) {
	out := new(GetWarningResponse)
	err := c.cc.Invoke(ctx, "/api.WarningService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warningServiceClient) Update(ctx context.Context, in *UpdateWarningRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.WarningService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warningServiceClient) Delete(ctx context.Context, in *DeleteWarningRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.WarningService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warningServiceClient) List(ctx context.Context, in *ListWarningRequest, opts ...grpc.CallOption) (*ListWarningResponse, error) {
	out := new(ListWarningResponse)
	err := c.cc.Invoke(ctx, "/api.WarningService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarningServiceServer is the server API for WarningService service.
type WarningServiceServer interface {
	// Create creates warning.
	Create(context.Context, *CreateWarningRequest) (*CreateWarningResponse, error)
	// Get returns the warning matching the given id.
	Get(context.Context, *GetWarningRequest) (*GetWarningResponse, error)
	// Update updates the given warning.
	Update(context.Context, *UpdateWarningRequest) (*empty.Empty, error)
	// Delete deletes the warning matching the given id.
	Delete(context.Context, *DeleteWarningRequest) (*empty.Empty, error)
	// List lists the available warning.
	List(context.Context, *ListWarningRequest) (*ListWarningResponse, error)
}

// UnimplementedWarningServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWarningServiceServer struct {
}

func (*UnimplementedWarningServiceServer) Create(ctx context.Context, req *CreateWarningRequest) (*CreateWarningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedWarningServiceServer) Get(ctx context.Context, req *GetWarningRequest) (*GetWarningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedWarningServiceServer) Update(ctx context.Context, req *UpdateWarningRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedWarningServiceServer) Delete(ctx context.Context, req *DeleteWarningRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedWarningServiceServer) List(ctx context.Context, req *ListWarningRequest) (*ListWarningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterWarningServiceServer(s *grpc.Server, srv WarningServiceServer) {
	s.RegisterService(&_WarningService_serviceDesc, srv)
}

func _WarningService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWarningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningServiceServer).Create(ctx, req.(*CreateWarningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarningService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWarningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningServiceServer).Get(ctx, req.(*GetWarningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarningService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWarningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningServiceServer).Update(ctx, req.(*UpdateWarningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarningService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWarningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningServiceServer).Delete(ctx, req.(*DeleteWarningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarningService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWarningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningServiceServer).List(ctx, req.(*ListWarningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WarningService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WarningService",
	HandlerType: (*WarningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WarningService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WarningService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _WarningService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WarningService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WarningService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/warning.proto",
}
