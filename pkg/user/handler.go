package user

import (
	"context"

	pb "gorpc/api/user"
)

type UserHandler struct {
	serv UserService
	pb.UnimplementedUserServiceServer
}

func NewHandler(serv UserService) *UserHandler {
	return &UserHandler{
		serv: serv,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	// Call the service to perform the business logic and database operation.
	err := h.serv.CreateUser(&User{
		name:  req.GetName(),
		email: req.GetEmail(),
		password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Message: "User created successfully",
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	// Call the service to perform the business logic and database operation.
	user, err := h.serv.GetUser(req.GetId())
	if err != nil {
		return nil, err
	}

	return  &pb.User{
			Id:    user.id,
			Name:  user.name,
			Email: user.email,
		}, nil
}
