package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json: "id"`
	Completed bool   `json: "completed"`
	Body      string `json: "body"`
}

func main() {
	app := fiber.New()
	err := godotenv.Load(".env")

	fmt.Println("teswtingggg")

	if err != nil {
		log.Fatal("error loading env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("teswtingggg", todos)
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			fmt.Println("POST", todo.Body)
			return c.Status(400).JSON(fiber.Map{"msg": "todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"msg": "not found"})

	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": "removed"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"msg": "not found"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
