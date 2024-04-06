package main

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/leandergangso/pongo2components"
	_ "github.com/leandergangso/pongo2components/demo/components"
)

func init() {
	pongo2components.Init("demo/components")
	// pongo2components.InitFS("demo/components")
}

func main() {
	set, err := pongo2.DefaultSet.FromFile("demo/views/simple.html")
	if err != nil {
		panic(err)
	}

	s, err := set.Execute(make(pongo2.Context))
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("content-type", "text/html")
		return c.SendString(s)
	})

	err = app.Listen(":8080")
	panic(err)
}
