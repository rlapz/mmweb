package model

type Todo struct {
	Id        int32  `json:"id"`
	IdUser    int32  `json:"id_user"`
	Label     string `json:"label"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt *int64 `json:"updated_at"`
}

type TodoItem struct {
	Id          int32  `json:"id"`
	IdTodo      int32  `json:"id_todo"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Flags       int32  `json:"flags"`
	CreatedAt   int64  `json:""`
	UpdatedAt   *int64 `json:""`
}
