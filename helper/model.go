package helper

import (
	"github.com/iniyusril/template/model/domain"
	"github.com/iniyusril/template/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   int(category.ID),
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:   int(user.ID),
		Name: user.Name,
		Company: web.CompanyResponse{
			ID:   int(user.Company.ID),
			Name: user.Company.Name,
		},
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
