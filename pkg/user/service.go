package user

import pb "gorpc/api/user"

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUser(id string) (*pb.User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	pbUser := &pb.User{
		Id: user.id,
		Username: user.username,
		Email: user.email,
		Name: user.name,
	}

	return pbUser, nil
}
