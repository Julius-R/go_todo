package todos

import (
	"fmt"
	"math/rand"
	"strconv"
)

type TodoList struct {
	list map[string]todo
}

func CreateList() TodoList {
	return TodoList{
		make(map[string]todo),
	}
}

func (tl *TodoList) AddTodo() {
	t, e := createTodo()
	if e != nil {
		fmt.Println(e)
		return
	}
	v := "G-" + strconv.Itoa(rand.Intn(500))
	t.id = v
	tl.list[v] = t
	fmt.Print("Todo created\n\n")
}

func (tl *TodoList) ViewTodos() {
	fmt.Println("Todos:")
	for _, t := range tl.list {
		t.displayTodo()
	}
}
