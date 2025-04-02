package db

import "gorm.io/gorm"

type Don_vi_tinh struct {
	gorm.Model

	Ten        string	`json:"ten"`

	San_pham []San_pham `json:"san_pham" gorm:"foreignKey:don_vi_tinh_id"`
}