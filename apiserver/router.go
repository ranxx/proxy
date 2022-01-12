package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/apiserver/clients"
	"github.com/ranxx/proxy/apiserver/tunnels"
	"github.com/ranxx/proxy/apiserver/users"
	websocket "github.com/ranxx/proxy/apiserver/ws"
	"github.com/ranxx/proxy/pkg/middleware"
)

// Router ...
func Router() *gin.Engine {
	engine := gin.New()

	api := engine.Group("/api")

	// ws
	ws := api.Group("/ws")
	{
		ws.GET("", middleware.EncodeGinMiddle(websocket.Ws).Gin)
	}

	// /docs/v1
	docs := api.Group("/docs/v1")
	{
		docs.StaticFS("", http.Dir("./docs"))
	}

	// /users/v1
	usersV1 := api.Group("/users/v1")
	{
		usersV1.POST("/register", middleware.EncodeGinResponse(users.Register).Gin)
		usersV1.POST("/login", middleware.EncodeGinResponse(users.Login).Gin)
	}

	// /tunnels/v1
	tunnelsV1 := api.Group("/tunnels/v1", middleware.EncodeGinMiddle(middleware.Auth).Gin)
	{
		tunnelsV1.GET("/list", middleware.EncodeGinResponse(tunnels.List).Gin)
		tunnelsV1.POST("/tunnel", middleware.EncodeGinResponse(tunnels.Create).Gin)
		tunnelsV1.PUT("/start/:id", middleware.EncodeGinResponse(tunnels.Start).Gin)
		tunnelsV1.PUT("/stop/:id", middleware.EncodeGinResponse(tunnels.Stop).Gin)
		tunnelsV1.DELETE("/tunnel/:id", middleware.EncodeGinResponse(tunnels.Delete).Gin)
	}

	// /clients/v1
	clientsV1 := api.Group("/clients/v1", middleware.EncodeGinMiddle(middleware.Auth).Gin)
	{
		clientsV1.GET("/list", middleware.EncodeGinResponse(clients.List).Gin)
		clientsV1.PUT("/stop/:id", middleware.EncodeGinResponse(clients.Stop).Gin)
		clientsV1.DELETE("/client/:id", middleware.EncodeGinResponse(clients.Delete).Gin)
	}
	return engine
}
