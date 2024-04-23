package todo

import "time"

type Todo struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Items       []Item    `json:"items"`
	FinishDate  time.Time `json:"finishDate"`
}

type Item struct {
	Description string `json:"description"`
}
