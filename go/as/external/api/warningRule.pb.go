// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/warningRule.proto

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

type WarningRule struct {
	// WarningRule ID.
	// Will be set automatically on create.
	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// Optional remark to store with the warning-rule.
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Status               int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	RuleName             string   `protobuf:"bytes,5,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	RuleParam            string   `protobuf:"bytes,7,opt,name=rule_param,json=ruleParam,proto3" json:"rule_param,omitempty"`
	MonitorControl       string   `protobuf:"bytes,8,opt,name=monitor_control,json=monitorControl,proto3" json:"monitor_control,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WarningRule) Reset()         { *m = WarningRule{} }
func (m *WarningRule) String() string { return proto.CompactTextString(m) }
func (*WarningRule) ProtoMessage()    {}
func (*WarningRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{0}
}

func (m *WarningRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WarningRule.Unmarshal(m, b)
}
func (m *WarningRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WarningRule.Marshal(b, m, deterministic)
}
func (m *WarningRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WarningRule.Merge(m, src)
}
func (m *WarningRule) XXX_Size() int {
	return xxx_messageInfo_WarningRule.Size(m)
}
func (m *WarningRule) XXX_DiscardUnknown() {
	xxx_messageInfo_WarningRule.DiscardUnknown(m)
}

var xxx_messageInfo_WarningRule proto.InternalMessageInfo

func (m *WarningRule) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WarningRule) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *WarningRule) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *WarningRule) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *WarningRule) GetRuleName() string {
	if m != nil {
		return m.RuleName
	}
	return ""
}

func (m *WarningRule) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WarningRule) GetRuleParam() string {
	if m != nil {
		return m.RuleParam
	}
	return ""
}

func (m *WarningRule) GetMonitorControl() string {
	if m != nil {
		return m.MonitorControl
	}
	return ""
}

type WarningRuleListItem struct {
	// User ID.
	// Will be set automatically on create.
	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// remark of the warning-rule.
	Remark         string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Status         int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	RuleName       string `protobuf:"bytes,5,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	Description    string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	RuleParam      string `protobuf:"bytes,7,opt,name=rule_param,json=ruleParam,proto3" json:"rule_param,omitempty"`
	MonitorControl string `protobuf:"bytes,8,opt,name=monitor_control,json=monitorControl,proto3" json:"monitor_control,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WarningRuleListItem) Reset()         { *m = WarningRuleListItem{} }
func (m *WarningRuleListItem) String() string { return proto.CompactTextString(m) }
func (*WarningRuleListItem) ProtoMessage()    {}
func (*WarningRuleListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{1}
}

func (m *WarningRuleListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WarningRuleListItem.Unmarshal(m, b)
}
func (m *WarningRuleListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WarningRuleListItem.Marshal(b, m, deterministic)
}
func (m *WarningRuleListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WarningRuleListItem.Merge(m, src)
}
func (m *WarningRuleListItem) XXX_Size() int {
	return xxx_messageInfo_WarningRuleListItem.Size(m)
}
func (m *WarningRuleListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_WarningRuleListItem.DiscardUnknown(m)
}

var xxx_messageInfo_WarningRuleListItem proto.InternalMessageInfo

func (m *WarningRuleListItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WarningRuleListItem) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *WarningRuleListItem) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *WarningRuleListItem) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *WarningRuleListItem) GetRuleName() string {
	if m != nil {
		return m.RuleName
	}
	return ""
}

func (m *WarningRuleListItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WarningRuleListItem) GetRuleParam() string {
	if m != nil {
		return m.RuleParam
	}
	return ""
}

func (m *WarningRuleListItem) GetMonitorControl() string {
	if m != nil {
		return m.MonitorControl
	}
	return ""
}

