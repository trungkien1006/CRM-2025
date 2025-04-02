package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"

	"gorm.io/gorm"
)

func CreateProductExec(req *requests.San_pham_create, res *responses.San_pham_create) error {
	//kiem tra upc san pham
	if result := helpers.GormDB.Debug().
		Table("san_pham").
		Where("upc = ?", req.Upc).
		First(&res.San_pham);
	result.RowsAffected > 0 {
		return errors.New("ma san pham da ton tai")
	}
	
	//tao danh sach doi tuong chi tiet san pham
	var chi_tiet_san_pham_arr 	[]db.Chi_tiet_san_pham

	if len(req.Chi_tiet_san_pham) > 0 {
		for _, value := range req.Chi_tiet_san_pham {
			chi_tiet_san_pham_arr = append(chi_tiet_san_pham_arr, db.Chi_tiet_san_pham{
				Ten_phan_loai: value.Ten_phan_loai,
				Hinh_anh: value.Hinh_anh,
				Gia_nhap: 0,
				Gia_ban: 0,
				So_luong: 0,
				Trang_thai: value.Trang_thai,
				Khong_phan_loai: 0,
			})
		}

	} else {
		chi_tiet_san_pham_arr = append(chi_tiet_san_pham_arr, db.Chi_tiet_san_pham{
			Ten_phan_loai: "",
			Hinh_anh: "",
			Gia_nhap: 0,
			Gia_ban: 0,
			So_luong: 0,
			Trang_thai: 0,
			Khong_phan_loai: 1,
		})
	}

	//tao doi tuong san pham
	var san_pham = db.San_pham{
		Ten: req.Ten,
		Upc: req.Upc,
		Loai_san_pham_id: req.Loai_san_pham_id,
		Hinh_anh: req.Hinh_anh,
		Don_vi_tinh_id: req.Don_vi_tinh_id,
		Vat: req.Vat,
		Mo_ta: req.Mo_ta,
		Trang_thai: req.Trang_thai,
		Loai_giam_gia_id: req.Loai_giam_gia_id,
		Thoi_gian_bao_hanh_id: req.Thoi_gian_bao_hanh_id,

		Chi_tiet_san_pham: chi_tiet_san_pham_arr,
	}

	//insert san pham
	if err := helpers.GormDB.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Create(&san_pham).Error; err != nil {
		return errors.New("khong the tao loai san pham: " + err.Error())
	}

	res.San_pham = san_pham

	return nil
}

