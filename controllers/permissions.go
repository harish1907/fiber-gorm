package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harish1907/go-udemy/database"
	"github.com/harish1907/go-udemy/models"
)

func AllViewPermissions(c *fiber.Ctx) error {
	var permissions []models.Permissions
	database.DB.Find(&permissions)
	return c.JSON(permissions)
}