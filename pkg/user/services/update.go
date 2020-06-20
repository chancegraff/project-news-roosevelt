package services

import (
	"context"

	"github.com/pronuu/roosevelt/internal/models"
)

// Update ...
func (s *service) Update(ctx context.Context, id uint, credentials models.Credentials) (*models.User, error) {
	user := models.User{
		Base: models.Base{
			ID: id,
		},
		Credentials: credentials,
	}
	err := s.Store.Database.Where(user).Update(user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
