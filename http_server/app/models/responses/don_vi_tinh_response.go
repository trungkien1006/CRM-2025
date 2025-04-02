package responses

import "admin-v1/app/models/db"

type Don_vi_tinh_filter struct {
	Don_vi_tinh		[]db.Don_vi_tinh
	Total_Page 		int
}

type Don_vi_tinh_create struct {
	Don_vi_tinh 	db.Don_vi_tinh
}