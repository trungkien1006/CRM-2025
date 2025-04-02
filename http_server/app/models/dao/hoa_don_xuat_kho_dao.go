package dao

import (
	"admin-v1/app/enums/datetime"
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
	"time"

	"gorm.io/gorm"
)

func CreateExportInvoice(req *requests.Hoa_don_xuat_kho_create, res *responses.Hoa_don_xuat_kho_create) error {
	var chi_tiet_hoa_don_xuat_kho []db.Chi_tiet_hoa_don_xuat_kho

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	var sp_ids 		[]int
	var ctsp_ids	[]int

	//tao 1 danh sach doi tuong cthdxk neu co danh sach tu client truyen len
	if len(req.Chi_tiet_hoa_don_xuat_kho) > 0 {
		for _, hdxk := range req.Chi_tiet_hoa_don_xuat_kho {
			sp_ids = append(sp_ids, hdxk.San_pham_id)
			ctsp_ids = append(ctsp_ids, hdxk.Ctsp_id)

			for _, sku := range hdxk.Ds_sku {
				//them tung cthdxk vao ds
				chi_tiet_hoa_don_xuat_kho = append(chi_tiet_hoa_don_xuat_kho, db.Chi_tiet_hoa_don_xuat_kho{	
					Ctsp_id: 		uint(hdxk.Ctsp_id),
					Sku: 			sku.Sku,
					Don_vi_tinh: 	hdxk.Don_vi_tinh,	
					So_luong_ban: 	sku.So_luong_ban,
					Gia_ban: 		sku.Gia_ban,
					Chiet_khau: 	hdxk.Chiet_khau,
					Thanh_tien: 	hdxk.Thanh_tien,
					// Gia_nhap: 		sku.Gia_nhap,
					Loi_nhuan: 		hdxk.Loi_nhuan,
					La_qua_tang: 	hdxk.La_qua_tang,
				})
			}
		}

		//kiem tra cac id cua san pham va ctsp co ton tai hay khong
		var count int64 = 0

		if err := tx.Debug().Table("san_pham").Where("id IN ?", sp_ids).Count(&count).Error; err != nil {
			tx.Rollback()
			return errors.New("loi khi kiem tra san pham: " + err.Error())
		}

		if count != int64(len(sp_ids)) {
			tx.Rollback()
			return errors.New("co id san pham khong ton tai")
		}

		count = 0

		if err := tx.Debug().Table("chi_tiet_san_pham").Where("id IN ?", ctsp_ids).Count(&count).Error; err != nil {
			tx.Rollback()
			return errors.New("loi khi kiem tra chi tiet san pham: " + err.Error())
		}

		if count != int64(len(ctsp_ids)) {
			tx.Rollback()
			return errors.New("co id chi tiet san pham khong ton tai")
		}
	}

	var hdxk_count_in_year int64 = 0

	if err := tx.Table("hoa_don_xuat_kho").
		Where("YEAR(created_at) = ?", helpers.GetCurrentTimeVN().Format(datetime.Year)).
		Count(&hdxk_count_in_year).Error;
	err != nil {
		tx.Rollback()
		return errors.New("loi khi tinh toan so luong hoa don trong nam: " + err.Error())
	}

	ngayXuat, err := time.Parse(time.RFC3339, req.Ngay_xuat)

	if err != nil {
		tx.Rollback()
		return errors.New("co loi khi chuyen doi thoi gian xuat: " + err.Error())
	}


	//tao doi tuong hdxk
	var hoa_don_xuat_kho = db.Hoa_don_xuat_kho{
		So_hoa_don: 			int(hdxk_count_in_year),					
		Ma_hoa_don: 			"HDX" + string(hdxk_count_in_year),
		Khach_hang_id: 			req.Khach_hang_id,
		Nhan_vien_sale_id: 		req.Nhan_vien_sale_id,
		Nhan_vien_giao_hang_id: req.Nhan_vien_giao_hang_id,
		Ngay_xuat: 				ngayXuat.Format(datetime.DateTime),
		Tong_tien: 				req.Tong_tien,	
		Vat: 					req.Vat,	
		Thanh_tien: 			req.Thanh_tien,				
		Tra_truoc: 				req.Tra_truoc,
		Con_lai: 				req.Tong_tien - req.Tra_truoc,
		Tong_gia_nhap: 			req.Tong_gia_nhap,		
		Loi_nhuan: 				req.Loi_nhuan,
		Ghi_chu: 				req.Ghi_chu,
		Da_giao_hang: 			req.Da_giao_hang,		
		Loai_chiet_khau: 		req.Loai_chiet_khau,
		Gia_tri_chiet_khau: 	req.Gia_tri_chiet_khau,
		Khoa_don: 				false,

		Chi_tiet_hoa_don_xuat_kho: chi_tiet_hoa_don_xuat_kho,
	}

	//insert hdxk
	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Create(&hoa_don_xuat_kho).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the tao hoa don xuat kho: " + err.Error())
	}

	//cap nhat lai chi tiet san pham: so luong, gia nhap, gia ban
	if len(hoa_don_xuat_kho.Chi_tiet_hoa_don_xuat_kho) > 0 {
		var ids_Ctsp []uint
		var chi_tiet_san_pham []db.Chi_tiet_san_pham

		for _, value := range hoa_don_xuat_kho.Chi_tiet_hoa_don_xuat_kho {
			ids_Ctsp = append(ids_Ctsp, value.Ctsp_id)
		}

		//lay ra ds ctsp theo ds id
		if err := tx.Debug().Table("chi_tiet_san_pham").
			Where("ID IN ?", ids_Ctsp).
			Find(&chi_tiet_san_pham).Error; 
		err != nil {
			tx.Rollback()
			return errors.New("khong the tim thay chi tiet san pham: " + err.Error())
		}

		//loc qua ds ctsp, cap nhat lai thong tin, tao ds ton kho
		for idx, value := range hoa_don_xuat_kho.Chi_tiet_hoa_don_xuat_kho {
			chi_tiet_san_pham[idx].So_luong = chi_tiet_san_pham[idx].So_luong - value.So_luong_ban
			// chi_tiet_san_pham[idx].Gia_nhap = value.Gia_nhap * ((100 - value.Chiet_khau) / 100)
			// chi_tiet_san_pham[idx].Gia_ban = value.Gia_ban

			//cap nhat so luong ton kho
			var ton_kho db.Ton_kho

			if err := tx.Table("ton_kho").
				Where("sku = ?", value.Sku).
				Find(&ton_kho).Error;
			err != nil {
				tx.Rollback()
				return errors.New("co loi khi truy xuat ton kho theo sku: " + err.Error())
			}

			ton_kho.So_luong_ton = ton_kho.So_luong_ton - value.So_luong_ban

			if err := tx.Save(&ton_kho).Error; err != nil {
				tx.Rollback()
				return errors.New("co loi khi luu ton kho da cap nhat: " + err.Error())
			}
		}

		//update ctsp: bulk update
		if err := tx.Debug().Table("chi_tiet_san_pham").Save(&chi_tiet_san_pham).Error; err != nil {
			tx.Rollback()
			return errors.New("khong the cap nhat chi tiet san pham: " + err.Error())
		}
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	res.Hoa_don_xuat_kho = hoa_don_xuat_kho

	return nil
}

