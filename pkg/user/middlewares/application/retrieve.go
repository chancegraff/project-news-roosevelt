package application

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Retrieve ...
func (m *Middleware) Retrieve(ctx context.Context, id uint, email string) (output *models.User, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Retrieve",
			"input", map[string]string{
				"id":    strconv.Itoa(int(id)),
				"email": email,
			},
			"output", fmt.Sprint(utils.GetSafeUser(output)),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Retrieve(ctx, id, email)
	return
}
