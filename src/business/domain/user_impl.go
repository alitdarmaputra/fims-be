package domain

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/entity"
	"gorm.io/gorm"
)

type UserDom struct{}

func InitUserDom() User {
	return &UserDom{}
}

func (domain *UserDom) Create(
	ctx context.Context,
	tx *gorm.DB,
	user entity.User,
) (entity.User, error) {
	if err := tx.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (domain *UserDom) FindOne(
	ctx context.Context,
	tx *gorm.DB,
	email string,
) (entity.User, error) {
	user := entity.User{}
	if err := tx.First(&user, "email = ?", email).Error; err != nil {
		return user, err
	}

	return user, nil
}
