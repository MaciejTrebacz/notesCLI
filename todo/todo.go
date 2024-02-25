package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Todo struct {
	Content string
}

func NewTodo(content string) (*Todo, error) {
	return &Todo{
		content,
	}, errors.New("Invalid input")
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {

		return err
	}
	return os.WriteFile(fileName, json, 0644) // error generating error too or nil if not error
}
