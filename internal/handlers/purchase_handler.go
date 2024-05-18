package handlers

import (
	"log"
	"tinder-apps/internal/middleware"
	"tinder-apps/internal/models"
	"tinder-apps/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetFeatures(service *services.PurchaseService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		features, err := service.GetFeaturesList()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get package",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Package found",
			"data":    features,
		})
	}
}

func CreatePurchase(service *services.MemberService, purchaseService *services.PurchaseService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var purchase models.PurchaseRequest
		if err := c.BodyParser(&purchase); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse request body",
			})
		}

		userID, err := middleware.GetUserID(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		member, err := service.GetMemberById(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Package not found",
			})
		}

		feature, err := purchaseService.GetFeatureById(purchase.PackageID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Package not found",
			})
		}

		if feature.Name == "Extend Quota" && member.Quota > 10 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "You have purchase this package",
			})
		} else if feature.Name == "Premium Member" && member.Label == "Premium" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "You have purchase this package",
			})
		}

		err = purchaseService.CreatePurchase(purchase, userID, feature.Price)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create purchase",
			})
		}

		var memberData models.Member

		if feature.Name == "Extend Quota" {
			memberData = models.Member{
				ID:    userID,
				Quota: 99,
				Label: member.Label,
			}
		} else if feature.Name == "Premium Member" {
			memberData = models.Member{
				ID:    userID,
				Quota: member.Quota,
				Label: "Premium",
			}
		}

		err = service.UpdateMember(&memberData)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update member status",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Package purchased",
			"data":    nil,
		})
	}
}
