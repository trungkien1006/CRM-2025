package requests

type Hoa_don_nhap_kho_create struct {
	So_hoa_don					int									`json:"so_hoa_don" binding:"omitempty"`
	Ma_hoa_don					string								`json:"ma_hoa_don" binding:"omitempty"`
	Nha_phan_phoi_id 			int									`json:"nha_phan_phoi_id"`
	Kho_id 						int									`json:"kho_id"`
	Ngay_nhap 					string								`json:"ngay_nhap"`
	Tong_tien 					float32								`json:"tong_tien" binding:"min=0"`
	Tra_truoc					float32								`json:"tra_truoc" binding:"min=0"`
	Con_lai						float32								`json:"con_lai" binding:"min=0"`
	Ghi_chu						string								`json:"ghi_chu" binding:"omitempty"`

	Ds_san_pham_nhap 			[]Chi_tiet_hoa_don_nhap_kho_create 	`json:"ds_san_pham_nhap"`
}

type Hoa_don_nhap_kho_update struct {
	Hoa_don_id					int									`json:"hoa_don_id"`
	Ngay_nhap 					string								`json:"ngay_nhap"`
	Tra_truoc					float32								`json:"tra_truoc" binding:"min=0"`
	Ghi_chu						string								`json:"ghi_chu"`
}

type Hoa_don_nhap_kho_lock struct {
	Hoa_don_id					int 								`json:"hoa_don_id"`
	Lock_or_open				string								`json:"lock_or_open" binding:"oneof=lock open"`
}

type Chi_tiet_san_pham_nhap_tra struct {
	Cthd_nhap_kho_id			int									`json:"cthd_nhap_kho_id"`
	Sku							string								`json:"sku"`
	So_luong_tra				int									`json:"so_luong_tra"`
}

type Tra_no_nhap_kho_request struct {
	Hoa_don_id					int									`json:"hoa_don_id"`
	Tien_tra					float32								`json:"tien_tra"`
}

type Tra_hang_nhap_kho_request struct {
	Hoa_don_id					int									`json:"hoa_don_id"`

	Ds_san_pham_tra				[]Chi_tiet_san_pham_nhap_tra		`json:"ds_san_pham_tra"`
}