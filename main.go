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

//testing some stuff 3

func main() {
	fmt.Println("Hi update 2")
	app := fiber.New()

	todos := []Todo{}

	//test()

	app.Get("/api/todos", func(req *fiber.Ctx) error {
		return req.Status(200).JSON(todos)
	})

	app.Post("/api/todos", func(req *fiber.Ctx) error {
		todo := &Todo{}

		fmt.Println(todo, &todo)

		if err := req.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return req.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		fmt.Println("Test", todo)
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return req.Status(201).JSON(todo)

	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "todo not found"})
	})

	log.Fatal(app.Listen(":8080"))
}

func test() {

	var x int = 5

	var p = &x

	fmt.Println(x, p, *p, &x)
}
