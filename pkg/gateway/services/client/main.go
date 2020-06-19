package client

import (
	"github.com/pronuu/roosevelt/internal/database"
	"github.com/pronuu/roosevelt/internal/models"
)

// Service implements the ranker interface
type Service interface {
	Create(distinctions models.Distinctions) (models.Client, error)
	Retrieve(id uint, ip, hash, userID string) (models.Client, error)
}

type service struct {
	Store *database.Store
}

// NewService instantiates the service with a connection to the database
func NewService(s *database.Store) Service {
	return &service{Store: s}
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
