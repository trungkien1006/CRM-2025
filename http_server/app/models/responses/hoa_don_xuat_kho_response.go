package responses

import "admin-v1/app/models/db"

type Hoa_don_xuat_kho_filter struct {
	Id                  		int     								`json:"ID"`
	So_hoa_don          		int     								`json:"so_hoa_don"`
	Ma_hoa_don          		string  								`json:"ma_hoa_don"`
	Khach_hang          		string  								`json:"khach_hang"`
	Nhan_vien_sale      		string  								`json:"nhan_vien_sale"`
	Nhan_vien_giao_hang 		string  								`json:"nhan_vien_giao_hang"`
	Ngay_xuat           		string  								`json:"ngay_xuat"`
	Tong_tien           		float32 								`json:"tong_tien"`
	Vat                 		float32 								`json:"vat"`
	Thanh_tien          		float32 								`json:"thanh_tien"`
	Tra_truoc           		float32 								`json:"tra_truoc"`
	Con_lai             		float32 								`json:"con_lai"`
	Tong_gia_nhap       		float32 								`json:"tong_gia_nhap"`
	Loi_nhuan           		float32 								`json:"loi_nhuan"`
	Ghi_chu             		string  								`json:"ghi_chu"`
	Da_giao_hang        		int     								`json:"da_giao_hang"`
	Loai_chiet_khau     		int     								`json:"loai_chiet_khau"`
	Gia_tri_chiet_khau  		string  								`json:"gia_tri_chiet_khau"`
	Khoa_don					bool									`json:"khoa_don"`
	Created_at					string									`json:"CreatedAt"`

	Chi_tiet_hoa_don_xuat_kho 	[]Chi_tiet_hoa_don_xuat_kho_response 	`json:"ds_san_pham_xuat" gorm:"-"`
}

type Chi_tiet_hoa_don_xuat_kho_response struct {
	Id							int										`json:"ID"`
	Hoa_don_id   				int     								`json:"hoa_don_id"`
	Ctsp_ten					string									`json:"ctsp_ten"`
	Ten_san_pham				string									`json:"ten_san_pham"`
	Ctsp_id      				int     								`json:"ctsp_id"`
	Sku          				string  								`json:"sku"`
	Don_vi_tinh  				string  								`json:"don_vi_tinh"`
	So_luong_ban 				int     								`json:"so_luong_ban"`
	Gia_ban      				float32 								`json:"gia_ban"`
	Chiet_khau   				float32 								`json:"chiet_khau"`
	Thanh_tien   				float32 								`json:"thanh_tien"`
	Gia_nhap     				float32 								`json:"gia_nhap"`
	Loi_nhuan    				float32 								`json:"loi_nhuan"`
	La_qua_tang  				int     								`json:"la_qua_tang"`
}

type Hoa_don_xuat_kho_create struct {
	Hoa_don_xuat_kho 			db.Hoa_don_xuat_kho 					`json:"hoa_don_xuat_kho"`
}