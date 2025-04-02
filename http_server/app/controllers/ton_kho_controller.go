package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get In Stock
// @Security BearerAuth
// @Description Get in stock by product detail id
// @Tags in stock
// @Accept application/x-www-form-urlencoded
// @Param Ctsp_id path int true "Product detail id"
// @Router /api/v1/ton-kho/{ctsp_id} [get]
func GetInStockByProductDetailId(c *gin.Context) {
	var req requests.Ton_kho_get_by_ctsp_id
	var res responses.Ton_kho_response

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.GetInStockByProductDetailIdExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"data" : res.Ds_ton_kho,
		},
		"message": "truy xuat danh sach ton kho thanh cong",
	})
}