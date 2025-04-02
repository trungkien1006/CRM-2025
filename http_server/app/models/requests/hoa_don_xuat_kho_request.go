package requests

type Hoa_don_xuat_kho_create struct {
	So_hoa_don					int							`json:"so_hoa_don" binding:"omitempty"`
	Ma_hoa_don					string						`json:"ma_hoa_don" binding:"omitempty"`
	Khach_hang_id				int							`json:"khach_hang_id"`
	Nhan_vien_sale_id			int							`json:"nhan_vien_sale_id"`
	Nhan_vien_giao_hang_id		int							`json:"nhan_vien_giao_hang_id"`
	Ngay_xuat					string						`json:"ngay_xuat"`
	Tong_tien					float32						`json:"tong_tien" binding:"min=0"`
	Vat							float32						`json:"vat" binding:"min=0"`
	Thanh_tien					float32						`json:"thanh_tien" binding:"min=0"`
	Tra_truoc 					float32						`json:"tra_truoc" binding:"min=0"`
	Con_lai						float32						`json:"con_lai" binding:"min=0"`
	Tong_gia_nhap				float32						`json:"tong_gia_nhap" binding:"min=0"`
	Loi_nhuan					float32						`json:"loi_nhuan" binding:"min=0"`
	Ghi_chu						string						`json:"ghi_chu" binding:"omitempty"`
	Da_giao_hang				bool						`json:"da_giao_hang"`
	Loai_chiet_khau				int							`json:"loai_chiet_khau"`
	Gia_tri_chiet_khau			int							`json:"gia_tri_chiet_khau" binding:"min=0"`

	Chi_tiet_hoa_don_xuat_kho 	[]Chi_tiet_hoa_don_xuat_kho_create 	`json:"ds_san_pham_xuat"`
}

type Hoa_don_xuat_kho_update struct {
	Hoa_don_id					int									`json:"hoa_don_id"`
	Khach_hang_id				int									`json:"khach_hang_id"`
	Nhan_vien_sale_id			int									`json:"nhan_vien_sale_id"`
	Nhan_vien_giao_hang_id		int									`json:"nhan_vien_giao_hang_id"`
	Ngay_xuat 					string								`json:"ngay_xuat"`
	Vat							float32								`json:"vat" binding:"min=0"`
	Tra_truoc					float32								`json:"tra_truoc" binding:"min=0"`
	Ghi_chu						string								`json:"ghi_chu" binding:"omitempty"`
	Da_giao_hang				bool								`json:"da_giao_hang"`
	Loai_chiet_khau				int									`json:"loai_chiet_khau" binding:"omitempty"`
	Gia_tri_chiet_khau			int									`json:"gia_tri_chiet_khau" binding:"min=0"`
}

type Hoa_don_xuat_kho_lock struct {
	Hoa_don_id					int 								`json:"hoa_don_id"`
	Lock_or_open				string								`json:"lock_or_open" binding:"oneof=lock open"`
}

type Chi_tiet_san_pham_xuat_tra struct {
	Cthd_xuat_kho_id			int									`json:"cthd_xuat_kho_id"`
	Sku							string								`json:"sku"`
	So_luong_tra				int									`json:"so_luong_tra"`
}

type Tra_no_xuat_kho_request struct {
	Hoa_don_id					int									`json:"hoa_don_id"`
	Tien_tra					float32								`json:"tien_tra"`
}

type Tra_hang_xuat_kho_request struct {
	Hoa_don_id					int									`json:"hoa_don_id"`

	Ds_san_pham_tra				[]Chi_tiet_san_pham_xuat_tra		`json:"ds_san_pham_tra"`
}