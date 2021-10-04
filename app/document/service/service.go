package service

import (
	"auth/app/document/model"
	"auth/app/document/repository"
	"auth/pkg/utils"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type IService interface {
	UpdateDocument(ctx *fiber.Ctx, req *model.Request) model.Response
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return &service{repo}
}

func (s *service) UpdateDocument(ctx *fiber.Ctx, req *model.Request) (result model.Response) {
	// Validate Model
	if err := req.Validate(); err != nil {
		return model.Response{IsError: true, ErrorMsg: "missing required field"}
	}

	// 1.Inquiry Datas
	tmpDoc, err := s.repo.Inquiry_TmpDocument(req)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return model.Response{IsError: true, ErrorMsg: err.Error()}
	}

	docVer, err := s.repo.Inquiry_DocumentVersion(req)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return model.Response{IsError: true, ErrorMsg: err.Error()}
	}

	total_DocVerLog, err := s.repo.Inquiry_DocumentVersionLog(req)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return model.Response{IsError: true, ErrorMsg: err.Error()}
	}

	// 2.Set Value
	tmpDoc.SeqID = utils.StringEmpty_SetDefault(req.SeqId, "")
	tmpDoc.ChronicleID = utils.StringEmpty_SetDefault(req.ChronicleId, "")
	tmpDoc.Version = utils.StringEmpty_SetDefault(req.Version, "")
	tmpDoc.UploadStatus = "Success"

	docVer.LastVersion = tmpDoc.Version
	docVer.DocFullName = tmpDoc.DocFullName
	docVer.UpdatedBy = "System"

	// 3.Init Value
	var docVerLog model.DocumentVersionlog
	if total_DocVerLog == 0 {
		docVerLog.DocumentID = tmpDoc.DocumentID
		docVerLog.DocTypeID = tmpDoc.DocTypeID
		docVerLog.DocNo = tmpDoc.DocNo
		docVerLog.StartDate = tmpDoc.StartDate
		docVerLog.Enddate = tmpDoc.Enddate
		docVerLog.FileName = tmpDoc.FileName
		docVerLog.DocRefNo = tmpDoc.DocRefNo
		docVerLog.ChronicleID = tmpDoc.ChronicleID
		docVerLog.SeqID = tmpDoc.SeqID
		docVerLog.Version = tmpDoc.Version
		docVerLog.FileSize = tmpDoc.FileSize
		docVerLog.Status = tmpDoc.Status
		docVerLog.UploadStatus = tmpDoc.Status
		docVerLog.ApproveStatus = sql.NullString{}
		docVerLog.CreatedBy = tmpDoc.CreatedBy
		docVerLog.UpdatedBy = tmpDoc.UpdatedBy
		docVerLog.DocFullName = tmpDoc.DocFullName
	}

	// 4. Save data in database.
	err = s.repo.Update_Document(&tmpDoc, &docVer, &docVerLog)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return model.Response{IsError: true, ErrorMsg: err.Error()}
	}

	return result
}
