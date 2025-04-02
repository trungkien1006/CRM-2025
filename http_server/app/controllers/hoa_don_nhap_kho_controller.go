package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Import Invoice
// @Security BearerAuth
// @Description Filter import invoice based on provided filters
// @Tags import invoice
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /api/v1/hoa-don-nhap-kho [get]
func FilterImportInvoice(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.Hoa_don_nhap_kho_filter]

	if err := Filter(&req, &res, c, "hoa_don_nhap_kho"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay hoa don nhap kho thanh cong",
	})
}

// @Summary Create Import Invoice
// @Security BearerAuth
// @Description Create a new import invoice entry
// @Tags import invoice
// @Accept  json
// @Produce json
// @Param Import_Invoice body requests.Hoa_don_nhap_kho_create true "Import Invoice data"
// @Router /api/v1/hoa-don-nhap-kho [post]
func CreateImportInvoice(c *gin.Context) {
	var req requests.Hoa_don_nhap_kho_create
	var res responses.Hoa_don_nhap_kho_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateImportInvoice(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them hoa don nhap kho thanh cong",
	})
}

// @Summary Update Import Invoice
// @Security BearerAuth
// @Description Update import invoice
// @Tags import invoice
// @Accept  json
// @Produce json
// @Param Update_Import_Invoice body requests.Hoa_don_nhap_kho_update true "Update Import Invoice Data"
// @Router /api/v1/hoa-don-nhap-kho [put]
func UpdateImportInvoice(c *gin.Context) {
	var req requests.Hoa_don_nhap_kho_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}
	
	if err := dao.UpdateImportInvoiceExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat hoa don nhap kho thanh cong",
	})
}

// @Summary Lock Import Invoice
// @Security BearerAuth
// @Description Lock import invoice
// @Tags import invoice
// @Accept  json
// @Produce json
// @Param Lock_Import_Invoice body requests.Hoa_don_nhap_kho_lock true "Lock Import Invoice Data"
// @Router /api/v1/hoa-don-nhap-kho/lock [patch]
func LockImportInvoice(c *gin.Context) {
	var req requests.Hoa_don_nhap_kho_lock

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}
	
	if err := dao.LockImportInvoiceExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat khoa hoa don nhap kho thanh cong",
	})
}

// @Summary Import Debt Payment
// @Security BearerAuth
// @Description Pay Debt for import invoice
// @Tags import invoice
// @Accept  json
// @Produce json
// @Param Debt_Payment body requests.Tra_no_nhap_kho_request true "Debt Payment Data"
// @Router /api/v1/hoa-don-nhap-kho/tra-no [patch]
func ImportDebtPayment(c *gin.Context) {
	var req requests.Tra_no_nhap_kho_request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.ImportDebtPaymentExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat tien no nhap kho thanh cong",
	})
}

// @Summary Import Return Product
// @Security BearerAuth
// @Description Return Product for import invoice
// @Tags import invoice
// @Accept  json
// @Produce json
// @Param Return_Product body requests.Tra_hang_nhap_kho_request true "Return Product Data"
// @Router /api/v1/hoa-don-nhap-kho/tra-hang [patch]
func ReturnImportProduct(c *gin.Context) {
	var req requests.Tra_hang_nhap_kho_request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.ReturnImportProductExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat tra hang nhap kho thanh cong",
	})
}