package services

import (
	"context"
	"errors"

	"github.com/pronuu/roosevelt/internal/models"
)

// Create ...
func (s *services) Create(ctx context.Context, credentials models.Credentials) (*models.User, error) {
	user := models.User{
		Credentials: credentials,
	}
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("Invalid request")
	}
	err := s.Store.Database.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
