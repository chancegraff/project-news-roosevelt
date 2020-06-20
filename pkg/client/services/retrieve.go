package services

import (
	"context"

	"github.com/pronuu/roosevelt/internal/models"
)

// Retrieve ...
func (s *services) Retrieve(ctx context.Context, id uint, ip, hash, userID string) (*models.Client, error) {
	client := models.Client{
		Base: models.Base{
			ID: id,
		},
		Distinctions: models.Distinctions{
			Hash:   hash,
			UserID: userID,
			IP:     ip,
		},
	}
	err := s.Store.Database.Where(client).First(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}
