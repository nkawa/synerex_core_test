// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synerex.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Response struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_synerex_fe137f295ad83d2f, []int{0}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Response) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ConfirmResponse struct {
	Ok                   bool               `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	MbusId               uint64             `protobuf:"fixed64,2,opt,name=mbus_id,json=mbusId,proto3" json:"mbus_id,omitempty"`
	Wait                 *duration.Duration `protobuf:"bytes,3,opt,name=wait,proto3" json:"wait,omitempty"`
	Err                  string             `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ConfirmResponse) Reset()         { *m = ConfirmResponse{} }
func (m *ConfirmResponse) String() string { return proto.CompactTextString(m) }
func (*ConfirmResponse) ProtoMessage()    {}
func (*ConfirmResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_synerex_fe137f295ad83d2f, []int{1}
}
func (m *ConfirmResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmResponse.Unmarshal(m, b)
}
func (m *ConfirmResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmResponse.Marshal(b, m, deterministic)
}
func (dst *ConfirmResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmResponse.Merge(dst, src)
}
func (m *ConfirmResponse) XXX_Size() int {
	return xxx_messageInfo_ConfirmResponse.Size(m)
}
func (m *ConfirmResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmResponse proto.InternalMessageInfo

func (m *ConfirmResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ConfirmResponse) GetMbusId() uint64 {
	if m != nil {
		return m.MbusId
	}
	return 0
}

func (m *ConfirmResponse) GetWait() *duration.Duration {
	if m != nil {
		return m.Wait
	}
	return nil
}

func (m *ConfirmResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type Content struct {
	Entity               []byte   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_synerex_fe137f295ad83d2f, []int{2}
}
func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (dst *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(dst, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetEntity() []byte {
	if m != nil {
		return m.Entity
	}
	return nil
}

type Supply struct {
	Id                   uint64               `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderId             uint64               `protobuf:"fixed64,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	TargetId             uint64               `protobuf:"fixed64,3,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	ChannelType          uint32               `protobuf:"varint,4,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	SupplyName           string               `protobuf:"bytes,5,opt,name=supply_name,json=supplyName,proto3" json:"supply_name,omitempty"`
	Ts                   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=ts,proto3" json:"ts,omitempty"`
	ArgJson              string               `protobuf:"bytes,7,opt,name=arg_json,json=argJson,proto3" json:"arg_json,omitempty"`
	MbusId               uint64               `protobuf:"fixed64,8,opt,name=mbus_id,json=mbusId,proto3" json:"mbus_id,omitempty"`
	Cdata                *Content             `protobuf:"bytes,9,opt,name=cdata,proto3" json:"cdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Supply) Reset()         { *m = Supply{} }
func (m *Supply) String() string { return proto.CompactTextString(m) }
func (*Supply) ProtoMessage()    {}
func (*Supply) Descriptor() ([]byte, []int) {
	return fileDescriptor_synerex_fe137f295ad83d2f, []int{3}
}
func (m *Supply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Supply.Unmarshal(m, b)
}
func (m *Supply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Supply.Marshal(b, m, deterministic)
}
func (dst *Supply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supply.Merge(dst, src)
}
func (m *Supply) XXX_Size() int {
	return xxx_messageInfo_Supply.Size(m)
}
func (m *Supply) XXX_DiscardUnknown() {
	xxx_messageInfo_Supply.DiscardUnknown(m)
}

var xxx_messageInfo_Supply proto.InternalMessageInfo

func (m *Supply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Supply) GetSenderId() uint64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *Supply) GetTargetId() uint64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Supply) GetChannelType() uint32 {
	if m != nil {
		return m.ChannelType
	}
	return 0
}

func (m *Supply) GetSupplyName() string {
	if m != nil {
		return m.SupplyName
	}
	return ""
}

func (m *Supply) GetTs() *timestamp.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *Supply) GetArgJson() string {
	if m != nil {
		return m.ArgJson
	}
	return ""
}

func (m *Supply) GetMbusId() uint64 {
	if m != nil {
		return m.MbusId
	}
	return 0
}

func (m *Supply) GetCdata() *Content {
	if m != nil {
		return m.Cdata
	}
	return nil
}

type Demand struct {
	Id                   uint64               `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderId             uint64               `protobuf:"fixed64,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	TargetId             uint64               `protobuf:"fixed64,3,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	ChannelType          uint32               `protobuf:"varint,4,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	DemandName           string               `protobuf:"bytes,5,opt,name=demand_name,json=demandName,proto3" json:"demand_name,omitempty"`
	Ts                   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=ts,proto3" json:"ts,omitempty"`
	ArgJson              string               `protobuf:"bytes,7,opt,name=arg_json,json=argJson,proto3" json:"arg_json,omitempty"`
	MbusId               uint64               `protobuf:"fixed64,8,opt,name=mbus_id,json=mbusId,proto3" json:"mbus_id,omitempty"`
	Cdata                *Content             `protobuf:"bytes,9,opt,name=cdata,proto3" json:"cdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Demand) Reset()         { *m = Demand{} }
func (m *Demand) String() string { return proto.CompactTextString(m) }
func (*Demand) ProtoMessage()    {}
func (*Demand) Descriptor() ([]byte, []int) {
	return fileDescriptor_synerex_fe137f295ad83d2f, []int{4}
}
func (m *Demand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Demand.Unmarshal(m, b)
}
func (m *Demand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Demand.Marshal(b, m, deterministic)
}
func (dst *Demand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Demand.Merge(dst, src)
}
func (m *Demand) XXX_Size() int {
	return xxx_messageInfo_Demand.Size(m)
}
func (m *Demand) XXX_DiscardUnknown() {
	xxx_messageInfo_Demand.DiscardUnknown(m)
}

var xxx_messageInfo_Demand proto.InternalMessageInfo

func (m *Demand) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Demand) GetSenderId() uint64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *Demand) GetTargetId() uint64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Demand) GetChannelType() uint32 {
	if m != nil {
		return m.ChannelType
	}
	return 0
}

func (m *Demand) GetDemandName() string {
	if m != nil {
		return m.DemandName
	}
	return ""
}

func (m *Demand) GetTs() *timestamp.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *Demand) GetArgJson() string {
	if m != nil {
		return m.ArgJson
	}
	return ""
}

func (m *Demand) GetMbusId() uint64 {
	if m != nil {
		return m.MbusId
	}
	return 0
}

func (m *Demand) GetCdata() *Content {
	if m != nil {
		return m.Cdata
	}
	return nil
}

type Target struct {
	Id                   uint64             `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderId             uint64             `protobuf:"fixed64,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	TargetId             uint64             `protobuf:"fixed64,3,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	ChannelType          uint32             `protobuf:"varint,4,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	Wait                 *duration.Duration `protobuf:"bytes,5,opt,name=wait,proto3" json:"wait,omitempty"`
	MbusId               uint64             `protobuf:"fixed64,6,opt,name=mbus_id,json=mbusId,proto3" json:"mbus_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_synerex_fe137f295ad83d2f, []int{5}
}
func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (dst *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(dst, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Target) GetSenderId() uint64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *Target) GetTargetId() uint64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Target) GetChannelType() uint32 {
	if m != nil {
		return m.ChannelType
	}
	return 0
}

func (m *Target) GetWait() *duration.Duration {
	if m != nil {
		return m.Wait
	}
	return nil
}

func (m *Target) GetMbusId() uint64 {
	if m != nil {
		return m.MbusId
	}
	return 0
}

type Channel struct {
	ClientId             uint64   `protobuf:"fixed64,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ChannelType          uint32   `protobuf:"varint,2,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	ArgJson              string   `protobuf:"bytes,3,opt,name=arg_json,json=argJson,proto3" json:"arg_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_synerex_fe137f295ad83d2f, []int{6}
}
func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (dst *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(dst, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetClientId() uint64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Channel) GetChannelType() uint32 {
	if m != nil {
		return m.ChannelType
	}
	return 0
}

func (m *Channel) GetArgJson() string {
	if m != nil {
		return m.ArgJson
	}
	return ""
}

type Mbus struct {
	ClientId             uint64   `protobuf:"fixed64,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	MbusId               uint64   `protobuf:"fixed64,2,opt,name=mbus_id,json=mbusId,proto3" json:"mbus_id,omitempty"`
	ArgJson              string   `protobuf:"bytes,3,opt,name=arg_json,json=argJson,proto3" json:"arg_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mbus) Reset()         { *m = Mbus{} }
func (m *Mbus) String() string { return proto.CompactTextString(m) }
func (*Mbus) ProtoMessage()    {}
func (*Mbus) Descriptor() ([]byte, []int) {
	return fileDescriptor_synerex_fe137f295ad83d2f, []int{7}
}
func (m *Mbus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mbus.Unmarshal(m, b)
}
func (m *Mbus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mbus.Marshal(b, m, deterministic)
}
func (dst *Mbus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mbus.Merge(dst, src)
}
func (m *Mbus) XXX_Size() int {
	return xxx_messageInfo_Mbus.Size(m)
}
func (m *Mbus) XXX_DiscardUnknown() {
	xxx_messageInfo_Mbus.DiscardUnknown(m)
}

var xxx_messageInfo_Mbus proto.InternalMessageInfo

func (m *Mbus) GetClientId() uint64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Mbus) GetMbusId() uint64 {
	if m != nil {
		return m.MbusId
	}
	return 0
}

func (m *Mbus) GetArgJson() string {
	if m != nil {
		return m.ArgJson
	}
	return ""
}

type MbusMsg struct {
	MsgId                uint64   `protobuf:"fixed64,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	SenderId             uint64   `protobuf:"fixed64,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	TargetId             uint64   `protobuf:"fixed64,3,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	MbusId               uint64   `protobuf:"fixed64,4,opt,name=mbus_id,json=mbusId,proto3" json:"mbus_id,omitempty"`
	MsgType              uint32   `protobuf:"varint,5,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	MsgInfo              string   `protobuf:"bytes,6,opt,name=msg_info,json=msgInfo,proto3" json:"msg_info,omitempty"`
	ArgJson              string   `protobuf:"bytes,7,opt,name=arg_json,json=argJson,proto3" json:"arg_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MbusMsg) Reset()         { *m = MbusMsg{} }
func (m *MbusMsg) String() string { return proto.CompactTextString(m) }
func (*MbusMsg) ProtoMessage()    {}
func (*MbusMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_synerex_fe137f295ad83d2f, []int{8}
}
func (m *MbusMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MbusMsg.Unmarshal(m, b)
}
func (m *MbusMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MbusMsg.Marshal(b, m, deterministic)
}
func (dst *MbusMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MbusMsg.Merge(dst, src)
}
func (m *MbusMsg) XXX_Size() int {
	return xxx_messageInfo_MbusMsg.Size(m)
}
func (m *MbusMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MbusMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MbusMsg proto.InternalMessageInfo

func (m *MbusMsg) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *MbusMsg) GetSenderId() uint64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *MbusMsg) GetTargetId() uint64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *MbusMsg) GetMbusId() uint64 {
	if m != nil {
		return m.MbusId
	}
	return 0
}

func (m *MbusMsg) GetMsgType() uint32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *MbusMsg) GetMsgInfo() string {
	if m != nil {
		return m.MsgInfo
	}
	return ""
}

func (m *MbusMsg) GetArgJson() string {
	if m != nil {
		return m.ArgJson
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "api.Response")
	proto.RegisterType((*ConfirmResponse)(nil), "api.ConfirmResponse")
	proto.RegisterType((*Content)(nil), "api.Content")
	proto.RegisterType((*Supply)(nil), "api.Supply")
	proto.RegisterType((*Demand)(nil), "api.Demand")
	proto.RegisterType((*Target)(nil), "api.Target")
	proto.RegisterType((*Channel)(nil), "api.Channel")
	proto.RegisterType((*Mbus)(nil), "api.Mbus")
	proto.RegisterType((*MbusMsg)(nil), "api.MbusMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SynerexClient is the client API for Synerex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SynerexClient interface {
	RegisterDemand(ctx context.Context, in *Demand, opts ...grpc.CallOption) (*Response, error)
	RegisterSupply(ctx context.Context, in *Supply, opts ...grpc.CallOption) (*Response, error)
	ProposeDemand(ctx context.Context, in *Demand, opts ...grpc.CallOption) (*Response, error)
	ProposeSupply(ctx context.Context, in *Supply, opts ...grpc.CallOption) (*Response, error)
	SelectSupply(ctx context.Context, in *Target, opts ...grpc.CallOption) (*ConfirmResponse, error)
	SelectDemand(ctx context.Context, in *Target, opts ...grpc.CallOption) (*ConfirmResponse, error)
	Confirm(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Response, error)
	SubscribeDemand(ctx context.Context, in *Channel, opts ...grpc.CallOption) (Synerex_SubscribeDemandClient, error)
	SubscribeSupply(ctx context.Context, in *Channel, opts ...grpc.CallOption) (Synerex_SubscribeSupplyClient, error)
	SubscribeMbus(ctx context.Context, in *Mbus, opts ...grpc.CallOption) (Synerex_SubscribeMbusClient, error)
	SendMsg(ctx context.Context, in *MbusMsg, opts ...grpc.CallOption) (*Response, error)
	CloseMbus(ctx context.Context, in *Mbus, opts ...grpc.CallOption) (*Response, error)
}

type synerexClient struct {
	cc *grpc.ClientConn
}

func NewSynerexClient(cc *grpc.ClientConn) SynerexClient {
	return &synerexClient{cc}
}

func (c *synerexClient) RegisterDemand(ctx context.Context, in *Demand, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Synerex/RegisterDemand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synerexClient) RegisterSupply(ctx context.Context, in *Supply, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Synerex/RegisterSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synerexClient) ProposeDemand(ctx context.Context, in *Demand, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Synerex/ProposeDemand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synerexClient) ProposeSupply(ctx context.Context, in *Supply, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Synerex/ProposeSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synerexClient) SelectSupply(ctx context.Context, in *Target, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, "/api.Synerex/SelectSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synerexClient) SelectDemand(ctx context.Context, in *Target, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, "/api.Synerex/SelectDemand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synerexClient) Confirm(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Synerex/Confirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synerexClient) SubscribeDemand(ctx context.Context, in *Channel, opts ...grpc.CallOption) (Synerex_SubscribeDemandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synerex_serviceDesc.Streams[0], "/api.Synerex/SubscribeDemand", opts...)
	if err != nil {
		return nil, err
	}
	x := &synerexSubscribeDemandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synerex_SubscribeDemandClient interface {
	Recv() (*Demand, error)
	grpc.ClientStream
}

type synerexSubscribeDemandClient struct {
	grpc.ClientStream
}

func (x *synerexSubscribeDemandClient) Recv() (*Demand, error) {
	m := new(Demand)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *synerexClient) SubscribeSupply(ctx context.Context, in *Channel, opts ...grpc.CallOption) (Synerex_SubscribeSupplyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synerex_serviceDesc.Streams[1], "/api.Synerex/SubscribeSupply", opts...)
	if err != nil {
		return nil, err
	}
	x := &synerexSubscribeSupplyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synerex_SubscribeSupplyClient interface {
	Recv() (*Supply, error)
	grpc.ClientStream
}

type synerexSubscribeSupplyClient struct {
	grpc.ClientStream
}

func (x *synerexSubscribeSupplyClient) Recv() (*Supply, error) {
	m := new(Supply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *synerexClient) SubscribeMbus(ctx context.Context, in *Mbus, opts ...grpc.CallOption) (Synerex_SubscribeMbusClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synerex_serviceDesc.Streams[2], "/api.Synerex/SubscribeMbus", opts...)
	if err != nil {
		return nil, err
	}
	x := &synerexSubscribeMbusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synerex_SubscribeMbusClient interface {
	Recv() (*MbusMsg, error)
	grpc.ClientStream
}

type synerexSubscribeMbusClient struct {
	grpc.ClientStream
}

func (x *synerexSubscribeMbusClient) Recv() (*MbusMsg, error) {
	m := new(MbusMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *synerexClient) SendMsg(ctx context.Context, in *MbusMsg, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Synerex/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synerexClient) CloseMbus(ctx context.Context, in *Mbus, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Synerex/CloseMbus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SynerexServer is the server API for Synerex service.
type SynerexServer interface {
	RegisterDemand(context.Context, *Demand) (*Response, error)
	RegisterSupply(context.Context, *Supply) (*Response, error)
	ProposeDemand(context.Context, *Demand) (*Response, error)
	ProposeSupply(context.Context, *Supply) (*Response, error)
	SelectSupply(context.Context, *Target) (*ConfirmResponse, error)
	SelectDemand(context.Context, *Target) (*ConfirmResponse, error)
	Confirm(context.Context, *Target) (*Response, error)
	SubscribeDemand(*Channel, Synerex_SubscribeDemandServer) error
	SubscribeSupply(*Channel, Synerex_SubscribeSupplyServer) error
	SubscribeMbus(*Mbus, Synerex_SubscribeMbusServer) error
	SendMsg(context.Context, *MbusMsg) (*Response, error)
	CloseMbus(context.Context, *Mbus) (*Response, error)
}

func RegisterSynerexServer(s *grpc.Server, srv SynerexServer) {
	s.RegisterService(&_Synerex_serviceDesc, srv)
}

func _Synerex_RegisterDemand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Demand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynerexServer).RegisterDemand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synerex/RegisterDemand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynerexServer).RegisterDemand(ctx, req.(*Demand))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synerex_RegisterSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynerexServer).RegisterSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synerex/RegisterSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynerexServer).RegisterSupply(ctx, req.(*Supply))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synerex_ProposeDemand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Demand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynerexServer).ProposeDemand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synerex/ProposeDemand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynerexServer).ProposeDemand(ctx, req.(*Demand))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synerex_ProposeSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynerexServer).ProposeSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synerex/ProposeSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynerexServer).ProposeSupply(ctx, req.(*Supply))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synerex_SelectSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynerexServer).SelectSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synerex/SelectSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynerexServer).SelectSupply(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synerex_SelectDemand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynerexServer).SelectDemand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synerex/SelectDemand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynerexServer).SelectDemand(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synerex_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynerexServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synerex/Confirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynerexServer).Confirm(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synerex_SubscribeDemand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SynerexServer).SubscribeDemand(m, &synerexSubscribeDemandServer{stream})
}

type Synerex_SubscribeDemandServer interface {
	Send(*Demand) error
	grpc.ServerStream
}

type synerexSubscribeDemandServer struct {
	grpc.ServerStream
}

func (x *synerexSubscribeDemandServer) Send(m *Demand) error {
	return x.ServerStream.SendMsg(m)
}

func _Synerex_SubscribeSupply_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SynerexServer).SubscribeSupply(m, &synerexSubscribeSupplyServer{stream})
}

type Synerex_SubscribeSupplyServer interface {
	Send(*Supply) error
	grpc.ServerStream
}

type synerexSubscribeSupplyServer struct {
	grpc.ServerStream
}

func (x *synerexSubscribeSupplyServer) Send(m *Supply) error {
	return x.ServerStream.SendMsg(m)
}

func _Synerex_SubscribeMbus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Mbus)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SynerexServer).SubscribeMbus(m, &synerexSubscribeMbusServer{stream})
}

type Synerex_SubscribeMbusServer interface {
	Send(*MbusMsg) error
	grpc.ServerStream
}

type synerexSubscribeMbusServer struct {
	grpc.ServerStream
}

func (x *synerexSubscribeMbusServer) Send(m *MbusMsg) error {
	return x.ServerStream.SendMsg(m)
}

func _Synerex_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MbusMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynerexServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synerex/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynerexServer).SendMsg(ctx, req.(*MbusMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synerex_CloseMbus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mbus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynerexServer).CloseMbus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synerex/CloseMbus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynerexServer).CloseMbus(ctx, req.(*Mbus))
	}
	return interceptor(ctx, in, info, handler)
}

var _Synerex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Synerex",
	HandlerType: (*SynerexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDemand",
			Handler:    _Synerex_RegisterDemand_Handler,
		},
		{
			MethodName: "RegisterSupply",
			Handler:    _Synerex_RegisterSupply_Handler,
		},
		{
			MethodName: "ProposeDemand",
			Handler:    _Synerex_ProposeDemand_Handler,
		},
		{
			MethodName: "ProposeSupply",
			Handler:    _Synerex_ProposeSupply_Handler,
		},
		{
			MethodName: "SelectSupply",
			Handler:    _Synerex_SelectSupply_Handler,
		},
		{
			MethodName: "SelectDemand",
			Handler:    _Synerex_SelectDemand_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _Synerex_Confirm_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _Synerex_SendMsg_Handler,
		},
		{
			MethodName: "CloseMbus",
			Handler:    _Synerex_CloseMbus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeDemand",
			Handler:       _Synerex_SubscribeDemand_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeSupply",
			Handler:       _Synerex_SubscribeSupply_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeMbus",
			Handler:       _Synerex_SubscribeMbus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "synerex.proto",
}

func init() { proto.RegisterFile("synerex.proto", fileDescriptor_synerex_fe137f295ad83d2f) }

var fileDescriptor_synerex_fe137f295ad83d2f = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xc1, 0x6a, 0xdb, 0x4c,
	0x10, 0x8e, 0x64, 0x5b, 0xb2, 0xc7, 0x76, 0xf2, 0x23, 0xfe, 0xb6, 0x8a, 0x0a, 0x4d, 0xa2, 0x4b,
	0x4c, 0x49, 0x94, 0x90, 0xbc, 0x41, 0x9d, 0x8b, 0x0b, 0x29, 0x45, 0x0e, 0x14, 0x7a, 0x31, 0x6b,
	0x6b, 0xad, 0xaa, 0xb1, 0x76, 0xc5, 0xee, 0x9a, 0xd6, 0xf4, 0x4d, 0xfa, 0x2e, 0xa5, 0x97, 0x3e,
	0x58, 0xd1, 0xac, 0x14, 0x4b, 0x71, 0x92, 0x06, 0x0a, 0xa1, 0xb7, 0xdd, 0x99, 0x6f, 0x66, 0xbe,
	0xf9, 0xb4, 0x33, 0x82, 0xbe, 0x5c, 0x31, 0x2a, 0xe8, 0xd7, 0x20, 0x13, 0x5c, 0x71, 0xa7, 0x41,
	0xb2, 0xc4, 0xdb, 0x8b, 0x39, 0x8f, 0x17, 0xf4, 0x04, 0x4d, 0xd3, 0xe5, 0xfc, 0x44, 0x25, 0x29,
	0x95, 0x8a, 0xa4, 0x99, 0x46, 0x79, 0xaf, 0x6e, 0x03, 0xa2, 0xa5, 0x20, 0x2a, 0xe1, 0x4c, 0xfb,
	0xfd, 0x23, 0x68, 0x87, 0x54, 0x66, 0x9c, 0x49, 0xea, 0x6c, 0x83, 0xc9, 0xaf, 0x5d, 0x63, 0xdf,
	0x18, 0xb4, 0x43, 0x93, 0x5f, 0x3b, 0xff, 0x41, 0x83, 0x0a, 0xe1, 0x9a, 0xfb, 0xc6, 0xa0, 0x13,
	0xe6, 0x47, 0xff, 0x1b, 0xec, 0x0c, 0x39, 0x9b, 0x27, 0x22, 0xbd, 0x37, 0xe8, 0x05, 0xd8, 0xe9,
	0x74, 0x29, 0x27, 0x49, 0x84, 0x81, 0x56, 0x68, 0xe5, 0xd7, 0x51, 0xe4, 0x1c, 0x43, 0xf3, 0x0b,
	0x49, 0x94, 0xdb, 0xd8, 0x37, 0x06, 0xdd, 0xb3, 0xdd, 0x40, 0x13, 0x0b, 0x4a, 0x62, 0xc1, 0x45,
	0x41, 0x2c, 0x44, 0x58, 0x59, 0xbc, 0xb9, 0x2e, 0x7e, 0x00, 0xf6, 0x90, 0x33, 0x45, 0x99, 0x72,
	0x9e, 0x83, 0x45, 0x99, 0x4a, 0xd4, 0x0a, 0x0b, 0xf7, 0xc2, 0xe2, 0xe6, 0x7f, 0x37, 0xc1, 0x1a,
	0x2f, 0xb3, 0x6c, 0xb1, 0xca, 0x79, 0x25, 0x11, 0xba, 0xad, 0xd0, 0x4c, 0x22, 0xe7, 0x25, 0x74,
	0x24, 0x65, 0x11, 0x15, 0x6b, 0x66, 0x6d, 0x6d, 0x18, 0xa1, 0x53, 0x11, 0x11, 0x53, 0x95, 0x3b,
	0x1b, 0xda, 0xa9, 0x0d, 0xa3, 0xc8, 0x39, 0x80, 0xde, 0xec, 0x13, 0x61, 0x8c, 0x2e, 0x26, 0x6a,
	0x95, 0x51, 0xa4, 0xd4, 0x0f, 0xbb, 0x85, 0xed, 0x6a, 0x95, 0x51, 0x67, 0x0f, 0xba, 0x12, 0xcb,
	0x4e, 0x18, 0x49, 0xa9, 0xdb, 0x42, 0xd2, 0xa0, 0x4d, 0xef, 0x48, 0x4a, 0x9d, 0xd7, 0x60, 0x2a,
	0xe9, 0x5a, 0xd8, 0xba, 0xb7, 0xd1, 0xfa, 0x55, 0xf9, 0xd1, 0x42, 0x53, 0x49, 0x67, 0x17, 0xda,
	0x44, 0xc4, 0x93, 0xcf, 0x92, 0x33, 0xd7, 0xc6, 0x4c, 0x36, 0x11, 0xf1, 0x5b, 0xc9, 0x59, 0x55,
	0xdc, 0x76, 0x4d, 0x5c, 0x1f, 0x5a, 0xb3, 0x88, 0x28, 0xe2, 0x76, 0xb0, 0x44, 0x2f, 0x20, 0x59,
	0x12, 0x14, 0x6a, 0x85, 0xda, 0x85, 0xe2, 0x5c, 0xd0, 0x94, 0xb0, 0xe8, 0xc9, 0xc5, 0x89, 0xb0,
	0x6c, 0x4d, 0x1c, 0x6d, 0xfa, 0x67, 0xc4, 0xf9, 0x61, 0x80, 0x75, 0x85, 0x4d, 0x3d, 0xad, 0x38,
	0xe5, 0x54, 0xb4, 0x1e, 0x37, 0x15, 0x95, 0x1e, 0xad, 0x6a, 0x8f, 0x7e, 0x04, 0xf6, 0x50, 0xa7,
	0xcd, 0x29, 0xcd, 0x16, 0x09, 0x65, 0x48, 0x49, 0xb7, 0xd1, 0xd6, 0x86, 0x3b, 0x28, 0x99, 0x9b,
	0x94, 0xaa, 0x12, 0x37, 0x6a, 0x12, 0xfb, 0x1f, 0xa0, 0x79, 0x39, 0x5d, 0xca, 0x87, 0x4b, 0xdc,
	0xbb, 0x01, 0x1e, 0x48, 0xfc, 0xcb, 0x00, 0x3b, 0xcf, 0x7c, 0x29, 0x63, 0xe7, 0x19, 0x58, 0xa9,
	0x8c, 0xd7, 0x99, 0x5b, 0xa9, 0x8c, 0x47, 0x7f, 0xf3, 0x19, 0x2a, 0x84, 0x9a, 0xb7, 0x09, 0xe5,
	0x95, 0x50, 0x88, 0x16, 0x0a, 0x61, 0xa7, 0x32, 0x2e, 0x45, 0x40, 0x12, 0x6c, 0xce, 0x51, 0xe9,
	0x0e, 0xba, 0x46, 0x6c, 0xce, 0x1f, 0x78, 0x82, 0x67, 0x3f, 0x9b, 0x60, 0x8f, 0xf5, 0x96, 0x76,
	0x02, 0xd8, 0x0e, 0x69, 0x9c, 0x48, 0x45, 0x45, 0x31, 0x75, 0x5d, 0x7c, 0x78, 0xfa, 0xe2, 0xf5,
	0xf1, 0x52, 0xae, 0x51, 0x7f, 0xab, 0x8a, 0x2f, 0x56, 0x98, 0xc6, 0xeb, 0xcb, 0x26, 0xfe, 0x18,
	0xfa, 0xef, 0x05, 0xcf, 0xb8, 0xa4, 0x8f, 0x4a, 0xbf, 0x86, 0x3f, 0x2a, 0xfb, 0x39, 0xf4, 0xc6,
	0x74, 0x41, 0x67, 0xaa, 0x86, 0xd6, 0x13, 0xe2, 0xfd, 0x5f, 0x4e, 0x50, 0xf5, 0x4f, 0x50, 0x0d,
	0xaa, 0x31, 0xfa, 0x43, 0xd0, 0x21, 0xae, 0xf5, 0xdc, 0x58, 0xc7, 0x6f, 0x50, 0x3a, 0x85, 0x9d,
	0xf1, 0x72, 0x2a, 0x67, 0x22, 0x99, 0x96, 0x2d, 0x17, 0xa3, 0xac, 0x1f, 0xaf, 0x57, 0x15, 0xc0,
	0xdf, 0x3a, 0x35, 0x6a, 0x11, 0x45, 0x1f, 0x77, 0x45, 0x68, 0x17, 0x46, 0x1c, 0x41, 0xff, 0x26,
	0x02, 0x5f, 0x7a, 0x07, 0x11, 0xf9, 0xd1, 0xeb, 0xdd, 0x1c, 0x2f, 0x65, 0x8c, 0xe8, 0x01, 0xd8,
	0x63, 0xca, 0xa2, 0xfc, 0xd1, 0xd6, 0x9c, 0x9b, 0xdc, 0x0f, 0xa1, 0x33, 0x5c, 0x70, 0xb9, 0x91,
	0xf3, 0x36, 0xf0, 0x4d, 0xeb, 0x63, 0xfe, 0x5f, 0x9f, 0x5a, 0xb8, 0x00, 0xce, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xb2, 0x85, 0x72, 0x04, 0xf4, 0x07, 0x00, 0x00,
}
