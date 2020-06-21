package services

import (
	"context"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Create ...
func (s *services) Create(ctx context.Context, client *models.Client) (output *models.Client, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = utils.NilPointer()
			output = nil
			return
		}
	}()
	if client.Hash == "" || client.IP == "" {
		err = utils.MissingKey()
		return
	}
	output = client
	output.CreatedAt = time.Now()
	output.ExpiredAt = utils.Tomorrow()
	err = s.Store.Database.Model(client).Create(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
