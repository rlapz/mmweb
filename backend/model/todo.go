package model

type Todo struct {
	Id        int32  `json:"id"`
	IdUser    int32  `json:"id_user"`
	Label     string `json:"label" validate:"required"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt *int64 `json:"updated_at"`
}

type TodoItem struct {
	Id          int32  `json:"id"`
	IdTodo      int32  `json:"id_todo" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Flags       int32  `json:"flags" validate:"required"`
	CreatedAt   int64  `json:""`
	UpdatedAt   *int64 `json:""`
}
