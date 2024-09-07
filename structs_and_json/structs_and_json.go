package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/engnrutkarsh/project-api/structs_and_json/note"
)

func userInput(prompt string) (string, error) {
	fmt.Print(prompt)
	input := bufio.NewReader(os.Stdin)
	inputText, err := input.ReadString('\n')
	if err != nil {
		error_statement := errors.New("empty input")
		return "", error_statement
	}
	inputText = strings.TrimSuffix(inputText, "\r\n")
	return inputText, nil
}
func main() {
	title, err := userInput("Note title:")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	content, err := userInput("Note content:")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	note_structure, err := note.New(title, content)
	if err != nil {
		return
	}
	err = note.SaveToFile(note_structure)
	if err != nil {
		return
	}
}
