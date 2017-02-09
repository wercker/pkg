package trace

import (
	"github.com/wercker/pkg/log"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

// ExposeInterceptor extracts the TraceID from the context and adds it to
// fields in the context.
func ExposeInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, next grpc.UnaryHandler) (interface{}, error) {
		traceID := ExtractTraceID(ctx)
		if traceID != "" {
			ctx, _ = log.AddFieldToCtx(ctx, TraceFieldKey, traceID)
		}

		return next(ctx, req)
	}
}
