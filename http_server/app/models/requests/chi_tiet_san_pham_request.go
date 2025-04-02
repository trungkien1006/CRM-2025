package requests

type Chi_tiet_san_pham_request struct {
	Id 				uint	`form:"id" binding:"required"`
	Ten_phan_loai 	string  `form:"ten_phan_loai" binding:"required"`
	Hinh_anh      	string 	`form:"hinh_anh" binding:"omitempty"`
	Trang_thai    	int     `form:"trang_thai" binding:"required"`
}

type Chi_tiet_san_pham_get_by_product_id struct {
	Product_id 		int		`uri:"product_id" binding:"required"`
}