package server

import (
	"flag"
	"fmt"

	"github.com/christiandwi/showcase/interfaces/container"
	"github.com/gin-gonic/gin"
)

func Start(container container.Container) *gin.Engine {
	addr := flag.String("addr: ", container.Config.App.Addr, "Address to listen and serve")
	app := gin.Default()

	// Setup Handler
	handler := setupHandler(container)

	// Setup Router
	setupRouter(app, handler)

	fmt.Println(app.Run(*addr))
	return app
}
