// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metacore/send.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SendStatus int32

const (
	SendStatus_Created   SendStatus = 0
	SendStatus_Finalized SendStatus = 1
	SendStatus_Mined     SendStatus = 2
	SendStatus_Confirmed SendStatus = 3
	SendStatus_Abort     SendStatus = 4
	SendStatus_Reverted  SendStatus = 5
	SendStatus_Aborted   SendStatus = 6
)

var SendStatus_name = map[int32]string{
	0: "Created",
	1: "Finalized",
	2: "Mined",
	3: "Confirmed",
	4: "Abort",
	5: "Reverted",
	6: "Aborted",
}

var SendStatus_value = map[string]int32{
	"Created":   0,
	"Finalized": 1,
	"Mined":     2,
	"Confirmed": 3,
	"Abort":     4,
	"Reverted":  5,
	"Aborted":   6,
}

func (x SendStatus) String() string {
	return proto.EnumName(SendStatus_name, int32(x))
}

func (SendStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c81f0328df818595, []int{0}
}

type Send struct {
	Creator             string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index               string     `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Sender              string     `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	SenderChain         string     `protobuf:"bytes,4,opt,name=senderChain,proto3" json:"senderChain,omitempty"`
	Receiver            string     `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	ReceiverChain       string     `protobuf:"bytes,6,opt,name=receiverChain,proto3" json:"receiverChain,omitempty"`
	MBurnt              string     `protobuf:"bytes,7,opt,name=mBurnt,proto3" json:"mBurnt,omitempty"`
	MMint               string     `protobuf:"bytes,8,opt,name=mMint,proto3" json:"mMint,omitempty"`
	Message             string     `protobuf:"bytes,9,opt,name=message,proto3" json:"message,omitempty"`
	InTxHash            string     `protobuf:"bytes,10,opt,name=inTxHash,proto3" json:"inTxHash,omitempty"`
	InBlockHeight       uint64     `protobuf:"varint,11,opt,name=inBlockHeight,proto3" json:"inBlockHeight,omitempty"`
	FinalizedMetaHeight uint64     `protobuf:"varint,12,opt,name=finalizedMetaHeight,proto3" json:"finalizedMetaHeight,omitempty"`
	Signers             []string   `protobuf:"bytes,13,rep,name=signers,proto3" json:"signers,omitempty"`
	Status              SendStatus `protobuf:"varint,14,opt,name=status,proto3,enum=MetaProtocol.metacore.metacore.SendStatus" json:"status,omitempty"`
	Broadcaster         uint64     `protobuf:"varint,15,opt,name=broadcaster,proto3" json:"broadcaster,omitempty"`
	Nonce               uint64     `protobuf:"varint,16,opt,name=nonce,proto3" json:"nonce,omitempty"`
	GasPrice            string     `protobuf:"bytes,17,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	StatusMessage       string     `protobuf:"bytes,18,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (m *Send) Reset()         { *m = Send{} }
func (m *Send) String() string { return proto.CompactTextString(m) }
func (*Send) ProtoMessage()    {}
func (*Send) Descriptor() ([]byte, []int) {
	return fileDescriptor_c81f0328df818595, []int{0}
}
func (m *Send) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Send) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Send.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Send) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send.Merge(m, src)
}
func (m *Send) XXX_Size() int {
	return m.Size()
}
func (m *Send) XXX_DiscardUnknown() {
	xxx_messageInfo_Send.DiscardUnknown(m)
}

var xxx_messageInfo_Send proto.InternalMessageInfo

func (m *Send) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Send) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Send) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Send) GetSenderChain() string {
	if m != nil {
		return m.SenderChain
	}
	return ""
}

func (m *Send) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Send) GetReceiverChain() string {
	if m != nil {
		return m.ReceiverChain
	}
	return ""
}

func (m *Send) GetMBurnt() string {
	if m != nil {
		return m.MBurnt
	}
	return ""
}

func (m *Send) GetMMint() string {
	if m != nil {
		return m.MMint
	}
	return ""
}

