// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2/session_entity_type.proto

package dialogflow // import "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The types of modifications for a session entity type.
type SessionEntityType_EntityOverrideMode int32

const (
	// Not specified. This value should be never used.
	SessionEntityType_ENTITY_OVERRIDE_MODE_UNSPECIFIED SessionEntityType_EntityOverrideMode = 0
	// The collection of session entities overrides the collection of entities
	// in the corresponding developer entity type.
	SessionEntityType_ENTITY_OVERRIDE_MODE_OVERRIDE SessionEntityType_EntityOverrideMode = 1
	// The collection of session entities extends the collection of entities in
	// the corresponding developer entity type.
	//
	// Note: Even in this override mode calls to `ListSessionEntityTypes`,
	// `GetSessionEntityType`, `CreateSessionEntityType` and
	// `UpdateSessionEntityType` only return the additional entities added in
	// this session entity type. If you want to get the supplemented list,
	// please call [EntityTypes.GetEntityType][google.cloud.dialogflow.v2.EntityTypes.GetEntityType] on the developer entity type
	// and merge.
	SessionEntityType_ENTITY_OVERRIDE_MODE_SUPPLEMENT SessionEntityType_EntityOverrideMode = 2
)

var SessionEntityType_EntityOverrideMode_name = map[int32]string{
	0: "ENTITY_OVERRIDE_MODE_UNSPECIFIED",
	1: "ENTITY_OVERRIDE_MODE_OVERRIDE",
	2: "ENTITY_OVERRIDE_MODE_SUPPLEMENT",
}
var SessionEntityType_EntityOverrideMode_value = map[string]int32{
	"ENTITY_OVERRIDE_MODE_UNSPECIFIED": 0,
	"ENTITY_OVERRIDE_MODE_OVERRIDE":    1,
	"ENTITY_OVERRIDE_MODE_SUPPLEMENT":  2,
}

func (x SessionEntityType_EntityOverrideMode) String() string {
	return proto.EnumName(SessionEntityType_EntityOverrideMode_name, int32(x))
}
func (SessionEntityType_EntityOverrideMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_session_entity_type_0b020acc3115ba2a, []int{0, 0}
}

