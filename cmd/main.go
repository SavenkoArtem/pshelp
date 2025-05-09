package main

import (
	"fmt"

	"github.com/SavenkoArtem/pshelp/configs"
	"github.com/SavenkoArtem/pshelp/internal/todo"
	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)
	app := fiber.New()

	todo.NewHandler(app)

	app.Listen(":3000")
}
