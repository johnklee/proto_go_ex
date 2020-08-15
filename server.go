package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    pb "proto_go_ex/chat"
)

func main(){
    lis, err := net.Listen("tcp", ":9000")
    if err != nil {
        log.Fatalf("Fail to listen on port 9000: %v", err)
    }

    grpcServer := grpc.NewServer()
    s := pb.Server{}
    pb.RegisterChatServiceServer(grpcServer, &s)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
    }

}
