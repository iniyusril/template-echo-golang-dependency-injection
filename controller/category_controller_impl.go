package controller

import (
	"strconv"

	"github.com/iniyusril/template/helper"
	"github.com/iniyusril/template/model/web"
	"github.com/iniyusril/template/service"
	"github.com/labstack/echo/v4"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(c echo.Context) error {
	categoryCreateRequest := web.CategoryCreateRequest{}
	c.Bind(&categoryCreateRequest)
	// helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(c.Request().Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return c.JSON(webResponse.Code, webResponse)
	// helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(c echo.Context) error {
	categoryUpdateRequest := web.CategoryUpdateRequest{}

	// helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	c.Bind(&categoryUpdateRequest)

	categoryId := c.Param("categoryId")

	// categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(c.Request().Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return c.JSON(webResponse.Code, webResponse)
	// helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(c echo.Context) error {
	categoryId := c.Param("categoryId")

	// categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(c.Request().Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return c.JSON(webResponse.Code, webResponse)

	// helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(c echo.Context) error {
	// categoryId := params.ByName("categoryId")
	categoryId := c.Param("categoryId")

	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(c.Request().Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return c.JSON(webResponse.Code, webResponse)

	// helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(c echo.Context) error {
	categoryResponses := controller.CategoryService.FindAll(c.Request().Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	return c.JSON(webResponse.Code, webResponse)

	// helper.WriteToResponseBody(writer, webResponse)
}
