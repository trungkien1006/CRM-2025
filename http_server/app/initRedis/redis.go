package initRedis

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/dao"
	"admin-v1/app/models/responses"
	"encoding/json"
	"fmt"
	"strconv"
)

func InitRolePermission() {
	var res responses.Quyen_by_chuc_vu_id

	//lay ra danh sach quyen theo id chuc vu
	if err := dao.GetFullPermissionByRoleId(&res); err != nil {
		fmt.Println("khong the lay duoc danh sach quyen: " + err.Error())

		return
	}

	//lap qua danh sach quyen va them vao redis
	for _, value := range res.Quyen_list {
		quyenJson, err := json.Marshal(value.Quyen)

		if err != nil {
			fmt.Println("co loi khi ma hoa quyen: " + err.Error())

			return
		}

		if e := helpers.Redis.Set(helpers.Ctx, "role:" + strconv.Itoa(int(value.Id)), quyenJson, 0).Err(); e != nil {
			fmt.Println("co loi khi them redis: " + e.Error())

			return
		}
	}

	fmt.Println("khoi tao du lieu quyen voi redis thanh cong")
}