//-----------------------------------------------------------------------------
// Copyright (c) 2017 Oracle and/or its affiliates.  All rights reserved.
// This program is free software: you can modify it and/or redistribute it
// under the terms of:
//
// (i)  the Universal Permissive License v 1.0 or at your option, any
//      later version (http://oss.oracle.com/licenses/upl); and/or
//
// (ii) the Apache License v 2.0. (http://www.apache.org/licenses/LICENSE-2.0)
//-----------------------------------------------------------------------------

package trace

import (
	"net/http"

	othttp "github.com/opentracing-contrib/go-stdlib/nethttp"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/wercker/pkg/log"
)

const (
	// TraceHTTPHeader is the header that will be used to expose the trace ID.
	TraceHTTPHeader = "X-Wercker-Trace-Id"

	// TraceFieldKey is the key that will be used for the field key.
	TraceFieldKey = "traceID"
)

// HTTPMiddleware adds a opentracing middleware, and exposes the TraceID.
func HTTPMiddleware(handler http.Handler, tracer opentracing.Tracer) http.Handler {
	handler = ExposeHandler(handler)          // expose traceID
	return othttp.Middleware(tracer, handler) // opentracing (incoming)
}

// ExposeHandler decorates another http.Handler. It will check the context
// defined on the incoming http.Request for a traceID. If it is found it will
// add this to the response header and to the fields in the context.
func ExposeHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		traceID := ExtractTraceID(ctx)
		if traceID != "" {
			res.Header().Set(TraceHTTPHeader, traceID)
			ctx, _ = log.AddFieldToCtx(ctx, TraceFieldKey, traceID)
			req = req.WithContext(ctx)
		}

		h.ServeHTTP(res, req)
	})
}
