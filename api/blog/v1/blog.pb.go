// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.0
// source: blog/v1/blog.proto

package v1

import (
	model "github.com/SoLikeWind/XuanXiang/api/common/model"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type ListArticleReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Tag           string                 `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListArticleReq) Reset() {
	*x = ListArticleReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleReq) ProtoMessage() {}

func (x *ListArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleReq.ProtoReflect.Descriptor instead.
func (*ListArticleReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{0}
}

func (x *ListArticleReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListArticleReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListArticleReq) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type ListArticleReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Articles      []*model.Article       `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListArticleReply) Reset() {
	*x = ListArticleReply{}
	mi := &file_blog_v1_blog_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleReply) ProtoMessage() {}

func (x *ListArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleReply.ProtoReflect.Descriptor instead.
func (*ListArticleReply) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{1}
}

func (x *ListArticleReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListArticleReply) GetArticles() []*model.Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

type GetArticleReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetArticleReq) Reset() {
	*x = GetArticleReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleReq) ProtoMessage() {}

func (x *GetArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleReq.ProtoReflect.Descriptor instead.
func (*GetArticleReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{2}
}

func (x *GetArticleReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetArticleReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Article       *model.Article         `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetArticleReply) Reset() {
	*x = GetArticleReply{}
	mi := &file_blog_v1_blog_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleReply) ProtoMessage() {}

func (x *GetArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleReply.ProtoReflect.Descriptor instead.
func (*GetArticleReply) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticleReply) GetArticle() *model.Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type CreateArticleReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Summary       string                 `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Image         *string                `protobuf:"bytes,3,opt,name=image,proto3,oneof" json:"image,omitempty"`
	ContentMd     string                 `protobuf:"bytes,4,opt,name=content_md,json=contentMd,proto3" json:"content_md,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateArticleReq) Reset() {
	*x = CreateArticleReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleReq) ProtoMessage() {}

func (x *CreateArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleReq.ProtoReflect.Descriptor instead.
func (*CreateArticleReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{4}
}

func (x *CreateArticleReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateArticleReq) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *CreateArticleReq) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *CreateArticleReq) GetContentMd() string {
	if x != nil {
		return x.ContentMd
	}
	return ""
}

type CreateArticleReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Article       *model.Article         `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateArticleReply) Reset() {
	*x = CreateArticleReply{}
	mi := &file_blog_v1_blog_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleReply) ProtoMessage() {}

func (x *CreateArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleReply.ProtoReflect.Descriptor instead.
func (*CreateArticleReply) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{5}
}

func (x *CreateArticleReply) GetArticle() *model.Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type UpdateArticleReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         *string                `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Summary       *string                `protobuf:"bytes,3,opt,name=summary,proto3,oneof" json:"summary,omitempty"`
	Image         *string                `protobuf:"bytes,4,opt,name=image,proto3,oneof" json:"image,omitempty"`
	ContentMd     *string                `protobuf:"bytes,5,opt,name=content_md,json=contentMd,proto3,oneof" json:"content_md,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateArticleReq) Reset() {
	*x = UpdateArticleReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleReq) ProtoMessage() {}

func (x *UpdateArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleReq.ProtoReflect.Descriptor instead.
func (*UpdateArticleReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateArticleReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateArticleReq) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdateArticleReq) GetSummary() string {
	if x != nil && x.Summary != nil {
		return *x.Summary
	}
	return ""
}

func (x *UpdateArticleReq) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *UpdateArticleReq) GetContentMd() string {
	if x != nil && x.ContentMd != nil {
		return *x.ContentMd
	}
	return ""
}

type UpdateArticleReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Article       *model.Article         `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateArticleReply) Reset() {
	*x = UpdateArticleReply{}
	mi := &file_blog_v1_blog_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleReply) ProtoMessage() {}

func (x *UpdateArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleReply.ProtoReflect.Descriptor instead.
func (*UpdateArticleReply) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateArticleReply) GetArticle() *model.Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type DeleteArticleReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteArticleReq) Reset() {
	*x = DeleteArticleReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleReq) ProtoMessage() {}

func (x *DeleteArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleReq.ProtoReflect.Descriptor instead.
func (*DeleteArticleReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteArticleReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListTagReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTagReq) Reset() {
	*x = ListTagReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagReq) ProtoMessage() {}

func (x *ListTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagReq.ProtoReflect.Descriptor instead.
func (*ListTagReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{9}
}

func (x *ListTagReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTagReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTagReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListTagReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Tags          []*model.Tag           `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTagReply) Reset() {
	*x = ListTagReply{}
	mi := &file_blog_v1_blog_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagReply) ProtoMessage() {}

