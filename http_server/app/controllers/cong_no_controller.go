package controllers

import (
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Customer Debt
// @Security BearerAuth
// @Description Filter customer debt based on provided filters
// @Tags debt
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/cong-no-khach-hang [get]
func FilterCustomerDebt(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.Cong_no_khach_hang_filter]

	if err := Filter(&req, &res, c, "cong_no_khach_hang"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay cong no khach hang thanh cong",
	})
}

// @Summary Filter Provider Debt
// @Security BearerAuth
// @Description Filter provider debt based on provided filters
// @Tags debt
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/cong-no-nha-phan-phoi [get]
func FilterProviderDebt(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.Cong_no_nha_phan_phoi_filter]

	if err := Filter(&req, &res, c, "cong_no_nha_phan_phoi"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay cong no nha phan phoi thanh cong",
	})
}