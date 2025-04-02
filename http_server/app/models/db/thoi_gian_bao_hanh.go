package db

import "gorm.io/gorm"

type Thoi_gian_bao_hanh struct {
	gorm.Model

	Ten string `json:"ten"`

	San_pham []San_pham `json:"san_pham" gorm:"foreignKey:thoi_gian_bao_hanh_id"`
}