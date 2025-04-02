package helpers

import (
	"admin-v1/app/models/requests"

	"gorm.io/gorm"
)



func Filter(query *gorm.DB, filters []requests.FilterStruc) {
	for _, value := range filters {
		if value.Field == "nhan_vien_sale" || value.Field == "nhan_vien_giao_hang" || value.Field == "khach_hang" || value.Field == "loai_san_pham" || value.Field == "don_vi_tinh" || value.Field == "loai_giam_gia" || value.Field == "thoi_gian_bao_hanh" || value.Field == "chuc_vu" || value.Field == "kho" {
			value.Field += ".ten"
		} else if value.Field == "nha_phan_phoi" {
			value.Field = "sp_npp.nha_phan_phoi_id"
		}

		switch value.Condition {
			case "contains":{
				query.Where(value.Field + " LIKE ?", "%" + value.Value + "%")
				break
			}
			case "notcontains":{
				query.Where(value.Field + " NOT LIKE ?", "%" + value.Value + "%")
				break
			}
			case "startswith":{
				query.Where(value.Field + " LIKE ?", value.Value + "%")
				break
			}
			case "endswith":{
				query.Where(value.Field + " LIKE ?", "%" + value.Value)
				break
			}
			case "=":{
				query.Where(value.Field + " = ?", value.Value)
				break
			}
			case "<>":{
				query.Where(value.Field + " != ?", value.Value)
				break
			}
			case ">":{
				query.Where(value.Field + " > ?", value.Value)
				break
			}
			case "<":{
				query.Where(value.Field + " < ?", value.Value)
				break
			}
			case ">=":{
				query.Where(value.Field + " >= ?", value.Value)
				break
			}
			case "<=":{
				query.Where(value.Field + " <= ?", value.Value)
				break
			}
		}
	}
}