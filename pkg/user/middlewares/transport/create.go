package transport

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
)

// MakeCreateMiddleware ...
func MakeCreateMiddleware(logger log.Logger, next endpoints.Endpoints) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (output interface{}, err error) {
		defer func(begin time.Time) {
			logger.Log(
				"method", "Create",
				"event", "end",
				"took", time.Since(begin),
			)
		}(time.Now())
		logger.Log(
			"method", "Create",
			"event", "start",
			"input", fmt.Sprint(rq),
			"output", fmt.Sprint(output),
			"err", err,
		)
		output, err = next.CreateEndpoint(ctx, rq)
		return
	}
}
