package user

import (
	"context"
	"log"

	pb "gorpc/api/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	service IUserService
	pb.UnimplementedUserServiceServer
}

func NewHandler(service IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := handler.service.GetUser(req.Id)
	log.Println("handler before: ", user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to retrieve user!")
	}

	log.Println("handler: ", user)

	pbUser := &pb.User{
		Id: user.id,
		Name: user.name,
		Email: user.email,
		Username: user.username,
		Password: user.password,
	}

	return pbUser, nil
}

func (handler *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	hashedPassword, err := handler.service.SecurePassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := User{
		 name: req.Name,
		 username: req.Username,
		 email: req.Email,
		 password: hashedPassword,
	}

	err = handler.service.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Message: "Success created user",
	}, nil
}
