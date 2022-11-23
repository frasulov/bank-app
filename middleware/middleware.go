package middleware

import (
	"BankApp/config"
	"BankApp/errors"
	"BankApp/globals"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var Common = func(c *fiber.Ctx) error {
	c.Request().Header.Add("Content-Type", "application/json")
	return c.Next()
}

var tokenConfig = config.Configuration.Token

var Protect = func(c *fiber.Ctx) error {
	authorization := c.Get(tokenConfig.AuthorizationHeaderKey, "")
	if len(authorization) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(errors.NewResponseByKey("no_authorization_token", "EN"))
	}

	fields := strings.Fields(authorization)
	if len(fields) < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(errors.NewResponseByKey("invalid_authorization_token", "EN"))
	}

	if tokenConfig.AuthorizationTypeBearer != strings.ToLower(fields[0]) {
		return c.Status(fiber.StatusUnauthorized).JSON(errors.NewResponseByKey("unsupported_authorization_type", "EN"))
	}

	accessToken := fields[1]
	payload, err := globals.TokenMaker.VerifyToken(accessToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(errors.NewResponseByKey(err.Error(), "EN"))
	}
	c.Locals(tokenConfig.AuthorizationPayloadKey, payload)
	return c.Next()
}
