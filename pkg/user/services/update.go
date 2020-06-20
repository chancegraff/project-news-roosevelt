package services

import (
	"context"
	"errors"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
)

// Update ...
func (s *services) Update(ctx context.Context, user *models.User) (*models.User, error) {
	if user.ID == 0 && (user.Email == "" || user.Password == "") {
		return nil, errors.New("Invalid request")
	}
	user.UpdatedAt = time.Now()
	err := s.Store.Database.Model(user).Where("id = ?", user.ID).Update(*user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
