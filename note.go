package main

import (
	"bufio"
	"encoding/json"
	"example.com/notesCLI/todo"
	"fmt"
	"os"
	"strings"
	"time"
)

type Saver interface {
	Save() error
}

type Note struct {
	Title     string `json:"titleTag___"`
	Content   string
	CreatedAt time.Time
}

func NewNote(title string, content string) *Note {
	return &Note{
		title,
		content,
		time.Now(),
	}
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {

		return err
	}

	return os.WriteFile(fileName, json, 0644) // error generating error too or nil if not error

}

func SaveData(data Saver) error {
	err := data.Save()
	HandleErrorPrompt(err, "")
	return err
}

func main() {
	title := getInput("Provide Title: ")
	content := getInput("Provide content: ")
	todoTest := getInput("Get todo text:")

	myTodo, err := todo.NewTodo(todoTest)
	HandleError(err)

	err = SaveData(myTodo)
	HandleError(err)

	note := NewNote(title, content)
	fmt.Printf("Your note titled %v has the following content: \n%v", note.Title, note.Content)
	err = SaveData(note)
	HandleError(err)
}

func HandleErrorPrompt(err error, prompt string) {
	if err != nil {
		fmt.Println(prompt)
	}
}
func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // on windows we have to remove sometimes /r as new line
	return text
}
