package user

import (
	"errors"

	"github.com/andikaraditya/personal-library/internal/api"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	req := new(User)

	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return api.SendError(ctx, 400, err.Error(), err)
	}

	err := Service.CreateUser(*req)
	if err != nil {
		if errors.Is(err, api.ErrPayload) {
			return api.SendError(ctx, 400, "user already exists", err)
		}
		return api.SendError(ctx, 500, "internal server error", err)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}

func Login(ctx *fiber.Ctx) error {
	req := new(User)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return api.SendError(ctx, 400, err.Error(), err)
	}

	token, err := Service.Login(*req)
	if err != nil {
		if errors.Is(err, api.ErrPayload) {
			return api.SendError(ctx, 400, "email or password is incorrect", err)
		}
		return api.SendError(ctx, 500, "internal server error", err)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"token": token,
	})
}
