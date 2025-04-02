package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateUnitExec(req *requests.Don_vi_tinh_create, res *responses.Don_vi_tinh_create) error {
	//kiem tra don vi tinh ton tai
	if result := helpers.GormDB.Debug().
		Table("don_vi_tinh").
		Where("ten = ?", req.Ten).
		First(&res.Don_vi_tinh);
	result.RowsAffected > 0 {
		return errors.New("don vi tinh da ton tai")
	}
	
	var don_vi_tinh = db.Don_vi_tinh{
		Ten: req.Ten,
	}

	//insertt don vi tinh
	if err := helpers.GormDB.Debug().Create(&don_vi_tinh).Error; err != nil {
		return errors.New("khong the tao don vi tinh: " + err.Error())
	}

	res.Don_vi_tinh = don_vi_tinh

	return nil
}

func UpdateUnitExec(req *requests.Don_vi_tinh_update) error {
	var don_vi_tinh db.Don_vi_tinh
	var don_vi_tinh_temp db.Don_vi_tinh

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra ten don vi tinh ton tai
	if result := tx.Debug().
		Table("don_vi_tinh").
		Where("ten = ?", req.Ten).
		Where("ID != ?", req.Id).
		First(&don_vi_tinh_temp);
	result.RowsAffected > 0 {
		return errors.New("ten don vi tinh da ton tai")
	}

	//kiem tra don vi tinh ton tai
	if result := tx.Debug().
		Table("don_vi_tinh").
		Where("id = ?", req.Id).
		First(&don_vi_tinh);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("don vi tinh khong ton tai")
	}

	don_vi_tinh.Ten = req.Ten

	//update don vi tinh
	if err := tx.Model(&don_vi_tinh).Debug().Updates(&don_vi_tinh).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat don vi tinh: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteUnitExec(req *requests.Don_vi_tinh_delete) error {
	var don_vi_tinh db.Don_vi_tinh

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra don vi tinh ton tai
	if result := tx.Debug().
		Table("don_vi_tinh").
		Where("id = ?", req.Id).
		First(&don_vi_tinh);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("don vi tinh khong ton tai")
	}

	//kiem tra don vi tinh da duoc su dung
	var count int64 = 0

	if err := tx.Table("san_pham").Where("don_vi_tinh_id = ?", req.Id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the kiem tra don vi tinh trong san pham: " + err.Error())
	}

	if count != 0 {
		tx.Rollback()
		return errors.New("don vi tinh da duoc su dung trong san pham")
	}

	//delete don vi tinh
	if err := tx.Model(&don_vi_tinh).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the xoa don vi tinh: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}