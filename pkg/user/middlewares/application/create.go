package application

import (
	"context"
	"fmt"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Create ...
func (m *Middleware) Create(ctx context.Context, credentials models.Credentials) (output *models.User, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Create",
			"input", map[string]string{
				"email": fmt.Sprint(credentials.Email),
			},
			"output", fmt.Sprint(utils.GetSafeUser(output)),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Create(ctx, credentials)
	return
}
