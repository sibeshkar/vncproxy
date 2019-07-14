// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package proto

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

type InitMsg struct {
	RfbHeader            string       `protobuf:"bytes,1,opt,name=RfbHeader,json=rfbHeader,proto3" json:"RfbHeader,omitempty"`
	RfbVersion           string       `protobuf:"bytes,2,opt,name=RfbVersion,json=rfbVersion,proto3" json:"RfbVersion,omitempty"`
	FBHeight             uint32       `protobuf:"varint,3,opt,name=FBHeight,json=fBHeight,proto3" json:"FBHeight,omitempty"`
	FBWidth              uint32       `protobuf:"varint,4,opt,name=FBWidth,json=fBWidth,proto3" json:"FBWidth,omitempty"`
	SecType              uint32       `protobuf:"varint,5,opt,name=SecType,json=secType,proto3" json:"SecType,omitempty"`
	StartTime            uint32       `protobuf:"varint,6,opt,name=StartTime,json=startTime,proto3" json:"StartTime,omitempty"`
	DesktopName          string       `protobuf:"bytes,7,opt,name=DesktopName,json=desktopName,proto3" json:"DesktopName,omitempty"`
	PixelFormat          *PixelFormat `protobuf:"bytes,8,opt,name=PixelFormat,json=pixelFormat,proto3" json:"PixelFormat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InitMsg) Reset()         { *m = InitMsg{} }
func (m *InitMsg) String() string { return proto.CompactTextString(m) }
func (*InitMsg) ProtoMessage()    {}
func (*InitMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *InitMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitMsg.Unmarshal(m, b)
}
func (m *InitMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitMsg.Marshal(b, m, deterministic)
}
func (m *InitMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitMsg.Merge(m, src)
}
func (m *InitMsg) XXX_Size() int {
	return xxx_messageInfo_InitMsg.Size(m)
}
func (m *InitMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_InitMsg.DiscardUnknown(m)
}

var xxx_messageInfo_InitMsg proto.InternalMessageInfo

func (m *InitMsg) GetRfbHeader() string {
	if m != nil {
		return m.RfbHeader
	}
	return ""
}

func (m *InitMsg) GetRfbVersion() string {
	if m != nil {
		return m.RfbVersion
	}
	return ""
}

func (m *InitMsg) GetFBHeight() uint32 {
	if m != nil {
		return m.FBHeight
	}
	return 0
}

func (m *InitMsg) GetFBWidth() uint32 {
	if m != nil {
		return m.FBWidth
	}
	return 0
}

func (m *InitMsg) GetSecType() uint32 {
	if m != nil {
		return m.SecType
	}
	return 0
}

func (m *InitMsg) GetStartTime() uint32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *InitMsg) GetDesktopName() string {
	if m != nil {
		return m.DesktopName
	}
	return ""
}

func (m *InitMsg) GetPixelFormat() *PixelFormat {
	if m != nil {
		return m.PixelFormat
	}
	return nil
}

type PointerEvent struct {
	Mask                 uint32   `protobuf:"varint,1,opt,name=Mask,json=mask,proto3" json:"Mask,omitempty"`
	X                    uint32   `protobuf:"varint,2,opt,name=X,json=x,proto3" json:"X,omitempty"`
	Y                    uint32   `protobuf:"varint,3,opt,name=Y,json=y,proto3" json:"Y,omitempty"`
	Timestamp            uint32   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PointerEvent) Reset()         { *m = PointerEvent{} }
func (m *PointerEvent) String() string { return proto.CompactTextString(m) }
func (*PointerEvent) ProtoMessage()    {}
func (*PointerEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *PointerEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointerEvent.Unmarshal(m, b)
}
func (m *PointerEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointerEvent.Marshal(b, m, deterministic)
}
func (m *PointerEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointerEvent.Merge(m, src)
}
func (m *PointerEvent) XXX_Size() int {
	return xxx_messageInfo_PointerEvent.Size(m)
}
func (m *PointerEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PointerEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PointerEvent proto.InternalMessageInfo

func (m *PointerEvent) GetMask() uint32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *PointerEvent) GetX() uint32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *PointerEvent) GetY() uint32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *PointerEvent) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type KeyEvent struct {
	Down                 uint32   `protobuf:"varint,1,opt,name=Down,json=down,proto3" json:"Down,omitempty"`
	Key                  uint32   `protobuf:"varint,2,opt,name=Key,json=key,proto3" json:"Key,omitempty"`
	Timestamp            uint32   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEvent) Reset()         { *m = KeyEvent{} }
func (m *KeyEvent) String() string { return proto.CompactTextString(m) }
func (*KeyEvent) ProtoMessage()    {}
func (*KeyEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{2}
}

func (m *KeyEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyEvent.Unmarshal(m, b)
}
func (m *KeyEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyEvent.Marshal(b, m, deterministic)
}
func (m *KeyEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEvent.Merge(m, src)
}
func (m *KeyEvent) XXX_Size() int {
	return xxx_messageInfo_KeyEvent.Size(m)
}
func (m *KeyEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEvent.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEvent proto.InternalMessageInfo

func (m *KeyEvent) GetDown() uint32 {
	if m != nil {
		return m.Down
	}
	return 0
}

func (m *KeyEvent) GetKey() uint32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *KeyEvent) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type FbsSegment struct {
	Bytes                []byte   `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Timestamp            uint32   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FbsSegment) Reset()         { *m = FbsSegment{} }
