package proto

import (
	"context"

	"github.com/chokey2nv/grpc_client/proto/v1"
)

type Server struct{}

func (s *Server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	a, b := req.GetA(), req.GetB()
	result := a + b
	return &proto.AddResponse{
		Result: result,
	}, nil
}
func (s *Server) Multiply(ctx context.Context, req *proto.MultiplyRequest) (*proto.MultiplyResponse, error) {
	a, b := req.GetA(), req.GetB()
	result := a * b
	return &proto.MultiplyResponse{
		Result: result,
	}, nil
}
