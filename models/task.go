package models

import (
	"context"
	"time"
)

type TaskStatus string
type Task struct {
	ID         string             `json:"id"`
	Status     TaskStatus         `json:"status"`
	Result     *string            `json:"result,omitempty"`
	Duration   string             `json:"duration"`
	CreatedAt  time.Time          `json:"created_at"`
	StartedAt  *time.Time         `json:"started_at,omitempty"`
	EndedAt    *time.Time         `json:"ended_at,omitempty"`
	Error      *string            `json:"error,omitempty"`
	CancelFunc context.CancelFunc `json:"-"`
}

const (
	Waiting TaskStatus = "waiting" // можно оставить для будущей реализации очереди задач
	Running TaskStatus = "running"
	Success TaskStatus = "success"
	Failed  TaskStatus = "error"            // можно оставить на будущие ошибки
	Deleted TaskStatus = "deleted/canceled" // можно оставить для будущей истории задач
)
