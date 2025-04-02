package controllers

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/dao"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Role
// @Security BearerAuth
// @Description Filter role based on provided filters
// @Tags role
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/chuc-vu [get]
func FilterRole(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[db.Chuc_vu]

	if err := Filter(&req, &res, c, "chuc_vu"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay chuc vu thanh cong",
	})
}

// @Summary Create Role
// @Security BearerAuth
// @Description Create a new role entry
// @Tags role
// @Accept  json
// @Produce json
// @Param Discount_Type body requests.Chuc_vu_create true "Role data"
// @Router /api/v1/chuc-vu [post]
func CreateRole(c *gin.Context) {
	var req requests.Chuc_vu_create
	var res responses.Chuc_vu_create

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "binding",
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateRoleExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them chuc vu thanh cong",
	})
}

// @Summary Update Role
// @Security BearerAuth
// @Description Update an existing role entry
// @Tags role
// @Accept  json
// @Produce json
// @Param Chuc_vu body requests.Chuc_vu_update true "Updated Role data"
// @Router /api/v1/chuc-vu [put]
func UpdateRole(c *gin.Context) {
	var req requests.Chuc_vu_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateRoleExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat chuc vu thanh cong",
	})
}

// @Summary Delete Role
// @Security BearerAuth
// @Description Delete an existing role entry
// @Tags role
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "Role ID to be deleted"
// @Router /api/v1/chuc-vu/{id} [delete]
func DeleteRole(c *gin.Context) {
	var req requests.Chuc_vu_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteRoleExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := helpers.Redis.Del(helpers.Ctx, "role:" + string(req.Id)).Err(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa chuc vu thanh cong",
	})
}