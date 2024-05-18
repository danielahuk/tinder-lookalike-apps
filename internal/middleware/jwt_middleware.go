package middleware

import (
	"log"
	"time"

	"tinder-apps/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id int, name string, email string, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = name
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		tokenString = tokenString[7:]

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("could not load config: %v", err)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		c.Locals("id", token.Claims.(jwt.MapClaims)["id"])
		return c.Next()
	}
}

func GetUserID(c *fiber.Ctx) (int, error) {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	tokenString = tokenString[7:]

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return int(userID), nil
}
