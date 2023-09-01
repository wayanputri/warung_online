package handler

import (
	"strconv"
	imagereview "warung_online/features/imageReview"
	"warung_online/features/structsEntity"
	"warung_online/helper"

	"github.com/labstack/echo/v4"
)

type ImageReviewHandler struct {
	imageReviewHandler imagereview.ImageReviewServiceInterface
}
func (handler *ImageReviewHandler)Add(c echo.Context)error{
	var entityImage structsEntity.ImageReviewEntity
	id:=c.Param("idReview")
	idConv,erConv:=strconv.Atoi(id)
	if erConv != nil{
		return helper.FailedRequest(c, "id not valid", nil)
	}

	entityImage.ReviewID=uint(idConv)
	link,errUploud:=helper.UploadImage(c)
	if errUploud != nil{
		return helper.FailedRequest(c,"error uploud image",nil)
	}
	entityImage.Link=link
	data,err:=handler.imageReviewHandler.Add(entityImage)
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	response:=structsEntity.ImageReviewEntityToResponse(data)
	return helper.SuccessCreate(c,"success create image review",response)
}

func (handler *ImageReviewHandler)Delete(c echo.Context)error{
	id:=c.Param("idImageReview")
	idConv,erConv:=strconv.Atoi(id)
	if erConv != nil{
		return helper.FailedRequest(c, "id not valid", nil)
	}
	err:=handler.imageReviewHandler.Delete(uint(idConv))
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	return helper.SuccessWithOutData(c,"success delete image review")
}

func New(image imagereview.ImageReviewServiceInterface)*ImageReviewHandler{
	return &ImageReviewHandler{
		imageReviewHandler: image,
	}
}