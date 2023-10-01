package todos

import (
	"errors"
	"fmt"
	"go_todo/inputGetters"
	"time"
)

type Todo struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"lastUpdated"`
}

func createTodo() (Todo, error) {
	fmt.Println("Enter description")
	desc, err := inputGetters.GetString()
	if err != nil {
		return Todo{}, fmt.Errorf("error with provided input, %v", err)
	}

	if len(desc) < 1 {
		return Todo{}, errors.New("description must not be empty")
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

func editTodo(t *Todo, editType int) (*Todo, error) {
	switch editType {
	case 1:
		{
			desc, err := inputGetters.GetString()
			if err != nil {
				return t, fmt.Errorf("error with provided input: %v", err)
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
	return t, nil
}
