package smtp

import (
	"github.com/alitdarmaputra/fims-be/src/business/entity"
	"github.com/alitdarmaputra/fims-be/src/business/model"
)

type SmtpUsecase interface {
	SendMail(user *model.User, data *entity.EmailData) error
	SendResetToken(user *model.User, data *entity.EmailData) error
	SendUpdate(user *model.User, data *entity.EmailData) error
}
