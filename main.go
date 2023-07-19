package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"codeberg.org/rimgo/rimgo/pages"
	"codeberg.org/rimgo/rimgo/static"
	"codeberg.org/rimgo/rimgo/utils"
	"codeberg.org/rimgo/rimgo/views"
	"github.com/aymerick/raymond"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/handlebars"
	"github.com/joho/godotenv"
)

func main() {
	envPath := flag.String("c", ".env", "Path to env file")
	_ := godotenv.Load(*envPath)
	utils.LoadConfig()
	
	pages.InitializeApiClient()

	views := http.FS(views.GetFiles())
	if os.Getenv("ENV") == "dev" {
		views = http.Dir("./views")
	}
	engine := handlebars.NewFileSystem(views, ".hbs")
	
	engine.AddFunc("noteq", func(a interface{}, b interface{}, options *raymond.Options) interface{} {
		if raymond.Str(a) != raymond.Str(b) {
			return options.Fn()
		}
		return ""
	})

	app := fiber.New(fiber.Config{
		Views:             engine,
		Prefork:           utils.Config.FiberPrefork,
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
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			fmt.Println(e)
		},
	}))
	
	if os.Getenv("ENV") == "dev" {
		app.Use("/static", filesystem.New(filesystem.Config{
			MaxAge: 2592000,
			Root: http.Dir("./static"),
		}))
		app.Get("/errors/429", func(c *fiber.Ctx) error {
			return c.Render("errors/429", nil)
		})
		app.Get("/errors/404", func(c *fiber.Ctx) error {
			return c.Render("errors/404", nil)
		})
		app.Get("/errors/error", func(c *fiber.Ctx) error {
			return c.Render("errors/error", fiber.Map{
				"err": "Test error",
			})
		})
	} else {
		app.Use("/static", filesystem.New(filesystem.Config{
			Root: http.FS(static.GetFiles()),
		}))
		app.Use(cache.New(cache.Config{
			Expiration:           30 * time.Minute,
			MaxBytes:             25000000,
			CacheControl:         true,
			StoreResponseHeaders: true,
		}))
	}

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
	app.Get("/about", pages.HandleAbout)
	app.Get("/privacy", pages.HandlePrivacy)
	app.Get("/a/:postID", pages.HandlePost)
	app.Get("/a/:postID/embed", pages.HandleEmbed)
	app.Get("/t/:tag", pages.HandleTag)
	app.Get("/t/:tag/:postID", pages.HandlePost)
	app.Get("/user/:userID", pages.HandleUser)
	app.Get("/r/:sub/:postID", pages.HandlePost)
	app.Get("/user/:userID/cover", pages.HandleUserCover)
	app.Get("/user/:userID/avatar", pages.HandleUserAvatar)
	app.Get("/gallery/:postID", pages.HandlePost)
	app.Get("/gallery/:postID/embed", pages.HandleEmbed)
	app.Get("/:postID.gifv", pages.HandleGifv)
	app.Get("/:baseName.:extension", pages.HandleMedia)
	app.Get("/stack/:baseName.:extension", pages.HandleMedia)
	app.Get("/:postID", pages.HandlePost)
	app.Get("/:postID/embed", pages.HandleEmbed)

	app.Listen(utils.Config.Addr + ":" + utils.Config.Port)
}
