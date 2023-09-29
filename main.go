package main

import (
	"fmt"
	"go_todo/helpers"
	"go_todo/todos"
)

func main() {
	t := todos.CreateList()
	for {
		fmt.Print(`Please choose an option:
[1] Create A Todo
[2] View Todos
[3] Update A Todo
[4] Quit
`)
		switch helpers.GetInt() {
		case 1:
			t.AddTodo()
		case 2:
			t.ViewTodos()
		case 3:
			t.UpdateTodo()
		case 4:
			fallthrough
		default:
			return
		}
		t.SaveTodos()
	}
}
