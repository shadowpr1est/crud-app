package main

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
}
