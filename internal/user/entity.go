package user

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Book string    `json:"book"`
}
