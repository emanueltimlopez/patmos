package book

import "github.com/google/uuid"

type Book struct {
	ID   uuid.UUID
	Name string
}
