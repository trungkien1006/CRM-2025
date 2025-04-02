package controllers

import (
	"admin-v1/app/models/dao"
	// "admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Provider
// @Security BearerAuth
// @Description Filter provider based on provided filters
// @Tags provider
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/nha-phan-phoi [get]
func FilterProvider(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.Nha_phan_phoi_filter]

	if err := Filter(&req, &res, c, "nha_phan_phoi"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay nha cung cap thanh cong",
	})
}

// @Summary Create Provider
// @Security BearerAuth
// @Description Create a new provider entry
// @Tags provider
// @Accept application/json
// @Produce json
// @Param Provider body requests.Nha_phan_phoi_create true "Provider data"
// @Router /api/v1/nha-phan-phoi [post]
func CreateProvider(c *gin.Context) {
	var req requests.Nha_phan_phoi_create
	var res responses.Nha_phan_phoi_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateProviderExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them nha cung cap thanh cong",
	})
}

// @Summary Update Provider
// @Security BearerAuth
// @Description Update an existing provider entry
// @Tags provider
// @Accept application/json
// @Produce json
// @Param Provider body requests.Nha_phan_phoi_update true "Updated provider data"
// @Router /api/v1/nha-phan-phoi [put]
func UpdateProvider(c *gin.Context) {
	var req requests.Nha_phan_phoi_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateProviderExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat nha cung cap thanh cong",
	})
}

// @Summary Delete Provider
// @Security BearerAuth
// @Description Delete an existing provider entry
// @Tags provider
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "Provider ID to be deleted"
// @Router /api/v1/nha-phan-phoi/{id} [delete]
func DeleteProvider(c *gin.Context) {
	var req requests.Nha_phan_phoi_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteProviderExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa nha cung cap thanh cong",
	})
}