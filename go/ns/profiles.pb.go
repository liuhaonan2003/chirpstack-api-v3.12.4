// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ns/profiles.proto

package ns

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

type RatePolicy int32

const (
	// Drop
	RatePolicy_DROP RatePolicy = 0
	// Mark
	RatePolicy_MARK RatePolicy = 1
)

var RatePolicy_name = map[int32]string{
	0: "DROP",
	1: "MARK",
}

var RatePolicy_value = map[string]int32{
	"DROP": 0,
	"MARK": 1,
}

func (x RatePolicy) String() string {
	return proto.EnumName(RatePolicy_name, int32(x))
}

func (RatePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ee79be696a10f18b, []int{0}
}

type ServiceProfile struct {
	// Service-profile ID.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Token bucket filling rate, including ACKs (packet/h).
	UlRate uint32 `protobuf:"varint,2,opt,name=ul_rate,json=ulRate,proto3" json:"ul_rate,omitempty"`
	// Token bucket burst size.
	UlBucketSize uint32 `protobuf:"varint,3,opt,name=ul_bucket_size,json=ulBucketSize,proto3" json:"ul_bucket_size,omitempty"`
	// Drop or mark when exceeding ULRate.
	UlRatePolicy RatePolicy `protobuf:"varint,4,opt,name=ul_rate_policy,json=ulRatePolicy,proto3,enum=ns.RatePolicy" json:"ul_rate_policy,omitempty"`
	// Token bucket filling rate, including ACKs (packet/h).
	DlRate uint32 `protobuf:"varint,5,opt,name=dl_rate,json=dlRate,proto3" json:"dl_rate,omitempty"`
	// Token bucket burst size.
	DlBucketSize uint32 `protobuf:"varint,6,opt,name=dl_bucket_size,json=dlBucketSize,proto3" json:"dl_bucket_size,omitempty"`
	// Drop or mark when exceeding DLRate.
	DlRatePolicy RatePolicy `protobuf:"varint,7,opt,name=dl_rate_policy,json=dlRatePolicy,proto3,enum=ns.RatePolicy" json:"dl_rate_policy,omitempty"`
	// GW metadata (RSSI, SNR, GW geoloc., etc.) are added to the packet sent to AS.
	AddGwMetadata bool `protobuf:"varint,8,opt,name=add_gw_metadata,json=addGwMetadata,proto3" json:"add_gw_metadata,omitempty"`
	// Frequency to initiate an End-Device status request (request/day).
	DevStatusReqFreq uint32 `protobuf:"varint,9,opt,name=dev_status_req_freq,json=devStatusReqFreq,proto3" json:"dev_status_req_freq,omitempty"`
	// Report End-Device battery level to AS.
	ReportDevStatusBattery bool `protobuf:"varint,10,opt,name=report_dev_status_battery,json=reportDevStatusBattery,proto3" json:"report_dev_status_battery,omitempty"`
	// Report End-Device margin to AS.
	ReportDevStatusMargin bool `protobuf:"varint,11,opt,name=report_dev_status_margin,json=reportDevStatusMargin,proto3" json:"report_dev_status_margin,omitempty"`
	// Minimum allowed data rate. Used for ADR.
	DrMin uint32 `protobuf:"varint,12,opt,name=dr_min,json=drMin,proto3" json:"dr_min,omitempty"`
	// Maximum allowed data rate. Used for ADR.
	DrMax uint32 `protobuf:"varint,13,opt,name=dr_max,json=drMax,proto3" json:"dr_max,omitempty"`
	// Channel mask. sNS does not have to obey (i.e., informative).
	ChannelMask []byte `protobuf:"bytes,14,opt,name=channel_mask,json=channelMask,proto3" json:"channel_mask,omitempty"`
	// Passive Roaming allowed.
	PrAllowed bool `protobuf:"varint,15,opt,name=pr_allowed,json=prAllowed,proto3" json:"pr_allowed,omitempty"`
	// Handover Roaming allowed.
	HrAllowed bool `protobuf:"varint,16,opt,name=hr_allowed,json=hrAllowed,proto3" json:"hr_allowed,omitempty"`
	// Roaming Activation allowed.
	RaAllowed bool `protobuf:"varint,17,opt,name=ra_allowed,json=raAllowed,proto3" json:"ra_allowed,omitempty"`
	// Enable network geolocation service.
	NwkGeoLoc bool `protobuf:"varint,18,opt,name=nwk_geo_loc,json=nwkGeoLoc,proto3" json:"nwk_geo_loc,omitempty"`
	// Target Packet Error Rate.
	TargetPer uint32 `protobuf:"varint,19,opt,name=target_per,json=targetPer,proto3" json:"target_per,omitempty"`
	// Minimum number of receiving GWs (informative).
	MinGwDiversity uint32 `protobuf:"varint,20,opt,name=min_gw_diversity,json=minGwDiversity,proto3" json:"min_gw_diversity,omitempty"`
	// Gateways under this service-profile are private.
	// This means that these gateways can only be used by devices under the
	// same service-profile.
	GwsPrivate           bool     `protobuf:"varint,21,opt,name=gws_private,json=gwsPrivate,proto3" json:"gws_private,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceProfile) Reset()         { *m = ServiceProfile{} }
func (m *ServiceProfile) String() string { return proto.CompactTextString(m) }
func (*ServiceProfile) ProtoMessage()    {}
func (*ServiceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee79be696a10f18b, []int{0}
}

func (m *ServiceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceProfile.Unmarshal(m, b)
}
func (m *ServiceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceProfile.Marshal(b, m, deterministic)
}
func (m *ServiceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceProfile.Merge(m, src)
}
func (m *ServiceProfile) XXX_Size() int {
	return xxx_messageInfo_ServiceProfile.Size(m)
}
func (m *ServiceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceProfile proto.InternalMessageInfo

func (m *ServiceProfile) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ServiceProfile) GetUlRate() uint32 {
	if m != nil {
		return m.UlRate
	}
	return 0
}

func (m *ServiceProfile) GetUlBucketSize() uint32 {
	if m != nil {
		return m.UlBucketSize
	}
	return 0
}

func (m *ServiceProfile) GetUlRatePolicy() RatePolicy {
	if m != nil {
		return m.UlRatePolicy
	}
	return RatePolicy_DROP
}

func (m *ServiceProfile) GetDlRate() uint32 {
	if m != nil {
		return m.DlRate
	}
	return 0
}

func (m *ServiceProfile) GetDlBucketSize() uint32 {
	if m != nil {
		return m.DlBucketSize
	}
	return 0
}

func (m *ServiceProfile) GetDlRatePolicy() RatePolicy {
	if m != nil {
		return m.DlRatePolicy
	}
	return RatePolicy_DROP
}

func (m *ServiceProfile) GetAddGwMetadata() bool {
	if m != nil {
		return m.AddGwMetadata
	}
	return false
}

func (m *ServiceProfile) GetDevStatusReqFreq() uint32 {
	if m != nil {
		return m.DevStatusReqFreq
	}
	return 0
}

func (m *ServiceProfile) GetReportDevStatusBattery() bool {
	if m != nil {
		return m.ReportDevStatusBattery
	}
	return false
}

func (m *ServiceProfile) GetReportDevStatusMargin() bool {
	if m != nil {
		return m.ReportDevStatusMargin
	}
	return false
}

func (m *ServiceProfile) GetDrMin() uint32 {
	if m != nil {
		return m.DrMin
	}
	return 0
}

func (m *ServiceProfile) GetDrMax() uint32 {
	if m != nil {
		return m.DrMax
	}
	return 0
}

func (m *ServiceProfile) GetChannelMask() []byte {
	if m != nil {
		return m.ChannelMask
	}
	return nil
}

func (m *ServiceProfile) GetPrAllowed() bool {
	if m != nil {
		return m.PrAllowed
	}
	return false
}

func (m *ServiceProfile) GetHrAllowed() bool {
	if m != nil {
		return m.HrAllowed
	}
	return false
}

func (m *ServiceProfile) GetRaAllowed() bool {
	if m != nil {
		return m.RaAllowed
	}
	return false
}

func (m *ServiceProfile) GetNwkGeoLoc() bool {
	if m != nil {
		return m.NwkGeoLoc
	}
	return false
}

func (m *ServiceProfile) GetTargetPer() uint32 {
	if m != nil {
		return m.TargetPer
	}
	return 0
}

func (m *ServiceProfile) GetMinGwDiversity() uint32 {
	if m != nil {
		return m.MinGwDiversity
	}
	return 0
}

func (m *ServiceProfile) GetGwsPrivate() bool {
	if m != nil {
		return m.GwsPrivate
	}
	return false
}

type DeviceProfile struct {
	// Device-profile ID.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// End-Device supports Class B.
	SupportsClassB bool `protobuf:"varint,2,opt,name=supports_class_b,json=supportsClassB,proto3" json:"supports_class_b,omitempty"`
	// Maximum delay for the End-Device to answer a MAC request or a confirmed DL frame (mandatory if class B mode supported).
	ClassBTimeout uint32 `protobuf:"varint,3,opt,name=class_b_timeout,json=classBTimeout,proto3" json:"class_b_timeout,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotPeriod uint32 `protobuf:"varint,4,opt,name=ping_slot_period,json=pingSlotPeriod,proto3" json:"ping_slot_period,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotDr uint32 `protobuf:"varint,5,opt,name=ping_slot_dr,json=pingSlotDr,proto3" json:"ping_slot_dr,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotFreq uint32 `protobuf:"varint,6,opt,name=ping_slot_freq,json=pingSlotFreq,proto3" json:"ping_slot_freq,omitempty"`
	// End-Device supports Class C.
	SupportsClassC bool `protobuf:"varint,7,opt,name=supports_class_c,json=supportsClassC,proto3" json:"supports_class_c,omitempty"`
	// Maximum delay for the End-Device to answer a MAC request or a confirmed DL frame (mandatory if class C mode supported).
	ClassCTimeout uint32 `protobuf:"varint,8,opt,name=class_c_timeout,json=classCTimeout,proto3" json:"class_c_timeout,omitempty"`
	// Version of the LoRaWAN supported by the End-Device.
	MacVersion string `protobuf:"bytes,9,opt,name=mac_version,json=macVersion,proto3" json:"mac_version,omitempty"`
	// Revision of the Regional Parameters document supported by the End-Device.
	RegParamsRevision string `protobuf:"bytes,10,opt,name=reg_params_revision,json=regParamsRevision,proto3" json:"reg_params_revision,omitempty"`
	// Class A RX1 delay (mandatory for ABP).
	RxDelay_1 uint32 `protobuf:"varint,11,opt,name=rx_delay_1,json=rxDelay1,proto3" json:"rx_delay_1,omitempty"`
	// RX1 data rate offset (mandatory for ABP).
	RxDrOffset_1 uint32 `protobuf:"varint,12,opt,name=rx_dr_offset_1,json=rxDrOffset1,proto3" json:"rx_dr_offset_1,omitempty"`
	// RX2 data rate (mandatory for ABP).
	RxDatarate_2 uint32 `protobuf:"varint,13,opt,name=rx_datarate_2,json=rxDatarate2,proto3" json:"rx_datarate_2,omitempty"`
	// RX2 channel frequency (mandatory for ABP).
	RxFreq_2 uint32 `protobuf:"varint,14,opt,name=rx_freq_2,json=rxFreq2,proto3" json:"rx_freq_2,omitempty"`
	// List of factory-preset frequencies (mandatory for ABP).
	FactoryPresetFreqs []uint32 `protobuf:"varint,15,rep,packed,name=factory_preset_freqs,json=factoryPresetFreqs,proto3" json:"factory_preset_freqs,omitempty"`
	// Maximum EIRP supported by the End-Device.
	MaxEirp uint32 `protobuf:"varint,16,opt,name=max_eirp,json=maxEirp,proto3" json:"max_eirp,omitempty"`
	// Maximum duty cycle supported by the End-Device.
	MaxDutyCycle uint32 `protobuf:"varint,17,opt,name=max_duty_cycle,json=maxDutyCycle,proto3" json:"max_duty_cycle,omitempty"`
	// End-Device supports Join (OTAA) or not (ABP).
	SupportsJoin bool `protobuf:"varint,18,opt,name=supports_join,json=supportsJoin,proto3" json:"supports_join,omitempty"`
	// RF region name.
	RfRegion string `protobuf:"bytes,19,opt,name=rf_region,json=rfRegion,proto3" json:"rf_region,omitempty"`
	// End-Device uses 32bit FCnt (mandatory for LoRaWAN 1.0 End-Device).
	Supports_32BitFCnt bool `protobuf:"varint,20,opt,name=supports_32bit_f_cnt,json=supports32bitFCnt,proto3" json:"supports_32bit_f_cnt,omitempty"`
	// ADR algorithm ID.
	// In case this is left blank, or is configured to a non-existing ADR
	// algorithm (plugin), then it falls back to 'default'.
	AdrAlgorithmId string `protobuf:"bytes,21,opt,name=adr_algorithm_id,json=adrAlgorithmId,proto3" json:"adr_algorithm_id,omitempty"`
	// End-Device supports Class D.
	// class d模式下行会发长包唤醒,目的是拓展低功耗实时唤醒(上行classA,下行classC)
	SupportsClassD bool `protobuf:"varint,22,opt,name=supports_class_d,json=supportsClassD,proto3" json:"supports_class_d,omitempty"`
	// RF preamble size .
	PreambleSize         uint32   `protobuf:"varint,23,opt,name=preamble_size,json=preambleSize,proto3" json:"preamble_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProfile) Reset()         { *m = DeviceProfile{} }
func (m *DeviceProfile) String() string { return proto.CompactTextString(m) }
func (*DeviceProfile) ProtoMessage()    {}
func (*DeviceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee79be696a10f18b, []int{1}
}

func (m *DeviceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProfile.Unmarshal(m, b)
}
func (m *DeviceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProfile.Marshal(b, m, deterministic)
}
func (m *DeviceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProfile.Merge(m, src)
}
func (m *DeviceProfile) XXX_Size() int {
	return xxx_messageInfo_DeviceProfile.Size(m)
}
func (m *DeviceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProfile proto.InternalMessageInfo

func (m *DeviceProfile) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DeviceProfile) GetSupportsClassB() bool {
	if m != nil {
		return m.SupportsClassB
	}
	return false
}

func (m *DeviceProfile) GetClassBTimeout() uint32 {
	if m != nil {
		return m.ClassBTimeout
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotPeriod() uint32 {
	if m != nil {
		return m.PingSlotPeriod
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotDr() uint32 {
	if m != nil {
		return m.PingSlotDr
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotFreq() uint32 {
	if m != nil {
		return m.PingSlotFreq
	}
	return 0
}

func (m *DeviceProfile) GetSupportsClassC() bool {
	if m != nil {
		return m.SupportsClassC
	}
	return false
}

func (m *DeviceProfile) GetClassCTimeout() uint32 {
	if m != nil {
		return m.ClassCTimeout
	}
	return 0
}

func (m *DeviceProfile) GetMacVersion() string {
	if m != nil {
		return m.MacVersion
	}
	return ""
}

func (m *DeviceProfile) GetRegParamsRevision() string {
	if m != nil {
		return m.RegParamsRevision
	}
	return ""
}

func (m *DeviceProfile) GetRxDelay_1() uint32 {
	if m != nil {
		return m.RxDelay_1
	}
	return 0
}

func (m *DeviceProfile) GetRxDrOffset_1() uint32 {
	if m != nil {
		return m.RxDrOffset_1
	}
	return 0
}

func (m *DeviceProfile) GetRxDatarate_2() uint32 {
	if m != nil {
		return m.RxDatarate_2
	}
	return 0
}

func (m *DeviceProfile) GetRxFreq_2() uint32 {
	if m != nil {
		return m.RxFreq_2
	}
	return 0
}

func (m *DeviceProfile) GetFactoryPresetFreqs() []uint32 {
	if m != nil {
		return m.FactoryPresetFreqs
	}
	return nil
}

func (m *DeviceProfile) GetMaxEirp() uint32 {
	if m != nil {
		return m.MaxEirp
	}
	return 0
}

func (m *DeviceProfile) GetMaxDutyCycle() uint32 {
	if m != nil {
		return m.MaxDutyCycle
	}
	return 0
}

func (m *DeviceProfile) GetSupportsJoin() bool {
	if m != nil {
		return m.SupportsJoin
	}
	return false
}

func (m *DeviceProfile) GetRfRegion() string {
	if m != nil {
		return m.RfRegion
	}
	return ""
}

func (m *DeviceProfile) GetSupports_32BitFCnt() bool {
	if m != nil {
		return m.Supports_32BitFCnt
	}
	return false
}

func (m *DeviceProfile) GetAdrAlgorithmId() string {
	if m != nil {
		return m.AdrAlgorithmId
	}
	return ""
}

func (m *DeviceProfile) GetSupportsClassD() bool {
	if m != nil {
		return m.SupportsClassD
	}
	return false
}

func (m *DeviceProfile) GetPreambleSize() uint32 {
	if m != nil {
		return m.PreambleSize
	}
	return 0
}

type RoutingProfile struct {
	// ID of the routing profile.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Application-server ID.
	AsId string `protobuf:"bytes,2,opt,name=as_id,json=asId,proto3" json:"as_id,omitempty"`
	// CA certificate for connecting to the AS.
	CaCert string `protobuf:"bytes,3,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// TLS certificate for connecting to the AS.
	TlsCert string `protobuf:"bytes,4,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	// TLS key for connecting to the AS.
	// Note: when retrieving the routing-profile, the tls_key is not returned
	// for security reasons. When updating the routing-profile, an empty tls_key
	// does not clear the certificate, unless the tls_cert is also left blank.
	TlsKey               string   `protobuf:"bytes,5,opt,name=tls_key,json=tlsKey,proto3" json:"tls_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutingProfile) Reset()         { *m = RoutingProfile{} }
func (m *RoutingProfile) String() string { return proto.CompactTextString(m) }
func (*RoutingProfile) ProtoMessage()    {}
func (*RoutingProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee79be696a10f18b, []int{2}
}

func (m *RoutingProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutingProfile.Unmarshal(m, b)
}
func (m *RoutingProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutingProfile.Marshal(b, m, deterministic)
}
func (m *RoutingProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutingProfile.Merge(m, src)
}
func (m *RoutingProfile) XXX_Size() int {
	return xxx_messageInfo_RoutingProfile.Size(m)
}
func (m *RoutingProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutingProfile.DiscardUnknown(m)
}

var xxx_messageInfo_RoutingProfile proto.InternalMessageInfo

func (m *RoutingProfile) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RoutingProfile) GetAsId() string {
	if m != nil {
		return m.AsId
	}
	return ""
}

func (m *RoutingProfile) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *RoutingProfile) GetTlsCert() string {
	if m != nil {
		return m.TlsCert
	}
	return ""
}

func (m *RoutingProfile) GetTlsKey() string {
	if m != nil {
		return m.TlsKey
	}
	return ""
}

func init() {
	proto.RegisterEnum("ns.RatePolicy", RatePolicy_name, RatePolicy_value)
	proto.RegisterType((*ServiceProfile)(nil), "ns.ServiceProfile")
	proto.RegisterType((*DeviceProfile)(nil), "ns.DeviceProfile")
	proto.RegisterType((*RoutingProfile)(nil), "ns.RoutingProfile")
}

func init() {
	proto.RegisterFile("ns/profiles.proto", fileDescriptor_ee79be696a10f18b)
}

var fileDescriptor_ee79be696a10f18b = []byte{
	// 1037 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x9d, 0xd3, 0xd4, 0xb1, 0x69, 0x4b, 0x75, 0x98, 0xb4, 0x55, 0xf7, 0x55, 0x2f, 0x2d, 0x06,
	0x63, 0x40, 0x9d, 0xc5, 0x1d, 0x30, 0xec, 0xb1, 0xb6, 0xd7, 0xa2, 0xeb, 0x8c, 0x1a, 0xca, 0xb0,
	0x87, 0xbd, 0x10, 0xb4, 0x48, 0xcb, 0x9c, 0x25, 0x51, 0xb9, 0xa4, 0xfc, 0xd1, 0xc7, 0xbd, 0xef,
	0x7f, 0xee, 0x67, 0x0c, 0xbc, 0x92, 0xed, 0xf4, 0x63, 0x7b, 0x93, 0xce, 0x39, 0x57, 0x87, 0x97,
	0xbc, 0x47, 0x24, 0xa7, 0x99, 0xb9, 0xcc, 0x41, 0xcf, 0x55, 0x22, 0x4d, 0x3f, 0x07, 0x6d, 0x35,
	0x3d, 0xca, 0xcc, 0xc5, 0xdf, 0x75, 0xe2, 0x5f, 0x4b, 0x58, 0xa9, 0x48, 0x4e, 0x4b, 0x96, 0xfa,
	0xe4, 0x48, 0x89, 0xa0, 0xd6, 0xad, 0xf5, 0xda, 0xe1, 0x91, 0x12, 0xf4, 0x21, 0x39, 0x29, 0x12,
	0x06, 0xdc, 0xca, 0xe0, 0xa8, 0x5b, 0xeb, 0x79, 0x61, 0xbd, 0x48, 0x42, 0x6e, 0x25, 0x7d, 0x4a,
	0xfc, 0x22, 0x61, 0xb3, 0x22, 0x5a, 0x4a, 0xcb, 0x8c, 0x7a, 0x27, 0x83, 0x3b, 0xc8, 0xb7, 0x8b,
	0x64, 0x88, 0xe0, 0xb5, 0x7a, 0x27, 0xe9, 0x0f, 0xa8, 0x72, 0xe5, 0x2c, 0xd7, 0x89, 0x8a, 0xb6,
	0xc1, 0x71, 0xb7, 0xd6, 0xf3, 0x07, 0x7e, 0x3f, 0x33, 0x7d, 0xf7, 0x9d, 0x29, 0xa2, 0xae, 0xea,
	0xf0, 0xe6, 0x4c, 0x45, 0x65, 0x7a, 0xb7, 0x34, 0x15, 0x7b, 0x53, 0xf1, 0xbe, 0x69, 0xbd, 0x34,
	0x15, 0x1f, 0x98, 0x8a, 0xf7, 0x4d, 0x4f, 0x3e, 0x6d, 0x2a, 0x6e, 0x9b, 0x7e, 0x4b, 0xee, 0x71,
	0x21, 0x58, 0xbc, 0x66, 0xa9, 0xb4, 0x5c, 0x70, 0xcb, 0x83, 0x46, 0xb7, 0xd6, 0x6b, 0x84, 0x1e,
	0x17, 0xe2, 0xd5, 0x7a, 0x52, 0x81, 0xf4, 0x19, 0x39, 0x13, 0x72, 0xc5, 0x8c, 0xe5, 0xb6, 0x30,
	0x0c, 0xe4, 0x0d, 0x9b, 0x83, 0xbc, 0x09, 0x9a, 0xb8, 0x90, 0x8e, 0x90, 0xab, 0x6b, 0x64, 0x42,
	0x79, 0xf3, 0x12, 0xe4, 0x0d, 0xfd, 0x89, 0x3c, 0x02, 0x99, 0x6b, 0xb0, 0xec, 0x56, 0xd5, 0x8c,
	0x5b, 0x2b, 0x61, 0x1b, 0x10, 0x34, 0x78, 0x50, 0x0a, 0xc6, 0xbb, 0xd2, 0x61, 0xc9, 0xd2, 0x1f,
	0x49, 0xf0, 0x71, 0x69, 0xca, 0x21, 0x56, 0x59, 0xd0, 0xc2, 0xca, 0xfb, 0x1f, 0x54, 0x4e, 0x90,
	0xa4, 0xf7, 0x49, 0x5d, 0x00, 0x4b, 0x55, 0x16, 0xb4, 0x71, 0x55, 0x77, 0x05, 0x4c, 0x0e, 0x30,
	0xdf, 0x04, 0xde, 0x1e, 0xe6, 0x1b, 0xfa, 0x0d, 0x69, 0x47, 0x0b, 0x9e, 0x65, 0x32, 0x61, 0x29,
	0x37, 0xcb, 0xc0, 0xc7, 0xc3, 0x6f, 0x55, 0xd8, 0x84, 0x9b, 0x25, 0xfd, 0x8a, 0x90, 0x1c, 0x18,
	0x4f, 0x12, 0xbd, 0x96, 0x22, 0xb8, 0x87, 0xde, 0xcd, 0x1c, 0x5e, 0x94, 0x80, 0xa3, 0x17, 0x07,
	0xba, 0x53, 0xd2, 0x8b, 0xdb, 0x34, 0xf0, 0x3d, 0x7d, 0x5a, 0xd2, 0xc0, 0x77, 0xf4, 0xd7, 0xa4,
	0x95, 0xad, 0x97, 0x2c, 0x96, 0x9a, 0x25, 0x3a, 0x0a, 0x68, 0xc9, 0x67, 0xeb, 0xe5, 0x2b, 0xa9,
	0x7f, 0xd5, 0x91, 0x2b, 0xb7, 0x1c, 0x62, 0x69, 0x59, 0x2e, 0x21, 0x38, 0xc3, 0xa5, 0x37, 0x4b,
	0x64, 0x2a, 0x81, 0xf6, 0x48, 0x27, 0x55, 0x99, 0x3b, 0x37, 0xa1, 0x56, 0x12, 0x8c, 0xb2, 0xdb,
	0xe0, 0x1c, 0x45, 0x7e, 0xaa, 0xb2, 0x57, 0xeb, 0xf1, 0x0e, 0xa5, 0x8f, 0x49, 0x2b, 0x5e, 0x1b,
	0x96, 0x83, 0x5a, 0xb9, 0xd1, 0xba, 0x8f, 0x46, 0x24, 0x5e, 0x9b, 0x69, 0x89, 0x5c, 0xfc, 0x53,
	0x27, 0xde, 0x58, 0xfe, 0x5f, 0x1c, 0x7a, 0xa4, 0x63, 0x8a, 0xdc, 0xed, 0xb9, 0x61, 0x51, 0xc2,
	0x8d, 0x61, 0x33, 0xcc, 0x45, 0x23, 0xf4, 0x77, 0xf8, 0xc8, 0xc1, 0x43, 0x37, 0x4e, 0x95, 0x80,
	0x59, 0x95, 0x4a, 0x5d, 0xd8, 0x2a, 0x20, 0x1e, 0xc2, 0xc3, 0xdf, 0x4a, 0xd0, 0x7d, 0x31, 0x57,
	0x59, 0xcc, 0x4c, 0xa2, 0xb1, 0x41, 0xa5, 0x05, 0x66, 0xc4, 0x0b, 0x7d, 0x87, 0x5f, 0x27, 0xda,
	0x75, 0xa9, 0xb4, 0xa0, 0x5d, 0xd2, 0x3e, 0x28, 0x05, 0x54, 0xd1, 0x20, 0x3b, 0xd5, 0x18, 0x5c,
	0x3c, 0x0e, 0x0a, 0x9c, 0xca, 0x2a, 0x1e, 0x3b, 0x0d, 0x4e, 0xe4, 0xc7, 0x3d, 0x44, 0x18, 0x90,
	0x0f, 0x7b, 0x18, 0x1d, 0x7a, 0x88, 0xf6, 0x3d, 0x34, 0x6e, 0xf5, 0x30, 0xda, 0xf5, 0xf0, 0x98,
	0xb4, 0x52, 0x1e, 0x31, 0xdc, 0x67, 0x9d, 0x61, 0x14, 0x9a, 0x21, 0x49, 0x79, 0xf4, 0x7b, 0x89,
	0xd0, 0x3e, 0x39, 0x03, 0x19, 0xb3, 0x9c, 0x03, 0x4f, 0x5d, 0x66, 0x56, 0x0a, 0x85, 0x04, 0x85,
	0xa7, 0x20, 0xe3, 0x29, 0x32, 0x61, 0x45, 0xd0, 0x2f, 0x09, 0x81, 0x0d, 0x13, 0x32, 0xe1, 0x5b,
	0x76, 0x85, 0xb3, 0xee, 0x85, 0x0d, 0xd8, 0x8c, 0x1d, 0x70, 0x45, 0x9f, 0x10, 0xdf, 0xb1, 0xc0,
	0xf4, 0x7c, 0x6e, 0xa4, 0x65, 0x57, 0xd5, 0x98, 0xb7, 0x60, 0x33, 0x86, 0xb7, 0x88, 0x5d, 0xd1,
	0x0b, 0xe2, 0x39, 0x11, 0xb7, 0x1c, 0x7f, 0x04, 0x83, 0x6a, 0xe6, 0x9d, 0xa6, 0xc2, 0x06, 0xf4,
	0x73, 0xd2, 0x84, 0x0d, 0x6e, 0x14, 0x1b, 0xe0, 0xd8, 0x7b, 0xe1, 0x09, 0x6c, 0xdc, 0x26, 0x0d,
	0xe8, 0xf7, 0xe4, 0x7c, 0xce, 0x23, 0xab, 0x61, 0xcb, 0x72, 0x90, 0xce, 0xc6, 0xe9, 0x4c, 0x70,
	0xaf, 0x7b, 0xa7, 0xe7, 0x85, 0xb4, 0xe2, 0xa6, 0x48, 0xb9, 0x0a, 0x43, 0x1f, 0x91, 0x46, 0xca,
	0x37, 0x4c, 0x2a, 0xc8, 0x31, 0x03, 0x5e, 0x78, 0x92, 0xf2, 0xcd, 0xcf, 0x0a, 0x72, 0x77, 0x30,
	0x8e, 0x12, 0x85, 0xdd, 0xb2, 0x68, 0x1b, 0x25, 0x12, 0x53, 0xe0, 0x85, 0xed, 0x94, 0x6f, 0xc6,
	0x85, 0xdd, 0x8e, 0x1c, 0x46, 0x9f, 0x10, 0x6f, 0x7f, 0x30, 0x7f, 0x6a, 0x95, 0x55, 0x51, 0x68,
	0xef, 0xc0, 0x5f, 0xb4, 0xca, 0xe8, 0x17, 0xa4, 0x09, 0x73, 0x06, 0x32, 0x76, 0x1b, 0x78, 0x86,
	0x1b, 0xd8, 0x80, 0x79, 0x88, 0xef, 0xf4, 0x92, 0x9c, 0xef, 0xbf, 0xf0, 0x7c, 0x30, 0x53, 0x96,
	0xcd, 0x59, 0x94, 0x59, 0xcc, 0x43, 0x23, 0x3c, 0xdd, 0x71, 0x48, 0xbd, 0x1c, 0x65, 0x38, 0x7d,
	0x5c, 0xb8, 0xe8, 0xc6, 0x1a, 0x94, 0x5d, 0xa4, 0x4c, 0x09, 0xcc, 0x45, 0x33, 0xf4, 0xb9, 0x80,
	0x17, 0x3b, 0xf8, 0xf5, 0xa7, 0x26, 0x5f, 0x04, 0x0f, 0x3e, 0x31, 0x35, 0x63, 0xd7, 0x46, 0x0e,
	0x92, 0xa7, 0xb3, 0x44, 0x96, 0xff, 0xe8, 0x87, 0xd5, 0x10, 0x56, 0xa0, 0xfb, 0x47, 0x5f, 0xfc,
	0x55, 0x23, 0x7e, 0xa8, 0x0b, 0xab, 0xb2, 0xf8, 0xbf, 0xb2, 0x76, 0x46, 0xee, 0x72, 0xe3, 0x16,
	0x74, 0x84, 0x0b, 0x3a, 0xe6, 0xe6, 0x35, 0xde, 0x47, 0x11, 0x67, 0x91, 0x84, 0x32, 0x4e, 0xcd,
	0xb0, 0x1e, 0xf1, 0x91, 0x04, 0xeb, 0x76, 0xdf, 0x26, 0xa6, 0x64, 0x8e, 0x91, 0x39, 0xb1, 0x89,
	0x41, 0xea, 0x21, 0x71, 0x8f, 0x6c, 0x29, 0xb7, 0x98, 0x99, 0x66, 0x58, 0xb7, 0x89, 0x79, 0x23,
	0xb7, 0xdf, 0x75, 0x09, 0xb9, 0x75, 0x01, 0x34, 0xc8, 0xf1, 0x38, 0x7c, 0x3b, 0xed, 0x7c, 0xe6,
	0x9e, 0x26, 0x2f, 0xc2, 0x37, 0x9d, 0xda, 0x70, 0x42, 0xce, 0x95, 0xee, 0x47, 0x0b, 0x05, 0xb9,
	0xb1, 0x3c, 0x5a, 0xf6, 0x79, 0xae, 0xfa, 0x99, 0x19, 0x7a, 0xd5, 0xa2, 0xcd, 0xd4, 0x5d, 0xa6,
	0xd3, 0xda, 0x1f, 0x4f, 0x63, 0x65, 0x17, 0xc5, 0xac, 0x1f, 0xe9, 0xf4, 0x72, 0x06, 0x3a, 0xe2,
	0x1c, 0x2e, 0x0f, 0x65, 0xcf, 0x78, 0xae, 0x2e, 0x33, 0x33, 0xab, 0xe3, 0xdd, 0xfb, 0xfc, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x6b, 0xad, 0xee, 0x90, 0x07, 0x00, 0x00,
}
