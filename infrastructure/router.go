package infrastructure

import "github.com/labstack/echo"

func Init() {
	echo.New()

	todoController := controllers.NewTodoController(NewSqlandler())
}
