package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wallace-jr/aguia-branca/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api")
	{
		trips := main.Group("trips")
		{
			trips.GET("/:id", controllers.ShowTrip)
		}
	}
	return router
}