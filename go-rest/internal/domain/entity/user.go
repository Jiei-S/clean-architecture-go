package entity

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Age       int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
