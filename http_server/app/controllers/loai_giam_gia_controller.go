package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Discount Type
// @Security BearerAuth
// @Description Filter discount type based on provided filters
// @Tags discount type
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/loai-giam-gia [get]
func FilterDiscountType(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[db.Loai_giam_gia]

	if err := Filter(&req, &res, c, "loai_giam_gia"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay loai giam gia thanh cong",
	})
}

// @Summary Create Discount Type
// @Security BearerAuth
// @Description Create a new discount type entry
// @Tags discount type
// @Accept application/json
// @Produce json
// @Param Discount_Type body requests.Loai_giam_gia_create true "Discount Type data"
// @Router /api/v1/loai-giam-gia [post]
func CreateDiscountType(c *gin.Context) {
	var req requests.Loai_giam_gia_create
	var res responses.Loai_giam_gia_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateDiscountTypeExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them loai giam gia thanh cong",
	})
}

// @Summary Update Discount Type
// @Security BearerAuth
// @Description Update an existing discount type entry
// @Tags discount type
// @Accept application/json
// @Produce json
// @Param Discount_Type body requests.Loai_giam_gia_update true "Updated Discount Type data"
// @Router /api/v1/loai-giam-gia [put]
func UpdateDiscountType(c *gin.Context) {
	var req requests.Loai_giam_gia_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateDiscountTypeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat loai giam gia thanh cong",
	})
}

// @Summary Delete Discount Type
// @Security BearerAuth
// @Description Delete an existing discount type entry
// @Tags discount type
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "discount type ID to be deleted"
// @Router /api/v1/loai-giam-gia/{id} [delete]
func DeleteDiscountType(c *gin.Context) {
	var req requests.Loai_giam_gia_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteDiscountTypeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa loai giam gia thanh cong",
	})
}