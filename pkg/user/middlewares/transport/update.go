package transport

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
)

// MakeUpdateMiddleware ...
func MakeUpdateMiddleware(logger log.Logger, next endpoints.Endpoints) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (output interface{}, err error) {
		defer func(begin time.Time) {
			logger.Log(
				"method", "Update",
				"input", fmt.Sprint(rq),
				"output", fmt.Sprint(output),
				"err", err,
				"took", time.Since(begin),
			)
		}(time.Now())
		output, err = next.UpdateEndpoint(ctx, rq)
		return
	}
}
