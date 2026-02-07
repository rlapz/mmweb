package model

import (
	"time"
)

type User struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Flags     int32     `json:"flags"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int64     `json:"created_by"`
}

type UserDetail struct {
	Id        int32  `json:"id"`
	IdUser    int32  `json:"id_user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte

	CreatedAt time.Time  `json:"created_at"`
	CreatedBy int64      `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy int64      `json:"updated_by"`
}
