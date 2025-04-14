package responses

type Tin_nhan_filter struct {
	ID          	int    		`json:"ID"`
	Sender_id   	int    		`json:"sender_id"`
	Receiver_id 	int    		`json:"receiver_id"`
	Content     	string   	`json:"content"`
	Is_read     	bool   		`json:"is_read"`
	Created_at  	string 		`json:"created_at"`
}