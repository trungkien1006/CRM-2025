package responses

type Dang_nhap struct {
	Token 			string		`json:"token"`
	Ds_quyen 		[]string	`json:"ds_quyen"`
}

type Get_me struct {
	Id 				uint		`json:"ID"`
	Avatar 			string		`json:"avatar"`	
	Chuc_vu			string		`json:"chuc_vu"`
	Chuc_vu_id		string		`json:"chuc_vu_id"`
	Dia_chi			string		`json:"dia_chi"`
	Dien_thoai		string		`json:"dien_thoai"`
	Email			string		`json:"email"`
	Ho_ten			string		`json:"ho_ten"`
	Ten_dang_nhap	string		`json:"ten_dang_nhap"`
	Created_at		string		`json:"CreatedAt"`
	Updated_at		string		`json:"UpdateAt"`
	Deleted_at		string		`json:"DeletedAt"`
	
	Quyen 			[]string	`json:"quyen"`
}
