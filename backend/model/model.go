package model

type Pagination struct {
	Page   int
	Limit  int
	Offset int
	Sort   string
	Order  string

	// generated
	PageCap int
	Len     int
	Cap     int
}
