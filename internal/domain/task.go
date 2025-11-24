package domain

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	State string `json:"state" binding:"required"`
}
