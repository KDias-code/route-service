package handler

import "github.com/gofiber/fiber/v3"

func (h *Handler) GetStops(c fiber.Ctx) error {
	resp, err := h.service.GetStops()
	if err != nil {
		h.logger.Error("Error getting stops", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON(resp)
}

func (h *Handler) NearStop(c fiber.Ctx) error {
	lat := c.Query("lat")
	if lat == "" {
		return c.Status(fiber.StatusBadRequest).JSON("null latitude")
	}

	lon := c.Query("lon")
	if lon == "" {
		return c.Status(fiber.StatusBadRequest).JSON("null longitude")
	}

	resp, err := h.service.NearStop(lat, lon)
	if err != nil {
		h.logger.Error("Error getting stops", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON(resp)
}
