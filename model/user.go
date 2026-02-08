package model

import (
	"time"
)

type User struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte
	Flags     int32      `json:"flags"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy int64      `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy int64      `json:"updated_by"`
}
