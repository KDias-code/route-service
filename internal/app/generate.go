package app

import (
	"diplom/route-service/internal/handler"
	"diplom/route-service/internal/service"
	"diplom/route-service/internal/store"
	"diplom/route-service/pkg/config"
	"diplom/route-service/pkg/db"
	"github.com/gofiber/fiber/v3"
	"github.com/hashicorp/go-hclog"
)

type server struct {
	app      *fiber.App
	logger   hclog.Logger
	handlers handler.IHandler
}

func (s *server) generate(conf *config.Config) {
	s.app = fiber.New()

	s.logger = hclog.New(&hclog.LoggerOptions{
		Level:      2,
		JSONFormat: true,
	})

	dbConnection, err := db.ConnectPostgres(conf.DB)
	if err != nil {
		s.logger.Error("Failed to connect to database", "error", err)
		return
	}

	stores := store.NewStore(dbConnection)
	services := service.NewService(stores)
	s.handlers = handler.NewHandler(s.logger, services)

	s.router()
}
