package model

type Course struct {
	ID          string  `json:"ID"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
