package handler

import (
	"strings"
	"warung_online/app/middleware"
	"warung_online/features/structsEntity"
	"warung_online/features/user"
	"warung_online/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userHandler user.UserServiceInterface
}

func (handler *UserHandler) GetById(c echo.Context) error{
	id:=middleware.ExtractTokenUserId(c)
	data,err:= handler.userHandler.SelectById(uint(id))
	if err != nil{
		return helper.InternalError(c,"get user failed",nil)
	}
	response:=structsEntity.UserEntityToResponse(data)
	return helper.Success(c,"success get user",response)
}

func (handler *UserHandler) GetAll(c echo.Context) error{
	
	data,err := handler.userHandler.GetAll()
	if err != nil{
		return helper.InternalError(c,"get user failed: "+err.Error(),nil)
	}

	var users []structsEntity.UserResponseAll
	for _,value := range data{
		users = append(users, structsEntity.UserEntityToResponseAll(value))
	}
	return helper.Success(c,"success get all data ",users)
}

func (handler *UserHandler) Login(c echo.Context) error{
	var login structsEntity.Login
	errBind:=c.Bind(&login)
	if errBind != nil{
		return helper.FailedRequest(c,"failed bind data",nil)
	}
	user:=structsEntity.LoginToUserEntity(login)
	id,err := handler.userHandler.Login(user)
	if err != nil{
		return helper.InternalError(c,"internal server error: "+err.Error(),nil)
	}
	token,errToken:=middleware.CreateToken(int(id))
	if errToken != nil{
		return helper.Forbidden(c,"failed create token",nil)
	}
	return helper.SuccessCreate(c,"success login",token)

}

func (handler *UserHandler) CreateUser(c echo.Context) error{
	var user structsEntity.UserRequest
	errBind:=c.Bind(&user)
	if errBind != nil{
		return helper.FailedNotFound(c, "error bind user", nil)
	}
	userEntity:=structsEntity.UserRequestToEntity(user)
	
	id,err:=handler.userHandler.Add(userEntity)
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	return helper.SuccessCreate(c,"success create data",id)
}

func (handler *UserHandler) Edit(c echo.Context) error{
	id:=middleware.ExtractTokenUserId(c)
	var user structsEntity.UserRequest
	errBind:=c.Bind(&user)
	if errBind != nil{
		return helper.FailedRequest(c, "error bind data",nil)
	}
	userEntity:=structsEntity.UserRequestToEntity(user)
	err:=handler.userHandler.Edit(userEntity,uint(id))
	if err != nil{
		return helper.InternalError(c,"error update data: "+err.Error(),nil)
	}
	return helper.Success(c,"success update data",nil)
}

func (handler *UserHandler) Delete(c echo.Context)error{
	id := middleware.ExtractTokenUserId(c)
	err:= handler.userHandler.Delete(uint(id))
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	return helper.SuccessWithOutData(c,"delete successfully")
}

func (handler *UserHandler) Upgrade(c echo.Context) error{
	id := middleware.ExtractTokenUserId(c)
	var user structsEntity.UserRequest
	link,errLink:=helper.UploadImage(c)
	if errLink != nil{
		return helper.FailedRequest(c,"error uploud data",nil)
	}
	userEntity :=structsEntity.UserRequestToEntity(user)
	userEntity.DataUpdate = link
	userEntity.Role = "pedagang"

	err := handler.userHandler.Upgrade(userEntity,uint(id))
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.FailedRequest(c,"failed: "+err.Error(),nil)
		}else{
			return helper.InternalError(c,"failed upgrade data"+err.Error(),nil)
		}
	}
	return helper.SuccessWithOutData(c,"success upgrade user")

}

func (handler *UserHandler) EditProfil(c echo.Context)error{
	id := middleware.ExtractTokenUserId(c)
	var user structsEntity.UserRequest
	link,errLink:=helper.UploadImage(c)
	if errLink != nil{
		return helper.FailedRequest(c,"error uploud data",nil)
	}
	userEntity :=structsEntity.UserRequestToEntity(user)
	userEntity.Profil = link

	err := handler.userHandler.UpdateProfil(userEntity,uint(id))
	if err != nil{
		return helper.InternalError(c,"failed update profil"+err.Error(),nil)
	}
	return helper.SuccessWithOutData(c,"success update profil user")
}
func New(user user.UserServiceInterface) *UserHandler{
	return &UserHandler{
		userHandler: user,
	}
}