package todos

import (
	"encoding/json"
	"fmt"
	"go_todo/helpers"
	"math/rand"
	"os"
	"strconv"
)

type TodoList struct {
	List map[string]*Todo `json:"list"`
}

func CreateList() TodoList {
	t := TodoList{
		make(map[string]*Todo),
	}
	file, err := os.ReadFile("./todos.json")
	if err != nil {
		return t
	}

	f := struct {
		Todos []*Todo `json:"todos"`
	}{}

	err = json.Unmarshal(file, &f)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	for _, v := range f.Todos {
		t.List[v.Id] = v
	}
	return t
}

func (tl *TodoList) AddTodo() {
	t, e := createTodo()
	if e != nil {
		fmt.Printf("error occured: %v\n", e)
		return
	}
	v := "G-" + strconv.Itoa(rand.Intn(500))
	t.Id = v
	tl.List[v] = &t
	fmt.Print("Todo created\n\n")
}

func (tl *TodoList) ViewTodos() {
	loopTodos(tl, "Todos")
}

func (tl *TodoList) UpdateTodo() {
	loopTodos(tl, "Please select Todo")
	fmt.Print("Enter the id of your Todo G-")
	todoId := "G-"
	usrInt := helpers.GetInt()
	todoId = todoId + strconv.Itoa(usrInt)
	// Validate that Todo exists
	todo, ok := tl.List[todoId]
	if !ok {
		fmt.Println("invalid id provided")
		return
	}
	//	choose update type
	fmt.Print(`
Please select from the following:
[1] Update Description
[2] Mark Complete
[3] Cancel
`)
	userInt := helpers.GetInt()
	todo = editTodo(todo, userInt)
}

func (tl *TodoList) makeSlice() []Todo {
	var sl []Todo
	for _, v := range tl.List {
		sl = append(sl, *v)
	}
	return sl
}

func (tl *TodoList) SaveTodos() {
	file, err := json.MarshalIndent(struct {
		Todos []Todo `json:"todos"`
	}{tl.makeSlice()}, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = os.WriteFile("./todos.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
