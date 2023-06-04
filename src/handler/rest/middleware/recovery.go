package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alitdarmaputra/fims-be/src/business/entity"
	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorHandler(ctx *gin.Context, err any) {
	if notFoundError(ctx, err) {
		return
	}

	if validationError(ctx, err) {
		return
	}

	if unauthorizedError(ctx, err) {
		return
	}

	if duplicateEntryError(ctx, err) {
		return
	}

	if badGateWayError(ctx, err) {
		return
	}

	if badRequestError(ctx, err) {
		return
	}

	internalServerError(ctx, err)
}

func notFoundError(ctx *gin.Context, err any) bool {
	exception, ok := err.(error)
	if ok && errors.Is(exception, gorm.ErrRecordNotFound) {
		common.JsonErrorResponse(ctx, http.StatusNotFound, "Not found", exception.Error())
		return true
	}

	exception, ok = err.(*entity.NotFoundError)
	if ok {
		common.JsonErrorResponse(ctx, http.StatusNotFound, "Not found", exception.Error())
		return true
	}

	return false
}

func validationError(ctx *gin.Context, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		var messages []string
		for _, fieldErr := range exception {
			messages = append(
				messages,
				fmt.Sprintf(
					"Validation error for field %s on tag %s",
					fieldErr.Field(),
					fieldErr.Tag(),
				),
			)
		}
		common.JsonErrorResponse(ctx, http.StatusBadRequest, "Bad request", messages)
		return true
	} else {
		return false
	}
}

func unauthorizedError(ctx *gin.Context, err any) bool {
	exception, ok := err.(*entity.UnauthorizedError)
	if ok {
		common.JsonErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", exception.Error())
		return true
	} else {
		return false
	}
}

func internalServerError(ctx *gin.Context, err any) {
	// TODO: Custom logger
	common.JsonErrorResponse(
		ctx,
		http.StatusInternalServerError,
		"Internal server error",
		"Internal server error",
	)
}

func duplicateEntryError(ctx *gin.Context, err any) bool {
	exception, ok := err.(*entity.DuplicateEntryError)
	if ok {
		common.JsonErrorResponse(ctx, http.StatusConflict, "Conflict", exception.Error())
		return true
	} else {
		return false
	}
}

func badGateWayError(ctx *gin.Context, err any) bool {
	exception, ok := err.(*entity.BadGateWayError)
	if ok {
		common.JsonErrorResponse(ctx, http.StatusBadGateway, "Bad Gateway", exception.Error())
		return true
	} else {
		return false
	}
}

func badRequestError(ctx *gin.Context, err any) bool {
	exception, ok := err.(*entity.BadRequestError)
	if ok {
		common.JsonErrorResponse(ctx, http.StatusBadRequest, "Bad Request", exception.Error())
		return true
	} else {
		return false
	}
}
