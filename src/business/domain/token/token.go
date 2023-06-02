package token

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type TokenDom interface {
	Save(c context.Context, tx *gorm.DB, resetToken model.Token)
	FindByToken(c context.Context, tx *gorm.DB, token string) (model.Token, error)
	DeleteAllByUserId(c context.Context, tx *gorm.DB, userId uint)
}
