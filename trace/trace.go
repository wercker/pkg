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
	"context"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go-opentracing"
)

// ExtractTraceID extracts the TraceID from a opentracing enabled context.
// Currently only the zipkin implementation is supported. Returns an empty
// string when no opentracing span was found, or a unsupported implementation
// was used.
func ExtractTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return ""
	}

	spancontext := span.Context()
	if spancontext == nil {
		return ""
	}

	switch s := spancontext.(type) {
	case zipkintracer.SpanContext:
		return s.TraceID.ToHex()
	default:
		return ""
	}
}
