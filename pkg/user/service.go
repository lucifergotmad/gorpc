package user

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user *User) error {
	// Business logic for creating a user.
	// Call the repository to perform the database operation.
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUser(userID int) (*User, error) {
	// Business logic for getting a user.
	// Call the repository to perform the database operation.
	return s.repo.GetUser(userID)
}
