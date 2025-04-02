package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateProductTypeExec(req *requests.Loai_san_pham_create, res *responses.Loai_san_pham_create) error {
	//kiem tra loai san pham ton tai
	if result := helpers.GormDB.Debug().
		Table("loai_san_pham").
		Where("ten = ?", req.Ten).
		First(&res.Loai_san_pham);
	result.RowsAffected > 0 {
		return errors.New("loai san pham da ton tai")
	}
	
	var loai_san_pham = db.Loai_san_pham{
		Ten: req.Ten,
		Hinh_anh: req.Hinh_anh,
	}

	//insert loai san pham
	if err := helpers.GormDB.Debug().Create(&loai_san_pham).Error; err != nil {
		return errors.New("khong the tao loai san pham: " + err.Error())
	}

	res.Loai_san_pham = loai_san_pham

	return nil
}

func UpdateProductTypeExec(req *requests.Loai_san_pham_update) error {
	var loai_san_pham db.Loai_san_pham
	var loai_san_pham_temp db.Loai_san_pham

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra ten loai san pham ton tai
	if result := tx.Debug().
		Table("loai_san_pham").
		Where("ten = ?", req.Ten).
		Where("ID != ?", req.Id).
		First(&loai_san_pham_temp);
	result.RowsAffected > 0 {
		tx.Rollback()
		return errors.New("ten loai san pham da ton tai")
	}

	//kiem tra loai san pham ton tai
	if result := tx.Debug().
		Table("loai_san_pham").
		Where("id = ?", req.Id).
		First(&loai_san_pham);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("loai san pham khong ton tai")
	}

	loai_san_pham.Ten = req.Ten
	loai_san_pham.Hinh_anh = req.Hinh_anh

	//update loai san pham
	if err := tx.Model(&loai_san_pham).Debug().Updates(&loai_san_pham).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat loai san pham: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteProductTypeExec(req *requests.Loai_san_pham_delete) error {
	var loai_san_pham db.Loai_san_pham

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra loai san pham ton tai
	if result := helpers.GormDB.Debug().
		Table("loai_san_pham").
		Where("id = ?", req.Id).
		First(&loai_san_pham);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("loai san pham khong ton tai")
	}

	//kiem tra loai san pham ton tai trong san pham
	var count int64 = 0

	if err := tx.Table("san_pham").Where("loai_san_pham_id = ?", req.Id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the kiem tra loai san pham trong san pham: " + err.Error())
	}

	if count != 0 {
		tx.Rollback()
		return errors.New("khong the xoa vi loai san pham ton tai trong san pham")
	}

	//delete loai san pham
	if err := helpers.GormDB.Model(&loai_san_pham).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the xoa loai san pham: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}