package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxts0gt/go-ambassador/src/database"
	"github.com/maxts0gt/go-ambassador/src/models"
)

func Ambassadors(c *fiber.Ctx) error {

	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(users)
}
