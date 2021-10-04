package model

import (
	"database/sql"
	"encoding/base64"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gookit/validate"
)

// ** ใช้สำหรับ query database on gorm.
// ===================== model entity =====================//
type Auth struct {
	Id         string       `gorm:"primaryKey; column:Id"`
	UserName   string       `gorm:"column:UserName" validate:"required" json:"userName"`
	Password   string       `gorm:"column:Password" validate:"required" json:"password"`
	CreateDate time.Time    `gorm:"column:CreateDate;autoCreateTime"`
	CreateBy   string       `gorm:"column:CreateBy; default:System"`
	UpdateDate sql.NullTime `gorm:"column:UpdateDate; autoUpdateTime;"`
	UpdateBy   string       `gorm:"column:UpdateBy"`
	IsActive   bool         `gorm:"column:IsActive; default:true"`
	UserId     string       `gorm:"column:UserId"`
	// ##foreign
	User User `gorm:"references:UserId;foreignKey:UserId;"`
}

func (a *Auth) TableName() string {
	return "Auth"
}

// Validation struct Auth
func (a *Auth) Validation(c *fiber.Ctx) error {
	// Get authorization header
	authReq := c.Get(fiber.HeaderAuthorization)

	// Check if the header contains content besides "basic".
	if len(authReq) <= 6 || strings.ToLower(authReq[:5]) != "basic" {
		return fiber.ErrUnauthorized
	}

	// Decode the header contents
	raw, err := base64.StdEncoding.DecodeString(authReq[6:])
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// Get the credentials
	creds := utils.UnsafeString(raw)

	// Check if the credentials are in the correct form
	// which is "username:password".
	index := strings.Index(creds, ":")
	if index == -1 {
		return fiber.ErrUnauthorized
	}

	// Get the username and password
	a.UserName = creds[:index]
	a.Password = creds[index+1:]

	v := validate.Struct(a)
	if !v.Validate() {
		return v.Errors
	}
	return nil
}

//##
type User struct {
	UserId     string       `gorm:"primaryKey; column:UserId"`
	Name       string       `gorm:"column:Name"`
	Last       string       `gorm:"column:Last"`
	CreateDate time.Time    `gorm:"column:CreateDate; autoCreateTime"`
	CreateBy   string       `gorm:"column:CreateBy;"`
	UpdateDate sql.NullTime `gorm:"column:UpdateDate; autoUpdateTime"`
	UpdateBy   string       `gorm:"column:UpdateBy"`
	IsActive   bool         `gorm:"column:IsActive; default:true"`
}

func (a *User) TableName() string {
	return "User"
}
