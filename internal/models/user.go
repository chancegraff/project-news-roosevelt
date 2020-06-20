package models

import (
	"fmt"
	"strings"
	"time"
)

// Credentials is an email and password combination
type Credentials struct {
	Email    string `json:"email,omitempty" gorm:"unique_index;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
}

// User is the representation of a user in the database
type User struct {
	Base
	Credentials
	VerifiedAt time.Time `json:"verified_at,omitempty"`
}

func (u *User) stringID() string {
	asString := fmt.Sprintf("ID: %d,", u.ID)
	if u.ID == 0 {
		asString = ""
	}
	return asString
}

func (u *User) stringEmail() string {
	asString := fmt.Sprintf("Email: %q,", u.Email)
	if u.Email == "" {
		asString = ""
	}
	return asString
}

func (u *User) stringCreatedAt() string {
	asString := fmt.Sprintf("CreatedAt: %q,", u.CreatedAt)
	if u.CreatedAt.IsZero() {
		asString = ""
	}
	return asString
}

func (u *User) stringUpdatedAt() string {
	asString := fmt.Sprintf("UpdatedAt: %q,", u.UpdatedAt)
	if u.UpdatedAt.IsZero() {
		asString = ""
	}
	return asString
}

func (u *User) stringVerifiedAt() string {
	asString := fmt.Sprintf("VerifiedAt: %q,", u.VerifiedAt)
	if u.VerifiedAt.IsZero() {
		asString = ""
	}
	return asString
}

// String ...
func (u *User) String() string {
	if u == nil {
		return ""
	}
	withLinesAndTabs := fmt.Sprintf(`
		User{
			%s
			%s
			%s
			%s
			%s
		}
	`,
		u.stringID(),
		u.stringEmail(),
		u.stringCreatedAt(),
		u.stringUpdatedAt(),
		u.stringVerifiedAt(),
	)
	withLines := strings.Replace(withLinesAndTabs, "\t", "", -1)
	sanitized := strings.Replace(withLines, "\n", "", -1)
	return sanitized
}
