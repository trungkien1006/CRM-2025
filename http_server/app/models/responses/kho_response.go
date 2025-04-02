package responses

import "admin-v1/app/models/db"

type Kho_create struct {
	Kho db.Kho	`json:"kho"`
}