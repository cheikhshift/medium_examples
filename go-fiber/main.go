package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Database struct {
	Name string
}

func (d *Database) Get() string {
	return d.Name
}

func main() {
	app := fiber.New()
	db := Database{"Virtual_DB"}

	app.Use(func(c *fiber.Ctx) error {

		SetLocal[Database](c, "db", db)
		// Go to next middleware:
		return c.Next()
	})

	app.Get("/", GetRoot)

	app.Listen(":3000")
}

func GetRoot(c *fiber.Ctx) error {

	db := GetLocal[Database](c, "db")
	fmt.Println(db.Get())

	return c.SendString("Hello, World!")
}

func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}

func GetLocal[T any](c *fiber.Ctx, key string) T {
	return c.Locals(key).(T)
}
