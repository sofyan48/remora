package entity

import "time"

// PlaybookResponse ...
type PlaybookResponse struct {
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
}
