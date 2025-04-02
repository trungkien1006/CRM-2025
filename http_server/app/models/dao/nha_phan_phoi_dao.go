package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"

	"gorm.io/gorm"
)

func CreateProviderExec(req *requests.Nha_phan_phoi_create, res *responses.Nha_phan_phoi_create) error {
	//kiem tra nha phan phoi ton tai
	if result := helpers.GormDB.Debug().
		Table("nha_phan_phoi").
		Where("ten = ?", req.Ten).
		First(&res.Nha_phan_phoi);
	result.RowsAffected > 0 {
		return errors.New("nha phan phoi da ton tai")
	}
	
	//select danh sach san pham 
	var ds_san_pham []db.San_pham

	if result := helpers.GormDB.Debug().
		Table("san_pham").
		Where("id IN ?", req.Ds_san_pham).
		Find(&ds_san_pham).Error; 
	result != nil {
		return errors.New("tim san pham bi loi")
	}

	//tao doi tuong nha phan phoi
	var nha_phan_phoi = db.Nha_phan_phoi{
		Ten: req.Ten,
		Dia_chi: req.Dia_chi,
		Dien_thoai: req.Dien_thoai,
		Email: req.Email,
		
		San_pham: ds_san_pham,
	}

	//insert nha phan phoi
	if err := helpers.GormDB.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Save(&nha_phan_phoi).Error; err != nil {
		return errors.New("khong the tao nha phan phoi: " + err.Error())
	}

	res.Nha_phan_phoi = nha_phan_phoi

	return nil
}

func UpdateProviderExec(req *requests.Nha_phan_phoi_update) error {
	var nha_phan_phoi db.Nha_phan_phoi
	var nha_phan_phoi_temp db.Nha_phan_phoi

	//bat dau transaction
	tx := helpers.GormDB.Begin()
	
	//kiem tra ten nha phan phoi ton tai
	if result := tx.Debug().
		Table("nha_phan_phoi").
		Where("ten = ?", req.Ten).
		Where("ID != ?", req.Id).
		First(&nha_phan_phoi_temp);
	result.RowsAffected > 0 {
		return errors.New("ten nha phan phoi da ton tai")
	}

	//select nha phan phoi ton tai
	if result := tx.Debug().
		Table("nha_phan_phoi").
		Where("id = ?", req.Id).
		First(&nha_phan_phoi);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("nha phan phoi khong ton tai")
	}

	var ids_sp		[]int

	for _, value := range req.Ds_san_pham {
		var San_pham_npp db.San_pham_nha_phan_phoi

		if res := tx.Table("san_pham_nha_phan_phoi").
			Where("nha_phan_phoi_id = ?", req.Id).
			Where("san_pham_id = ?", value).
			Find(&San_pham_npp).RowsAffected;
		res == 0 {
			if err := tx.Debug().Model(&db.San_pham_nha_phan_phoi{}).
				Create(&San_pham_npp).Error; 
			err != nil {
				tx.Rollback()
				return errors.New("khong the cap nhat san pham cua nha phan phoi: " + err.Error())
			}
		}
		
		ids_sp = append(ids_sp, value)
	}

	//xoa cac san pham nha phan phoi
	if err := tx.Debug().Table("san_pham_nha_phan_phoi").
		Where("nha_phan_phoi_id = ?", req.Id).
		Where("san_pham_id NOT IN ?", ids_sp).
		Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; 
	err != nil {
		tx.Rollback()
		return errors.New("khong the xoa san pham cua nha phan phoi: " + err.Error())
	}

	nha_phan_phoi.Ten = req.Ten
	nha_phan_phoi.Dia_chi = req.Dia_chi
	nha_phan_phoi.Email = req.Email
	nha_phan_phoi.Dien_thoai = req.Dien_thoai

	//update nha phan phoi
	if err := tx.Debug().Model(&nha_phan_phoi).Updates(&nha_phan_phoi).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the cap nhat nha phan phoi: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func DeleteProviderExec(req *requests.Nha_phan_phoi_delete) error {
	var nha_phan_phoi db.Nha_phan_phoi

	tx := helpers.GormDB.Begin()

	//kiem tra nha phan phoi ton tai
	if result := tx.Debug().
		Table("nha_phan_phoi").
		Where("id = ?", req.Id).
		First(&nha_phan_phoi);
	result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("nha phan phoi khong ton tai")
	}

	//kiem tra nha phan phoi ton tai trong hdnk
	var count int64 = 0

	if err := tx.Table("hoa_don_nhap_kho").Where("nha_phan_phoi_id = ?", req.Id).Count(&count).Error; err != nil {
		tx.Rollback()
		return errors.New("loi khi kiem tra nha phan phoi trong hdnk: " + err.Error())
	}

	if count != 0 {
		tx.Rollback()
		return errors.New("khong the xoa nha phan phoi vi co anh huong den hdnk")
	}

	//delete nha phan phoi
	if err := tx.Model(&nha_phan_phoi).Debug().Update("deleted_at", helpers.GetCurrentTimeVN().String()).Error; err != nil {
		tx.Rollback()
		return errors.New("khong the xoa nha phan phoi: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}