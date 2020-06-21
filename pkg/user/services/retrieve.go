package services

import (
	"context"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Retrieve ...
func (s *services) Retrieve(ctx context.Context, user *models.User) (output *models.User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = utils.NilPointer()
			output = nil
			return
		}
	}()
	if user.ID == 0 && user.Email == "" {
		err = utils.MissingKey()
		return
	}
	output = user
	err = s.Store.Database.Model(user).Where(user).First(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
