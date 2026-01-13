package domain

import (
	"errors"
	"time"
)

// Comment errors
var (
	ErrCommentBodyEmpty = errors.New("comment body cannot be empty")
	ErrNotCommentAuthor = errors.New("user is not the comment author")
)

// Comment represents a comment on a task.
type Comment struct {
	ID        string
	TaskID    string
	AuthorID  string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewComment creates a new comment with validation.
func NewComment(id, taskID, authorID, body string) (*Comment, error) {
	if body == "" {
		return nil, ErrCommentBodyEmpty
	}
	now := time.Now()
	return &Comment{
		ID:        id,
		TaskID:    taskID,
		AuthorID:  authorID,
		Body:      body,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// IsAuthor checks if the given user is the comment author.
func (c *Comment) IsAuthor(userID string) bool {
	return c.AuthorID == userID
}

// CanEdit checks if the given user can edit this comment.
// Only the author can edit their own comments.
func (c *Comment) CanEdit(userID string) bool {
	return c.IsAuthor(userID)
}

// CanDelete checks if the given user can delete this comment.
// Only the author can delete their own comments.
func (c *Comment) CanDelete(userID string) bool {
	return c.IsAuthor(userID)
}

// Update updates the comment body if the user is the author.
func (c *Comment) Update(userID, newBody string) error {
	if !c.IsAuthor(userID) {
		return ErrNotCommentAuthor
	}
	if newBody == "" {
		return ErrCommentBodyEmpty
	}
	c.Body = newBody
	c.UpdatedAt = time.Now()
	return nil
}
