package services

import (
	"context"

	"github.com/pronuu/roosevelt/internal/database"
	"github.com/pronuu/roosevelt/internal/models"
)

// Services implements the collector interface
type Services interface {
	Create(ctx context.Context, credentials models.Credentials) (*models.User, error)
	Retrieve(ctx context.Context, id uint, email string) (*models.User, error)
	Update(ctx context.Context, id uint, credentials models.Credentials) (*models.User, error)
}

type services struct {
	Store *database.Store
}

// NewServices instantiates the service with a connection to the database
func NewServices(s *database.Store) Services {
	return &services{
		Store: s,
	}
}

// Middleware is a chainable middleware for Service
type Middleware func(Services) Services
