package usecase

import (
	"github.com/dionomusuko/todos-go/domain"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Add(u domain.Todo) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.Store(u)
	return
}

func (interactor *TodoInteractor) Todos() (todos domain.Todos, err error) {
	todos, err = interactor.TodoRepository.FindAll()
	return
}

func (interactor *TodoInteractor) TodoById(id int) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.FindById(id)
	return
}
