package requests

type Khach_hang_create struct {
	Ho_ten     string `json:"ho_ten" binding:"required"`
	Dien_thoai string `json:"dien_thoai" binding:"required"`
	Dia_chi    string `json:"dia_chi" binding:"required"`
}

type Khach_hang_update struct {
	Id         string `json:"id" binding:"required"`
	Ho_ten     string `json:"ho_ten" binding:"required"`
	Dien_thoai string `json:"dien_thoai" binding:"required"`
	Dia_chi    string `json:"dia_chi" binding:"required"`
}

type Khach_hang_delete struct {
	Id         string `uri:"id" binding:"required"`
}