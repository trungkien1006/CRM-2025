package requests

type San_pham_create struct {
	Ten   					string 						`form:"ten" binding:"required"`
	Upc                   	string 						`form:"upc" binding:"required"`
	Loai_san_pham_id      	int							`form:"loai_san_pham_id" binding:"required"`
	Hinh_anh              	string 						`form:"hinh_anh" binding:"required"`
	Don_vi_tinh_id        	int							`form:"don_vi_tinh_id" binding:"required"`
	Vat                   	float32						`form:"vat" binding:"omitempty"`
	Mo_ta                 	string						`form:"mo_ta" binding:"omitempty"`
	Trang_thai            	int							`form:"trang_thai"`
	Loai_giam_gia_id      	int							`form:"loai_giam_gia_id" binding:"omitempty"`
	Thoi_gian_bao_hanh_id 	int							`form:"thoi_gian_bao_hanh_id" binding:"omitempty"`

	Chi_tiet_san_pham 		[]Chi_tiet_san_pham_request `form:"chi_tiet_san_pham" binding:"omitempty"`
}

type San_pham_update struct {
	Id    					int    						`form:"id" binding:"required"`
	Ten   					string 						`form:"ten" binding:"required"`
	Upc                   	string 						`form:"upc" binding:"required"`
	Loai_san_pham_id      	int							`form:"loai_san_pham_id"`
	Hinh_anh              	string 						`form:"hinh_anh" binding:"omitempty"`
	Don_vi_tinh_id        	int							`form:"don_vi_tinh_id"`
	Vat                   	float32						`form:"vat" binding:"omitempty"`
	Mo_ta                 	string						`form:"mo_ta" binding:"omitempty"`
	Trang_thai            	int							`form:"trang_thai"`
	Loai_giam_gia_id      	int							`form:"loai_giam_gia_id"`
	Thoi_gian_bao_hanh_id 	int							`form:"thoi_gian_bao_hanh_id"`

	Chi_tiet_san_pham 		[]Chi_tiet_san_pham_request `form:"chi_tiet_san_pham" binding:"omitempty"`
}

type San_pham_delete struct {
	Id int `uri:"id" binding:"required"`
}