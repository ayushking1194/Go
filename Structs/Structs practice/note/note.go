package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title  string `json:"title"` // Title of the note
	Content string `json:"content"` // Content of the note
	CreatedAt time.Time `json:"created_at"` // Timestamp when the note was created

}
func (note Note) Display()  {
	fmt.Printf("Your note title is %v and content is %v", note.Title, note.Content)
}

func (note Note) Save() error{
	fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		Title:   title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}
