package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harish1907/go-udemy/database"
	"github.com/harish1907/go-udemy/models"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Roles
	database.DB.Find(&roles)
	return c.JSON(roles)
}

func CreateRoles(c *fiber.Ctx) error {
	var role models.Roles
	if err := c.BodyParser(&role); err != nil {
		return err
	}
	database.DB.Create(&role)
	return c.JSON(role)
}

func GetSingleRole(c *fiber.Ctx) error {
	var role models.Roles
	database.DB.Find(&role, c.Params("id"))
	return c.JSON(role)
}

func UpdateRoles(c *fiber.Ctx) error {
	var role models.Roles
	database.DB.Find(&role, c.Params("id"))
	if err := c.BodyParser(&role); err != nil {
		return err
	}
	database.DB.Model(&role).Updates(role)
	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	var role models.Roles
	database.DB.Delete(&role, c.Params("id"))
	return c.JSON(fiber.Map{
		"message": "Role delete successfully.",
	})
}
