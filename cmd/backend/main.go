package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	bortlespb "github.com/mgazz0la/BortleBot/proto"
)

var port = flag.Int("port", 9001, "server port")

type (
	bortlesServer struct {
		bortlespb.UnimplementedBortlesServiceServer
	}
)

func (b *bortlesServer) HelloWorld(ctx context.Context, req *bortlespb.HelloRequest) (*bortlespb.HelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Request is nil")
	}

	log.Println(req.Msg)

	return &bortlespb.HelloResponse{
		Ts: timestamppb.Now(),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Unable to open port [%d]", *port)
	}
	log.Printf("listening on port [%d]", *port)

	s := grpc.NewServer()
	bs := &bortlesServer{}
	bortlespb.RegisterBortlesServiceServer(s, bs)

	log.Fatal(s.Serve(lis))
}
