package controller

import (
	"github.com/iniyusril/template/model/web"
	"github.com/iniyusril/template/service"
	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) Create(c echo.Context) error {
	userCreateRequest := web.CreateUserRequest{}
	c.Bind(&userCreateRequest)

	categoryResponse := controller.userService.Create(c.Request().Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return c.JSON(webResponse.Code, webResponse)
	// helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(c echo.Context) error {
	userUpdateRequest := web.UpdateUserRequest{}

	// helper.ReadFromRequestBody(request, &userUpdateRequest)

	c.Bind(&userUpdateRequest)

	id := c.Param("userId")

	userUpdateRequest.ID = id

	userResponse := controller.userService.Update(c.Request().Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return c.JSON(webResponse.Code, webResponse)
	// helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(c echo.Context) error {
	id := c.Param("userId")

	controller.userService.Delete(c.Request().Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return c.JSON(webResponse.Code, webResponse)

	// helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(c echo.Context) error {
	// userId := params.ByName("categoryId")
	userId := c.Param("userId")

	categoryResponse := controller.userService.FindById(c.Request().Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return c.JSON(webResponse.Code, webResponse)

	// helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(c echo.Context) error {
	categoryResponses := controller.userService.FindAll(c.Request().Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	return c.JSON(webResponse.Code, webResponse)

	// helper.WriteToResponseBody(writer, webResponse)
}
