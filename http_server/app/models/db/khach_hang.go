package db

import "gorm.io/gorm"

type Khach_hang struct {
	gorm.Model

	Ho_ten        		string				`json:"ho_ten"`
	Dien_thoai    		string				`json:"dien_thoai"`
	Dia_chi       		string				`json:"dia_chi"`

	Hoa_don_xuat_kho	[]Hoa_don_xuat_kho	`json:"hoa_don_xuat_kho" gorm:"khach_hang_id"`
}