package services

import (
	"context"
	"fmt"

	"github.com/gabrieldesousah/fc2-grpc/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func(*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Ex.: Insert on database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}