package application

import (
	"context"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
)

// Retrieve ...
func (m *Middleware) Retrieve(ctx context.Context, client *models.Client) (output *models.Client, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Retrieve",
			"input", client.String(),
			"output", output.String(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Retrieve(ctx, client)
	return
}
