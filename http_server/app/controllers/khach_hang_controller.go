package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Customer
// @Security BearerAuth
// @Description Filter customer based on provided filters
// @Tags customer
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/khach-hang [get]
func FilterCustomer(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[db.Khach_hang]

	if err := Filter(&req, &res, c, "khach_hang"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay khach hang thanh cong",
	})
}

// @Summary Create Customer
// @Security BearerAuth
// @Description Create a new customer entry
// @Tags customer
// @Accept  json
// @Produce json
// @Param Discount_Type body requests.Khach_hang_create true "Customer data"
// @Router /api/v1/khach-hang [post]
func CreateCustomer(c *gin.Context) {
	var req requests.Khach_hang_create
	var res responses.Khach_hang_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateCustomerExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay khach hang thanh cong",
	})
}

// @Summary Update Customer
// @Security BearerAuth
// @Description Update an existing customer entry
// @Tags customer
// @Accept  json
// @Produce json
// @Param Customer body requests.Khach_hang_update true "Updated customer data"
// @Router /api/v1/khach-hang [put]
func UpdateCustomer(c *gin.Context) {
	var req requests.Khach_hang_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateCustomerExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat khach hang thanh cong",
	})
}


// @Summary Delete Customer
// @Security BearerAuth
// @Description Delete an existing customer entry
// @Tags customer
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "Customer ID to be deleted"
// @Router /api/v1/khach-hang/{id} [delete]
func DeleteCustomer(c *gin.Context) {
	var req requests.Khach_hang_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteCustomerExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa khach hang thanh cong",
	})
}