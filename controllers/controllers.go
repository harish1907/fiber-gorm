package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harish1907/go-udemy/database"
	"github.com/harish1907/go-udemy/models"
	"github.com/harish1907/go-udemy/util"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(err)
	}

	if data["password"] != data["cpassword"] {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Password don't match",
		})
	}

	user := models.MyUser{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	user.SetPassword(data["password"])
	database.DB.Create(&user)
	return c.JSON(user)
}

func LoginAPI(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(err)
	}

	var user models.MyUser
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	token, err := util.GenerateToken(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJWT(cookie)
	var user models.MyUser
	database.DB.Find(&user, id)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
