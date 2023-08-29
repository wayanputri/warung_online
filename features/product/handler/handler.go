package handler

import (
	"strconv"
	"strings"
	"warung_online/app/middleware"
	"warung_online/features/product"
	"warung_online/features/structsEntity"
	"warung_online/helper"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productHandler product.ProductServiceInterface
}

func (handler *ProductHandler) Edit(c echo.Context) error{
	id:=middleware.ExtractTokenUserId(c)
	var product structsEntity.ProductEntity
	errBind:=c.Bind(&product)
	if errBind != nil{
		return helper.FailedRequest(c, "error bind data", nil)
	}
	idProduct:=c.Param("id_product")
	idConv,errConv:=strconv.Atoi(idProduct)
	if errConv != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	product.UserID = uint(id)
	err:=handler.productHandler.Edit(product,uint(idConv))
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.FailedRequest(c,err.Error(),nil)
		}else{
			return helper.InternalError(c,"failed update data "+err.Error(),nil)
		}
	}
	return helper.SuccessWithOutData(c,"success update product")
}

func (handler *ProductHandler) Delete(c echo.Context) error{
	id:=middleware.ExtractTokenUserId(c)
	var product structsEntity.ProductEntity
	errBind:=c.Bind(&product)
	if errBind != nil{
		return helper.FailedRequest(c, "error bind data", nil)
	}
	idProduct:=c.Param("id_product")
	idConv,errConv:=strconv.Atoi(idProduct)
	if errConv != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	err:=handler.productHandler.Delete(uint(id),uint(idConv))
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.FailedRequest(c,err.Error(),nil)
		}else{
			return helper.InternalError(c,"failed delete data "+err.Error(),nil)
		}
	}
	return helper.SuccessWithOutData(c,"success delete product")
}

func (handler *ProductHandler) Add(c echo.Context) error{
	id:=middleware.ExtractTokenUserId(c)
	var product structsEntity.ProductEntity
	errBind:=c.Bind(&product)
	if errBind != nil{
		return helper.FailedRequest(c, "error bind data", nil)
	}

	product.UserID = uint(id)
	err:=handler.productHandler.Add(product)
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.FailedRequest(c,err.Error(),nil)
		}else{
			return helper.InternalError(c,"failed insert data "+err.Error(),nil)
		}
	}
	return helper.SuccessCreate(c,"success create product",nil)
}

func (handler *ProductHandler) GetAll(c echo.Context)error{
	name:=c.QueryParam("name")
	data,err:= handler.productHandler.GetAll(name)
	if err != nil{
		return helper.InternalError(c,"failed get all data "+err.Error(),nil)
	}
	var response []structsEntity.ProductResponseAll
	for _,value:= range data{
		response= append(response, structsEntity.ProductEntityToResponseAll(value)) 
	}
	
	return helper.Success(c,"success get product",response)
}

func (handler *ProductHandler) GetById(c echo.Context)error{
	idProduct:=c.Param("id_product")
	idConv,errConv:=strconv.Atoi(idProduct)
	if errConv != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	data,err:= handler.productHandler.SelectById(uint(idConv))
	if err != nil{
		return helper.InternalError(c,"failed get data by id"+err.Error(),nil)
	}
	dataEntity:= structsEntity.ProductEntityToResponse(data)
	return helper.Success(c,"success get product",dataEntity)
}

func New(product product.ProductServiceInterface) *ProductHandler{
	return &ProductHandler{
		productHandler: product,
	}
}