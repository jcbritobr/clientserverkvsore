package model

const (
	ErrNotFound   = "not found"
	ErrEmptyQuery = "empty query"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type UuidMessage struct {
	Uuid string `json:"uuid"`
}

type DoneundoneResult struct {
	Done bool `json:"done"`
}
