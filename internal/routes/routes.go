package routes

import (
	"tinder-apps/internal/handlers"
	"tinder-apps/internal/repositories"
	"tinder-apps/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(app *fiber.App, db *sqlx.DB, secret string) {
	memberRepo := repositories.NewMemberRepository(db)
	memberService := services.NewMemberService(memberRepo)

	app.Post("/login", handlers.Login(memberService, secret))
	app.Post("/register", handlers.CreateMember(memberService))
}

func MemberRoutes(app *fiber.App, db *sqlx.DB) {
	memberRepo := repositories.NewMemberRepository(db)
	partnerRepo := repositories.NewPartnerRepository(db)
	purchaseRepo := repositories.NewPurchaseRepository(db)

	memberService := services.NewMemberService(memberRepo)
	partnerService := services.NewPartnerService(partnerRepo)
	purchaseService := services.NewPurchaseService(purchaseRepo)

	app.Get("/members", handlers.GetMembers(memberService))
	app.Get("/view", handlers.ViewPartner(memberService, partnerService))
	app.Post("/swipe", handlers.SwipeAction(memberService, partnerService))

	//Package & Purchase
	app.Get("/package", handlers.GetFeatures(purchaseService))
	app.Post("/purchase", handlers.CreatePurchase(memberService, purchaseService))
}
