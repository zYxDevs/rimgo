package pages

import (
	"strconv"

	"codeberg.org/rimgo/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleTrending(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("X-Frame-Options", "DENY")
	c.Set("Cache-Control", "public,max-age=604800")
	c.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'self'; style-src 'unsafe-inline' 'self'; media-src 'self'; img-src 'self'; manifest-src 'self'; block-all-mixed-content")

	page := "1"
	if c.Query("page") != "" {
		page = c.Query("page")
	}

	pageNumber, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageNumber = 1
	}

	section := c.Query("section")
	switch section {
	case "hot", "new", "top":
	default:
		section = "hot"
	}
	sort := c.Query("sort")
	switch sort {
	case "newest", "best", "popular":
	default:
		sort = "popular"
	}

	results, err := ApiClient.FetchTrending(section, sort, page)
	if err != nil {
		return err
	}

	return c.Render("trending", fiber.Map{
		"results":     results,
		"section":     section,
		"sort":        sort,
		"page":        pageNumber,
		"nextPage":    pageNumber + 1,
		"prevPage":    pageNumber - 1,
	})
}
