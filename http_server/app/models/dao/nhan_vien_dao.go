package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func CreateEmployeeExec(req *requests.Nhan_vien_create, res *responses.Nhan_vien_create) error {
	//kiem tra ten dang nhap
	if result := helpers.GormDB.Debug().
		Table("nhan_vien").
		Where("ten_dang_nhap = ?", req.Ten_dang_nhap).
		First(&res.Nhan_vien);
	result.RowsAffected > 0 {
		return errors.New("nhan vien da ton tai")
	}
	
	//kiem tra hash mk hop le
	var password string

	if p, err := helpers.HashPassword(req.Dien_thoai); err == nil {
		password = p
	}

	//tao cau truc nhan vien 
	var nhan_vien = db.Nhan_vien{
		Ten_dang_nhap: req.Ten_dang_nhap,
		Avatar: req.Avatar,
		Mat_khau: password,
		Ho_ten: req.Ho_ten,
		Email: req.Email,        
		Dien_thoai: req.Dien_thoai,   
		Dia_chi: req.Dia_chi,       
		Chuc_vu_id: req.Chuc_vu_id,
	}

	//insert nhan vien
	if err := helpers.GormDB.Debug().Create(&nhan_vien).Error; err != nil {
		return errors.New("khong the tao nhan vien: " + err.Error())
	}

	res.Nhan_vien = nhan_vien

	return nil
}

func UpdateEmployeeExec(req *requests.Nhan_vien_update) error {
	var nhan_vien db.Nhan_vien
	var nhan_vien_temp db.Nhan_vien

	//bat dau transaction
	tx := helpers.GormDB.Begin()

	//kiem tra ten dang nhap
	if result := tx.Debug().
		Table("nhan_vien").
		Where("ten_dang_nhap = ?", req.Ten_dang_nhap).
		Where("ID != ?", req.Id).
		First(&nhan_vien_temp);
	result.RowsAffected > 0 {
		tx.Rollback()
		return errors.New("ten nhan vien da ton tai")
	}

	//kiem tra nhan vien ton tai
	if result := tx.Debug().
		Table("nhan_vien").
		Where("id = ?", req.Id).
		First(&nhan_vien);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("nhan vien khong ton tai")
	}

	//kiem tra chuc vu ton tai
	var count int64 = 0

	if err := tx.Table("chuc_vu").Where("id = ?", req.Chuc_vu_id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("loi khi kiem tra chuc vu: " + err.Error())
	}

	if count == 0 {
		tx.Rollback()
		return errors.New("chuc vu khong on tai")
	}

	nhan_vien.Ten_dang_nhap = req.Ten_dang_nhap
	nhan_vien.Ho_ten = req.Ho_ten
	nhan_vien.Email = req.Email        
	nhan_vien.Dien_thoai = req.Dien_thoai   
	nhan_vien.Dia_chi = req.Dia_chi       
	nhan_vien.Chuc_vu_id = req.Chuc_vu_id

	//update nhan vien
	if err := tx.Model(&nhan_vien).Debug().Updates(&nhan_vien).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat nhan vien: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteEmployeeExec(req *requests.Nhan_vien_delete) error {
	var nhan_vien db.Nhan_vien

	//kiem tra nhan vien ton tai
	if result := helpers.GormDB.Debug().
		Table("nhan_vien").
		Where("id = ?", req.Id).
		First(&nhan_vien);
	result.RowsAffected == 0 {
		return errors.New("nhan vien khong ton tai")
	}

	//delete nhan vien
	if err := helpers.GormDB.Model(&nhan_vien).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		return errors.New("khong the xoa nhan vien: " + err.Error())
	}

	return nil
}