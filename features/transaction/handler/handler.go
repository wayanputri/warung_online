package handler

import (
	"strconv"
	"warung_online/app/middleware"
	"warung_online/features/structsEntity"
	"warung_online/features/transaction"
	"warung_online/helper"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionHandler transaction.TransactionServiceInterface
}

func (handler *TransactionHandler) Add(c echo.Context)error{
	var input structsEntity.TransactionRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return helper.FailedNotFound(c, "error bind data", nil)
	}
	inputCore := structsEntity.TransactionRequestToEntity(input)
	idProduct:=c.Param("product_id")
	idConv,errConv:= strconv.Atoi(idProduct)
	if errConv != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	idUser:=middleware.ExtractTokenUserId(c)
	inputCore.ProductID = uint(idConv)
	inputCore.UserID = uint(idUser)
	data,err:= handler.transactionHandler.Add(inputCore)
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	response:= structsEntity.TransactionEntityToResponse(data)
	return helper.SuccessCreate(c,"success create transaction",response)
}

func (handler *TransactionHandler) GetAll(c echo.Context)error{
	idUser:=middleware.ExtractTokenUserId(c)
	data,err:=handler.transactionHandler.GetAll(uint(idUser))
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	var response []structsEntity.TransactionResponseAll
	for _,value:=range data{
		response = append(response, structsEntity.TransactionEntityToResponseAll(value))
	}
	return helper.Success(c,"success get all transaction",response)
}

func (handler *TransactionHandler)GetById(c echo.Context)error{
	idUser:=middleware.ExtractTokenUserId(c)
	id:=c.Param("id")
	idConv,errCon:=strconv.Atoi(id)
	if errCon != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	data,err:=handler.transactionHandler.GetById(uint(idConv),uint(idUser))
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	response:=structsEntity.TransactionEntityToResponse(data)
	return helper.Success(c,"success get transaction by id",response)

}

func (handler *TransactionHandler)Edit(c echo.Context)error{
	var input structsEntity.TransactionRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return helper.FailedNotFound(c, "error bind data", nil)
	}
	id:=c.Param("id")
	idConv,errCon:=strconv.Atoi(id)
	if errCon != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	idProduct:=c.Param("product_id")
	idCon,errC:=strconv.Atoi(idProduct)
	if errC != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	idUser:=middleware.ExtractTokenUserId(c)
	inputCore := structsEntity.TransactionRequestToEntity(input)	
	inputCore.UserID = uint(idUser)
	inputCore.Id = uint(idConv)
	inputCore.ProductID = uint(idCon)
	data,err:=handler.transactionHandler.Edit(inputCore)
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	response:=structsEntity.TransactionEntityToResponse(data)
	return helper.Success(c,"success update transaction",response)
}

func (handler *TransactionHandler)Delete(c echo.Context)error{
	idUser:=middleware.ExtractTokenUserId(c)
	id:=c.Param("id")
	idConv,errCon:=strconv.Atoi(id)
	if errCon != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	err:=handler.transactionHandler.Delete(uint(idConv),uint(idUser))
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	return helper.SuccessWithOutData(c,"success deleted transaction")
}

func New(handler transaction.TransactionServiceInterface) *TransactionHandler{
	return &TransactionHandler{
		transactionHandler: handler,
	}
}