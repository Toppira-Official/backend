package model

type Message struct {
	Title    string  `json:"title"`
	Body     string  `json:"body"`
	ImageURL *string `json:"image_url,omitempty"`
}
