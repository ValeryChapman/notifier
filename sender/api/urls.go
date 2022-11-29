package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"sender/api/core"
)

func InitRoutes() *gin.Engine {
	router := gin.New()
	gin.SetMode(viper.GetString("server.mode"))

	router.SetTrustedProxies([]string{
		viper.GetString("server.proxies"),
	})

	router.NoRoute(func(c *gin.Context) {
		core.ErrorResponse(c, core.RouteNotFound)
	})

	router.NoMethod(func(c *gin.Context) {
		core.ErrorResponse(c, core.MethodNotAllowed)
	})

	api := router.Group("/api")
	{
		notifications := api.Group("/notifications")
		{
			notifications.POST("/", CreateNotification)
			notifications.GET("/", GetNotifications)
			notifications.GET("/:id", GetNotification)
		}
	}
	return router
}