func (m *WarningRuleListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *WarningRuleListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateWarningRuleRequest struct {
	// WarningRule object to create.
	WarningRule          *WarningRule `protobuf:"bytes,1,opt,name=warning_rule,json=warningRule,proto3" json:"warning_rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateWarningRuleRequest) Reset()         { *m = CreateWarningRuleRequest{} }
func (m *CreateWarningRuleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWarningRuleRequest) ProtoMessage()    {}
func (*CreateWarningRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{2}
}

func (m *CreateWarningRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWarningRuleRequest.Unmarshal(m, b)
}
func (m *CreateWarningRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWarningRuleRequest.Marshal(b, m, deterministic)
}
func (m *CreateWarningRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWarningRuleRequest.Merge(m, src)
}
func (m *CreateWarningRuleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWarningRuleRequest.Size(m)
}
func (m *CreateWarningRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWarningRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWarningRuleRequest proto.InternalMessageInfo

func (m *CreateWarningRuleRequest) GetWarningRule() *WarningRule {
	if m != nil {
		return m.WarningRule
	}
	return nil
}

type CreateWarningRuleResponse struct {
	// WarningRule ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWarningRuleResponse) Reset()         { *m = CreateWarningRuleResponse{} }
func (m *CreateWarningRuleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateWarningRuleResponse) ProtoMessage()    {}
func (*CreateWarningRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{3}
}

func (m *CreateWarningRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWarningRuleResponse.Unmarshal(m, b)
}
func (m *CreateWarningRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWarningRuleResponse.Marshal(b, m, deterministic)
}
func (m *CreateWarningRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWarningRuleResponse.Merge(m, src)
}
func (m *CreateWarningRuleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateWarningRuleResponse.Size(m)
}
func (m *CreateWarningRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWarningRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWarningRuleResponse proto.InternalMessageInfo

func (m *CreateWarningRuleResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetWarningRuleRequest struct {
	// WarningRule ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWarningRuleRequest) Reset()         { *m = GetWarningRuleRequest{} }
func (m *GetWarningRuleRequest) String() string { return proto.CompactTextString(m) }
func (*GetWarningRuleRequest) ProtoMessage()    {}
func (*GetWarningRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{4}
}

func (m *GetWarningRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWarningRuleRequest.Unmarshal(m, b)
}
func (m *GetWarningRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWarningRuleRequest.Marshal(b, m, deterministic)
}
func (m *GetWarningRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWarningRuleRequest.Merge(m, src)
}
func (m *GetWarningRuleRequest) XXX_Size() int {
	return xxx_messageInfo_GetWarningRuleRequest.Size(m)
}
func (m *GetWarningRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWarningRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWarningRuleRequest proto.InternalMessageInfo

func (m *GetWarningRuleRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetWarningRuleResponse struct {
	// WarningRule object.
	WarningRule *WarningRule `protobuf:"bytes,1,opt,name=warning_rule,json=warningRule,proto3" json:"warning_rule,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetWarningRuleResponse) Reset()         { *m = GetWarningRuleResponse{} }
func (m *GetWarningRuleResponse) String() string { return proto.CompactTextString(m) }
func (*GetWarningRuleResponse) ProtoMessage()    {}
func (*GetWarningRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{5}
}

func (m *GetWarningRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWarningRuleResponse.Unmarshal(m, b)
}
func (m *GetWarningRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWarningRuleResponse.Marshal(b, m, deterministic)
}
func (m *GetWarningRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWarningRuleResponse.Merge(m, src)
}
func (m *GetWarningRuleResponse) XXX_Size() int {
	return xxx_messageInfo_GetWarningRuleResponse.Size(m)
}
func (m *GetWarningRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWarningRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWarningRuleResponse proto.InternalMessageInfo

func (m *GetWarningRuleResponse) GetWarningRule() *WarningRule {
	if m != nil {
		return m.WarningRule
	}
	return nil
}

func (m *GetWarningRuleResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetWarningRuleResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateWarningRuleRequest struct {
	// WarningRule object to update.
	WarningRule          *WarningRule `protobuf:"bytes,1,opt,name=warning_rule,json=warningRule,proto3" json:"warning_rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateWarningRuleRequest) Reset()         { *m = UpdateWarningRuleRequest{} }
func (m *UpdateWarningRuleRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateWarningRuleRequest) ProtoMessage()    {}
func (*UpdateWarningRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{6}
}

func (m *UpdateWarningRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWarningRuleRequest.Unmarshal(m, b)
}
func (m *UpdateWarningRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWarningRuleRequest.Marshal(b, m, deterministic)
}
func (m *UpdateWarningRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWarningRuleRequest.Merge(m, src)
}
func (m *UpdateWarningRuleRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateWarningRuleRequest.Size(m)
}
func (m *UpdateWarningRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWarningRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWarningRuleRequest proto.InternalMessageInfo

func (m *UpdateWarningRuleRequest) GetWarningRule() *WarningRule {
	if m != nil {
		return m.WarningRule
	}
	return nil
}

type DeleteWarningRuleRequest struct {
	// WarningRule ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteWarningRuleRequest) Reset()         { *m = DeleteWarningRuleRequest{} }
func (m *DeleteWarningRuleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteWarningRuleRequest) ProtoMessage()    {}
func (*DeleteWarningRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{7}
}

func (m *DeleteWarningRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteWarningRuleRequest.Unmarshal(m, b)
}
func (m *DeleteWarningRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteWarningRuleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteWarningRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteWarningRuleRequest.Merge(m, src)
}
func (m *DeleteWarningRuleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteWarningRuleRequest.Size(m)
}
func (m *DeleteWarningRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteWarningRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteWarningRuleRequest proto.InternalMessageInfo

func (m *DeleteWarningRuleRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListWarningRuleRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWarningRuleRequest) Reset()         { *m = ListWarningRuleRequest{} }
func (m *ListWarningRuleRequest) String() string { return proto.CompactTextString(m) }
func (*ListWarningRuleRequest) ProtoMessage()    {}
func (*ListWarningRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{8}
}

func (m *ListWarningRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWarningRuleRequest.Unmarshal(m, b)
}
func (m *ListWarningRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWarningRuleRequest.Marshal(b, m, deterministic)
}
func (m *ListWarningRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWarningRuleRequest.Merge(m, src)
}
func (m *ListWarningRuleRequest) XXX_Size() int {
	return xxx_messageInfo_ListWarningRuleRequest.Size(m)
}
func (m *ListWarningRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWarningRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWarningRuleRequest proto.InternalMessageInfo

func (m *ListWarningRuleRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListWarningRuleRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListWarningRuleResponse struct {
	// Total number of warning-rule.
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// WarningRule within the result-set.
	Result               []*WarningRuleListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListWarningRuleResponse) Reset()         { *m = ListWarningRuleResponse{} }
func (m *ListWarningRuleResponse) String() string { return proto.CompactTextString(m) }
func (*ListWarningRuleResponse) ProtoMessage()    {}
func (*ListWarningRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3c796c56f82e6ee, []int{9}
}

func (m *ListWarningRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWarningRuleResponse.Unmarshal(m, b)
}
func (m *ListWarningRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWarningRuleResponse.Marshal(b, m, deterministic)
}
func (m *ListWarningRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWarningRuleResponse.Merge(m, src)
}
func (m *ListWarningRuleResponse) XXX_Size() int {
	return xxx_messageInfo_ListWarningRuleResponse.Size(m)
}
func (m *ListWarningRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWarningRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWarningRuleResponse proto.InternalMessageInfo

func (m *ListWarningRuleResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListWarningRuleResponse) GetResult() []*WarningRuleListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*WarningRule)(nil), "api.WarningRule")
	proto.RegisterType((*WarningRuleListItem)(nil), "api.WarningRuleListItem")
	proto.RegisterType((*CreateWarningRuleRequest)(nil), "api.CreateWarningRuleRequest")
	proto.RegisterType((*CreateWarningRuleResponse)(nil), "api.CreateWarningRuleResponse")
	proto.RegisterType((*GetWarningRuleRequest)(nil), "api.GetWarningRuleRequest")
	proto.RegisterType((*GetWarningRuleResponse)(nil), "api.GetWarningRuleResponse")
	proto.RegisterType((*UpdateWarningRuleRequest)(nil), "api.UpdateWarningRuleRequest")
	proto.RegisterType((*DeleteWarningRuleRequest)(nil), "api.DeleteWarningRuleRequest")
	proto.RegisterType((*ListWarningRuleRequest)(nil), "api.ListWarningRuleRequest")
	proto.RegisterType((*ListWarningRuleResponse)(nil), "api.ListWarningRuleResponse")
}

func init() {
	proto.RegisterFile("as/external/api/warningRule.proto", fileDescriptor_e3c796c56f82e6ee)
}

var fileDescriptor_e3c796c56f82e6ee = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0x93, 0x36, 0xb7, 0x39, 0xb9, 0xea, 0xed, 0x9d, 0xde, 0x1b, 0x5c, 0xf7, 0x2f, 0x35,
	0x8b, 0x46, 0x45, 0xb5, 0x51, 0xba, 0x40, 0xb0, 0x6b, 0x0b, 0x54, 0x48, 0x08, 0x2a, 0x03, 0x42,
	0x42, 0x88, 0x68, 0x62, 0x4f, 0xd3, 0x51, 0x6d, 0x8f, 0x99, 0x19, 0x17, 0x10, 0xea, 0x86, 0x2d,
	0x4b, 0x1e, 0x89, 0x47, 0xe0, 0x15, 0x58, 0x21, 0x1e, 0x02, 0xcd, 0x78, 0xa2, 0x5a, 0x8e, 0x23,
	0x44, 0xc5, 0x8a, 0xe5, 0x39, 0xe7, 0x9b, 0xef, 0xf8, 0x7c, 0xf3, 0x1d, 0x0f, 0x6c, 0x61, 0xe1,
	0x93, 0xb7, 0x92, 0xf0, 0x14, 0xc7, 0x3e, 0xce, 0xa8, 0xff, 0x06, 0xf3, 0x94, 0xa6, 0xe3, 0x20,
	0x8f, 0x89, 0x97, 0x71, 0x26, 0x19, 0x6a, 0xe2, 0x8c, 0x3a, 0x6b, 0x63, 0xc6, 0xc6, 0x31, 0xd1,
	0x10, 0x9c, 0xa6, 0x4c, 0x62, 0x49, 0x59, 0x2a, 0x0a, 0x88, 0xb3, 0x69, 0xaa, 0x3a, 0x1a, 0xe5,
	0x27, 0xbe, 0xa4, 0x09, 0x11, 0x12, 0x27, 0x99, 0x01, 0xac, 0x56, 0x01, 0x24, 0xc9, 0xe4, 0xbb,
	0xa2, 0xe8, 0x7e, 0xb7, 0xa0, 0xf3, 0xfc, 0xb2, 0x2d, 0x5a, 0x84, 0x06, 0x8d, 0x6c, 0xab, 0x67,
	0xf5, 0x9b, 0x41, 0x83, 0x46, 0xa8, 0x0b, 0x2d, 0x9c, 0xcb, 0x53, 0xc6, 0xed, 0x46, 0xcf, 0xea,
	0xb7, 0x03, 0x13, 0xa9, 0x3c, 0x27, 0x09, 0xe6, 0x67, 0x76, 0xb3, 0xc8, 0x17, 0x91, 0xca, 0x0b,
	0x89, 0x65, 0x2e, 0xec, 0xb9, 0x9e, 0xd5, 0x9f, 0x0f, 0x4c, 0x84, 0x56, 0xa1, 0xcd, 0xf3, 0x98,
	0x0c, 0x53, 0x9c, 0x10, 0x7b, 0x5e, 0x1f, 0x59, 0x50, 0x89, 0x47, 0x38, 0x21, 0xa8, 0x07, 0x9d,
	0x88, 0x88, 0x90, 0xd3, 0x4c, 0x0d, 0x66, 0xb7, 0x74, 0xb9, 0x9c, 0x42, 0xeb, 0x00, 0xfa, 0x78,
	0x86, 0x39, 0x4e, 0xec, 0xbf, 0x34, 0x40, 0x13, 0x1e, 0xab, 0x04, 0xda, 0x86, 0x7f, 0x12, 0x96,
	0x52, 0xc9, 0xf8, 0x30, 0x64, 0xa9, 0xe4, 0x2c, 0xb6, 0x17, 0x34, 0x66, 0xd1, 0xa4, 0x0f, 0x8b,
	0xac, 0xfb, 0xad, 0x01, 0xcb, 0xa5, 0x71, 0x1f, 0x52, 0x21, 0x1f, 0x48, 0x92, 0xfc, 0xe1, 0x63,
	0xa3, 0xdb, 0x00, 0x21, 0x27, 0x58, 0x92, 0x68, 0x88, 0xa5, 0xdd, 0xee, 0x59, 0xfd, 0xce, 0xc0,
	0xf1, 0x0a, 0x5f, 0x78, 0x13, 0x5f, 0x78, 0x4f, 0x27, 0xc6, 0x09, 0xda, 0x06, 0xbd, 0x2f, 0xd5,
	0xd1, 0x3c, 0x8b, 0x26, 0x47, 0xe1, 0xe7, 0x47, 0x0d, 0x7a, 0x5f, 0xba, 0x8f, 0xc1, 0x3e, 0xd4,
	0x3c, 0x25, 0xc5, 0x03, 0xf2, 0x3a, 0x27, 0x42, 0xa2, 0x3d, 0xf8, 0xdb, 0xb8, 0x7d, 0xa8, 0xe6,
	0xd1, 0xd2, 0x77, 0x06, 0x4b, 0x1e, 0xce, 0xa8, 0x57, 0x86, 0x77, 0x4a, 0x3b, 0xe1, 0xde, 0x80,
	0x95, 0x1a, 0x42, 0x91, 0xb1, 0x54, 0x4c, 0x39, 0xd7, 0xdd, 0x86, 0xff, 0x8f, 0x88, 0xac, 0x69,
	0x5d, 0x05, 0x7e, 0xb6, 0xa0, 0x5b, 0x45, 0x1a, 0xce, 0xab, 0x7c, 0x65, 0x45, 0xec, 0xc6, 0xd5,
	0xc5, 0x6e, 0xfe, 0xa2, 0xd8, 0xcf, 0x74, 0xf0, 0xbb, 0xc4, 0xde, 0x01, 0xfb, 0x2e, 0x89, 0x49,
	0x2d, 0x61, 0x55, 0xc2, 0xfb, 0xd0, 0x55, 0xab, 0x54, 0x83, 0xfc, 0x0f, 0xe6, 0x63, 0x9a, 0x50,
	0x69, 0xc0, 0x45, 0xa0, 0xd6, 0x85, 0x9d, 0x9c, 0x08, 0x52, 0xc8, 0xd3, 0x0c, 0x4c, 0xe4, 0xc6,
	0x70, 0x6d, 0x8a, 0xc7, 0x5c, 0xc5, 0x26, 0x74, 0x24, 0x93, 0x38, 0x1e, 0x86, 0x2c, 0x4f, 0x27,
	0x74, 0xa0, 0x53, 0x87, 0x2a, 0x83, 0x6e, 0xaa, 0xd5, 0x14, 0x79, 0xac, 0x38, 0x9b, 0xfd, 0xce,
	0xc0, 0xae, 0x8e, 0x37, 0x59, 0xf6, 0xc0, 0xe0, 0x06, 0x1f, 0xe7, 0x00, 0x95, 0xea, 0x4f, 0x08,
	0x3f, 0xa7, 0x21, 0x41, 0x04, 0x5a, 0x85, 0xcb, 0xd0, 0xba, 0xa6, 0x98, 0xe5, 0x61, 0x67, 0x63,
	0x56, 0xb9, 0xf8, 0x64, 0x77, 0xed, 0xc3, 0x97, 0xaf, 0x9f, 0x1a, 0x5d, 0xf7, 0xdf, 0xf2, 0xcf,
	0x7d, 0x57, 0xdd, 0xc0, 0x1d, 0x6b, 0x07, 0xbd, 0x82, 0xe6, 0x11, 0x91, 0xc8, 0xd1, 0x24, 0xb5,
	0x4e, 0x75, 0x56, 0x6b, 0x6b, 0x86, 0x7d, 0x43, 0xb3, 0xdb, 0xa8, 0x3b, 0xc5, 0xee, 0xbf, 0xa7,
	0xd1, 0x05, 0x62, 0xd0, 0x2a, 0x0c, 0x61, 0xc6, 0x98, 0xe5, 0x0e, 0xa7, 0x3b, 0x65, 0xb0, 0x7b,
	0xea, 0x81, 0x70, 0x3d, 0xdd, 0xa0, 0xef, 0x5c, 0xaf, 0x69, 0x50, 0xb6, 0x93, 0x47, 0xa3, 0x0b,
	0x35, 0xd0, 0x10, 0x5a, 0x85, 0x61, 0x4c, 0xc3, 0x59, 0xee, 0x99, 0xd9, 0xd0, 0x4c, 0xb4, 0x33,
	0x6b, 0xa2, 0x97, 0x30, 0xa7, 0xee, 0x10, 0x15, 0xb2, 0xd4, 0x1b, 0xce, 0x59, 0xab, 0x2f, 0x1a,
	0xd1, 0x56, 0x74, 0x8b, 0x65, 0x34, 0x7d, 0x25, 0x07, 0x19, 0x6c, 0x51, 0xe6, 0x85, 0xa7, 0x94,
	0x67, 0x42, 0xe2, 0xf0, 0x4c, 0xf3, 0x60, 0xe1, 0x4d, 0x5e, 0x68, 0x15, 0x1f, 0x2c, 0x95, 0x48,
	0x8f, 0xd5, 0xd7, 0x1f, 0x5b, 0x2f, 0x6e, 0x8d, 0xa9, 0x3c, 0xcd, 0x47, 0x5e, 0xc8, 0x12, 0x7f,
	0xc4, 0x59, 0x88, 0x31, 0xf7, 0x2f, 0x69, 0x76, 0x55, 0xa3, 0x31, 0xf3, 0xcf, 0xf7, 0xfc, 0xca,
	0x73, 0x3f, 0x6a, 0xe9, 0xf9, 0xf7, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0xea, 0x71, 0xbb, 0xf5,
	0x08, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WarningRuleServiceClient is the client API for WarningRuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WarningRuleServiceClient interface {
	// Create creates warning-rule.
	Create(ctx context.Context, in *CreateWarningRuleRequest, opts ...grpc.CallOption) (*CreateWarningRuleResponse, error)
	// Get returns the warning-rule matching the given id.
	Get(ctx context.Context, in *GetWarningRuleRequest, opts ...grpc.CallOption) (*GetWarningRuleResponse, error)
	// Update updates the given warning-rule.
	Update(ctx context.Context, in *UpdateWarningRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the warning-rule matching the given id.
	Delete(ctx context.Context, in *DeleteWarningRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available warning-rule.
	List(ctx context.Context, in *ListWarningRuleRequest, opts ...grpc.CallOption) (*ListWarningRuleResponse, error)
}

type warningRuleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWarningRuleServiceClient(cc grpc.ClientConnInterface) WarningRuleServiceClient {
	return &warningRuleServiceClient{cc}
}

func (c *warningRuleServiceClient) Create(ctx context.Context, in *CreateWarningRuleRequest, opts ...grpc.CallOption) (*CreateWarningRuleResponse, error) {
	out := new(CreateWarningRuleResponse)
	err := c.cc.Invoke(ctx, "/api.WarningRuleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warningRuleServiceClient) Get(ctx context.Context, in *GetWarningRuleRequest, opts ...grpc.CallOption) (*GetWarningRuleResponse, error) {
	out := new(GetWarningRuleResponse)
	err := c.cc.Invoke(ctx, "/api.WarningRuleService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warningRuleServiceClient) Update(ctx context.Context, in *UpdateWarningRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.WarningRuleService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warningRuleServiceClient) Delete(ctx context.Context, in *DeleteWarningRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.WarningRuleService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warningRuleServiceClient) List(ctx context.Context, in *ListWarningRuleRequest, opts ...grpc.CallOption) (*ListWarningRuleResponse, error) {
	out := new(ListWarningRuleResponse)
	err := c.cc.Invoke(ctx, "/api.WarningRuleService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarningRuleServiceServer is the server API for WarningRuleService service.
type WarningRuleServiceServer interface {
	// Create creates warning-rule.
	Create(context.Context, *CreateWarningRuleRequest) (*CreateWarningRuleResponse, error)
	// Get returns the warning-rule matching the given id.
	Get(context.Context, *GetWarningRuleRequest) (*GetWarningRuleResponse, error)
	// Update updates the given warning-rule.
	Update(context.Context, *UpdateWarningRuleRequest) (*empty.Empty, error)
	// Delete deletes the warning-rule matching the given id.
	Delete(context.Context, *DeleteWarningRuleRequest) (*empty.Empty, error)
	// List lists the available warning-rule.
	List(context.Context, *ListWarningRuleRequest) (*ListWarningRuleResponse, error)
}

// UnimplementedWarningRuleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWarningRuleServiceServer struct {
}

func (*UnimplementedWarningRuleServiceServer) Create(ctx context.Context, req *CreateWarningRuleRequest) (*CreateWarningRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedWarningRuleServiceServer) Get(ctx context.Context, req *GetWarningRuleRequest) (*GetWarningRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedWarningRuleServiceServer) Update(ctx context.Context, req *UpdateWarningRuleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedWarningRuleServiceServer) Delete(ctx context.Context, req *DeleteWarningRuleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedWarningRuleServiceServer) List(ctx context.Context, req *ListWarningRuleRequest) (*ListWarningRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterWarningRuleServiceServer(s *grpc.Server, srv WarningRuleServiceServer) {
	s.RegisterService(&_WarningRuleService_serviceDesc, srv)
}

func _WarningRuleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWarningRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningRuleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningRuleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningRuleServiceServer).Create(ctx, req.(*CreateWarningRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarningRuleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWarningRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningRuleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningRuleService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningRuleServiceServer).Get(ctx, req.(*GetWarningRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarningRuleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWarningRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningRuleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningRuleService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningRuleServiceServer).Update(ctx, req.(*UpdateWarningRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarningRuleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWarningRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningRuleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningRuleService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningRuleServiceServer).Delete(ctx, req.(*DeleteWarningRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarningRuleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWarningRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarningRuleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WarningRuleService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarningRuleServiceServer).List(ctx, req.(*ListWarningRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WarningRuleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WarningRuleService",
	HandlerType: (*WarningRuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WarningRuleService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WarningRuleService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _WarningRuleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WarningRuleService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WarningRuleService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/warningRule.proto",
}