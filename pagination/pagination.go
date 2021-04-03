package pagination

type PaginationMeta struct {
	PerPage    int `json:"per_page"`
	Page       int `json:"page"`
	MaxPage    int `json:"max_page"`
	TotalItems int `json:"total_items"`
}
