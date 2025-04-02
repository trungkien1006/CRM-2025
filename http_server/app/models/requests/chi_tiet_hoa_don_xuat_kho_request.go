package requests

type Chi_tiet_hoa_don_xuat_kho_create struct {
	San_pham_id			int					`json:"san_pham_id" binding:"required"`
	Ctsp_id				int					`json:"ctsp_id" binding:"required"`
	Sku					string				`json:"sku" binding:"required"`
	Don_vi_tinh			string				`json:"don_vi_tinh" binding:"required"`
	So_luong_ban		int					`json:"so_luong_ban" binding:"required"`
	Gia_ban				float32				`json:"gia_ban" binding:"min=0"`
	Chiet_khau			float32				`json:"chiet_khau" binding:"min=0"`
	Thanh_tien			float32				`json:"thanh_tien" binding:"min=0"`
	Gia_nhap			float32				`json:"gia_nhap" binding:"min=0"`
	Loi_nhuan			float32				`json:"loi_nhuan" binding:"min=0"`
	La_qua_tang			bool				`json:"la_qua_tang" binding:"min=false"`

	Ds_sku				[]Chi_tiet_sku		`json:"ds_sku" binding:"required"`			
}

type Chi_tiet_sku struct {
	Sku					string				`json:"sku" binding:"omitempty"`
	So_luong_ban		int					`json:"so_luong_ban" binding:"omitempty"`
	// Gia_nhap			float32				`json:"gia_nhap" binding:"omitempty"`
	Gia_ban				float32				`json:"gia_ban" binding:"omitempty"`
}