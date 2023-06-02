package token

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type TokenDomImpl struct {
}

func InitTokenDom() TokenDom {
	return &TokenDomImpl{}
}

func (repository *TokenDomImpl) Save(
	ctx context.Context,
	tx *gorm.DB,
	resetToken model.Token,
) {
	_ = tx.Create(&resetToken)
}

func (repository *TokenDomImpl) FindByToken(
	ctx context.Context,
	tx *gorm.DB,
	token string,
) (model.Token, error) {
	var resetToken model.Token

	if err := tx.First(&resetToken, "token = ?", token).Error; err != nil {
		return resetToken, err
	}

	return resetToken, nil
}

func (repository *TokenDomImpl) DeleteAllByUserId(
	ctx context.Context,
	tx *gorm.DB,
	userId uint,
) {
	tx.Delete(&[]model.Token{}, "user_id", userId)
}
