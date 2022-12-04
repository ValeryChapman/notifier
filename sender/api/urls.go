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
		response := make(chan core.Response)
		go func(context *gin.Context) {
			response <- core.ErrorResponse(core.RouteNotFound)
		}(c.Copy())
		core.SendResponse(c, <-response)
	})

	router.NoMethod(func(c *gin.Context) {
		response := make(chan core.Response)
		go func(context *gin.Context) {
			response <- core.ErrorResponse(core.MethodNotAllowed)
		}(c.Copy())
		core.SendResponse(c, <-response)
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
