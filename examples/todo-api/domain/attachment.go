package domain

import (
	"errors"
	"path/filepath"
	"time"
)

// Attachment errors
var (
	ErrAttachmentFileNameEmpty = errors.New("file name cannot be empty")
	ErrAttachmentFileTooLarge  = errors.New("file size exceeds maximum allowed")
	ErrAttachmentInvalidType   = errors.New("file type not allowed")
	ErrNotAttachmentUploader   = errors.New("user is not the attachment uploader")
)

// MaxAttachmentSize is the maximum allowed file size (10MB).
const MaxAttachmentSize = 10 * 1024 * 1024

// AllowedMimeTypes defines the allowed file types.
var AllowedMimeTypes = map[string]bool{
	"image/jpeg":      true,
	"image/png":       true,
	"image/gif":       true,
	"application/pdf": true,
	"text/plain":      true,
}

// AttachmentTargetType represents the type of entity an attachment belongs to.
type AttachmentTargetType string

const (
	AttachmentTargetTask    AttachmentTargetType = "task"
	AttachmentTargetComment AttachmentTargetType = "comment"
)

// Attachment represents a file attached to a task or comment.
type Attachment struct {
	ID         string
	TargetType AttachmentTargetType
	TargetID   string
	UploaderID string
	FileName   string
	FileSize   int64
	MimeType   string
	URL        string
	CreatedAt  time.Time
}

// NewAttachment creates a new attachment with validation.
func NewAttachment(id string, targetType AttachmentTargetType, targetID, uploaderID, fileName string, fileSize int64, mimeType, url string) (*Attachment, error) {
	if fileName == "" {
		return nil, ErrAttachmentFileNameEmpty
	}
	if fileSize > MaxAttachmentSize {
		return nil, ErrAttachmentFileTooLarge
	}
	if !AllowedMimeTypes[mimeType] {
		return nil, ErrAttachmentInvalidType
	}
	return &Attachment{
		ID:         id,
		TargetType: targetType,
		TargetID:   targetID,
		UploaderID: uploaderID,
		FileName:   fileName,
		FileSize:   fileSize,
		MimeType:   mimeType,
		URL:        url,
		CreatedAt:  time.Now(),
	}, nil
}

// IsUploader checks if the given user uploaded this attachment.
func (a *Attachment) IsUploader(userID string) bool {
	return a.UploaderID == userID
}

// CanDelete checks if the given user can delete this attachment.
func (a *Attachment) CanDelete(userID string) bool {
	return a.IsUploader(userID)
}

// Extension returns the file extension.
func (a *Attachment) Extension() string {
	return filepath.Ext(a.FileName)
}

// IsImage returns true if the attachment is an image.
func (a *Attachment) IsImage() bool {
	switch a.MimeType {
	case "image/jpeg", "image/png", "image/gif":
		return true
	}
	return false
}
