package infrastructure

import (
	"github.com/dionomusuko/todos-go/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	echo.New()

	todoController := controllers.NewTodoiController(NewSqlandler())

	// Middleware
	e.Todo(middleware.Logger())
	e.Todo(middleware.Recover())

	e.GET("/", hello)
	e.GET("/todos", func(c echo.Context) error { return todoController.Index(c) })

}
