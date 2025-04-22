package services

import (
	"context"
	entities "task2/internal/entities/client"
	"task2/internal/exceptions"
	repositories "task2/internal/repositories/client"

	"github.com/google/uuid"
)

type ClientService interface {
	RegisterClient(ctx context.Context, name string, redirectURIs []string) (*entities.Client, exceptions.HttpError)
	ValidateClient(ctx context.Context, clientID, clientSecret string) (bool, exceptions.HttpError)
}

type clientService struct {
	repo repositories.ClientRepository
}

func NewClientService(repo repositories.ClientRepository) ClientService {
	return &clientService{repo}
}

func (s *clientService) RegisterClient(ctx context.Context, name string, redirectURIs []string) (*entities.Client, exceptions.HttpError) {
	client := &entities.Client{
		ID:           uuid.NewString(),
		Name:         name,
		ClientID:     uuid.NewString(),
		ClientSecret: uuid.NewString(),
		RedirectURIs: redirectURIs,
	}
	err := s.repo.CreateClient(ctx, client)
	return client, err
}

func (s *clientService) ValidateClient(ctx context.Context, clientID, clientSecret string) (bool, exceptions.HttpError) {
	client, err := s.repo.GetClientByClientID(ctx, clientID)
	if err.Code != 0 {
		return false, err
	}
	return client.ClientSecret == clientSecret, exceptions.HttpError{}
}
