package ws

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/pkg/middleware"
	"github.com/ranxx/proxy/pkg/token"
	websocket "github.com/ranxx/proxy/service/ws"
)

// Ws websocket
func Ws(ctx *gin.Context) error {
	xctx, err := middleware.AuthToken(ctx, ctx.Query("token"))
	if err != nil {
		return err
	}

	ws, err := websocket.NewWebSocket(ctx.Writer, ctx.Request)
	if err != nil {
		return err
	}

	claim := token.GetToken(xctx)

	fmt.Println(claim)

	websocket.Global.Append(ctx, claim.UserID, ws)
	return nil
}
