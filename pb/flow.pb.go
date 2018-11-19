// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow.proto

package flowprotob

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

type FlowMessage_FlowType int32

const (
	FlowMessage_FLOWUNKNOWN FlowMessage_FlowType = 0
	FlowMessage_NFV9        FlowMessage_FlowType = 9
	FlowMessage_IPFIX       FlowMessage_FlowType = 10
	FlowMessage_SFLOW       FlowMessage_FlowType = 5
)

var FlowMessage_FlowType_name = map[int32]string{
	0:  "FLOWUNKNOWN",
	9:  "NFV9",
	10: "IPFIX",
	5:  "SFLOW",
}
var FlowMessage_FlowType_value = map[string]int32{
	"FLOWUNKNOWN": 0,
	"NFV9":        9,
	"IPFIX":       10,
	"SFLOW":       5,
}

func (x FlowMessage_FlowType) String() string {
	return proto.EnumName(FlowMessage_FlowType_name, int32(x))
}
func (FlowMessage_FlowType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_flow_afe3427e2abb0c00, []int{0, 0}
}

// To be deprecated
type FlowMessage_IPType int32

const (
	FlowMessage_IPUNKNOWN FlowMessage_IPType = 0
	FlowMessage_IPv4      FlowMessage_IPType = 4
	FlowMessage_IPv6      FlowMessage_IPType = 6
)

var FlowMessage_IPType_name = map[int32]string{
	0: "IPUNKNOWN",
	4: "IPv4",
	6: "IPv6",
}
var FlowMessage_IPType_value = map[string]int32{
	"IPUNKNOWN": 0,
	"IPv4":      4,
	"IPv6":      6,
}

func (x FlowMessage_IPType) String() string {
	return proto.EnumName(FlowMessage_IPType_name, int32(x))
}
func (FlowMessage_IPType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_flow_afe3427e2abb0c00, []int{0, 1}
}

