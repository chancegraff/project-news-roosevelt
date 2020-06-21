package services

import (
	"context"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Update ...
func (s *services) Update(ctx context.Context, user *models.User) (output *models.User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = utils.NilPointer()
			output = nil
			return
		}
	}()
	if user.ID == 0 && (user.Email == "" || user.Password == "") {
		err = utils.MissingKey()
		return
	}
	output = user
	output.UpdatedAt = time.Now()
	err = s.Store.Database.Model(user).Where("id = ?", user.ID).Update(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
