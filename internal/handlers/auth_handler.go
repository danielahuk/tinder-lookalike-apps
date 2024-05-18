package handlers

import (
	"tinder-apps/internal/middleware"
	"tinder-apps/internal/models"
	"tinder-apps/internal/services"

	"github.com/gofiber/fiber/v2"
)

func Login(service *services.MemberService, secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var member models.Login
		if err := c.BodyParser(&member); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse request body",
			})
		}

		members, err := service.GetMemberByEmailPassword(member.Email, member.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Member not found",
			})
		}

		token, err := middleware.GenerateToken(members.ID, members.Name, members.Email, secret)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"mesasage": "Failed to generate token",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Login success",
			"data": fiber.Map{
				"token": token,
			},
		})
	}
}