type FlowMessage struct {
	Type         FlowMessage_FlowType `protobuf:"varint,1,opt,name=Type,proto3,enum=flowprotob.FlowMessage_FlowType" json:"Type,omitempty"`
	TimeRecvd    uint64               `protobuf:"varint,2,opt,name=TimeRecvd,proto3" json:"TimeRecvd,omitempty"`
	SamplingRate uint64               `protobuf:"varint,3,opt,name=SamplingRate,proto3" json:"SamplingRate,omitempty"`
	SequenceNum  uint32               `protobuf:"varint,4,opt,name=SequenceNum,proto3" json:"SequenceNum,omitempty"`
	// Found inside packet
	TimeFlow uint64 `protobuf:"varint,5,opt,name=TimeFlow,proto3" json:"TimeFlow,omitempty"`
	// Source/destination addresses
	SrcIP     []byte             `protobuf:"bytes,6,opt,name=SrcIP,proto3" json:"SrcIP,omitempty"`
	DstIP     []byte             `protobuf:"bytes,7,opt,name=DstIP,proto3" json:"DstIP,omitempty"`
	IPversion FlowMessage_IPType `protobuf:"varint,8,opt,name=IPversion,proto3,enum=flowprotob.FlowMessage_IPType" json:"IPversion,omitempty"`
	// Size of the sampled packet
	Bytes   uint64 `protobuf:"varint,9,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
	Packets uint64 `protobuf:"varint,10,opt,name=Packets,proto3" json:"Packets,omitempty"`
	// Routing information
	RouterAddr []byte `protobuf:"bytes,11,opt,name=RouterAddr,proto3" json:"RouterAddr,omitempty"`
	NextHop    []byte `protobuf:"bytes,12,opt,name=NextHop,proto3" json:"NextHop,omitempty"`
	NextHopAS  uint32 `protobuf:"varint,13,opt,name=NextHopAS,proto3" json:"NextHopAS,omitempty"`
	// Autonomous system information
	SrcAS uint32 `protobuf:"varint,14,opt,name=SrcAS,proto3" json:"SrcAS,omitempty"`
	DstAS uint32 `protobuf:"varint,15,opt,name=DstAS,proto3" json:"DstAS,omitempty"`
	// Prefix size
	SrcNet uint32 `protobuf:"varint,16,opt,name=SrcNet,proto3" json:"SrcNet,omitempty"`
	DstNet uint32 `protobuf:"varint,17,opt,name=DstNet,proto3" json:"DstNet,omitempty"`
	// Interfaces
	SrcIf uint32 `protobuf:"varint,18,opt,name=SrcIf,proto3" json:"SrcIf,omitempty"`
	DstIf uint32 `protobuf:"varint,19,opt,name=DstIf,proto3" json:"DstIf,omitempty"`
	// Layer 4 protocol
	Proto uint32 `protobuf:"varint,20,opt,name=Proto,proto3" json:"Proto,omitempty"`
	// Port for UDP and TCP
	SrcPort uint32 `protobuf:"varint,21,opt,name=SrcPort,proto3" json:"SrcPort,omitempty"`
	DstPort uint32 `protobuf:"varint,22,opt,name=DstPort,proto3" json:"DstPort,omitempty"`
	// IP and TCP special flags
	IPTos            uint32 `protobuf:"varint,23,opt,name=IPTos,proto3" json:"IPTos,omitempty"`
	ForwardingStatus uint32 `protobuf:"varint,24,opt,name=ForwardingStatus,proto3" json:"ForwardingStatus,omitempty"`
	IPTTL            uint32 `protobuf:"varint,25,opt,name=IPTTL,proto3" json:"IPTTL,omitempty"`
	TCPFlags         uint32 `protobuf:"varint,26,opt,name=TCPFlags,proto3" json:"TCPFlags,omitempty"`
	// Ethernet information
	SrcMac uint64 `protobuf:"varint,27,opt,name=SrcMac,proto3" json:"SrcMac,omitempty"`
	DstMac uint64 `protobuf:"varint,28,opt,name=DstMac,proto3" json:"DstMac,omitempty"`
	// 802.1q VLAN in sampled packet
	VlanId uint32 `protobuf:"varint,29,opt,name=VlanId,proto3" json:"VlanId,omitempty"`
	// Layer 3 protocol (IPv4/IPv6/ARP/...)
	Etype    uint32 `protobuf:"varint,30,opt,name=Etype,proto3" json:"Etype,omitempty"`
	IcmpType uint32 `protobuf:"varint,31,opt,name=IcmpType,proto3" json:"IcmpType,omitempty"`
	IcmpCode uint32 `protobuf:"varint,32,opt,name=IcmpCode,proto3" json:"IcmpCode,omitempty"`
	// Vlan
	SrcVlan uint32 `protobuf:"varint,33,opt,name=SrcVlan,proto3" json:"SrcVlan,omitempty"`
	DstVlan uint32 `protobuf:"varint,34,opt,name=DstVlan,proto3" json:"DstVlan,omitempty"`
	// Fragments (IPv4/IPv6)
	FragmentId           uint32   `protobuf:"varint,35,opt,name=FragmentId,proto3" json:"FragmentId,omitempty"`
	FragmentOffset       uint32   `protobuf:"varint,36,opt,name=FragmentOffset,proto3" json:"FragmentOffset,omitempty"`
	IPv6FlowLabel        uint32   `protobuf:"varint,37,opt,name=IPv6FlowLabel,proto3" json:"IPv6FlowLabel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowMessage) Reset()         { *m = FlowMessage{} }
func (m *FlowMessage) String() string { return proto.CompactTextString(m) }
func (*FlowMessage) ProtoMessage()    {}
func (*FlowMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_flow_afe3427e2abb0c00, []int{0}
}
func (m *FlowMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowMessage.Unmarshal(m, b)
}
func (m *FlowMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowMessage.Marshal(b, m, deterministic)
}
func (dst *FlowMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowMessage.Merge(dst, src)
}
func (m *FlowMessage) XXX_Size() int {
	return xxx_messageInfo_FlowMessage.Size(m)
}
func (m *FlowMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FlowMessage proto.InternalMessageInfo

func (m *FlowMessage) GetType() FlowMessage_FlowType {
	if m != nil {
		return m.Type
	}
	return FlowMessage_FLOWUNKNOWN
}

func (m *FlowMessage) GetTimeRecvd() uint64 {
	if m != nil {
		return m.TimeRecvd
	}
	return 0
}

func (m *FlowMessage) GetSamplingRate() uint64 {
	if m != nil {
		return m.SamplingRate
	}
	return 0
}

func (m *FlowMessage) GetSequenceNum() uint32 {
	if m != nil {
		return m.SequenceNum
	}
	return 0
}

func (m *FlowMessage) GetTimeFlow() uint64 {
	if m != nil {
		return m.TimeFlow
	}
	return 0
}

func (m *FlowMessage) GetSrcIP() []byte {
	if m != nil {
		return m.SrcIP
	}
	return nil
}

func (m *FlowMessage) GetDstIP() []byte {
	if m != nil {
		return m.DstIP
	}
	return nil
}

func (m *FlowMessage) GetIPversion() FlowMessage_IPType {
	if m != nil {
		return m.IPversion
	}
	return FlowMessage_IPUNKNOWN
}

func (m *FlowMessage) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *FlowMessage) GetPackets() uint64 {
	if m != nil {
		return m.Packets
	}
	return 0
}

func (m *FlowMessage) GetRouterAddr() []byte {
	if m != nil {
		return m.RouterAddr
	}
	return nil
}

func (m *FlowMessage) GetNextHop() []byte {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *FlowMessage) GetNextHopAS() uint32 {
	if m != nil {
		return m.NextHopAS
	}
	return 0
}

func (m *FlowMessage) GetSrcAS() uint32 {
	if m != nil {
		return m.SrcAS
	}
	return 0
}

func (m *FlowMessage) GetDstAS() uint32 {
	if m != nil {
		return m.DstAS
	}
	return 0
}

func (m *FlowMessage) GetSrcNet() uint32 {
	if m != nil {
		return m.SrcNet
	}
	return 0
}

func (m *FlowMessage) GetDstNet() uint32 {
	if m != nil {
		return m.DstNet
	}
	return 0
}

func (m *FlowMessage) GetSrcIf() uint32 {
	if m != nil {
		return m.SrcIf
	}
	return 0
}

func (m *FlowMessage) GetDstIf() uint32 {
	if m != nil {
		return m.DstIf
	}
	return 0
}

func (m *FlowMessage) GetProto() uint32 {
	if m != nil {
		return m.Proto
	}
	return 0
}

func (m *FlowMessage) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *FlowMessage) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *FlowMessage) GetIPTos() uint32 {
	if m != nil {
		return m.IPTos
	}
	return 0
}

