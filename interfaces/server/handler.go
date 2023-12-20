package server

import (
	"github.com/christiandwi/showcase/interfaces/container"
)

type handler struct {
	guestHandler *guestHandler
}

func setupHandler(container container.Container) *handler {
	guestHandler := newGuestHandler(container.GuestService)
	return &handler{
		guestHandler: guestHandler,
	}
}
