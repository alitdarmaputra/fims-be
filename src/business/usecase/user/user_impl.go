package user

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/alitdarmaputra/fims-be/src/business/domain/token"
	"github.com/alitdarmaputra/fims-be/src/business/domain/user"
	"github.com/alitdarmaputra/fims-be/src/business/entity"
	"github.com/alitdarmaputra/fims-be/src/business/model"
	"github.com/alitdarmaputra/fims-be/src/business/usecase/smtp"
	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/config"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/request"
	"github.com/alitdarmaputra/fims-be/src/handler/rest/response"
	"github.com/alitdarmaputra/fims-be/utils"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	defaultSecretKey  = "default-secret-key"
	defaultJWTExpired = 24 * time.Hour
)

type UserUsecaseImpl struct {
	UserDom      user.UserDom
	TokenDom     token.TokenDom
	SmtpUsecase  smtp.SmtpUsecase
	DB           *gorm.DB
	jwtSecretKey string
	jwtExpired   time.Duration
	cfg          *config.Api
}

func InitUserUsecase(
	userDom user.UserDom,
	smtpUsecase smtp.SmtpUsecase,
	tokenDom token.TokenDom,
	db *gorm.DB,
	cfg *config.Api,
) UserUsecase {
	return &UserUsecaseImpl{
		UserDom:      userDom,
		TokenDom:     tokenDom,
		SmtpUsecase:  smtpUsecase,
		DB:           db,
		jwtSecretKey: defaultSecretKey,
		jwtExpired:   defaultJWTExpired,
		cfg:          cfg,
	}
}

func (usecase *UserUsecaseImpl) SetJWTConfig(secret string, expired time.Duration) {
	usecase.jwtSecretKey = secret
	usecase.jwtExpired = expired
}

func (usecase *UserUsecaseImpl) Create(
	c context.Context,
	request request.HTTPUserCreateRequest,
) {
	var user model.User

	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	// Check if user has register but not verified

	user, err := usecase.UserDom.FindOne(c, tx, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}

	if user.VerifiedAt.Valid {
		panic(gorm.ErrDuplicatedKey)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	utils.PanicIfError(err)

	user.Email = request.Email
	user.Password = string(hash)
	user.Name = request.Name
	user.ProfileImg = "https://sbcf.fr/wp-content/uploads/2018/03/sbcf-default-avatar.png"

	user, err = usecase.UserDom.Save(c, tx, user)
	utils.PanicIfError(err)

	token, err := usecase.GenerateToken(user)
	utils.PanicIfError(err)

	data := entity.EmailData{
		Name:    user.Name,
		URL:     usecase.cfg.SMTP.ClientOrigin + "/register/verification/" + token,
		Subject: "Verifikasi Email Fims",
	}

	err = usecase.SmtpUsecase.SendMail(&user, &data)
	utils.PanicIfError(err)
}

func (usecase *UserUsecaseImpl) FindById(
	c context.Context,
	userId uint,
) response.HTTPUserDetailResponse {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := usecase.UserDom.FindById(c, tx, userId)
	utils.PanicIfError(err)

	return response.ToUserResponse(user)
}

func (usecase *UserUsecaseImpl) Login(
	c context.Context,
	request request.HTTPUserLoginRequest,
) *Token {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := usecase.UserDom.FindOne(c, tx, request.Email)
	if !user.VerifiedAt.Valid {
		panic(entity.NewUnauthorizedError("Incorrect email and password entered"))
	}

	if err != nil {
		panic(entity.NewUnauthorizedError("Incorrect email and password entered"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		panic(entity.NewUnauthorizedError("Incorrect email and password entered"))
	}

	token, err := usecase.GenerateToken(user)
	utils.PanicIfError(err)

	return &Token{
		Token: token,
	}
}

func (usecase *UserUsecaseImpl) GenerateToken(user model.User) (string, error) {
	eJWT := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  user.ID,
			"exp": time.Now().Add(usecase.jwtExpired).Unix(),
		},
	)
	return eJWT.SignedString([]byte(usecase.jwtSecretKey))
}

func (usecase *UserUsecaseImpl) ChangePassword(
	c context.Context,
	request request.HTTPChangePasswordRequest,
	userId uint,
) {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := usecase.UserDom.FindById(c, tx, userId)
	utils.PanicIfError(err)

	hash, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.MinCost)
	utils.PanicIfError(err)

	user.Password = string(hash)

	_, err = usecase.UserDom.Update(c, tx, user)
	utils.PanicIfError(err)
}

