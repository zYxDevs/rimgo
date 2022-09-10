package pages

import (
	"strings"

	"codeberg.org/video-prize-ranch/rimgo/api"
	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleEmbed(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("Cache-Control", "public,max-age=31557600")
	c.Set("Content-Security-Policy", "default-src 'none'; media-src 'self'; style-src 'self'; img-src 'self'; font-src 'self'; block-all-mixed-content")

	post, err := api.Album{}, error(nil)
	switch {
	case strings.HasPrefix(c.Path(), "/a"):
		post, err = api.FetchAlbum(c.Params("postID"))
	case strings.HasPrefix(c.Path(), "/gallery"):
		post, err = api.FetchPosts(c.Params("postID"))
	default:
		post, err = api.FetchMedia(c.Params("postID"))
	}
	if err != nil && err.Error() == "ratelimited by imgur" {
		return c.Status(429).Render("errors/429", nil)
	}
	if post.Id == "" || (err != nil && strings.Contains(err.Error(), "404")) {
		return c.Status(404).Render("errors/404", nil)
	}
 	if err != nil {
		return err
	}

	return c.Render("embed", fiber.Map{
		"post": post,
	})
}

func HandleGifv(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("Cache-Control", "public,max-age=31557600")
	c.Set("Content-Security-Policy", "default-src 'none'; media-src 'self'; style-src 'self'; img-src 'self'; font-src 'self'; block-all-mixed-content")

	return c.Render("gifv", fiber.Map{
		"id": c.Params("postID"),
	})
}