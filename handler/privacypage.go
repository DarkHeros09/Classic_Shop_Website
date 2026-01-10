package handler

import (
	"cshop-website/utils"
	"cshop-website/view/privacypage"

	"github.com/gofiber/fiber/v2"
)

type PrivacyPageHandler struct{}

func (l *PrivacyPageHandler) HandlerPrivacyPageShow(ctx *fiber.Ctx) error {
	lang := ctx.Params("lang", "ar")

	// Fallback logic for unsupported languages
	if lang != "en" && lang != "ar" {
		lang = "ar"
	}

	// 2. Fetch the correct content
	content, ok := utils.LandingPage[lang]
	if !ok {
		content = utils.LandingPage["ar"] // Fallback
	}

	// 2. Fetch the correct content
	privacyPolicy, ok := utils.PrivacyContent[lang]
	if !ok {
		privacyPolicy = utils.PrivacyContent["ar"] // Fallback
	}

	// 3. Set the Content-Type to HTML and render the Templ component
	ctx.Set("Content-Type", "text/html")
	return utils.Render(ctx, privacypage.Show(lang, content, privacyPolicy))
}
