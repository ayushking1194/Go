package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("\nNote saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter note title: ")

	content := getUserInput("Enter note content: ")

	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	input = strings.TrimSpace(input) // Remove any trailing newline or spaces
	if input == "" {
		fmt.Println("Input cannot be empty. Please try again.")
		return getUserInput(prompt) // Recursive call to get valid input
	}
	return input
}