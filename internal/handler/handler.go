package handler

import (
	"diplom/route-service/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/hashicorp/go-hclog"
)

type IHandler interface {
	GetStops(c fiber.Ctx) error
	GetSchedule(c fiber.Ctx) error
	UpdateSchedule(c fiber.Ctx) error
	NearStop(c fiber.Ctx) error
}
type Handler struct {
	logger  hclog.Logger
	service service.IService
}

func NewHandler(logger hclog.Logger, svc service.IService) *Handler {
	return &Handler{
		logger:  logger,
		service: svc,
	}
}
