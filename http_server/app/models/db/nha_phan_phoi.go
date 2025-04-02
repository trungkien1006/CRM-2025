package db

import (
	"gorm.io/gorm"
)

type Nha_phan_phoi struct {
	gorm.Model

	Ten        			string				`json:"ten"`
	Dia_chi    			string				`json:"dia_chi"`
	Dien_thoai 			string				`json:"dien_thoai"`
	Email      			string				`json:"email"`

	San_pham   			[]San_pham			`json:"san_pham" gorm:"many2many:san_pham_nha_phan_phoi;constraint:OnUpdate:CASCADE"`
	Hoa_don_nhap_kho 	[]Hoa_don_nhap_kho	`json:"hoa_don_nhap_kho" gorm:"foreignKey:nha_phan_phoi_id"`
}