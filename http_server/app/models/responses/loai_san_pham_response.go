package responses

import "admin-v1/app/models/db"

type Loai_san_pham_filter struct {
	Loai_san_pham 	[]db.Loai_san_pham `json:"loai_san_pham"`
	Total_Page 		int	`json:"total_page"`
}

type Loai_san_pham_create struct {
	Loai_san_pham 	db.Loai_san_pham `json:"loai_san_pham"`
}