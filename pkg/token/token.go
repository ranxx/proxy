package token

import (
	"context"

	"github.com/ranxx/proxy/constant"
	"github.com/ranxx/proxy/pkg/jwt"
)

// WithToken token
func WithToken(ctx context.Context, token *jwt.Claim) context.Context {
	return context.WithValue(ctx, constant.TokenKey, token)
}

// GetToken token
func GetToken(ctx context.Context) *jwt.Claim {
	value, ok := ctx.Value(constant.TokenKey).(*jwt.Claim)
	if ok {
		return value
	}
	return &jwt.Claim{}
}
