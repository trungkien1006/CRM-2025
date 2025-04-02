package requests

type Thoi_gian_bao_hanh_create struct {
	Ten string `json:"ten" binding:"required"`
}

type Thoi_gian_bao_hanh_update struct {
	Id  int    `json:"id" binding:"required"`
	Ten string `json:"ten" binding:"required"`
}

type Thoi_gian_bao_hanh_delete struct {
	Id int `uri:"id" binding:"required"`
}