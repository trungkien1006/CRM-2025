package responses

type Filter[T any] struct {
	Data       []T 	`json:"data"`
	Total_Page int 	`json:"total_page"`
}