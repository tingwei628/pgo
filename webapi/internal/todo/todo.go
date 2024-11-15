package todo

import (
	"errors"
	"strings"
)

const (
	Pending    string = "Pending"
	InProgress string = "InProgress"
	Completed  string = "Completed"
	Failed     string = "Failed"
)

type Item struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}
type TodoService struct {
	todos []Item
}

func NewTodoService() *TodoService {
	return &TodoService{
		todos: []Item{},
	}
}
func (service *TodoService) Add(todo Item) error {

	for _, otherTodo := range service.todos {
		if todo.Task == otherTodo.Task {
			return errors.New("duplicate todo")
		}
	}
	service.todos = append(service.todos, Item{
		Task:   todo.Task,
		Status: Completed,
	})
	return nil
}

func (service *TodoService) GetAll() []Item {
	return service.todos
}

func (service *TodoService) Search(query string) []Item {
	var result []Item

	for _, todo := range service.todos {
		if strings.Contains(todo.Task, query) {
			result = append(result, todo)
		}
	}

	return result
}
