package helper

type Pagination struct {
	Page       int     `json:"page" query:"page"`
	Row        int     `json:"row"`
	Sort       string  `default:"name asc" json:"sort" query:"sort"`
	TotalRows  int64   `json:"total_rows"`
	TotalPages float64 `json:"total_page"`
}
