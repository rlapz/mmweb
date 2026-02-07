package model

import "time"

type Todo struct {
	Id          int32     `json:"id"`
	IdUser      int32     `json:"id_user"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Flags       int32     `json:"flags"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int32     `json:"created_by"`
}
