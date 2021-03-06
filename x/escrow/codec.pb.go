// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/escrow/codec.proto

/*
	Package escrow is a generated protocol buffer package.

	It is generated from these files:
		x/escrow/codec.proto

	It has these top-level messages:
		Escrow
		CreateEscrowMsg
		ReleaseEscrowMsg
		ReturnEscrowMsg
		UpdateEscrowPartiesMsg
*/
package escrow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import x "github.com/iov-one/weave/x"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Escrow holds some coins.
// The arbiter or sender can release them to the recipient.
// The recipient can return them to the sender.
// Upon timeout, they will be returned to the sender.
//
// Note that if the arbiter is a Hashlock permission, we have
// an HTLC ;)
type Escrow struct {
	// Sender, Arbiter, Recipient are all weave.Permission
	Sender    []byte `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Arbiter   []byte `protobuf:"bytes,2,opt,name=arbiter,proto3" json:"arbiter,omitempty"`
	Recipient []byte `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// amount may contain multiple token types
	Amount []*x.Coin `protobuf:"bytes,4,rep,name=amount" json:"amount,omitempty"`
	// if unreleased before timeout, will return to sender
	// timeout stored here is absolute block height
	Timeout int64 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// max length 128 character
	Memo string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *Escrow) Reset()                    { *m = Escrow{} }
func (m *Escrow) String() string            { return proto.CompactTextString(m) }
func (*Escrow) ProtoMessage()               {}
func (*Escrow) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{0} }

func (m *Escrow) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Escrow) GetArbiter() []byte {
	if m != nil {
		return m.Arbiter
	}
	return nil
}

func (m *Escrow) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Escrow) GetAmount() []*x.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Escrow) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Escrow) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

// CreateEscrowMsg is a request to create an Escrow with some tokens.
// If sender is not defined, it defaults to the first signer
// The rest must be defined
type CreateEscrowMsg struct {
	// Sender, Arbiter, Recipient are all weave.Permission
	Sender    []byte `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Arbiter   []byte `protobuf:"bytes,2,opt,name=arbiter,proto3" json:"arbiter,omitempty"`
	Recipient []byte `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// amount may contain multiple token types
	Amount []*x.Coin `protobuf:"bytes,4,rep,name=amount" json:"amount,omitempty"`
	// if unreleased before timeout, will return to sender
	Timeout int64 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// max length 128 character
	Memo string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *CreateEscrowMsg) Reset()                    { *m = CreateEscrowMsg{} }
func (m *CreateEscrowMsg) String() string            { return proto.CompactTextString(m) }
func (*CreateEscrowMsg) ProtoMessage()               {}
func (*CreateEscrowMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{1} }

func (m *CreateEscrowMsg) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *CreateEscrowMsg) GetArbiter() []byte {
	if m != nil {
		return m.Arbiter
	}
	return nil
}

func (m *CreateEscrowMsg) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *CreateEscrowMsg) GetAmount() []*x.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *CreateEscrowMsg) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *CreateEscrowMsg) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

// ReleaseEscrowMsg releases the content to the recipient.
// Must be authorized by sender or arbiter.
// If amount not provided, defaults to entire escrow,
// May be a subset of the current balance.
type ReleaseEscrowMsg struct {
	EscrowId []byte    `protobuf:"bytes,1,opt,name=escrow_id,json=escrowId,proto3" json:"escrow_id,omitempty"`
	Amount   []*x.Coin `protobuf:"bytes,2,rep,name=amount" json:"amount,omitempty"`
}

func (m *ReleaseEscrowMsg) Reset()                    { *m = ReleaseEscrowMsg{} }
func (m *ReleaseEscrowMsg) String() string            { return proto.CompactTextString(m) }
func (*ReleaseEscrowMsg) ProtoMessage()               {}
func (*ReleaseEscrowMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{2} }

func (m *ReleaseEscrowMsg) GetEscrowId() []byte {
	if m != nil {
		return m.EscrowId
	}
	return nil
}

func (m *ReleaseEscrowMsg) GetAmount() []*x.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

// ReturnEscrowMsg returns the content to the sender.
// Must be authorized by the sender or an expired timeout
type ReturnEscrowMsg struct {
	EscrowId []byte `protobuf:"bytes,1,opt,name=escrow_id,json=escrowId,proto3" json:"escrow_id,omitempty"`
}

func (m *ReturnEscrowMsg) Reset()                    { *m = ReturnEscrowMsg{} }
func (m *ReturnEscrowMsg) String() string            { return proto.CompactTextString(m) }
func (*ReturnEscrowMsg) ProtoMessage()               {}
func (*ReturnEscrowMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{3} }

func (m *ReturnEscrowMsg) GetEscrowId() []byte {
	if m != nil {
		return m.EscrowId
	}
	return nil
}

// UpdateEscrowPartiesMsg changes any of the parties of the escrow:
// sender, arbiter, recipient. This must be authorized by the current
// holder of that position (eg. only sender can update sender).
//
// Represents delegating responsibility
type UpdateEscrowPartiesMsg struct {
	EscrowId  []byte `protobuf:"bytes,1,opt,name=escrow_id,json=escrowId,proto3" json:"escrow_id,omitempty"`
	Sender    []byte `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Arbiter   []byte `protobuf:"bytes,3,opt,name=arbiter,proto3" json:"arbiter,omitempty"`
	Recipient []byte `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
}

