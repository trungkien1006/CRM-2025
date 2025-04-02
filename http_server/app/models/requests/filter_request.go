package requests

type FilterStruc struct {
	Field     string `form:"field"`
	Condition string `form:"condition"`
	Value     string `form:"value"`
}

type Filter struct {
	Filters string `form:"filters" binding:"omitempty"`
	Sort    string `form:"sort" binding:"omitempty"`
	Order   string `form:"order" binding:"omitempty"`
	Page    int    `form:"page" binding:"omitempty"`
	Limit   int    `form:"limit" binding:"omitempty"`
}