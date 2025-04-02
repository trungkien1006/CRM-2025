package responses

import "admin-v1/app/models/db"

type Loai_giam_gia_filter struct {
	Loai_giam_gia 	[]db.Loai_giam_gia 	`json:"loai_giam_gia"`
	Total_Page 		int	`json:"total_page"`
}

type Loai_giam_gia_create struct {
	Loai_giam_gia 	db.Loai_giam_gia `json:"loai_giam_gia"`
}