func (m *FlowMessage) GetForwardingStatus() uint32 {
	if m != nil {
		return m.ForwardingStatus
	}
	return 0
}

func (m *FlowMessage) GetIPTTL() uint32 {
	if m != nil {
		return m.IPTTL
	}
	return 0
}

func (m *FlowMessage) GetTCPFlags() uint32 {
	if m != nil {
		return m.TCPFlags
	}
	return 0
}

func (m *FlowMessage) GetSrcMac() uint64 {
	if m != nil {
		return m.SrcMac
	}
	return 0
}

func (m *FlowMessage) GetDstMac() uint64 {
	if m != nil {
		return m.DstMac
	}
	return 0
}

func (m *FlowMessage) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

func (m *FlowMessage) GetEtype() uint32 {
	if m != nil {
		return m.Etype
	}
	return 0
}

func (m *FlowMessage) GetIcmpType() uint32 {
	if m != nil {
		return m.IcmpType
	}
	return 0
}

func (m *FlowMessage) GetIcmpCode() uint32 {
	if m != nil {
		return m.IcmpCode
	}
	return 0
}

func (m *FlowMessage) GetSrcVlan() uint32 {
	if m != nil {
		return m.SrcVlan
	}
	return 0
}

func (m *FlowMessage) GetDstVlan() uint32 {
	if m != nil {
		return m.DstVlan
	}
	return 0
}

func (m *FlowMessage) GetFragmentId() uint32 {
	if m != nil {
		return m.FragmentId
	}
	return 0
}

func (m *FlowMessage) GetFragmentOffset() uint32 {
	if m != nil {
		return m.FragmentOffset
	}
	return 0
}

func (m *FlowMessage) GetIPv6FlowLabel() uint32 {
	if m != nil {
		return m.IPv6FlowLabel
	}
	return 0
}

func init() {
	proto.RegisterType((*FlowMessage)(nil), "flowprotob.FlowMessage")
	proto.RegisterEnum("flowprotob.FlowMessage_FlowType", FlowMessage_FlowType_name, FlowMessage_FlowType_value)
	proto.RegisterEnum("flowprotob.FlowMessage_IPType", FlowMessage_IPType_name, FlowMessage_IPType_value)
}

func init() { proto.RegisterFile("flow.proto", fileDescriptor_flow_afe3427e2abb0c00) }

