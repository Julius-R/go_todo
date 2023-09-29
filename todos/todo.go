package todos

import (
	"errors"
	"fmt"
	"go_todo/helpers"
	"time"
)

type Todo struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"lastUpdated"`
}

func getDescription() (string, error) {
	desc := helpers.GetString()

	if len(desc) < 1 {
		return "", errors.New("empty string")
	}

	return desc, nil
}

func createTodo() (Todo, error) {
	fmt.Println("Enter description")
	desc, err := getDescription()
	if err != nil {
		return Todo{}, err
	}

	return Todo{
		"",
		desc,
		false,
		time.Now().Local(),
		time.Now().Local(),
	}, nil
}

func (t Todo) displayTodo() {
	fmt.Printf(
		`ID: %v
Description: %v
Completed: %v
Created: %v
Last Modified: %v

`, t.Id, t.Description, t.IsCompleted, t.Created.Format(time.RFC822), t.LastUpdated.Format(time.RFC822))
}

func loopTodos(tl *TodoList, loopMsg string) {
	fmt.Printf("%v:\n\n", loopMsg)
	for _, t := range tl.List {
		t.displayTodo()
	}
}

func editTodo(t *Todo, editType int) *Todo {
	switch editType {
	case 1:
		{
			desc, err := getDescription()
			if err != nil {
				fmt.Println(err)
				return t
			}

			t.Description = desc
			t.LastUpdated = time.Now().Local()
			fmt.Println("Description updated")
		}
	case 2:
		{
			if t.IsCompleted {
				fmt.Println("Already completed")
			}
			t.IsCompleted = true
			t.LastUpdated = time.Now().Local()
			fmt.Println("Marked completed")
		}
	case 3:
		fmt.Println("Returning")
	default:
		fmt.Println("Invalid entry")
	}
	return t
}
