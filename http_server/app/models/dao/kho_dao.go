package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateWareHouseExec(req *requests.Kho_create, res *responses.Kho_create) error {
	//kiem tra ten kho ton tai
	if result := helpers.GormDB.Debug().
		Table("kho").
		Where("ten = ?", req.Ten).
		First(&res.Kho);
	result.RowsAffected > 0 {
		return errors.New("kho da ton tai")
	}
	
	var kho = db.Kho{
		Ten: req.Ten,
		Dia_chi: req.Dia_chi,
	}

	//insert kho
	if err := helpers.GormDB.Debug().Create(&kho).Error; err != nil {
		return errors.New("khong the tao kho: " + err.Error())
	}

	res.Kho = kho

	return nil
}

func UpdateWareHouseExec(req *requests.Kho_update) error {
	var kho db.Kho
	var kho_temp db.Kho

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra ten kho ton tai
	if result := tx.Debug().
		Table("kho").
		Where("ten = ?", req.Ten).
		Where("ID != ?", req.Id).
		First(&kho_temp);
	result.RowsAffected > 0 {
		return errors.New("ten kho da ton tai")
	}

	//kiem tra kho ton tai
	if result := tx.Debug().
		Table("kho").
		Where("id = ?", req.Id).
		First(&kho);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("kho khong ton tai")
	}

	kho.Ten = req.Ten
	kho.Dia_chi = req.Dia_chi

	//update kho
	if err := tx.Model(&kho).Debug().Updates(&kho).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat kho: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteWareHouseExec(req *requests.Kho_delete) error {
	var kho db.Kho

	//bat dau transaction
	tx := helpers.GormDB.Begin()
	
	//kiem tra kho ton tai
	if result := tx.Debug().
		Table("kho").
		Where("id = ?", req.Id).
		First(&kho);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("kho khong ton tai")
	}

	//kiem tra kho ton tai trong hdnk
	var count int64 = 0

	if err := tx.Table("hoa_don_nhap_kho").Where("kho_id").Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("co loi khi kim tra kho trong hoa don nhap kho: " + err.Error())
	}

	if count != 0 {
		tx.Rollback()
		return errors.New("khong the xoa kho vi co ton tai trong hoa don nhap kho")
	}

	//delete kho
	if err := tx.Model(&kho).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the xoa kho: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}