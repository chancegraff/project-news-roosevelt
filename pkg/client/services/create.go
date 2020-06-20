package services

import (
	"context"
	"errors"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Create ...
func (s *services) Create(ctx context.Context, client *models.Client) (*models.Client, error) {
	if client.Hash == "" || client.IP == "" {
		return nil, errors.New("Invalid request")
	}
	client.CreatedAt = time.Now()
	client.ExpiredAt = utils.Tomorrow()
	err := s.Store.Database.Create(client).Error
	if err != nil {
		return nil, err
	}
	return client, nil
}
