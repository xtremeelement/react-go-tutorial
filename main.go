package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json: "id"`
	Completed bool   `json: "completed"`
	Body      string `json: "body"`
}

func main() {
	fmt.Println("Hi update 2")
	app := fiber.New()

	todos := []Todo{}

	test()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello World"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		fmt.Println(todo, *todo)

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		if todo == nil {
			return c.Status(400).JSON(fiber.Map{"error": "retrun some shit"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)

	})

	log.Fatal(app.Listen(":8080"))
}

func test() {

	var x int = 5

	var p = &x

	fmt.Println(x, p, *p)
}
