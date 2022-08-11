package todo

import "time"

type TodoFormatter struct {
	ID              int        `json:"id"`
	ActivityGroupID int        `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        bool       `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

func FormatTodo(todo Todo) TodoFormatter {
	return TodoFormatter{
		ID:              todo.ID,
		ActivityGroupID: todo.ActivityGroupID,
		Title:           todo.Title,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		CreatedAt:       todo.CreatedAt,
		UpdatedAt:       todo.UpdatedAt,
		DeletedAt:       todo.DeletedAt,
	}
}

func FormatTodos(todos []Todo) []TodoFormatter {
	todosFormatter := []TodoFormatter{}
	for _, todo := range todos {
		todosFormatter = append(todosFormatter, FormatTodo(todo))
	}
	return todosFormatter
}
