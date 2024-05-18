package handlers

import (
	"tinder-apps/internal/models"
	"tinder-apps/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetMembers(service *services.MemberService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		members, err := service.GetAllMembers()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get members",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Members found",
			"data":    members,
		})
	}
}

func CreateMember(service *services.MemberService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var member models.Member
		if err := c.BodyParser(&member); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse request body",
			})
		}

		if err := service.CreateMember(&member); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create member",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Member created",
			"data":    member,
		})
	}
}
