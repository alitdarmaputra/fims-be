package usecase

import (
	"github.com/alitdarmaputra/fims-be/src/business/domain"
	"github.com/alitdarmaputra/fims-be/src/business/usecase/smtp"
	"github.com/alitdarmaputra/fims-be/src/business/usecase/user"
	"github.com/alitdarmaputra/fims-be/src/config"
	"gorm.io/gorm"
)

type Usecase struct {
	User user.UserUsecase
}

func Init(
	dom *domain.Domain,
	db *gorm.DB,
	cfg *config.Api,
) *Usecase {
	smtpUsecase := smtp.InitSMTPUsecase(cfg.SMTP)
	userUsecase := user.InitUserUsecase(dom.User, smtpUsecase, dom.Token, db, cfg)

	return &Usecase{
		User: userUsecase,
	}
}
