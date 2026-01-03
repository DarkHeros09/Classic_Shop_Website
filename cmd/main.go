package main

import (
	"cshop-website/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/assets", "./assets",
		fiber.Static{
			Compress:      true,
			ByteRange:     true,
			CacheDuration: 10 * 7 * 24 * 60 * 60, // Better caching for fonts
		})

	orgPageHandler := handler.OrgPageHandler{}
	app.Get("/org", orgPageHandler.HandlerOrgPageShow)
	landingPageHandler := handler.LandingPageHandler{}
	app.Get("/", landingPageHandler.HandlerLandingPageShow)
	app.Get("/:lang", landingPageHandler.HandlerLandingPageShow)
	app.Listen(":3000")
}
