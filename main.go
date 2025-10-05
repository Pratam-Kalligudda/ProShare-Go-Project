package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)
func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/health",func (ctx fiber.Ctx)error{
		return ctx.Status(fiber.StatusOK).JSON(map[string]string{"message":"healthy"})
	})
	app.Listen(":8080")
	fmt.Println("hello world")
}
