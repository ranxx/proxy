package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/constant"
	"github.com/ranxx/proxy/pkg/jwt"
	"github.com/ranxx/proxy/pkg/token"
)

// AuthToken ...
func AuthToken(ctx context.Context, tokenStr string) (context.Context, error) {
	if len(tokenStr) <= 0 {
		return ctx, fmt.Errorf("token is empty")
	}

	claim := jwt.Claim{}

	key := []byte(config.GetConfig().Auth.Key)

	if err := jwt.Parse(key, tokenStr, &claim, nil); err != nil {
		return ctx, err
	}
	return token.WithToken(ctx, &claim), nil
}

// Auth 校验
func Auth(ctx *gin.Context) error {
	tokenStr := ctx.GetHeader("Authorization")
	if len(tokenStr) <= 0 {
		return fmt.Errorf("token is empty")
	}

	claim := jwt.Claim{}

	key := []byte(config.GetConfig().Auth.Key)

	if err := jwt.Parse(key, tokenStr, &claim, nil); err != nil {
		return err
	}

	token.WithToken(ctx, &claim)
	ctx.Set(constant.TokenKey, &claim)

	ctx.Next()
	return nil
}
