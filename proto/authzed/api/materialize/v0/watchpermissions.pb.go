// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: authzed/api/materialize/v0/watchpermissions.proto

package v0

import (
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
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

type PermissionChange_Permissionship int32

const (
	PermissionChange_PERMISSIONSHIP_UNSPECIFIED            PermissionChange_Permissionship = 0
	PermissionChange_PERMISSIONSHIP_NO_PERMISSION          PermissionChange_Permissionship = 1
	PermissionChange_PERMISSIONSHIP_HAS_PERMISSION         PermissionChange_Permissionship = 2
	PermissionChange_PERMISSIONSHIP_CONDITIONAL_PERMISSION PermissionChange_Permissionship = 3
)

// Enum value maps for PermissionChange_Permissionship.
var (
	PermissionChange_Permissionship_name = map[int32]string{
		0: "PERMISSIONSHIP_UNSPECIFIED",
		1: "PERMISSIONSHIP_NO_PERMISSION",
		2: "PERMISSIONSHIP_HAS_PERMISSION",
		3: "PERMISSIONSHIP_CONDITIONAL_PERMISSION",
	}
	PermissionChange_Permissionship_value = map[string]int32{
		"PERMISSIONSHIP_UNSPECIFIED":            0,
		"PERMISSIONSHIP_NO_PERMISSION":          1,
		"PERMISSIONSHIP_HAS_PERMISSION":         2,
		"PERMISSIONSHIP_CONDITIONAL_PERMISSION": 3,
	}
)

func (x PermissionChange_Permissionship) Enum() *PermissionChange_Permissionship {
	p := new(PermissionChange_Permissionship)
	*p = x
	return p
}

func (x PermissionChange_Permissionship) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionChange_Permissionship) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_materialize_v0_watchpermissions_proto_enumTypes[0].Descriptor()
}

func (PermissionChange_Permissionship) Type() protoreflect.EnumType {
	return &file_authzed_api_materialize_v0_watchpermissions_proto_enumTypes[0]
}

func (x PermissionChange_Permissionship) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionChange_Permissionship.Descriptor instead.
func (PermissionChange_Permissionship) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_materialize_v0_watchpermissions_proto_rawDescGZIP(), []int{3, 0}
}

type WatchPermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// permissions is a list of permissions to watch for changes. At least one permission must be specified, and it must
	// be a subset or equal to the permissions that were enabled for the service.
	Permissions []*WatchedPermission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// optional_starting_after is the revision token to start watching from. If not provided, the stream
	// will start from the current revision at the moment of the request.
	OptionalStartingAfter *v1.ZedToken `protobuf:"bytes,2,opt,name=optional_starting_after,json=optionalStartingAfter,proto3" json:"optional_starting_after,omitempty"`
}

func (x *WatchPermissionsRequest) Reset() {
	*x = WatchPermissionsRequest{}
	mi := &file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchPermissionsRequest) ProtoMessage() {}

func (x *WatchPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchPermissionsRequest.ProtoReflect.Descriptor instead.
func (*WatchPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_materialize_v0_watchpermissions_proto_rawDescGZIP(), []int{0}
}

func (x *WatchPermissionsRequest) GetPermissions() []*WatchedPermission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *WatchPermissionsRequest) GetOptionalStartingAfter() *v1.ZedToken {
	if x != nil {
		return x.OptionalStartingAfter
	}
	return nil
}

type WatchedPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource_type is the type of the resource to watch for changes.
	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// permission is the permission to watch for changes.
	Permission string `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	// subject_type is the type of the subject to watch for changes.
	SubjectType string `protobuf:"bytes,3,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"`
	// optional_subject_relation is the relation on the subject to watch for changes.
	OptionalSubjectRelation string `protobuf:"bytes,4,opt,name=optional_subject_relation,json=optionalSubjectRelation,proto3" json:"optional_subject_relation,omitempty"`
}

func (x *WatchedPermission) Reset() {
	*x = WatchedPermission{}
	mi := &file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchedPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchedPermission) ProtoMessage() {}

func (x *WatchedPermission) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchedPermission.ProtoReflect.Descriptor instead.
func (*WatchedPermission) Descriptor() ([]byte, []int) {
	return file_authzed_api_materialize_v0_watchpermissions_proto_rawDescGZIP(), []int{1}
}

func (x *WatchedPermission) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *WatchedPermission) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *WatchedPermission) GetSubjectType() string {
	if x != nil {
		return x.SubjectType
	}
	return ""
}

func (x *WatchedPermission) GetOptionalSubjectRelation() string {
	if x != nil {
		return x.OptionalSubjectRelation
	}
	return ""
}

type WatchPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*WatchPermissionsResponse_Change
	//	*WatchPermissionsResponse_CompletedRevision
	Response isWatchPermissionsResponse_Response `protobuf_oneof:"response"`
}

func (x *WatchPermissionsResponse) Reset() {
	*x = WatchPermissionsResponse{}
	mi := &file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchPermissionsResponse) ProtoMessage() {}

func (x *WatchPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchPermissionsResponse.ProtoReflect.Descriptor instead.
func (*WatchPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_materialize_v0_watchpermissions_proto_rawDescGZIP(), []int{2}
}

func (m *WatchPermissionsResponse) GetResponse() isWatchPermissionsResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *WatchPermissionsResponse) GetChange() *PermissionChange {
	if x, ok := x.GetResponse().(*WatchPermissionsResponse_Change); ok {
		return x.Change
	}
	return nil
}

func (x *WatchPermissionsResponse) GetCompletedRevision() *v1.ZedToken {
	if x, ok := x.GetResponse().(*WatchPermissionsResponse_CompletedRevision); ok {
		return x.CompletedRevision
	}
	return nil
}

type isWatchPermissionsResponse_Response interface {
	isWatchPermissionsResponse_Response()
}

type WatchPermissionsResponse_Change struct {
	// change is the computed permission delta that has occurred as result of a mutation in origin SpiceDB.
	// The consumer should apply this change to the current state of the computed permissions in their target system.
	// Once an event arrives with completed_revision instead, the consumer shall consider there are not more changes
	// originating from that revision.
	//
	// The consumer should keep track of the revision in order to resume streaming in the event of consumer restarts.
	Change *PermissionChange `protobuf:"bytes,1,opt,name=change,proto3,oneof"`
}

type WatchPermissionsResponse_CompletedRevision struct {
	// completed_revision is the revision token that indicates all changes originating from a revision have been
	// streamed and thus the revision should be considered completed. It may also be
	// received without accompanying set of changes, indicating that a mutation in the origin SpiceDB cluster did
	// not yield any effective changes in the computed permissions
	CompletedRevision *v1.ZedToken `protobuf:"bytes,2,opt,name=completed_revision,json=completedRevision,proto3,oneof"`
}

func (*WatchPermissionsResponse_Change) isWatchPermissionsResponse_Response() {}

func (*WatchPermissionsResponse_CompletedRevision) isWatchPermissionsResponse_Response() {}

type PermissionChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// revision represents the revision at which the change occurred.
	Revision *v1.ZedToken `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// resource is the resource that the permission change is related to.
	Resource *v1.ObjectReference `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// permission is the permission that has changed.
	Permission string `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
	// subject is the subject that the permission change is related to.
	Subject *v1.SubjectReference `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	// permissionship is the new permissionship of the subject over the resource after the change.
	Permissionship PermissionChange_Permissionship `protobuf:"varint,5,opt,name=permissionship,proto3,enum=authzed.api.materialize.v0.PermissionChange_Permissionship" json:"permissionship,omitempty"`
}

func (x *PermissionChange) Reset() {
	*x = PermissionChange{}
	mi := &file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionChange) ProtoMessage() {}

func (x *PermissionChange) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionChange.ProtoReflect.Descriptor instead.
func (*PermissionChange) Descriptor() ([]byte, []int) {
	return file_authzed_api_materialize_v0_watchpermissions_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionChange) GetRevision() *v1.ZedToken {
	if x != nil {
		return x.Revision
	}
	return nil
}

func (x *PermissionChange) GetResource() *v1.ObjectReference {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *PermissionChange) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *PermissionChange) GetSubject() *v1.SubjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *PermissionChange) GetPermissionship() PermissionChange_Permissionship {
	if x != nil {
		return x.Permissionship
	}
	return PermissionChange_PERMISSIONSHIP_UNSPECIFIED
}

var File_authzed_api_materialize_v0_watchpermissions_proto protoreflect.FileDescriptor

var file_authzed_api_materialize_v0_watchpermissions_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x30, 0x2f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x30, 0x1a,
	0x19, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x17, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x50, 0x0a, 0x17, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x5a, 0x65, 0x64, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0xb7, 0x01, 0x0a, 0x11, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xb9, 0x01, 0x0a, 0x18, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x5a, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00,
	0x52, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xe9, 0x03, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x5a, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x63, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x1e, 0x0a, 0x1a, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x4e, 0x4f,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a,
	0x1d, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f,
	0x48, 0x41, 0x53, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x29, 0x0a, 0x25, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48,
	0x49, 0x50, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x32, 0x9d, 0x01, 0x0a, 0x17,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x10, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x34, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x62, 0x0a, 0x1e, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x30, 0x50, 0x01, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x30, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authzed_api_materialize_v0_watchpermissions_proto_rawDescOnce sync.Once
	file_authzed_api_materialize_v0_watchpermissions_proto_rawDescData = file_authzed_api_materialize_v0_watchpermissions_proto_rawDesc
)

func file_authzed_api_materialize_v0_watchpermissions_proto_rawDescGZIP() []byte {
	file_authzed_api_materialize_v0_watchpermissions_proto_rawDescOnce.Do(func() {
		file_authzed_api_materialize_v0_watchpermissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_materialize_v0_watchpermissions_proto_rawDescData)
	})
	return file_authzed_api_materialize_v0_watchpermissions_proto_rawDescData
}

var file_authzed_api_materialize_v0_watchpermissions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_authzed_api_materialize_v0_watchpermissions_proto_goTypes = []any{
	(PermissionChange_Permissionship)(0), // 0: authzed.api.materialize.v0.PermissionChange.Permissionship
	(*WatchPermissionsRequest)(nil),      // 1: authzed.api.materialize.v0.WatchPermissionsRequest
	(*WatchedPermission)(nil),            // 2: authzed.api.materialize.v0.WatchedPermission
	(*WatchPermissionsResponse)(nil),     // 3: authzed.api.materialize.v0.WatchPermissionsResponse
	(*PermissionChange)(nil),             // 4: authzed.api.materialize.v0.PermissionChange
	(*v1.ZedToken)(nil),                  // 5: authzed.api.v1.ZedToken
	(*v1.ObjectReference)(nil),           // 6: authzed.api.v1.ObjectReference
	(*v1.SubjectReference)(nil),          // 7: authzed.api.v1.SubjectReference
}
var file_authzed_api_materialize_v0_watchpermissions_proto_depIdxs = []int32{
	2, // 0: authzed.api.materialize.v0.WatchPermissionsRequest.permissions:type_name -> authzed.api.materialize.v0.WatchedPermission
	5, // 1: authzed.api.materialize.v0.WatchPermissionsRequest.optional_starting_after:type_name -> authzed.api.v1.ZedToken
	4, // 2: authzed.api.materialize.v0.WatchPermissionsResponse.change:type_name -> authzed.api.materialize.v0.PermissionChange
	5, // 3: authzed.api.materialize.v0.WatchPermissionsResponse.completed_revision:type_name -> authzed.api.v1.ZedToken
	5, // 4: authzed.api.materialize.v0.PermissionChange.revision:type_name -> authzed.api.v1.ZedToken
	6, // 5: authzed.api.materialize.v0.PermissionChange.resource:type_name -> authzed.api.v1.ObjectReference
	7, // 6: authzed.api.materialize.v0.PermissionChange.subject:type_name -> authzed.api.v1.SubjectReference
	0, // 7: authzed.api.materialize.v0.PermissionChange.permissionship:type_name -> authzed.api.materialize.v0.PermissionChange.Permissionship
	1, // 8: authzed.api.materialize.v0.WatchPermissionsService.WatchPermissions:input_type -> authzed.api.materialize.v0.WatchPermissionsRequest
	3, // 9: authzed.api.materialize.v0.WatchPermissionsService.WatchPermissions:output_type -> authzed.api.materialize.v0.WatchPermissionsResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_authzed_api_materialize_v0_watchpermissions_proto_init() }
func file_authzed_api_materialize_v0_watchpermissions_proto_init() {
	if File_authzed_api_materialize_v0_watchpermissions_proto != nil {
		return
	}
	file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes[2].OneofWrappers = []any{
		(*WatchPermissionsResponse_Change)(nil),
		(*WatchPermissionsResponse_CompletedRevision)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authzed_api_materialize_v0_watchpermissions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authzed_api_materialize_v0_watchpermissions_proto_goTypes,
		DependencyIndexes: file_authzed_api_materialize_v0_watchpermissions_proto_depIdxs,
		EnumInfos:         file_authzed_api_materialize_v0_watchpermissions_proto_enumTypes,
		MessageInfos:      file_authzed_api_materialize_v0_watchpermissions_proto_msgTypes,
	}.Build()
	File_authzed_api_materialize_v0_watchpermissions_proto = out.File
	file_authzed_api_materialize_v0_watchpermissions_proto_rawDesc = nil
	file_authzed_api_materialize_v0_watchpermissions_proto_goTypes = nil
	file_authzed_api_materialize_v0_watchpermissions_proto_depIdxs = nil
}
