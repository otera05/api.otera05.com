package entity

import "time"

type TodoID string

func (t TodoID) String() string { return string(t) }

type Todo struct {
	ID      TodoID    `json:"id"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	DueDate time.Time `json:"due_date"`
	Done    bool      `json:"done"`
}

type TodoList []Todo
