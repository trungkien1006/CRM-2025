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

func CreateImportInvoice(req *requests.Hoa_don_nhap_kho_create, res *responses.Hoa_don_nhap_kho_create) error {
	var chi_tiet_hoa_don_nhap_kho []db.Chi_tiet_hoa_don_nhap_kho

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//tao 1 danh sach doi tuong cthdnk neu co danh sach tu client truyen len
	if len(req.Ds_san_pham_nhap) > 0 {
		//tao tham so date 1 phan cua sku
		var date = helpers.GetCurrentTimeVN().Format(datetime.Date)

		//tao 2 ds chua cac id de kiem tra ton tai hay chua
		var ds_san_pham_id 	[]int64
		var ds_ctsp_id		[]int64

		for _, value := range req.Ds_san_pham_nhap {
			//tao tham so counter 1 phan cua sku
			var counter int64

			if err := tx.Table("chi_tiet_hoa_don_nhap_kho").
				Where("ctsp_id = ?", value.Ctsp_id).
				Where("DATE(created_at) = ?", date).
				Count(&counter).Error;
			err != nil {
				return errors.New("khong the tim chi tiet hoa don nhap kho: " + err.Error())
			}

			hanSuDung, err := time.Parse(time.RFC3339, value.Han_su_dung)

			if err != nil {
				tx.Rollback()
				return errors.New("co loi khi chuyen doi han su dung: " + err.Error())
			}

			//them tung cthdnk vao ds
			chi_tiet_hoa_don_nhap_kho = append(chi_tiet_hoa_don_nhap_kho, db.Chi_tiet_hoa_don_nhap_kho{
				San_pham_id: uint(value.San_pham_id),
				Ctsp_id: uint(value.Ctsp_id),
				Sku: helpers.GenerateSKU(value.Upc, value.Ctsp_id, counter),
				So_luong: value.So_luong,
				Don_vi_tinh: value.Don_vi_tinh,
				Ke: value.Ke,
				Gia_nhap: value.Gia_nhap,
				Gia_ban: value.Gia_ban,
				Chiet_khau: value.Chiet_khau,
				Thanh_tien: value.Thanh_tien,
				La_qua_tang: value.La_qua_tang,
				Han_su_dung: hanSuDung.Format(datetime.DateTime),
			})
		}

		//kiem tra cac id cua san pham va ctsp co ton tai hay khong
		var count int64 = 0

		if err := tx.Debug().Table("san_pham").Where("id IN ?", ds_san_pham_id).Count(&count).Error; err != nil {
			tx.Rollback()
			return errors.New("loi khi kiem tra san pham: " + err.Error())
		}

		if count != int64(len(ds_san_pham_id)) {
			tx.Rollback()
			return errors.New("co id san pham khong ton tai")
		}

		count = 0

		if err := tx.Debug().Table("chi_tiet_san_pham").Where("id IN ?", ds_ctsp_id).Count(&count).Error; err != nil {
			tx.Rollback()
			return errors.New("loi khi kiem tra chi tiet san pham: " + err.Error())
		}

		if count != int64(len(ds_ctsp_id)) {
			tx.Rollback()
			return errors.New("co id chi tiet san pham khong ton tai")
		}
	}

	var hdnk_count_in_year int64 = 0

	if err := tx.Table("hoa_don_nhap_kho").
		Where("YEAR(created_at) = ?", helpers.GetCurrentTimeVN().Format(datetime.Year)).
		Count(&hdnk_count_in_year).Error;
	err != nil {
		tx.Rollback()
		return errors.New("loi khi tinh toan so luong hoa don trong nam: " + err.Error())
	}

	ngayNhap, err := time.Parse(time.RFC3339, req.Ngay_nhap)

	if err != nil {
		tx.Rollback()
		return errors.New("co loi khi chuyen doi thoi gian nhap: " + err.Error())
	}

	//tao doi tuong hdnk
	var hoa_don_nhap_kho = db.Hoa_don_nhap_kho{
		So_hoa_don: int(hdnk_count_in_year),
		Ma_hoa_don: "HDN" + string(hdnk_count_in_year),
		Nha_phan_phoi_id: req.Nha_phan_phoi_id,	
		Kho_id: req.Kho_id,
		Ngay_nhap: ngayNhap.Format(datetime.DateTime),
		Tong_tien: req.Tong_tien,
		Tra_truoc: req.Tra_truoc,
		Con_lai: req.Tong_tien - req.Tra_truoc,
		Ghi_chu: req.Ghi_chu,
		Khoa_don: false,

		Chi_tiet_hoa_don_nhap_kho: chi_tiet_hoa_don_nhap_kho,
	}

	//insert hdnk
	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Create(&hoa_don_nhap_kho).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the tao hoa don nhap kho: " + err.Error())
	}

	//cap nhat lai chi tiet san pham: so luong, gia nhap, gia ban
	if len(hoa_don_nhap_kho.Chi_tiet_hoa_don_nhap_kho) > 0 {
		var ids_Ctsp []uint
		var chi_tiet_san_pham []db.Chi_tiet_san_pham

		for _, value := range hoa_don_nhap_kho.Chi_tiet_hoa_don_nhap_kho {
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

		var ds_ton_kho []db.Ton_kho

		//loc qua ds ctsp, cap nhat lai thong tin, tao ds ton kho
		for idx, value := range hoa_don_nhap_kho.Chi_tiet_hoa_don_nhap_kho {
			chi_tiet_san_pham[idx].So_luong = chi_tiet_san_pham[idx].So_luong + value.So_luong
			chi_tiet_san_pham[idx].Gia_nhap = value.Gia_nhap * ((100 - value.Chiet_khau) / 100)
			chi_tiet_san_pham[idx].Gia_ban = value.Gia_ban

			ds_ton_kho = append(ds_ton_kho, db.Ton_kho{
				San_pham_id: uint(value.San_pham_id),
				Ctsp_id: uint(value.Ctsp_id),
				Sku: value.Sku,
				So_luong_ton: value.So_luong,
			})
		}

		//update ctsp: bulk update
		if err := tx.Debug().Model(&db.Chi_tiet_san_pham{}).Save(&chi_tiet_san_pham).Error; err != nil {
			tx.Rollback()
			return errors.New("khong the cap nhat chi tiet san pham: " + err.Error())
		}

		//insert ton kho
		if err := tx.Debug().Model(&db.Ton_kho{}).Create(&ds_ton_kho).Error; err != nil {
			tx.Rollback()
			return errors.New("khong the them danh sach ton kho: " + err.Error())
		}
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	res.Hoa_don_nhap_kho = hoa_don_nhap_kho

	return nil
}

func UpdateImportInvoiceExec(req *requests.Hoa_don_nhap_kho_update) error {
	var hdnk db.Hoa_don_nhap_kho

	tx := helpers.GormDB.Begin()

	//kiem tra hoa don nhap kho ton tai
	if err := tx.Table("hoa_don_nhap_kho").
		Where("id = ?", req.Hoa_don_id).
		Find(&hdnk).Error;
	err != nil {
		tx.Rollback()
		return errors.New("hoa don nhap kho khong ton tai: " + err.Error())
	}


	//kiem tra tien tra truoc hop le
	if req.Tra_truoc > hdnk.Tong_tien {
		tx.Rollback()
		return errors.New("so tien tra truoc khong hop le: > tong tien hoa don")
	}

	//convert ngay nhap sang time voi layout datetime
	ngayNhap, err := time.Parse(time.RFC3339, req.Ngay_nhap)

	if err != nil {
		tx.Rollback()
		return errors.New("co loi khi convert ngay nhap sang time")
	}

	hdnk.Tra_truoc = req.Tra_truoc
	hdnk.Ngay_nhap = ngayNhap.Format(datetime.DateTime)
	hdnk.Ghi_chu = req.Ghi_chu
	hdnk.Con_lai = hdnk.Tong_tien - req.Tra_truoc

	//update hoa don nhap kho
	if err := tx.Save(&hdnk).Error; err != nil {
		tx.Rollback()
		return errors.New("loi khi cap nhat hoa don nhap kho: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func LockImportInvoiceExec(req *requests.Hoa_don_nhap_kho_lock) error {
	var hdnk db.Hoa_don_nhap_kho

	tx := helpers.GormDB.Begin()

	//kiem tra hoa don nhap kho ton tai
	if err := tx.Debug().Table("hoa_don_nhap_kho").
		Where("id = ?", req.Hoa_don_id).
		First(&hdnk).Error;
	err != nil {
		tx.Rollback()
		return errors.New("hoa don nhap kho khong ton tai: " + err.Error())
	}

	if req.Lock_or_open == "lock" {
		hdnk.Khoa_don = true
	} else {
		hdnk.Khoa_don = false
	}

	//update hoa don nhap kho 
	if err := tx.Debug().Save(&hdnk).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the updade tien no hoa don: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func ImportDebtPaymentExec(req *requests.Tra_no_nhap_kho_request) error {
	var hdnk db.Hoa_don_nhap_kho

	tx := helpers.GormDB.Begin()

	//kiem tra hoa don nhap kho ton tai
	if err := tx.Debug().Table("hoa_don_nhap_kho").
		Where("id = ?", req.Hoa_don_id).
		First(&hdnk).Error;
	err != nil {
		tx.Rollback()
		return errors.New("hoa don nhap kho khong ton tai: " + err.Error())
	}

	//kiem tra hoa don co cho phep chinh sua
	if hdnk.Khoa_don == true {
		tx.Rollback()
		return errors.New("hoa don bi khoa: khong duoc chinh sua")
	}

	//kiem tra tien tra no hop le
	if hdnk.Con_lai <= req.Tien_tra {
		return errors.New("tien tra no khong hop le: > tien con lai")
	}

	hdnk.Con_lai = hdnk.Con_lai - req.Tien_tra

	//update hoa don nhap kho 
	if err := tx.Debug().Save(&hdnk).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the updade tien no hoa don: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func ReturnImportProductExec(req *requests.Tra_hang_nhap_kho_request) error {
	var hdnk db.Hoa_don_nhap_kho

	tx := helpers.GormDB.Begin()

	//kiem tra hoa don nhap kho ton tai
	if err := tx.Debug().Table("hoa_don_nhap_kho").
		Where("id = ?", req.Hoa_don_id).
		First(&hdnk).Error;
	err != nil {
		tx.Rollback()
		return errors.New("hoa don nhap kho khong ton tai: " + err.Error())
	}

	for _, value := range req.Ds_san_pham_tra {
		var cthd_nhap_kho 	db.Chi_tiet_hoa_don_nhap_kho
		var ton_kho			db.Ton_kho

		//kiem tra chi tiet hoa don nhap kho va ton kho ton tai
		if err := tx.Debug().Table("chi_tiet_hoa_don_nhap_kho").
			Where("id = ?", value.Cthd_nhap_kho_id).
			Find(&cthd_nhap_kho).Error;
		err != nil {
			tx.Rollback()
			return errors.New("chi tiet hoa don nhap kho khong ton tai: " + err.Error())
		}

		if err := tx.Debug().Table("ton_kho").
			Where("id = ?", value.Sku).
			Find(&ton_kho).Error;
		err != nil {
			tx.Rollback()
			return errors.New("ton kho khong ton tai: " + err.Error())
		}

		//neu so luong tra lon hon so luong ton kho se bao loi
		if cthd_nhap_kho.So_luong < value.So_luong_tra || ton_kho.So_luong_ton < value.So_luong_tra {
			return errors.New("so luong tra hang khong hop le: > so luong ton kho")
		}

		//neu so luong tra != 0 se tien hanh cap nhat
		if value.So_luong_tra != 0 {
			//neu so luong tra == so luong ton kho sex tien hanh cap nhat sl ton kho = 0 va xoa chi tiet hdnk
			if value.So_luong_tra == cthd_nhap_kho.So_luong {
				ton_kho.So_luong_ton = 0

				if err := tx.Debug().Save(&ton_kho).Error; err != nil {
					tx.Rollback()
					return errors.New("cap nhat ton kho gap loi: " + err.Error())
				}

				if err := tx.Debug().Table("chi_tiet_hoa_don_nhap_kho").
					Where("id = ?", cthd_nhap_kho.ID).
					Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error;
				err != nil {
					tx.Rollback()
					return errors.New("xoa chi tiet hoa don nhap kho gap loi: " + err.Error())
				}
			} else {
				ton_kho.So_luong_ton = ton_kho.So_luong_ton - value.So_luong_tra
				cthd_nhap_kho.So_luong = cthd_nhap_kho.So_luong - value.So_luong_tra

				if err := tx.Debug().Save(&ton_kho).Error; err != nil {
					tx.Rollback()
					return errors.New("cap nhat ton kho gap loi: " + err.Error())
				}

				if err := tx.Debug().Save(&cthd_nhap_kho).Error; err != nil {
					tx.Rollback()
					return errors.New("cap nhat chi tiet hoa don nhap kho gap loi: " + err.Error())
				}
			}
		}
	}

	return nil
}

