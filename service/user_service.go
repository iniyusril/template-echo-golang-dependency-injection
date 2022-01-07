package service

import (
	"context"

	"github.com/iniyusril/template/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.CreateUserRequest) web.UserResponse
	Update(ctx context.Context, request web.UpdateUserRequest) web.UserResponse
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
