package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pooladkhay/tinier-auth-service/domain"
	"github.com/pooladkhay/tinier-auth-service/service"
)

type Auth interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type auth struct {
	service service.Auth
}

func NewAuth(s service.Auth) Auth {
	return &auth{service: s}
}

func (a *auth) Login(c *fiber.Ctx) error {
	reqBody := new(domain.LoginRequest)

	if err := c.BodyParser(reqBody); err != nil {
		return err
	}
	if err := reqBody.Validate(); err != nil {
		return c.Status(err.Status).JSON(err)
	}

	resp, err := a.service.Login(reqBody)
	if err != nil {
		return c.Status(err.Status).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(resp)
}

func (a *auth) Register(c *fiber.Ctx) error {
	reqBody := new(domain.RegisterRequest)

	if err := c.BodyParser(reqBody); err != nil {
		return err
	}
	if err := reqBody.Validate(); err != nil {
		return c.Status(err.Status).JSON(err)
	}

	err := a.service.Register(reqBody)
	if err != nil {
		return c.Status(err.Status).JSON(err)
	}

	// Login automatically after successful registration
	login := &domain.LoginRequest{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	}
	resp, err := a.service.Login(login)
	if err != nil {
		return c.Status(err.Status).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(resp)
}
