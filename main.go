package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, body := getNoteData()

	userNote, err := note.New(title, body)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	userNote.PrintNote()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	body := getUserInput("Note body: ")

	return title, body
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}
