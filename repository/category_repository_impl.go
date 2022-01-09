package repository

import (
	"context"

	"github.com/iniyusril/template/helper"
	"github.com/iniyusril/template/model/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category {

	tx.Create(&category)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category {

	id := category.ID

	categoryNew := domain.Category{}
	if err := tx.Where("id = ?", id).First(&categoryNew).Error; err != nil {
		helper.PanicIfError(err)
	}
	categoryNew.Name = category.Name

	tx.Save(categoryNew)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, category domain.Category) {

	var categoryNew domain.Category

	if err := tx.Where("id = ?", category.ID).First(&categoryNew).Error; err != nil {
		helper.PanicIfError(err)
	}
	tx.Delete(&categoryNew)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, categoryId int) (domain.Category, error) {

	var category domain.Category

	if err := tx.Where("id = ?", categoryId).First(&category).Error; err != nil {
		helper.PanicIfError(err)
	}

	return category, nil

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Category {
	var categories []domain.Category
	if err := tx.Find(&categories).Error; err != nil {
		helper.PanicIfError(err)
	}
	return categories

}
