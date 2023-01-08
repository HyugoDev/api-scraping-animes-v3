package main

import (
	"github.com/HyugoDev/api-scraping-animes-v3/config"
	"github.com/HyugoDev/api-scraping-animes-v3/script/FLV"
	"github.com/HyugoDev/api-scraping-animes-v3/script/JKanime"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	app := fiber.New()

	app.Get("/JKanime/UDirectory", func(e *fiber.Ctx) error {
		return e.JSON(JKanime.GetDirectory())
	})

	app.Get("/JKanime/ERecientes", func(e *fiber.Ctx) error {
		return e.JSON(JKanime.GetERecientes())
	})

	app.Get("/FLV/ERecientes", func(e *fiber.Ctx) error {
		return e.JSON(FLV.GetERecientes())
	})

	app.Listen(config.GetPort())
}
