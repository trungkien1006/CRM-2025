package responses

import "admin-v1/app/models/db"

type Nha_phan_phoi_create struct {
	Nha_phan_phoi db.Nha_phan_phoi	`json:"nha_phan_phoi"`
}

type Nha_phan_phoi_filter struct {
	Id					int									`json:"ID"`
	Ten        			string								`json:"ten"`
	Dia_chi    			string								`json:"dia_chi"`
	Dien_thoai 			string								`json:"dien_thoai"`
	Email      			string								`json:"email"`
	Created_at			string								`json:"CreatedAt"`
	Deleted_at			string								`json:"DeletedAt"`
	Updated_at			string								`json:"UpdatedAt"`

	Ds_san_pham			[]Nha_phan_phoi_san_pham_response	`json:"ds_san_pham" gorm:"-"`
}

type Nha_phan_phoi_san_pham_response struct {
	Id					int									`json:"ID"`
	Ten					string								`json:"ten"`
	Upc					string								`json:"upc"`
	Don_vi_tinh			string								`json:"don_vi_tinh"`
	Nha_phan_phoi_id	int									`json:"nha_phan_phoi_id"`
}