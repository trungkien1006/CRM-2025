package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func LoginExec(req *requests.Dang_nhap, res *responses.Dang_nhap) (uint, int, error) {
	var nhan_vien db.Nhan_vien

	if result := helpers.GormDB.Debug().Table("nhan_vien").
		Where("ten_dang_nhap = ?", req.Ten_dang_nhap).
		First(&nhan_vien).RowsAffected; 
	result == 0 {
		return 0, 0, errors.New("ten dang nhap khong ton tai")
	}

	if !helpers.CheckPasswordHash(req.Mat_khau, nhan_vien.Mat_khau) {
		return 0, 0, errors.New("sai mat khau")
	}

	var ds_quyen []string

	if err := helpers.GormDB.Debug().Table("quyen").
		Where("chuc_vu_id = ?", nhan_vien.Chuc_vu_id).
		Joins("JOIN chuc_nang ON chuc_nang.id = quyen.chuc_nang_id").
		Select("chuc_nang.code").
		Find(&ds_quyen).Error; 
	err != nil {
		return 0, 0, errors.New("khong the tim thay quyen")
	}

	var user = helpers.UserJWTSubject{
		Id: nhan_vien.ID,
		Name: nhan_vien.Ten_dang_nhap,
	}

	res.Token = helpers.GenerateToken(user)
	res.Ds_quyen = ds_quyen

	return nhan_vien.ID, nhan_vien.Chuc_vu_id, nil
}

func GetMeExec(nhan_vien_id int, res *responses.Get_me) error {

	if err := helpers.GormDB.Debug().
		Table("nhan_vien").
		Where("nhan_vien.id = ?", nhan_vien_id).
		Joins("JOIN chuc_vu ON chuc_vu.id = nhan_vien.chuc_vu_id").
		Select("nhan_vien.*, chuc_vu.ten as chuc_vu").
		Find(&res).Error;
	err != nil {
		return errors.New("loi khi truy van du lieu nhan vien: " + err.Error())
	}

	if err := helpers.GormDB.Debug().
		Table("chuc_nang").
		Joins("JOIN quyen ON chuc_nang.id = quyen.chuc_nang_id").
		Where("quyen.chuc_vu_id = ?", res.Chuc_vu_id).
		Select("chuc_nang.code").
		Find(&res.Quyen).Error; 
	err != nil {
		return errors.New("loi khi truy van du lieu quyen cua nhan vien: " + err.Error())
	}

	return nil
}