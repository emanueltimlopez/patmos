package book

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Pages  int    `json:"pages"`
	Words  int    `json:"words"`
	Author string `json:"author"`
	Image  string `json:"image"`
	Isbn   string `json:"isbn"`
}

type UserBooksRelation struct {
	UserID string `json:"user_id"`
	BookID string `json:"book_id"`
	Folder string `json:"folder"`
}