func (m *FbsSegment) String() string { return proto.CompactTextString(m) }
func (*FbsSegment) ProtoMessage()    {}
func (*FbsSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{3}
}

func (m *FbsSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FbsSegment.Unmarshal(m, b)
}
func (m *FbsSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FbsSegment.Marshal(b, m, deterministic)
}
func (m *FbsSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FbsSegment.Merge(m, src)
}
func (m *FbsSegment) XXX_Size() int {
	return xxx_messageInfo_FbsSegment.Size(m)
}
func (m *FbsSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_FbsSegment.DiscardUnknown(m)
}

var xxx_messageInfo_FbsSegment proto.InternalMessageInfo

func (m *FbsSegment) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *FbsSegment) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PixelFormat struct {
	BPP                  uint32   `protobuf:"varint,1,opt,name=BPP,json=bPP,proto3" json:"BPP,omitempty"`
	Depth                uint32   `protobuf:"varint,2,opt,name=Depth,json=depth,proto3" json:"Depth,omitempty"`
	BigEndian            uint32   `protobuf:"varint,3,opt,name=BigEndian,json=bigEndian,proto3" json:"BigEndian,omitempty"`
	TrueColor            uint32   `protobuf:"varint,4,opt,name=TrueColor,json=trueColor,proto3" json:"TrueColor,omitempty"`
	RedMax               uint32   `protobuf:"varint,5,opt,name=RedMax,json=redMax,proto3" json:"RedMax,omitempty"`
	GreenMax             uint32   `protobuf:"varint,6,opt,name=GreenMax,json=greenMax,proto3" json:"GreenMax,omitempty"`
	BlueMax              uint32   `protobuf:"varint,7,opt,name=BlueMax,json=blueMax,proto3" json:"BlueMax,omitempty"`
	RedShift             uint32   `protobuf:"varint,8,opt,name=RedShift,json=redShift,proto3" json:"RedShift,omitempty"`
	GreenShift           uint32   `protobuf:"varint,9,opt,name=GreenShift,json=greenShift,proto3" json:"GreenShift,omitempty"`
	BlueShift            uint32   `protobuf:"varint,10,opt,name=BlueShift,json=blueShift,proto3" json:"BlueShift,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PixelFormat) Reset()         { *m = PixelFormat{} }
func (m *PixelFormat) String() string { return proto.CompactTextString(m) }
func (*PixelFormat) ProtoMessage()    {}
func (*PixelFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{4}
}

func (m *PixelFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PixelFormat.Unmarshal(m, b)
}
func (m *PixelFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PixelFormat.Marshal(b, m, deterministic)
}
func (m *PixelFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PixelFormat.Merge(m, src)
}
func (m *PixelFormat) XXX_Size() int {
	return xxx_messageInfo_PixelFormat.Size(m)
}
func (m *PixelFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_PixelFormat.DiscardUnknown(m)
}

var xxx_messageInfo_PixelFormat proto.InternalMessageInfo

func (m *PixelFormat) GetBPP() uint32 {
	if m != nil {
		return m.BPP
	}
	return 0
}

func (m *PixelFormat) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *PixelFormat) GetBigEndian() uint32 {
	if m != nil {
		return m.BigEndian
	}
	return 0
}

func (m *PixelFormat) GetTrueColor() uint32 {
	if m != nil {
		return m.TrueColor
	}
	return 0
}

func (m *PixelFormat) GetRedMax() uint32 {
	if m != nil {
		return m.RedMax
	}
	return 0
}

func (m *PixelFormat) GetGreenMax() uint32 {
	if m != nil {
		return m.GreenMax
	}
	return 0
}

func (m *PixelFormat) GetBlueMax() uint32 {
	if m != nil {
		return m.BlueMax
	}
	return 0
}

func (m *PixelFormat) GetRedShift() uint32 {
	if m != nil {
		return m.RedShift
	}
	return 0
}

func (m *PixelFormat) GetGreenShift() uint32 {
	if m != nil {
		return m.GreenShift
	}
	return 0
}

func (m *PixelFormat) GetBlueShift() uint32 {
	if m != nil {
		return m.BlueShift
	}
	return 0
}

type Rectangle struct {
	X                    uint32   `protobuf:"varint,1,opt,name=X,json=x,proto3" json:"X,omitempty"`
	Y                    uint32   `protobuf:"varint,2,opt,name=Y,json=y,proto3" json:"Y,omitempty"`
	Width                uint32   `protobuf:"varint,3,opt,name=Width,json=width,proto3" json:"Width,omitempty"`
	Height               uint32   `protobuf:"varint,4,opt,name=Height,json=height,proto3" json:"Height,omitempty"`
	Enc                  uint32   `protobuf:"varint,5,opt,name=Enc,json=enc,proto3" json:"Enc,omitempty"`
	Bytes                []byte   `protobuf:"bytes,6,opt,name=Bytes,json=bytes,proto3" json:"Bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}
func (*Rectangle) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{5}
}

