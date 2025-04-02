package responses

type Quyen_by_chuc_vu_id struct {
	Quyen_list []Quyen_list
}

type Quyen_list struct {
	Id 	uint	`gorm:"id"`
	Quyen		[]string
}