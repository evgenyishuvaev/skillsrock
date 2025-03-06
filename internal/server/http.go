package server

import (
	"log"
	tasker "skillsrock/internal/app"
	"skillsrock/internal/config"
	sql_storage "skillsrock/internal/storage/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewHttpServer(config config.Config) *fiber.App {

	app := fiber.New()
	app.Use(logger.New())

	storage, err := sql_storage.NewSQLStorage(config.DB.DSN())
	if err != nil {
		log.Fatalf("Cannot create Storage instance: %v", err)
	}
	tasker, err := tasker.NewTasker(storage)
	if err != nil {
		log.Fatalf("Cannot initialize App instance: %v", err)
	}

	api := app.Group("/api")

	api.Post("/tasks", CreateTask(tasker))
	api.Get("/tasks", ListTask(tasker))
	api.Delete("/tasks/:id", DeleteTask(tasker))
	api.Patch("/tasks/:id", UpdateTask(tasker))

	return app
}
