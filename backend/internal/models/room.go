package models

type Room struct {
	Message   []Message `json:"messages"`
	Name      string    `json:"name"`
	Topic     string    `json:"topic"`
	Author    string    `json:"author"`
	Active    bool      `json:"active"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
}
