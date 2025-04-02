package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateDiscountTypeExec(req *requests.Loai_giam_gia_create, res *responses.Loai_giam_gia_create) error {
	//kiem tra ten loai giam gia ton tai
	if result := helpers.GormDB.Debug().
		Table("loai_giam_gia").
		Where("ten = ?", req.Ten).
		First(&res.Loai_giam_gia);
	result.RowsAffected > 0 {
		return errors.New("loai giam gia da ton tai")
	}
	
	var loai_giam_gia = db.Loai_giam_gia{
		Ten: req.Ten,
		Gia_tri: req.Gia_tri,
	}

	//insert lgg
	if err := helpers.GormDB.Debug().Create(&loai_giam_gia).Error; err != nil {
		return errors.New("khong the tao loai giam gia: " + err.Error())
	}

	res.Loai_giam_gia = loai_giam_gia

	return nil
}

func UpdateDiscountTypeExec(req *requests.Loai_giam_gia_update) error {
	var loai_giam_gia db.Loai_giam_gia
	var loai_giam_gia_temp db.Loai_giam_gia

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra ten loai san pham ton tai
	if result := tx.Debug().
		Table("loai_giam_gia").
		Where("ten = ?", req.Ten).
		Where("ID != ?", req.Id).
		First(&loai_giam_gia_temp);
	result.RowsAffected > 0 {
		tx.Rollback()
		return errors.New("ten loai giam gia da ton tai")
	}

	//kiem tra loai giam gia ton tai
	if result := tx.Debug().
		Table("loai_giam_gia").
		Where("id = ?", req.Id).
		First(&loai_giam_gia);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("loai giam gia khong ton tai")
	}

	loai_giam_gia.Ten = req.Ten
	loai_giam_gia.Gia_tri = req.Gia_tri

	//update loai giam gia
	if err := tx.Model(&loai_giam_gia).Debug().Updates(&loai_giam_gia).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat loai giam gia: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteDiscountTypeExec(req *requests.Loai_giam_gia_delete) error {
	var loai_giam_gia db.Loai_giam_gia

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra loai giam gia ton tai
	if result := tx.Debug().
		Table("loai_giam_gia").
		Where("id = ?", req.Id).
		First(&loai_giam_gia);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("loai giam gia khong ton tai")
	}

	//kiem tra loai giam gia ton tai trong san pham
	var count int64 = 0

	if err := tx.Table("san_pham").Where("loai_giam_gia_id = ?", req.Id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the kiem tra loai giam gia trong san pham: " + err.Error())
	}

	if count != 0 {
		tx.Rollback()
		return errors.New("khong the xoa vi loai giam gia ton tai trong san pham")
	}

	//delete loai giam gia
	if err := tx.Model(&loai_giam_gia).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the xoa loai giam gia: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}