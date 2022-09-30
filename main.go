package main

import (
	"flag"
	"fmt"
	"net/http"

	"codeberg.org/video-prize-ranch/rimgo/pages"
	"codeberg.org/video-prize-ranch/rimgo/static"
	"codeberg.org/video-prize-ranch/rimgo/utils"
	"codeberg.org/video-prize-ranch/rimgo/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/handlebars"
	"github.com/joho/godotenv"
)

func main() {
	envPath := flag.String("c", ".env", "Path to env file")
	err := godotenv.Load(*envPath)
	if err != nil {
		fmt.Println(err)
	}
	utils.LoadConfig()

	engine := handlebars.NewFileSystem(http.FS(views.GetFiles()), ".hbs")
	app := fiber.New(fiber.Config{
		Views:             engine,
		Prefork:           utils.Config["fiberPrefork"].(bool),
		UnescapePath:      true,
		StreamRequestBody: true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			err = ctx.Status(code).Render("errors/error", fiber.Map{
				"err": err,
			})
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			return nil
		},
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func (c *fiber.Ctx, e interface{})  {
			fmt.Println(e)
		},
	}))
	app.Use("/static", filesystem.New(filesystem.Config{
		Next: func(c *fiber.Ctx) bool {
			c.Response().Header.Add("Cache-Control", "public,max-age=2592000")
			return false
		},
		Root: http.FS(static.GetFiles()),
	}))

	app.Get("/robots.txt", func(c *fiber.Ctx) error {
		file, _ := static.GetFiles().ReadFile("robots.txt")
		_, err := c.Write(file)
		return err
	})
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		file, _ := static.GetFiles().ReadFile("favicon/favicon.ico")
		_, err := c.Write(file)
		return err
	})

	app.Get("/", pages.HandleFrontpage)
	app.Get("/privacy", pages.HandlePrivacy)
	app.Get("/:postID.gifv", pages.HandleGifv)
	app.Get("/:baseName.:extension", pages.HandleMedia)
	app.Get("/stack/:baseName.:extension", pages.HandleMedia)
	app.Get("/:postID", pages.HandlePost)
	app.Get("/:postID/embed", pages.HandleEmbed)
	app.Get("/a/:postID", pages.HandlePost)
	app.Get("/a/:postID/embed", pages.HandleEmbed)
	app.Get("/t/:tag", pages.HandleTag)
	app.Get("/user/:userID", pages.HandleUser)
	app.Get("/r/:sub/:postID", pages.HandlePost)
	app.Get("/user/:userID/cover", pages.HandleUserCover)
	app.Get("/user/:userID/avatar", pages.HandleUserAvatar)
	app.Get("/gallery/:postID", pages.HandlePost)
	app.Get("/gallery/:postID/embed", pages.HandleEmbed)

	app.Listen(utils.Config["addr"].(string) + ":" + utils.Config["port"].(string))
}
