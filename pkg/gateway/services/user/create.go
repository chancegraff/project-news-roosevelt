package user

import (
	"github.com/pronuu/roosevelt/internal/models"
)

// Create ...
func (s *service) Create(credentials models.Credentials) (models.User, error) {
	user := models.User{
		Credentials: credentials,
	}
	err := s.Store.Database.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
