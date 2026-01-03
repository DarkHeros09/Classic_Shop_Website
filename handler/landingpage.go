package handler

import (
	"cshop-website/view/landingpage"

	"github.com/gofiber/fiber/v2"
)

type LandingPageHandler struct{}

func (l *LandingPageHandler) HandlerLandingPageShow(ctx *fiber.Ctx) error {
	// 1. Get language from query or default to Arabic
	lang := ctx.Params("lang", "ar")

	// Fallback logic for unsupported languages
	if lang != "en" && lang != "ar" {
		lang = "ar"
	}

	// 2. Fetch the correct content
	content, ok := translations[lang]
	if !ok {
		content = translations["ar"] // Fallback
	}

	// 3. Set the Content-Type to HTML and render the Templ component
	ctx.Set("Content-Type", "text/html")
	return render(ctx, landingpage.Show(lang, content))
}
