package db

import "gorm.io/gorm"

type Chi_tiet_hoa_don_nhap_kho struct {
	gorm.Model

	Hoa_don_id  		uint				`json:"hoa_don_id"`
	San_pham_id 		uint				`json:"san_pham_id"`
	Ctsp_id     		uint				`json:"ctsp_id"`
	Sku         		string				`json:"sku"`
	So_luong    		int					`json:"so_luong"`
	Don_vi_tinh 		string				`json:"don_vi_tinh"`
	Ke          		string				`json:"ke"`
	Gia_nhap    		float32				`json:"gia_nhap"`
	Gia_ban     		float32				`json:"gia_ban"`
	Chiet_khau  		float32				`json:"chiet_khau"`
	Thanh_tien  		float32				`json:"thanh_tien"`
	La_qua_tang 		bool				`json:"la_qua_tang"`
	Han_su_dung			string				`json:"han_su_dung"`

	Hoa_don_nhap_kho 	Hoa_don_nhap_kho 	`json:"hoa_don_nhap_kho" gorm:"foreignKey:Hoa_don_id"`
	San_pham 			San_pham 			`json:"san_pham" gorm:"foreignKey:San_pham_id"`
	Chi_tiet_san_pham 	Chi_tiet_san_pham 	`json:"chi_tiet_san_pham" gorm:"foreignKey:Ctsp_id"`
}