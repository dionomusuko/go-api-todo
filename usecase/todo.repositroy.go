package usecase

import "github.com/dionomusuko/todos-go/domain"

type TodoRepository interface {
	Store(domain.Todo) (int, error)
	FindById(int) (domain.Todo, error)
	FindAll() (domain.Todos, error)
}
