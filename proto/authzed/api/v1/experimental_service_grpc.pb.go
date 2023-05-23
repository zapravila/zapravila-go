// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: authzed/api/v1/experimental_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ExperimentalService_BulkImportRelationships_FullMethodName = "/authzed.api.v1.ExperimentalService/BulkImportRelationships"
	ExperimentalService_BulkExportRelationships_FullMethodName = "/authzed.api.v1.ExperimentalService/BulkExportRelationships"
)

// ExperimentalServiceClient is the client API for ExperimentalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExperimentalServiceClient interface {
	// BulkImportRelationships is a faster path to writing a large number of
	// relationships at once. It is both batched and streaming. For maximum
	// performance, the caller should attempt to write relationships in as close
	// to relationship sort order as possible: (resource.object_type,
	// resource.object_id, relation, subject.object.object_type,
	// subject.object.object_id, subject.optional_relation)
	//
	// EXPERIMENTAL
	// https://github.com/authzed/spicedb/issues/1303
	BulkImportRelationships(ctx context.Context, opts ...grpc.CallOption) (ExperimentalService_BulkImportRelationshipsClient, error)
	// BulkExportRelationships is the fastest path available to exporting
	// relationships from the server. It is resumable, and will return results
	// in an order determined by the server.
	BulkExportRelationships(ctx context.Context, in *BulkExportRelationshipsRequest, opts ...grpc.CallOption) (ExperimentalService_BulkExportRelationshipsClient, error)
}

type experimentalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExperimentalServiceClient(cc grpc.ClientConnInterface) ExperimentalServiceClient {
	return &experimentalServiceClient{cc}
}

func (c *experimentalServiceClient) BulkImportRelationships(ctx context.Context, opts ...grpc.CallOption) (ExperimentalService_BulkImportRelationshipsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExperimentalService_ServiceDesc.Streams[0], ExperimentalService_BulkImportRelationships_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &experimentalServiceBulkImportRelationshipsClient{stream}
	return x, nil
}

type ExperimentalService_BulkImportRelationshipsClient interface {
	Send(*BulkImportRelationshipsRequest) error
	CloseAndRecv() (*BulkImportRelationshipsResponse, error)
	grpc.ClientStream
}

type experimentalServiceBulkImportRelationshipsClient struct {
	grpc.ClientStream
}

func (x *experimentalServiceBulkImportRelationshipsClient) Send(m *BulkImportRelationshipsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *experimentalServiceBulkImportRelationshipsClient) CloseAndRecv() (*BulkImportRelationshipsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BulkImportRelationshipsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *experimentalServiceClient) BulkExportRelationships(ctx context.Context, in *BulkExportRelationshipsRequest, opts ...grpc.CallOption) (ExperimentalService_BulkExportRelationshipsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExperimentalService_ServiceDesc.Streams[1], ExperimentalService_BulkExportRelationships_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &experimentalServiceBulkExportRelationshipsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExperimentalService_BulkExportRelationshipsClient interface {
	Recv() (*BulkExportRelationshipsResponse, error)
	grpc.ClientStream
}

type experimentalServiceBulkExportRelationshipsClient struct {
	grpc.ClientStream
}

func (x *experimentalServiceBulkExportRelationshipsClient) Recv() (*BulkExportRelationshipsResponse, error) {
	m := new(BulkExportRelationshipsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExperimentalServiceServer is the server API for ExperimentalService service.
// All implementations must embed UnimplementedExperimentalServiceServer
// for forward compatibility
type ExperimentalServiceServer interface {
	// BulkImportRelationships is a faster path to writing a large number of
	// relationships at once. It is both batched and streaming. For maximum
	// performance, the caller should attempt to write relationships in as close
	// to relationship sort order as possible: (resource.object_type,
	// resource.object_id, relation, subject.object.object_type,
	// subject.object.object_id, subject.optional_relation)
	//
	// EXPERIMENTAL
	// https://github.com/authzed/spicedb/issues/1303
	BulkImportRelationships(ExperimentalService_BulkImportRelationshipsServer) error
	// BulkExportRelationships is the fastest path available to exporting
	// relationships from the server. It is resumable, and will return results
	// in an order determined by the server.
	BulkExportRelationships(*BulkExportRelationshipsRequest, ExperimentalService_BulkExportRelationshipsServer) error
	mustEmbedUnimplementedExperimentalServiceServer()
}

// UnimplementedExperimentalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExperimentalServiceServer struct {
}

func (UnimplementedExperimentalServiceServer) BulkImportRelationships(ExperimentalService_BulkImportRelationshipsServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkImportRelationships not implemented")
}
func (UnimplementedExperimentalServiceServer) BulkExportRelationships(*BulkExportRelationshipsRequest, ExperimentalService_BulkExportRelationshipsServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkExportRelationships not implemented")
}
func (UnimplementedExperimentalServiceServer) mustEmbedUnimplementedExperimentalServiceServer() {}

// UnsafeExperimentalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExperimentalServiceServer will
// result in compilation errors.
type UnsafeExperimentalServiceServer interface {
	mustEmbedUnimplementedExperimentalServiceServer()
}

func RegisterExperimentalServiceServer(s grpc.ServiceRegistrar, srv ExperimentalServiceServer) {
	s.RegisterService(&ExperimentalService_ServiceDesc, srv)
}

func _ExperimentalService_BulkImportRelationships_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExperimentalServiceServer).BulkImportRelationships(&experimentalServiceBulkImportRelationshipsServer{stream})
}

type ExperimentalService_BulkImportRelationshipsServer interface {
	SendAndClose(*BulkImportRelationshipsResponse) error
	Recv() (*BulkImportRelationshipsRequest, error)
	grpc.ServerStream
}

type experimentalServiceBulkImportRelationshipsServer struct {
	grpc.ServerStream
}

func (x *experimentalServiceBulkImportRelationshipsServer) SendAndClose(m *BulkImportRelationshipsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *experimentalServiceBulkImportRelationshipsServer) Recv() (*BulkImportRelationshipsRequest, error) {
	m := new(BulkImportRelationshipsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExperimentalService_BulkExportRelationships_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BulkExportRelationshipsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExperimentalServiceServer).BulkExportRelationships(m, &experimentalServiceBulkExportRelationshipsServer{stream})
}

type ExperimentalService_BulkExportRelationshipsServer interface {
	Send(*BulkExportRelationshipsResponse) error
	grpc.ServerStream
}

type experimentalServiceBulkExportRelationshipsServer struct {
	grpc.ServerStream
}

func (x *experimentalServiceBulkExportRelationshipsServer) Send(m *BulkExportRelationshipsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ExperimentalService_ServiceDesc is the grpc.ServiceDesc for ExperimentalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExperimentalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authzed.api.v1.ExperimentalService",
	HandlerType: (*ExperimentalServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkImportRelationships",
			Handler:       _ExperimentalService_BulkImportRelationships_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BulkExportRelationships",
			Handler:       _ExperimentalService_BulkExportRelationships_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "authzed/api/v1/experimental_service.proto",
}
