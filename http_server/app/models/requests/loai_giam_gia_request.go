package requests

type Loai_giam_gia_create struct {
	Ten     string `json:"ten" binding:"required"`
	Gia_tri float32    `json:"gia_tri" binding:"required"`
}

type Loai_giam_gia_update struct {
	Id      int    `json:"id" binding:"required"`
	Ten     string `json:"ten" binding:"required"`
	Gia_tri float32    `json:"gia_tri" binding:"required"`
}

type Loai_giam_gia_delete struct {
	Id int `uri:"id" binding:"required"`
}