package main

import (
	"cshop-website/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
	app := fiber.New()

	app.Static("/assets", "./assets")

	// Secure your Fiber app by adding security headers
	app.Use(helmet.New(helmet.ConfigDefault))

	// Or extend your config for customization
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// // Or extend your config for customization
	// app.Use(favicon.New(favicon.Config{
	// 	File: "./assets/png/logo.png",
	// 	URL:  "./assets/png/logo.png",
	// }))

	landingPageHandler := handler.LandingPageHandler{}
	privacyPageHandler := handler.PrivacyPageHandler{}
	termsPageHandler := handler.TermsOfUsePageHandler{}
	app.Get("/", landingPageHandler.HandlerLandingPageShow)
	app.Get("/privacy", privacyPageHandler.HandlerPrivacyPageShow)
	app.Get("/terms", termsPageHandler.HandlerTermsOfUsePageShow)

	app.Get("/:lang", landingPageHandler.HandlerLandingPageShow)
	app.Get("/:lang/privacy", privacyPageHandler.HandlerPrivacyPageShow)
	app.Get("/:lang/terms", termsPageHandler.HandlerTermsOfUsePageShow)

	app.Listen(":3000")
}
