package db

import "gorm.io/gorm"

type Hoa_don_nhap_kho struct {
	gorm.Model

	So_hoa_don					int							`json:"so_hoa_don"`
	Ma_hoa_don					string						`json:"ma_hoa_don"`
	Nha_phan_phoi_id 			int							`json:"nha_phan_phoi_id"`
	Kho_id 						int							`json:"kho_id"`
	Ngay_nhap 					string						`json:"ngay_nhap"`
	Tong_tien 					float32						`json:"tong_tien"`
	Tra_truoc					float32						`json:"tra_truoc"`
	Con_lai						float32						`json:"con_lai"`
	Ghi_chu						string						`json:"ghi_chu"`
	Khoa_don					bool						`json:"khoa_don"`

	Nha_phan_phoi 				Nha_phan_phoi				`json:"nha_phan_phoi"`
	Kho							Kho							`json:"kho"`
	Chi_tiet_hoa_don_nhap_kho	[]Chi_tiet_hoa_don_nhap_kho	`json:"chi_tiet_hoa_don_nhap_kho" gorm:"foreignKey:hoa_don_id;constraint:OnUpdate:CASCADE"`
}