package trace

import (
	"net/http"

	"github.com/wercker/pkg/log"
)

const (
	// TraceHTTPHeader is the header that will be used to expose the trace ID.
	TraceHTTPHeader = "X-Wercker-Trace-Id"

	// TraceFieldKey is the key that will be used for the field key.
	TraceFieldKey = "traceID"
)

// ExposeHandler decorates another http.Handler. It will check the context
// defined on the incoming http.Request for a traceID. If it is found it will
// add this to the response header and to the fields in the context.
func ExposeHandler(h http.Handler) http.Handler {
	return &traceExposer{h}
}

type traceExposer struct {
	h http.Handler
}

func (e *traceExposer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	traceID := ExtractTraceID(ctx)
	if traceID != "" {
		res.Header().Set(TraceHTTPHeader, traceID)
		ctx, _ = log.AddFieldToCtx(ctx, TraceFieldKey, traceID)
		req = req.WithContext(ctx)
	}

	e.h.ServeHTTP(res, req)
}

var _ http.Handler = (*traceExposer)(nil)
