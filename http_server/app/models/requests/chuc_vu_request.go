package requests

type Chuc_vu_create struct {
	Ten   string `form:"ten" binding:"required"`
}

type Chuc_vu_update struct {
	Id    int    `form:"id" binding:"required"`
	Ten   string `form:"ten" binding:"required"`
}

type Chuc_vu_delete struct {
	Id int `uri:"id" binding:"required"`
}