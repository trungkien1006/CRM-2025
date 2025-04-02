package responses

import "admin-v1/app/models/db"

type Hoa_don_nhap_kho_filter struct {
	Id							int										`json:"ID"`
	Nha_phan_phoi 				string									`json:"nha_phan_phoi"`
	Kho 						string									`json:"kho"`
	So_hoa_don					int										`json:"so_hoa_don"`
	Ma_hoa_don					string									`json:"ma_hoa_don"`
	Ngay_nhap 					string									`json:"ngay_nhap"`
	Tong_tien 					float32									`json:"tong_tien"`
	Tra_truoc					float32									`json:"tra_truoc"`
	Con_lai						float32									`json:"con_lai"`
	Ghi_chu						string									`json:"ghi_chu"`
	Khoa_don					bool									`json:"khoa_don"`
	Created_at					string									`json:"CreatedAt"`

	Chi_tiet_hoa_don_nhap_kho	[]Chi_tiet_hoa_don_nhap_kho_response	`json:"ds_san_pham_nhap" gorm:"-"`
}

type Chi_tiet_hoa_don_nhap_kho_response struct {
	Id					int												`json:"ID"`
	Hoa_don_id			int												`json:"hoa_don_id"`
	Sku					string											`json:"sku"`
	Ctsp_ten					string									`json:"ctsp_ten"`
	Ten_san_pham		string											`json:"ten_san_pham"`
	So_luong    		int												`json:"so_luong"`
	Don_vi_tinh 		string											`json:"don_vi_tinh"`
	Ke          		string											`json:"ke"`
	Gia_nhap    		float32											`json:"gia_nhap"`
	Gia_ban     		float32											`json:"gia_ban"`
	Chiet_khau  		float32											`json:"chiet_khau"`
	Thanh_tien  		string											`json:"thanh_tien"`
	La_qua_tang 		int												`json:"la_qua_tang"`
	Han_su_dung			string											`json:"han_su_dung"`
}

type Hoa_don_nhap_kho_create struct {
	Hoa_don_nhap_kho 	db.Hoa_don_nhap_kho 							`json:"hoa_don_nhap_kho"`
}