func (usecase *UserUsecaseImpl) VerifyEmail(
	c context.Context,
	verificationCode string,
) {
	// Decode token
	if verificationCode == "" {
		panic(entity.NewBadRequestError("Invalid URL"))
	}

	token, err := jwt.Parse(verificationCode, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok ||
			method != jwt.SigningMethodHS256 {
			return nil, entity.NewUnauthorizedError("Signing method invalid")
		}

		return []byte(usecase.jwtSecretKey), nil
	})

	if err != nil {
		panic(entity.NewUnauthorizedError(err.Error()))
	}

	claims := token.Claims.(jwt.MapClaims)
	res := new(entity.Token)
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(&claims)
	json.NewDecoder(buff).Decode(&res)

	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := usecase.UserDom.FindUnverifiedById(c, tx, res.Id)
	utils.PanicIfError(err)

	user.VerifiedAt = sql.NullTime{Time: time.Now(), Valid: true}

	_, err = usecase.UserDom.Update(c, tx, user)
	utils.PanicIfError(err)
}

func (usecase *UserUsecaseImpl) SendResetToken(
	c context.Context,
	request request.HTTPResetTokenRequest,
) {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	// check if user exist
	user, err := usecase.UserDom.FindOne(c, tx, request.Email)
	utils.PanicIfError(err)

	// generate reset token
	token := utils.RandStringRunes(30)

	//store token in database
	resetToken := model.Token{
		UserId:      user.ID,
		Token:       token,
		TokenExpiry: time.Now().Add(time.Minute * time.Duration(usecase.cfg.ResetTokenExpiredTime)),
	}

	usecase.TokenDom.Save(c, tx, resetToken)

	// send email with password reset link
	data := entity.EmailData{
		Name:    user.Name,
		URL:     usecase.cfg.SMTP.ClientOrigin + "/reset-password/" + resetToken.Token,
		Subject: "Reset Password Fims",
	}
	if err := usecase.SmtpUsecase.SendResetToken(&user, &data); err != nil {
		panic(err)
	}
}

func (usecase *UserUsecaseImpl) RedeemToken(
	c context.Context,
	request request.HTTPRedeemTokenRequest,
) {
	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	token, err := usecase.TokenDom.FindByToken(c, tx, request.Token)
	utils.PanicIfError(err)

	if time.Now().After(token.TokenExpiry) {
		panic(entity.NewUnauthorizedError("Invalid token"))
	}

	user, err := usecase.UserDom.FindById(c, tx, token.UserId)
	utils.PanicIfError(err)

	usecase.TokenDom.DeleteAllByUserId(c, tx, user.ID)

	hash, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.MinCost)
	utils.PanicIfError(err)

	user.Password = string(hash)

	if _, err = usecase.UserDom.Update(c, tx, user); err != nil {
		panic(err)
	}
}

func (usecase *UserUsecaseImpl) FindAll(
	c context.Context,
	page, perPage int,
	querySearch string,
) ([]response.HTTPUserDetailResponse, common.Meta) {
	offset := utils.CountOffset(page, perPage)

	tx := usecase.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	users, count := usecase.UserDom.FindAll(c, tx, offset, perPage, querySearch)

	meta := common.Meta{
		Page:      page,
		PerPage:   perPage,
		Total:     count,
		TotalPage: utils.CountTotalPage(count, perPage),
	}
	return response.ToUserResponses(users), meta
}
