package requests

type Don_vi_tinh_create struct {
	Ten     string  `json:"ten" binding:"required"`
}

type Don_vi_tinh_update struct {
	Id      int     `json:"id" binding:"required"`
	Ten     string  `json:"ten" binding:"required"`
}

type Don_vi_tinh_delete struct {
	Id int `uri:"id" binding:"required"`
}