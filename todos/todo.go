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

/*
func createTodo() (*todos, error) {
	fmt.Println("Enter description")
	value := getString()
	if len(value) == 0 {
		return &todos{}, errors.New("empty input provided")
	}

	return &todos{
		value,
		false,
		time.Now().Local(),
		time.Now().Local(),
	}, nil
}

func (t todos) displayTodo() {
	fmt.Printf(`
Description: %v
Completed: %v
Created: %v
Last Updated: %v

`, t.description, t.isCompleted, t.created.Format(time.DateTime), t.lastUpdated.Format(time.DateTime))
}

func updateTodo(tl []*todos) {
	options := []string{"Update Description", "Update Completion Status", "Delete", "Cancel"}
	fmt.Println("Todos:")
	for i, v := range tl {
		fmt.Printf("ID: %v", i+1)
		v.displayTodo()
	}
	fmt.Println("Select a todos")
	choice := getInt() - 1

	if choice >= len(tl) {
		fmt.Println("provided value does not exist")
		return
	}

	tod := tl[choice]
	tod.displayTodo()
	fmt.Println("Options:")
	for i, v := range options {
		fmt.Printf("[%v] %v\n", i, v)
	}
	fmt.Println("\nPlease chose edit option")
	updateOption := getInt()

	if updateOption >= len(options) {
		fmt.Println("provided value does not exist")
		return
	}

	switch updateOption {
	case 0:
		fmt.Println("Provide new description")
		desc := getString()
		if len(desc) == 0 {
			fmt.Println("empty value provided, update cancelled")
		} else {
			tod.description = desc
			tod.lastUpdated = time.Now().Local()
		}
	}
}
*/
