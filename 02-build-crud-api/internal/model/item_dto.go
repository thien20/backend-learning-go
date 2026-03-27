package model

// CreateItemRequest and UpdateItemRequest are transport shapes.
//
// They are separate from Item so the student can practice DTO-to-domain mapping.
// Do not assume every HTTP field should exist on the domain model forever.
type CreateItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateItemRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type ItemResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ListItemsResponse struct {
	Items []ItemResponse `json:"items"`
}

// TODO: Add request-specific validation rules here or in the service layer.
// The important part is being explicit about where validation belongs.
