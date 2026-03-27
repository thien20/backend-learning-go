// Package model contains domain types and shared sentinel errors.
//
// Keep models focused on business meaning, not HTTP or database concerns.
package model

import (
	"errors"
	"time"
)

var (
	ErrItemNotFound   = errors.New("item not found")
	ErrInvalidItemID  = errors.New("invalid item id")
	ErrNotImplemented = errors.New("student task not implemented")
)

type Item struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}


// TODO: Decide what makes an Item valid in this stage.
// Examples to consider:
// - Is an empty ID ever allowed? - OK
// - Should names be trimmed before storage?
// - Should timestamps be set by the service or the DB?
