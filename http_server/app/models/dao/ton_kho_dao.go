package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"errors"
)

func GetInStockByProductDetailIdExec(req *requests.Ton_kho_get_by_ctsp_id, res *responses.Ton_kho_response) error {
	var count int64 = 0

	if err := helpers.GormDB.Debug().Table("chi_tiet_san_pham").
		Where("id = ?", req.Ctsp_id).
		Count(&count).Error; 
	err != nil {
		return errors.New("co loi khi kiem tra ctsp: " + err.Error())
	}

	if count == 0 {
		return errors.New("chi tiet san pham khong ton tai")
	}

	if err := helpers.GormDB.Debug().Table("ton_kho as tk").
		Where("tk.ctsp_id = ?", req.Ctsp_id).
		Joins("JOIN chi_tiet_hoa_don_nhap_kho as ct_hdnk ON ct_hdnk.sku = tk.sku").
		Select("tk.id, tk.san_pham_id, tk.ctsp_id, tk.sku, tk.so_luong_ton, ct_hdnk.han_su_dung, ct_hdnk.don_vi_tinh, tk.created_at, ct_hdnk.gia_ban").
		Find(&res.Ds_ton_kho).Error;
	err != nil {
		return errors.New("co loi khi truy xuat ton kho: " + err.Error())
	}

	return nil
}