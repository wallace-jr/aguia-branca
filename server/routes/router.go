package routes

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api")
	{
		trips := main.Group("trips")
		{
			trips.GET("/")
		}
	}
	return router
}