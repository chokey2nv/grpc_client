package main

import (
	"fmt"
	"log"
	"net"

	"github.com/chokey2nv/grpc_client/proto/v1"
	actions "github.com/chokey2nv/grpc_client/proto/v1/actions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	PORT = 8000
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatal("listener error: ", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterServiceServer(grpcServer, &actions.Server{})
	reflection.Register(grpcServer) //serialize and deserialize

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("grpc server: ", err)
	}
}
