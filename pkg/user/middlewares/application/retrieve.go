package application

import (
	"context"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
)

// Retrieve ...
func (m *Middleware) Retrieve(ctx context.Context, user *models.User) (output *models.User, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Retrieve",
			"input", user.String(),
			"output", output.String(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Retrieve(ctx, user)
	return
}
