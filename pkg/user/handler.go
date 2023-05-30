package user

import (
	"context"

	pb "gorpc/api/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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


func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	user, err := h.serv.GetUser(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to retrieve user!")
	}

	return user, nil
}
