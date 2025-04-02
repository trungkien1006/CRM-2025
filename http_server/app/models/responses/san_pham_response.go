package responses

import "admin-v1/app/models/db"

type San_pham_filter struct {
	ID 							int 						`json:"ID"`
	Ten                   		string						`json:"ten"`
	Upc                   		string						`json:"upc"`
	Loai_san_pham     			string						`json:"loai_san_pham"`
	Loai_san_pham_id     		int							`json:"loai_san_pham_id"`
	Hinh_anh              		string						`json:"hinh_anh"`
	Don_vi_tinh        			string						`json:"don_vi_tinh"`
	Don_vi_tinh_id     			int							`json:"don_vi_tinh_id"`
	Vat                   		float32						`json:"vat"`
	Mo_ta                 		string						`json:"mo_ta"`
	Trang_thai            		int							`json:"trang_thai"`
	Loai_giam_gia      			string						`json:"loai_giam_gia"`
	Loai_giam_gia_id     		int							`json:"loai_giam_gia_id"`
	Thoi_gian_bao_hanh 			string						`json:"thoi_gian_bao_hanh"`
	Thoi_gian_bao_hanh_id     	int							`json:"thoi_gian_bao_hanh_id"`
	Created_At					string						`json:"created_at"`
}

type San_pham_create struct {
	San_pham 	db.San_pham `json:"san_pham"`
}