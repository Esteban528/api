package types

type Resource struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Image_URL   string `json:"image_url"`
}
