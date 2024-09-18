package task

import "time"

type Status string

const (
	STATUS_TODO        = "todo"
	STATUS_DONE        = "done"
	STATUS_IN_PROGRESS = "in-progress"
)

type Task struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
