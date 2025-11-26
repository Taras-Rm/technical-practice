package domain

import "time"

// type TaskStatus string

// const (
// 	StatusTodo       TaskStatus = "todo"
// 	StatusInProgress TaskStatus = "in-progrress"
// 	StatusDone       TaskStatus = "done"
// )

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
