package webhooks

import (
	"bytes"
	"encoding/json"
	"estebandev_api/db"
	"estebandev_api/events"
	"fmt"
	"net/http"
	"os"
)

type SubCreate[T any] struct{}

func (SubCreate[T]) Update(value T) {
	post := any(value).(*db.Post)
	url := fmt.Sprintf("https://%s/blog/post/%d", os.Getenv("TEXT_SERVERDOMAIN"), +post.ID)
	Send(post.Title, post.Description, "A new post got released", url, "By " + post.Author)
}

func LoadEvents() {
	events.Subscribe("create_post", SubCreate[*db.Post]{})
}

type Footer struct {
	Text string `json:"text"`
}

type embedT struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Color       int    `json:"color"`
	Footer      Footer `json:"footer"`
}

type payloadT struct {
	Embeds    []embedT `json:"embeds"`
	Username  string   `json:"username"`
	Content   string   `json:"content"`
	AvatarURL string   `json:"avatar_url"`
}

func Send(title, content, message, embedUrl, footText string) {
	url := os.Getenv("WEBHOOKS_DISCORD_URL")

	embed := embedT{
		Title:       title,
		Description: content,
		URL:         embedUrl,
		Color:       0x00AED9,
		Footer:      Footer{Text: footText},
	}

	payload := payloadT{
		Embeds:    []embedT{embed},
		Username:  os.Getenv("WEBHOOKS_DISCORD_USERNAME"),
		Content:   message,
		AvatarURL: os.Getenv("WEBHOOKS_DISCORD_AVATAR_URL"),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error marshaling:", err)
		return
	}

	http.Post(url, "application/json", bytes.NewBuffer(data))
}
