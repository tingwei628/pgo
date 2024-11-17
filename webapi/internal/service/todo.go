package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/tingwei628/pgo/webapi/internal/database/repository"
	"github.com/tingwei628/pgo/webapi/internal/entity"
)

const (
	Pending    string = "Pending"
	InProgress string = "InProgress"
	Completed  string = "Completed"
	Failed     string = "Failed"
)

type TodoService struct {
	db repository.Todo
}

func NewTodoService(db repository.Todo) *TodoService {
	return &TodoService{
		db: db,
	}
}
func (service *TodoService) Add(todo entity.Item) error {

	items, err := service.GetAll()

	if err != nil {
		return err
	}

	for _, otherTodo := range items {
		if todo.Task == otherTodo.Task {
			return errors.New("duplicate todo")
		}
	}

	err = service.db.InsertItem(context.Background(), entity.Item{
		Task:   todo.Task,
		Status: Completed,
	})

	if err != nil {
		return err
	}

	return nil
}

func (service *TodoService) GetAll() ([]entity.Item, error) {

	items, err := service.db.GetAllItems(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}

	return items, nil
}

func (service *TodoService) Search(keyword string) ([]entity.Item, error) {

	items, err := service.db.GetItemsByKeyword(context.Background(), keyword)

	if err != nil {
		return nil, err
	}

	return items, nil
}
