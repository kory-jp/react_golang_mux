package domain

type Tag struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
	Label string `json:"label"`
}

type Tags []Tag
