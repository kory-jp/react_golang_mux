package domain

type Todo struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Todos []Todo
