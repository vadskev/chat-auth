package main

import (
	"context"
	"log"
	"time"

	desc "github.com/vadskev/chat-auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
	userID  = 1
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed toconnect to server: %v", err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalf("failed to close server: %v", err)
		}
	}()

	c := desc.NewUserV1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: userID})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}

	user := r
	log.Printf("User info:\nID: %d\nName: %s\nEmail: %s\nRole: %s\nCreated At: %v\nUpdated At: %v\n",
		user.GetId(), user.GetName(), user.GetEmail(), user.GetRole().String(),
		user.GetCreatedAt().AsTime(), user.GetUpdatedAt().AsTime())
}