func (x *ListTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagReply.ProtoReflect.Descriptor instead.
func (*ListTagReply) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{10}
}

func (x *ListTagReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListTagReply) GetTags() []*model.Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type GetTagReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTagReq) Reset() {
	*x = GetTagReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagReq) ProtoMessage() {}

func (x *GetTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagReq.ProtoReflect.Descriptor instead.
func (*GetTagReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{11}
}

func (x *GetTagReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTagReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           *model.Tag             `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTagReply) Reset() {
	*x = GetTagReply{}
	mi := &file_blog_v1_blog_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagReply) ProtoMessage() {}

func (x *GetTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagReply.ProtoReflect.Descriptor instead.
func (*GetTagReply) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{12}
}

func (x *GetTagReply) GetTag() *model.Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type CreateTagReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTagReq) Reset() {
	*x = CreateTagReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagReq) ProtoMessage() {}

func (x *CreateTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagReq.ProtoReflect.Descriptor instead.
func (*CreateTagReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{13}
}

func (x *CreateTagReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateTagReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           *model.Tag             `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTagReply) Reset() {
	*x = CreateTagReply{}
	mi := &file_blog_v1_blog_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagReply) ProtoMessage() {}

func (x *CreateTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagReply.ProtoReflect.Descriptor instead.
func (*CreateTagReply) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{14}
}

func (x *CreateTagReply) GetTag() *model.Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type UpdateTagReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTagReq) Reset() {
	*x = UpdateTagReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTagReq) ProtoMessage() {}

