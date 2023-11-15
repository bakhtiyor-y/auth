package jwt

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

// ExtractTokenFromContext извлечь токен из контекста
func ExtractTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("No metadata found in context")
	}

	tokens := md.Get("authorization")
	if len(tokens) == 0 {
		return "", fmt.Errorf("No authorization token found in metadata")
	}

	token := strings.TrimPrefix(tokens[0], "Bearer ")

	return token, nil
}
