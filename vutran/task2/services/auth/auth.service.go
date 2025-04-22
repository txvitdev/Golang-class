package services

import (
	"context"
	dto "task2/dtos/auth"
	entities "task2/entities/user"
	"task2/exceptions"
	repositories "task2/repositories/user"
	services "task2/services/jwt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService{
	return &AuthService{userRepository}
}

func (authService *AuthService) SignUp(ctx context.Context, signUpDto *dto.SignUpDto) (entities.User, exceptions.HttpError) {
	// Check email existed
	isExisted, err := authService.userRepository.IsEmailExisted(ctx, signUpDto.Email)

	if err.Code != 0 {
		return entities.User{}, err
	}

	if isExisted {
		return entities.User{}, exceptions.NewConflict("Email existed")
	}

	hashString, cryptErr := hash(signUpDto.Password)

	if cryptErr != nil {
		return entities.User{}, exceptions.NewInternal("Something went wrong when sign up")
	}

	signUpDto.Password = hashString

	// Create user
	return  authService.userRepository.Save(ctx, signUpDto)
}

func (authService *AuthService) SignIn(ctx context.Context, signInDto *dto.SignInDto) (dto.SignInResponse, error) {
	user, err := authService.userRepository.FindByEmail(ctx, signInDto.Email)

	if err.Code != 0 {
		return dto.SignInResponse{}, err
	}

	// Generate Jwt token
	token, err := services.GenerateJwt(dto.JwtPayloadDto{
		Sub: user.Id,
	})

	if err.Code != 0 {
		return dto.SignInResponse{}, err
	}

	return dto.SignInResponse{
		AccessToken: token,
	}, exceptions.HttpError{}
}

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func verifyHash(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}