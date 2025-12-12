package types

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Visit_URL   string `json:"visit_url"`
	Source_URL  string `json:"source_url"`
	Youtube_URL string `json:"youtube_url"`
	Image_URL   string `json:"image_url"`
}
