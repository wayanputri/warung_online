package handler

import (
	"strconv"
	"strings"
	"warung_online/app/middleware"
	imageproduct "warung_online/features/imageProduct"
	"warung_online/features/structsEntity"
	"warung_online/helper"

	"github.com/labstack/echo/v4"
)

type ImageProductHandler struct {
	imageProductHandler imageproduct.ImageProductServiceInterface
}

func (handler *ImageProductHandler) Add(c echo.Context)error{
	idUser:=middleware.ExtractTokenUserId(c)
	link,errUp:=helper.UploadImage(c)
	if errUp != nil{
		return helper.Forbidden(c,"error uploud file",nil)
	}
	id:=c.Param("product_id")
	idConv,errCon:=strconv.Atoi(id)
	if errCon != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}

	var imageEntity structsEntity.ImageProductEntity
	imageEntity.Link = link
	imageEntity.ProductID = uint(idConv)

	data,err:=handler.imageProductHandler.Add(imageEntity,uint(idUser))
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.FailedRequest(c,err.Error(),nil)
		}else{
			return helper.InternalError(c,err.Error(),nil)
		}
	}
	response:=structsEntity.ImageProductEntityToResponse(data)
	return helper.SuccessCreate(c,"success insert image product",response)
}
func (handler *ImageProductHandler) Edit(c echo.Context)error{
	idUser:=middleware.ExtractTokenUserId(c)
	var image structsEntity.ImageProductEntity
	errBind:=c.Bind(&image)
	if errBind != nil{
		return helper.FailedRequest(c,"error bind data",nil)
	}
	id:=c.Param("image_id")
	idCon,errCon:=strconv.Atoi(id)
	if errCon != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	idProduct:=c.Param("product_id")
	idConvProduct,errConProduct:=strconv.Atoi(idProduct)
	if errConProduct != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	link,errUp:=helper.UploadImage(c)
	if errUp != nil{
		return helper.FailedRequest(c,"err upload data",nil)
	}
	image.Link = link
	image.ProductID = uint(idConvProduct)
	data,err:=handler.imageProductHandler.Edit(image,uint(idCon),uint(idUser))
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.FailedRequest(c,err.Error(),nil)
		}else{
			return helper.InternalError(c,err.Error(),nil)
		}
	}
	response:=structsEntity.ImageProductEntityToResponse(data)
	return helper.SuccessCreate(c,"success update image product",response)
}

func (handler *ImageProductHandler) Delete(c echo.Context)error{
	id:=c.Param("image_id")
	idCon,errCon:=strconv.Atoi(id)
	if errCon != nil{
		return helper.FailedRequest(c,"id not valid",nil)
	}
	err:=handler.imageProductHandler.Delete(uint(idCon))
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	return helper.SuccessWithOutData(c,"delete successfully")
}

func New(image imageproduct.ImageProductServiceInterface)*ImageProductHandler{
	return &ImageProductHandler{
		imageProductHandler: image,
	}
}