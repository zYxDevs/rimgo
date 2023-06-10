package pages

import (
	"codeberg.org/rimgo/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

var VersionInfo string

func HandleFrontpage(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("X-Frame-Options", "DENY")
	c.Set("Cache-Control", "public,max-age=31557600")
	c.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'none'; style-src 'self'; img-src 'self'; manifest-src 'self'; block-all-mixed-content")

	return c.Render("frontpage", fiber.Map{
		"config":  utils.Config,
		"version": VersionInfo,
	})
}