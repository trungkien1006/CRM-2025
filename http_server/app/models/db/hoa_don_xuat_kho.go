package db

import "gorm.io/gorm"

type Hoa_don_xuat_kho struct {
	gorm.Model

	So_hoa_don					int							`json:"so_hoa_don"`
	Ma_hoa_don					string						`json:"ma_hoa_don"`
	Khach_hang_id				int							`json:"khach_hang_id"`
	Nhan_vien_sale_id			int							`json:"nhan_vien_sale_id"`
	Nhan_vien_giao_hang_id		int							`json:"nhan_vien_giao_hang_id"`
	Ngay_xuat					string						`json:"ngay_xuat"`
	Tong_tien					float32						`json:"tong_tien"`
	Vat							float32						`json:"vat"`
	Thanh_tien					float32						`json:"thanh_tien"`
	Tra_truoc 					float32						`json:"tra_truoc"`
	Con_lai						float32						`json:"con_lai"`
	Tong_gia_nhap				float32						`json:"tong_gia_nhap"`
	Loi_nhuan					float32						`json:"loi_nhuan"`
	Ghi_chu						string						`json:"ghi_chu"`
	Da_giao_hang				bool						`json:"da_giao_hang"`
	Loai_chiet_khau				int							`json:"loai_chiet_khau"`
	Gia_tri_chiet_khau			int							`json:"gia_tri_chiet_khau"`
	Khoa_don					bool						`json:"khoa_don"`

	Khach_hang					Khach_hang					`json:"khach_hang"`
	Nhan_vien_sale				Nhan_vien					`json:"nhan_vien_sale"`
	Nhan_vien_giao_hang			Nhan_vien					`json:"nhan_vien_giao_hang"`
	Chi_tiet_hoa_don_xuat_kho	[]Chi_tiet_hoa_don_xuat_kho	`json:"chi_tiet_hoa_don_xuat_kho" gorm:"foreignKey:hoa_don_id"`
}