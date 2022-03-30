package main

import (
	"bytes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func testHook(c *fiber.Ctx) error {
	body := c.Body()
	body = bytes.ReplaceAll(body, []byte("\n"), []byte("\\n"))
	log.Print(body)
	return nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := fiber.New()

	app.Post("/test", testHook)

	err := app.Listen(":3000")
	if err != nil {
		log.Error().Err(err).Msg("Error launching app")
	}
}
