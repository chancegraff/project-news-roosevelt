package utils

import (
	"os"
	"os/signal"

	"github.com/go-kit/kit/log"
	"github.com/pronuu/roosevelt/internal/models"
)

// GetEnvVar ...
func GetEnvVar(key string, def string) (string, bool) {
	variable := os.Getenv(key)
	if variable == "" {
		return def, false
	}
	return variable, true
}

// GetLogger will return a gokit logger with default params
func GetLogger(serviceName string) log.Logger {
	lgr := log.NewLogfmtLogger(os.Stderr)
	lgr = log.WithPrefix(lgr, "service", serviceName)
	lgr = log.WithPrefix(lgr, "timestamp", log.DefaultTimestampUTC)
	return lgr
}

// GetDoneChannel will create a new channel to listen for done signals on
func GetDoneChannel() *chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	return &done
}

// GetSafeUser will pull out and return the safe attributes from a user instance
func GetSafeUser(u *models.User) *models.User {
	if u == nil {
		return nil
	}
	return &models.User{
		Credentials: models.Credentials{
			Email: u.Email,
		},
		Base: models.Base{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		VerifiedAt: u.VerifiedAt,
	}
}
