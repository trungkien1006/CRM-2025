package responses

import "admin-v1/app/models/db"

type Chi_tiet_san_pham_get_by_product_id struct {
	Chi_tiet_san_pham []db.Chi_tiet_san_pham
}