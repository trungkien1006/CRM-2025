package db

import "gorm.io/gorm"

type San_pham struct {
	gorm.Model

	Ten                   		string						`json:"ten"`
	Upc                   		string						`json:"upc"`
	Loai_san_pham_id      		int							`json:"loai_san_pham_id"`
	Hinh_anh              		string						`json:"hinh_anh"`
	Don_vi_tinh_id        		int							`json:"don_vi_tinh_id"`
	Vat                   		float32						`json:"vat"`
	Mo_ta                 		string						`json:"mo_ta"`
	Trang_thai            		int							`json:"trang_thai"`
	Loai_giam_gia_id      		int							`json:"loai_giam_gia_id"`
	Thoi_gian_bao_hanh_id 		int							`json:"thoi_gian_bao_hanh_id"`

	Nha_phan_phoi	  			[]Nha_phan_phoi	  			`json:"nha_phan_phoi" gorm:"many2many:san_pham_nha_phan_phoi"`	
	Chi_tiet_san_pham 			[]Chi_tiet_san_pham 		`json:"chi_tiet_san_pham" gorm:"foreignKey:san_pham_id;constraint:OnUpdate:CASCADE"`
	Loai_san_pham      			Loai_san_pham      			`json:"loai_san_pham"`
	Don_vi_tinh        			Don_vi_tinh        			`json:"don_vi_tinh"`
	Loai_giam_gia      			Loai_giam_gia      			`json:"loai_giam_gia"`
	Thoi_gian_bao_hanh 			Thoi_gian_bao_hanh 			`json:"thoi_gian_bao_hanh"`
	Chi_tiet_hoa_don_nhap_kho 	[]Chi_tiet_hoa_don_nhap_kho `json:"chi_tiet_hoa_don_nhap_kho" gorm:"foreignKey:san_pham_id"`
	Chi_tiet_hoa_don_xuat_kho 	[]Chi_tiet_hoa_don_xuat_kho `json:"chi_tiet_hoa_don_xuat_kho" gorm:"foreignKey:san_pham_id"`
	Ton_kho						[]Ton_kho					`json:"ton_kho" gorm:"foreignKey:san_pham_id"`
}