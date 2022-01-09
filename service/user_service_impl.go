package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/iniyusril/template/exception"
	"github.com/iniyusril/template/helper"
	"github.com/iniyusril/template/model/domain"
	"github.com/iniyusril/template/model/web"
	"github.com/iniyusril/template/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name: request.Name,
		Company: domain.Company{
			ID: request.Company.ID,
		},
	}

	userRes := service.UserRepository.Save(tx, user)

	return helper.ToUserResponse(userRes)

}

func (service *UserServiceImpl) Update(request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userData := domain.User{
		Model: gorm.Model{ID: uint(request.ID)},
		Name:  request.Name,
	}

	user := service.UserRepository.Update(tx, userData)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(userId int) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(tx, user)
}

func (service *UserServiceImpl) FindById(userId int) web.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll() []web.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(tx)

	return helper.ToUserResponses(users)
}
