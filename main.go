package main

import (
	"fmt"
	"go_todo/inputGetters"
	"go_todo/todos"
	"log"
	"os"
)

func main() {
	t, err := todos.CreateList()
	if err != nil {
		log.Fatalf("Error creating or loading todo list: %v", err)
	}
	for {
		fmt.Print(`Please choose an option:
[1] Create A Todo
[2] View Todos
[3] Update A Todo
[4] Quit
`)
		i, e := inputGetters.GetInt()
		if e != nil {
			log.Printf("error with input: %v", e)
			continue
		}
		switch i {
		case 1:
			if err := t.AddTodo(); err != nil {
				log.Printf("error adding todo: %v", err)
				continue
			}
		case 2:
			t.ViewTodos("Todos")
		case 3:
			if err := t.UpdateTodo(); err != nil {
				log.Printf("error updating todo: %v", err)
				continue
			}
		case 4:
			if err := t.SaveTodos(); err != nil {
				log.Printf("Error saving todos on program exit: %v", err)
				os.Exit(2) // Terminate the program with a non-zero status code.
			}
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
			continue
		}

	}
}
