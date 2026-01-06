package handler

import (
	"cshop-website/utils"
	"cshop-website/view/termspage"

	"github.com/gofiber/fiber/v2"
)

type TermsOfUsePageHandler struct{}

func (l *TermsOfUsePageHandler) HandlerTermsOfUsePageShow(ctx *fiber.Ctx) error {
	lang := ctx.Params("lang", "ar")

	// Fallback logic for unsupported languages
	if lang != "en" && lang != "ar" {
		lang = "ar"
	}

	// 2. Fetch the correct content
	content, ok := utils.Translations[lang]
	if !ok {
		content = utils.Translations["ar"] // Fallback
	}

	// 2. Fetch the correct content
	termsOfUse, ok := utils.TermsContent[lang]
	if !ok {
		termsOfUse = utils.TermsContent["ar"] // Fallback
	}

	// 3. Set the Content-Type to HTML and render the Templ component
	ctx.Set("Content-Type", "text/html")
	return utils.Render(ctx, termspage.Show(lang, content, termsOfUse))
}
