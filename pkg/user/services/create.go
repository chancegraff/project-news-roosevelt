package services

import (
	"context"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Create ...
func (s *services) Create(ctx context.Context, user *models.User) (output *models.User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = utils.NilPointer()
			output = nil
			return
		}
	}()
	if user.Email == "" || user.Password == "" {
		err = utils.MissingKey()
		return
	}
	output = user
	output.CreatedAt = time.Now()
	err = s.Store.Database.Model(user).Create(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
