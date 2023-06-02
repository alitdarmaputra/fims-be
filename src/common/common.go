package common

import (
	"time"

	"github.com/gin-gonic/gin"
)

type BasicResponse struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

type BasicData struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type Meta struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type PageData struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
	Meta      Meta        `json:"meta"`
}

type ErrorResponse struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Message   interface{} `json:"message"`
	Timestamp int64       `json:"timestamp"`
}

type PathParam struct {
	Id uint `uri:"id" binding:"required"`
}

func JsonBasicResponse(ctx *gin.Context, code int, status string) {
	ctx.JSON(
		code,
		BasicResponse{
			Code:      code,
			Status:    status,
			Timestamp: time.Now().UnixNano(),
		},
	)
}

func JsonBasicData(ctx *gin.Context, code int, status string, data interface{}) {
	ctx.JSON(
		code,
		BasicData{
			Code:      code,
			Status:    status,
			Timestamp: time.Now().UnixNano(),
			Data:      data,
		},
	)
}

func JsonPageData(ctx *gin.Context, code int, status string, data interface{}, meta Meta) {
	ctx.JSON(
		code,
		PageData{
			Code:      code,
			Status:    status,
			Timestamp: time.Now().UnixNano(),
			Data:      data,
			Meta:      meta,
		},
	)
}

func JsonErrorResponse(ctx *gin.Context, code int, status string, message interface{}) {
	ctx.JSON(
		code,
		ErrorResponse{
			Code:      code,
			Status:    status,
			Message:   message,
			Timestamp: time.Now().UnixNano(),
		},
	)
}
