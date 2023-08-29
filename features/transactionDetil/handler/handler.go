package handler

import (
	"net/http"
	"strconv"
	"strings"
	"warung_online/app/middleware"
	"warung_online/features/structsEntity"
	transactiondetil "warung_online/features/transactionDetil"
	"warung_online/helper"

	"github.com/labstack/echo/v4"
)

type TransactionDetilHandler struct {
	transactionDetilHandler transactiondetil.TransactionDetilServiceInterface
}

func (handler *TransactionDetilHandler)Add(c echo.Context)error{
	idUser:=middleware.ExtractTokenUserId(c)
	ids := c.QueryParam("ids")

	if ids == "" {
		return c.String(http.StatusBadRequest, "IDs parameter is missing")
	}

	idStrings  := strings.Split(ids, ",")
	var idList []uint

	for _, idStr := range idStrings {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid ID format")
		}
		idList = append(idList, uint(id))
	}
	var input structsEntity.TransactionFinalEntity
	data,err:=handler.transactionDetilHandler.Add(idList,uint(idUser),input)
	if err != nil{
		return helper.InternalError(c, err.Error(), nil)
	}
	response:=structsEntity.TransactionFinalEntityToResponse(data)
	return helper.SuccessCreate(c,"success create transaction detil",response)
}
func New(handler transactiondetil.TransactionDetilServiceInterface)*TransactionDetilHandler{
	return &TransactionDetilHandler{
		transactionDetilHandler: handler,
	}
}