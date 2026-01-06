package handler

import (
	"cshop-website/view/landingpage2"

	"github.com/gofiber/fiber/v2"

	"cshop-website/utils"
)

type LandingPage2Handler struct{}

func (l *LandingPage2Handler) HandlerLandingPage2Show(ctx *fiber.Ctx) error {
	// 1. Get language from query or default to Arabic
	lang := ctx.Params("lang", "ar")

	// Fallback logic for unsupported languages
	if lang != "en" && lang != "ar" {
		lang = "ar"
	}

	// 2. Fetch the correct content
	_, ok := utils.Translations[lang]
	if !ok {
		_ = utils.Translations["ar"] // Fallback
	}

	// 3. Set the Content-Type to HTML and render the Templ component
	ctx.Set("Content-Type", "text/html")
	return utils.Render(ctx, landingpage2.Show())
}
