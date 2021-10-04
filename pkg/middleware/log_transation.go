package middleware

import (
	"auth/pkg/enum"
	"auth/pkg/utils"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

//===============================================================//
//=========== ใช้สำหรับ Track Logs => request and respose =========//
//==============================================================//
type transationLogs struct {
	TransactionId string                 `json:"TransactionId,omitempty"`
	UserId        int64                  `json:"UserId,omitempty"`
	Level         string                 `json:"Level"`
	Method        string                 `json:"Method,omitempty"`
	HttpStatus    int                    `json:"HttpStatus,omitempty"`
	RequestURI    string                 `json:"RequestURI,omitempty"`
	Timestamp     time.Time              `json:"@timestamp"`
	Code          string                 `json:"Code,omitempty"`
	Request       map[string]interface{} `json:"Request"`
	Response      interface{}            `json:"Response"`
}

func Logger(c *fiber.Ctx) error {
	var tLog transationLogs
	var request map[string]interface{}
	c.BodyParser(&request)
	//set log
	tLog.Method = c.Method()
	tLog.RequestURI = c.OriginalURL()
	tLog.Request = request

	//Set InitDefaultUserInfo
	utils.InitDefaultUserInfo(c)

	err := c.Next()

	var response interface{}
	responseBody := c.Response().Body()
	json.Unmarshal([]byte(string(string(responseBody))), &response)
	statusCode := c.Response().Header.StatusCode()

	//set log
	tLog.Response = response
	tLog.HttpStatus = statusCode
	tLog.Timestamp = utils.GetTimeStampLog()
	userInfo := utils.GetUserInfo(c)
	tLog.TransactionId = userInfo.TransactionId
	tLog.UserId = userInfo.UserId

	switch statusCode {
	case fiber.StatusOK:
		tLog.Level = enum.LogLv_Info
	case fiber.StatusBadRequest,
		fiber.StatusUnauthorized:
		tLog.Level = enum.LogLv_Warn
	default:
		tLog.Level = enum.LogLv_Err
	}

	b, _ := json.Marshal(tLog)
	fmt.Println(string(b))
	return err
}
