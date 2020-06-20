package services

import (
	"context"
	"errors"

	"github.com/pronuu/roosevelt/internal/models"
)

// Retrieve ...
func (s *services) Retrieve(ctx context.Context, id uint, email string) (*models.User, error) {
	user := models.User{
		Base: models.Base{
			ID: id,
		},
		Credentials: models.Credentials{
			Email: email,
		},
	}
	if user.ID == 0 && user.Email == "" {
		return nil, errors.New("Invalid request")
	}
	err := s.Store.Database.Where(user).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
