package services

import (
	"context"

	"github.com/pronuu/roosevelt/internal/models"
)

// Create ...
func (s *service) Create(ctx context.Context, credentials models.Credentials) (*models.User, error) {
	user := models.User{
		Credentials: credentials,
	}
	err := s.Store.Database.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
