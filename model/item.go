package model

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func NewItem() *Item {
	return nil
}
