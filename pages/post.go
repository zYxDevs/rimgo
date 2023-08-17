package pages

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"

	"codeberg.org/rimgo/rimgo/api"
	"codeberg.org/rimgo/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

// Cursed function
func nextInTag(client *api.Client, tagname, sort, page, I string) string {
	i, err := strconv.Atoi(I)
	if err != nil || i < 0 {
		return ""
	}
	tag, err := client.FetchTag(tagname, sort, page)
	if err != nil {
		return ""
	}
	if i >= len(tag.Posts)-1 {
		pageNumber, _ := strconv.Atoi(page)
		tagn, err := client.FetchTag(tagname, sort, strconv.Itoa(pageNumber+1))
		if err != nil {
			return ""
		}
		return tagn.Posts[0].Link
	}
	return tag.Posts[i+1].Link
}

func HandlePost(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("X-Frame-Options", "DENY")

	post, err := api.Album{}, error(nil)
	switch {
	case strings.HasPrefix(c.Path(), "/a"):
		post, err = ApiClient.FetchAlbum(c.Params("postID"))
	case strings.HasPrefix(c.Path(), "/gallery"):
		post, err = ApiClient.FetchPosts(c.Params("postID"))
	case strings.HasPrefix(c.Path(), "/t"):
		post, err = ApiClient.FetchPosts(c.Params("postID"))
	default:
		post, err = ApiClient.FetchMedia(c.Params("postID"))
	}
	if err != nil && err.Error() == "ratelimited by imgur" {
		return c.Status(429).Render("errors/429", fiber.Map{
			"path": c.Path(),
		})
	}
	if err != nil && post.Id == "" && strings.Contains(err.Error(), "404") {
		return c.Status(404).Render("errors/404", nil)
	}
	if err != nil {
		return err
	}

	comments := []api.Comment{}
	if post.SharedWithCommunity {
		c.Set("Cache-Control", "public,max-age=604800")
		comments, err = ApiClient.FetchComments(c.Params("postID"))
		if err != nil {
			return err
		}
	} else {
		c.Set("Cache-Control", "public,max-age=31557600")
	}

	nonce := ""
	csp := "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'self'; media-src 'self'; img-src 'self'; manifest-src 'self'; block-all-mixed-content; style-src 'self'"
	if len(post.Tags) != 0 {
		b := make([]byte, 8)
		rand.Read(b)
		nonce = fmt.Sprintf("%x", b)
		csp = csp + " 'nonce-" + nonce + "'"
	}
	c.Set("Content-Security-Policy", csp)

	var next string
	tagParam := strings.Split(c.Query("tag"), ".")
	if len(tagParam) == 4 {
		tag, sort, page, index := tagParam[0], tagParam[1], tagParam[2], tagParam[3]
		next = nextInTag(ApiClient, tag, sort, page, index)
	}

	return c.Render("post", fiber.Map{
		"post":     post,
		"next":     next,
		"comments": comments,
		"nonce":    nonce,
	})
}
