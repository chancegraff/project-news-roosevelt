package client

import (
	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Create ...
func (s *service) Create(distinctions models.Distinctions) (*models.Client, error) {
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
