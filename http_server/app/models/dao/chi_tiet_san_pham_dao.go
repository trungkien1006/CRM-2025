package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func GetProductDetailExec(req *requests.Chi_tiet_san_pham_get_by_product_id, res *responses.Chi_tiet_san_pham_get_by_product_id) error {
	var isProductExist int64

	helpers.GormDB.Debug().
		Table("san_pham").
		Where("id = ?", req.Product_id).
		Count(&isProductExist)

	if isProductExist == 0 {
		return errors.New("san pham khong ton tai")
	}

	if err := helpers.GormDB.Debug().Table("chi_tiet_san_pham").
		Where("san_pham_id = ?", req.Product_id).
		Where("khong_phan_loai = 0").
		Find(&res.Chi_tiet_san_pham).Error;
	err != nil {
		return errors.New("chi tiet san pham khong ton tai")
	}

	return nil
}