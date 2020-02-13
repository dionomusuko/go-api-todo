package controllers

import "strconv"

type TodoController struct {
	Interactor Todocase.TodoInteractor
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

func (controller *TodoController) Create(c Context) {
	u := domain.Todo{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201)
}

func (controller *TodoController) Index(c Context) {
	Todos, err := controller.Interactor.Todos()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Todos)
}

func (controller *TodoController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Todo, err := controller.Interactor.TodoById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Todo)
}
