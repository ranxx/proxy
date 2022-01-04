package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/config"
	jwt "github.com/ranxx/proxy/pkg/jwt"
)

// Auth 校验
func Auth(ctx *gin.Context) error {
	tokenStr := ctx.GetHeader("Authorization")

	claim := jwt.Claim{}

	key := []byte(config.GetConfig().Auth.Key)

	if err := jwt.Parse(key, tokenStr, &claim, nil); err != nil {
		return err
	}

	ctx.Set("Token", claim)

	ctx.Next()
	return nil
}
