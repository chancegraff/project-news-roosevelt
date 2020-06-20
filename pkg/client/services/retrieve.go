package services

import (
	"context"
	"errors"

	"github.com/pronuu/roosevelt/internal/models"
)

// Retrieve ...
func (s *services) Retrieve(ctx context.Context, client *models.Client) (*models.Client, error) {
	if client.ID == 0 && client.Hash == "" && client.UserID == "" && client.IP == "" {
		return nil, errors.New("Invalid request")
	}
	err := s.Store.Database.Where(client).First(client).Error
	if err != nil {
		return nil, err
	}
	return client, nil
}
