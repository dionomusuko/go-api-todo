package usecase

import "github.com/dionomusuko/todos-go/domain"

type TodoRepository interface {
	Store(domain.Todo) (domain.Todo, error)
	FindById(id int) (domain.Todo, error)
	FindAll() (domain.Todos, error)
}
