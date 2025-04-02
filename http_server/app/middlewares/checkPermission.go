package middlewares

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/dao"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// "golang.org/x/text/cases"
)

func CheckPermission(c *gin.Context) {
	//lay ra jwt tu header
	var jwt string = c.GetHeader("Authorization")

	//parse jwt thanh object
	var userSub helpers.UserJWTSubject = helpers.GetTokenSubject(jwt)

	//lay role_id cua user
	role_id, err := helpers.Redis.Get(helpers.Ctx, "user:" + strconv.Itoa(int(userSub.Id))).Result()

	//khoi tao bien chua danh sach quyen
	var permissions []string

	//neu viec get tu redis co loi
	if err != nil {
		//neu loi do la do khong ton tai key
		if err.Error() == "redis: nil" {
			//truy van db de lay role_id va danh sach quyen
			if role_id_DB, permissionsDB, perErr := dao.CheckPermissionByUserId(userSub.Id); perErr != nil {
				//log loi truy van db
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "gap loi khi truy van csdl de check quyen: " + perErr.Error(),
				})

				c.Abort()
				return
			} else {
				//kiem tra role_id va danh sach quyen co ton tai
				if role_id_DB == 0 || len(permissionsDB) == 0 {
					//neu khong se luu user_id la -1 vao redis de tranh cache penetration
					if e := helpers.Redis.Set(helpers.Ctx, "user:" + strconv.Itoa(int(userSub.Id)), -1, time.Hour).Err(); e != nil {
						c.JSON(http.StatusForbidden, gin.H{
							"error": "gap loi khi set null ngan chan cache penetration: " + e.Error(),
						})

						c.Abort()
						return
					}
				} else {
					//neu co du lieu se bat dau set redis user:<int> va role:<[]string>
					if e := helpers.Redis.Set(helpers.Ctx, "user:" + strconv.Itoa(int(userSub.Id)), strconv.Itoa(int(role_id_DB)), 0).Err(); e != nil {
						c.JSON(http.StatusForbidden, gin.H{
							"error": "gap loi khi set user role id redis: " + e.Error(),
						})

						c.Abort()
						return
					}

					permissionDBJSON, err := json.Marshal(permissionsDB)

					if err != nil {
						c.JSON(http.StatusForbidden, gin.H{
							"error": "gap loi khi ma hoa danh sach quyen thanh json",
						})

						c.Abort()
						return
					}

					if e := helpers.Redis.Set(helpers.Ctx, "role:" + strconv.Itoa(int(role_id_DB)), permissionDBJSON, 0).Err(); e != nil {
						c.JSON(http.StatusForbidden, gin.H{
							"error": "gap loi khi set user role id redis: " + e.Error(),
						})

						c.Abort()
						return
					}

					//luu ds quyen vao ds chung
					permissions = permissionsDB
				}
			}
		} else {
			//log loi cua redis get
			c.JSON(http.StatusForbidden, gin.H{
				"error": "loi khi kiem tra quyen 1: " + err.Error(),
			})

			c.Abort()
			return
		}
	} else {
		//kiem tra xem user_id do co bi lock hay khong
		if role_id == "-1" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "ban da bi lock",
			})

			c.Abort()
			return
		} else {
			//lay ds quyen cua user từ role_id duoi dang json
			permissionJson, err := helpers.Redis.Get(helpers.Ctx, "role:" + role_id).Result()

			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"error": "loi khi kiem tra quyen 2: " + err.Error(),
				})

				c.Abort()
				return
			}

			//giai ma json
			json.Unmarshal([]byte(permissionJson), &permissions)
		}
	}

	//lay ra method va path cua request
	method := c.Request.Method 
    path := strings.Split(c.Request.URL.Path, "/")

	var action 		string = ""

	//tao code cua quyen
	switch(method) {
		case "GET": {
			action = "view-" +  path[3]
			break
		}
		case "POST": {
			action = "create-" +  path[3]
			break
		}
		case "PUT": {
			action = "update-" +  path[3]
			break
		}
		case "DELETE": {
			action = "delete-" +  path[3]
			break
		}
		case "PATCH": {
			action = path[4] + "-" + path[3]
		}
	}

	var permissionExist bool = false

	//kiem tra quyen da ton tai hay chua
	for _, value := range permissions {
		if action == value {
			permissionExist = true 
			break
		}
	}

	if !permissionExist {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "ban khong co quyen thuc hien chuc nang",
		})

		c.Abort()
		return
	} 
	
	c.Next()
}