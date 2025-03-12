package entities

import "time"

type Category struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
