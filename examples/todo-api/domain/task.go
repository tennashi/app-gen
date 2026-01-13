package domain

import (
	"errors"
	"time"
)

// TaskStatus represents the status of a task.
type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

// Task errors
var (
	ErrInvalidStatusTransition = errors.New("invalid status transition")
	ErrTaskAlreadyCompleted    = errors.New("task is already completed")
	ErrTaskOverdue             = errors.New("task is overdue")
)

// Task represents a task within a project.
type Task struct {
	ID           string
	ProjectID    string
	ParentTaskID string
	ReporterID   string
	AssigneeID   string
	Title        string
	Description  string
	Status       TaskStatus
	DueDate      *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CanTransitionTo checks if the task can transition to the given status.
// Valid transitions:
//   - todo -> in_progress
//   - in_progress -> todo (revert)
//   - in_progress -> done
//   - done -> in_progress (reopen)
func (t *Task) CanTransitionTo(newStatus TaskStatus) bool {
	switch t.Status {
	case TaskStatusTodo:
		return newStatus == TaskStatusInProgress
	case TaskStatusInProgress:
		return newStatus == TaskStatusTodo || newStatus == TaskStatusDone
	case TaskStatusDone:
		return newStatus == TaskStatusInProgress
	}
	return false
}

// TransitionTo changes the task status if the transition is valid.
func (t *Task) TransitionTo(newStatus TaskStatus) error {
	if !t.CanTransitionTo(newStatus) {
		return ErrInvalidStatusTransition
	}
	t.Status = newStatus
	t.UpdatedAt = time.Now()
	return nil
}

// IsOverdue returns true if the task has a due date that has passed.
func (t *Task) IsOverdue() bool {
	if t.DueDate == nil {
		return false
	}
	return time.Now().After(*t.DueDate) && t.Status != TaskStatusDone
}

// Complete marks the task as done.
func (t *Task) Complete() error {
	if t.Status == TaskStatusDone {
		return ErrTaskAlreadyCompleted
	}
	// Allow direct completion from any status
	t.Status = TaskStatusDone
	t.UpdatedAt = time.Now()
	return nil
}
