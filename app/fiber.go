package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pooladkhay/tinier-auth-service/client"
	"github.com/pooladkhay/tinier-auth-service/controller"
	"github.com/pooladkhay/tinier-auth-service/repository"
	"github.com/pooladkhay/tinier-auth-service/service"
)

func StartFiber() {
	cassandraSession := client.CassandraSession()
	userRepo := repository.NewUser(cassandraSession)
	authService := service.NewAuth(userRepo)
	authController := controller.NewAuth(authService)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(cors.New())
	app.Use(recover.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/login", authController.Login)
	auth.Post("/register", authController.Register)

	app.All("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})

	p := os.Getenv("PORT")
	fmt.Printf("Auth api is listening on %s...\r\n", p)
	err := app.Listen(fmt.Sprintf(":%s", p))
	if err != nil {
		log.Fatalf("Err listening on %s: %s\r\n", p, err)
	}
}