func (m *Rectangle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rectangle.Unmarshal(m, b)
}
func (m *Rectangle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rectangle.Marshal(b, m, deterministic)
}
func (m *Rectangle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rectangle.Merge(m, src)
}
func (m *Rectangle) XXX_Size() int {
	return xxx_messageInfo_Rectangle.Size(m)
}
func (m *Rectangle) XXX_DiscardUnknown() {
	xxx_messageInfo_Rectangle.DiscardUnknown(m)
}

var xxx_messageInfo_Rectangle proto.InternalMessageInfo

func (m *Rectangle) GetX() uint32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Rectangle) GetY() uint32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Rectangle) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Rectangle) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Rectangle) GetEnc() uint32 {
	if m != nil {
		return m.Enc
	}
	return 0
}

func (m *Rectangle) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

type FramebufferUpdate struct {
	Rectangles           []*Rectangle `protobuf:"bytes,1,rep,name=rectangles,proto3" json:"rectangles,omitempty"`
	Timestamp            uint32       `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FramebufferUpdate) Reset()         { *m = FramebufferUpdate{} }
func (m *FramebufferUpdate) String() string { return proto.CompactTextString(m) }
func (*FramebufferUpdate) ProtoMessage()    {}
func (*FramebufferUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{6}
}

func (m *FramebufferUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FramebufferUpdate.Unmarshal(m, b)
}
func (m *FramebufferUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FramebufferUpdate.Marshal(b, m, deterministic)
}
func (m *FramebufferUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FramebufferUpdate.Merge(m, src)
}
func (m *FramebufferUpdate) XXX_Size() int {
	return xxx_messageInfo_FramebufferUpdate.Size(m)
}
func (m *FramebufferUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_FramebufferUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_FramebufferUpdate proto.InternalMessageInfo

func (m *FramebufferUpdate) GetRectangles() []*Rectangle {
	if m != nil {
		return m.Rectangles
	}
	return nil
}

func (m *FramebufferUpdate) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Demonstration struct {
	Initmsg              *InitMsg             `protobuf:"bytes,1,opt,name=initmsg,proto3" json:"initmsg,omitempty"`
	Fbupdates            []*FramebufferUpdate `protobuf:"bytes,2,rep,name=fbupdates,proto3" json:"fbupdates,omitempty"`
	Segments             []*FbsSegment        `protobuf:"bytes,3,rep,name=segments,proto3" json:"segments,omitempty"`
	Pointerevents        []*PointerEvent      `protobuf:"bytes,4,rep,name=pointerevents,proto3" json:"pointerevents,omitempty"`
	Keyevents            []*KeyEvent          `protobuf:"bytes,5,rep,name=keyevents,proto3" json:"keyevents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Demonstration) Reset()         { *m = Demonstration{} }
func (m *Demonstration) String() string { return proto.CompactTextString(m) }
func (*Demonstration) ProtoMessage()    {}
func (*Demonstration) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{7}
}

