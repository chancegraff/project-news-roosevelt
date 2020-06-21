package services

import (
	"context"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Retrieve ...
func (s *services) Retrieve(ctx context.Context, client *models.Client) (output *models.Client, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = utils.NilPointer()
			output = nil
			return
		}
	}()
	if client.ID == 0 && client.Hash == "" && client.UserID == "" && client.IP == "" {
		err = utils.MissingKey()
		return
	}
	output = client
	err = s.Store.Database.Model(client).Where(client).First(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
