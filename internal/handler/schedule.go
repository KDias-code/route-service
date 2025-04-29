package handler

import (
	"diplom/route-service/internal/models"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetSchedule(c fiber.Ctx) error {
	resp, err := h.service.GetSchedule()
	if err != nil {
		h.logger.Error("failed to get schedule", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *Handler) UpdateSchedule(c fiber.Ctx) error {
	var data models.Schedule
	body := c.Body()
	err := json.Unmarshal(body, &data)
	if err != nil {
		h.logger.Error("failed to unmarshal body", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if data.EndTime == "" || data.StartTime == "" || data.Id == 0 || data.RouteNumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON("empty one of the data")
	}

	err = h.service.ChangeSchedule(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(data)
}
