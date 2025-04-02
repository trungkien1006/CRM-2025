package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Export Invoice
// @Security BearerAuth
// @Description Filter export invoice based on provided filters
// @Tags export invoice
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/hoa-don-xuat-kho [get]
func FilterExportInvoice(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.Hoa_don_xuat_kho_filter]

	if err := Filter(&req, &res, c, "hoa_don_xuat_kho"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay hoa don xuat kho thanh cong",
	})
}

// @Summary Create Export Invoice
// @Security BearerAuth
// @Description Create a new export invoice entry
// @Tags export invoice
// @Accept  json
// @Produce json
// @Param Export_Invoice body requests.Hoa_don_xuat_kho_create true "Export Invoice data"
// @Router /api/v1/hoa-don-xuat-kho [post]
func CreateExportInvoice(c *gin.Context) {
	var req requests.Hoa_don_xuat_kho_create
	var res responses.Hoa_don_xuat_kho_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateExportInvoice(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them hoa don xuat kho thanh cong",
	})
}

// @Summary Update Export Invoice
// @Security BearerAuth
// @Description Update export invoice
// @Tags export invoice
// @Accept  json
// @Produce json
// @Param Update_Export_Invoice body requests.Hoa_don_xuat_kho_update true "Update Export Invoice Data"
// @Router /api/v1/hoa-don-xuat-kho [put]
func UpdateExportInvoice(c *gin.Context) {
	var req requests.Hoa_don_xuat_kho_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}
	
	if err := dao.UpdateExportInvoiceExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat hoa don xuat kho thanh cong",
	})
}

// @Summary Lock Export Invoice
// @Security BearerAuth
// @Description Lock export invoice
// @Tags export invoice
// @Accept  json
// @Produce json
// @Param Lock_Export_Invoice body requests.Hoa_don_xuat_kho_lock true "Lock Export Invoice Data"
// @Router /api/v1/hoa-don-xuat-kho/lock [patch]
func LockExportInvoice(c *gin.Context) {
	var req requests.Hoa_don_xuat_kho_lock

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}
	
	if err := dao.LockExportInvoiceExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat khoa hoa don xuat kho thanh cong",
	})
}

// @Summary Export Debt Payment
// @Security BearerAuth
// @Description Pay Debt for export invoice
// @Tags export invoice
// @Accept  json
// @Produce json
// @Param Debt_Payment body requests.Tra_no_xuat_kho_request true "Debt Payment Data"
// @Router /api/v1/hoa-don-xuat-kho/tra-no [patch]
func ExportDebtPayment(c *gin.Context) {
	var req requests.Tra_no_xuat_kho_request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.ExportDebtPaymentExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat tien no xuat kho thanh cong",
	})
}

// @Summary Export Return Product
// @Security BearerAuth
// @Description Return Product for export invoice
// @Tags export invoice
// @Accept  json
// @Produce json
// @Param Return_Product body requests.Tra_hang_xuat_kho_request true "Return Product Data"
// @Router /api/v1/hoa-don-xuat-kho/tra-hang [patch]
func ReturnExportProduct(c *gin.Context) {
	var req requests.Tra_hang_xuat_kho_request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.ReturnExportProductExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat tra hang xuat kho thanh cong",
	})
}