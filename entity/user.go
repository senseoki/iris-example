package entity

import "time"

// User is ...
type User struct {
	ID        uint64 `gorm:"primary_key"`
	Email     string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
