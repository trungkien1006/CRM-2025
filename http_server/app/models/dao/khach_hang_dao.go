package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateCustomerExec(req *requests.Khach_hang_create, res *responses.Khach_hang_create) error {
	var khach_hang = db.Khach_hang{
		Ho_ten: req.Ho_ten,
		Dien_thoai: req.Dien_thoai,
		Dia_chi: req.Dia_chi,
	}

	//insert khach hang
	if err := helpers.GormDB.Debug().Create(&khach_hang).Error; err != nil {
		return errors.New("khong the tao khach hang: " + err.Error())
	}

	res.Khach_hang = khach_hang

	return nil
}

func UpdateCustomerExec(req *requests.Khach_hang_update) error {
	var khach_hang db.Khach_hang

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra khach hang ton tai
	if result := tx.Debug().
		Table("khach_hang").
		Where("id = ?", req.Id).
		First(&khach_hang);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("khach hang khong ton tai")
	}

	khach_hang.Ho_ten = req.Ho_ten
	khach_hang.Dien_thoai = req.Dien_thoai
	khach_hang.Dia_chi = req.Dia_chi

	//update khach hang
	if err := tx.Model(&khach_hang).Debug().Updates(&khach_hang).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat khach hang: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteCustomerExec(req *requests.Khach_hang_delete) error {
	var khach_hang db.Khach_hang

	//kiem tra khach hang ton tai
	if result := helpers.GormDB.Debug().
		Table("khach_hang").
		Where("id = ?", req.Id).
		First(&khach_hang);
	result.RowsAffected == 0 {
		return errors.New("khach hang khong ton tai")
	}

	//delete khach hang
	if err := helpers.GormDB.Model(&khach_hang).Debug().
		Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; 
	err != nil {
		return errors.New("khong the xoa khach hang: " + err.Error())
	}

	return nil
}