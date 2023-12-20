package container

import (
	"github.com/christiandwi/showcase/config"
	"github.com/christiandwi/showcase/database"
	"github.com/christiandwi/showcase/infrastructure/repo"
	"github.com/christiandwi/showcase/lib/event"
	"github.com/christiandwi/showcase/usecase/guest"
)

type Container struct {
	Config       *config.Config
	GuestService guest.Service
	RabbitMq     event.RabbitMq
}

func SetupContainer() (out Container) {
	cfg := config.SetupConfig()

	db := database.SetupDatabase(cfg)

	rabbitMq := event.RabbitMqInit(cfg)

	userRepo := repo.SetupUserRepo(db)

	//setup asynq package
	// asynq := asynq.NewServer(
	// 	asynq.RedisClientOpt{
	// 		Addr: cfg.Redis.Address,
	// 	},
	// 	asynq.Config{
	// 		Concurrency: cfg.Asynq.Concurrency,
	// 	},
	// )

	//setup guest
	guestService := guest.NewGuestService(userRepo, rabbitMq)

	return Container{
		Config:       cfg,
		GuestService: guestService,
		RabbitMq:     rabbitMq,
	}
}
