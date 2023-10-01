package todos

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_todo/inputGetters"
	"os"
	"strconv"
)

type Todos struct {
	List map[string]*Todo `json:"list"`
}

func CreateList() (*Todos, error) {
	t := &Todos{
		make(map[string]*Todo),
	}
	file, err := os.ReadFile("./todos.json")
	if err != nil {
		if os.IsNotExist(err) {
			emptyJSON := []byte(`{"todos": []}`)
			if err := os.WriteFile("./todos.json", emptyJSON, 0644); err != nil {
				return t, fmt.Errorf("error creating file: %w", err)
			}
			return t, nil
		}
		return t, fmt.Errorf("error occured while reading file: %w", err)
	}

	f := struct {
		Todos []*Todo `json:"todos"`
	}{}

	err = json.Unmarshal(file, &f)

	if err != nil {
		return t, fmt.Errorf("error occurred while unmarshaling JSON: %w", err)
	}

	for _, v := range f.Todos {
		t.List[v.Id] = v
	}
	return t, nil
}

func (tl *Todos) AddTodo() error {
	t, e := createTodo()
	if e != nil {
		return fmt.Errorf("error creating todo: %v", e)
	}
	v := "G-" + strconv.Itoa(len(tl.List)+1)
	t.Id = v
	tl.List[v] = &t
	fmt.Print("Todo created\n\n")
	return nil
}

func (tl *Todos) ViewTodos(loopMsg string) {
	fmt.Printf("%v:\n\n", loopMsg)
	for _, t := range tl.List {
		t.displayTodo()
	}
}

func (tl *Todos) UpdateTodo() error {
	tl.ViewTodos("Please select Todo")
	fmt.Print("Enter the id of your Todo G-")
	todoId := "G-"
	usrInt, err := inputGetters.GetInt()
	if err != nil {
		return err
	}
	todoId = todoId + strconv.Itoa(usrInt)
	todo, ok := tl.List[todoId]
	if !ok {
		return errors.New("invalid id provided")
	}
	fmt.Print(`
Please select from the following:
[1] Update Description
[2] Mark Complete
[3] Cancel
`)
	userInt, err := inputGetters.GetInt()
	if err != nil {
		return err
	}
	todo, err = editTodo(todo, userInt)
	return nil
}

func (tl *Todos) makeSlice() []Todo {
	var sl []Todo
	for _, v := range tl.List {
		sl = append(sl, *v)
	}
	return sl
}

func (tl *Todos) SaveTodos() error {
	file, err := json.MarshalIndent(struct {
		Todos []Todo `json:"todos"`
	}{tl.makeSlice()}, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("./todos.json", file, 0644)

}
