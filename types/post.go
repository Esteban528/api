package types

type Post struct {
	ID          int    `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Date        string `json:"date"`
}

