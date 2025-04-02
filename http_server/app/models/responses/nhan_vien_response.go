package responses

import "admin-v1/app/models/db"

type Nhan_vien_filter struct {
	Id				int		`json:"ID"`
	Ten_dang_nhap 	string	`json:"ten_dang_nhap"`
	Mat_khau      	string	`json:"mat_khau"`
	Ho_ten        	string	`json:"ho_ten"`
	Email         	string	`json:"email"`
	Dien_thoai    	string	`json:"dien_thoai"`
	Dia_chi       	string	`json:"dia_chi"`
	Avatar        	string	`json:"avatar"`
	Chuc_vu    		string	`json:"chuc_vu"`
	Chuc_vu_id    	string	`json:"chuc_vu_id"`
	Created_at		string	`json:"CreatedAt"`
}

type Nhan_vien_create struct {
	Nhan_vien 	db.Nhan_vien `json:"nhan_vien"`
}