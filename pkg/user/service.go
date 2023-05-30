package user

type IUserService interface {
	GetUser(id uint64) (*User, error)
}

type UserService struct {
	repo IUserRepository
}

func NewService(repo IUserRepository) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUser(id uint64) (*User, error) {
	return s.repo.GetUserByID(id)
}
