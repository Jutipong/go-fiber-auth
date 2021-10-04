package model

import (
	"auth/pkg/enum"
	"auth/pkg/utils"
	"auth/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	CallbackRefid string `json:"callbackRefid"`
	SeqId         string `json:"seqId"`
	ChronicleId   string `json:"chronicleId"`
	Version       string `json:"version"`
	Action        string `json:"action"`
	AppOwner      string `json:"appOwner"`
	DocType       string `json:"docType"`
	PageCount     string `json:"pageCount"`
}

type Response struct {
	IsError  bool   `json:"isError"`
	ErrorMsg string `json:"errorMsg"`
	RefID    string `json:"refId"`
}

func (h *Request) Validate() error {
	var errs []string
	ruleId := validate(h, &errs)
	if len(errs) != 0 {
		json := utils.JsonSerialize(errs)
		return fiber.NewError(ruleId, json)
	} else {
		return nil
	}
}

func validate(req *Request, errs *[]string) int {

	//Rule 1 => Required
	ruleId := validation.Required(&[]validation.RequiredRule{
		{FieldName: "CallbackRefid", Value: req.CallbackRefid},
		{FieldName: "SeqId", Value: req.SeqId},
		{FieldName: "ChronicleId", Value: req.ChronicleId},
		{FieldName: "Version", Value: req.Version},
		{FieldName: "Action", Value: req.Action},
		{FieldName: "AppOwner", Value: req.AppOwner},
		{FieldName: "DocType", Value: req.DocType},
		{FieldName: "PageCount", Value: req.PageCount},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	return enum.Ok
}
