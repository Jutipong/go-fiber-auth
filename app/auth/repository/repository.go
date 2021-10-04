package repository

import (
	"auth/app/auth/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IRepository interface {
	Inquiry_Auth(userName string) (model.Auth, error)
	Create_UserAndAuth(auth *model.Auth, uaser *model.User) error
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
	err = r.db.Preload("User", "IsActive = ?", true).
		Find(&result, "UserName = ?", userName).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *repository) Create_UserAndAuth(auth *model.Auth, user *model.User) error {
	tx := r.db.Debug().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Check UserName Duplicate
	userTotal := int64(0)
	err := r.db.Model(&model.Auth{}).Where(&model.Auth{UserName: auth.UserName}).Count(&userTotal).Error
	if err != nil {
		return err
	}
	if userTotal > 0 {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("UserName:%s existing in table: Auth", auth.UserName))
	}

	// Table User
	err = tx.Debug().Omit("UpdateDate", "UpdateBy", "Last").Create(user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Table Auth
	err = tx.Omit("UpdateDate", "UpdateBy").Create(auth).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
