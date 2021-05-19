package pagination

import "gorm.io/gorm"

type Page struct {
	URL    string
	Number int
}

type ViewData struct {
	HasPages   bool
	Next       Page
	HasNext    bool
	Prev       Page
	HasPrev    bool
	Current    Page
	TotalCount int64
	TotalPage  int
}

type Pagination struct {
	BaseURL string
	PerPage int
	Page    int
	Count   int64
	db      *gorm.DB
}
