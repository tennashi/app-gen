package domain

import (
	"errors"
	"time"
)

// Tag errors
var (
	ErrTagNameEmpty     = errors.New("tag name cannot be empty")
	ErrTagAlreadyExists = errors.New("tag with this name already exists in project")
	ErrInvalidTagColor  = errors.New("invalid tag color format")
)

// Tag represents a label that can be attached to tasks.
type Tag struct {
	ID        string
	ProjectID string
	Name      string
	Color     string
	CreatedAt time.Time
}

// NewTag creates a new tag with validation.
func NewTag(id, projectID, name, color string) (*Tag, error) {
	if name == "" {
		return nil, ErrTagNameEmpty
	}
	if color != "" && !isValidHexColor(color) {
		return nil, ErrInvalidTagColor
	}
	return &Tag{
		ID:        id,
		ProjectID: projectID,
		Name:      name,
		Color:     color,
		CreatedAt: time.Now(),
	}, nil
}

func isValidHexColor(color string) bool {
	if len(color) != 7 || color[0] != '#' {
		return false
	}
	for _, c := range color[1:] {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// TaskTag represents the many-to-many relationship between Task and Tag.
type TaskTag struct {
	TaskID    string
	TagID     string
	CreatedAt time.Time
}

// NewTaskTag creates a new task-tag association.
func NewTaskTag(taskID, tagID string) *TaskTag {
	return &TaskTag{
		TaskID:    taskID,
		TagID:     tagID,
		CreatedAt: time.Now(),
	}
}
