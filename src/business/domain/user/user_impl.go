package user

import (
	"context"

	"github.com/alitdarmaputra/fims-be/src/business/model"
	"gorm.io/gorm"
)

type UserDomImpl struct{}

func InitUserDom() UserDom {
	return &UserDomImpl{}
}

func (domain *UserDomImpl) Create(
	c context.Context,
	tx *gorm.DB,
	user model.User,
) (model.User, error) {
	if err := tx.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (domain *UserDomImpl) Save(
	c context.Context,
	tx *gorm.DB,
	user model.User,
) (model.User, error) {
	if err := tx.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (domain *UserDomImpl) FindOne(
	c context.Context,
	tx *gorm.DB,
	email string,
) (model.User, error) {
	user := model.User{}
	if err := tx.First(&user, "email = ?", email).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (domain *UserDomImpl) FindById(
	c context.Context,
	tx *gorm.DB,
	user_id uint,
) (model.User, error) {
	user := model.User{}
	if err := tx.First(&user, user_id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (domain *UserDomImpl) Update(
	c context.Context,
	tx *gorm.DB,
	user model.User,
) (model.User, error) {
	if err := tx.Updates(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (domain *UserDomImpl) FindUnverifiedById(
	c context.Context,
	tx *gorm.DB,
	userId uint,
) (model.User, error) {
	var user model.User
	if err := tx.Where("verified_at IS NULL AND id = ?", userId).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
func (domain *UserDomImpl) FindAll(
	c context.Context,
	tx *gorm.DB,
	offset, limit int,
	search string,
) ([]model.User, int) {
	var users []model.User = []model.User{}
	result := tx.Find(&users)

	if search != "" {
		search = "%" + search + "%"
		result = tx.Limit(limit).
			Offset(offset).
			Where("name LIKE ?", search).
			Find(&users)
	}

	return users, int(result.RowsAffected)
}