var fileDescriptor_flow_afe3427e2abb0c00 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xe1, 0x6f, 0xd2, 0x50,
	0x10, 0xc0, 0x45, 0x19, 0x83, 0xc7, 0xd8, 0xea, 0x73, 0xce, 0x73, 0x4e, 0xac, 0x38, 0x0d, 0x71,
	0x09, 0x1f, 0x74, 0x59, 0x62, 0xf4, 0x0b, 0xdb, 0x6c, 0x6c, 0x64, 0x5d, 0xd3, 0xe2, 0xe6, 0xd7,
	0xd2, 0xbe, 0x36, 0xc4, 0xd2, 0x62, 0xfb, 0x60, 0xee, 0x8f, 0xf1, 0x7f, 0x35, 0x77, 0xaf, 0x2d,
	0x4c, 0xe3, 0xb7, 0xf7, 0xfb, 0xdd, 0xf5, 0x5e, 0xef, 0xca, 0xc1, 0x58, 0x18, 0xa7, 0x37, 0x83,
	0x79, 0x96, 0xca, 0x94, 0xd3, 0x99, 0x8e, 0x93, 0xde, 0xef, 0x16, 0x6b, 0x1b, 0x71, 0x7a, 0x73,
	0x21, 0xf2, 0xdc, 0x8b, 0x04, 0x3f, 0x66, 0xf5, 0xf1, 0xed, 0x5c, 0x40, 0x4d, 0xaf, 0xf5, 0xb7,
	0xdf, 0xe9, 0x83, 0x55, 0xea, 0x60, 0x2d, 0x8d, 0xce, 0x98, 0xe7, 0x50, 0x36, 0x3f, 0x60, 0xad,
	0xf1, 0x74, 0x26, 0x1c, 0xe1, 0x2f, 0x03, 0xb8, 0xaf, 0xd7, 0xfa, 0x75, 0x67, 0x25, 0x78, 0x8f,
	0x6d, 0xb9, 0xde, 0x6c, 0x1e, 0x4f, 0x93, 0xc8, 0xf1, 0xa4, 0x80, 0x07, 0x94, 0x70, 0xc7, 0x71,
	0x9d, 0xb5, 0x5d, 0xf1, 0x73, 0x21, 0x12, 0x5f, 0x58, 0x8b, 0x19, 0xd4, 0xf5, 0x5a, 0xbf, 0xe3,
	0xac, 0x2b, 0xbe, 0xcf, 0x9a, 0x58, 0x12, 0x6f, 0x86, 0x0d, 0xaa, 0x50, 0x31, 0xdf, 0x65, 0x1b,
	0x6e, 0xe6, 0x9b, 0x36, 0x34, 0xf4, 0x5a, 0x7f, 0xcb, 0x51, 0x80, 0xf6, 0x3c, 0x97, 0xa6, 0x0d,
	0x9b, 0xca, 0x12, 0xf0, 0x4f, 0xac, 0x65, 0xda, 0x4b, 0x91, 0xe5, 0xd3, 0x34, 0x81, 0x26, 0xb5,
	0xd9, 0xfd, 0x5f, 0x9b, 0xa6, 0x4d, 0x4d, 0xae, 0x1e, 0xc0, 0x9a, 0xa7, 0xb7, 0x52, 0xe4, 0xd0,
	0xa2, 0x57, 0x50, 0xc0, 0x81, 0x6d, 0xda, 0x9e, 0xff, 0x43, 0xc8, 0x1c, 0x18, 0xf9, 0x12, 0x79,
	0x97, 0x31, 0x27, 0x5d, 0x48, 0x91, 0x0d, 0x83, 0x20, 0x83, 0x36, 0xbd, 0xc8, 0x9a, 0xc1, 0x27,
	0x2d, 0xf1, 0x4b, 0x7e, 0x49, 0xe7, 0xb0, 0x45, 0xc1, 0x12, 0x71, 0xa6, 0xc5, 0x71, 0xe8, 0x42,
	0x87, 0xe6, 0xb1, 0x12, 0x45, 0xc7, 0x43, 0x17, 0xb6, 0x29, 0xa2, 0xa0, 0xe8, 0x78, 0xe8, 0xc2,
	0x8e, 0xb2, 0x04, 0x7c, 0x8f, 0x35, 0xdc, 0xcc, 0xb7, 0x84, 0x04, 0x8d, 0x74, 0x41, 0xe8, 0xcf,
	0x73, 0x89, 0xfe, 0xa1, 0xf2, 0x8a, 0xca, 0x69, 0x86, 0xc0, 0xab, 0xda, 0x66, 0x58, 0x4e, 0x33,
	0x84, 0x47, 0x55, 0x6d, 0x65, 0x6d, 0x9c, 0x1b, 0xec, 0x2a, 0x4b, 0x80, 0x5d, 0xb9, 0x99, 0x6f,
	0xa7, 0x99, 0x84, 0xc7, 0xe4, 0x4b, 0xc4, 0xc8, 0x79, 0x2e, 0x29, 0xb2, 0xa7, 0x22, 0x05, 0x62,
	0x25, 0xd3, 0x1e, 0xa7, 0x39, 0x3c, 0x51, 0x95, 0x08, 0xf8, 0x5b, 0xa6, 0x19, 0x69, 0x76, 0xe3,
	0x65, 0xc1, 0x34, 0x89, 0x5c, 0xe9, 0xc9, 0x45, 0x0e, 0x40, 0x09, 0xff, 0xf8, 0xa2, 0xc2, 0x78,
	0x04, 0x4f, 0xab, 0x0a, 0xe3, 0x11, 0xfd, 0x6e, 0xce, 0x6c, 0x23, 0xf6, 0xa2, 0x1c, 0xf6, 0x29,
	0x50, 0x71, 0x31, 0x99, 0x0b, 0xcf, 0x87, 0x67, 0xf4, 0xd9, 0x0a, 0x2a, 0x26, 0x83, 0xfe, 0x40,
	0x79, 0x45, 0xe8, 0xaf, 0x62, 0x2f, 0x31, 0x03, 0x78, 0xae, 0x26, 0xa6, 0x08, 0x6f, 0xfe, 0x2c,
	0x71, 0x6d, 0xba, 0xea, 0x66, 0x02, 0xbc, 0xd9, 0xf4, 0x67, 0x73, 0xda, 0xa7, 0x17, 0xea, 0xe6,
	0x92, 0xcb, 0xd8, 0x59, 0x1a, 0x08, 0xd0, 0x57, 0x31, 0xe4, 0x62, 0x7a, 0x58, 0x1a, 0x5e, 0x56,
	0xd3, 0x43, 0x2c, 0xa6, 0x47, 0x91, 0x5e, 0x35, 0x3d, 0x8a, 0x74, 0x19, 0x33, 0x32, 0x2f, 0x9a,
	0x89, 0x44, 0x9a, 0x01, 0xbc, 0xa2, 0xe0, 0x9a, 0xe1, 0x6f, 0xd8, 0x76, 0x49, 0x97, 0x61, 0x98,
	0x0b, 0x09, 0x87, 0x94, 0xf3, 0x97, 0xe5, 0x87, 0xac, 0x63, 0xda, 0xcb, 0x13, 0x5c, 0x82, 0x91,
	0x37, 0x11, 0x31, 0xbc, 0xa6, 0xb4, 0xbb, 0xb2, 0xf7, 0x91, 0x35, 0xcb, 0x7f, 0x00, 0xbe, 0xc3,
	0xda, 0xc6, 0xe8, 0xf2, 0xfa, 0x9b, 0xf5, 0xd5, 0xba, 0xbc, 0xb6, 0xb4, 0x7b, 0xbc, 0xc9, 0xea,
	0x96, 0x71, 0xf5, 0x41, 0x6b, 0xf1, 0x16, 0x7e, 0x10, 0xc3, 0xfc, 0xae, 0x31, 0x3c, 0xba, 0x98,
	0xa6, 0x6d, 0xf4, 0x8e, 0x58, 0x43, 0xed, 0x15, 0xef, 0xe0, 0x2a, 0xde, 0x79, 0xd0, 0xb4, 0x97,
	0xc7, 0x5a, 0xbd, 0x38, 0x9d, 0x68, 0x8d, 0xd3, 0x23, 0xb6, 0xef, 0xa7, 0xb3, 0x81, 0x1f, 0xa7,
	0x8b, 0x20, 0x8c, 0xbd, 0x4c, 0x0c, 0x12, 0x21, 0x69, 0x5d, 0xbd, 0x28, 0x3a, 0xed, 0xac, 0x2d,
	0xab, 0x3d, 0x99, 0x34, 0x68, 0x85, 0xdf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x5c, 0x7b,
	0x89, 0xed, 0x04, 0x00, 0x00,
}
