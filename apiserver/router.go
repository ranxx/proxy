package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/apiserver/tunnels"
	"github.com/ranxx/proxy/apiserver/users"
	"github.com/ranxx/proxy/pkg/middleware"
)

// Router ...
func Router() *gin.Engine {
	engine := gin.New()

	api := engine.Group("/api")

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
	}

	return engine
}
