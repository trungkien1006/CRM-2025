package db

import "gorm.io/gorm"

type Loai_giam_gia struct {
	gorm.Model

	Ten        string `json:"ten"`
	Gia_tri    float32 `json:"gia_tri"`

	San_pham []San_pham `json:"san_pham" gorm:"foreignKey:loai_giam_gia_id"`
}