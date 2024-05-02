package helper

type Pagination struct {
	Page       int     `json:"page" query:"page" default:"1"`
	Row        int     `json:"row" default:"5"`
	Sort       string  `default:"name asc" json:"sort" query:"sort"`
	TotalRows  int64   `json:"total_rows"`
	TotalPages float64 `json:"total_page"`
}

type Fil struct {
	Totalpage     int
	Prevpage      int
	Nextpage      int
	Totalrows     int
	TotalNextpage int
	Totalprevpage int
}
