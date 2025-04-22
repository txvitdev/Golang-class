package services

import (
	"context"
	dto "task2/internal/dtos/auth"
	entities "task2/internal/entities/user"
	"task2/internal/exceptions"
	clientRepo "task2/internal/repositories/client"
	userRepo "task2/internal/repositories/user"
	services "task2/internal/services/jwt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, email, password, clientID string) (dto.SignInResponse, exceptions.HttpError)
	Register(ctx context.Context, email, password, fullName string) exceptions.HttpError
}

type authService struct {
	userRepo   userRepo.UserRepository
	clientRepo clientRepo.ClientRepository
	jwtService services.JWTMaker
}

func NewAuthService(userRepo userRepo.UserRepository, clientRepo clientRepo.ClientRepository, jwtMaker services.JWTMaker) AuthService {
	return &authService{userRepo, clientRepo, jwtMaker}
}

func (s *authService) Login(ctx context.Context, email, password, clientID string) (dto.SignInResponse, exceptions.HttpError) {
	signInResponse := dto.SignInResponse{
		AccessToken: "",
	}

	client, err := s.clientRepo.GetClientByClientID(ctx, clientID)
	if err.Code != 0 {
		return signInResponse, err
	}

	user, err := s.userRepo.GetByEmail(ctx, email)
	if err.Code != 0 {
		return signInResponse, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return signInResponse, exceptions.NewBadRequest("Invalid credentials")
	}

	tokenStr, jwtErr := s.jwtService.CreateToken(user.Id, user.Email, client.ClientID)
	if jwtErr != nil {
		return signInResponse, exceptions.NewBadRequest("")
	}

	signInResponse.AccessToken = tokenStr

	return signInResponse, exceptions.HttpError{}
}

func (s *authService) Register(ctx context.Context, email, password, fullName string) exceptions.HttpError {
	_, err := s.userRepo.GetByEmail(ctx, email)
	if err.Code == 0 {
		return exceptions.NewConflict("Email already exists")
	}

	hashed, bcryptErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if bcryptErr != nil {
		return exceptions.NewInternal("")
	}

	user := &entities.User{
		Id:              uuid.NewString(),
		Email:           email,
		Password:        string(hashed),
		FullName:        fullName,
		IsEmailVerified: false,
	}

	return s.userRepo.CreateUser(ctx, user)
}
