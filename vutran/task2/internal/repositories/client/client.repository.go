package repositories

import (
	"context"
	"fmt"
	entities "task2/internal/entities/client"
	"task2/internal/exceptions"

	"github.com/jmoiron/sqlx"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, client *entities.Client) exceptions.HttpError
	GetClientByID(ctx context.Context, id string) (*entities.Client, exceptions.HttpError)
	GetClientByClientID(ctx context.Context, clientID string) (*entities.Client, exceptions.HttpError)
}

type clientRepo struct {
	db *sqlx.DB
}

func NewClientRepository(db *sqlx.DB) ClientRepository {
	return &clientRepo{db}
}

func (r *clientRepo) CreateClient(ctx context.Context, client *entities.Client) exceptions.HttpError {
	query := `
	INSERT INTO clients (id, name, client_id, client_secret, redirect_uris)
	VALUES (:id, :name, :client_id, :client_secret, :redirect_uris)`
	_, err := r.db.NamedExecContext(ctx, query, client)
	if err != nil {
		fmt.Println(err)
		return exceptions.NewInternal("")
	}

	return exceptions.HttpError{}
}

func (r *clientRepo) GetClientByID(ctx context.Context, id string) (*entities.Client, exceptions.HttpError) {
	var client entities.Client
	err := r.db.GetContext(ctx, &client, "SELECT * FROM clients WHERE id=$1", id)

	if err != nil {
		return &entities.Client{}, exceptions.NewNotFound(fmt.Sprintf("Not fount client with id = %s", id))
	}
	return &client, exceptions.HttpError{}
}

func (r *clientRepo) GetClientByClientID(ctx context.Context, clientID string) (*entities.Client, exceptions.HttpError) {
	var client entities.Client
	err := r.db.GetContext(ctx, &client, "SELECT * FROM clients WHERE client_id=$1", clientID)

	if err != nil {
		return &client, exceptions.NewNotFound(fmt.Sprintf("Not fount client with clientId = %s", clientID))
	}
	return &client, exceptions.HttpError{}
}
