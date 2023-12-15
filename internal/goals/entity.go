package goals

type Goal struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Quantity int    `json:"quantity"`
	Days     int    `json:"days"`
	UserID   string `json:"user_id"`
}
