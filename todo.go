package main

import (
	"errors"
	"fmt"
	"time"
)

type todo struct {
	description          string
	isCompleted          bool
	created, lastUpdated time.Time
}

func createTodo() (*todo, error) {
	fmt.Println("Enter description")
	value := getString()
	if len(value) == 0 {
		return &todo{}, errors.New("empty input provided")
	}

	return &todo{
		value,
		false,
		time.Now().Local(),
		time.Now().Local(),
	}, nil
}

func (t todo) displayTodo() {
	fmt.Printf(`
Description: %v
Completed: %v
Created: %v
Last Updated: %v

`, t.description, t.isCompleted, t.created.Format(time.DateTime), t.lastUpdated.Format(time.DateTime))
}

func updateTodo(tl []*todo) {
	options := []string{"Update Description", "Update Completion Status", "Delete", "Cancel"}
	fmt.Println("Todos:")
	for i, v := range tl {
		fmt.Printf("ID: %v", i+1)
		v.displayTodo()
	}
	fmt.Println("Select a todo")
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
