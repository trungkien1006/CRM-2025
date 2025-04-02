package db

import "gorm.io/gorm"

type San_pham_nha_phan_phoi struct {
	gorm.Model

	Nha_phan_phoi_id 	int				`json:"nha_phan_phoi_id" gorm:"primaryKey"`
	San_pham_id      	int				`json:"san_pham_id" gorm:"primaryKey"`

	Nha_phan_phoi		Nha_phan_phoi	`json:"nha_phan_phoi" gorm:"foreignKey:Nha_phan_phoi_id"`
	San_pham			San_pham		`json:"san_pham" gorm:"foreignKey:San_pham_id"`
}