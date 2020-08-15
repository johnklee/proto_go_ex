package main

import (
    "log"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "proto_go_ex/chat"
)

func main(){
    var conn *grpc.ClientConn
    conn, err := grpc.Dial(":9000", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("could not connect: %s", err)
    }

    defer conn.Close()
    c := pb.NewChatServiceClient(conn)
    message := pb.Message {
        Body: "Hello from the client!",
    }
    resp, err := c.SayHello(context.Background(), &message)
    if err != nil {
        log.Fatalf("Error when calling SayHello: %s", err)
    }

    log.Printf("Resp from Server: %s", resp.Body)
}
