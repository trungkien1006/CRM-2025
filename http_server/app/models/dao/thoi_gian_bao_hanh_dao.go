package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateWarrantyTimeExec(req *requests.Thoi_gian_bao_hanh_create, res *responses.Thoi_gian_bao_hanh_create) error {
	//kiem tra ten tgbh ton tai
	if result := helpers.GormDB.Debug().
		Table("thoi_gian_bao_hanh").
		Where("ten = ?", req.Ten).
		First(&res.Thoi_gian_bao_hanh);
	result.RowsAffected > 0 {
		return errors.New("thoi gian bao hanh da ton tai")
	}
	
	var thoi_gian_bao_hanh = db.Thoi_gian_bao_hanh{
		Ten: req.Ten,
	}

	//insert tgbh
	if err := helpers.GormDB.Debug().Create(&thoi_gian_bao_hanh).Error; err != nil {
		return errors.New("khong the tao thoi gian bao hanh: " + err.Error())
	}

	res.Thoi_gian_bao_hanh = thoi_gian_bao_hanh

	return nil
}

func UpdateWarrantyTimeExec(req *requests.Thoi_gian_bao_hanh_update) error {
	var thoi_gian_bao_hanh db.Thoi_gian_bao_hanh
	var thoi_gian_bao_hanh_temp db.Thoi_gian_bao_hanh

	//bat dau transaction
	tx := helpers.GormDB.Begin()
	
	//kiem tra ten tgbh ton tai
	if result := tx.Debug().
		Table("thoi_gian_bao_hanh").
		Where("ten = ?", req.Ten).
		Where("ID != ?", req.Id).
		First(&thoi_gian_bao_hanh_temp);
	result.RowsAffected > 0 {
		tx.Rollback()
		return errors.New("ten thoi gian bao hanh da ton tai")
	}

	//kiem tra tgbh ton tai
	if result := tx.Debug().
		Table("thoi_gian_bao_hanh").
		Where("id = ?", req.Id).
		First(&thoi_gian_bao_hanh);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("thoi gian bao hanh khong ton tai")
	}

	thoi_gian_bao_hanh.Ten = req.Ten

	//update tgbh
	if err := tx.Model(&thoi_gian_bao_hanh).Debug().Update("ten", thoi_gian_bao_hanh.Ten).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat thoi gian bao hanh: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteWarrantyTimeExec(req *requests.Thoi_gian_bao_hanh_delete) error {
	var thoi_gian_bao_hanh db.Thoi_gian_bao_hanh

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra thoi gian bao hanh ton tai
	if result := tx.Debug().
		Table("thoi_gian_bao_hanh").
		Where("id = ?", req.Id).
		First(&thoi_gian_bao_hanh);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("thoi gian bao hanh khong ton tai")
	}

	//kiem tra tgbh ton tai trong san pham
	var count int64 = 0

	if err := tx.Table("san_pham").Where("thoi_gian_bao_hanh_id = ?", req.Id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the kiem tra thoi gian bao hanh trong san pham: " + err.Error())
	}

	if count != 0 {
		tx.Rollback()
		return errors.New("khong the xoa vi thoi gian bao hanh ton tai trong san pham")
	}

	//delete thoi gian bao hanh
	if err := tx.Model(&thoi_gian_bao_hanh).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the xoa thoi gian bao hanh: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}