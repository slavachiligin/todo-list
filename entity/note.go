package entity

import "time"

type Note struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	Time        time.Time
	Done        bool
}
