package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Product
// @Security BearerAuth
// @Description Filter product based on provided filters
// @Tags product
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/san-pham [get]
func FilterProduct(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.San_pham_filter]

	if err := Filter(&req, &res, c, "san_pham"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay san pham thanh cong",
	})
}

// @Summary Create Product
// @Security BearerAuth
// @Description Create a new product entry
// @Tags product
// @Accept  json
// @Produce json
// @Param CreateProduct body requests.San_pham_create true "Product Create Data"
// @Success 200 {object} map[string]interface{} "data: San_pham_create, message: them san pham thanh cong"
// @Failure 400 {object} map[string]string "message: error message"
// @Router /api/v1/san-pham [post]
func CreateProduct(c *gin.Context) {
	var req requests.San_pham_create
	var res responses.San_pham_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateProductExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them san pham thanh cong",
	})
}

// @Summary Update Product
// @Security BearerAuth
// @Description Update an existing product entry
// @Tags product
// @Accept  json
// @Produce json
// @Param UpdateProduct body requests.San_pham_update true "Product Update Data"
// @Success 200 {object} map[string]interface{} "message: cap nhat san pham thanh cong"
// @Failure 400 {object} map[string]string "message: error message"
// @Router /api/v1/san-pham [put]
func UpdateProduct(c *gin.Context) {
	var req requests.San_pham_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateProductExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat san pham thanh cong",
	})
}

// @Summary Delete Product
// @Security BearerAuth
// @Description Delete an existing product entry
// @Tags product
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "product ID to be deleted"
// @Router /api/v1/san-pham/{id} [delete]
func DeleteProduct(c *gin.Context) {
	var req requests.San_pham_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteProductExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa san pham thanh cong",
	})
}