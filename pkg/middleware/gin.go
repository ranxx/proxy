package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/pkg/components"
)

// EncodeGinResponse ...
type EncodeGinResponse func(ctx *gin.Context) (interface{}, error)

// Gin wrap
func (e EncodeGinResponse) Gin(ctx *gin.Context) {
	resp, err := e(ctx)
	if err != nil {
		components.ServerErrorEncoder(ctx, err, ctx.Writer)
		ctx.Abort()
		return
	}
	components.EncodeHTTPGenericResponse(ctx, ctx.Writer, resp)
}

// EncodeGinMiddle ...
type EncodeGinMiddle func(ctx *gin.Context) error

// Gin wrap
func (e EncodeGinMiddle) Gin(ctx *gin.Context) {
	err := e(ctx)
	if err == nil {
		return
	}
	components.ServerErrorEncoder(ctx, err, ctx.Writer)
	ctx.Abort()
}
