package todo

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (todo Todo) Display() {
	fmt.Println(todo.Content)
}
func (todo Todo) HandleErrorPrompt(err error, prompt string) {
	if err != nil {
		fmt.Println(prompt)
	}
}
