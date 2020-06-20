package application

import (
	"context"
	"fmt"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
)

// Create ...
func (m *Middleware) Create(ctx context.Context, distinctions models.Distinctions) (output *models.Client, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Create",
			"input", fmt.Sprint(distinctions),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Create(ctx, distinctions)
	return
}
