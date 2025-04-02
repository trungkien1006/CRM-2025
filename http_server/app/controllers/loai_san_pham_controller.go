package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Product Type
// @Security BearerAuth
// @Description Filter product type based on provided filters
// @Tags product type
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/loai-san-pham [get]
func FilterProductType(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[db.Loai_san_pham]

	if err := Filter(&req, &res, c, "loai_san_pham"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error": "binding",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay loai san pham thanh cong",
	})
}

// @Summary Create Product Type
// @Security BearerAuth
// @Description Create a new product type entry
// @Tags product type
// @Accept  json
// @Produce json
// @Param CreateProductType body requests.Loai_san_pham_create true "Product Type Create Data"
// @Success 200 {object} map[string]interface{} "data: Loai_san_pham_create, message: them loai san pham thanh cong"
// @Failure 400 {object} map[string]string "message: error message"
// @Router /api/v1/loai-san-pham [post]
func CreateProductType(c *gin.Context) {
	var req requests.Loai_san_pham_create
	var res responses.Loai_san_pham_create

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "binding",
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateProductTypeExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them loai san pham thanh cong",
	})
}

// @Summary Update Product Type
// @Security BearerAuth
// @Description Update an existing product type entry
// @Tags product type
// @Accept  json
// @Produce json
// @Param UpdateProductType body requests.Loai_san_pham_update true "Product Type Update Data"
// @Success 200 {object} map[string]string "message: cap nhat loai san pham thanh cong"
// @Failure 400 {object} map[string]string "message: error message"
// @Router /api/v1/loai-san-pham [put]
func UpdateProductType(c *gin.Context) {
	var req requests.Loai_san_pham_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateProductTypeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat loai san pham thanh cong",
	})
}

// @Summary Delete Product Type
// @Security BearerAuth
// @Description Delete an existing product type entry
// @Tags product type
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "product type ID to be deleted"
// @Router /api/v1/loai-san-pham/{id} [delete]
func DeleteProductType(c *gin.Context) {
	var req requests.Loai_san_pham_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteProductTypeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa loai san pham thanh cong",
	})
}