package main

import (
	"log"
	"project/internal/database"
	"project/internal/handlers"
	"project/internal/taskService"
	"project/internal/userService"
	"project/internal/web/tasks"
	"project/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	// database.DB.AutoMigrate(&taskService.Task{})
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	if err := database.DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Инициализируем task
	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Инициализируем user
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
