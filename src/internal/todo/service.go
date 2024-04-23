package todo

import "time"

func GetAllTodos() (todos []Todo, err error) {
	todos = BuildTodos()
	return
}

func GetTodoById(id int64) (todo Todo, err error) {
	todos := BuildTodos()
	for _, t := range todos {
		if t.Id == id {
			todo = t
			return
		}
	}

	return
}

func BuildTodos() (todos []Todo) {
	var items []Item
	items = append(items, Item{Description: "Daily"})
	items = append(items, Item{Description: "Coding"})
	items = append(items, Item{Description: "Tests"})
	todos = append(todos, Todo{Id: 1, Description: "Work", Items: items, FinishDate: time.Now()})
	return
}
