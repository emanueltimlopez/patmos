package book

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Pages  int    `json:"pages"`
	Words  int    `json:"words"`
	Author string `json:"author"`
}

type BookFromAPI struct {
	Cover       int      `json:"cover_i"`
	FullText    bool     `json:"has_fulltext"`
	Edition     int      `json:"edition_count"`
	Title       string   `json:"title"`
	Author      []string `json:"author_name"`
	PublishYear int      `json:"first_publish_year"`
	Key         string   `json:"key"`
	Ia          []string `json:"ia"`
	AuthorKey   []string `json:"author_key"`
	PublicScan  bool     `json:"public_scan_b"`
}

type BooksFromAPI struct {
	Start    int           `json:"start"`
	NumFound int           `json:"num_found"`
	Docs     []BookFromAPI `json:"docs"`
}
