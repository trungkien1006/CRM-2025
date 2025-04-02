package db

import "gorm.io/gorm"

type Nhan_vien struct {
	gorm.Model

	Ten_dang_nhap 				string				`json:"ten_dang_nhap"`
	Mat_khau      				string				`json:"mat_khau"`
	Ho_ten        				string				`json:"ho_ten"`
	Email         				string				`json:"email"`
	Dien_thoai    				string				`json:"dien_thoai"`
	Dia_chi       				string				`json:"dia_chi"`
	Avatar        				string				`json:"avatar"`
	Chuc_vu_id    				int					`json:"chuc_vu_id"`

	Chuc_vu						Chuc_vu				`json:"chuc_vu"`
	Hoa_don_xuat_kho_sale		[]Hoa_don_xuat_kho	`json:"hoa_don_xuat_kho_sale" gorm:"foreignKey:nhan_vien_sale_id"`
	Hoa_don_xuat_kho_giao_hang	[]Hoa_don_xuat_kho	`json:"hoa_don_xuat_kho_giao_hang" gorm:"foreignKey:nhan_vien_giao_hang_id"`
}