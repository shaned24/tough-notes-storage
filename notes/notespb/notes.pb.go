// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notes/notespb/notes.proto

package notespb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Note struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Contents             string   `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
	AuthorId             string   `protobuf:"bytes,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_09d7226a476f5dbc, []int{0}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Note) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Note) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func (m *Note) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func init() {
	proto.RegisterType((*Note)(nil), "storage.Note")
}

func init() { proto.RegisterFile("notes/notespb/notes.proto", fileDescriptor_09d7226a476f5dbc) }

var fileDescriptor_09d7226a476f5dbc = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcb, 0x2f, 0x49,
	0x2d, 0xd6, 0x07, 0x93, 0x05, 0x49, 0x10, 0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0xbd,
	0xb8, 0x24, 0xbf, 0x28, 0x31, 0x3d, 0x55, 0x29, 0x95, 0x8b, 0xc5, 0x2f, 0xbf, 0x24, 0x55, 0x88,
	0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48,
	0x84, 0x8b, 0xb5, 0x24, 0xb3, 0x24, 0x27, 0x55, 0x82, 0x09, 0x2c, 0x04, 0xe1, 0x08, 0x49, 0x71,
	0x71, 0x24, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x14, 0x4b, 0x30, 0x83, 0x25, 0xe0, 0x7c, 0x21,
	0x69, 0x2e, 0xce, 0xc4, 0xd2, 0x92, 0x8c, 0xfc, 0xa2, 0xf8, 0xcc, 0x14, 0x09, 0x16, 0x88, 0x24,
	0x44, 0xc0, 0x33, 0xc5, 0x88, 0x97, 0x8b, 0x1b, 0x64, 0x4d, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72,
	0xaa, 0x13, 0x67, 0x14, 0x3b, 0xd4, 0x55, 0x49, 0x6c, 0x60, 0x07, 0x19, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xd5, 0xa7, 0x40, 0x46, 0xad, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NoteServiceClient is the client API for NoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoteServiceClient interface {
}

type noteServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoteServiceClient(cc *grpc.ClientConn) NoteServiceClient {
	return &noteServiceClient{cc}
}

// NoteServiceServer is the server API for NoteService service.
type NoteServiceServer interface {
}

func RegisterNoteServiceServer(s *grpc.Server, srv NoteServiceServer) {
	s.RegisterService(&_NoteService_serviceDesc, srv)
}

var _NoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "notes/notespb/notes.proto",
}
