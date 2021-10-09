package controllers

import (
	"app/database"
	"app/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/smtp"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token := RandStringRunes(36)

	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}
	database.DB.Create(&passwordReset)

	from := "admin@goauthbackend.com"
	to := []string{data["email"]}
	url := "http://127.0.0.1:3000/reset/" + token
	message := []byte("Click <a href=\"" + url + "\">here</a> to reset your password!")

	err := smtp.SendMail("host.docker.internal:1025", nil, from, to, message)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Token successfully generated!",
	})
}

func Reset(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}

	var passwordReset = models.PasswordReset{}
	if err := database.DB.Where("token = ?", data["token"]).Last(&passwordReset); err.Error != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Token",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Password Changed!",
	})

}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func asciiLowerCase() {
	for char := 'a'; char <= 'z'; char++ {
		fmt.Printf("%c", char)
	}
}

func asciiUpperCase() {
	for char := 'A'; char <= 'Z'; char++ {
		fmt.Printf("%c", char)
	}
}
