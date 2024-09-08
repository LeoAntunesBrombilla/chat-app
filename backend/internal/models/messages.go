package models

type Message struct {
	ID        string `json:"id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Visible   bool   `json:"visible"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
