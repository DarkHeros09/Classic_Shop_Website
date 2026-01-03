package handler

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func render(ctx *fiber.Ctx, component templ.Component) error {
	return component.Render(ctx.Context(), ctx.Response().BodyWriter())
}
