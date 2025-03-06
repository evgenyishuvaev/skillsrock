package server

import (
	"skillsrock/internal/app"
	"skillsrock/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(t *app.Tasker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := storage.Task{}
		if err := c.BodyParser(&task); err != nil {
			return c.Status(400).SendString("Неверный формат запроса")
		}

		err := t.CreateTask(task)
		if err != nil {
			return c.Status(500).SendString("Не удалось создать задачу.")
		}
		return nil
	}
}

func UpdateTask(t *app.Tasker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		task := storage.Task{}
		if err := c.BodyParser(&task); err != nil {
			return c.Status(400).SendString("Неверный формат запроса")
		}
		err := t.UpdateTask(id, task)
		if err != nil {
			return c.Status(500).SendString("Не удалось обновить задачу.")
		}
		return nil
	}
}

func DeleteTask(t *app.Tasker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := t.DeleteTask(id)
		if err != nil {
			return err
		}
		return nil
	}
}

func ListTask(t *app.Tasker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasks, err := t.ListTask()
		if err != nil {
			return c.Status(500).SendString("Не удалось получить список задач")
		}
		return c.JSON(tasks)
	}
}
