package main

import (
	"bufio"
	"example.com/notesCLI/todo"
	"fmt"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}
type outputtable interface {
	Saver
	Display()
	HandleErrorPrompt(err error, prompt string)
}

func main() {
	title := getInput("Provide Title: ")
	content := getInput("Provide content: ")
	todoTest := getInput("Get todo text:")

	myTodo, err := todo.NewTodo(todoTest)
	myTodo.HandleErrorPrompt(err, "outputtable error")
	outputData(myTodo)
	note := NewNote(title, content)
	outputData(note)
	result := add(2, 5)
	fmt.Println(result)
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

func outputData(data outputtable) {
	data.Display()
	SaveData(data)
}

func printSomething(value any) { // can be interface{} instead any
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case string:
		fmt.Println("String: ", value)
	}
	fmt.Println(value)

	// extracting type information from values

	floatval, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", floatval)
	}

	strinval, ok := value.(string)
	if ok {
		fmt.Println("String: ", strinval)
	}
}

// GENERICS

func add[T int | float64 | string](a, b T) T {
	return a + b
}
