// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weather_prediction.proto

package xbospb

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

type Weather_Prediction_State struct {
	Predictions          []*Weather_Prediction_State_Prediction `protobuf:"bytes,1,rep,name=predictions,proto3" json:"predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *Weather_Prediction_State) Reset()         { *m = Weather_Prediction_State{} }
func (m *Weather_Prediction_State) String() string { return proto.CompactTextString(m) }
func (*Weather_Prediction_State) ProtoMessage()    {}
func (*Weather_Prediction_State) Descriptor() ([]byte, []int) {
	return fileDescriptor_weather_prediction_b2716913910afe24, []int{0}
}
func (m *Weather_Prediction_State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Weather_Prediction_State.Unmarshal(m, b)
}
func (m *Weather_Prediction_State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Weather_Prediction_State.Marshal(b, m, deterministic)
}
func (dst *Weather_Prediction_State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Weather_Prediction_State.Merge(dst, src)
}
func (m *Weather_Prediction_State) XXX_Size() int {
	return xxx_messageInfo_Weather_Prediction_State.Size(m)
}
func (m *Weather_Prediction_State) XXX_DiscardUnknown() {
	xxx_messageInfo_Weather_Prediction_State.DiscardUnknown(m)
}

var xxx_messageInfo_Weather_Prediction_State proto.InternalMessageInfo

func (m *Weather_Prediction_State) GetPredictions() []*Weather_Prediction_State_Prediction {
	if m != nil {
		return m.Predictions
	}
	return nil
}

type Weather_Prediction_State_Prediction struct {
	PredictionTime       uint64                 `protobuf:"varint,1,opt,name=prediction_time,json=predictionTime,proto3" json:"prediction_time,omitempty"`
	Prediction           *Weather_Current_State `protobuf:"bytes,2,opt,name=prediction,proto3" json:"prediction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Weather_Prediction_State_Prediction) Reset()         { *m = Weather_Prediction_State_Prediction{} }
func (m *Weather_Prediction_State_Prediction) String() string { return proto.CompactTextString(m) }
func (*Weather_Prediction_State_Prediction) ProtoMessage()    {}
func (*Weather_Prediction_State_Prediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_weather_prediction_b2716913910afe24, []int{0, 0}
}
func (m *Weather_Prediction_State_Prediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Weather_Prediction_State_Prediction.Unmarshal(m, b)
}
func (m *Weather_Prediction_State_Prediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Weather_Prediction_State_Prediction.Marshal(b, m, deterministic)
}
func (dst *Weather_Prediction_State_Prediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Weather_Prediction_State_Prediction.Merge(dst, src)
}
func (m *Weather_Prediction_State_Prediction) XXX_Size() int {
	return xxx_messageInfo_Weather_Prediction_State_Prediction.Size(m)
}
func (m *Weather_Prediction_State_Prediction) XXX_DiscardUnknown() {
	xxx_messageInfo_Weather_Prediction_State_Prediction.DiscardUnknown(m)
}

var xxx_messageInfo_Weather_Prediction_State_Prediction proto.InternalMessageInfo

func (m *Weather_Prediction_State_Prediction) GetPredictionTime() uint64 {
	if m != nil {
		return m.PredictionTime
	}
	return 0
}

func (m *Weather_Prediction_State_Prediction) GetPrediction() *Weather_Current_State {
	if m != nil {
		return m.Prediction
	}
	return nil
}

func init() {
	proto.RegisterType((*Weather_Prediction_State)(nil), "xbospb.Weather_Prediction_State")
	proto.RegisterType((*Weather_Prediction_State_Prediction)(nil), "xbospb.Weather_Prediction_State.Prediction")
}

func init() {
	proto.RegisterFile("weather_prediction.proto", fileDescriptor_weather_prediction_b2716913910afe24)
}

var fileDescriptor_weather_prediction_b2716913910afe24 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x4f, 0x4d, 0x2c,
	0xc9, 0x48, 0x2d, 0x8a, 0x2f, 0x28, 0x4a, 0x4d, 0xc9, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xab, 0x48, 0xca, 0x2f, 0x2e, 0x48, 0x92, 0x12, 0xce, 0x2b,
	0xcd, 0xc9, 0x49, 0x4c, 0xca, 0x49, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x48, 0x4a, 0x89, 0xc2,
	0xb4, 0x25, 0x97, 0x16, 0x15, 0xa5, 0xe6, 0x95, 0x40, 0x84, 0x95, 0xee, 0x33, 0x72, 0x49, 0x84,
	0x43, 0x65, 0x02, 0xe0, 0x06, 0xc6, 0x07, 0x97, 0x24, 0x96, 0xa4, 0x0a, 0xf9, 0x72, 0x71, 0x23,
	0x2c, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xd2, 0xd6, 0x83, 0x58, 0xa3, 0x87, 0x4b,
	0x9b, 0x1e, 0x42, 0x20, 0x08, 0x59, 0xbf, 0x54, 0x09, 0x17, 0x17, 0x42, 0x4a, 0x48, 0x9d, 0x8b,
	0x1f, 0x21, 0x19, 0x5f, 0x92, 0x99, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x87,
	0x10, 0x0e, 0xc9, 0xcc, 0x4d, 0x15, 0xb2, 0xe5, 0xe2, 0x42, 0x88, 0x48, 0x30, 0x29, 0x30, 0x6a,
	0x70, 0x1b, 0xc9, 0xa2, 0x3b, 0xc2, 0x19, 0xe2, 0x2b, 0x88, 0x0b, 0x82, 0x90, 0x34, 0x24, 0xb1,
	0x81, 0x3d, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x06, 0x8d, 0x39, 0x38, 0x01, 0x00,
	0x00,
}
