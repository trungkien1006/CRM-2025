package db

import "gorm.io/gorm"

type Chuc_nang struct {
	gorm.Model

	Ten 			string
	Code 			string	
	Loai 			string	
	Show_in_menu 	string	
	Thuoc_chuc_vu 	bool	`gorm:"column:thuoc_chuc_vu"`

	Chuc_vu			[]Chuc_vu	`gorm:"many2many:quyen"`
}