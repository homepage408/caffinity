package models

type CaffeResponse struct {
	ID          int32    `json:"id"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Address     *string  `json:"address,omitempty"`
	Description *string  `json:"description,omitempty"`
	Rating      *float64 `json:"rating,omitempty"`
	CreatedAt   *string  `json:"created_at,omitempty"`
}
