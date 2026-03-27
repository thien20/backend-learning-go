package model

import (
	"errors"
	"time"
)

var (
	ErrItemNotFound   = errors.New("item not found")
	ErrInvalidItemID  = errors.New("invalid item id")
	ErrInvalidItem    = errors.New("invalid item")
	ErrNotImplemented = errors.New("student task not implemented")
)

type Item struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TODO: Decide which invariants belong to the domain model and which belong to
// the transport layer. This becomes important once DTOs and persistence diverge.