func (m *Send) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Send) GetInTxHash() string {
	if m != nil {
		return m.InTxHash
	}
	return ""
}

func (m *Send) GetInBlockHeight() uint64 {
	if m != nil {
		return m.InBlockHeight
	}
	return 0
}

func (m *Send) GetFinalizedMetaHeight() uint64 {
	if m != nil {
		return m.FinalizedMetaHeight
	}
	return 0
}

func (m *Send) GetSigners() []string {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *Send) GetStatus() SendStatus {
	if m != nil {
		return m.Status
	}
	return SendStatus_Created
}

func (m *Send) GetBroadcaster() uint64 {
	if m != nil {
		return m.Broadcaster
	}
	return 0
}

func (m *Send) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Send) GetGasPrice() string {
	if m != nil {
		return m.GasPrice
	}
	return ""
}

func (m *Send) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

func init() {
	proto.RegisterEnum("MetaProtocol.metacore.metacore.SendStatus", SendStatus_name, SendStatus_value)
	proto.RegisterType((*Send)(nil), "MetaProtocol.metacore.metacore.Send")
}

func init() { proto.RegisterFile("metacore/send.proto", fileDescriptor_c81f0328df818595) }

var fileDescriptor_c81f0328df818595 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x5b, 0xc7, 0x49, 0x36, 0x4d, 0x31, 0xdb, 0x0a, 0xad, 0x72, 0xb0, 0x2c, 0xc4, 0x21,
	0xaa, 0x84, 0x83, 0xca, 0x17, 0x90, 0x48, 0xa8, 0x07, 0x22, 0x55, 0x29, 0x27, 0x6e, 0x1b, 0x7b,
	0xea, 0xac, 0x88, 0x77, 0xab, 0xdd, 0x4d, 0x15, 0xf8, 0x0a, 0x3e, 0x82, 0x03, 0x5f, 0x82, 0x38,
	0xf6, 0xc8, 0x11, 0x25, 0x3f, 0x82, 0x66, 0xd7, 0x4e, 0x1b, 0x09, 0x71, 0x9b, 0xf7, 0xde, 0xcc,
	0xee, 0xf3, 0x1b, 0x2f, 0x39, 0xab, 0xc0, 0xf2, 0x5c, 0x69, 0x18, 0x1b, 0x90, 0x45, 0x76, 0xa7,
	0x95, 0x55, 0x34, 0x99, 0x81, 0xe5, 0xd7, 0x58, 0xe6, 0x6a, 0x95, 0x35, 0x1d, 0xfb, 0x62, 0x78,
	0x5e, 0xaa, 0x52, 0xb9, 0xd6, 0x31, 0x56, 0x7e, 0xea, 0xe5, 0xcf, 0x90, 0x84, 0x37, 0x20, 0x0b,
	0xca, 0x48, 0x27, 0xd7, 0xc0, 0xad, 0xd2, 0x2c, 0x48, 0x83, 0x51, 0x6f, 0xde, 0x40, 0x7a, 0x4e,
	0xda, 0x42, 0x16, 0xb0, 0x61, 0x47, 0x8e, 0xf7, 0x80, 0xbe, 0x20, 0x11, 0x5e, 0x0e, 0x9a, 0x1d,
	0x3b, 0xba, 0x46, 0x34, 0x25, 0x7d, 0x5f, 0x4d, 0x97, 0x5c, 0x48, 0x16, 0x3a, 0xf1, 0x29, 0x45,
	0x87, 0xa4, 0xab, 0x21, 0x07, 0x71, 0x0f, 0x9a, 0xb5, 0x9d, 0xbc, 0xc7, 0xf4, 0x15, 0x19, 0x34,
	0xb5, 0x9f, 0x8f, 0x5c, 0xc3, 0x21, 0x89, 0x77, 0x57, 0x93, 0xb5, 0x96, 0x96, 0x75, 0xfc, 0xdd,
	0x1e, 0xa1, 0xd3, 0x6a, 0x26, 0xa4, 0x65, 0x5d, 0xef, 0xd4, 0x01, 0xfc, 0xb2, 0x0a, 0x8c, 0xe1,
	0x25, 0xb0, 0x9e, 0xff, 0xb2, 0x1a, 0xa2, 0x13, 0x21, 0x3f, 0x6e, 0xae, 0xb8, 0x59, 0x32, 0xe2,
	0x9d, 0x34, 0x18, 0x9d, 0x08, 0x39, 0x59, 0xa9, 0xfc, 0xf3, 0x15, 0x88, 0x72, 0x69, 0x59, 0x3f,
	0x0d, 0x46, 0xe1, 0xfc, 0x90, 0xa4, 0x6f, 0xc8, 0xd9, 0xad, 0x90, 0x7c, 0x25, 0xbe, 0x42, 0x81,
	0xf9, 0xd7, 0xbd, 0x27, 0xae, 0xf7, 0x5f, 0x12, 0xba, 0x31, 0xa2, 0x94, 0xa0, 0x0d, 0x1b, 0xa4,
	0xc7, 0xe8, 0xa6, 0x86, 0x74, 0x42, 0x22, 0x63, 0xb9, 0x5d, 0x1b, 0x76, 0x9a, 0x06, 0xa3, 0xd3,
	0xcb, 0x8b, 0xec, 0xff, 0x1b, 0xcd, 0x70, 0x6f, 0x37, 0x6e, 0x62, 0x5e, 0x4f, 0x62, 0xfa, 0x0b,
	0xad, 0x78, 0x91, 0x73, 0x63, 0x41, 0xb3, 0x67, 0xce, 0xc7, 0x53, 0x0a, 0x33, 0x92, 0x4a, 0xe6,
	0xc0, 0x62, 0xa7, 0x79, 0x80, 0x49, 0x94, 0xdc, 0x5c, 0x6b, 0x91, 0x03, 0x7b, 0xee, 0x93, 0x68,
	0x30, 0x26, 0xe1, 0x4f, 0x9f, 0xd5, 0x29, 0x52, 0xbf, 0x93, 0x03, 0xf2, 0x42, 0x12, 0xf2, 0xe8,
	0x87, 0xf6, 0x49, 0x67, 0x8a, 0xbf, 0x0f, 0x14, 0x71, 0x8b, 0x0e, 0x48, 0xef, 0x7d, 0x93, 0x44,
	0x1c, 0xd0, 0x1e, 0x69, 0xcf, 0x84, 0x84, 0x22, 0x3e, 0x42, 0x65, 0xaa, 0xe4, 0xad, 0xd0, 0x15,
	0x14, 0xf1, 0x31, 0x2a, 0xef, 0x16, 0x4a, 0xdb, 0x38, 0xa4, 0x27, 0xa4, 0x3b, 0x87, 0x7b, 0xd0,
	0x78, 0x42, 0x1b, 0x8f, 0x73, 0x02, 0x14, 0x71, 0x34, 0x0c, 0x7f, 0x7c, 0x4f, 0x82, 0xc9, 0x87,
	0x5f, 0xdb, 0x24, 0x78, 0xd8, 0x26, 0xc1, 0x9f, 0x6d, 0x12, 0x7c, 0xdb, 0x25, 0xad, 0x87, 0x5d,
	0xd2, 0xfa, 0xbd, 0x4b, 0x5a, 0x9f, 0x2e, 0x4b, 0x61, 0x97, 0xeb, 0x45, 0x96, 0xab, 0x6a, 0x8c,
	0x09, 0xbe, 0x6e, 0x22, 0x1c, 0xef, 0x9f, 0xcd, 0xe6, 0xb1, 0xb4, 0x5f, 0xee, 0xc0, 0x2c, 0x22,
	0xf7, 0x1a, 0xde, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x36, 0x42, 0x92, 0x5a, 0x03, 0x00,
	0x00,
}

