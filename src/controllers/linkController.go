package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maxts0gt/go-ambassador/src/database"
	"github.com/maxts0gt/go-ambassador/src/models"
)

func Link(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var links []models.Link

	database.DB.Where("user_id = ?", id).Find(&links)

	return c.JSON(links)
}
