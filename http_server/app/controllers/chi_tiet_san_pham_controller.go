package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProductDetail lấy chi tiết sản phẩm theo Product ID
// @Summary Get Product Detail
// @Security BearerAuth
// @Description API này lấy thông tin chi tiết của một sản phẩm theo ID
// @Tags product detail
// @Accept  json
// @Produce  json
// @Param  product_id path int true "ID của sản phẩm"
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/chi-tiet-san-pham/{product_id} [get]
func GetProductDetail(c *gin.Context) {
	var req requests.Chi_tiet_san_pham_get_by_product_id
	var res responses.Chi_tiet_san_pham_get_by_product_id

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.GetProductDetailExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    gin.H{
			"data": res.Chi_tiet_san_pham,
		},
		"message": "lay chi tiet san pham thanh cong",
	})
}