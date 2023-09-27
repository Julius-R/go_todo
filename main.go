package main

import "go_todo/todos"

func main() {
	t := todos.CreateList()
	t.AddTodo()
	t.AddTodo()
	t.ViewTodos()
}
