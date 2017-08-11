package grpc

import (
	"context"
	"fmt"

	"github.com/wercker/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// DebugUnaryServerInterceptor adds the fields in the log context
// as headers to the response.
// Keep in mind that these fields are only the ones that are
// set by earlier interceptors in the chain.
func DebugUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	fields, ok := log.FieldsFromContext(ctx)
	if !ok {
		return resp, err
	}

	m := map[string]string{}
	for k, v := range fields {
		if vs, ok := v.(string); ok {
			m[fmt.Sprintf("X-Debug-%s", k)] = vs
		}
	}
	md := metadata.New(m)

	grpc.SendHeader(ctx, md)
	return resp, err
}
