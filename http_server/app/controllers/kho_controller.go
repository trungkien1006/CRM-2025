package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Ware House
// @Security BearerAuth
// @Description Filter ware house based on provided filters
// @Tags ware house
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/kho [get]
func FilterWareHouse(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[db.Kho]

	if err := Filter(&req, &res, c, "kho"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay kho thanh cong",
	})
}

// @Summary Create Ware House
// @Security BearerAuth
// @Description Create a new ware house entry
// @Tags ware house
// @Accept application/json
// @Produce json
// @Param Ware_House body requests.Kho_create true "ware house data"
// @Router /api/v1/kho [post]
func CreateWareHouse(c *gin.Context) {
	var req requests.Kho_create
	var res responses.Kho_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateWareHouseExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them kho thanh cong",
	})
}

// @Summary Update WareHouse
// @Security BearerAuth
// @Description Update an existing ware house entry
// @Tags ware house
// @Accept application/json
// @Produce json
// @Param Ware_House body requests.Kho_update true "Updated ware house data"
// @Router /api/v1/kho [put]
func UpdateWareHouse(c *gin.Context) {
	var req requests.Kho_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateWareHouseExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat kho thanh cong",
	})
}

// @Summary Delete Ware House
// @Security BearerAuth
// @Description Delete an existing ware house entry
// @Tags ware house
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "Ware House ID to be deleted"
// @Router /api/v1/kho/{id} [delete]
func DeleteWareHouse(c *gin.Context) {
	var req requests.Kho_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteWareHouseExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa kho thanh cong",
	})
}