package application

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/internal/utils"
)

// Update ...
func (m *Middleware) Update(ctx context.Context, id uint, credentials models.Credentials) (output *models.User, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Update",
			"input", fmt.Sprint(map[string]string{
				"id":    strconv.Itoa(int(id)),
				"email": credentials.Email,
			}),
			"output", fmt.Sprint(utils.GetSafeUser(output)),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Update(ctx, id, credentials)
	return
}
