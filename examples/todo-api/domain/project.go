// Package domain contains the core business entities and logic.
package domain

import (
	"errors"
	"slices"
	"time"
)

// Project errors
var (
	ErrNotProjectOwner    = errors.New("user is not the project owner")
	ErrProjectNameEmpty   = errors.New("project name cannot be empty")
	ErrUserAlreadyMember  = errors.New("user is already a member")
	ErrUserNotMember      = errors.New("user is not a member")
	ErrCannotRemoveOwner  = errors.New("cannot remove owner from project")
)

// Project represents a project that contains tasks.
type Project struct {
	ID          string
	Name        string
	Description string
	OwnerID     string
	MemberIDs   []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewProject creates a new project with the owner as the first member.
func NewProject(id, name, description, ownerID string) (*Project, error) {
	if name == "" {
		return nil, ErrProjectNameEmpty
	}
	now := time.Now()
	return &Project{
		ID:          id,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		MemberIDs:   []string{ownerID},
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// IsOwner checks if the given user is the project owner.
func (p *Project) IsOwner(userID string) bool {
	return p.OwnerID == userID
}

// IsMember checks if the given user is a project member.
func (p *Project) IsMember(userID string) bool {
	return slices.Contains(p.MemberIDs, userID)
}

// AddMember adds a user to the project.
func (p *Project) AddMember(userID string) error {
	if p.IsMember(userID) {
		return ErrUserAlreadyMember
	}
	p.MemberIDs = append(p.MemberIDs, userID)
	p.UpdatedAt = time.Now()
	return nil
}

// RemoveMember removes a user from the project.
func (p *Project) RemoveMember(userID string) error {
	if userID == p.OwnerID {
		return ErrCannotRemoveOwner
	}
	if !p.IsMember(userID) {
		return ErrUserNotMember
	}
	newMembers := make([]string, 0, len(p.MemberIDs)-1)
	for _, id := range p.MemberIDs {
		if id != userID {
			newMembers = append(newMembers, id)
		}
	}
	p.MemberIDs = newMembers
	p.UpdatedAt = time.Now()
	return nil
}

// CanDelete checks if the given user can delete this project.
func (p *Project) CanDelete(userID string) bool {
	return p.IsOwner(userID)
}
