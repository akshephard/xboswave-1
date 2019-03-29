// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xbos.proto

package xbospb

/*
This defines the protocol buffers used by XBOS with WaveMQ. The schema for
a wavemq xbos proto message is:
  xbosproto/RootMessageType

Version 1.0
*/

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

// Root union type
// We reserve 50 fields per imported file
type XBOS struct {
	// Hamilton Fields
	HamiltonData        *HamiltonData        `protobuf:"bytes,50,opt,name=HamiltonData,json=hamiltonData,proto3" json:"HamiltonData,omitempty"`
	HamiltonBRLinkStats *HamiltonBRLinkStats `protobuf:"bytes,51,opt,name=HamiltonBRLinkStats,json=hamiltonBRLinkStats,proto3" json:"HamiltonBRLinkStats,omitempty"`
	HamiltonBRMessage   *HamiltonBRMessage   `protobuf:"bytes,52,opt,name=HamiltonBRMessage,json=hamiltonBRMessage,proto3" json:"HamiltonBRMessage,omitempty"`
	// IoT Fields
	XBOSIoTDeviceState     *XBOSIoTDeviceState     `protobuf:"bytes,100,opt,name=XBOSIoTDeviceState,json=xBOSIoTDeviceState,proto3" json:"XBOSIoTDeviceState,omitempty"`
	XBOSIoTDeviceActuation *XBOSIoTDeviceActuation `protobuf:"bytes,101,opt,name=XBOSIoTDeviceActuation,json=xBOSIoTDeviceActuation,proto3" json:"XBOSIoTDeviceActuation,omitempty"`
	XBOSIoTContext         *XBOSIoTContext         `protobuf:"bytes,102,opt,name=XBOSIoTContext,json=xBOSIoTContext,proto3" json:"XBOSIoTContext,omitempty"`
	// Dent Fields
	DentMeterState *DentMeterState `protobuf:"bytes,150,opt,name=DentMeterState,json=dentMeterState,proto3" json:"DentMeterState,omitempty"`
	// System Status Fields
	BasicServerStatus    *BasicServerStatus `protobuf:"bytes,200,opt,name=BasicServerStatus,json=basicServerStatus,proto3" json:"BasicServerStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *XBOS) Reset()         { *m = XBOS{} }
func (m *XBOS) String() string { return proto.CompactTextString(m) }
func (*XBOS) ProtoMessage()    {}
func (*XBOS) Descriptor() ([]byte, []int) {
	return fileDescriptor_xbos_eda7eb804a957f42, []int{0}
}
func (m *XBOS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XBOS.Unmarshal(m, b)
}
func (m *XBOS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XBOS.Marshal(b, m, deterministic)
}
func (dst *XBOS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XBOS.Merge(dst, src)
}
func (m *XBOS) XXX_Size() int {
	return xxx_messageInfo_XBOS.Size(m)
}
func (m *XBOS) XXX_DiscardUnknown() {
	xxx_messageInfo_XBOS.DiscardUnknown(m)
}

var xxx_messageInfo_XBOS proto.InternalMessageInfo

func (m *XBOS) GetHamiltonData() *HamiltonData {
	if m != nil {
		return m.HamiltonData
	}
	return nil
}

func (m *XBOS) GetHamiltonBRLinkStats() *HamiltonBRLinkStats {
	if m != nil {
		return m.HamiltonBRLinkStats
	}
	return nil
}

func (m *XBOS) GetHamiltonBRMessage() *HamiltonBRMessage {
	if m != nil {
		return m.HamiltonBRMessage
	}
	return nil
}

func (m *XBOS) GetXBOSIoTDeviceState() *XBOSIoTDeviceState {
	if m != nil {
		return m.XBOSIoTDeviceState
	}
	return nil
}

func (m *XBOS) GetXBOSIoTDeviceActuation() *XBOSIoTDeviceActuation {
	if m != nil {
		return m.XBOSIoTDeviceActuation
	}
	return nil
}

func (m *XBOS) GetXBOSIoTContext() *XBOSIoTContext {
	if m != nil {
		return m.XBOSIoTContext
	}
	return nil
}

func (m *XBOS) GetDentMeterState() *DentMeterState {
	if m != nil {
		return m.DentMeterState
	}
	return nil
}

func (m *XBOS) GetBasicServerStatus() *BasicServerStatus {
	if m != nil {
		return m.BasicServerStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*XBOS)(nil), "xbospb.XBOS")
}

func init() { proto.RegisterFile("xbos.proto", fileDescriptor_xbos_eda7eb804a957f42) }

var fileDescriptor_xbos_eda7eb804a957f42 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x86, 0x19, 0x8c, 0xc1, 0x97, 0x6f, 0x54, 0x9a, 0x8d, 0x11, 0x27, 0x88, 0x78, 0xe5, 0xd5,
	0x2e, 0x36, 0x2f, 0xbc, 0x52, 0x9c, 0x03, 0xa7, 0x38, 0x84, 0x4c, 0xc4, 0x3b, 0x49, 0xb7, 0xa3,
	0x0d, 0xda, 0x64, 0x34, 0x67, 0xa3, 0xfe, 0x11, 0xff, 0x92, 0xfe, 0x2c, 0x69, 0x9a, 0x69, 0xd3,
	0xee, 0x2e, 0xe7, 0x79, 0xdf, 0x3e, 0xe5, 0xb4, 0x21, 0x24, 0x8b, 0xb4, 0x19, 0xac, 0x52, 0x8d,
	0x9a, 0xb6, 0xf2, 0xf3, 0x2a, 0xea, 0x07, 0xb1, 0x48, 0xe4, 0x3b, 0x6a, 0x55, 0xf0, 0xfe, 0x3f,
	0xa9, 0xd1, 0x1d, 0xf7, 0x96, 0xa0, 0x30, 0x01, 0x84, 0xd4, 0x81, 0xae, 0xf9, 0x30, 0x08, 0xc9,
	0x73, 0xa2, 0x95, 0x44, 0xed, 0xe8, 0xf1, 0x57, 0x93, 0x34, 0x9f, 0xc6, 0xf7, 0x73, 0x7a, 0x46,
	0xda, 0x53, 0x27, 0x9b, 0x08, 0x14, 0x6c, 0x78, 0xd4, 0x38, 0xf9, 0x3f, 0xec, 0x0e, 0x8a, 0x37,
	0x0d, 0xca, 0x19, 0x6f, 0xc7, 0xa5, 0x89, 0xce, 0x48, 0x67, 0x9b, 0x8e, 0xf9, 0x9d, 0x54, 0x6f,
	0x73, 0x14, 0x68, 0xd8, 0xc8, 0x0a, 0x0e, 0xaa, 0x82, 0x52, 0x85, 0x77, 0xe2, 0x3a, 0xa4, 0xd7,
	0x24, 0xfc, 0xeb, 0xce, 0xc0, 0x18, 0xf1, 0x0a, 0xec, 0xd4, 0xca, 0xf6, 0xeb, 0x32, 0x57, 0xe0,
	0x61, 0x5c, 0x45, 0xf4, 0x96, 0xd0, 0x7c, 0xb3, 0x1b, 0xfd, 0x30, 0x81, 0x8d, 0x5c, 0x40, 0xae,
	0x07, 0xb6, 0xb4, 0xa6, 0xfe, 0xd6, 0x54, 0x6f, 0x70, 0x9a, 0xd5, 0x18, 0x7d, 0x24, 0x3d, 0xaf,
	0x79, 0xb9, 0xc0, 0xb5, 0x40, 0xa9, 0x15, 0x03, 0xeb, 0x3b, 0xdc, 0xe9, 0xfb, 0x6d, 0xf1, 0x5e,
	0xb6, 0x93, 0xd3, 0x73, 0x12, 0xb8, 0x27, 0xae, 0xb4, 0x42, 0xc8, 0x90, 0xbd, 0x58, 0x5f, 0xaf,
	0xe2, 0x73, 0x29, 0x0f, 0x32, 0x6f, 0xa6, 0x17, 0x24, 0x98, 0x80, 0xc2, 0x59, 0xfe, 0x9f, 0x8b,
	0xfd, 0x3e, 0x1b, 0xbe, 0xc0, 0x8f, 0x79, 0xb0, 0xf4, 0x66, 0x3a, 0x25, 0xe1, 0x58, 0x18, 0xb9,
	0x98, 0x43, 0xba, 0x29, 0xd8, 0xda, 0xb0, 0xef, 0x86, 0xff, 0xb9, 0x6b, 0x0d, 0x1e, 0x46, 0x55,
	0x14, 0xb5, 0xec, 0x85, 0x1a, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xea, 0x93, 0xe0, 0x8e, 0xa8,
	0x02, 0x00, 0x00,
}
