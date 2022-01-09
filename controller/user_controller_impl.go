package controller

import (
	"github.com/iniyusril/template/helper"
	"github.com/iniyusril/template/model/web"
	"github.com/iniyusril/template/service"
	"github.com/labstack/echo/v4"
	"strconv"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController{
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(c echo.Context) error {
	userCreateRequest := web.UserCreateRequest{}
	c.Bind(&userCreateRequest)
	// helper.ReadFromRequestBody(request, &categoryCreateRequest)

	userResponse := controller.UserService.Create(userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.JSON(webResponse.Code, webResponse)}

func (controller *UserControllerImpl) Update(c echo.Context) error {
	userUpdateRequest := web.UserUpdateRequest{}


	c.Bind(&userUpdateRequest)

	userId := c.Param("userId")

	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.ID = id

	userResponse := controller.UserService.Update( userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.JSON(webResponse.Code, webResponse)}

func (controller *UserControllerImpl) Delete(c echo.Context) error {
	userId := c.Param("userId")

	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete( id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return c.JSON(webResponse.Code, webResponse)
}

func (controller *UserControllerImpl) FindById(c echo.Context) error {
	userId := c.Param("userId")

	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById( id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.JSON(webResponse.Code, webResponse)}

func (controller *UserControllerImpl) FindAll(c echo.Context) error {
	userResponses := controller.UserService.FindAll()
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	return c.JSON(webResponse.Code, webResponse)}

