// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.0
// source: common/model/blog_field.proto

package model

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Article struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Summary       string                 `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	Image         string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	ContentMd     string                 `protobuf:"bytes,5,opt,name=content_md,json=contentMd,proto3" json:"content_md,omitempty"`
	ContentHtml   string                 `protobuf:"bytes,6,opt,name=content_html,json=contentHtml,proto3" json:"content_html,omitempty"`
	Views         int64                  `protobuf:"varint,7,opt,name=views,proto3" json:"views,omitempty"`
	Tags          []*Tag                 `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Article) Reset() {
	*x = Article{}
	mi := &file_common_model_blog_field_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_blog_field_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_common_model_blog_field_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Article) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Article) GetContentMd() string {
	if x != nil {
		return x.ContentMd
	}
	return ""
}

func (x *Article) GetContentHtml() string {
	if x != nil {
		return x.ContentHtml
	}
	return ""
}

func (x *Article) GetViews() int64 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *Article) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Article) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Article) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tag) Reset() {
	*x = Tag{}
	mi := &file_common_model_blog_field_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_blog_field_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_common_model_blog_field_proto_rawDescGZIP(), []int{1}
}

func (x *Tag) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_common_model_blog_field_proto protoreflect.FileDescriptor

const file_common_model_blog_field_proto_rawDesc = "" +
	"\n" +
	"\x1dcommon/model/blog_field.proto\x12\fcommon.model\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x17validate/validate.proto\"\xdd\x02\n" +
	"\aArticle\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\asummary\x18\x03 \x01(\tR\asummary\x12\x14\n" +
	"\x05image\x18\x04 \x01(\tR\x05image\x12\x1d\n" +
	"\n" +
	"content_md\x18\x05 \x01(\tR\tcontentMd\x12!\n" +
	"\fcontent_html\x18\x06 \x01(\tR\vcontentHtml\x12\x14\n" +
	"\x05views\x18\a \x01(\x03R\x05views\x12%\n" +
	"\x04tags\x18\b \x03(\v2\x11.common.model.TagR\x04tags\x129\n" +
	"\n" +
	"created_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\")\n" +
	"\x03Tag\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04nameB8Z6github.com/SoLikeWind/XuanXiang/api/common/model;modelb\x06proto3"

var (
	file_common_model_blog_field_proto_rawDescOnce sync.Once
	file_common_model_blog_field_proto_rawDescData []byte
)

func file_common_model_blog_field_proto_rawDescGZIP() []byte {
	file_common_model_blog_field_proto_rawDescOnce.Do(func() {
		file_common_model_blog_field_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_model_blog_field_proto_rawDesc), len(file_common_model_blog_field_proto_rawDesc)))
	})
	return file_common_model_blog_field_proto_rawDescData
}

var file_common_model_blog_field_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_model_blog_field_proto_goTypes = []any{
	(*Article)(nil),               // 0: common.model.Article
	(*Tag)(nil),                   // 1: common.model.Tag
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_common_model_blog_field_proto_depIdxs = []int32{
	1, // 0: common.model.Article.tags:type_name -> common.model.Tag
	2, // 1: common.model.Article.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: common.model.Article.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_model_blog_field_proto_init() }
func file_common_model_blog_field_proto_init() {
	if File_common_model_blog_field_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_model_blog_field_proto_rawDesc), len(file_common_model_blog_field_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_model_blog_field_proto_goTypes,
		DependencyIndexes: file_common_model_blog_field_proto_depIdxs,
		MessageInfos:      file_common_model_blog_field_proto_msgTypes,
	}.Build()
	File_common_model_blog_field_proto = out.File
	file_common_model_blog_field_proto_goTypes = nil
	file_common_model_blog_field_proto_depIdxs = nil
}
