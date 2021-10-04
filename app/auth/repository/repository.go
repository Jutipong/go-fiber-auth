package repository

import (
	"auth/app/auth/model"

	"gorm.io/gorm"
)

type IRepository interface {
	Inquiry_Auth(userName string) (model.Auth, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &repository{db: db}
}

func (r *repository) Inquiry_Auth(userName string) (result model.Auth, err error) {
	//## หากต้องการดู string query ที่ gorm generate ให้ให้ใช่ .Debuger()
	//## ตัวอย่าง err = r.db.Debug().Find(&result, "UserName = ?", userName).Error
	err = r.db.Preload("User").Find(&result, "UserName = ?", userName).Error
	if err != nil {
		return result, err
	}
	return result, nil
}
