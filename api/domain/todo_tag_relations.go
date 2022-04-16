package domain

type TodoTagRelation struct {
	ID     int `json:"id"`
	TodoID int `json:"todo_id"`
	TagID  int `json:"tag_id"`
}

type TodoTagRelations []TodoTagRelation
