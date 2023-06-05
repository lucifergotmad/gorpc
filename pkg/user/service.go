package user

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetUser(id uint64) (*User, error)
	CreateUser(user *User) (error)
	GenerateToken(user *User) (string, error)
	SecurePassword(password string) (string, error)
	CheckPasswordHash(password, hash string) (bool)
}

type Token struct {
	userID uint64
	jwt.StandardClaims
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

func (s *UserService) CreateUser(user *User) (error)  {
	return s.repo.CreateUser(user)
}

func (s *UserService) GenerateToken(user *User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	payload := &Token{userID: user.id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), payload)
	return token.SignedString([]byte(secretKey))
}

func (s *UserService) SecurePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *UserService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}