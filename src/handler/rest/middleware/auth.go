package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/alitdarmaputra/fims-be/src/business/entity"
	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Token struct {
	Id     uint `json:"id"`
	Exp    int  `json:"exp"`
	RoleId uint `json:"role"`
}

type Authetication interface {
	ExtractJWTUser(ctx *gin.Context) (*Token, error)
}

type AutheticationImpl struct {
	secretKey string
}

func NewAuthentication(secretKey string) Authetication {
	return &AutheticationImpl{
		secretKey: secretKey,
	}
}

func (authentication *AutheticationImpl) ExtractJWTUser(ctx *gin.Context) (*Token, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return nil, entity.NewUnauthorizedError("User token not found")
	}

	if _, ok := user.(*jwt.Token); !ok {
		return nil, entity.NewUnauthorizedError("User token not found")
	}

	claims := user.(*jwt.Token).Claims.(jwt.MapClaims)

	res := new(Token)
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(&claims)
	json.NewDecoder(buff).Decode(res)

	return res, nil
}

func JWTMiddlewareAuth(jwtSecretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := strings.Replace(
			ctx.GetHeader("Authorization"),
			"Bearer ",
			"",
			1,
		)

		if token = strings.TrimSpace(token); token == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				common.ErrorResponse{
					Code:      http.StatusUnauthorized,
					Status:    "Unauthorized",
					Message:   "Invalid Token",
					Timestamp: time.Now().Unix(),
				},
			)
			return
		}

		res, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok ||
				method != jwt.SigningMethodHS256 {
				return nil, errors.New("signing method invalid")
			}

			return []byte(jwtSecretKey), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				common.ErrorResponse{
					Code:      http.StatusUnauthorized,
					Status:    "Unauthorized",
					Message:   err.Error(),
					Timestamp: time.Now().Unix(),
				},
			)
			return
		}

		ctx.Set("user", res)
		ctx.Next()
	}
}
