package responses

import "admin-v1/app/models/db"

type Khach_hang_create struct {
	Khach_hang 	db.Khach_hang	`json:"khach_hang"`	
}