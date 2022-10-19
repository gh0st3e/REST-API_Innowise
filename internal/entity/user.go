package entity

import (
	uuid2 "github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid2.UUID `json:"id"`
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Email     string     `json:"email"`
	Age       uint       `json:"age"`
	Created   time.Time  `json:"created"`
}
