package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"titleTag____"`
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

func (note Note) Display() {
	fmt.Println(note.Content)
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
	return err
}
func (note Note) HandleErrorPrompt(err error, prompt string) {
	if err != nil {
		fmt.Println(prompt)
	}
}
