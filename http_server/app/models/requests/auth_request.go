package requests

type Dang_nhap struct {
	Ten_dang_nhap	string	`json:"ten_dang_nhap" binding:"required"`
	Mat_khau		string	`json:"mat_khau" binding:"required"`
}