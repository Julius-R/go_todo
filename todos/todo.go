package todos

import (
	"errors"
	"fmt"
	"time"
)

type todo struct {
	id, description      string
	isCompleted          bool
	created, lastUpdated time.Time
}

func getDescription() (string, error) {
	desc := getString()

	if len(desc) < 1 {
		return "", errors.New("empty string")
	}

	return desc, nil
}

func createTodo() (todo, error) {
	fmt.Println("Enter description")
	desc, err := getDescription()
	if err != nil {
		return todo{}, err
	}

	return todo{
		"",
		desc,
		false,
		time.Now().Local(),
		time.Now().Local(),
	}, nil
}

func (t todo) displayTodo() {
	fmt.Printf(`
ID: %v
Description: %v
Completed: %v
Created: %v
Last Modified: %v
`, t.id, t.description, t.isCompleted, t.created.Format(time.RFC822), t.lastUpdated.Format(time.RFC822))
}
