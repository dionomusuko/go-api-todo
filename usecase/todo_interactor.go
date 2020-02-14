package usecase

import (
	"github.com/dionomusuko/todos-go/domain"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Add(u domain.Todo) (err error) {
	_, err = interactor.TodoRepository.Store(u)
	return
}

func (interactor *TodoInteractor) Todos() (todo Todo.domain, err error) {
	todo, err = interactor.TodoRepository.FindAll()
	return
}

func (interactor *TodoInteractor) TodoById(identifier int) (todo domein.Todo, err error) {
	todo, err = interactor.TodoRepository.FindById(identifier)
	return
}
