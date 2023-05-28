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
		Name:  req.GetName(),
		Email: req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Message: "User created successfully",
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Call the service to perform the business logic and database operation.
	user, err := h.serv.GetUser(req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
