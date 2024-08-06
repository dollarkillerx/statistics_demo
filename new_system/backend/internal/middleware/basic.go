package middleware

import (
	"github.com/dollarkillerx/common/pkg/open_telemetry"
	"github.com/dollarkillerx/common/pkg/resp"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"fmt"
	"runtime"
	"runtime/debug"
)

// HttpRecover recover
func HttpRecover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				_, span := open_telemetry.Tracer.Start(ctx, fmt.Sprintf("HttpRecover%s", ctx.Request.URL.Path))

				e1 := errors.New(err.(string))
				stackTrace := debug.Stack()
				runtime.Stack(stackTrace, true)

				// add attributes to the span event
				attributes := []attribute.KeyValue{
					attribute.String("exception.stacktrace", string(stackTrace)),
				}

				log.Error().Msgf("HttpRecover url: %s stackTrace %s", ctx.Request.URL.Path, string(stackTrace))
				options := trace.WithAttributes(attributes...)

				// add error to span event
				span.RecordError(e1, options)

				// set event status to error
				span.SetStatus(codes.Error, e1.Error())

				span.End()

				resp.Return(ctx, 500, "Internal Server Error", nil)
			}
		}()
		ctx.Next()
	}
}
