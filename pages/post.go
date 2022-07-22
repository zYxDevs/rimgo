package pages

import (
	"crypto/rand"
	"fmt"
	"strings"

	"codeberg.org/video-prize-ranch/rimgo/api"
	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandlePost(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("X-Frame-Options", "DENY")

	post, err := api.Album{}, error(nil)
	switch {
	case strings.HasPrefix(c.Path(), "/a"):
		post, err = api.FetchAlbum(c.Params("postID"))
	case strings.HasPrefix(c.Path(), "/gallery"):
		post, err = api.FetchPosts(c.Params("postID"))
	default:
		post, err = api.FetchMedia(c.Params("postID"))
	}
	if post.Id == "" || (err != nil && strings.Contains(err.Error(), "404")) {
		c.Status(404)
		return c.Render("errors/404", nil)
	}
	if err != nil {
		return err
	}

	comments := []api.Comment{}
	if post.SharedWithCommunity {
		c.Set("Cache-Control", "public,max-age=604800")
		comments, err = api.FetchComments(c.Params("postID"))
		if err != nil {
			return err
		}
	} else {
		c.Set("Cache-Control", "public,max-age=31557600")
	}

	nonce := ""
	csp := "default-src 'none'; media-src 'self'; img-src 'self'; font-src 'self'; manifest-src 'self'; block-all-mixed-content; style-src 'self'"
	if len(post.Tags) != 0 {
		b := make([]byte, 8)
		rand.Read(b)
		nonce = fmt.Sprintf("%x", b)
		csp = csp + " 'nonce-" + nonce + "'"
	}
	c.Set("Content-Security-Policy", csp)

	return c.Render("post", fiber.Map{
		"post":     post,
		"comments": comments,
		"nonce":    nonce,
	})
}
