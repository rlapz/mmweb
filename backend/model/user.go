package model

type User struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Flags     int32  `json:"flags"`
	CreatedAt int64  `json:"created_at"`
	CreatedBy int32  `json:"created_by"`
	UpdatedAt *int64 `json:"updated_at"`
	UpdatedBy int32  `json:"updated_by"`
}
