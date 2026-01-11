package main

import (
	"cshop-website/builder"
	"cshop-website/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
	app := fiber.New()

	app.Static("/assets", "/assets")

	app.Use(helmet.New(helmet.ConfigDefault))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Run the generator before starting the server
	builder.RunStaticGenerator()

	landingPageHandler := handler.LandingPageHandler{}
	privacyPageHandler := handler.PrivacyPageHandler{}
	termsPageHandler := handler.TermsOfUsePageHandler{}

	app.Get("/", landingPageHandler.HandlerLandingPageShow)
	app.Get("/privacy", privacyPageHandler.HandlerPrivacyPageShow)
	app.Get("/terms", termsPageHandler.HandlerTermsOfUsePageShow)

	app.Get("/:lang", landingPageHandler.HandlerLandingPageShow)
	app.Get("/:lang/privacy", privacyPageHandler.HandlerPrivacyPageShow)
	app.Get("/:lang/terms", termsPageHandler.HandlerTermsOfUsePageShow)

	log.Fatal(app.Listen(":3000"))
}
