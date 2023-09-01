package handler

import (
	"strconv"
	"warung_online/app/middleware"
	"warung_online/features/review"
	"warung_online/features/structsEntity"
	"warung_online/helper"

	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	reviewHandler review.ReviewServiceInterface
}

func (handler *ReviewHandler)Add(c echo.Context)error{
	idUser:=middleware.ExtractTokenUserId(c)
	id:=c.Param("product_id")
	idProduct,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return helper.FailedRequest(c, "id not valid", nil)
	}
	idPayment:=c.Param("payment_id")
	idPaymentConv,errPaymentConv:=strconv.Atoi(idPayment)
	if errPaymentConv != nil{
		return helper.FailedRequest(c, "id not valid", nil)
	}
	var request structsEntity.ReviewRequest
	errBind:=c.Bind(&request)
	if errBind != nil{
		return helper.FailedRequest(c,"error bind data",nil)
	}
	input:=structsEntity.ReviewRequestToEntity(request)
	input.ProductID=uint(idProduct)
	input.PaymentID = uint(idPaymentConv)
	data,err:=handler.reviewHandler.Add(input,uint(idUser))
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	response:=structsEntity.ReviewEntityToResponse(data)
	return helper.SuccessCreate(c,"success create review",response)
}

func (handler *ReviewHandler)Delete(c echo.Context)error{
	idUser:=middleware.ExtractTokenUserId(c)
	id:=c.Param("id")
	idReview,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return helper.FailedRequest(c, "id not valid", nil)
	}
	err:=handler.reviewHandler.Delete(uint(idReview),uint(idUser))
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	return helper.SuccessWithOutData(c,"success delete review")
}

func New(review review.ReviewServiceInterface)*ReviewHandler{
	return &ReviewHandler{
		reviewHandler: review,
	}
}