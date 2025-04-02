package responses

type Ton_kho_response struct {
	Ds_ton_kho 					[]Ton_kho_item				`json:"ds_ton_kho"`
}

type Ton_kho_item struct {
	Id							int							`json:"id"`
	San_pham_id					int							`json:"san_pham_id"`
	Ctsp_id						int							`json:"ctsp_id"`
	Sku 						string						`json:"sku"`
	So_luong_ton				int							`json:"so_luong_ton"`
	Han_su_dung					string						`json:"han_su_dung"`
	Don_vi_tinh					string						`json:"don_vi_tinh"`
	Gia_ban						float32						`json:"gia_ban"`
	Created_at					string						`json:"CreatedAt"`
}