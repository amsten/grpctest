package main

import (
	"context"
	"log"
	"net"

	"github.com/amsten/grpctest/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, in *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("Hello World"),
		Docx: []byte("Hello World from docx"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Could not listen on port 8080: %v", err)
	}
	defer lis.Close()

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}
