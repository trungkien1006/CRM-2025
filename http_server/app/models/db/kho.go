package db

import "gorm.io/gorm"

type Kho struct {
	gorm.Model

	Ten        			string				`json:"ten"`
	Dia_chi    			string				`json:"dia_chi"`

	Hoa_don_nhap_kho	[]Hoa_don_nhap_kho	`json:"hoa_don_nhap_kho" gorm:"foreignKey:kho_id"`
}