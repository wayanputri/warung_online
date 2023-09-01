package router

import (
	"warung_online/app/middleware"
	dataUser "warung_online/features/user/data"
	handlerUser "warung_online/features/user/handler"
	serviceUser "warung_online/features/user/service"

	dataProduct "warung_online/features/product/data"
	handlerProduct "warung_online/features/product/handler"
	serviceProduct "warung_online/features/product/service"

	dataImageProduct "warung_online/features/imageProduct/data"
	handlerImageProduct "warung_online/features/imageProduct/handler"
	serviceImageProduct "warung_online/features/imageProduct/service"

	dataTransaction "warung_online/features/transaction/data"
	handlerTransaction "warung_online/features/transaction/handler"
	serviceTransaction "warung_online/features/transaction/service"

	dataTransactionDetil "warung_online/features/transactionDetil/data"
	handlerTransactionDetil "warung_online/features/transactionDetil/handler"
	serviceTransactionDetil "warung_online/features/transactionDetil/service"

	dataPayment "warung_online/features/payment/data"
	handlerPayment "warung_online/features/payment/handler"
	servicePayment "warung_online/features/payment/service"

	dataReview "warung_online/features/review/data"
	handlerReview "warung_online/features/review/handler"
	serviceReview "warung_online/features/review/service"

	dataImageReview "warung_online/features/imageReview/data"
	handlerImageReview "warung_online/features/imageReview/handler"
	serviceImageReview "warung_online/features/imageReview/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo){
	dUser := dataUser.New(db)
	sUser := serviceUser.New(dUser)
	hUser := handlerUser.New(sUser)

	c.POST("/users",hUser.CreateUser)
	c.POST("/login",hUser.Login)
	c.GET("/users",hUser.GetAll)
	c.GET("/user",hUser.GetById,middleware.JWTMiddleware())
	c.PUT("/users",hUser.Edit,middleware.JWTMiddleware())
	c.DELETE("/users",hUser.Delete,middleware.JWTMiddleware())
	c.PUT("/upgrade",hUser.Upgrade,middleware.JWTMiddleware())
	c.PUT("/profil",hUser.EditProfil,middleware.JWTMiddleware())

	dProduct := dataProduct.New(db)
	sProduct := serviceProduct.New(dProduct)
	hProduct := handlerProduct.New(sProduct)

	c.POST("/products",hProduct.Add,middleware.JWTMiddleware())
	c.PUT("/products/:id_product",hProduct.Edit,middleware.JWTMiddleware())
	c.DELETE("/products/:id_product",hProduct.Delete,middleware.JWTMiddleware())
	c.GET("/products",hProduct.GetAll)
	c.GET("/products/:id_product",hProduct.GetById)

	dImageProduct := dataImageProduct.New(db)
	sImageProduct := serviceImageProduct.New(dImageProduct)
	hImageProduct := handlerImageProduct.New(sImageProduct)

	c.POST("/products/:product_id/imageproducts",hImageProduct.Add,middleware.JWTMiddleware())
	c.PUT("/products/:product_id/imageproducts/:image_id",hImageProduct.Edit,middleware.JWTMiddleware())
	c.DELETE("/imageproducts/:image_id",hImageProduct.Delete,middleware.JWTMiddleware())

	dTransaction := dataTransaction.New(db)
	sTransaction := serviceTransaction.New(dTransaction)
	hTransaction := handlerTransaction.New(sTransaction)

	c.POST("/products/:product_id/transactions",hTransaction.Add,middleware.JWTMiddleware())
	c.GET("/transactions",hTransaction.GetAll,middleware.JWTMiddleware())
	c.GET("/transactions/:id",hTransaction.GetById,middleware.JWTMiddleware())
	c.PUT("/products/:product_id/transactions/:id",hTransaction.Edit,middleware.JWTMiddleware())
	c.DELETE("/transactions/:id",hTransaction.Delete,middleware.JWTMiddleware())

	dTransactionDetil := dataTransactionDetil.New(db)
	sTransactionDetil := serviceTransactionDetil.New(dTransactionDetil)
	hTransactionDetil := handlerTransactionDetil.New(sTransactionDetil)

	c.POST("/transactiondetils",hTransactionDetil.Add,middleware.JWTMiddleware())

	dPayment := dataPayment.New(db)
	sPayment := servicePayment.New(dPayment)
	hPayment := handlerPayment.New(sPayment)

	c.POST("/transactiondetils/:transaction_id/payments",hPayment.Add,middleware.JWTMiddleware())
	c.POST("/notifications",hPayment.Notification)

	dReview := dataReview.New(db)
	sReview := serviceReview.New(dReview)
	hReview := handlerReview.New(sReview)
	c.POST("/products/:product_id/payments/:payment_id/reviews",hReview.Add,middleware.JWTMiddleware())
	c.DELETE("/reviews/:id",hReview.Delete,middleware.JWTMiddleware())

	dImageReview := dataImageReview.New(db)
	sImageReview := serviceImageReview.New(dImageReview)
	hImageReview := handlerImageReview.New(sImageReview)

	c.POST("/imagereviews/:idReview",hImageReview.Add,middleware.JWTMiddleware())
	c.DELETE("/imagereviews/:idImageReview",hImageReview.Delete,middleware.JWTMiddleware())
}