package main

import (
	"bufio"
	"fmt"
	"os"
	"project/InterfacesAndGenericCode/note"
	"project/InterfacesAndGenericCode/todo"
	"strings"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

func printSomething(value any) {
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ",value)
	// case float64:
	// 	fmt.Println("Float64: ",value)
	// case string:
	// 	fmt.Println("String: ",value)
	// }

	intVal,ok := value.(int)
	if ok {
		fmt.Println("Integer: ",intVal)
		return
	}
	float64Val,ok := value.(float64)
	if ok {
		fmt.Println("Float64: ",float64Val)
		return
	}
	stringVal,ok := value.(string)
	if ok {
		fmt.Println("String: ",stringVal)
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}


func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the node failed")
		return err
	}

	fmt.Println("Saving the node succeeded")
	return nil
}

func getNoteData() (string,string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text,err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r")

	return text
}