func UpdateProductExec(req *requests.San_pham_update) error {
	var san_pham db.San_pham
	var san_pham_temp db.San_pham

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra upc san pham
	if result := tx.Debug().
		Table("san_pham").
		Where("upc = ?", req.Upc).
		Where("ID != ?", req.Id).
		First(&san_pham_temp);
	result.RowsAffected > 0 {
		tx.Rollback()
		return errors.New("ma san pham da ton tai")
	}

	//kiem tra san pham ton tai
	if result := tx.Debug().
		Table("san_pham").
		Preload("Chi_tiet_san_pham").
		Where("id = ?", req.Id).
		First(&san_pham);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("san pham khong ton tai")
	}

	//kiem tra loai san pham ton tai
	var count int64 = 0

	if err := tx.Table("loai_san_pham").Where("id = ?", req.Loai_san_pham_id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("loi khi kiem tra loai san pham: " + err.Error())
	}

	if count == 0 {
		tx.Rollback()
		return errors.New("loai san pham khong ton tai")
	}

	//kiem tra don vi tinh ton tai
	count = 0

	if err := tx.Table("don_vi_tinh").Where("id = ?", req.Don_vi_tinh_id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("loi khi kiem tra don vi tinh: " + err.Error())
	}

	if count == 0 {
		tx.Rollback()
		return errors.New("don vi tinh khong ton tai")
	}

	//kiem tra loai giam gia ton tai
	count = 0

	if err := tx.Table("loai_giam_gia").Where("id = ?", req.Loai_giam_gia_id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("loi khi kiem tra loai giam gia: " + err.Error())
	}

	if count == 0 {
		tx.Rollback()
		return errors.New("loai giam gia khong ton tai")
	}

	//kiem tra thoi gian bao hanh ton tai
	count = 0

	if err := tx.Table("thoi_gian_bao_hanh").Where("id = ?", req.Thoi_gian_bao_hanh_id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("loi khi kiem tra thoi gian bao hanh: " + err.Error())
	}

	if count == 0 {
		tx.Rollback()
		return errors.New("thoi gian bao hanh khong ton tai")
	}

	//tao danh sach doi tuong ctsp
	var chi_tiet_san_pham_update 	[]db.Chi_tiet_san_pham
	var chi_tiet_san_pham_insert	[]db.Chi_tiet_san_pham
	var ids							[]uint

	//neu la chi tiet san pham moi se them vao bang insert nguoc lai them vao bang update
	if len(req.Chi_tiet_san_pham) > 0 {
		for _, value := range req.Chi_tiet_san_pham {
			if value.Id != 0 {
				ids = append(ids, value.Id)

				if value.Hinh_anh != "" {
					chi_tiet_san_pham_update = append(chi_tiet_san_pham_update, db.Chi_tiet_san_pham{
						ID: value.Id,
						San_pham_id: uint(req.Id),
						Ten_phan_loai: value.Ten_phan_loai,
						Hinh_anh: value.Hinh_anh,
						Gia_nhap: 0,
						Gia_ban: 0,
						So_luong: 0,
						Trang_thai: value.Trang_thai,
						Khong_phan_loai: 0,
					})
				} else {
						chi_tiet_san_pham_update = append(chi_tiet_san_pham_update, db.Chi_tiet_san_pham{
						ID: value.Id,
						San_pham_id: uint(req.Id),
						Ten_phan_loai: value.Ten_phan_loai,
						Gia_nhap: 0,
						Gia_ban: 0,
						So_luong: 0,
						Trang_thai: value.Trang_thai,
						Khong_phan_loai: 0,
					})
				}				
			} else {
				chi_tiet_san_pham_insert = append(chi_tiet_san_pham_insert, db.Chi_tiet_san_pham{
					San_pham_id: uint(req.Id),
					Ten_phan_loai: value.Ten_phan_loai,
					Hinh_anh: value.Hinh_anh,
					Gia_nhap: 0,
					Gia_ban: 0,
					So_luong: 0,
					Trang_thai: value.Trang_thai,
					Khong_phan_loai: 0,
				})
			}
		}
	} else {
		chi_tiet_san_pham_insert = append(chi_tiet_san_pham_insert, db.Chi_tiet_san_pham{
			San_pham_id: uint(req.Id),
			Ten_phan_loai: "",
			Hinh_anh: "",
			Gia_nhap: 0,
			Gia_ban: 0,
			So_luong: 0,
			Trang_thai: 0,
			Khong_phan_loai: 1,
		})
	}

	//kiem tea chi tiet san pham ton tai
	count = 0

	if err := helpers.GormDB.Table("chi_tiet_san_pham").Where("ID IN ?", ids).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("kiem tra chi tiet san pham gap loi: " + err.Error())
	}

	if count != int64(len(ids)) {
		tx.Rollback()
		return errors.New("chi tiet san pham khong ton tai")
	}

	san_pham.Ten = req.Ten
	san_pham.Upc = req.Upc
	san_pham.Loai_san_pham_id = req.Loai_san_pham_id
	san_pham.Hinh_anh = req.Hinh_anh
	san_pham.Don_vi_tinh_id = req.Don_vi_tinh_id
	san_pham.Vat = req.Vat
	san_pham.Mo_ta = req.Mo_ta
	san_pham.Trang_thai = req.Trang_thai
	san_pham.Loai_giam_gia_id = req.Loai_giam_gia_id
	san_pham.Thoi_gian_bao_hanh_id = req.Thoi_gian_bao_hanh_id

	san_pham.Chi_tiet_san_pham = chi_tiet_san_pham_update

	//update san pham
	if err := tx.Model(&san_pham).Debug().Session(&gorm.Session{FullSaveAssociations: true}).Save(&san_pham).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat san pham: " + err.Error())
	}
	
	//delete cac chi tiet san pham
	if len(ids) == 0 {
		if err := tx.Debug().Where("san_pham_id = ?", req.Id).Delete(&db.Chi_tiet_san_pham{}).Error; err != nil {
			tx.Rollback()
			return errors.New("khong the xoa chi tiet san pham: " + err.Error())
		}
	} else {
		if err := tx.Debug().Where("id NOT IN ?", ids).Delete(&db.Chi_tiet_san_pham{}).Error; err != nil {
			tx.Rollback()
			return errors.New("khong the xoa chi tiet san pham: " + err.Error())
		}
	}

	//insert cac chi tiet san pham moi
	if len(chi_tiet_san_pham_insert) > 0 {
		if err := tx.Model(&db.Chi_tiet_san_pham{}).Debug().Create(&chi_tiet_san_pham_insert).Error; err != nil {
			tx.Rollback()
			return errors.New("khong the them chi tiet san pham: " + err.Error())
		}
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteProductExec(req *requests.San_pham_delete) error {
	var san_pham db.San_pham

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra san pham ton tai
	if result := tx.Debug().
		Table("san_pham").
		Where("id = ?", req.Id).
		First(&san_pham);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("san pham khong ton tai")
	}

	//kiem tra san pham thuoc nha phan phoi
	var count int64 = 0

	if err := tx.Table("san_pham_nha_phan_phoi").Where("san_pham_id = ?", req.Id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the kiem tra san pham trong nha phan phoi: " + err.Error())
	}

	if count != 0 {
		tx.Rollback()
		return errors.New("khong the xoa san pham vi lien quan den nha phan phoi")
	}

	//kiem tra san pham thuoc chi tiet san pham
	count = 0

	if err := tx.Table("chi_tiet_san_pham").Where("san_pham_id = ?", req.Id).Where("khong_phan_loai = 0").Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the kiem tra san pham trong chi tiet san pham: " + err.Error())
	}

	if count != 0 {
		tx.Rollback()
		return errors.New("khong the xoa san pham vi lien quan den chi tiet san pham")
	}

	//kiem tra san pham thuoc hoa don nhap kho
	count = 0

	if err := tx.Table("chi_tiet_hoa_don_nhap_kho").Where("san_pham_id = ?", req.Id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the kiem tra san pham trong hoa don nhap kho: " + err.Error())
	}

	if count != 0 {
		tx.Rollback()
		return errors.New("khong the xoa san pham vi lien quan den hoa don nhap kho")
	}

	//delete san pham
	if err := tx.Model(&san_pham).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the xoa san pham: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}