package pages

import (
	"codeberg.org/rimgo/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleSearch(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("X-Frame-Options", "DENY")
	c.Set("Cache-Control", "public,max-age=604800")
	c.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'self'; style-src 'unsafe-inline' 'self'; media-src 'self'; img-src 'self'; manifest-src 'self'; block-all-mixed-content")

	query := c.Query("q")

	if utils.ImgurRe.MatchString(query) {
		return c.Redirect(utils.ImgurRe.ReplaceAllString(query, ""))
	}

	/*page := "1"
	if c.Query("page") != "" {
		page = c.Query("page")
	}

	pageNumber, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageNumber = 0
	}

	displayPrevPage := true
	if page == "1" {
		displayPrevPage = false
	}*/

	results, err := ApiClient.Search(query)
	if err != nil {
		return err
	}

	return c.Render("search", fiber.Map{
		"query":	query,
		"results":        results,
		/*"page":        page,
		"displayPrev": displayPrevPage,
		"nextPage":    pageNumber + 1,
		"prevPage":    pageNumber - 1,*/
	})
}
