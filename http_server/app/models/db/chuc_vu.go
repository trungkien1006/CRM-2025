package db

import "gorm.io/gorm"

type Chuc_vu struct {
	gorm.Model

	Ten  		string		`json:"ten"`

	Chuc_nang 	[]Chuc_nang	`json:"chuc_nang" gorm:"many2many:quyen;"`
	Nhan_vien	[]Nhan_vien	`json:"nhan_vien" gorm:"chuc_vu_id"`
}