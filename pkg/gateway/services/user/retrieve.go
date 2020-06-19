package user

import "github.com/pronuu/roosevelt/internal/models"

// Retrieve ...
func (s *service) Retrieve(id uint, email string) (models.User, error) {
	user := models.User{
		Base: models.Base{
			ID: id,
		},
		Credentials: models.Credentials{
			Email: email,
		},
	}
	err := s.Store.Database.Where(user).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
