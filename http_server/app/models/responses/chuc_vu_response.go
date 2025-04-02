package responses

import "admin-v1/app/models/db"

type Chuc_vu_create struct {
	Chuc_vu db.Chuc_vu `json:"chuc_vu"`
}