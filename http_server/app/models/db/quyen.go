package db

import "gorm.io/gorm"

type Quyen struct {
	gorm.Model

	Chuc_vu_id   	int 		`gorm:"primaryKey"`
	Chuc_nang_id 	int			`gorm:"primaryKey"`

	ChucVu   		Chuc_vu   	`gorm:"foreignKey:Chuc_vu_id"`
	ChucNang 		Chuc_nang 	`gorm:"foreignKey:Chuc_nang_id"`
}