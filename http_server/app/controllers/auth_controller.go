package controllers

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Login
// @Description Login API
// @Tags auth
// @Accept  json
// @Produce  json
// @Param Login_data body requests.Dang_nhap true "Login data include username, password"
// @Success 200 {object} responses.Dang_nhap
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/dang-nhap [post]
func Login(c *gin.Context) {
	var req requests.Dang_nhap
	var res responses.Dang_nhap

	//validate params
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "binding",
			"message": err.Error(),
		})

		return
	}

	//truy van thong tin dang nhap
	if id, role_id, err := dao.LoginExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	} else {
		//neu thanh cong se luu id chuc vu vao redis với key la "user:<nhanvien_id>"
		//kiem tra key da ton tai hay chua, neu chua moi them vao redis
		isExist, redisErr := helpers.Redis.Exists(helpers.Ctx, "user:" + strconv.Itoa(int(id))).Result()

		if redisErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "redis check exist",
				"message": redisErr.Error(),
			})

			return
		}

		if isExist == 0 {
			if e := helpers.Redis.Set(helpers.Ctx, "user:" + strconv.Itoa(int(id)), strconv.Itoa(role_id), 0).Err(); e != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "redis save",
					"message": e.Error(),
				})

				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    gin.H{
			"data": res,
		},
		"message": "dang nhap thanh cong",
	})
}

// @Summary Get Me
// @Description Get Me API by sending JWT
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.Get_me
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/thong-tin-nhan-vien [get]
func GetMe(c *gin.Context) {
	//lay ra jwt tu header
	var jwt string = c.GetHeader("Authorization")

	//parse jwt thanh object
	var userSub helpers.UserJWTSubject = helpers.GetTokenSubject(jwt)

	var res responses.Get_me

	if err := dao.GetMeExec(int(userSub.Id), &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"data": res,
		},
		"message": "get me successfull",
	})
}