package infrastructure

import (
	"github.com/dionomusuko/todos-go/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	todoController := controllers.NewTodoController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/todos", func(c echo.Context) error { return todoController.Index(c) })
	e.GET("/todos/:id", func(c echo.Context) error { return todoController.Show(c) })
	e.POST("/create", func(c echo.Context) error { return todoController.Create(c) })

	e.Logger.Fatal(e.Start(":1323"))
}
