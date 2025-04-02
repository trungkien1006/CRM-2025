package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateRoleExec(req *requests.Chuc_vu_create, res *responses.Chuc_vu_create) error {
	//kiem tra ten chuc vu da ton tai
	if result := helpers.GormDB.Debug().
		Table("chuc_vu").
		Where("ten = ?", req.Ten).
		First(&res.Chuc_vu);
	result.RowsAffected > 0 {
		return errors.New("chuc vu da ton tai")
	}
	
	var chuc_vu = db.Chuc_vu{
		Ten: req.Ten,
	}

	//insert chuc vu
	if err := helpers.GormDB.Debug().Create(&chuc_vu).Error; err != nil {
		return errors.New("khong the tao chuc vu: " + err.Error())
	}

	res.Chuc_vu = chuc_vu

	return nil
}

func UpdateRoleExec(req *requests.Chuc_vu_update) error {
	var chuc_vu db.Chuc_vu
	var chuc_vu_temp db.Chuc_vu

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra ten chuc vu ton tai
	if result := tx.Debug().
		Table("chuc_vu").
		Where("ten = ?", req.Ten).
		Where("ID != ?", req.Id).
		First(&chuc_vu_temp);
	result.RowsAffected > 0 {
		return errors.New("ten chuc vu da ton tai")
	}

	//kiem tra chuc vu ton tai
	if result := tx.Debug().
		Table("chuc_vu").
		Where("id = ?", req.Id).
		First(&chuc_vu);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("chuc vu khong ton tai")
	}

	chuc_vu.Ten = req.Ten

	//update chuc vu
	if err := tx.Model(&chuc_vu).Debug().Updates(&chuc_vu).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat chuc vu: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteRoleExec(req *requests.Chuc_vu_delete) error {
	var chuc_vu db.Chuc_vu

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra chuc vu ton tai
	if result := tx.Debug().
		Table("chuc_vu").
		Where("id = ?", req.Id).
		First(&chuc_vu);
	result.RowsAffected == 0 {
		return errors.New("chuc vu khong ton tai")
	}

	//kiem tra chuc vu da duoc su dung
	var count int64 = 0

	if err := tx.Table("nhan_vien").Where("chuc_vu_id = ?", req.Id).Count(&count).Error; err != nil {
		return errors.New("kiem tra chuc vu gap loi: " + err.Error())
	}

	if count != 0 {
		return errors.New("chuc vu da duoc su dung khong the xoa")
	}

	//delete chuc vu
	if err := tx.Model(&chuc_vu).Debug().Delete(&db.Chuc_vu{}, req.Id).Error; err != nil {
		return errors.New("khong the xoa chuc vu: " + err.Error())
	}

	//xoa cac quyen lien quan chuc vu
	if err := tx.Where("chuc_vu_id = ?", req.Id).Delete(&db.Quyen{}).Error; err != nil {
		return errors.New("khong the xoa quyen: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}