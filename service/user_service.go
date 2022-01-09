package service

import "github.com/iniyusril/template/model/web"

type UserService interface {
	Create( request web.UserCreateRequest) web.UserResponse
	Update( request web.UserUpdateRequest) web.UserResponse
	Delete(userId int)
	FindById( userId int) web.UserResponse
	FindAll() []web.UserResponse
}

