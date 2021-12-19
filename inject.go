//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/iniyusril/template/app"
	"github.com/iniyusril/template/controller"
	"github.com/iniyusril/template/repository"
	"github.com/iniyusril/template/service"
	"github.com/labstack/echo/v4"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	// wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	// wire.Bind(new(service.CategoryService), new(*repository.CategoryRepositoryImpl)),
	controller.NewCategoryController,
	// wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(app.NewDB, validator.New, categorySet, app.NewRouter)
	return nil
}
