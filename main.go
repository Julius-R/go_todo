package main

import (
	"fmt"
	"os"
)

func main() {
	var tl []*todo
	t, err := createTodo()

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	tl = append(tl, t)

	updateTodo(tl)
	tl[0].displayTodo()
}
