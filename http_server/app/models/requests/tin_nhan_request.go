package requests

type Tin_nhan_create struct {
	Sender_id		int					`json:"sender_id" binding:"required"`
	Receiver_id		int					`json:"receiver_id" binding:"required"`
	Content			int					`json:"content" binding:"required"`
}

type Tin_nhan_create_batch struct {
	Tin_nhan_batch 	[]Tin_nhan_create	`json:"tin_nhan_batch"`
} 