func UpdateExportInvoiceExec(req *requests.Hoa_don_xuat_kho_update) error {
	var hdxk db.Hoa_don_xuat_kho

	tx := helpers.GormDB.Begin()

	//kiem tra hoa don xuat kho ton tai
	if err := tx.Table("hoa_don_xuat_kho").
		Where("id = ?", req.Hoa_don_id).
		First(&hdxk).Error;
	err != nil {
		tx.Rollback()
		return errors.New("hoa don xuat kho khong ton tai: " + err.Error())
	}

	//kiem tra hoa don co cho phep chinh sua
	if hdxk.Khoa_don == true {
		tx.Rollback()
		return errors.New("hoa don bi khoa: khong duoc chinh sua")
	}

	//kiem tra tien tra truoc hop le
	if req.Tra_truoc > hdxk.Thanh_tien {
		tx.Rollback()
		return errors.New("so tien tra truoc khong hop le: > thanh tien hoa don")
	}

	//convert ngay xuat sang time voi layout datetime
	ngayNhap, err := time.Parse(time.RFC3339, req.Ngay_xuat)

	if err != nil {
		tx.Rollback()
		return errors.New("co loi khi convert ngay xuat sang time")
	}

	hdxk.Tra_truoc 				= req.Tra_truoc
	hdxk.Ngay_xuat	 			= ngayNhap.Format(datetime.DateTime)
	hdxk.Ghi_chu 				= req.Ghi_chu
	hdxk.Con_lai 				= (hdxk.Tong_tien * ((100 - req.Vat) / 100)) - req.Tra_truoc
	hdxk.Khach_hang_id 			= req.Khach_hang_id
	hdxk.Nhan_vien_sale_id 		= req.Nhan_vien_sale_id
	hdxk.Nhan_vien_giao_hang_id = req.Nhan_vien_giao_hang_id
	hdxk.Vat 					= req.Vat
	hdxk.Loai_chiet_khau 		= req.Loai_chiet_khau
	hdxk.Gia_tri_chiet_khau 	= req.Gia_tri_chiet_khau
	hdxk.Da_giao_hang 			= req.Da_giao_hang

	//update hoa don xuat kho
	if err := tx.Save(&hdxk).Error; err != nil {
		tx.Rollback()
		return errors.New("loi khi cap nhat hoa don xuat kho: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func LockExportInvoiceExec(req *requests.Hoa_don_xuat_kho_lock) error {
	var hdxk db.Hoa_don_xuat_kho

	tx := helpers.GormDB.Begin()

	//kiem tra hoa don xuat kho ton tai
	if err := tx.Table("hoa_don_xuat_kho").
		Where("id = ?", req.Hoa_don_id).
		First(&hdxk).Error;
	err != nil {
		tx.Rollback()
		return errors.New("hoa don xuat kho khong ton tai: " + err.Error())
	}

	if req.Lock_or_open == "lock" {
		hdxk.Khoa_don = true
	} else {
		hdxk.Khoa_don = false
	}

	//update hoa don xuat kho 
	if err := tx.Save(&hdxk).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the lock hoa don: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}


func ExportDebtPaymentExec(req *requests.Tra_no_xuat_kho_request) error {
	var hdxk db.Hoa_don_xuat_kho

	tx := helpers.GormDB.Begin()

	//kiem tra hoa don xuat kho ton tai
	if err := tx.Table("hoa_don_xuat_kho").
		Where("id = ?", req.Hoa_don_id).
		First(&hdxk).Error;
	err != nil {
		tx.Rollback()
		return errors.New("hoa don xuat kho khong ton tai: " + err.Error())
	}

	//kiem tra tien tra no hop le
	if hdxk.Con_lai <= req.Tien_tra {
		return errors.New("tien tra no khong hop le: > tien con lai")
	}

	hdxk.Con_lai = hdxk.Con_lai - req.Tien_tra

	//update hoa don xuat kho 
	if err := tx.Save(&hdxk).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the updade tien no hoa don: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func ReturnExportProductExec(req *requests.Tra_hang_xuat_kho_request) error {
	var hdxk db.Hoa_don_xuat_kho

	tx := helpers.GormDB.Begin()

	//kiem tra hoa don xuat kho ton tai
	if err := tx.Table("hoa_don_xuat_kho").
		Where("id = ?", req.Hoa_don_id).
		Find(&hdxk).Error;
	err != nil {
		tx.Rollback()
		return errors.New("hoa don xuat kho khong ton tai: " + err.Error())
	}

	for _, value := range req.Ds_san_pham_tra {
		var cthd_xuat_kho 	db.Chi_tiet_hoa_don_xuat_kho
		var ton_kho			db.Ton_kho

		//kiem tra chi tiet hoa don xuat kho va ton kho ton tai
		if err := tx.Table("chi_tiet_hoa_don_xuat_kho").
			Where("id = ?", value.Cthd_xuat_kho_id).
			Find(&cthd_xuat_kho).Error;
		err != nil {
			tx.Rollback()
			return errors.New("chi tiet hoa don xuat kho khong ton tai: " + err.Error())
		}

		if err := tx.Table("ton_kho").
			Where("id = ?", value.Sku).
			Find(&ton_kho).Error;
		err != nil {
			tx.Rollback()
			return errors.New("ton kho khong ton tai: " + err.Error())
		}

		//neu so luong tra lon hon so luong da xuat se bao loi
		if cthd_xuat_kho.So_luong_ban < value.So_luong_tra {
			return errors.New("so luong tra hang khong hop le: > so luong ban")
		}

		//neu so luong tra != 0 se tien hanh cap nhat
		if value.So_luong_tra != 0 {
			//neu so luong tra == so luong ton kho sex tien hanh cap nhat sl ton kho = 0 va xoa chi tiet hdxk
			if value.So_luong_tra == cthd_xuat_kho.So_luong_ban {
				ton_kho.So_luong_ton = 0

				if err := tx.Debug().Save(&ton_kho).Error; err != nil {
					tx.Rollback()
					return errors.New("cap nhat ton kho gap loi: " + err.Error())
				}

				if err := tx.Debug().Table("chi_tiet_hoa_don_xuat_kho").
					Where("id = ?", cthd_xuat_kho.ID).
					Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error;
				err != nil {
					tx.Rollback()
					return errors.New("xoa chi tiet hoa don xuat kho gap loi: " + err.Error())
				}
			} else {
				ton_kho.So_luong_ton = ton_kho.So_luong_ton + value.So_luong_tra
				cthd_xuat_kho.So_luong_ban = cthd_xuat_kho.So_luong_ban + value.So_luong_tra

				if err := tx.Debug().Save(&ton_kho).Error; err != nil {
					tx.Rollback()
					return errors.New("cap nhat ton kho gap loi: " + err.Error())
				}

				if err := tx.Debug().Save(&cthd_xuat_kho).Error; err != nil {
					tx.Rollback()
					return errors.New("cap nhat chi tiet hoa don xuat kho gap loi: " + err.Error())
				}
			}
		}
	}

	return nil
}