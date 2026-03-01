package model

const (
	TODO_ITEM_STATUS_PENDING   = 1 << 1
	TODO_ITEM_STATUS_POSTPONED = 1 << 2
	TODO_ITEM_STATUS_CANCELLED = 1 << 3
	TODO_ITEM_STATUS_DONE      = 1 << 4
)

type Todo struct {
	Id        int32  `json:"id"`
	IdUser    int32  `json:"id_user"`
	Label     string `json:"label" validate:"required"`
	IsActive  bool   `json:"is_active"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt *int64 `json:"updated_at"`
}

type TodoItem struct {
	Id          int32  `json:"id" validate:"required"`
	IdTodo      int32  `json:"id_todo" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Deadline    *int64 `json:"deadline"`
	Status      int32  `json:"status"`
	CreatedAt   int64  `json:""`
	UpdatedAt   *int64 `json:""`
}
