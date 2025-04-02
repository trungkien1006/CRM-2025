package db

import "gorm.io/gorm"

type Chi_tiet_san_pham struct {
	gorm.Model
	
	ID							uint						`json:"ID"`
	San_pham_id   				uint						`json:"san_pham_id"`
	Ten_phan_loai 				string 						`json:"ten_phan_loai"`
	Hinh_anh					string						`json:"hinh_anh"`
	Gia_nhap					float32						`json:"gia_nhap"`
	Gia_ban						float32						`json:"gia_ban"`
	So_luong 					int							`json:"so_luong"`
	Trang_thai					int							`json:"trang_thai"`
	Khong_phan_loai				int							`json:"khong_phan_loai"`

	San_pham 					San_pham 					`json:"san_pham"`
	Chi_tiet_hoa_don_nhap_kho 	[]Chi_tiet_hoa_don_nhap_kho `json:"chi_tiet_hoa_don_nhap_kho" gorm:"foreignKey:ctsp_id"`
	Chi_tiet_hoa_don_xuat_kho 	[]Chi_tiet_hoa_don_xuat_kho `json:"chi_tiet_hoa_don_xuat_kho" gorm:"foreignKey:ctsp_id"`
	Ton_kho						[]Ton_kho					`json:"ton_kho" gorm:"foreignKey:ctsp_id"`
}
