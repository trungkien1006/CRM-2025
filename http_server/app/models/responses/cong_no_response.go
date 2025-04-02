package responses

type Cong_no_khach_hang_filter struct {
	Khach_hang			string		`json:"khach_hang"`
	Khach_hang_id		int			`json:"khach_hang_id"`
	Tong_hoa_don		int			`json:"tong_hoa_don"`
	Tong_tien			float32		`json:"tong_tien"`
	Con_lai				float32		`json:"con_lai"`
	Tra_truoc			float32		`json:"tra_truoc"`
}

type Cong_no_nha_phan_phoi_filter struct {
	Nha_phan_phoi		string		`json:"nha_phan_phoi"`
	Nha_phan_phoi_id	int			`json:"nha_phan_phoi_id"`
	Tong_hoa_don		int			`json:"tong_hoa_don"`
	Tong_tien			float32		`json:"tong_tien"`
	Con_lai				float32		`json:"con_lai"`
	Tra_truoc			float32		`json:"tra_truoc"`
}