func (m *Demonstration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Demonstration.Unmarshal(m, b)
}
func (m *Demonstration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Demonstration.Marshal(b, m, deterministic)
}
func (m *Demonstration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Demonstration.Merge(m, src)
}
func (m *Demonstration) XXX_Size() int {
	return xxx_messageInfo_Demonstration.Size(m)
}
func (m *Demonstration) XXX_DiscardUnknown() {
	xxx_messageInfo_Demonstration.DiscardUnknown(m)
}

var xxx_messageInfo_Demonstration proto.InternalMessageInfo

func (m *Demonstration) GetInitmsg() *InitMsg {
	if m != nil {
		return m.Initmsg
	}
	return nil
}

func (m *Demonstration) GetFbupdates() []*FramebufferUpdate {
	if m != nil {
		return m.Fbupdates
	}
	return nil
}

func (m *Demonstration) GetSegments() []*FbsSegment {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *Demonstration) GetPointerevents() []*PointerEvent {
	if m != nil {
		return m.Pointerevents
	}
	return nil
}

func (m *Demonstration) GetKeyevents() []*KeyEvent {
	if m != nil {
		return m.Keyevents
	}
	return nil
}

func init() {
	proto.RegisterType((*InitMsg)(nil), "proto.InitMsg")
	proto.RegisterType((*PointerEvent)(nil), "proto.PointerEvent")
	proto.RegisterType((*KeyEvent)(nil), "proto.KeyEvent")
	proto.RegisterType((*FbsSegment)(nil), "proto.FbsSegment")
	proto.RegisterType((*PixelFormat)(nil), "proto.PixelFormat")
	proto.RegisterType((*Rectangle)(nil), "proto.Rectangle")
	proto.RegisterType((*FramebufferUpdate)(nil), "proto.FramebufferUpdate")
	proto.RegisterType((*Demonstration)(nil), "proto.Demonstration")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x45, 0x96, 0x5f, 0xba, 0x8a, 0xdb, 0x64, 0x1a, 0x82, 0x28, 0xa5, 0x18, 0xaf, 0xbc, 0x49,
	0x28, 0x69, 0x29, 0x74, 0x57, 0x5c, 0xc7, 0x4d, 0x09, 0x09, 0x66, 0x9c, 0xbe, 0xa0, 0x1b, 0xc9,
	0xba, 0x92, 0x85, 0xad, 0x91, 0x98, 0x19, 0x27, 0xf1, 0xa2, 0x3f, 0x50, 0xe8, 0x07, 0xf4, 0x6f,
	0xcb, 0x3c, 0x24, 0xdb, 0xe9, 0xa2, 0x2b, 0x71, 0xee, 0x99, 0x7b, 0x38, 0xf7, 0x25, 0x80, 0x18,
	0xf3, 0xe2, 0xac, 0xe4, 0x85, 0x2c, 0x48, 0x4b, 0x7f, 0x06, 0xbf, 0x1b, 0xd0, 0xf9, 0xc4, 0x32,
	0x79, 0x2d, 0x52, 0xf2, 0x02, 0x3c, 0x9a, 0x44, 0x97, 0x18, 0xc6, 0xc8, 0x03, 0xa7, 0xef, 0x0c,
	0x3d, 0xea, 0xf1, 0x2a, 0x40, 0x5e, 0x02, 0xd0, 0x24, 0xfa, 0x82, 0x5c, 0x64, 0x05, 0x0b, 0x1a,
	0x9a, 0x06, 0x5e, 0x47, 0xc8, 0x73, 0xe8, 0x4e, 0x46, 0x97, 0x98, 0xa5, 0x0b, 0x19, 0xb8, 0x7d,
	0x67, 0xd8, 0xa3, 0xdd, 0xc4, 0x62, 0x12, 0x40, 0x67, 0x32, 0xfa, 0x9a, 0xc5, 0x72, 0x11, 0x34,
	0x35, 0xd5, 0x49, 0x0c, 0x54, 0xcc, 0x0c, 0xe7, 0xb7, 0x9b, 0x12, 0x83, 0x96, 0x61, 0x84, 0x81,
	0xca, 0xcd, 0x4c, 0x86, 0x5c, 0xde, 0x66, 0x39, 0x06, 0x6d, 0xcd, 0x79, 0xa2, 0x0a, 0x90, 0x3e,
	0xf8, 0x63, 0x14, 0x4b, 0x59, 0x94, 0x37, 0x61, 0x8e, 0x41, 0x47, 0xdb, 0xf1, 0xe3, 0x6d, 0x88,
	0xbc, 0x01, 0x7f, 0x9a, 0x3d, 0xe0, 0x6a, 0x52, 0xf0, 0x3c, 0x94, 0x41, 0xb7, 0xef, 0x0c, 0xfd,
	0x73, 0x62, 0xaa, 0x3f, 0xdb, 0x61, 0xa8, 0x5f, 0x6e, 0xc1, 0xe0, 0x07, 0x1c, 0x4c, 0x8b, 0x8c,
	0x49, 0xe4, 0x17, 0x77, 0xc8, 0x24, 0x21, 0xd0, 0xbc, 0x0e, 0xc5, 0x52, 0xb7, 0xa3, 0x47, 0x9b,
	0x79, 0x28, 0x96, 0xe4, 0x00, 0x9c, 0x6f, 0xba, 0x01, 0x3d, 0xea, 0x3c, 0x28, 0xf4, 0xdd, 0x16,
	0xec, 0x6c, 0x94, 0x6b, 0x99, 0xe5, 0x28, 0x64, 0x98, 0x97, 0xb6, 0xd6, 0x6d, 0x60, 0x70, 0x03,
	0xdd, 0x2b, 0xdc, 0xd4, 0xca, 0xe3, 0xe2, 0x9e, 0x55, 0xca, 0x71, 0x71, 0xcf, 0xc8, 0x21, 0xb8,
	0x57, 0xb8, 0xb1, 0xda, 0xee, 0x12, 0x1f, 0xe9, 0xb9, 0x8f, 0xf5, 0xde, 0x03, 0x4c, 0x22, 0x31,
	0xc3, 0x34, 0x57, 0x8a, 0xc7, 0xd0, 0x8a, 0x36, 0x12, 0x85, 0x96, 0x3c, 0xa0, 0x06, 0xec, 0x2b,
	0x34, 0x1e, 0x2b, 0xfc, 0x69, 0xec, 0xb5, 0x49, 0x39, 0x18, 0x4d, 0xa7, 0xd6, 0x94, 0x1b, 0x4d,
	0xa7, 0x4a, 0x75, 0x8c, 0xa5, 0x5c, 0xd8, 0xdc, 0x56, 0xac, 0x80, 0x52, 0x1d, 0x65, 0xe9, 0x05,
	0x8b, 0xb3, 0x90, 0x55, 0xbe, 0xa2, 0x2a, 0xa0, 0xd8, 0x5b, 0xbe, 0xc6, 0x0f, 0xc5, 0xaa, 0xe0,
	0x75, 0x17, 0xaa, 0x00, 0x39, 0x81, 0x36, 0xc5, 0xf8, 0x3a, 0x7c, 0xb0, 0x23, 0x6f, 0x73, 0x8d,
	0xd4, 0x06, 0x7d, 0xe4, 0x88, 0x4c, 0x31, 0x66, 0xe0, 0xdd, 0xd4, 0x62, 0xb5, 0x27, 0xa3, 0xd5,
	0x1a, 0x15, 0xd5, 0x31, 0x7b, 0x12, 0x19, 0xa8, 0xb2, 0x28, 0xc6, 0xb3, 0x45, 0x96, 0x98, 0x21,
	0xf7, 0x68, 0x97, 0x5b, 0xac, 0x76, 0x56, 0x2b, 0x1a, 0xd6, 0xd3, 0x2c, 0xa4, 0x75, 0x44, 0x57,
	0xb1, 0x5a, 0xa3, 0xa1, 0xc1, 0x56, 0x51, 0x05, 0x06, 0x3f, 0xc1, 0xa3, 0x38, 0x97, 0x21, 0x4b,
	0x57, 0x68, 0x86, 0xee, 0xec, 0x0d, 0xbd, 0x51, 0x0d, 0xfd, 0x18, 0x5a, 0x66, 0xb9, 0x4d, 0x23,
	0x5a, 0xf7, 0x7a, 0xb5, 0x4f, 0xa0, 0x6d, 0xcf, 0xc1, 0x74, 0xa0, 0xbd, 0x30, 0xc7, 0x70, 0x08,
	0xee, 0x05, 0x9b, 0xdb, 0xda, 0x5d, 0x64, 0x73, 0x95, 0x3f, 0xd2, 0x83, 0x6b, 0xef, 0x0c, 0x6e,
	0x30, 0x87, 0xa3, 0x09, 0x0f, 0x73, 0x8c, 0xd6, 0x49, 0x82, 0xfc, 0x73, 0x19, 0x87, 0x12, 0xc9,
	0x2b, 0x00, 0x5e, 0x79, 0x52, 0x83, 0x76, 0x87, 0xfe, 0xf9, 0xa1, 0x5d, 0xea, 0xda, 0x2c, 0xdd,
	0x79, 0xf3, 0x9f, 0xf9, 0xff, 0x6a, 0x40, 0x6f, 0x8c, 0x79, 0xc1, 0x84, 0xe4, 0xa1, 0x54, 0x77,
	0x3c, 0x84, 0x4e, 0xc6, 0x32, 0x99, 0x8b, 0x54, 0x97, 0xeb, 0x9f, 0x3f, 0xb1, 0xf2, 0xf6, 0x37,
	0x41, 0x2b, 0x9a, 0xbc, 0x05, 0x2f, 0x89, 0xd6, 0xda, 0x97, 0x08, 0x1a, 0xda, 0x4a, 0x60, 0xdf,
	0xfe, 0x63, 0x9c, 0x6e, 0x9f, 0x92, 0x53, 0xe8, 0x0a, 0xb3, 0xb2, 0x22, 0x70, 0x75, 0xda, 0x51,
	0x95, 0x56, 0x2f, 0x33, 0xad, 0x9f, 0x90, 0x77, 0xd0, 0x2b, 0xcd, 0x49, 0xe2, 0x9d, 0xce, 0x69,
	0xea, 0x9c, 0x67, 0xd5, 0x29, 0xef, 0x9c, 0x2b, 0xdd, 0x7f, 0x49, 0x4e, 0xc1, 0x5b, 0xe2, 0xc6,
	0xa6, 0xb5, 0x74, 0xda, 0x53, 0x9b, 0x56, 0xdd, 0x21, 0xdd, 0xbe, 0x88, 0xda, 0x9a, 0x7a, 0xfd,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x19, 0x99, 0xbe, 0x28, 0x05, 0x00, 0x00,
}
