package requests

type Nha_phan_phoi_create struct {
	Ten        	string	`json:"ten" binding:"required"`
	Dia_chi    	string	`json:"dia_chi" binding:"required"`
	Dien_thoai 	string	`json:"dien_thoai" binding:"required"`
	Email      	string	`json:"email" binding:"required"`

	Ds_san_pham	[]int 	`json:"ds_san_pham" binding:"omitempty"`
}	

type Nha_phan_phoi_update struct {
	Id 			int		`json:"id" binding:"required"`
	Ten        	string	`json:"ten" binding:"required"`
	Dia_chi    	string	`json:"dia_chi" binding:"required"`
	Dien_thoai 	string	`json:"dien_thoai" binding:"required"`
	Email      	string	`json:"email" binding:"required"`

	Ds_san_pham	[]int 	`json:"ds_san_pham" binding:"omitempty"`
}

type Nha_phan_phoi_delete struct {
	Id			int		`uri:"id" binding:"required"`
}