// Represents a session entity type.
//
// Extends or replaces a developer entity type at the user session level (we
// refer to the entity types defined at the agent level as "developer entity
// types").
//
// Note: session entity types apply to all queries, regardless of the language.
type SessionEntityType struct {
	// Required. The unique identifier of this session entity type. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	//
	// `<Entity Type Display Name>` must be the display name of an existing entity
	// type in the same agent that will be overridden or supplemented.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Indicates whether the additional data should override or
	// supplement the developer entity type definition.
	EntityOverrideMode SessionEntityType_EntityOverrideMode `protobuf:"varint,2,opt,name=entity_override_mode,json=entityOverrideMode,proto3,enum=google.cloud.dialogflow.v2.SessionEntityType_EntityOverrideMode" json:"entity_override_mode,omitempty"`
	// Required. The collection of entities associated with this session entity
	// type.
	Entities             []*EntityType_Entity `protobuf:"bytes,3,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SessionEntityType) Reset()         { *m = SessionEntityType{} }
func (m *SessionEntityType) String() string { return proto.CompactTextString(m) }
func (*SessionEntityType) ProtoMessage()    {}
func (*SessionEntityType) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_entity_type_0b020acc3115ba2a, []int{0}
}
func (m *SessionEntityType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionEntityType.Unmarshal(m, b)
}
func (m *SessionEntityType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionEntityType.Marshal(b, m, deterministic)
}
func (dst *SessionEntityType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionEntityType.Merge(dst, src)
}
func (m *SessionEntityType) XXX_Size() int {
	return xxx_messageInfo_SessionEntityType.Size(m)
}
func (m *SessionEntityType) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionEntityType.DiscardUnknown(m)
}

var xxx_messageInfo_SessionEntityType proto.InternalMessageInfo

func (m *SessionEntityType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SessionEntityType) GetEntityOverrideMode() SessionEntityType_EntityOverrideMode {
	if m != nil {
		return m.EntityOverrideMode
	}
	return SessionEntityType_ENTITY_OVERRIDE_MODE_UNSPECIFIED
}

func (m *SessionEntityType) GetEntities() []*EntityType_Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

// The request message for [SessionEntityTypes.ListSessionEntityTypes][google.cloud.dialogflow.v2.SessionEntityTypes.ListSessionEntityTypes].
type ListSessionEntityTypesRequest struct {
	// Required. The session to list all session entity types from.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of items to return in a single page. By
	// default 100 and at most 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSessionEntityTypesRequest) Reset()         { *m = ListSessionEntityTypesRequest{} }
func (m *ListSessionEntityTypesRequest) String() string { return proto.CompactTextString(m) }
func (*ListSessionEntityTypesRequest) ProtoMessage()    {}
func (*ListSessionEntityTypesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_entity_type_0b020acc3115ba2a, []int{1}
}
func (m *ListSessionEntityTypesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSessionEntityTypesRequest.Unmarshal(m, b)
}
func (m *ListSessionEntityTypesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSessionEntityTypesRequest.Marshal(b, m, deterministic)
}
func (dst *ListSessionEntityTypesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSessionEntityTypesRequest.Merge(dst, src)
}
func (m *ListSessionEntityTypesRequest) XXX_Size() int {
	return xxx_messageInfo_ListSessionEntityTypesRequest.Size(m)
}
func (m *ListSessionEntityTypesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSessionEntityTypesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSessionEntityTypesRequest proto.InternalMessageInfo

func (m *ListSessionEntityTypesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListSessionEntityTypesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListSessionEntityTypesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [SessionEntityTypes.ListSessionEntityTypes][google.cloud.dialogflow.v2.SessionEntityTypes.ListSessionEntityTypes].
type ListSessionEntityTypesResponse struct {
	// The list of session entity types. There will be a maximum number of items
	// returned based on the page_size field in the request.
	SessionEntityTypes []*SessionEntityType `protobuf:"bytes,1,rep,name=session_entity_types,json=sessionEntityTypes,proto3" json:"session_entity_types,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSessionEntityTypesResponse) Reset()         { *m = ListSessionEntityTypesResponse{} }
func (m *ListSessionEntityTypesResponse) String() string { return proto.CompactTextString(m) }
func (*ListSessionEntityTypesResponse) ProtoMessage()    {}
func (*ListSessionEntityTypesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_entity_type_0b020acc3115ba2a, []int{2}
}
func (m *ListSessionEntityTypesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSessionEntityTypesResponse.Unmarshal(m, b)
}
func (m *ListSessionEntityTypesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSessionEntityTypesResponse.Marshal(b, m, deterministic)
}
func (dst *ListSessionEntityTypesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSessionEntityTypesResponse.Merge(dst, src)
}
func (m *ListSessionEntityTypesResponse) XXX_Size() int {
	return xxx_messageInfo_ListSessionEntityTypesResponse.Size(m)
}
func (m *ListSessionEntityTypesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSessionEntityTypesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSessionEntityTypesResponse proto.InternalMessageInfo

func (m *ListSessionEntityTypesResponse) GetSessionEntityTypes() []*SessionEntityType {
	if m != nil {
		return m.SessionEntityTypes
	}
	return nil
}

func (m *ListSessionEntityTypesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [SessionEntityTypes.GetSessionEntityType][google.cloud.dialogflow.v2.SessionEntityTypes.GetSessionEntityType].
type GetSessionEntityTypeRequest struct {
	// Required. The name of the session entity type. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionEntityTypeRequest) Reset()         { *m = GetSessionEntityTypeRequest{} }
func (m *GetSessionEntityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionEntityTypeRequest) ProtoMessage()    {}
func (*GetSessionEntityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_entity_type_0b020acc3115ba2a, []int{3}
}
func (m *GetSessionEntityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionEntityTypeRequest.Unmarshal(m, b)
}
func (m *GetSessionEntityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionEntityTypeRequest.Marshal(b, m, deterministic)
}
func (dst *GetSessionEntityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionEntityTypeRequest.Merge(dst, src)
}
func (m *GetSessionEntityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionEntityTypeRequest.Size(m)
}
func (m *GetSessionEntityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionEntityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionEntityTypeRequest proto.InternalMessageInfo

func (m *GetSessionEntityTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [SessionEntityTypes.CreateSessionEntityType][google.cloud.dialogflow.v2.SessionEntityTypes.CreateSessionEntityType].
type CreateSessionEntityTypeRequest struct {
	// Required. The session to create a session entity type for.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The session entity type to create.
	SessionEntityType    *SessionEntityType `protobuf:"bytes,2,opt,name=session_entity_type,json=sessionEntityType,proto3" json:"session_entity_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateSessionEntityTypeRequest) Reset()         { *m = CreateSessionEntityTypeRequest{} }
func (m *CreateSessionEntityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSessionEntityTypeRequest) ProtoMessage()    {}
func (*CreateSessionEntityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_entity_type_0b020acc3115ba2a, []int{4}
}
func (m *CreateSessionEntityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionEntityTypeRequest.Unmarshal(m, b)
}
func (m *CreateSessionEntityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionEntityTypeRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSessionEntityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionEntityTypeRequest.Merge(dst, src)
}
func (m *CreateSessionEntityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSessionEntityTypeRequest.Size(m)
}
func (m *CreateSessionEntityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionEntityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionEntityTypeRequest proto.InternalMessageInfo

func (m *CreateSessionEntityTypeRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateSessionEntityTypeRequest) GetSessionEntityType() *SessionEntityType {
	if m != nil {
		return m.SessionEntityType
	}
	return nil
}

// The request message for [SessionEntityTypes.UpdateSessionEntityType][google.cloud.dialogflow.v2.SessionEntityTypes.UpdateSessionEntityType].
type UpdateSessionEntityTypeRequest struct {
	// Required. The entity type to update. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	SessionEntityType *SessionEntityType `protobuf:"bytes,1,opt,name=session_entity_type,json=sessionEntityType,proto3" json:"session_entity_type,omitempty"`
	// Optional. The mask to control which fields get updated.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateSessionEntityTypeRequest) Reset()         { *m = UpdateSessionEntityTypeRequest{} }
func (m *UpdateSessionEntityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSessionEntityTypeRequest) ProtoMessage()    {}
func (*UpdateSessionEntityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_entity_type_0b020acc3115ba2a, []int{5}
}
func (m *UpdateSessionEntityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSessionEntityTypeRequest.Unmarshal(m, b)
}
func (m *UpdateSessionEntityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSessionEntityTypeRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateSessionEntityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSessionEntityTypeRequest.Merge(dst, src)
}
func (m *UpdateSessionEntityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSessionEntityTypeRequest.Size(m)
}
func (m *UpdateSessionEntityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSessionEntityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSessionEntityTypeRequest proto.InternalMessageInfo

func (m *UpdateSessionEntityTypeRequest) GetSessionEntityType() *SessionEntityType {
	if m != nil {
		return m.SessionEntityType
	}
	return nil
}

func (m *UpdateSessionEntityTypeRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The request message for [SessionEntityTypes.DeleteSessionEntityType][google.cloud.dialogflow.v2.SessionEntityTypes.DeleteSessionEntityType].
type DeleteSessionEntityTypeRequest struct {
	// Required. The name of the entity type to delete. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSessionEntityTypeRequest) Reset()         { *m = DeleteSessionEntityTypeRequest{} }
func (m *DeleteSessionEntityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSessionEntityTypeRequest) ProtoMessage()    {}
func (*DeleteSessionEntityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_entity_type_0b020acc3115ba2a, []int{6}
}
func (m *DeleteSessionEntityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSessionEntityTypeRequest.Unmarshal(m, b)
}
func (m *DeleteSessionEntityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSessionEntityTypeRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteSessionEntityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSessionEntityTypeRequest.Merge(dst, src)
}
func (m *DeleteSessionEntityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSessionEntityTypeRequest.Size(m)
}
func (m *DeleteSessionEntityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSessionEntityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSessionEntityTypeRequest proto.InternalMessageInfo

func (m *DeleteSessionEntityTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionEntityType)(nil), "google.cloud.dialogflow.v2.SessionEntityType")
	proto.RegisterType((*ListSessionEntityTypesRequest)(nil), "google.cloud.dialogflow.v2.ListSessionEntityTypesRequest")
	proto.RegisterType((*ListSessionEntityTypesResponse)(nil), "google.cloud.dialogflow.v2.ListSessionEntityTypesResponse")
	proto.RegisterType((*GetSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2.GetSessionEntityTypeRequest")
	proto.RegisterType((*CreateSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2.CreateSessionEntityTypeRequest")
	proto.RegisterType((*UpdateSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2.UpdateSessionEntityTypeRequest")
	proto.RegisterType((*DeleteSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2.DeleteSessionEntityTypeRequest")
	proto.RegisterEnum("google.cloud.dialogflow.v2.SessionEntityType_EntityOverrideMode", SessionEntityType_EntityOverrideMode_name, SessionEntityType_EntityOverrideMode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionEntityTypesClient is the client API for SessionEntityTypes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionEntityTypesClient interface {
	// Returns the list of all session entity types in the specified session.
	ListSessionEntityTypes(ctx context.Context, in *ListSessionEntityTypesRequest, opts ...grpc.CallOption) (*ListSessionEntityTypesResponse, error)
	// Retrieves the specified session entity type.
	GetSessionEntityType(ctx context.Context, in *GetSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error)
	// Creates a session entity type.
	//
	// If the specified session entity type already exists, overrides the session
	// entity type.
	CreateSessionEntityType(ctx context.Context, in *CreateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error)
	// Updates the specified session entity type.
	UpdateSessionEntityType(ctx context.Context, in *UpdateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error)
	// Deletes the specified session entity type.
	DeleteSessionEntityType(ctx context.Context, in *DeleteSessionEntityTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sessionEntityTypesClient struct {
	cc *grpc.ClientConn
}

func NewSessionEntityTypesClient(cc *grpc.ClientConn) SessionEntityTypesClient {
	return &sessionEntityTypesClient{cc}
}

func (c *sessionEntityTypesClient) ListSessionEntityTypes(ctx context.Context, in *ListSessionEntityTypesRequest, opts ...grpc.CallOption) (*ListSessionEntityTypesResponse, error) {
	out := new(ListSessionEntityTypesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/ListSessionEntityTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) GetSessionEntityType(ctx context.Context, in *GetSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error) {
	out := new(SessionEntityType)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/GetSessionEntityType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) CreateSessionEntityType(ctx context.Context, in *CreateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error) {
	out := new(SessionEntityType)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/CreateSessionEntityType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) UpdateSessionEntityType(ctx context.Context, in *UpdateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error) {
	out := new(SessionEntityType)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/UpdateSessionEntityType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) DeleteSessionEntityType(ctx context.Context, in *DeleteSessionEntityTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/DeleteSessionEntityType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionEntityTypesServer is the server API for SessionEntityTypes service.
type SessionEntityTypesServer interface {
	// Returns the list of all session entity types in the specified session.
	ListSessionEntityTypes(context.Context, *ListSessionEntityTypesRequest) (*ListSessionEntityTypesResponse, error)
	// Retrieves the specified session entity type.
	GetSessionEntityType(context.Context, *GetSessionEntityTypeRequest) (*SessionEntityType, error)
	// Creates a session entity type.
	//
	// If the specified session entity type already exists, overrides the session
	// entity type.
	CreateSessionEntityType(context.Context, *CreateSessionEntityTypeRequest) (*SessionEntityType, error)
	// Updates the specified session entity type.
	UpdateSessionEntityType(context.Context, *UpdateSessionEntityTypeRequest) (*SessionEntityType, error)
	// Deletes the specified session entity type.
	DeleteSessionEntityType(context.Context, *DeleteSessionEntityTypeRequest) (*empty.Empty, error)
}

func RegisterSessionEntityTypesServer(s *grpc.Server, srv SessionEntityTypesServer) {
	s.RegisterService(&_SessionEntityTypes_serviceDesc, srv)
}

func _SessionEntityTypes_ListSessionEntityTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSessionEntityTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).ListSessionEntityTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/ListSessionEntityTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).ListSessionEntityTypes(ctx, req.(*ListSessionEntityTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_GetSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).GetSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/GetSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).GetSessionEntityType(ctx, req.(*GetSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_CreateSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).CreateSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/CreateSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).CreateSessionEntityType(ctx, req.(*CreateSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_UpdateSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).UpdateSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/UpdateSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).UpdateSessionEntityType(ctx, req.(*UpdateSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_DeleteSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).DeleteSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/DeleteSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).DeleteSessionEntityType(ctx, req.(*DeleteSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionEntityTypes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.v2.SessionEntityTypes",
	HandlerType: (*SessionEntityTypesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSessionEntityTypes",
			Handler:    _SessionEntityTypes_ListSessionEntityTypes_Handler,
		},
		{
			MethodName: "GetSessionEntityType",
			Handler:    _SessionEntityTypes_GetSessionEntityType_Handler,
		},
		{
			MethodName: "CreateSessionEntityType",
			Handler:    _SessionEntityTypes_CreateSessionEntityType_Handler,
		},
		{
			MethodName: "UpdateSessionEntityType",
			Handler:    _SessionEntityTypes_UpdateSessionEntityType_Handler,
		},
		{
			MethodName: "DeleteSessionEntityType",
			Handler:    _SessionEntityTypes_DeleteSessionEntityType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/v2/session_entity_type.proto",
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2/session_entity_type.proto", fileDescriptor_session_entity_type_0b020acc3115ba2a)
}

var fileDescriptor_session_entity_type_0b020acc3115ba2a = []byte{
	// 808 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcd, 0x4e, 0xdb, 0x58,
	0x14, 0xc7, 0xe7, 0x9a, 0x19, 0x04, 0x17, 0xcd, 0x0c, 0xdc, 0x41, 0x21, 0x63, 0x86, 0x4c, 0xc6,
	0x33, 0x1a, 0xa1, 0xa8, 0xb5, 0xd5, 0x14, 0xb5, 0x2a, 0xb4, 0x52, 0x05, 0x31, 0x34, 0x2a, 0x81,
	0xd4, 0x09, 0x48, 0xad, 0x54, 0x59, 0x26, 0x39, 0x58, 0x2e, 0x89, 0xaf, 0xeb, 0xeb, 0xd0, 0x86,
	0x8a, 0x0d, 0xaf, 0xd0, 0x45, 0xd5, 0x6d, 0xa5, 0x2e, 0xda, 0xbe, 0x02, 0xaf, 0xd0, 0x55, 0x17,
	0x7d, 0x81, 0xee, 0x78, 0x81, 0x2e, 0x2b, 0x5f, 0x3b, 0x80, 0xb0, 0x7d, 0xab, 0xa0, 0xee, 0x7c,
	0x3f, 0xce, 0x39, 0xbf, 0xff, 0xb9, 0xe7, 0x1c, 0x19, 0x2f, 0xd8, 0x94, 0xda, 0x1d, 0xd0, 0x5a,
	0x1d, 0xda, 0x6b, 0x6b, 0x6d, 0xc7, 0xea, 0x50, 0x7b, 0xb7, 0x43, 0x9f, 0x69, 0xfb, 0x65, 0x8d,
	0x01, 0x63, 0x0e, 0x75, 0x4d, 0x70, 0x03, 0x27, 0xe8, 0x9b, 0x41, 0xdf, 0x03, 0xd5, 0xf3, 0x69,
	0x40, 0x89, 0x1c, 0x59, 0xa9, 0xdc, 0x4a, 0x3d, 0xb3, 0x52, 0xf7, 0xcb, 0xf2, 0x5f, 0xb1, 0x47,
	0xcb, 0x73, 0x34, 0xcb, 0x75, 0x69, 0x60, 0x05, 0x0e, 0x75, 0x59, 0x64, 0x29, 0xff, 0x79, 0xee,
	0xd4, 0x07, 0x46, 0x7b, 0x7e, 0x2b, 0x76, 0x2a, 0x5f, 0x11, 0xa0, 0x24, 0x10, 0xe4, 0xd9, 0xf8,
	0x36, 0x5f, 0xed, 0xf4, 0x76, 0x35, 0xe8, 0x7a, 0x41, 0x3f, 0x3e, 0x2c, 0x5e, 0x3c, 0xdc, 0x75,
	0xa0, 0xd3, 0x36, 0xbb, 0x16, 0xdb, 0x8b, 0x6e, 0x28, 0x27, 0x12, 0x9e, 0x6a, 0x44, 0xfa, 0x74,
	0xee, 0xbb, 0xd9, 0xf7, 0x80, 0x10, 0xfc, 0xb3, 0x6b, 0x75, 0x21, 0x8f, 0x8a, 0x68, 0x7e, 0xdc,
	0xe0, 0xdf, 0xc4, 0xc7, 0xd3, 0x71, 0x74, 0xba, 0x0f, 0xbe, 0xef, 0xb4, 0xc1, 0xec, 0xd2, 0x36,
	0xe4, 0xa5, 0x22, 0x9a, 0xff, 0xad, 0x7c, 0x57, 0xcd, 0x4e, 0x85, 0x9a, 0x08, 0xa0, 0x46, 0x9f,
	0x9b, 0xb1, 0xa3, 0x1a, 0x6d, 0x83, 0x41, 0x20, 0xb1, 0x47, 0xaa, 0x78, 0x8c, 0xef, 0x3a, 0xc0,
	0xf2, 0x23, 0xc5, 0x91, 0xf9, 0x89, 0xf2, 0x55, 0x51, 0x9c, 0x44, 0x00, 0xe3, 0xd4, 0x5c, 0x39,
	0x42, 0x98, 0x24, 0xa3, 0x92, 0xff, 0x70, 0x51, 0xdf, 0x68, 0x56, 0x9b, 0x0f, 0xcd, 0xcd, 0x6d,
	0xdd, 0x30, 0xaa, 0x15, 0xdd, 0xac, 0x6d, 0x56, 0x74, 0x73, 0x6b, 0xa3, 0x51, 0xd7, 0x57, 0xaa,
	0xab, 0x55, 0xbd, 0x32, 0xf9, 0x13, 0xf9, 0x07, 0xcf, 0xa5, 0xde, 0x1a, 0xac, 0x26, 0x11, 0xf9,
	0x17, 0xff, 0x9d, 0x7a, 0xa5, 0xb1, 0x55, 0xaf, 0xaf, 0xeb, 0x35, 0x7d, 0xa3, 0x39, 0x29, 0x29,
	0x0c, 0xcf, 0xad, 0x3b, 0x2c, 0x48, 0xe4, 0x83, 0x19, 0xf0, 0xb4, 0x07, 0x2c, 0x20, 0x39, 0x3c,
	0xea, 0x59, 0x3e, 0xb8, 0x41, 0x9c, 0xfa, 0x78, 0x45, 0x66, 0xf1, 0xb8, 0x67, 0xd9, 0x60, 0x32,
	0xe7, 0x20, 0xca, 0xf8, 0x2f, 0xc6, 0x58, 0xb8, 0xd1, 0x70, 0x0e, 0x80, 0xcc, 0x61, 0xcc, 0x0f,
	0x03, 0xba, 0x07, 0x6e, 0x7e, 0x84, 0x1b, 0xf2, 0xeb, 0xcd, 0x70, 0x43, 0x79, 0x87, 0x70, 0x21,
	0x2b, 0x2a, 0xf3, 0xa8, 0xcb, 0x80, 0x98, 0x78, 0x3a, 0xa5, 0xc8, 0x59, 0x1e, 0x7d, 0x3f, 0xe7,
	0x09, 0xaf, 0x06, 0x61, 0x89, 0x40, 0xe4, 0x7f, 0xfc, 0xbb, 0x0b, 0xcf, 0x03, 0xf3, 0x1c, 0xa7,
	0xc4, 0x39, 0x7f, 0x0d, 0xb7, 0xeb, 0xa7, 0xac, 0xd7, 0xf0, 0xec, 0x1a, 0x24, 0x49, 0x07, 0xe9,
	0x49, 0xa9, 0x4b, 0xe5, 0x15, 0xc2, 0x85, 0x15, 0x1f, 0xac, 0x00, 0x32, 0xcd, 0xb2, 0xb2, 0xfa,
	0x18, 0xff, 0x91, 0x22, 0x9b, 0x93, 0x0d, 0xad, 0x7a, 0x2a, 0xa1, 0x5a, 0x39, 0x46, 0xb8, 0xb0,
	0xe5, 0xb5, 0x45, 0x64, 0x19, 0x04, 0xe8, 0xc7, 0x10, 0x90, 0x25, 0x3c, 0xd1, 0xe3, 0x00, 0xbc,
	0xe5, 0x63, 0x61, 0xf2, 0xc0, 0xed, 0x60, 0x2a, 0xa8, 0xab, 0xe1, 0x54, 0xa8, 0x59, 0x6c, 0xcf,
	0xc0, 0xd1, 0xf5, 0xf0, 0x5b, 0x59, 0xc0, 0x85, 0x0a, 0x74, 0x40, 0x40, 0x9f, 0xf2, 0x1c, 0xe5,
	0xd7, 0x63, 0x98, 0x24, 0x2b, 0x8d, 0x7c, 0x44, 0x38, 0x97, 0x5e, 0x84, 0xe4, 0x96, 0x48, 0xa6,
	0xb0, 0x5d, 0xe4, 0xc5, 0xcb, 0x98, 0x46, 0x35, 0xaf, 0xdc, 0x3e, 0xfa, 0xf4, 0xe5, 0xa5, 0x74,
	0x83, 0x2c, 0x84, 0x73, 0xf5, 0x45, 0x54, 0x11, 0x77, 0x3c, 0x9f, 0x3e, 0x81, 0x56, 0xc0, 0xb4,
	0x92, 0x66, 0xd9, 0xe0, 0x06, 0x83, 0xd9, 0xcf, 0xb4, 0xd2, 0x61, 0x3c, 0x7c, 0x23, 0xe8, 0x63,
	0x84, 0xa7, 0xd3, 0x2a, 0x95, 0xdc, 0x14, 0x21, 0x09, 0x6a, 0x5b, 0x1e, 0xee, 0xb5, 0x2f, 0xe0,
	0x87, 0xa9, 0x17, 0xc1, 0x9f, 0x67, 0xd7, 0x4a, 0x87, 0xe4, 0x33, 0xc2, 0x33, 0x19, 0x4d, 0x43,
	0x84, 0x49, 0x15, 0x77, 0xda, 0xb0, 0x22, 0x1e, 0x70, 0x11, 0xf7, 0x95, 0x4b, 0xbd, 0xc1, 0x62,
	0x5a, 0xeb, 0x90, 0x13, 0x84, 0x67, 0x32, 0x9a, 0x4e, 0xac, 0x4c, 0xdc, 0xa9, 0xc3, 0x2a, 0x03,
	0xae, 0xcc, 0x2c, 0xdf, 0xe3, 0xca, 0xd2, 0xfe, 0x20, 0x86, 0x7c, 0xb2, 0x74, 0xb5, 0x1f, 0x10,
	0x9e, 0xc9, 0x68, 0x52, 0xb1, 0x5a, 0x71, 0x67, 0xcb, 0xb9, 0xc4, 0x8c, 0xd0, 0xc3, 0xdf, 0x8a,
	0x41, 0xd5, 0x95, 0x2e, 0x55, 0x75, 0xcb, 0x6f, 0x11, 0x2e, 0xb4, 0x68, 0x57, 0xc0, 0xb5, 0x9c,
	0x4b, 0x20, 0xd5, 0x43, 0x82, 0x3a, 0x7a, 0x54, 0x89, 0xad, 0x6c, 0xda, 0xb1, 0x5c, 0x5b, 0xa5,
	0xbe, 0xad, 0xd9, 0xe0, 0x72, 0x3e, 0x2d, 0x3a, 0xb2, 0x3c, 0x87, 0xa5, 0xfd, 0x35, 0x2d, 0x9d,
	0xad, 0xbe, 0x22, 0xf4, 0x46, 0x92, 0x2a, 0xab, 0xef, 0x25, 0x79, 0x2d, 0x72, 0xb7, 0xc2, 0x21,
	0x2a, 0x67, 0x10, 0xdb, 0xe5, 0x9d, 0x51, 0xee, 0xf5, 0xfa, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0xde, 0xc9, 0xc2, 0x15, 0x0a, 0x00, 0x00,
}