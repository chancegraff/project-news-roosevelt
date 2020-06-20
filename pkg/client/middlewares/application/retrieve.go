package application

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pronuu/roosevelt/internal/models"
)

// Retrieve ...
func (m *Middleware) Retrieve(ctx context.Context, id uint, ip, hash, userID string) (output *models.Client, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Retrieve",
			"input", fmt.Sprint(map[string]string{
				"id":     strconv.Itoa(int(id)),
				"ip":     ip,
				"hash":   hash,
				"userID": userID,
			}),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Retrieve(ctx, id, ip, hash, userID)
	return
}
