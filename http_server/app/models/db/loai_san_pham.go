package db

import "gorm.io/gorm"

type Loai_san_pham struct {
	gorm.Model

	Ten        string `json:"ten"`
	Hinh_anh   string `json:"hinh_anh"`

	San_pham []San_pham `json:"san_pham" gorm:"foreignKey:loai_san_pham_id"`
}