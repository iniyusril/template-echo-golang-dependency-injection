package domain

import  "gorm.io/gorm"


type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company `gorm:"foreignKey:CompanyID"`
}

