package main

import (
	"github.com/christiandwi/showcase/interfaces/container"
	"github.com/christiandwi/showcase/interfaces/server"
)

func main() {
	server.Start(container.SetupContainer())
}
