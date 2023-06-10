package pages

import (
	"github.com/gofiber/fiber/v2"

	"codeberg.org/video-prize-ranch/rimgo/utils"
)

func HandlePrivacy(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("X-Frame-Options", "DENY")

	return c.Render("privacy", fiber.Map{
		"config":  utils.Config,
		"version": VersionInfo,
	})
}
