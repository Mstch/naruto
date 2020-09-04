// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg.proto

package msg

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

type Cmd_Opt int32

const (
	Get Cmd_Opt = 0
	Set Cmd_Opt = 1
)

var Cmd_Opt_name = map[int32]string{
	0: "Get",
	1: "Set",
}

var Cmd_Opt_value = map[string]int32{
	"Get": 0,
	"Set": 1,
}

func (Cmd_Opt) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0, 0}
}

type Cmd struct {
	Opt   Cmd_Opt `protobuf:"varint,1,opt,name=opt,proto3,enum=Cmd_Opt" json:"opt,omitempty"`
	Key   string  `protobuf:"buffer,2,opt,name=key,proto3" json:"key,omitempty"`
	Value string  `protobuf:"buffer,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Cmd) Reset()      { *m = Cmd{} }
func (*Cmd) ProtoMessage() {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}
func (m *Cmd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cmd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmd.Merge(m, src)
}
func (m *Cmd) XXX_Size() int {
	return m.Size()
}
func (m *Cmd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmd.DiscardUnknown(m)
}

var xxx_messageInfo_Cmd proto.InternalMessageInfo

func (m *Cmd) GetOpt() Cmd_Opt {
	if m != nil {
		return m.Opt
	}
	return Get
}

func (m *Cmd) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Cmd) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CmdResp struct {
	Res        string `protobuf:"buffer,1,opt,name=res,proto3" json:"res,omitempty"`
	IsLeader   bool   `protobuf:"varint,2,opt,name=isLeader,proto3" json:"isLeader,omitempty"`
	LeaderAddr string `protobuf:"buffer,3,opt,name=leaderAddr,proto3" json:"leaderAddr,omitempty"`
}

func (m *CmdResp) Reset()      { *m = CmdResp{} }
func (*CmdResp) ProtoMessage() {}
func (*CmdResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}
func (m *CmdResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CmdResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CmdResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CmdResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdResp.Merge(m, src)
}
func (m *CmdResp) XXX_Size() int {
	return m.Size()
}
func (m *CmdResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdResp.DiscardUnknown(m)
}

var xxx_messageInfo_CmdResp proto.InternalMessageInfo

func (m *CmdResp) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func (m *CmdResp) GetIsLeader() bool {
	if m != nil {
		return m.IsLeader
	}
	return false
}

func (m *CmdResp) GetLeaderAddr() string {
	if m != nil {
		return m.LeaderAddr
	}
	return ""
}

func init() {
	proto.RegisterEnum("Cmd_Opt", Cmd_Opt_name, Cmd_Opt_value)
	proto.RegisterType((*Cmd)(nil), "Cmd")
	proto.RegisterType((*CmdResp)(nil), "CmdResp")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 238 buffer of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x77, 0x5c, 0xf4, 0x92, 0x29, 0x24, 0x2c, 0x82, 0xe1, 0x8a, 0xe1, 0x48, 0x75, 0x55,
	0x0a, 0xf5, 0x05, 0x34, 0x85, 0x8d, 0x70, 0xb0, 0x16, 0xd6, 0x77, 0xec, 0x20, 0xe2, 0x85, 0x2c,
	0xd9, 0x55, 0xb0, 0xf3, 0x11, 0x7c, 0x0c, 0x1f, 0xc5, 0x32, 0xe5, 0x95, 0x66, 0xd3, 0x58, 0xde,
	0x23, 0xc8, 0xae, 0x22, 0x76, 0xdf, 0xff, 0x15, 0xf3, 0x31, 0x98, 0xb7, 0xee, 0xbe, 0xb6, 0x7d,
	0xe7, 0xbb, 0x6a, 0x83, 0xb2, 0x69, 0x8d, 0x9a, 0xa3, 0xec, 0xac, 0x2f, 0x61, 0x01, 0xcb, 0xe3,
	0xb3, 0xac, 0x6e, 0x5a, 0x53, 0xaf, 0xac, 0xd7, 0x51, 0xaa, 0x02, 0xe5, 0x23, 0xbf, 0x94, 0x07,
	0x0b, 0x58, 0xe6, 0x3a, 0xa2, 0x3a, 0xc1, 0xc3, 0xe7, 0xf5, 0xf6, 0x89, 0x4b, 0x99, 0xdc, 0xcf,
	0xa8, 0x4e, 0x51, 0xae, 0xac, 0x57, 0x33, 0x94, 0xd7, 0xec, 0x0b, 0x11, 0xe1, 0x96, 0x7d, 0x01,
	0xd5, 0x1d, 0xce, 0x9a, 0xd6, 0x68, 0x76, 0x36, 0xde, 0xea, 0xd9, 0xa5, 0x4e, 0xae, 0x23, 0xaa,
	0x39, 0x66, 0x0f, 0xee, 0x86, 0xd7, 0x86, 0xfb, 0x94, 0xc8, 0xf4, 0xdf, 0x56, 0x84, 0xb8, 0x4d,
	0x74, 0x69, 0x4c, 0xff, 0x1b, 0xfb, 0x67, 0xae, 0x2e, 0x86, 0x91, 0xc4, 0x6e, 0x24, 0xb1, 0x1f,
	0x09, 0x5e, 0x03, 0xc1, 0x7b, 0x20, 0xf8, 0x08, 0x04, 0x43, 0x20, 0xf8, 0x0c, 0x04, 0x5f, 0x81,
	0xc4, 0x3e, 0x10, 0xbc, 0x4d, 0x24, 0x86, 0x89, 0xc4, 0x6e, 0x22, 0xb1, 0x39, 0x4a, 0x9f, 0x9f,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x42, 0xf5, 0x1a, 0x49, 0x06, 0x01, 0x00, 0x00,
}

func (x Cmd_Opt) String() string {
	s, ok := Cmd_Opt_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Cmd) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Cmd)
	if !ok {
		that2, ok := that.(Cmd)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Opt != that1.Opt {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *CmdResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CmdResp)
	if !ok {
		that2, ok := that.(CmdResp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Res != that1.Res {
		return false
	}
	if this.IsLeader != that1.IsLeader {
		return false
	}
	if this.LeaderAddr != that1.LeaderAddr {
		return false
	}
	return true
}
func (this *Cmd) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&msg.Cmd{")
	s = append(s, "Opt: "+fmt.Sprintf("%#v", this.Opt)+",\n")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CmdResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&msg.CmdResp{")
	s = append(s, "Res: "+fmt.Sprintf("%#v", this.Res)+",\n")
	s = append(s, "IsLeader: "+fmt.Sprintf("%#v", this.IsLeader)+",\n")
	s = append(s, "LeaderAddr: "+fmt.Sprintf("%#v", this.LeaderAddr)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMsg(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Cmd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cmd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cmd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if m.Opt != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Opt))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CmdResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CmdResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CmdResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LeaderAddr) > 0 {
		i -= len(m.LeaderAddr)
		copy(dAtA[i:], m.LeaderAddr)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.LeaderAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.IsLeader {
		i--
		if m.IsLeader {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Res) > 0 {
		i -= len(m.Res)
		copy(dAtA[i:], m.Res)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Res)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cmd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Opt != 0 {
		n += 1 + sovMsg(uint64(m.Opt))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *CmdResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Res)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.IsLeader {
		n += 2
	}
	l = len(m.LeaderAddr)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Cmd) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Cmd{`,
		`Opt:` + fmt.Sprintf("%v", this.Opt) + `,`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CmdResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CmdResp{`,
		`Res:` + fmt.Sprintf("%v", this.Res) + `,`,
		`IsLeader:` + fmt.Sprintf("%v", this.IsLeader) + `,`,
		`LeaderAddr:` + fmt.Sprintf("%v", this.LeaderAddr) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMsg(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Cmd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: Cmd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cmd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Opt", wireType)
			}
			m.Opt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Opt |= Cmd_Opt(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *CmdResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: CmdResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CmdResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Res", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Res = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLeader", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLeader = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LeaderAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)