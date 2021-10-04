package utils

import (
	"auth/pkg/enum"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

//==============================================================//
//================= ใช้สำหรับ Track logs => ทั่วไป =================//
//=============================================================//
type logs struct {
	TransactionId string    `json:"TransactionId,omitempty"`
	UserId        int64     `json:"UserId,omitempty"`
	Level         string    `json:"Level" default:"INFO"` //เช่น info, err, warn
	Timestamp     time.Time `json:"@timestamp"`
	// Code          string    `json:"Code,omitempty"`
	Message string `json:"Message"`
}

func LogInfo(Msg string) {
	fmt.Println(JsonSerialize(logs{
		Level:     enum.LogLv_Info,
		Timestamp: GetTimeStampLog(),
		Message:   Msg,
	}))
}
func LogInfoCtx(c *fiber.Ctx, Msg string) {
	u := GetUserInfo(c)
	fmt.Println(JsonSerialize(logs{
		TransactionId: u.TransactionId,
		UserId:        u.UserId,
		Level:         enum.LogLv_Info,
		Timestamp:     GetTimeStampLog(),
		Message:       Msg,
	}))
}

func LogErr(Msg string) {
	fmt.Println(JsonSerialize(logs{
		Level:     enum.LogLv_Err,
		Timestamp: GetTimeStampLog(),
		Message:   Msg,
	}))
}
func LogErrCtx(c *fiber.Ctx, Msg string) {
	u := GetUserInfo(c)
	fmt.Println(JsonSerialize(logs{
		TransactionId: u.TransactionId,
		UserId:        u.UserId,
		Level:         enum.LogLv_Err,
		Timestamp:     GetTimeStampLog(),
		Message:       Msg,
	}))
}

func LogWarn(Msg string) {
	fmt.Println(JsonSerialize(logs{
		Level:     enum.LogLv_Warn,
		Timestamp: GetTimeStampLog(),
		Message:   Msg,
	}))
}
func LogWarnCtx(c *fiber.Ctx, Msg string) {
	u := GetUserInfo(c)
	fmt.Println(JsonSerialize(logs{
		TransactionId: u.TransactionId,
		UserId:        u.UserId,
		Level:         enum.LogLv_Warn,
		Timestamp:     GetTimeStampLog(),
		Message:       Msg,
	}))
}

func GetTimeStampLog() time.Time {
	dataTime, err := time.Parse(time.RFC3339Nano, time.Now().UTC().Format(time.RFC3339Nano))
	if err != nil {
		panic(fmt.Sprintf("GetTimeStampLog: %v", err))
	}
	return dataTime
}
