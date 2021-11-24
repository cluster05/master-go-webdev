package api

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
	ID    string `json:"id"`
}
