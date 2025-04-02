package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Unit
// @Security BearerAuth
// @Description Filter unit based on provided filters
// @Tags unit
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/don-vi-tinh [get]
func FilterUnit(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[db.Don_vi_tinh]

	if err := Filter(&req, &res, c, "don_vi_tinh"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay don vi tinh thanh cong",
	})
}

// @Summary Create unit
// @Security BearerAuth
// @Description Create a new unit entry
// @Tags unit
// @Accept application/json
// @Produce json
// @Param Unit body requests.Don_vi_tinh_create true "Unit data"
// @Router /api/v1/don-vi-tinh [post]
func CreateUnit(c *gin.Context) {
	var req requests.Don_vi_tinh_create
	var res responses.Don_vi_tinh_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateUnitExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them don vi tinh thanh cong",
	})
}

// @Summary Update Unit
// @Security BearerAuth
// @Description Update an existing unit entry
// @Tags unit
// @Accept application/json
// @Produce json
// @Param Unit body requests.Don_vi_tinh_update true "Updated unit data"
// @Router /api/v1/don-vi-tinh [put]
func UpdateUnit(c *gin.Context) {
	var req requests.Don_vi_tinh_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateUnitExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat don vi tinh thanh cong",
	})
}

// @Summary Delete Unit
// @Security BearerAuth
// @Description Delete an existing unit entry
// @Tags unit
// @Accept application/x-www-form-urlencoded
// @Param id path int true "unit ID to be deleted"
// @Router /api/v1/don-vi-tinh/{id} [delete]
func DeleteUnit(c *gin.Context) {
	var req requests.Don_vi_tinh_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteUnitExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa don vi tinh thanh cong",
	})
}