package model

import (
	"database/sql"
	"time"
)

//======================= Model Database =======================
type TmpDocument struct {
	DocRefNo          string       `gorm:"primaryKey; column:DocRefNo"`
	DocumentID        string       `gorm:"column:DocumentID"`
	CustomerProfileID string       `gorm:"column:CustomerProfileID"`
	DocTypeID         int64        `gorm:"column:DocTypeID"`
	DocNo             string       `gorm:"column:DocNo"`
	FileName          string       `gorm:"column:FileName"`
	StartDate         sql.NullTime `gorm:"column:StartDate"`
	Enddate           sql.NullTime `gorm:"column:Enddate"`
	CreatedDate       time.Time    `gorm:"column:CreatedDate"`
	CreatedBy         string       `gorm:"column:CreatedBy"`
	UpdatedDate       time.Time    `gorm:"autoUpdateTime; column:UpdatedDate"`
	UpdatedBy         string       `gorm:"column:UpdatedBy"`
	ChronicleID       string       `gorm:"column:ChronicleID"`
	SeqID             string       `gorm:"column:SeqID"`
	Version           string       `gorm:"column:Version"`
	FileSize          int64        `gorm:"column:FileSize"`
	Status            string       `gorm:"column:Status"`
	UploadStatus      string       `gorm:"column:UploadStatus"`
	OrderNo           string       `gorm:"column:OrderNo"`
	DocumentGUID      string       `gorm:"column:DocumentGUID"`
	DocFullName       string       `gorm:"column:DocFullName"`
}

func (c *TmpDocument) TableName() string {
	return "tmpDocument"
}

type DocumentVersion struct {
	DocRefNo       string    `gorm:"primaryKey; column:DocRefNo"`
	DocumentID     string    `gorm:"column:DocumentID"`
	LastVersion    string    `gorm:"column:LastVersion"`
	ApproveVersion string    `gorm:"column:ApproveVersion"`
	CreatedDate    time.Time `gorm:"column:CreatedDate"`
	CreatedBy      string    `gorm:"column:CreatedBy"`
	UpdatedDate    time.Time `gorm:"autoUpdateTime; column:UpdatedDate"`
	UpdatedBy      string    `gorm:"column:UpdatedBy"`
	DocFullName    string    `gorm:"column:DocFullName"`
}

func (c *DocumentVersion) TableName() string {
	return "DocumentVersion"
}

type DocumentVersionlog struct {
	DocRefNo             string         `gorm:"column:DocRefNo"`
	DocumentVersionlogID int64          `gorm:"primaryKey; column:DocumentVersionlogID"`
	DocumentID           string         `gorm:"column:DocumentID"`
	DocTypeID            int64          `gorm:"column:DocTypeID"`
	DocNo                string         `gorm:"column:DocNo"`
	StartDate            sql.NullTime   `gorm:"column:StartDate"`
	Enddate              sql.NullTime   `gorm:"column:Enddate"`
	FileName             string         `gorm:"column:FileName"`
	ChronicleID          string         `gorm:"column:ChronicleID"`
	SeqID                string         `gorm:"column:SeqID"`
	Version              string         `gorm:"column:Version"`
	FileSize             int64          `gorm:"column:FileSize"`
	Status               string         `gorm:"column:Status"`
	UploadStatus         string         `gorm:"column:UploadStatus"`
	ApproveStatus        sql.NullString `gorm:"column:ApproveStatus"`
	CreatedDate          time.Time      `gorm:"autoCreateTime; column:CreatedDate"`
	CreatedBy            string         `gorm:"column:CreatedBy"`
	UpdatedDate          sql.NullTime   `gorm:"autoUpdateTime; column:UpdatedDate"`
	UpdatedBy            string         `gorm:"column:UpdatedBy"`
	DocFullName          string         `gorm:"column:DocFullName"`
}

func (c *DocumentVersionlog) TableName() string {
	return "DocumentVersionlog"
}

// type ServiceLog struct {
// 	Id              string         `gorm:"column:ID"`
// 	UserID          string         `gorm:"column:UserID"`
// 	MenuNameID      string         `gorm:"column:MenuNameID"`
// 	ServiceName     string         `gorm:"column:ServiceName"`
// 	RequestMessage  string         `gorm:"column:RequestMessage"`
// 	ResponseMessage string         `gorm:"column:ResponseMessage"`
// 	ResultMessage   sql.NullString `gorm:"column:ResultMessage"`
// 	PCode           sql.NullString `gorm:"column:PCode"`
// 	RequestDate     time.Time      `gorm:"column:RequestDate"`
// 	ResponseDate    time.Time      `gorm:"column:ResponseDate"`
// }

// func (c *ServiceLog) TableName() string {
// 	return "ServiceLog"
// }
