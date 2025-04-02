package routes

import (
	"net/http"

	"admin-v1/app/controllers"
	"admin-v1/app/middlewares"
	// "admin-v1/app/middlewares"
	_ "admin-v1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// Thêm header CORS cho mỗi request
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                       // Cho phép tất cả các origin
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS") // Các phương thức HTTP cho phép
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")            // Các header cho phép

		// Xử lý preflight OPTIONS
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	r.Static("/public/images", "./public/images")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//version 1.0
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/", middlewares.AuthGuard, middlewares.CheckPermission)
		{
			//thoi gian bao hanh
			auth.GET("/thoi-gian-bao-hanh", controllers.FilterWarrantyTime)
			auth.POST("/thoi-gian-bao-hanh", controllers.CreateWarrantyTime)
			auth.PUT("/thoi-gian-bao-hanh", controllers.UpdateWarrantyTime)
			auth.DELETE("/thoi-gian-bao-hanh/:id", controllers.DeleteWarrantyTime)

			//loai giam gia
			auth.GET("/loai-giam-gia", controllers.FilterDiscountType)
			auth.POST("/loai-giam-gia", controllers.CreateDiscountType)
			auth.PUT("/loai-giam-gia", controllers.UpdateDiscountType)
			auth.DELETE("/loai-giam-gia/:id", controllers.DeleteDiscountType)

			//don vi tinh
			auth.GET("/don-vi-tinh", controllers.FilterUnit)
			auth.POST("/don-vi-tinh", controllers.CreateUnit)
			auth.PUT("/don-vi-tinh", controllers.UpdateUnit)
			auth.DELETE("/don-vi-tinh/:id", controllers.DeleteUnit)

			//loai san pham
			auth.GET("/loai-san-pham", controllers.FilterProductType)
			auth.POST("/loai-san-pham", controllers.CreateProductType)
			auth.PUT("/loai-san-pham", controllers.UpdateProductType)
			auth.DELETE("/loai-san-pham/:id", controllers.DeleteProductType)

			//san pham
			auth.GET("/san-pham", controllers.FilterProduct)
			auth.POST("/san-pham", controllers.CreateProduct)
			auth.PUT("/san-pham", controllers.UpdateProduct)
			auth.DELETE("/san-pham/:id", controllers.DeleteProduct)

			//nhan vien
			auth.GET("/nhan-vien", controllers.FilterEmployee)
			auth.POST("/nhan-vien", controllers.CreateEmployee)
			auth.PUT("/nhan-vien", controllers.UpdateEmployee)
			auth.DELETE("/nhan-vien/:id", controllers.DeleteEmployee)

			//chuc vu
			auth.GET("/chuc-vu", controllers.FilterRole)
			auth.POST("/chuc-vu", controllers.CreateRole)
			auth.PUT("/chuc-vu", controllers.UpdateRole)
			auth.DELETE("/chuc-vu/:id", controllers.DeleteRole)

			//nha phan phoi
			auth.GET("/nha-phan-phoi", controllers.FilterProvider)
			auth.POST("/nha-phan-phoi", controllers.CreateProvider)
			auth.PUT("/nha-phan-phoi", controllers.UpdateProvider)
			auth.DELETE("/nha-phan-phoi/:id", controllers.DeleteProvider)

			//khach hang
			auth.GET("/khach-hang", controllers.FilterCustomer)
			auth.POST("/khach-hang", controllers.CreateCustomer)
			auth.PUT("/khach-hang", controllers.UpdateCustomer)
			auth.DELETE("/khach-hang/:id", controllers.DeleteCustomer)

			//kho
			auth.GET("/kho", controllers.FilterWareHouse)
			auth.POST("/kho", controllers.CreateWareHouse)
			auth.PUT("/kho", controllers.UpdateWareHouse)
			auth.DELETE("/kho/:id", controllers.DeleteWareHouse)

			//quyen
			auth.GET("/quyen/:chuc_vu_id", controllers.GetPermission)
			auth.PATCH("/quyen/modify", controllers.ModifyPermission)

			//hoa don nhap kho
			auth.GET("/hoa-don-nhap-kho", controllers.FilterImportInvoice)
			auth.POST("/hoa-don-nhap-kho", controllers.CreateImportInvoice)
			auth.PUT("/hoa-don-nhap-kho", controllers.UpdateImportInvoice)
			auth.PATCH("/hoa-don-nhap-kho/lock", controllers.LockImportInvoice)
			auth.PATCH("/hoa-don-nhap-kho/tra-no", controllers.ImportDebtPayment)
			auth.PATCH("/hoa-don-nhap-kho/tra-hang", controllers.ReturnImportProduct)

			//hoa don xuat kho
			auth.GET("/hoa-don-xuat-kho", controllers.FilterExportInvoice)
			auth.POST("/hoa-don-xuat-kho", controllers.CreateExportInvoice)
			auth.PUT("/hoa-don-xuat-kho", controllers.UpdateExportInvoice)
			auth.PATCH("/hoa-don-xuat-kho/lock", controllers.LockExportInvoice)
			auth.PATCH("/hoa-don-xuat-kho/tra-no", controllers.ExportDebtPayment)
			auth.PATCH("/hoa-don-xuat-kho/tra-hang", controllers.ReturnExportProduct)

			//chi tiet san pham
			auth.GET("/chi-tiet-san-pham/:product_id", controllers.GetProductDetail)

			//ton kho
			auth.GET("/ton-kho/:ctsp_id", controllers.GetInStockByProductDetailId)

			//cong no
			auth.GET("/cong-no-khach-hang", controllers.FilterCustomerDebt)
			auth.GET("/cong-no-nha-phan-phoi", controllers.FilterProviderDebt)
		}

		//auth
		v1.POST("/dang-nhap", controllers.Login)
		v1.GET("/thong-tin-nhan-vien", controllers.GetMe)

		//test middleware
		// v1.GET("/check-permission", middlewares.CheckPermission, func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"message": "ok",
		// 	})
		// })
	}

	//test
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}