package user

import (
	"errors"
	"fmt"

	"spot_sync/internal/auth"
	"spot_sync/internal/user/dto"
)

var ErrInvalidCredentials = fmt.Errorf("invalid email or password")
type service struct {
	repo Repository
	jwtService auth.JWTService
}

func NewService(repo Repository,jwtService auth.JWTService) *service {
	return &service{repo: repo, jwtService: jwtService}
}

func (s *service) CreateUser(req *dto.RegisterUserRequest) (*dto.Response, error) {
	 user :=User{
	  Name: req.Name,
	  Email: req.Email,}
	  err := user.hashPassword(req.Password)
  if err != nil {
	  return nil, err
  }
  err = s.repo.CreateUser(&user)
  if err != nil {
	  return nil, err
  }
  response := dto.Response{
	  ID: user.ID,
	  Name: user.Name,
	  Email: user.Email,
	  CreatedAt: user.CreatedAt.String(),
  }

  return &response, nil
}

func (s *service) LoginUser(request *dto.LoginRequest)(*dto.Response, error) {
 user,err:=s.repo.GetUserByEmail(request.Email)
 if err != nil {
	 return nil, err
 }
 	if user == nil {
		return nil, ErrInvalidCredentials // User not found
	}
	err = user.checkPassword(request.Password)

	if err != nil {
		return nil, ErrInvalidCredentials
	}
	// generate token
	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Name)
	  fmt.Printf("DEBUG token: %q err: %v\n", token, err) // ← ADD
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}
	return &dto.Response{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Token: token,
		CreatedAt: user.CreatedAt.String(),
	}, nil
}


func (s *service) Update(id uint, req dto.UpdateUserRequest) (*dto.Response, error) {
	return nil, errors.New("not implemented")
}

func (s *service) Delete(id uint) error {
	return errors.New("not implemented")
}
