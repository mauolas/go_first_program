package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, body string) (*Note, error) {

	if title == "" || body == "" {
		return &Note{}, errors.New("title and body are required")
	}

	return &Note{
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonData, err := json.Marshal(n)

	if err != nil {
		println("Error marshalling note:", err)
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func (n Note) PrintNote() {
	println("Title:", n.Title)
	println("Body:", n.Body)

	createdAtFormatted := n.CreatedAt.Format("January 2, 2006 at 3:04pm")
	println("Created at:", createdAtFormatted)
}
