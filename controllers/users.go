package controllers

import (
	"fmt"
	"pagoda/auth"
	"pagoda/database"
	"pagoda/middleware"
	"pagoda/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB

	user := new(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	})

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and password are required",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot hash password",
		})
	}

	newUser := models.User{
		Username: user.Username,
		Password: string(hashedPassword),
	}

	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Username already exists",
		})
	}

	sessionId, err := auth.GenerateSession(newUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create session",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": newUser,
		"key":  sessionId,
	})
}

func LoginUser(c *fiber.Ctx) error {
	db := database.DB

	user := new(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	})

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and password are required",
		})
	}

	var existingUser models.User

	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	sessionId, err := auth.GenerateSession(existingUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create session",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": existingUser,
		"key":  sessionId,
	})
}

func CurrentUser(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	return c.Status(fiber.StatusOK).JSON(user)
}

func LogoutUser(c *fiber.Ctx) error {
	session := new(middleware.Session)

	if err := c.CookieParser(session); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	auth.DeleteSession(session.SessionId)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged out",
	})
}
