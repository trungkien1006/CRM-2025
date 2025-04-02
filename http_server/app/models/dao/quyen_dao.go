package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
	"strings"

	"gorm.io/gorm"
)

func GetPermissionExec(req *requests.Quyen_read ,res *responses.Quyen_read) error {
	var chuc_nang []db.Chuc_nang

	//lay ra tat ca chuc nang theo id chuc vu
	if err := helpers.GormDB.
		Table("chuc_nang").
		Select("chuc_nang.*, CASE WHEN quyen.chuc_nang_id IS NOT NULL THEN true ELSE false END as thuoc_chuc_vu").
		Joins("LEFT JOIN quyen ON chuc_nang.id = quyen.chuc_nang_id AND quyen.chuc_vu_id = ?", req.Chuc_vu_id).
		Find(&chuc_nang).Error; 
	err != nil {
		return errors.New("khong co chuc nang nao")
	}

	group := make(map[string][]responses.Quyen_sub)

	//group cac chuc nang theo loai chuc nang
	for _, item := range chuc_nang {
		group[item.Show_in_menu + "+" + item.Loai] = append(group[item.Show_in_menu + "+" + item.Loai], responses.Quyen_sub{
			Id: item.ID,
			Ten: item.Ten,
			Code: item.Code,
			Trang_thai: item.Thuoc_chuc_vu,
		})
	}

	//map cac chuc nang vao cac loai chuc nang tuong ung
	for key, value := range group {
		res.Quyen = append(res.Quyen, responses.Quyen{
			Hien_thi_menu: strings.Split(key, "+")[0],
			Loai_quyen: strings.Split(key, "+")[1],
			Quyen: value,
		})
	}

	return nil
}

func ModifyPermissionExec(req *requests.Quyen_modify, ds_code_quyen *[]string) error {
	var caseWhenClauses []interface{}
	var ids []interface{}

	//tao 1 danh sach id, ds cac tham so update
	for _, quyen := range req.Quyen {
		if quyen.Active == 1 {
			caseWhenClauses = append(caseWhenClauses, quyen.Quyen_id, nil)
		} else {
			caseWhenClauses = append(caseWhenClauses, quyen.Quyen_id, helpers.GetCurrentTimeVN().String())
		}

		ids = append(ids, quyen.Quyen_id)
	}

	//kiem tra danh sach cac quyen ton tai
	var countPermission int64 = 0

	if err := helpers.GormDB.Debug().Model(&db.Quyen{}).Where("chuc_nang_id IN ?", ids).Count(&countPermission).Error; err != nil {
		return errors.New("loi khi kiem tra quyen: " + err.Error())
	}

	if countPermission != int64(len(ids)) {
		return errors.New("co quyen khong ton tai")
	}

	//modify quyen
	if err := helpers.GormDB.Debug().
		Model(&db.Quyen{}).
		Where("chuc_vu_id = ?", req.Chuc_vu_id).
		Where("chuc_nang_id IN ?", ids).
		Updates(map[string]interface{}{
			"deleted_at": gorm.Expr(
				`CASE
					WHEN chuc_nang_id = ? AND deleted_at IS NOT NULL THEN NULL
					WHEN chuc_nang_id = ? AND deleted_at IS NULL THEN ?
					ELSE deleted_at
				END`, 
				caseWhenClauses...
			),
		},
	).Error; err != nil {
		return errors.New("khong the thuc hien chinh sua quyen: " + err.Error())
	}
	
	if err := helpers.GormDB.Debug().
		Table("chuc_nang").
		Joins("JOIN quyen ON chuc_nang.id = quyen.chuc_nang_id").
		Where("quyen.chuc_vu_id = ?", req.Chuc_vu_id).
		Select("chuc_nang.code").
		Find(&ds_code_quyen).Error; 
	err != nil {
		return errors.New("loi khi truy van du lieu quyen: " + err.Error())
	}

	return nil
}

func GetFullPermissionByRoleId(res *responses.Quyen_by_chuc_vu_id) error {
	if err := helpers.GormDB.Debug().Table("chuc_vu").Select("chuc_vu.id").Find(&res.Quyen_list).Error; err != nil {
		return errors.New("co loi khi lay chuc vu id: " + err.Error())
	}

	for index, value := range res.Quyen_list {
		if err := helpers.GormDB.Debug().Table("quyen").
			Joins("JOIN chuc_nang ON chuc_nang.id = quyen.chuc_nang_id").
			Where("quyen.chuc_vu_id = ?", value.Id).
			Select("chuc_nang.code").
			Find(&res.Quyen_list[index].Quyen).Error; 
		err != nil {
			return errors.New("co loi khi lay danh sach quyen: " + err.Error())
		}
	}

	return nil
}

func CheckPermissionByUserId(nhan_vien_id uint) (uint, []string, error) {
	var chuc_vu_id uint

	if err := helpers.GormDB.Debug().Table("nhan_vien").
		Where("nhan_vien.id = ?", nhan_vien_id).
		Select("nhan_vien.chuc_vu_id").
		Find(&chuc_vu_id).Error; 
	err != nil {
		return 0, nil, errors.New("loi khi lay chuc vu id cua nhan vien: " + err.Error())
	}

	var ds_quyen []string

	if err := helpers.GormDB.Debug().Table("quyen").
			Joins("JOIN chuc_nang ON chuc_nang.id = quyen.chuc_nang_id").
			Where("quyen.chuc_vu_id = ?", chuc_vu_id).
			Select("chuc_nang.code").
			Find(&ds_quyen).Error; 
	err != nil {
		return 0, nil, errors.New("co loi khi lay danh sach quyen: " + err.Error())
	}

	return chuc_vu_id, ds_quyen, nil
}