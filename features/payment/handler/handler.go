package handler

import (
	"strconv"
	"warung_online/features/payment"
	"warung_online/features/structsEntity"
	"warung_online/helper"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct{
	paymentHandler payment.PaymentServiceInterface
}

func (handler *PaymentHandler)Add(c echo.Context)error{
	id:=c.Param("transaction_id")
	idConv,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	var request structsEntity.PaymentRequest
	errBind:=c.Bind(&request)
	if errBind != nil{
		return helper.FailedRequest(c, "error bind data", nil)
	}
	inputEntity:=structsEntity.PaymentRequestToEntity(request)
	inputEntity.TransactionFinalID=uint(idConv)
	data,err:=handler.paymentHandler.Add(inputEntity)
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	response:=structsEntity.PaymentEntityToResponse(data)
	return helper.SuccessCreate(c,"success create payment",response)
}

func (handler *PaymentHandler) Notification(c echo.Context)error{
	var notificationPayload map[string]interface{}

	err := c.Bind(&notificationPayload)
	if err != nil {

		return helper.FailedRequest(c,"gagal bind data",nil)
	}
	data,errNotification:=handler.paymentHandler.Notification(notificationPayload)
	if errNotification !=nil{
		return helper.InternalError(c,errNotification.Error(),nil)
	}
	response:=structsEntity.PaymentEntityToResponse(data)
	return helper.Success(c,"success notification",response)
}
func New(handler payment.PaymentServiceInterface)*PaymentHandler{
	return &PaymentHandler{
		paymentHandler: handler,
	}
}