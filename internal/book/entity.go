package book

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Pages  int    `json:"pages"`
	Words  int    `json:"words"`
	Author string `json:"author"`
}
