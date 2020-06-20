package services

import (
	"context"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Create ...
func (s *services) Create(ctx context.Context, distinctions models.Distinctions) (*models.Client, error) {
	client := models.Client{
		Distinctions: distinctions,
		ExpiredAt:    utils.Tomorrow(),
	}
	err := s.Store.Database.Create(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}
