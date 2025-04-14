package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Message
// @Security BearerAuth
// @Description Filter message based on provided filters
// @Tags message
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/tin-nhan [get]
func FilterMessage(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.Tin_nhan_filter]

	if err := Filter(&req, &res, c, "tin_nhan"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay tin nhan thanh cong",
	})
}

// @Summary Create Message
// @Security BearerAuth
// @Description Create a new message entry
// @Tags message
// @Accept  json
// @Produce json
// @Param CreateMessage body requests.Tin_nhan_create true "Message Create Data"
// @Success 200 {object} map[string]interface{} "data: Tin_nhan_create, message: them tin nhan thanh cong"
// @Failure 400 {object} map[string]string "message: error message"
// @Router /api/v1/tin-nhan [post]
func CreateMessage(c *gin.Context) {
	var req requests.Tin_nhan_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateMessageExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "them tin nhan thanh cong",
	})
}

func CreateBatchMessage(c *gin.Context) {
	var req requests.Tin_nhan_create_batch

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateBatchMessageExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "them batch tin nhan thanh cong",
	})
}