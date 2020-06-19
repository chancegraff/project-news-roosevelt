package models

import (
	"time"
)

// Distinctions is a collection of specific information that identifies a unique client
type Distinctions struct {
	Hash   string `json:"hash" gorm:"unique_index"`
	UserID string `json:"user"`
	IP     string `json:"ip"`
}

// Client is a representation of a unique machine in the database
type Client struct {
	Base
	Distinctions
	ExpiredAt time.Time `json:"expiredAt"`
}
