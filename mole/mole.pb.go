// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mole.proto

package mole

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ClientScore int32

const (
	ClientScore_MOLE_BONKED ClientScore = 0
	ClientScore_MOLE_MISSED ClientScore = 1
	ClientScore_EMPTY_HIT   ClientScore = 2
)

var ClientScore_name = map[int32]string{
	0: "MOLE_BONKED",
	1: "MOLE_MISSED",
	2: "EMPTY_HIT",
}

var ClientScore_value = map[string]int32{
	"MOLE_BONKED": 0,
	"MOLE_MISSED": 1,
	"EMPTY_HIT":   2,
}

func (x ClientScore) String() string {
	return proto.EnumName(ClientScore_name, int32(x))
}

func (ClientScore) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_362b8b7844cc62c4, []int{0}
}

// Screen dimension in the fourth quadrant
// X: +ve & Y: -ve
type ScreenDimension struct {
	Height               int32    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScreenDimension) Reset()         { *m = ScreenDimension{} }
func (m *ScreenDimension) String() string { return proto.CompactTextString(m) }
func (*ScreenDimension) ProtoMessage()    {}
func (*ScreenDimension) Descriptor() ([]byte, []int) {
	return fileDescriptor_362b8b7844cc62c4, []int{0}
}

func (m *ScreenDimension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScreenDimension.Unmarshal(m, b)
}
func (m *ScreenDimension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScreenDimension.Marshal(b, m, deterministic)
}
func (m *ScreenDimension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScreenDimension.Merge(m, src)
}
func (m *ScreenDimension) XXX_Size() int {
	return xxx_messageInfo_ScreenDimension.Size(m)
}
func (m *ScreenDimension) XXX_DiscardUnknown() {
	xxx_messageInfo_ScreenDimension.DiscardUnknown(m)
}

var xxx_messageInfo_ScreenDimension proto.InternalMessageInfo

func (m *ScreenDimension) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ScreenDimension) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