func (m *Send) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Send) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Send) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StatusMessage) > 0 {
		i -= len(m.StatusMessage)
		copy(dAtA[i:], m.StatusMessage)
		i = encodeVarintSend(dAtA, i, uint64(len(m.StatusMessage)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if len(m.GasPrice) > 0 {
		i -= len(m.GasPrice)
		copy(dAtA[i:], m.GasPrice)
		i = encodeVarintSend(dAtA, i, uint64(len(m.GasPrice)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if m.Nonce != 0 {
		i = encodeVarintSend(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.Broadcaster != 0 {
		i = encodeVarintSend(dAtA, i, uint64(m.Broadcaster))
		i--
		dAtA[i] = 0x78
	}
	if m.Status != 0 {
		i = encodeVarintSend(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x70
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintSend(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x6a
		}
	}
	if m.FinalizedMetaHeight != 0 {
		i = encodeVarintSend(dAtA, i, uint64(m.FinalizedMetaHeight))
		i--
		dAtA[i] = 0x60
	}
	if m.InBlockHeight != 0 {
		i = encodeVarintSend(dAtA, i, uint64(m.InBlockHeight))
		i--
		dAtA[i] = 0x58
	}
	if len(m.InTxHash) > 0 {
		i -= len(m.InTxHash)
		copy(dAtA[i:], m.InTxHash)
		i = encodeVarintSend(dAtA, i, uint64(len(m.InTxHash)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintSend(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.MMint) > 0 {
		i -= len(m.MMint)
		copy(dAtA[i:], m.MMint)
		i = encodeVarintSend(dAtA, i, uint64(len(m.MMint)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.MBurnt) > 0 {
		i -= len(m.MBurnt)
		copy(dAtA[i:], m.MBurnt)
		i = encodeVarintSend(dAtA, i, uint64(len(m.MBurnt)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ReceiverChain) > 0 {
		i -= len(m.ReceiverChain)
		copy(dAtA[i:], m.ReceiverChain)
		i = encodeVarintSend(dAtA, i, uint64(len(m.ReceiverChain)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintSend(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SenderChain) > 0 {
		i -= len(m.SenderChain)
		copy(dAtA[i:], m.SenderChain)
		i = encodeVarintSend(dAtA, i, uint64(len(m.SenderChain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintSend(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSend(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSend(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSend(dAtA []byte, offset int, v uint64) int {
	offset -= sovSend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Send) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	l = len(m.SenderChain)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	l = len(m.ReceiverChain)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	l = len(m.MBurnt)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	l = len(m.MMint)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	l = len(m.InTxHash)
	if l > 0 {
		n += 1 + l + sovSend(uint64(l))
	}
	if m.InBlockHeight != 0 {
		n += 1 + sovSend(uint64(m.InBlockHeight))
	}
	if m.FinalizedMetaHeight != 0 {
		n += 1 + sovSend(uint64(m.FinalizedMetaHeight))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovSend(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovSend(uint64(m.Status))
	}
	if m.Broadcaster != 0 {
		n += 1 + sovSend(uint64(m.Broadcaster))
	}
	if m.Nonce != 0 {
		n += 2 + sovSend(uint64(m.Nonce))
	}
	l = len(m.GasPrice)
	if l > 0 {
		n += 2 + l + sovSend(uint64(l))
	}
	l = len(m.StatusMessage)
	if l > 0 {
		n += 2 + l + sovSend(uint64(l))
	}
	return n
}

func sovSend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSend(x uint64) (n int) {
	return sovSend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Send) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Send: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Send: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiverChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiverChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MBurnt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MBurnt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MMint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MMint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InBlockHeight", wireType)
			}
			m.InBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedMetaHeight", wireType)
			}
			m.FinalizedMetaHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalizedMetaHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SendStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Broadcaster", wireType)
			}
			m.Broadcaster = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Broadcaster |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSend
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSend
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSend
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSend
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSend        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSend          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSend = fmt.Errorf("proto: unexpected end of group")
)