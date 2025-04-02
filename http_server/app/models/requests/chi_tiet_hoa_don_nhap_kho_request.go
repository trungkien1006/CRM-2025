package requests

type Chi_tiet_hoa_don_nhap_kho_create struct {
	Hoa_don_id  		uint			`json:"hoa_don_id" binding:"required"`
	San_pham_id 		int				`json:"san_pham_id" binding:"required"`
	Upc 				string			`json:"upc" binding:"required"`
	Ctsp_id     		int				`json:"ctsp_id" binding:"required"`
	Sku         		string			`json:"sku" binding:"omitempty"`
	So_luong    		int				`json:"so_luong" binding:"required"`
	Don_vi_tinh 		string			`json:"don_vi_tinh" binding:"required"`
	Ke          		string			`json:"ke" binding:"required"`
	Gia_nhap    		float32			`json:"gia_nhap" binding:"min=0"`
	Gia_ban     		float32			`json:"gia_ban" binding:"min=0"` 
	Chiet_khau  		float32			`json:"chiet_khau" binding:"min=0"`
	Thanh_tien  		float32			`json:"thanh_tien" binding:"min=0"`
	La_qua_tang 		bool			`json:"la_qua_tang" binding:"min=false"`
	Han_su_dung			string			`json:"han_su_dung" binding:"required"`
}