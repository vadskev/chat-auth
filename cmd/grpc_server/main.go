package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/brianvoe/gofakeit"
	desc "github.com/vadskev/chat-auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

// Create
func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create: name: %s, email: %s, password: %s, password_confirm: %s, role: %s",
		req.Name, req.Email, req.Password, req.PasswordConfirmed, req.Role)

	id := genRandomID()
	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// Get
func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())
	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.UserRole_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

// Update
func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	log.Printf("Update user ID %d:", req.GetId())
	log.Printf("Update: id: %d, name: %s, email: %s, role: %s", req.Id, gofakeit.Name(), gofakeit.Email(), desc.UserRole(gofakeit.Number(0, 1)))
	//log.Printf("Update: id: %d, name: %d, email: %d, role: %d", req.Id, req.Name, req.Email, req.Role)
	return &desc.UpdateResponse{}, nil
}

// Delete
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	log.Printf("Delete user ID %d:", req.GetId())
	return &desc.DeleteResponse{}, nil
}

func genRandomID() int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
