package handler

import (
	"cshop-website/view/org"

	"github.com/gofiber/fiber/v2"
)

type OrgPageHandler struct{}

func (l *OrgPageHandler) HandlerOrgPageShow(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html")
	return render(ctx, org.Show())
}