func (x *UpdateTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTagReq.ProtoReflect.Descriptor instead.
func (*UpdateTagReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{15}
}

func (x *UpdateTagReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTagReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteTagReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTagReq) Reset() {
	*x = DeleteTagReq{}
	mi := &file_blog_v1_blog_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagReq) ProtoMessage() {}

func (x *DeleteTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagReq.ProtoReflect.Descriptor instead.
func (*DeleteTagReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{16}
}

func (x *DeleteTagReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_blog_v1_blog_proto protoreflect.FileDescriptor

const file_blog_v1_blog_proto_rawDesc = "" +
	"\n" +
	"\x12blog/v1/blog.proto\x12\ablog.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x17validate/validate.proto\x1a\x1dcommon/model/blog_field.proto\"S\n" +
	"\x0eListArticleReq\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x03R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x03R\bpageSize\x12\x10\n" +
	"\x03tag\x18\x03 \x01(\tR\x03tag\"[\n" +
	"\x10ListArticleReply\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x03R\x05total\x121\n" +
	"\barticles\x18\x02 \x03(\v2\x15.common.model.ArticleR\barticles\"(\n" +
	"\rGetArticleReq\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x02id\"B\n" +
	"\x0fGetArticleReply\x12/\n" +
	"\aarticle\x18\x01 \x01(\v2\x15.common.model.ArticleR\aarticle\"\x86\x01\n" +
	"\x10CreateArticleReq\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\asummary\x18\x02 \x01(\tR\asummary\x12\x19\n" +
	"\x05image\x18\x03 \x01(\tH\x00R\x05image\x88\x01\x01\x12\x1d\n" +
	"\n" +
	"content_md\x18\x04 \x01(\tR\tcontentMdB\b\n" +
	"\x06_image\"E\n" +
	"\x12CreateArticleReply\x12/\n" +
	"\aarticle\x18\x01 \x01(\v2\x15.common.model.ArticleR\aarticle\"\x82\x02\n" +
	"\x10UpdateArticleReq\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x02id\x12$\n" +
	"\x05title\x18\x02 \x01(\tB\t\xfaB\x06r\x04\x10\x01\x18xH\x00R\x05title\x88\x01\x01\x12)\n" +
	"\asummary\x18\x03 \x01(\tB\n" +
	"\xfaB\ar\x05\x10\x01\x18\xf0\x01H\x01R\asummary\x88\x01\x01\x12$\n" +
	"\x05image\x18\x04 \x01(\tB\t\xfaB\x06r\x04\x10\x01\x18xH\x02R\x05image\x88\x01\x01\x12/\n" +
	"\n" +
	"content_md\x18\x05 \x01(\tB\v\xfaB\br\x06\x10\x01\x18\xa0\x8d\x06H\x03R\tcontentMd\x88\x01\x01B\b\n" +
	"\x06_titleB\n" +
	"\n" +
	"\b_summaryB\b\n" +
	"\x06_imageB\r\n" +
	"\v_content_md\"E\n" +
	"\x12UpdateArticleReply\x12/\n" +
	"\aarticle\x18\x01 \x01(\v2\x15.common.model.ArticleR\aarticle\"+\n" +
	"\x10DeleteArticleReq\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x02id\"Q\n" +
	"\n" +
	"ListTagReq\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x03R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x03R\bpageSize\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\"K\n" +
	"\fListTagReply\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x03R\x05total\x12%\n" +
	"\x04tags\x18\x02 \x03(\v2\x11.common.model.TagR\x04tags\"\x1b\n" +
	"\tGetTagReq\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"2\n" +
	"\vGetTagReply\x12#\n" +
	"\x03tag\x18\x01 \x01(\v2\x11.common.model.TagR\x03tag\"\"\n" +
	"\fCreateTagReq\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"5\n" +
	"\x0eCreateTagReply\x12#\n" +
	"\x03tag\x18\x01 \x01(\v2\x11.common.model.TagR\x03tag\"2\n" +
	"\fUpdateTagReq\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"\x1e\n" +
	"\fDeleteTagReq\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\x94\a\n" +
	"\x04Blog\x12[\n" +
	"\vListArticle\x12\x17.blog.v1.ListArticleReq\x1a\x19.blog.v1.ListArticleReply\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/v1/blog/article\x12]\n" +
	"\n" +
	"GetArticle\x12\x16.blog.v1.GetArticleReq\x1a\x18.blog.v1.GetArticleReply\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/blog/article/{id}\x12d\n" +
	"\rCreateArticle\x12\x19.blog.v1.CreateArticleReq\x1a\x1b.blog.v1.CreateArticleReply\"\x1b\x82\xd3\xe4\x93\x02\x15:\x01*\"\x10/v1/blog/article\x12d\n" +
	"\rUpdateArticle\x12\x19.blog.v1.UpdateArticleReq\x1a\x16.google.protobuf.Empty\" \x82\xd3\xe4\x93\x02\x1a:\x01*\x1a\x15/v1/blog/article/{id}\x12a\n" +
	"\rDeleteArticle\x12\x19.blog.v1.DeleteArticleReq\x1a\x16.google.protobuf.Empty\"\x1d\x82\xd3\xe4\x93\x02\x17*\x15/v1/blog/article/{id}\x12K\n" +
	"\aListTag\x12\x13.blog.v1.ListTagReq\x1a\x15.blog.v1.ListTagReply\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\f/v1/blog/tag\x12M\n" +
	"\x06GetTag\x12\x12.blog.v1.GetTagReq\x1a\x14.blog.v1.GetTagReply\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/blog/tag/{id}\x12T\n" +
	"\tCreateTag\x12\x15.blog.v1.CreateTagReq\x1a\x17.blog.v1.CreateTagReply\"\x17\x82\xd3\xe4\x93\x02\x11:\x01*\"\f/v1/blog/tag\x12X\n" +
	"\tUpdateTag\x12\x15.blog.v1.UpdateTagReq\x1a\x16.google.protobuf.Empty\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\x1a\x11/v1/blog/tag/{id}\x12U\n" +
	"\tDeleteTag\x12\x15.blog.v1.DeleteTagReq\x1a\x16.google.protobuf.Empty\"\x19\x82\xd3\xe4\x93\x02\x13*\x11/v1/blog/tag/{id}B0Z.github.com/SoLikeWind/XuanXiang/api/blog/v1;v1b\x06proto3"

var (
	file_blog_v1_blog_proto_rawDescOnce sync.Once
	file_blog_v1_blog_proto_rawDescData []byte
)

func file_blog_v1_blog_proto_rawDescGZIP() []byte {
	file_blog_v1_blog_proto_rawDescOnce.Do(func() {
		file_blog_v1_blog_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_blog_v1_blog_proto_rawDesc), len(file_blog_v1_blog_proto_rawDesc)))
	})
	return file_blog_v1_blog_proto_rawDescData
}

