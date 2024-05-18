package handlers

import (
	"database/sql"
	"errors"
	"tinder-apps/internal/middleware"
	"tinder-apps/internal/models"
	"tinder-apps/internal/services"

	"github.com/gofiber/fiber/v2"
)

func ViewPartner(service *services.MemberService, partnerService *services.PartnerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := middleware.GetUserID(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		partnerList, err := partnerService.GetPartnerList(userID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "No new member",
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get members",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Members found",
			"data":    partnerList,
		})
	}
}

func SwipeAction(service *services.MemberService, partnerService *services.PartnerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var memberTarget models.PartnerRequest

		userID, err := middleware.GetUserID(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		if err := c.BodyParser(&memberTarget); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse request body",
			})
		}

		member_source, err := service.GetMemberById(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Member source not found",
			})
		}

		member_target, err := service.GetMemberById(memberTarget.MemberID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Member target not found",
			})
		}

		partnerCount, err := partnerService.GetPartnerCount(userID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Insufficient quota",
				})
			}
		}

		if partnerCount >= member_source.Quota {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Insufficient quota",
			})
		}

		partnerCheck, err := partnerService.GetPartnerCheck(userID, memberTarget.MemberID)

		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Member and target already connected",
				})
			}
		}

		if partnerCheck > 0 {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Member and target already connected",
			})
		}

		if member_source.ID == member_target.ID || member_source.Gender == member_target.Gender {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Member source cannot same with member target",
			})
		}

		if err := partnerService.UpdatePartnership(member_source.ID, &member_target, memberTarget.Direction); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update status",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Partnering member succeeded",
			"data":    memberTarget,
		})
	}
}