type MoleFromClient struct {
	Score                ClientScore          `protobuf:"varint,1,opt,name=score,proto3,enum=mole.ClientScore" json:"score,omitempty"`
	DisplayedOnScreen    *timestamp.Timestamp `protobuf:"bytes,2,opt,name=displayedOnScreen,proto3" json:"displayedOnScreen,omitempty"`
	ActionedByUser       *timestamp.Timestamp `protobuf:"bytes,3,opt,name=actionedByUser,proto3" json:"actionedByUser,omitempty"`
	Screen               *ScreenDimension     `protobuf:"bytes,4,opt,name=screen,proto3" json:"screen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MoleFromClient) Reset()         { *m = MoleFromClient{} }
func (m *MoleFromClient) String() string { return proto.CompactTextString(m) }
func (*MoleFromClient) ProtoMessage()    {}
func (*MoleFromClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_362b8b7844cc62c4, []int{1}
}

func (m *MoleFromClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoleFromClient.Unmarshal(m, b)
}
func (m *MoleFromClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoleFromClient.Marshal(b, m, deterministic)
}
func (m *MoleFromClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoleFromClient.Merge(m, src)
}
func (m *MoleFromClient) XXX_Size() int {
	return xxx_messageInfo_MoleFromClient.Size(m)
}
func (m *MoleFromClient) XXX_DiscardUnknown() {
	xxx_messageInfo_MoleFromClient.DiscardUnknown(m)
}

var xxx_messageInfo_MoleFromClient proto.InternalMessageInfo

func (m *MoleFromClient) GetScore() ClientScore {
	if m != nil {
		return m.Score
	}
	return ClientScore_MOLE_BONKED
}

func (m *MoleFromClient) GetDisplayedOnScreen() *timestamp.Timestamp {
	if m != nil {
		return m.DisplayedOnScreen
	}
	return nil
}

func (m *MoleFromClient) GetActionedByUser() *timestamp.Timestamp {
	if m != nil {
		return m.ActionedByUser
	}
	return nil
}

func (m *MoleFromClient) GetScreen() *ScreenDimension {
	if m != nil {
		return m.Screen
	}
	return nil
}

type MoleFromServer struct {
	Img  string `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
	XPos int32  `protobuf:"varint,2,opt,name=xPos,proto3" json:"xPos,omitempty"`
	YPos int32  `protobuf:"varint,3,opt,name=yPos,proto3" json:"yPos,omitempty"`
	//Peep time in milli seconds
	PopUp                int64                `protobuf:"varint,4,opt,name=popUp,proto3" json:"popUp,omitempty"`
	ServerSideCreation   *timestamp.Timestamp `protobuf:"bytes,5,opt,name=serverSideCreation,proto3" json:"serverSideCreation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MoleFromServer) Reset()         { *m = MoleFromServer{} }
func (m *MoleFromServer) String() string { return proto.CompactTextString(m) }
func (*MoleFromServer) ProtoMessage()    {}
func (*MoleFromServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_362b8b7844cc62c4, []int{2}
}

func (m *MoleFromServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoleFromServer.Unmarshal(m, b)
}
func (m *MoleFromServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoleFromServer.Marshal(b, m, deterministic)
}
func (m *MoleFromServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoleFromServer.Merge(m, src)
}
func (m *MoleFromServer) XXX_Size() int {
	return xxx_messageInfo_MoleFromServer.Size(m)
}
func (m *MoleFromServer) XXX_DiscardUnknown() {
	xxx_messageInfo_MoleFromServer.DiscardUnknown(m)
}

var xxx_messageInfo_MoleFromServer proto.InternalMessageInfo

func (m *MoleFromServer) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *MoleFromServer) GetXPos() int32 {
	if m != nil {
		return m.XPos
	}
	return 0
}

func (m *MoleFromServer) GetYPos() int32 {
	if m != nil {
		return m.YPos
	}
	return 0
}

func (m *MoleFromServer) GetPopUp() int64 {
	if m != nil {
		return m.PopUp
	}
	return 0
}

func (m *MoleFromServer) GetServerSideCreation() *timestamp.Timestamp {
	if m != nil {
		return m.ServerSideCreation
	}
	return nil
}

func init() {
	proto.RegisterEnum("mole.ClientScore", ClientScore_name, ClientScore_value)
	proto.RegisterType((*ScreenDimension)(nil), "mole.ScreenDimension")
	proto.RegisterType((*MoleFromClient)(nil), "mole.MoleFromClient")
	proto.RegisterType((*MoleFromServer)(nil), "mole.MoleFromServer")
}

func init() { proto.RegisterFile("mole.proto", fileDescriptor_362b8b7844cc62c4) }

var fileDescriptor_362b8b7844cc62c4 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xb5, 0x1d, 0xd1, 0x89, 0x48, 0xdd, 0x51, 0x41, 0x56, 0x2e, 0xa0, 0x5c, 0xa8, 0x90,
	0x70, 0x51, 0xb8, 0x71, 0x00, 0xc9, 0x8d, 0xa1, 0x05, 0x4c, 0x2a, 0x3b, 0x39, 0x70, 0x8a, 0x9c,
	0x78, 0x70, 0x56, 0xb2, 0xbd, 0xd6, 0xee, 0xf2, 0x91, 0xff, 0xc4, 0xef, 0xe3, 0x8c, 0x76, 0x37,
	0x41, 0x21, 0x20, 0xe5, 0xf6, 0xe6, 0xcd, 0xbe, 0xd9, 0x37, 0x6f, 0x00, 0x1a, 0x5e, 0x53, 0xd4,
	0x09, 0xae, 0x38, 0x7a, 0x1a, 0x0f, 0x1f, 0x57, 0x9c, 0x57, 0x35, 0x5d, 0x19, 0x6e, 0xf9, 0xf5,
	0xcb, 0x95, 0x62, 0x0d, 0x49, 0x55, 0x34, 0x9d, 0x7d, 0x36, 0x7a, 0x03, 0x67, 0xf9, 0x4a, 0x10,
	0xb5, 0x13, 0xd6, 0x50, 0x2b, 0x19, 0x6f, 0xf1, 0x11, 0xf4, 0xd6, 0xc4, 0xaa, 0xb5, 0x0a, 0x9d,
	0x27, 0xce, 0xa5, 0x9f, 0x6d, 0x2b, 0xbc, 0x00, 0xff, 0x3b, 0x2b, 0xd5, 0x3a, 0x3c, 0x31, 0xb4,
	0x2d, 0x46, 0xbf, 0x1c, 0x18, 0xa4, 0xbc, 0xa6, 0xb7, 0x82, 0x37, 0xd7, 0x35, 0xa3, 0x56, 0xe1,
	0x53, 0xf0, 0xe5, 0x8a, 0x0b, 0x32, 0xfa, 0xc1, 0xf8, 0x3c, 0x32, 0xb6, 0x6c, 0x33, 0xd7, 0x8d,
	0xcc, 0xf6, 0xf1, 0x06, 0xce, 0x4b, 0x26, 0xbb, 0xba, 0xd8, 0x50, 0x39, 0x6d, 0xad, 0x0f, 0x33,
	0xbd, 0x3f, 0x1e, 0x46, 0xd6, 0x79, 0xb4, 0x73, 0x1e, 0xcd, 0x76, 0xce, 0xb3, 0x7f, 0x45, 0x18,
	0xc3, 0xa0, 0x58, 0x29, 0xc6, 0x5b, 0x2a, 0xe3, 0xcd, 0x5c, 0x92, 0x08, 0xdd, 0xa3, 0x63, 0x0e,
	0x14, 0xf8, 0x1c, 0x7a, 0xd2, 0x5a, 0xf0, 0x8c, 0xf6, 0xa1, 0xf5, 0x7d, 0x10, 0x4f, 0xb6, 0x7d,
	0x34, 0xfa, 0xb9, 0xb7, 0x78, 0x4e, 0xe2, 0x1b, 0x09, 0x0c, 0xc0, 0x65, 0x4d, 0x65, 0xd6, 0x3e,
	0xcd, 0x34, 0x44, 0x04, 0xef, 0xc7, 0x1d, 0x97, 0xdb, 0xc8, 0x0c, 0xd6, 0xdc, 0x46, 0x73, 0xae,
	0xe5, 0x34, 0xd6, 0xd9, 0x76, 0xbc, 0x9b, 0x77, 0xe6, 0x6b, 0x37, 0xb3, 0x05, 0xbe, 0x07, 0x94,
	0x66, 0x72, 0xce, 0x4a, 0xba, 0x16, 0x54, 0x68, 0xbf, 0xa1, 0x7f, 0x74, 0xb3, 0xff, 0xa8, 0x9e,
	0xbd, 0x86, 0xfe, 0xde, 0x05, 0xf0, 0x0c, 0xfa, 0xe9, 0xf4, 0x63, 0xb2, 0x88, 0xa7, 0x9f, 0x3e,
	0x24, 0x93, 0xe0, 0xde, 0x1f, 0x22, 0xbd, 0xcd, 0xf3, 0x64, 0x12, 0x38, 0xf8, 0x00, 0x4e, 0x93,
	0xf4, 0x6e, 0xf6, 0x79, 0x71, 0x73, 0x3b, 0x0b, 0x4e, 0xc6, 0x31, 0x78, 0x7a, 0x5b, 0x7c, 0x05,
	0xf7, 0xdf, 0x91, 0xd2, 0x50, 0xe2, 0x85, 0x4d, 0xe8, 0xef, 0xf3, 0x0f, 0x0f, 0x58, 0x9b, 0xcd,
	0xa5, 0xf3, 0xc2, 0x59, 0xf6, 0x8c, 0xd7, 0x97, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x20,
	0x3e, 0xa2, 0xa8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MoleClient is the client API for Mole service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoleClient interface {
	GetMoles(ctx context.Context, opts ...grpc.CallOption) (Mole_GetMolesClient, error)
}

type moleClient struct {
	cc grpc.ClientConnInterface
}

func NewMoleClient(cc grpc.ClientConnInterface) MoleClient {
	return &moleClient{cc}
}

func (c *moleClient) GetMoles(ctx context.Context, opts ...grpc.CallOption) (Mole_GetMolesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Mole_serviceDesc.Streams[0], "/mole.Mole/GetMoles", opts...)
	if err != nil {
		return nil, err
	}
	x := &moleGetMolesClient{stream}
	return x, nil
}

type Mole_GetMolesClient interface {
	Send(*MoleFromClient) error
	Recv() (*MoleFromServer, error)
	grpc.ClientStream
}

type moleGetMolesClient struct {
	grpc.ClientStream
}

func (x *moleGetMolesClient) Send(m *MoleFromClient) error {
	return x.ClientStream.SendMsg(m)
}

func (x *moleGetMolesClient) Recv() (*MoleFromServer, error) {
	m := new(MoleFromServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MoleServer is the server API for Mole service.
type MoleServer interface {
	GetMoles(Mole_GetMolesServer) error
}

// UnimplementedMoleServer can be embedded to have forward compatible implementations.
type UnimplementedMoleServer struct {
}

func (*UnimplementedMoleServer) GetMoles(srv Mole_GetMolesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMoles not implemented")
}

func RegisterMoleServer(s *grpc.Server, srv MoleServer) {
	s.RegisterService(&_Mole_serviceDesc, srv)
}

func _Mole_GetMoles_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MoleServer).GetMoles(&moleGetMolesServer{stream})
}

type Mole_GetMolesServer interface {
	Send(*MoleFromServer) error
	Recv() (*MoleFromClient, error)
	grpc.ServerStream
}

type moleGetMolesServer struct {
	grpc.ServerStream
}

func (x *moleGetMolesServer) Send(m *MoleFromServer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *moleGetMolesServer) Recv() (*MoleFromClient, error) {
	m := new(MoleFromClient)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Mole_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mole.Mole",
	HandlerType: (*MoleServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMoles",
			Handler:       _Mole_GetMoles_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mole.proto",
}
