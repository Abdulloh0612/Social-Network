// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/comment-service/comment.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId     string `protobuf:"bytes,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	ViewsCount string `protobuf:"bytes,5,opt,name=views_count,json=viewsCount,proto3" json:"views_count,omitempty"`
	CreatedAt  string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt  string `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_comment_service_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_protos_comment_service_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_protos_comment_service_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Comment) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetViewsCount() string {
	if x != nil {
		return x.ViewsCount
	}
	return ""
}

func (x *Comment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Comment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Comment) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type CreateComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId  string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateComment) Reset() {
	*x = CreateComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_comment_service_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateComment) ProtoMessage() {}

func (x *CreateComment) ProtoReflect() protoreflect.Message {
	mi := &file_protos_comment_service_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateComment.ProtoReflect.Descriptor instead.
func (*CreateComment) Descriptor() ([]byte, []int) {
	return file_protos_comment_service_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateComment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateComment) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *CreateComment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateComment) Reset() {
	*x = UpdateComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_comment_service_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateComment) ProtoMessage() {}

func (x *UpdateComment) ProtoReflect() protoreflect.Message {
	mi := &file_protos_comment_service_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateComment.ProtoReflect.Descriptor instead.
func (*UpdateComment) Descriptor() ([]byte, []int) {
	return file_protos_comment_service_comment_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateComment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateComment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DeleteCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCommentReq) Reset() {
	*x = DeleteCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_comment_service_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentReq) ProtoMessage() {}

func (x *DeleteCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_comment_service_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentReq.ProtoReflect.Descriptor instead.
func (*DeleteCommentReq) Descriptor() ([]byte, []int) {
	return file_protos_comment_service_comment_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCommentReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCommentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteCommentRes) Reset() {
	*x = DeleteCommentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_comment_service_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRes) ProtoMessage() {}

func (x *DeleteCommentRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_comment_service_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRes.ProtoReflect.Descriptor instead.
func (*DeleteCommentRes) Descriptor() ([]byte, []int) {
	return file_protos_comment_service_comment_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCommentRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetCommentReq) Reset() {
	*x = GetCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_comment_service_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentReq) ProtoMessage() {}

func (x *GetCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_comment_service_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentReq.ProtoReflect.Descriptor instead.
func (*GetCommentReq) Descriptor() ([]byte, []int) {
	return file_protos_comment_service_comment_proto_rawDescGZIP(), []int{5}
}

func (x *GetCommentReq) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *GetCommentReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetAllCommentsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit     int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Field     string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	Value     string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Sortby    string `protobuf:"bytes,5,opt,name=sortby,proto3" json:"sortby,omitempty"`
	StartedAt string `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt   string `protobuf:"bytes,7,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
}

func (x *GetAllCommentsReq) Reset() {
	*x = GetAllCommentsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_comment_service_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCommentsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCommentsReq) ProtoMessage() {}

func (x *GetAllCommentsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_comment_service_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCommentsReq.ProtoReflect.Descriptor instead.
func (*GetAllCommentsReq) Descriptor() ([]byte, []int) {
	return file_protos_comment_service_comment_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllCommentsReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllCommentsReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllCommentsReq) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *GetAllCommentsReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetAllCommentsReq) GetSortby() string {
	if x != nil {
		return x.Sortby
	}
	return ""
}

func (x *GetAllCommentsReq) GetStartedAt() string {
	if x != nil {
		return x.StartedAt
	}
	return ""
}

func (x *GetAllCommentsReq) GetEndedAt() string {
	if x != nil {
		return x.EndedAt
	}
	return ""
}

type GetAllCommentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Count    int64      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllCommentRes) Reset() {
	*x = GetAllCommentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_comment_service_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCommentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCommentRes) ProtoMessage() {}

func (x *GetAllCommentRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_comment_service_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCommentRes.ProtoReflect.Descriptor instead.
func (*GetAllCommentRes) Descriptor() ([]byte, []int) {
	return file_protos_comment_service_comment_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllCommentRes) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetAllCommentRes) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_protos_comment_service_comment_proto protoreflect.FileDescriptor

var file_protos_comment_service_comment_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0xe3, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x39, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x22, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xbb, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74,
	0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x62, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x56, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2c,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xaa, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3f,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x42,
	0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_comment_service_comment_proto_rawDescOnce sync.Once
	file_protos_comment_service_comment_proto_rawDescData = file_protos_comment_service_comment_proto_rawDesc
)

func file_protos_comment_service_comment_proto_rawDescGZIP() []byte {
	file_protos_comment_service_comment_proto_rawDescOnce.Do(func() {
		file_protos_comment_service_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_comment_service_comment_proto_rawDescData)
	})
	return file_protos_comment_service_comment_proto_rawDescData
}

var file_protos_comment_service_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_comment_service_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),           // 0: comment.Comment
	(*CreateComment)(nil),     // 1: comment.CreateComment
	(*UpdateComment)(nil),     // 2: comment.UpdateComment
	(*DeleteCommentReq)(nil),  // 3: comment.DeleteCommentReq
	(*DeleteCommentRes)(nil),  // 4: comment.DeleteCommentRes
	(*GetCommentReq)(nil),     // 5: comment.GetCommentReq
	(*GetAllCommentsReq)(nil), // 6: comment.GetAllCommentsReq
	(*GetAllCommentRes)(nil),  // 7: comment.GetAllCommentRes
}
var file_protos_comment_service_comment_proto_depIdxs = []int32{
	0, // 0: comment.GetAllCommentRes.comments:type_name -> comment.Comment
	1, // 1: comment.CommentService.Create:input_type -> comment.CreateComment
	2, // 2: comment.CommentService.Update:input_type -> comment.UpdateComment
	3, // 3: comment.CommentService.Delete:input_type -> comment.DeleteCommentReq
	5, // 4: comment.CommentService.Get:input_type -> comment.GetCommentReq
	6, // 5: comment.CommentService.GetAll:input_type -> comment.GetAllCommentsReq
	0, // 6: comment.CommentService.Create:output_type -> comment.Comment
	0, // 7: comment.CommentService.Update:output_type -> comment.Comment
	4, // 8: comment.CommentService.Delete:output_type -> comment.DeleteCommentRes
	0, // 9: comment.CommentService.Get:output_type -> comment.Comment
	7, // 10: comment.CommentService.GetAll:output_type -> comment.GetAllCommentRes
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_comment_service_comment_proto_init() }
func file_protos_comment_service_comment_proto_init() {
	if File_protos_comment_service_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_comment_service_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_comment_service_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateComment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_comment_service_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateComment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_comment_service_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_comment_service_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_comment_service_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_comment_service_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCommentsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_comment_service_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCommentRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_comment_service_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_comment_service_comment_proto_goTypes,
		DependencyIndexes: file_protos_comment_service_comment_proto_depIdxs,
		MessageInfos:      file_protos_comment_service_comment_proto_msgTypes,
	}.Build()
	File_protos_comment_service_comment_proto = out.File
	file_protos_comment_service_comment_proto_rawDesc = nil
	file_protos_comment_service_comment_proto_goTypes = nil
	file_protos_comment_service_comment_proto_depIdxs = nil
}
