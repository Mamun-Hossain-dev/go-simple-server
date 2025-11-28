package user

import "fmt"

type Service interface {
	RegisterUser(u User) User
	LoginUser(email, pass string) (*User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) Service {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) RegisterUser(u User) User {
	return s.repo.Store(u)
}

func (s *UserService) LoginUser(email, pass string) (*User, error) {
	user := s.repo.Find(email, pass)

	if user == nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return user, nil
}
