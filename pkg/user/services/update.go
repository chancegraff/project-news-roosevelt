package services

import (
	"context"
	"errors"

	"github.com/pronuu/roosevelt/internal/models"
)

// Update ...
func (s *services) Update(ctx context.Context, id uint, credentials models.Credentials) (*models.User, error) {
	user := models.User{
		Base: models.Base{
			ID: id,
		},
		Credentials: credentials,
	}
	if user.ID == 0 || user.Email == "" || user.Password == "" {
		return nil, errors.New("Invalid request")
	}
	err := s.Store.Database.Where(user).Update(user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
