package services

import (
	"context"

	"github.com/pronuu/roosevelt/internal/database"
	"github.com/pronuu/roosevelt/internal/models"
)

// Service implements the collector interface
type Service interface {
	Create(ctx context.Context, credentials models.Credentials) (*models.User, error)
	Retrieve(ctx context.Context, id uint, email string) (*models.User, error)
	Update(ctx context.Context, id uint, credentials models.Credentials) (*models.User, error)
}

type service struct {
	Store *database.Store
}

// NewService instantiates the service with a connection to the database
func NewService(s *database.Store) Service {
	return &service{
		Store: s,
	}
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
