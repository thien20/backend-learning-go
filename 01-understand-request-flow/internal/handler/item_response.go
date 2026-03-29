package handler

import (
	"time"

	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/model"
)

type itemResponse struct {
	ID        string    `json:"id,omitempty`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetItemResponse(item model.Item) itemResponse {
	return itemResponse{
		ID:        item.ID,
		Name:      item.Name,
		UpdatedAt: item.UpdatedAt,
	}
}
