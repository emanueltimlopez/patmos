package user

type User struct {
	ID   string `json:"id"`
	Book string `json:"book,omitempty"`
}
