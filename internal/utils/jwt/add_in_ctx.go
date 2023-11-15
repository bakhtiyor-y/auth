package jwt

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func AddTokenToContext(ctx context.Context, token string) context.Context {
	md := metadata.Pairs("authorization", "Bearer "+token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}
