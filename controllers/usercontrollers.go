package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harish1907/go-udemy/database"
	"github.com/harish1907/go-udemy/models"
)

func AllUser(c *fiber.Ctx) error {
	var users []models.MyUser
	database.DB.Find(&users)
	return c.JSON(users)
}

func CreateUsers(c *fiber.Ctx) error {
	var user models.MyUser
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.SetPassword("1234")
	database.DB.Create(&user)
	return c.JSON(user)
}


func GetSingleUser(c *fiber.Ctx) error {
	var user models.MyUser
	database.DB.Find(&user, c.Params("id"))
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.MyUser
	database.DB.Find(&user, c.Params("id"))
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.MyUser
	database.DB.Delete(&user, c.Params("id"))
	return c.JSON(fiber.Map{
		"message": "User delete successfully.",
	})
}