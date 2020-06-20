package transport

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/pronuu/roosevelt/pkg/client/endpoints"
)

// MakeRetrieveMiddleware ...
func MakeRetrieveMiddleware(logger log.Logger, next endpoints.Endpoints) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (output interface{}, err error) {
		defer func(begin time.Time) {
			logger.Log(
				"method", "Retrieve",
				"event", "stop",
				"took", time.Since(begin),
			)
		}(time.Now())
		logger.Log(
			"method", "Retrieve",
			"event", "start",
			"input", fmt.Sprint(rq),
			"output", fmt.Sprint(output),
			"err", err,
		)
		output, err = next.RetrieveEndpoint(ctx, rq)
		return
	}
}
