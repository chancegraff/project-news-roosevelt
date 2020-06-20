package models

import (
	"fmt"
	"strings"
	"time"
)

// Distinctions is a collection of specific information that identifies a unique client
type Distinctions struct {
	Hash   string `json:"hash,omitempty" gorm:"unique_index;not null"`
	UserID string `json:"user,omitempty"`
	IP     string `json:"ip,omitempty" gorm:"not null;index:ip"`
}

// Client is a representation of a unique machine in the database
type Client struct {
	Base
	Distinctions
	ExpiredAt time.Time `json:"expiredAt,omitempty"`
}

func (c *Client) stringID() string {
	asString := fmt.Sprintf("ID: %d,", c.ID)
	if c.ID == 0 {
		asString = ""
	}
	return asString
}

func (c *Client) stringHash() string {
	asString := fmt.Sprintf("Hash: %q,", c.Hash)
	if c.Hash == "" {
		asString = ""
	}
	return asString
}

func (c *Client) stringUserID() string {
	asString := fmt.Sprintf("UserID: %q,", c.UserID)
	if c.UserID == "" {
		asString = ""
	}
	return asString
}

func (c *Client) stringIP() string {
	asString := fmt.Sprintf("IP: %q,", c.IP)
	if c.IP == "" {
		asString = ""
	}
	return asString
}

func (c *Client) stringCreatedAt() string {
	asString := fmt.Sprintf("CreatedAt: %q,", c.CreatedAt)
	if c.CreatedAt.IsZero() {
		asString = ""
	}
	return asString
}

func (c *Client) stringUpdatedAt() string {
	asString := fmt.Sprintf("UpdatedAt: %q,", c.UpdatedAt)
	if c.UpdatedAt.IsZero() {
		asString = ""
	}
	return asString
}

func (c *Client) stringExpiredAt() string {
	asString := fmt.Sprintf("ExpiredAt: %q,", c.ExpiredAt)
	if c.ExpiredAt.IsZero() {
		asString = ""
	}
	return asString
}

// String ...
func (c *Client) String() string {
	if c == nil {
		return ""
	}
	withLinesAndTabs := fmt.Sprintf(`
		Client{
			%s
			%s
			%s
			%s
			%s
			%s
			%s
		}
	`,
		c.stringID(),
		c.stringHash(),
		c.stringUserID(),
		c.stringIP(),
		c.stringCreatedAt(),
		c.stringUpdatedAt(),
		c.stringExpiredAt(),
	)
	withLines := strings.Replace(withLinesAndTabs, "\t", "", -1)
	sanitized := strings.Replace(withLines, "\n", "", -1)
	return sanitized
}
