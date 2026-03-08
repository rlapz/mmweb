package model

type Pagination struct {
	Page      int
	PageCap   int
	ListLimit int
	ListLen   int
	ListCap   int
	Sort      string
	Order     string
}
