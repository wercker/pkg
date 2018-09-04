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
	grpcmw "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/wercker/pkg/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Interceptor adds a opentracing middleware, and exposes the TraceID.
func Interceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return grpcmw.ChainUnaryServer(
		otgrpc.OpenTracingServerInterceptor(tracer), // opentracing (incoming)
		ExposeInterceptor(),                         // expose traceID
	)
}

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
