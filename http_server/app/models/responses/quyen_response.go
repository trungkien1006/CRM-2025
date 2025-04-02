package responses

type Quyen_read struct {
	Quyen []Quyen	`json:"quyen"`
}

type Quyen struct {
	Hien_thi_menu string      `json:"hien_thi_menu"`
	Loai_quyen    string      `json:"loai_quyen"`
	Quyen         []Quyen_sub `json:"quyen"`
}

type Quyen_sub struct {
	Id         uint   `json:"ID"`
	Ten        string `json:"ten"`
	Code       string `json:"code"`
	Trang_thai bool   `json:"trang_thai"`
}