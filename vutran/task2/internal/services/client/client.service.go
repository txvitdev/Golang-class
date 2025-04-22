package services

import (
	"context"
	"fmt"
	clientEntity "task2/internal/entities/client"
	userEntity "task2/internal/entities/user"
	"task2/internal/exceptions"
	clientRepositories "task2/internal/repositories/client"
	roleRepositories "task2/internal/repositories/role"
	userRepositories "task2/internal/repositories/user"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ClientService interface {
	RegisterClient(ctx context.Context, name string, redirectURIs []string, adminEmail string,
		adminPassword string,
		adminFullName string) (*clientEntity.Client, exceptions.HttpError)
	ValidateClient(ctx context.Context, clientID, clientSecret string) (bool, exceptions.HttpError)
}

type clientService struct {
	clientRepo clientRepositories.ClientRepository
	userRepo   userRepositories.UserRepository
	roleRepo   roleRepositories.RoleRepository
}

func NewClientService(clientRepo clientRepositories.ClientRepository, userRepo userRepositories.UserRepository, roleRepo roleRepositories.RoleRepository) ClientService {
	return &clientService{clientRepo, userRepo, roleRepo}
}

func (s *clientService) RegisterClient(ctx context.Context, name string, redirectURIs []string, adminEmail string,
	adminPassword string,
	adminFullName string) (*clientEntity.Client, exceptions.HttpError) {
	client := &clientEntity.Client{
		ID:           uuid.NewString(),
		Name:         name,
		ClientID:     uuid.NewString(),
		ClientSecret: uuid.NewString(),
		RedirectURIs: redirectURIs,
	}
	err := s.clientRepo.CreateClient(ctx, client)

	if err.Code != 0 {
		return nil, err
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	adminUser := &userEntity.User{
		Id:              uuid.NewString(),
		Email:           adminEmail,
		Password:        string(hashed),
		FullName:        adminFullName,
		IsEmailVerified: true,
		CreatedAt:       time.Now(),
	}
	if err := s.userRepo.CreateUser(ctx, adminUser); err.Code != 0 {
		return nil, err
	}

	role, err := s.roleRepo.GetByName(ctx, "superadmin")
	if err.Code != 0 || role.ID == "" {
		_ = s.roleRepo.InsertRole(ctx, "superadmin")
	}
	err1 := s.roleRepo.AssignUserRole(ctx, adminUser.Id, role.ID)

	if err1.Code != 0 {
		fmt.Println(err1)
		return &clientEntity.Client{}, err1
	}

	return client, exceptions.HttpError{}
}

func (s *clientService) ValidateClient(ctx context.Context, clientID, clientSecret string) (bool, exceptions.HttpError) {
	client, err := s.clientRepo.GetClientByClientID(ctx, clientID)
	if err.Code != 0 {
		return false, err
	}
	return client.ClientSecret == clientSecret, exceptions.HttpError{}
}
