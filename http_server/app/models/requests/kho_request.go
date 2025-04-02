package requests

type Kho_create struct {
	Ten		string	`json:"ten" binding:"required"`
	Dia_chi	string	`json:"dia_chi" binding:"required"`
}

type Kho_update struct {
	Id 		int		`json:"id" binding:"required"`
	Ten		string	`json:"ten" binding:"required"`
	Dia_chi	string	`json:"dia_chi" binding:"required"`
}

type Kho_delete struct {
	Id 		int		`uri:"id" binding:"required"`
}