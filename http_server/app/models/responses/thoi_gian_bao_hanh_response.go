package responses

import "admin-v1/app/models/db"

type Thoi_gian_bao_hanh_filter struct {
	Thoi_gian_bao_hanh 	[]db.Thoi_gian_bao_hanh `json:"thoi_gian_bao_hanh"`
	Total_Page 			int	`json:"total_page"`
}

type Thoi_gian_bao_hanh_create struct {
	Thoi_gian_bao_hanh db.Thoi_gian_bao_hanh `json:"thoi_gian_bao_hanh"`
}