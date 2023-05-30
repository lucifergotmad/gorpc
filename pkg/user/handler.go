package user

import (
	"context"

	pb "gorpc/api/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	serv IUserService
	pb.UnimplementedUserServiceServer
}

func NewHandler(serv IUserService) *UserHandler {
	return &UserHandler{
		serv: serv,
	}
}


func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := h.serv.GetUser(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to retrieve user!")
	}

	pbUser := &pb.User{
		Id: user.id,
		Name: user.name,
		Email: user.email,
		Username: user.username,
		Password: user.password,
	}

	return pbUser, nil
}
