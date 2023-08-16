package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	bortlespb "github.com/mgazz0la/BortleBot/proto"
)

var (
	addr = flag.String("addr", "localhost:9001", "Address for the BortleBot backend")
	msg  = flag.String("msg", "Howdy y'all", "message to send")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := bortlespb.NewBortlesServiceClient(conn)

	resp, err := client.HelloWorld(context.Background(), &bortlespb.HelloRequest{Msg: *msg})
	if err != nil {
		log.Fatalf("hello world failed: %v", err)
	}

	log.Printf("[%s] received at %v", *msg, resp.Ts.AsTime())
}
