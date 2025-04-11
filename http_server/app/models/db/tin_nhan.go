package db

import "gorm.io/gorm"

type Tin_nhan struct {
	gorm.Model

	Sender_id		int			`json:"sender_id"`
	Receiver_id		int			`json:"receiver_id"`
	Content			int			`json:"content"`
	Is_read			bool		`json:"is_read"`

	Sender			Nhan_vien	`json:"sender" gorm:"foreignKey:Nhan_vien_id"`
	Receiver		Nhan_vien	`json:"receiver" gorm:"foreignKey:Nhan_vien_id"`
}