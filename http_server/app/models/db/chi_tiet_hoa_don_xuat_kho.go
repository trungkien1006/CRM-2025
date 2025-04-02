package db

import "gorm.io/gorm"

type Chi_tiet_hoa_don_xuat_kho struct {
	gorm.Model

	Hoa_don_id			uint				`json:"hoa_don_id"`
	San_pham_id			uint				`json:"san_pham_id"`
	Ctsp_id				uint				`json:"ctsp_id"`
	Sku					string				`json:"sku"`
	Don_vi_tinh			string				`json:"don_vi_tinh"`
	So_luong_ban		int					`json:"so_luong_ban"`
	Gia_ban				float32				`json:"gia_ban"`
	Chiet_khau			float32				`json:"chiet_khau"`
	Thanh_tien			float32				`json:"thanh_tien"`
	Gia_nhap			float32				`json:"gia_nhap"`
	Loi_nhuan			float32				`json:"loi_nhuan"`
	La_qua_tang			bool				`json:"la_qua_tang"`

	Hoa_don_xuat_kho	Hoa_don_xuat_kho	`json:"hoa_don_xuat_kho" gorm:"foreignKey:Hoa_don_id"`
	San_pham			San_pham			`json:"san_pham" gorm:"foreignKey:San_pham_id"`
	Chi_tiet_san_pham	Chi_tiet_san_pham	`json:"chi_tiet_san_pham" gorm:"foreignKey:Ctsp_id"`
}