package services

import (
	"context"
	"errors"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
)

// Create ...
func (s *services) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("Invalid request")
	}
	user.CreatedAt = time.Now()
	err := s.Store.Database.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
