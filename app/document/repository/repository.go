package repository

import (
	"auth/app/document/model"
	"auth/pkg/enum"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IRepository interface {
	Inquiry_TmpDocument(req *model.Request) (model.TmpDocument, error)
	Inquiry_DocumentVersion(req *model.Request) (model.DocumentVersion, error)
	Inquiry_DocumentVersionLog(req *model.Request) (int64, error)
	Update_Document(tmpDoc *model.TmpDocument, docVer *model.DocumentVersion, docVerLog *model.DocumentVersionlog) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &repository{db: db}
}

var tmplateMessage = "Table %s  data not found was condition is DocRefNo: %s"

func (r *repository) Inquiry_TmpDocument(req *model.Request) (result model.TmpDocument, err error) {
	err = r.db.First(&result, &model.TmpDocument{DocRefNo: req.CallbackRefid}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(enum.DataNotFound, fmt.Sprintf(tmplateMessage, "TmpDocument", req.CallbackRefid))
		}
		return
	}

	return result, nil
}

func (r *repository) Inquiry_DocumentVersion(req *model.Request) (result model.DocumentVersion, err error) {
	err = r.db.Find(&result, &model.DocumentVersion{DocRefNo: req.CallbackRefid}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(enum.DataNotFound, fmt.Sprintf(tmplateMessage, "DocumentVersion", req.CallbackRefid))
		}
		return
	}

	return result, nil
}

func (r *repository) Inquiry_DocumentVersionLog(req *model.Request) (result int64, err error) {
	enity := r.db.Model(&model.DocumentVersionlog{}).Where(&model.DocumentVersionlog{DocRefNo: req.CallbackRefid, Version: req.Version}).Count(&result)
	if enity.Error != nil {
		err = enity.Error
		return
	}

	return result, nil
}

func (r *repository) Update_Document(tmpDoc *model.TmpDocument, docVer *model.DocumentVersion, docVerLog *model.DocumentVersionlog) error {

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 1.Update Document
	err := tx.Select("SeqID", "ChronicleID", "Version", "UploadStatus", "UpdatedDate").Updates(tmpDoc).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Select("LastVersion", "UpdatedDate", "UpdatedBy", "DocFullName").Updates(docVer).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2.Create Document Log
	if docVerLog.DocumentID != "" {
		err := tx.Create(docVerLog).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
