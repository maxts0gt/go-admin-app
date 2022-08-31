package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxts0gt/go-ambassador/src/database"
	"github.com/maxts0gt/go-ambassador/src/models"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Find(&orders)

	return c.JSON(orders)
}
