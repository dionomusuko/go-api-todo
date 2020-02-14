package controllers

import (
	"github.com/dionomusuko/todos-go/domain"
	"github.com/dionomusuko/todos-go/interfaces/database"
	"github.com/dionomusuko/todos-go/usecase"
	"github.com/labstack/echo"
	"strconv"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) Create(c echo.Context) (err error) {
	u := domain.Todo{}
	c.Bind(&u)
	todo, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, todo)
	return
}

func (controller *TodoController) Index(c echo.Context) (err error) {
	todos, err := controller.Interactor.Todos()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, todos)
	return
}

func (controller *TodoController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := controller.Interactor.TodoById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, todo)
	return
}
