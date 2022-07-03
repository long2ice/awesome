package error

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/schema"
)

func Handler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	if _, ok := err.(validator.ValidationErrors); ok {
		code = fiber.StatusBadRequest
	}
	err = ctx.Status(code).JSON(schema.ErrorResponse{Error: err.Error()})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return nil
}