var file_blog_v1_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_blog_v1_blog_proto_goTypes = []any{
	(*ListArticleReq)(nil),     // 0: blog.v1.ListArticleReq
	(*ListArticleReply)(nil),   // 1: blog.v1.ListArticleReply
	(*GetArticleReq)(nil),      // 2: blog.v1.GetArticleReq
	(*GetArticleReply)(nil),    // 3: blog.v1.GetArticleReply
	(*CreateArticleReq)(nil),   // 4: blog.v1.CreateArticleReq
	(*CreateArticleReply)(nil), // 5: blog.v1.CreateArticleReply
	(*UpdateArticleReq)(nil),   // 6: blog.v1.UpdateArticleReq
	(*UpdateArticleReply)(nil), // 7: blog.v1.UpdateArticleReply
	(*DeleteArticleReq)(nil),   // 8: blog.v1.DeleteArticleReq
	(*ListTagReq)(nil),         // 9: blog.v1.ListTagReq
	(*ListTagReply)(nil),       // 10: blog.v1.ListTagReply
	(*GetTagReq)(nil),          // 11: blog.v1.GetTagReq
	(*GetTagReply)(nil),        // 12: blog.v1.GetTagReply
	(*CreateTagReq)(nil),       // 13: blog.v1.CreateTagReq
	(*CreateTagReply)(nil),     // 14: blog.v1.CreateTagReply
	(*UpdateTagReq)(nil),       // 15: blog.v1.UpdateTagReq
	(*DeleteTagReq)(nil),       // 16: blog.v1.DeleteTagReq
	(*model.Article)(nil),      // 17: common.model.Article
	(*model.Tag)(nil),          // 18: common.model.Tag
	(*emptypb.Empty)(nil),      // 19: google.protobuf.Empty
}
var file_blog_v1_blog_proto_depIdxs = []int32{
	17, // 0: blog.v1.ListArticleReply.articles:type_name -> common.model.Article
	17, // 1: blog.v1.GetArticleReply.article:type_name -> common.model.Article
	17, // 2: blog.v1.CreateArticleReply.article:type_name -> common.model.Article
	17, // 3: blog.v1.UpdateArticleReply.article:type_name -> common.model.Article
	18, // 4: blog.v1.ListTagReply.tags:type_name -> common.model.Tag
	18, // 5: blog.v1.GetTagReply.tag:type_name -> common.model.Tag
	18, // 6: blog.v1.CreateTagReply.tag:type_name -> common.model.Tag
	0,  // 7: blog.v1.Blog.ListArticle:input_type -> blog.v1.ListArticleReq
	2,  // 8: blog.v1.Blog.GetArticle:input_type -> blog.v1.GetArticleReq
	4,  // 9: blog.v1.Blog.CreateArticle:input_type -> blog.v1.CreateArticleReq
	6,  // 10: blog.v1.Blog.UpdateArticle:input_type -> blog.v1.UpdateArticleReq
	8,  // 11: blog.v1.Blog.DeleteArticle:input_type -> blog.v1.DeleteArticleReq
	9,  // 12: blog.v1.Blog.ListTag:input_type -> blog.v1.ListTagReq
	11, // 13: blog.v1.Blog.GetTag:input_type -> blog.v1.GetTagReq
	13, // 14: blog.v1.Blog.CreateTag:input_type -> blog.v1.CreateTagReq
	15, // 15: blog.v1.Blog.UpdateTag:input_type -> blog.v1.UpdateTagReq
	16, // 16: blog.v1.Blog.DeleteTag:input_type -> blog.v1.DeleteTagReq
	1,  // 17: blog.v1.Blog.ListArticle:output_type -> blog.v1.ListArticleReply
	3,  // 18: blog.v1.Blog.GetArticle:output_type -> blog.v1.GetArticleReply
	5,  // 19: blog.v1.Blog.CreateArticle:output_type -> blog.v1.CreateArticleReply
	19, // 20: blog.v1.Blog.UpdateArticle:output_type -> google.protobuf.Empty
	19, // 21: blog.v1.Blog.DeleteArticle:output_type -> google.protobuf.Empty
	10, // 22: blog.v1.Blog.ListTag:output_type -> blog.v1.ListTagReply
	12, // 23: blog.v1.Blog.GetTag:output_type -> blog.v1.GetTagReply
	14, // 24: blog.v1.Blog.CreateTag:output_type -> blog.v1.CreateTagReply
	19, // 25: blog.v1.Blog.UpdateTag:output_type -> google.protobuf.Empty
	19, // 26: blog.v1.Blog.DeleteTag:output_type -> google.protobuf.Empty
	17, // [17:27] is the sub-list for method output_type
	7,  // [7:17] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_blog_v1_blog_proto_init() }
func file_blog_v1_blog_proto_init() {
	if File_blog_v1_blog_proto != nil {
		return
	}
	file_blog_v1_blog_proto_msgTypes[4].OneofWrappers = []any{}
	file_blog_v1_blog_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_blog_v1_blog_proto_rawDesc), len(file_blog_v1_blog_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_v1_blog_proto_goTypes,
		DependencyIndexes: file_blog_v1_blog_proto_depIdxs,
		MessageInfos:      file_blog_v1_blog_proto_msgTypes,
	}.Build()
	File_blog_v1_blog_proto = out.File
	file_blog_v1_blog_proto_goTypes = nil
	file_blog_v1_blog_proto_depIdxs = nil
}
