package repository

import (
	"github.com/iniyusril/template/helper"
	"github.com/iniyusril/template/model/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(tx *gorm.DB, user domain.User) domain.User {
	tx.Create(&user)
	return user
}

func (u *UserRepositoryImpl) Update(tx *gorm.DB, user domain.User) domain.User {
	var newUser domain.User
	if err := tx.Where("id = ?", user.ID).First(&newUser).Error; err != nil {
		helper.PanicIfError(err)
	}

	newUser.Name = user.Name

	tx.Omit(clause.Associations).Save(&newUser)
	return newUser
}

func (u *UserRepositoryImpl) Delete(tx *gorm.DB, user domain.User) {
	var newUser domain.User
	if err := tx.Where("id = ?", user.ID).First(&newUser).Error; err != nil {
		helper.PanicIfError(err)
	}
	tx.Delete(&newUser)
}

func (u *UserRepositoryImpl) FindById(tx *gorm.DB, userId int) (domain.User, error) {
	var newUser domain.User
	if err := tx.Preload("Company").Where("id = ?", userId).First(&newUser).Error; err != nil {
		helper.PanicIfError(err)
	}
	return newUser, nil
}

func (u *UserRepositoryImpl) FindAll(tx *gorm.DB) []domain.User {
	var users []domain.User
	if err := tx.Preload("Company").Find(&users).Error; err != nil {
		helper.PanicIfError(err)
	}

	return users
}
