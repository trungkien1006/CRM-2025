package db

import "gorm.io/gorm"

type Tin_nhan struct {
	gorm.Model

	Sender_id		int			`json:"sender_id"`
	Receiver_id		int			`json:"receiver_id"`
	Content			string		`json:"content"`
	Is_read			bool		`json:"is_read"`

	Sender			Nhan_vien	`json:"sender" gorm:"foreignKey:Sender_id"`
	Receiver		Nhan_vien	`json:"receiver" gorm:"foreignKey:Receiver_id"`
}