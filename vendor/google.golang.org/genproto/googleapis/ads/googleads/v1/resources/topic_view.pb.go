// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/topic_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A topic view.
type TopicView struct {
	// The resource name of the topic view.
	// Topic view resource names have the form:
	//
	// `customers/{customer_id}/topicViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicView) Reset()         { *m = TopicView{} }
func (m *TopicView) String() string { return proto.CompactTextString(m) }
func (*TopicView) ProtoMessage()    {}
func (*TopicView) Descriptor() ([]byte, []int) {
	return fileDescriptor_topic_view_8b901d02d0a1a21d, []int{0}
}
func (m *TopicView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicView.Unmarshal(m, b)
}
func (m *TopicView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicView.Marshal(b, m, deterministic)
}
func (dst *TopicView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicView.Merge(dst, src)
}
func (m *TopicView) XXX_Size() int {
	return xxx_messageInfo_TopicView.Size(m)
}
func (m *TopicView) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicView.DiscardUnknown(m)
}

var xxx_messageInfo_TopicView proto.InternalMessageInfo

func (m *TopicView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*TopicView)(nil), "google.ads.googleads.v1.resources.TopicView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/topic_view.proto", fileDescriptor_topic_view_8b901d02d0a1a21d)
}

var fileDescriptor_topic_view_8b901d02d0a1a21d = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0xb9, 0x13, 0x84, 0x1e, 0xea, 0xd0, 0x49, 0xc4, 0xc1, 0x2a, 0x05, 0xa7, 0x17, 0xa3,
	0x5b, 0x9c, 0xd2, 0xa5, 0xe0, 0x20, 0xa5, 0xc8, 0x0d, 0x72, 0x50, 0xe2, 0x25, 0x84, 0x40, 0x2f,
	0xef, 0xb8, 0x77, 0x5e, 0xbf, 0x8f, 0xa3, 0x1f, 0xc5, 0x8f, 0xe2, 0x57, 0x70, 0x91, 0x6b, 0x4c,
	0x46, 0xbb, 0xfd, 0x49, 0x7e, 0xff, 0xdf, 0x7b, 0x49, 0x71, 0x6f, 0x11, 0xed, 0xd6, 0x30, 0xa5,
	0x89, 0x85, 0x38, 0xa6, 0x81, 0xb3, 0xce, 0x10, 0xbe, 0x77, 0xb5, 0x21, 0xd6, 0x63, 0xeb, 0xea,
	0xcd, 0xe0, 0xcc, 0x0e, 0xda, 0x0e, 0x7b, 0x9c, 0xce, 0x02, 0x08, 0x4a, 0x13, 0xa4, 0x0e, 0x0c,
	0x1c, 0x52, 0xe7, 0xe2, 0x32, 0x6a, 0x5b, 0xc7, 0x94, 0xf7, 0xd8, 0xab, 0xde, 0xa1, 0xa7, 0x20,
	0xb8, 0xbe, 0x2b, 0x26, 0x2f, 0xa3, 0xb4, 0x74, 0x66, 0x37, 0xbd, 0x29, 0x4e, 0x63, 0x6f, 0xe3,
	0x55, 0x63, 0xce, 0xb3, 0xab, 0xec, 0x76, 0xb2, 0x3e, 0x89, 0x87, 0xcf, 0xaa, 0x31, 0x8b, 0x9f,
	0xac, 0x98, 0xd7, 0xd8, 0xc0, 0xc1, 0xc9, 0x8b, 0xb3, 0x64, 0x5e, 0x8d, 0xb3, 0x56, 0xd9, 0xeb,
	0xd3, 0x5f, 0xc9, 0xe2, 0x56, 0x79, 0x0b, 0xd8, 0x59, 0x66, 0x8d, 0xdf, 0x6f, 0x12, 0x9f, 0xdc,
	0x3a, 0xfa, 0xe7, 0x07, 0x1e, 0x53, 0xfa, 0xc8, 0x8f, 0x96, 0x52, 0x7e, 0xe6, 0xb3, 0x65, 0x50,
	0x4a, 0x4d, 0x10, 0xe2, 0x98, 0x4a, 0x0e, 0xeb, 0x48, 0x7e, 0x45, 0xa6, 0x92, 0x9a, 0xaa, 0xc4,
	0x54, 0x25, 0xaf, 0x12, 0xf3, 0x9d, 0xcf, 0xc3, 0x85, 0x10, 0x52, 0x93, 0x10, 0x89, 0x12, 0xa2,
	0xe4, 0x42, 0x24, 0xee, 0xed, 0x78, 0xbf, 0xec, 0xc3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0xd4, 0x08, 0xb0, 0xad, 0x01, 0x00, 0x00,
}