func (m *UpdateEscrowPartiesMsg) Reset()                    { *m = UpdateEscrowPartiesMsg{} }
func (m *UpdateEscrowPartiesMsg) String() string            { return proto.CompactTextString(m) }
func (*UpdateEscrowPartiesMsg) ProtoMessage()               {}
func (*UpdateEscrowPartiesMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{4} }

func (m *UpdateEscrowPartiesMsg) GetEscrowId() []byte {
	if m != nil {
		return m.EscrowId
	}
	return nil
}

func (m *UpdateEscrowPartiesMsg) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *UpdateEscrowPartiesMsg) GetArbiter() []byte {
	if m != nil {
		return m.Arbiter
	}
	return nil
}

func (m *UpdateEscrowPartiesMsg) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func init() {
	proto.RegisterType((*Escrow)(nil), "escrow.Escrow")
	proto.RegisterType((*CreateEscrowMsg)(nil), "escrow.CreateEscrowMsg")
	proto.RegisterType((*ReleaseEscrowMsg)(nil), "escrow.ReleaseEscrowMsg")
	proto.RegisterType((*ReturnEscrowMsg)(nil), "escrow.ReturnEscrowMsg")
	proto.RegisterType((*UpdateEscrowPartiesMsg)(nil), "escrow.UpdateEscrowPartiesMsg")
}
func (m *Escrow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Escrow) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Sender)))
		i += copy(dAtA[i:], m.Sender)
	}
	if len(m.Arbiter) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Arbiter)))
		i += copy(dAtA[i:], m.Arbiter)
	}
	if len(m.Recipient) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Recipient)))
		i += copy(dAtA[i:], m.Recipient)
	}
	if len(m.Amount) > 0 {
		for _, msg := range m.Amount {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Timeout))
	}
	if len(m.Memo) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Memo)))
		i += copy(dAtA[i:], m.Memo)
	}
	return i, nil
}

func (m *CreateEscrowMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateEscrowMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Sender)))
		i += copy(dAtA[i:], m.Sender)
	}
	if len(m.Arbiter) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Arbiter)))
		i += copy(dAtA[i:], m.Arbiter)
	}
	if len(m.Recipient) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Recipient)))
		i += copy(dAtA[i:], m.Recipient)
	}
	if len(m.Amount) > 0 {
		for _, msg := range m.Amount {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Timeout))
	}
	if len(m.Memo) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Memo)))
		i += copy(dAtA[i:], m.Memo)
	}
	return i, nil
}

func (m *ReleaseEscrowMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReleaseEscrowMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.EscrowId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.EscrowId)))
		i += copy(dAtA[i:], m.EscrowId)
	}
	if len(m.Amount) > 0 {
		for _, msg := range m.Amount {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ReturnEscrowMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReturnEscrowMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.EscrowId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.EscrowId)))
		i += copy(dAtA[i:], m.EscrowId)
	}
	return i, nil
}

func (m *UpdateEscrowPartiesMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateEscrowPartiesMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.EscrowId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.EscrowId)))
		i += copy(dAtA[i:], m.EscrowId)
	}
	if len(m.Sender) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Sender)))
		i += copy(dAtA[i:], m.Sender)
	}
	if len(m.Arbiter) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Arbiter)))
		i += copy(dAtA[i:], m.Arbiter)
	}
	if len(m.Recipient) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Recipient)))
		i += copy(dAtA[i:], m.Recipient)
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Escrow) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Arbiter)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if m.Timeout != 0 {
		n += 1 + sovCodec(uint64(m.Timeout))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *CreateEscrowMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Arbiter)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if m.Timeout != 0 {
		n += 1 + sovCodec(uint64(m.Timeout))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *ReleaseEscrowMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.EscrowId)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *ReturnEscrowMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.EscrowId)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *UpdateEscrowPartiesMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.EscrowId)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Arbiter)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Escrow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Escrow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Escrow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arbiter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Arbiter = append(m.Arbiter[:0], dAtA[iNdEx:postIndex]...)
			if m.Arbiter == nil {
				m.Arbiter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = append(m.Recipient[:0], dAtA[iNdEx:postIndex]...)
			if m.Recipient == nil {
				m.Recipient = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, &x.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *CreateEscrowMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateEscrowMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateEscrowMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arbiter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Arbiter = append(m.Arbiter[:0], dAtA[iNdEx:postIndex]...)
			if m.Arbiter == nil {
				m.Arbiter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = append(m.Recipient[:0], dAtA[iNdEx:postIndex]...)
			if m.Recipient == nil {
				m.Recipient = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, &x.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *ReleaseEscrowMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReleaseEscrowMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReleaseEscrowMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowId = append(m.EscrowId[:0], dAtA[iNdEx:postIndex]...)
			if m.EscrowId == nil {
				m.EscrowId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, &x.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *ReturnEscrowMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReturnEscrowMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReturnEscrowMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowId = append(m.EscrowId[:0], dAtA[iNdEx:postIndex]...)
			if m.EscrowId == nil {
				m.EscrowId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *UpdateEscrowPartiesMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateEscrowPartiesMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateEscrowPartiesMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowId = append(m.EscrowId[:0], dAtA[iNdEx:postIndex]...)
			if m.EscrowId == nil {
				m.EscrowId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arbiter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Arbiter = append(m.Arbiter[:0], dAtA[iNdEx:postIndex]...)
			if m.Arbiter == nil {
				m.Arbiter = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = append(m.Recipient[:0], dAtA[iNdEx:postIndex]...)
			if m.Recipient == nil {
				m.Recipient = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("x/escrow/codec.proto", fileDescriptorCodec) }

var fileDescriptorCodec = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x4f, 0x4a, 0x03, 0x31,
	0x14, 0xc6, 0x4d, 0xa7, 0xa6, 0xf6, 0x29, 0xb4, 0x04, 0x29, 0x41, 0x65, 0x1c, 0x66, 0x21, 0xb3,
	0x71, 0x06, 0xf4, 0x06, 0x16, 0x17, 0x2e, 0x84, 0x12, 0x70, 0x2d, 0xe9, 0xcc, 0xa3, 0x06, 0x9c,
	0xa4, 0x64, 0x32, 0x6d, 0x2f, 0xe0, 0xde, 0x5b, 0xb8, 0xf5, 0x18, 0x2e, 0x3d, 0x82, 0xd4, 0x8b,
	0x88, 0x33, 0xad, 0x2d, 0xe2, 0xbf, 0xa5, 0xbb, 0x7c, 0xdf, 0x4b, 0xde, 0xfb, 0xf8, 0xe5, 0xc1,
	0xee, 0x2c, 0xc1, 0x22, 0xb5, 0x66, 0x9a, 0xa4, 0x26, 0xc3, 0x34, 0x1e, 0x5b, 0xe3, 0x0c, 0xa3,
	0xb5, 0xb7, 0x77, 0x34, 0x52, 0xee, 0xa6, 0x1c, 0xc6, 0xa9, 0xc9, 0x13, 0x65, 0x26, 0xc7, 0x46,
	0x63, 0x32, 0x45, 0x39, 0xc1, 0x64, 0xb6, 0x7e, 0x3f, 0x7c, 0x20, 0x40, 0xcf, 0xab, 0x27, 0xac,
	0x07, 0xb4, 0x40, 0x9d, 0xa1, 0xe5, 0x24, 0x20, 0xd1, 0x8e, 0x58, 0x28, 0xc6, 0xa1, 0x25, 0xed,
	0x50, 0x39, 0xb4, 0xbc, 0x51, 0x15, 0x96, 0x92, 0x1d, 0x40, 0xdb, 0x62, 0xaa, 0xc6, 0x0a, 0xb5,
	0xe3, 0x5e, 0x55, 0x5b, 0x19, 0xec, 0x10, 0xa8, 0xcc, 0x4d, 0xa9, 0x1d, 0x6f, 0x06, 0x5e, 0xb4,
	0x7d, 0xd2, 0x8a, 0x67, 0x71, 0xdf, 0x28, 0x2d, 0x16, 0xf6, 0x7b, 0x63, 0xa7, 0x72, 0x34, 0xa5,
	0xe3, 0x9b, 0x01, 0x89, 0x3c, 0xb1, 0x94, 0x8c, 0x41, 0x33, 0xc7, 0xdc, 0x70, 0x1a, 0x90, 0xa8,
	0x2d, 0xaa, 0x73, 0xf8, 0x48, 0xa0, 0xd3, 0xb7, 0x28, 0x1d, 0xd6, 0x79, 0x2f, 0x8b, 0xd1, 0x7f,
	0x8f, 0x3c, 0x80, 0xae, 0xc0, 0x5b, 0x94, 0xc5, 0x5a, 0xe4, 0x7d, 0x68, 0xd7, 0x5f, 0x74, 0xad,
	0xb2, 0x45, 0xea, 0xad, 0xda, 0xb8, 0xc8, 0xd6, 0xe6, 0x37, 0xbe, 0x9c, 0x1f, 0xc6, 0xd0, 0x11,
	0xe8, 0x4a, 0xab, 0xff, 0xd6, 0x30, 0xbc, 0x23, 0xd0, 0xbb, 0x1a, 0x67, 0x1f, 0xd0, 0x06, 0xd2,
	0x3a, 0x85, 0xc5, 0xaf, 0x41, 0x56, 0x60, 0x1b, 0xdf, 0x81, 0xf5, 0x7e, 0x00, 0xdb, 0xfc, 0x04,
	0xf6, 0xac, 0xfb, 0x34, 0xf7, 0xc9, 0xf3, 0xdc, 0x27, 0x2f, 0x73, 0x9f, 0xdc, 0xbf, 0xfa, 0x1b,
	0x43, 0x5a, 0xed, 0xdf, 0xe9, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0xfe, 0xa2, 0x61, 0xc7,
	0x02, 0x00, 0x00,
}
