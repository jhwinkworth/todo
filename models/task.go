package models

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}
