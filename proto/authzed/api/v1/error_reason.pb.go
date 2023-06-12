// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: authzed/api/v1/error_reason.proto

package v1

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

// Defines the supported values for `google.rpc.ErrorInfo.reason` for the
// `authzed.com` error domain.
type ErrorReason int32

const (
	// Do not use this default value.
	ErrorReason_ERROR_REASON_UNSPECIFIED ErrorReason = 0
	// The request gave a schema that could not be parsed.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_SCHEMA_PARSE_ERROR",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "start_line_number": "1",
	//	    "start_column_position": "19",
	//	    "end_line_number": "1",
	//	    "end_column_position": "19",
	//	    "source_code": "somedefinition",
	//	  }
	//	}
	//
	// The line numbers and column positions are 0-indexed and may not be present.
	ErrorReason_ERROR_REASON_SCHEMA_PARSE_ERROR ErrorReason = 1
	// The request contains a schema with a type error.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_SCHEMA_TYPE_ERROR",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "definition_name": "somedefinition",
	//	    ... additional keys based on the kind of type error ...
	//	  }
	//	}
	ErrorReason_ERROR_REASON_SCHEMA_TYPE_ERROR ErrorReason = 2
	// The request referenced an unknown object definition in the schema.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_UNKNOWN_DEFINITION",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "definition_name": "somedefinition"
	//	  }
	//	}
	ErrorReason_ERROR_REASON_UNKNOWN_DEFINITION ErrorReason = 3
	// The request referenced an unknown relation or permission under a definition in the schema.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_UNKNOWN_RELATION_OR_PERMISSION",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "definition_name": "somedefinition",
	//	    "relation_or_permission_name": "somepermission"
	//	  }
	//	}
	ErrorReason_ERROR_REASON_UNKNOWN_RELATION_OR_PERMISSION ErrorReason = 4
	// The WriteRelationships request contained more updates than the maximum configured.
	//
	// Example of an ErrorInfo:
	//
	//	{ "reason": "ERROR_REASON_TOO_MANY_UPDATES_IN_REQUEST",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "update_count": "525",
	//	    "maximum_updates_allowed": "500",
	//	  }
	//	}
	ErrorReason_ERROR_REASON_TOO_MANY_UPDATES_IN_REQUEST ErrorReason = 5
	// The request contained more preconditions than the maximum configured.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_TOO_MANY_PRECONDITIONS_IN_REQUEST",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "precondition_count": "525",
	//	    "maximum_preconditions_allowed": "500",
	//	  }
	//	}
	ErrorReason_ERROR_REASON_TOO_MANY_PRECONDITIONS_IN_REQUEST ErrorReason = 6
	// The request contained a precondition that failed.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_WRITE_OR_DELETE_PRECONDITION_FAILURE",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "precondition_resource_type": "document",
	//	    ... other fields for the filter ...
	//	    "precondition_operation": "MUST_EXIST",
	//	  }
	//	}
	ErrorReason_ERROR_REASON_WRITE_OR_DELETE_PRECONDITION_FAILURE ErrorReason = 7
	// A write or delete request was made to an instance that is deployed in read-only mode.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_SERVICE_READ_ONLY",
	//	  "domain": "authzed.com"
	//	}
	ErrorReason_ERROR_REASON_SERVICE_READ_ONLY ErrorReason = 8
	// The request referenced an unknown caveat in the schema.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_UNKNOWN_CAVEAT",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "caveat_name": "somecaveat"
	//	  }
	//	}
	ErrorReason_ERROR_REASON_UNKNOWN_CAVEAT ErrorReason = 9
	// The request tries to use a subject type that was not valid for a relation.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_INVALID_SUBJECT_TYPE",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "definition_name": "somedefinition",
	//	    "relation_name": "somerelation",
	//	    "subject_type": "user:*"
	//	  }
	//	}
	ErrorReason_ERROR_REASON_INVALID_SUBJECT_TYPE ErrorReason = 10
	// The request tries to specify a caveat parameter value with the wrong type.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_CAVEAT_PARAMETER_TYPE_ERROR",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "definition_name": "somedefinition",
	//	    "relation_name": "somerelation",
	//	    "caveat_name": "somecaveat",
	//	    "parameter_name": "someparameter",
	//	    "expected_type": "int",
	//	  }
	//	}
	ErrorReason_ERROR_REASON_CAVEAT_PARAMETER_TYPE_ERROR ErrorReason = 11
	// The request tries to perform two or more updates on the same relationship in the same WriteRelationships call.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_UPDATES_ON_SAME_RELATIONSHIP",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "definition_name": "somedefinition",
	//	    "relationship": "somerelationship",
	//	  }
	//	}
	ErrorReason_ERROR_REASON_UPDATES_ON_SAME_RELATIONSHIP ErrorReason = 12
	// The request tries to write a relationship on a permission instead of a relation.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_CANNOT_UPDATE_PERMISSION",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "definition_name": "somedefinition",
	//	    "permission_name": "somerelation",
	//	  }
	//	}
	ErrorReason_ERROR_REASON_CANNOT_UPDATE_PERMISSION ErrorReason = 13
	// The request failed to evaluate a caveat expression due to an error.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_CAVEAT_EVALUATION_ERROR",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "caveat_name": "somecaveat",
	//	  }
	//	}
	ErrorReason_ERROR_REASON_CAVEAT_EVALUATION_ERROR ErrorReason = 14
	// The request failed because the provided cursor was invalid in some way.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_INVALID_CURSOR",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	     ... additional keys based on the kind of cursor error ...
	//	  }
	//	}
	ErrorReason_ERROR_REASON_INVALID_CURSOR ErrorReason = 15
	// The request failed because there are too many matching relationships to be
	// deleted within a single transactional deletion call. To avoid, set
	// `optional_allow_partial_deletions` to true on the DeleteRelationships call.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_TOO_MANY_RELATIONSHIPS_FOR_TRANSACTIONAL_DELETE",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	     ... fields for the filter ...
	//	  }
	//	}
	ErrorReason_ERROR_REASON_TOO_MANY_RELATIONSHIPS_FOR_TRANSACTIONAL_DELETE ErrorReason = 16
	// The request failed because the client attempted to write a relationship
	// with a context that exceeded the configured server limit.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_MAX_RELATIONSHIP_CONTEXT_SIZE",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "relationship":     "relationship_exceeding_the_limit",
	//	    "max_allowed_size": "server_max_allowed_context_size",
	//	    "context_size":     "actual_relationship_context_size" ,
	//	  }
	//	}
	ErrorReason_ERROR_REASON_MAX_RELATIONSHIP_CONTEXT_SIZE ErrorReason = 17
	// The request failed because a relationship marked to be CREATEd
	// was already present within the datastore.
	//
	// Example of an ErrorInfo:
	//
	//	{
	//	  "reason": "ERROR_REASON_ATTEMPT_TO_RECREATE_RELATIONSHIP",
	//	  "domain": "authzed.com",
	//	  "metadata": {
	//	    "relationship":          "relationship_that_already_existed",
	//	    "resource_type":         "resource type",
	//	    "resource_object_id":    "resource object id",
	//	    ... additional decomposed relationship fields ...
	//	  }
	//	}
	ErrorReason_ERROR_REASON_ATTEMPT_TO_RECREATE_RELATIONSHIP ErrorReason = 18
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:  "ERROR_REASON_UNSPECIFIED",
		1:  "ERROR_REASON_SCHEMA_PARSE_ERROR",
		2:  "ERROR_REASON_SCHEMA_TYPE_ERROR",
		3:  "ERROR_REASON_UNKNOWN_DEFINITION",
		4:  "ERROR_REASON_UNKNOWN_RELATION_OR_PERMISSION",
		5:  "ERROR_REASON_TOO_MANY_UPDATES_IN_REQUEST",
		6:  "ERROR_REASON_TOO_MANY_PRECONDITIONS_IN_REQUEST",
		7:  "ERROR_REASON_WRITE_OR_DELETE_PRECONDITION_FAILURE",
		8:  "ERROR_REASON_SERVICE_READ_ONLY",
		9:  "ERROR_REASON_UNKNOWN_CAVEAT",
		10: "ERROR_REASON_INVALID_SUBJECT_TYPE",
		11: "ERROR_REASON_CAVEAT_PARAMETER_TYPE_ERROR",
		12: "ERROR_REASON_UPDATES_ON_SAME_RELATIONSHIP",
		13: "ERROR_REASON_CANNOT_UPDATE_PERMISSION",
		14: "ERROR_REASON_CAVEAT_EVALUATION_ERROR",
		15: "ERROR_REASON_INVALID_CURSOR",
		16: "ERROR_REASON_TOO_MANY_RELATIONSHIPS_FOR_TRANSACTIONAL_DELETE",
		17: "ERROR_REASON_MAX_RELATIONSHIP_CONTEXT_SIZE",
		18: "ERROR_REASON_ATTEMPT_TO_RECREATE_RELATIONSHIP",
	}
	ErrorReason_value = map[string]int32{
		"ERROR_REASON_UNSPECIFIED":                                     0,
		"ERROR_REASON_SCHEMA_PARSE_ERROR":                              1,
		"ERROR_REASON_SCHEMA_TYPE_ERROR":                               2,
		"ERROR_REASON_UNKNOWN_DEFINITION":                              3,
		"ERROR_REASON_UNKNOWN_RELATION_OR_PERMISSION":                  4,
		"ERROR_REASON_TOO_MANY_UPDATES_IN_REQUEST":                     5,
		"ERROR_REASON_TOO_MANY_PRECONDITIONS_IN_REQUEST":               6,
		"ERROR_REASON_WRITE_OR_DELETE_PRECONDITION_FAILURE":            7,
		"ERROR_REASON_SERVICE_READ_ONLY":                               8,
		"ERROR_REASON_UNKNOWN_CAVEAT":                                  9,
		"ERROR_REASON_INVALID_SUBJECT_TYPE":                            10,
		"ERROR_REASON_CAVEAT_PARAMETER_TYPE_ERROR":                     11,
		"ERROR_REASON_UPDATES_ON_SAME_RELATIONSHIP":                    12,
		"ERROR_REASON_CANNOT_UPDATE_PERMISSION":                        13,
		"ERROR_REASON_CAVEAT_EVALUATION_ERROR":                         14,
		"ERROR_REASON_INVALID_CURSOR":                                  15,
		"ERROR_REASON_TOO_MANY_RELATIONSHIPS_FOR_TRANSACTIONAL_DELETE": 16,
		"ERROR_REASON_MAX_RELATIONSHIP_CONTEXT_SIZE":                   17,
		"ERROR_REASON_ATTEMPT_TO_RECREATE_RELATIONSHIP":                18,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_authzed_api_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_authzed_api_v1_error_reason_proto protoreflect.FileDescriptor

var file_authzed_api_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2a, 0xc7, 0x06, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x50, 0x41, 0x52, 0x53, 0x45, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12,
	0x2f, 0x0a, 0x2b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4f, 0x52, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x04,
	0x12, 0x2c, 0x0a, 0x28, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x05, 0x12, 0x32,
	0x0a, 0x2e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54,
	0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x50, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x06, 0x12, 0x35, 0x0a, 0x31, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x07, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x08, 0x12, 0x1f, 0x0a,
	0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x43, 0x41, 0x56, 0x45, 0x41, 0x54, 0x10, 0x09, 0x12, 0x25,
	0x0a, 0x21, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x0a, 0x12, 0x2c, 0x0a, 0x28, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x56, 0x45, 0x41, 0x54, 0x5f, 0x50, 0x41, 0x52,
	0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x0b, 0x12, 0x2d, 0x0a, 0x29, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x53, 0x5f, 0x4f, 0x4e, 0x5f, 0x53,
	0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50,
	0x10, 0x0c, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x28, 0x0a,
	0x24, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x41,
	0x56, 0x45, 0x41, 0x54, 0x5f, 0x45, 0x56, 0x41, 0x4c, 0x55, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0e, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x43, 0x55, 0x52, 0x53, 0x4f, 0x52, 0x10, 0x0f, 0x12, 0x40, 0x0a, 0x3c, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e,
	0x59, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x53, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x41,
	0x4c, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x10, 0x12, 0x2e, 0x0a, 0x2a, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x43, 0x4f, 0x4e, 0x54,
	0x45, 0x58, 0x54, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x11, 0x12, 0x31, 0x0a, 0x2d, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x41, 0x54, 0x54, 0x45, 0x4d,
	0x50, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x10, 0x12, 0x42, 0x48, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2d,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authzed_api_v1_error_reason_proto_rawDescOnce sync.Once
	file_authzed_api_v1_error_reason_proto_rawDescData = file_authzed_api_v1_error_reason_proto_rawDesc
)

func file_authzed_api_v1_error_reason_proto_rawDescGZIP() []byte {
	file_authzed_api_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_authzed_api_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v1_error_reason_proto_rawDescData)
	})
	return file_authzed_api_v1_error_reason_proto_rawDescData
}

var file_authzed_api_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authzed_api_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: authzed.api.v1.ErrorReason
}
var file_authzed_api_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authzed_api_v1_error_reason_proto_init() }
func file_authzed_api_v1_error_reason_proto_init() {
	if File_authzed_api_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authzed_api_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authzed_api_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_authzed_api_v1_error_reason_proto_enumTypes,
	}.Build()
	File_authzed_api_v1_error_reason_proto = out.File
	file_authzed_api_v1_error_reason_proto_rawDesc = nil
	file_authzed_api_v1_error_reason_proto_goTypes = nil
	file_authzed_api_v1_error_reason_proto_depIdxs = nil
}
