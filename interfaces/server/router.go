package server

import "github.com/gin-gonic/gin"

func setupRouter(app *gin.Engine, handler *handler) {
	appRouter := app.Group("/api")
	v1 := appRouter.Group("/v1")

	//Guest grouping
	guest := v1.Group("/user")
	guest.GET("/login", handler.guestHandler.Login())
	guest.POST("/register", handler.guestHandler.CreateUser())
}
