package utils

import (
	"cshop-website/model"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(ctx *fiber.Ctx, component templ.Component) error {
	return component.Render(ctx.Context(), ctx.Response().BodyWriter())
}

func GetDir(lang string) string {
	if lang == "ar" {
		return "rtl"
	}
	return "ltr"
}

var yearStr = strconv.Itoa(time.Now().Year())

func CurrentYear(content model.Content) string {
	// Highly performant: only replaces if the placeholder exists
	return strings.ReplaceAll(content.FooterText, "{{year}}", yearStr)
}
